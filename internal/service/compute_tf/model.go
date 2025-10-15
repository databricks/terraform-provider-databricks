// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package compute_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AddInstanceProfile struct {
	// The AWS IAM role ARN of the role associated with the instance profile.
	// This field is required if your role name and instance profile name do not
	// match and you want to use the instance profile with [Databricks SQL
	// Serverless].
	//
	// Otherwise, this field is optional.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	IamRoleArn types.String `tfsdk:"iam_role_arn"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile types.Bool `tfsdk:"is_meta_instance_profile"`
	// By default, Databricks validates that it has sufficient permissions to
	// launch instances with the instance profile. This validation uses AWS
	// dry-run mode for the RunInstances API. If validation fails with an error
	// message that does not indicate an IAM related permission issue, (e.g.
	// “Your requested instance type is not supported in your requested
	// availability zone”), you can pass this flag to skip the validation and
	// forcibly add the instance profile.
	SkipValidation types.Bool `tfsdk:"skip_validation"`
}

func (to *AddInstanceProfile) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddInstanceProfile) {
}

func (to *AddInstanceProfile) SyncFieldsDuringRead(ctx context.Context, from AddInstanceProfile) {
}

func (m AddInstanceProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["iam_role_arn"] = attrs["iam_role_arn"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetRequired()
	attrs["is_meta_instance_profile"] = attrs["is_meta_instance_profile"].SetOptional()
	attrs["skip_validation"] = attrs["skip_validation"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddInstanceProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AddInstanceProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddInstanceProfile
// only implements ToObjectValue() and Type().
func (m AddInstanceProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"iam_role_arn":             m.IamRoleArn,
			"instance_profile_arn":     m.InstanceProfileArn,
			"is_meta_instance_profile": m.IsMetaInstanceProfile,
			"skip_validation":          m.SkipValidation,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AddInstanceProfile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"iam_role_arn":             types.StringType,
			"instance_profile_arn":     types.StringType,
			"is_meta_instance_profile": types.BoolType,
			"skip_validation":          types.BoolType,
		},
	}
}

type AddResponse struct {
}

func (to *AddResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddResponse) {
}

func (to *AddResponse) SyncFieldsDuringRead(ctx context.Context, from AddResponse) {
}

func (m AddResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AddResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddResponse
// only implements ToObjectValue() and Type().
func (m AddResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m AddResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A storage location in Adls Gen2
type Adlsgen2Info struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	Destination types.String `tfsdk:"destination"`
}

func (to *Adlsgen2Info) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Adlsgen2Info) {
}

func (to *Adlsgen2Info) SyncFieldsDuringRead(ctx context.Context, from Adlsgen2Info) {
}

func (m Adlsgen2Info) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination"] = attrs["destination"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Adlsgen2Info.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Adlsgen2Info) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Adlsgen2Info
// only implements ToObjectValue() and Type().
func (m Adlsgen2Info) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": m.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Adlsgen2Info) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type AutoScale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. Note that `max_workers` must be strictly greater than
	// `min_workers`.
	MaxWorkers types.Int64 `tfsdk:"max_workers"`
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers types.Int64 `tfsdk:"min_workers"`
}

func (to *AutoScale) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutoScale) {
}

func (to *AutoScale) SyncFieldsDuringRead(ctx context.Context, from AutoScale) {
}

func (m AutoScale) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["max_workers"] = attrs["max_workers"].SetOptional()
	attrs["min_workers"] = attrs["min_workers"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoScale.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AutoScale) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoScale
// only implements ToObjectValue() and Type().
func (m AutoScale) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_workers": m.MaxWorkers,
			"min_workers": m.MinWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AutoScale) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_workers": types.Int64Type,
			"min_workers": types.Int64Type,
		},
	}
}

// Attributes set during cluster creation which are related to Amazon Web
// Services.
type AwsAttributes struct {
	Availability types.String `tfsdk:"availability"`
	// The number of volumes launched for each instance. Users can choose up to
	// 10 volumes. This feature is only enabled for supported node types. Legacy
	// node types cannot specify custom EBS volumes. For node types with no
	// instance store, at least one EBS volume needs to be specified; otherwise,
	// cluster creation will fail.
	//
	// These EBS volumes will be mounted at `/ebs0`, `/ebs1`, and etc. Instance
	// store volumes will be mounted at `/local_disk0`, `/local_disk1`, and etc.
	//
	// If EBS volumes are attached, Databricks will configure Spark to use only
	// the EBS volumes for scratch storage because heterogenously sized scratch
	// devices can lead to inefficient disk utilization. If no EBS volumes are
	// attached, Databricks will configure Spark to use instance store volumes.
	//
	// Please note that if EBS volumes are specified, then the Spark
	// configuration `spark.local.dir` will be overridden.
	EbsVolumeCount types.Int64 `tfsdk:"ebs_volume_count"`
	// If using gp3 volumes, what IOPS to use for the disk. If this is not set,
	// the maximum performance of a gp2 volume with the same volume size will be
	// used.
	EbsVolumeIops types.Int64 `tfsdk:"ebs_volume_iops"`
	// The size of each EBS volume (in GiB) launched for each instance. For
	// general purpose SSD, this value must be within the range 100 - 4096. For
	// throughput optimized HDD, this value must be within the range 500 - 4096.
	EbsVolumeSize types.Int64 `tfsdk:"ebs_volume_size"`
	// If using gp3 volumes, what throughput to use for the disk. If this is not
	// set, the maximum performance of a gp2 volume with the same volume size
	// will be used.
	EbsVolumeThroughput types.Int64 `tfsdk:"ebs_volume_throughput"`
	// The type of EBS volumes that will be launched with this cluster.
	EbsVolumeType types.String `tfsdk:"ebs_volume_type"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. If this value is greater than 0, the cluster driver
	// node in particular will be placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand types.Int64 `tfsdk:"first_on_demand"`
	// Nodes for this cluster will only be placed on AWS instances with this
	// instance profile. If ommitted, nodes will be placed on instances without
	// an IAM instance profile. The instance profile must have previously been
	// added to the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// The bid price for AWS spot instances, as a percentage of the
	// corresponding instance type's on-demand price. For example, if this field
	// is set to 50, and the cluster needs a new `r3.xlarge` spot instance, then
	// the bid price is half of the price of on-demand `r3.xlarge` instances.
	// Similarly, if this field is set to 200, the bid price is twice the price
	// of on-demand `r3.xlarge` instances. If not specified, the default value
	// is 100. When spot instances are requested for this cluster, only spot
	// instances whose bid price percentage matches this field will be
	// considered. Note that, for safety, we enforce this field to be no more
	// than 10000.
	SpotBidPricePercent types.Int64 `tfsdk:"spot_bid_price_percent"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. If the zone specified is "auto", will try to place cluster
	// in a zone with high availability, and will retry placement in a different
	// AZ if there is not enough capacity.
	//
	// The list of available zones as well as the default value can be found by
	// using the `List Zones` method.
	ZoneId types.String `tfsdk:"zone_id"`
}

func (to *AwsAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsAttributes) {
}

func (to *AwsAttributes) SyncFieldsDuringRead(ctx context.Context, from AwsAttributes) {
}

func (m AwsAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability"] = attrs["availability"].SetOptional()
	attrs["ebs_volume_count"] = attrs["ebs_volume_count"].SetOptional()
	attrs["ebs_volume_iops"] = attrs["ebs_volume_iops"].SetOptional()
	attrs["ebs_volume_size"] = attrs["ebs_volume_size"].SetOptional()
	attrs["ebs_volume_throughput"] = attrs["ebs_volume_throughput"].SetOptional()
	attrs["ebs_volume_type"] = attrs["ebs_volume_type"].SetOptional()
	attrs["first_on_demand"] = attrs["first_on_demand"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["spot_bid_price_percent"] = attrs["spot_bid_price_percent"].SetOptional()
	attrs["zone_id"] = attrs["zone_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AwsAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsAttributes
// only implements ToObjectValue() and Type().
func (m AwsAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":           m.Availability,
			"ebs_volume_count":       m.EbsVolumeCount,
			"ebs_volume_iops":        m.EbsVolumeIops,
			"ebs_volume_size":        m.EbsVolumeSize,
			"ebs_volume_throughput":  m.EbsVolumeThroughput,
			"ebs_volume_type":        m.EbsVolumeType,
			"first_on_demand":        m.FirstOnDemand,
			"instance_profile_arn":   m.InstanceProfileArn,
			"spot_bid_price_percent": m.SpotBidPricePercent,
			"zone_id":                m.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AwsAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":           types.StringType,
			"ebs_volume_count":       types.Int64Type,
			"ebs_volume_iops":        types.Int64Type,
			"ebs_volume_size":        types.Int64Type,
			"ebs_volume_throughput":  types.Int64Type,
			"ebs_volume_type":        types.StringType,
			"first_on_demand":        types.Int64Type,
			"instance_profile_arn":   types.StringType,
			"spot_bid_price_percent": types.Int64Type,
			"zone_id":                types.StringType,
		},
	}
}

// Attributes set during cluster creation which are related to Microsoft Azure.
type AzureAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones. Note: If `first_on_demand` is zero, this
	// availability type will be used for the entire cluster.
	Availability types.String `tfsdk:"availability"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. This value should be greater than 0, to make sure
	// the cluster driver node is placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand types.Int64 `tfsdk:"first_on_demand"`
	// Defines values necessary to configure and run Azure Log Analytics agent
	LogAnalyticsInfo types.Object `tfsdk:"log_analytics_info"`
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	SpotBidMaxPrice types.Float64 `tfsdk:"spot_bid_max_price"`
}

func (to *AzureAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureAttributes) {
	if !from.LogAnalyticsInfo.IsNull() && !from.LogAnalyticsInfo.IsUnknown() {
		if toLogAnalyticsInfo, ok := to.GetLogAnalyticsInfo(ctx); ok {
			if fromLogAnalyticsInfo, ok := from.GetLogAnalyticsInfo(ctx); ok {
				// Recursively sync the fields of LogAnalyticsInfo
				toLogAnalyticsInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromLogAnalyticsInfo)
				to.SetLogAnalyticsInfo(ctx, toLogAnalyticsInfo)
			}
		}
	}
}

func (to *AzureAttributes) SyncFieldsDuringRead(ctx context.Context, from AzureAttributes) {
	if !from.LogAnalyticsInfo.IsNull() && !from.LogAnalyticsInfo.IsUnknown() {
		if toLogAnalyticsInfo, ok := to.GetLogAnalyticsInfo(ctx); ok {
			if fromLogAnalyticsInfo, ok := from.GetLogAnalyticsInfo(ctx); ok {
				toLogAnalyticsInfo.SyncFieldsDuringRead(ctx, fromLogAnalyticsInfo)
				to.SetLogAnalyticsInfo(ctx, toLogAnalyticsInfo)
			}
		}
	}
}

func (m AzureAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability"] = attrs["availability"].SetOptional()
	attrs["first_on_demand"] = attrs["first_on_demand"].SetOptional()
	attrs["log_analytics_info"] = attrs["log_analytics_info"].SetOptional()
	attrs["spot_bid_max_price"] = attrs["spot_bid_max_price"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AzureAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_analytics_info": reflect.TypeOf(LogAnalyticsInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureAttributes
// only implements ToObjectValue() and Type().
func (m AzureAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":       m.Availability,
			"first_on_demand":    m.FirstOnDemand,
			"log_analytics_info": m.LogAnalyticsInfo,
			"spot_bid_max_price": m.SpotBidMaxPrice,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AzureAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":       types.StringType,
			"first_on_demand":    types.Int64Type,
			"log_analytics_info": LogAnalyticsInfo{}.Type(ctx),
			"spot_bid_max_price": types.Float64Type,
		},
	}
}

// GetLogAnalyticsInfo returns the value of the LogAnalyticsInfo field in AzureAttributes as
// a LogAnalyticsInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *AzureAttributes) GetLogAnalyticsInfo(ctx context.Context) (LogAnalyticsInfo, bool) {
	var e LogAnalyticsInfo
	if m.LogAnalyticsInfo.IsNull() || m.LogAnalyticsInfo.IsUnknown() {
		return e, false
	}
	var v LogAnalyticsInfo
	d := m.LogAnalyticsInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogAnalyticsInfo sets the value of the LogAnalyticsInfo field in AzureAttributes.
func (m *AzureAttributes) SetLogAnalyticsInfo(ctx context.Context, v LogAnalyticsInfo) {
	vs := v.ToObjectValue(ctx)
	m.LogAnalyticsInfo = vs
}

type CancelCommand struct {
	ClusterId types.String `tfsdk:"cluster_id"`

	CommandId types.String `tfsdk:"command_id"`

	ContextId types.String `tfsdk:"context_id"`
}

func (to *CancelCommand) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelCommand) {
}

func (to *CancelCommand) SyncFieldsDuringRead(ctx context.Context, from CancelCommand) {
}

func (m CancelCommand) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["command_id"] = attrs["command_id"].SetOptional()
	attrs["context_id"] = attrs["context_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelCommand.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelCommand) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelCommand
// only implements ToObjectValue() and Type().
func (m CancelCommand) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
			"command_id": m.CommandId,
			"context_id": m.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelCommand) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"command_id": types.StringType,
			"context_id": types.StringType,
		},
	}
}

type CancelResponse struct {
}

func (to *CancelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelResponse) {
}

func (to *CancelResponse) SyncFieldsDuringRead(ctx context.Context, from CancelResponse) {
}

func (m CancelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelResponse
// only implements ToObjectValue() and Type().
func (m CancelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m CancelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ChangeClusterOwner struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	// New owner of the cluster_id after this RPC.
	OwnerUsername types.String `tfsdk:"owner_username"`
}

func (to *ChangeClusterOwner) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ChangeClusterOwner) {
}

func (to *ChangeClusterOwner) SyncFieldsDuringRead(ctx context.Context, from ChangeClusterOwner) {
}

func (m ChangeClusterOwner) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["owner_username"] = attrs["owner_username"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChangeClusterOwner.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ChangeClusterOwner) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChangeClusterOwner
// only implements ToObjectValue() and Type().
func (m ChangeClusterOwner) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":     m.ClusterId,
			"owner_username": m.OwnerUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ChangeClusterOwner) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":     types.StringType,
			"owner_username": types.StringType,
		},
	}
}

type ChangeClusterOwnerResponse struct {
}

func (to *ChangeClusterOwnerResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ChangeClusterOwnerResponse) {
}

func (to *ChangeClusterOwnerResponse) SyncFieldsDuringRead(ctx context.Context, from ChangeClusterOwnerResponse) {
}

func (m ChangeClusterOwnerResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChangeClusterOwnerResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ChangeClusterOwnerResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChangeClusterOwnerResponse
// only implements ToObjectValue() and Type().
func (m ChangeClusterOwnerResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ChangeClusterOwnerResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ClientsTypes struct {
	// With jobs set, the cluster can be used for jobs
	Jobs types.Bool `tfsdk:"jobs"`
	// With notebooks set, this cluster can be used for notebooks
	Notebooks types.Bool `tfsdk:"notebooks"`
}

func (to *ClientsTypes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClientsTypes) {
}

func (to *ClientsTypes) SyncFieldsDuringRead(ctx context.Context, from ClientsTypes) {
}

func (m ClientsTypes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["jobs"] = attrs["jobs"].SetOptional()
	attrs["notebooks"] = attrs["notebooks"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClientsTypes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClientsTypes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClientsTypes
// only implements ToObjectValue() and Type().
func (m ClientsTypes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"jobs":      m.Jobs,
			"notebooks": m.Notebooks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClientsTypes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"jobs":      types.BoolType,
			"notebooks": types.BoolType,
		},
	}
}

type CloneCluster struct {
	// The cluster that is being cloned.
	SourceClusterId types.String `tfsdk:"source_cluster_id"`
}

func (to *CloneCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CloneCluster) {
}

func (to *CloneCluster) SyncFieldsDuringRead(ctx context.Context, from CloneCluster) {
}

func (m CloneCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["source_cluster_id"] = attrs["source_cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloneCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CloneCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloneCluster
// only implements ToObjectValue() and Type().
func (m CloneCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"source_cluster_id": m.SourceClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CloneCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"source_cluster_id": types.StringType,
		},
	}
}

type CloudProviderNodeInfo struct {
	// Status as reported by the cloud provider
	Status types.List `tfsdk:"status"`
}

func (to *CloudProviderNodeInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CloudProviderNodeInfo) {
	if !from.Status.IsNull() && !from.Status.IsUnknown() && to.Status.IsNull() && len(from.Status.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Status, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Status = from.Status
	}
}

func (to *CloudProviderNodeInfo) SyncFieldsDuringRead(ctx context.Context, from CloudProviderNodeInfo) {
	if !from.Status.IsNull() && !from.Status.IsUnknown() && to.Status.IsNull() && len(from.Status.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Status, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Status = from.Status
	}
}

func (m CloudProviderNodeInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloudProviderNodeInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CloudProviderNodeInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"status": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudProviderNodeInfo
// only implements ToObjectValue() and Type().
func (m CloudProviderNodeInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CloudProviderNodeInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetStatus returns the value of the Status field in CloudProviderNodeInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CloudProviderNodeInfo) GetStatus(ctx context.Context) ([]types.String, bool) {
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in CloudProviderNodeInfo.
func (m *CloudProviderNodeInfo) SetStatus(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Status = types.ListValueMust(t, vs)
}

type ClusterAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *ClusterAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAccessControlRequest) {
}

func (to *ClusterAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from ClusterAccessControlRequest) {
}

func (m ClusterAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAccessControlRequest
// only implements ToObjectValue() and Type().
func (m ClusterAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ClusterAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *ClusterAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *ClusterAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from ClusterAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m ClusterAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ClusterPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAccessControlResponse
// only implements ToObjectValue() and Type().
func (m ClusterAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ClusterPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in ClusterAccessControlResponse as
// a slice of ClusterPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAccessControlResponse) GetAllPermissions(ctx context.Context) ([]ClusterPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ClusterPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ClusterAccessControlResponse.
func (m *ClusterAccessControlResponse) SetAllPermissions(ctx context.Context, v []ClusterPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

// Common set of attributes set during cluster creation. These attributes cannot
// be changed over the lifetime of a cluster.
type ClusterAttributes struct {
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.Object `tfsdk:"cluster_log_conf"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName types.String `tfsdk:"cluster_name"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags"`

	DataSecurityMode types.String `tfsdk:"data_security_mode"`
	// Custom docker image BYOC
	DockerImage types.Object `tfsdk:"docker_image"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode types.Bool `tfsdk:"is_single_node"`

	Kind types.String `tfsdk:"kind"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars types.Map `tfsdk:"spark_env_vars"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime types.Bool `tfsdk:"use_ml_runtime"`

	WorkloadType types.Object `tfsdk:"workload_type"`
}

func (to *ClusterAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAttributes) {
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				// Recursively sync the fields of ClusterLogConf
				toClusterLogConf.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				// Recursively sync the fields of DockerImage
				toDockerImage.SyncFieldsDuringCreateOrUpdate(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				// Recursively sync the fields of WorkloadType
				toWorkloadType.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (to *ClusterAttributes) SyncFieldsDuringRead(ctx context.Context, from ClusterAttributes) {
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				toClusterLogConf.SyncFieldsDuringRead(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				toDockerImage.SyncFieldsDuringRead(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				toWorkloadType.SyncFieldsDuringRead(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (m ClusterAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["is_single_node"] = attrs["is_single_node"].SetOptional()
	attrs["kind"] = attrs["kind"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["runtime_engine"] = attrs["runtime_engine"].SetOptional()
	attrs["single_user_name"] = attrs["single_user_name"].SetOptional()
	attrs["spark_conf"] = attrs["spark_conf"].SetOptional()
	attrs["spark_env_vars"] = attrs["spark_env_vars"].SetOptional()
	attrs["spark_version"] = attrs["spark_version"].SetRequired()
	attrs["ssh_public_keys"] = attrs["ssh_public_keys"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()
	attrs["use_ml_runtime"] = attrs["use_ml_runtime"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":   reflect.TypeOf(AwsAttributes{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAttributes
// only implements ToObjectValue() and Type().
func (m ClusterAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autotermination_minutes":        m.AutoterminationMinutes,
			"aws_attributes":                 m.AwsAttributes,
			"azure_attributes":               m.AzureAttributes,
			"cluster_log_conf":               m.ClusterLogConf,
			"cluster_name":                   m.ClusterName,
			"custom_tags":                    m.CustomTags,
			"data_security_mode":             m.DataSecurityMode,
			"docker_image":                   m.DockerImage,
			"driver_instance_pool_id":        m.DriverInstancePoolId,
			"driver_node_type_id":            m.DriverNodeTypeId,
			"enable_elastic_disk":            m.EnableElasticDisk,
			"enable_local_disk_encryption":   m.EnableLocalDiskEncryption,
			"gcp_attributes":                 m.GcpAttributes,
			"init_scripts":                   m.InitScripts,
			"instance_pool_id":               m.InstancePoolId,
			"is_single_node":                 m.IsSingleNode,
			"kind":                           m.Kind,
			"node_type_id":                   m.NodeTypeId,
			"policy_id":                      m.PolicyId,
			"remote_disk_throughput":         m.RemoteDiskThroughput,
			"runtime_engine":                 m.RuntimeEngine,
			"single_user_name":               m.SingleUserName,
			"spark_conf":                     m.SparkConf,
			"spark_env_vars":                 m.SparkEnvVars,
			"spark_version":                  m.SparkVersion,
			"ssh_public_keys":                m.SshPublicKeys,
			"total_initial_remote_disk_size": m.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 m.UseMlRuntime,
			"workload_type":                  m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autotermination_minutes": types.Int64Type,
			"aws_attributes":          AwsAttributes{}.Type(ctx),
			"azure_attributes":        AzureAttributes{}.Type(ctx),
			"cluster_log_conf":        ClusterLogConf{}.Type(ctx),
			"cluster_name":            types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode":           types.StringType,
			"docker_image":                 DockerImage{}.Type(ctx),
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes":               GcpAttributes{}.Type(ctx),
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id":       types.StringType,
			"is_single_node":         types.BoolType,
			"kind":                   types.StringType,
			"node_type_id":           types.StringType,
			"policy_id":              types.StringType,
			"remote_disk_throughput": types.Int64Type,
			"runtime_engine":         types.StringType,
			"single_user_name":       types.StringType,
			"spark_conf": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_env_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_version": types.StringType,
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"total_initial_remote_disk_size": types.Int64Type,
			"use_ml_runtime":                 types.BoolType,
			"workload_type":                  WorkloadType{}.Type(ctx),
		},
	}
}

// GetAwsAttributes returns the value of the AwsAttributes field in ClusterAttributes as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v AwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterAttributes.
func (m *ClusterAttributes) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterAttributes as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v AzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterAttributes.
func (m *ClusterAttributes) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterAttributes as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if m.ClusterLogConf.IsNull() || m.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v ClusterLogConf
	d := m.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterAttributes.
func (m *ClusterAttributes) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	m.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in ClusterAttributes as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in ClusterAttributes.
func (m *ClusterAttributes) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterAttributes as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if m.DockerImage.IsNull() || m.DockerImage.IsUnknown() {
		return e, false
	}
	var v DockerImage
	d := m.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDockerImage sets the value of the DockerImage field in ClusterAttributes.
func (m *ClusterAttributes) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	m.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterAttributes as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v GcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterAttributes.
func (m *ClusterAttributes) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in ClusterAttributes as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := m.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterAttributes.
func (m *ClusterAttributes) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterAttributes as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkConf.IsNull() || m.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in ClusterAttributes.
func (m *ClusterAttributes) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterAttributes as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkEnvVars.IsNull() || m.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterAttributes.
func (m *ClusterAttributes) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterAttributes as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if m.SshPublicKeys.IsNull() || m.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterAttributes.
func (m *ClusterAttributes) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterAttributes as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAttributes) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if m.WorkloadType.IsNull() || m.WorkloadType.IsUnknown() {
		return e, false
	}
	var v WorkloadType
	d := m.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterAttributes.
func (m *ClusterAttributes) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	m.WorkloadType = vs
}

type ClusterCompliance struct {
	// Canonical unique identifier for a cluster.
	ClusterId types.String `tfsdk:"cluster_id"`
	// Whether this cluster is in compliance with the latest version of its
	// policy.
	IsCompliant types.Bool `tfsdk:"is_compliant"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	Violations types.Map `tfsdk:"violations"`
}

func (to *ClusterCompliance) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterCompliance) {
}

func (to *ClusterCompliance) SyncFieldsDuringRead(ctx context.Context, from ClusterCompliance) {
}

func (m ClusterCompliance) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["is_compliant"] = attrs["is_compliant"].SetOptional()
	attrs["violations"] = attrs["violations"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterCompliance.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterCompliance) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterCompliance
// only implements ToObjectValue() and Type().
func (m ClusterCompliance) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":   m.ClusterId,
			"is_compliant": m.IsCompliant,
			"violations":   m.Violations,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterCompliance) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":   types.StringType,
			"is_compliant": types.BoolType,
			"violations": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetViolations returns the value of the Violations field in ClusterCompliance as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterCompliance) GetViolations(ctx context.Context) (map[string]types.String, bool) {
	if m.Violations.IsNull() || m.Violations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Violations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViolations sets the value of the Violations field in ClusterCompliance.
func (m *ClusterCompliance) SetViolations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["violations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Violations = types.MapValueMust(t, vs)
}

// Describes all of the metadata about a single Spark cluster in Databricks.
type ClusterDetails struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// Number of CPU cores available for this cluster. Note that this can be
	// fractional, e.g. 7.5 cores, since certain node types are configured to
	// share cores between Spark nodes on the same instance.
	ClusterCores types.Float64 `tfsdk:"cluster_cores"`
	// Canonical identifier for the cluster. This id is retained during cluster
	// restarts and resizes, while each new cluster has a globally unique id.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.Object `tfsdk:"cluster_log_conf"`
	// Cluster log delivery status.
	ClusterLogStatus types.Object `tfsdk:"cluster_log_status"`
	// Total amount of cluster memory, in megabytes
	ClusterMemoryMb types.Int64 `tfsdk:"cluster_memory_mb"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName types.String `tfsdk:"cluster_name"`
	// Determines whether the cluster was created by a user through the UI,
	// created by the Databricks Jobs Scheduler, or through an API request.
	ClusterSource types.String `tfsdk:"cluster_source"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags"`

	DataSecurityMode types.String `tfsdk:"data_security_mode"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - Creator: <username_of_creator>
	//
	// - ClusterName: <name_of_cluster>
	//
	// - ClusterId: <id_of_cluster>
	//
	// - Name: <Databricks internal use>
	DefaultTags types.Map `tfsdk:"default_tags"`
	// Custom docker image BYOC
	DockerImage types.Object `tfsdk:"docker_image"`
	// Node on which the Spark driver resides. The driver node contains the
	// Spark master and the Databricks application that manages the per-notebook
	// Spark REPLs.
	Driver types.Object `tfsdk:"driver"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption"`
	// Nodes on which the Spark executors reside.
	Executors types.List `tfsdk:"executors"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode types.Bool `tfsdk:"is_single_node"`
	// Port on which Spark JDBC server is listening, in the driver nod. No
	// service will be listeningon on this port in executor nodes.
	JdbcPort types.Int64 `tfsdk:"jdbc_port"`

	Kind types.String `tfsdk:"kind"`
	// the timestamp that the cluster was started/restarted
	LastRestartedTime types.Int64 `tfsdk:"last_restarted_time"`
	// Time when the cluster driver last lost its state (due to a restart or
	// driver failure).
	LastStateLossTime types.Int64 `tfsdk:"last_state_loss_time"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf"`
	// A canonical SparkContext identifier. This value *does* change when the
	// Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
	// globally unique identifier over all Spark contexts.
	SparkContextId types.Int64 `tfsdk:"spark_context_id"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars types.Map `tfsdk:"spark_env_vars"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version"`
	// The spec contains a snapshot of the latest user specified settings that
	// were used to create/edit the cluster. Note: not included in the response
	// of the ListClusters API.
	Spec types.Object `tfsdk:"spec"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys"`
	// Time (in epoch milliseconds) when the cluster creation request was
	// received (when the cluster entered a `PENDING` state).
	StartTime types.Int64 `tfsdk:"start_time"`
	// Current state of the cluster.
	State types.String `tfsdk:"state"`
	// A message associated with the most recent state transition (e.g., the
	// reason why the cluster entered a `TERMINATED` state).
	StateMessage types.String `tfsdk:"state_message"`
	// Time (in epoch milliseconds) when the cluster was terminated, if
	// applicable.
	TerminatedTime types.Int64 `tfsdk:"terminated_time"`
	// Information about why the cluster was terminated. This field only appears
	// when the cluster is in a `TERMINATING` or `TERMINATED` state.
	TerminationReason types.Object `tfsdk:"termination_reason"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime types.Bool `tfsdk:"use_ml_runtime"`

	WorkloadType types.Object `tfsdk:"workload_type"`
}

func (to *ClusterDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterDetails) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				// Recursively sync the fields of Autoscale
				toAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				// Recursively sync the fields of ClusterLogConf
				toClusterLogConf.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.ClusterLogStatus.IsNull() && !from.ClusterLogStatus.IsUnknown() {
		if toClusterLogStatus, ok := to.GetClusterLogStatus(ctx); ok {
			if fromClusterLogStatus, ok := from.GetClusterLogStatus(ctx); ok {
				// Recursively sync the fields of ClusterLogStatus
				toClusterLogStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterLogStatus)
				to.SetClusterLogStatus(ctx, toClusterLogStatus)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				// Recursively sync the fields of DockerImage
				toDockerImage.SyncFieldsDuringCreateOrUpdate(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.Driver.IsNull() && !from.Driver.IsUnknown() {
		if toDriver, ok := to.GetDriver(ctx); ok {
			if fromDriver, ok := from.GetDriver(ctx); ok {
				// Recursively sync the fields of Driver
				toDriver.SyncFieldsDuringCreateOrUpdate(ctx, fromDriver)
				to.SetDriver(ctx, toDriver)
			}
		}
	}
	if !from.Executors.IsNull() && !from.Executors.IsUnknown() && to.Executors.IsNull() && len(from.Executors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Executors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Executors = from.Executors
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.TerminationReason.IsNull() && !from.TerminationReason.IsUnknown() {
		if toTerminationReason, ok := to.GetTerminationReason(ctx); ok {
			if fromTerminationReason, ok := from.GetTerminationReason(ctx); ok {
				// Recursively sync the fields of TerminationReason
				toTerminationReason.SyncFieldsDuringCreateOrUpdate(ctx, fromTerminationReason)
				to.SetTerminationReason(ctx, toTerminationReason)
			}
		}
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				// Recursively sync the fields of WorkloadType
				toWorkloadType.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (to *ClusterDetails) SyncFieldsDuringRead(ctx context.Context, from ClusterDetails) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				toAutoscale.SyncFieldsDuringRead(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				toClusterLogConf.SyncFieldsDuringRead(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.ClusterLogStatus.IsNull() && !from.ClusterLogStatus.IsUnknown() {
		if toClusterLogStatus, ok := to.GetClusterLogStatus(ctx); ok {
			if fromClusterLogStatus, ok := from.GetClusterLogStatus(ctx); ok {
				toClusterLogStatus.SyncFieldsDuringRead(ctx, fromClusterLogStatus)
				to.SetClusterLogStatus(ctx, toClusterLogStatus)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				toDockerImage.SyncFieldsDuringRead(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.Driver.IsNull() && !from.Driver.IsUnknown() {
		if toDriver, ok := to.GetDriver(ctx); ok {
			if fromDriver, ok := from.GetDriver(ctx); ok {
				toDriver.SyncFieldsDuringRead(ctx, fromDriver)
				to.SetDriver(ctx, toDriver)
			}
		}
	}
	if !from.Executors.IsNull() && !from.Executors.IsUnknown() && to.Executors.IsNull() && len(from.Executors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Executors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Executors = from.Executors
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.TerminationReason.IsNull() && !from.TerminationReason.IsUnknown() {
		if toTerminationReason, ok := to.GetTerminationReason(ctx); ok {
			if fromTerminationReason, ok := from.GetTerminationReason(ctx); ok {
				toTerminationReason.SyncFieldsDuringRead(ctx, fromTerminationReason)
				to.SetTerminationReason(ctx, toTerminationReason)
			}
		}
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				toWorkloadType.SyncFieldsDuringRead(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (m ClusterDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["cluster_cores"] = attrs["cluster_cores"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_log_status"] = attrs["cluster_log_status"].SetOptional()
	attrs["cluster_memory_mb"] = attrs["cluster_memory_mb"].SetOptional()
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["cluster_source"] = attrs["cluster_source"].SetOptional()
	attrs["creator_user_name"] = attrs["creator_user_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["default_tags"] = attrs["default_tags"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["driver"] = attrs["driver"].SetOptional()
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["executors"] = attrs["executors"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["is_single_node"] = attrs["is_single_node"].SetOptional()
	attrs["jdbc_port"] = attrs["jdbc_port"].SetOptional()
	attrs["kind"] = attrs["kind"].SetOptional()
	attrs["last_restarted_time"] = attrs["last_restarted_time"].SetOptional()
	attrs["last_state_loss_time"] = attrs["last_state_loss_time"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["num_workers"] = attrs["num_workers"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["runtime_engine"] = attrs["runtime_engine"].SetOptional()
	attrs["single_user_name"] = attrs["single_user_name"].SetOptional()
	attrs["spark_conf"] = attrs["spark_conf"].SetOptional()
	attrs["spark_context_id"] = attrs["spark_context_id"].SetOptional()
	attrs["spark_env_vars"] = attrs["spark_env_vars"].SetOptional()
	attrs["spark_version"] = attrs["spark_version"].SetOptional()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["ssh_public_keys"] = attrs["ssh_public_keys"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["state_message"] = attrs["state_message"].SetOptional()
	attrs["terminated_time"] = attrs["terminated_time"].SetOptional()
	attrs["termination_reason"] = attrs["termination_reason"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()
	attrs["use_ml_runtime"] = attrs["use_ml_runtime"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":          reflect.TypeOf(AutoScale{}),
		"aws_attributes":     reflect.TypeOf(AwsAttributes{}),
		"azure_attributes":   reflect.TypeOf(AzureAttributes{}),
		"cluster_log_conf":   reflect.TypeOf(ClusterLogConf{}),
		"cluster_log_status": reflect.TypeOf(LogSyncStatus{}),
		"custom_tags":        reflect.TypeOf(types.String{}),
		"default_tags":       reflect.TypeOf(types.String{}),
		"docker_image":       reflect.TypeOf(DockerImage{}),
		"driver":             reflect.TypeOf(SparkNode{}),
		"executors":          reflect.TypeOf(SparkNode{}),
		"gcp_attributes":     reflect.TypeOf(GcpAttributes{}),
		"init_scripts":       reflect.TypeOf(InitScriptInfo{}),
		"spark_conf":         reflect.TypeOf(types.String{}),
		"spark_env_vars":     reflect.TypeOf(types.String{}),
		"spec":               reflect.TypeOf(ClusterSpec{}),
		"ssh_public_keys":    reflect.TypeOf(types.String{}),
		"termination_reason": reflect.TypeOf(TerminationReason{}),
		"workload_type":      reflect.TypeOf(WorkloadType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterDetails
// only implements ToObjectValue() and Type().
func (m ClusterDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":                      m.Autoscale,
			"autotermination_minutes":        m.AutoterminationMinutes,
			"aws_attributes":                 m.AwsAttributes,
			"azure_attributes":               m.AzureAttributes,
			"cluster_cores":                  m.ClusterCores,
			"cluster_id":                     m.ClusterId,
			"cluster_log_conf":               m.ClusterLogConf,
			"cluster_log_status":             m.ClusterLogStatus,
			"cluster_memory_mb":              m.ClusterMemoryMb,
			"cluster_name":                   m.ClusterName,
			"cluster_source":                 m.ClusterSource,
			"creator_user_name":              m.CreatorUserName,
			"custom_tags":                    m.CustomTags,
			"data_security_mode":             m.DataSecurityMode,
			"default_tags":                   m.DefaultTags,
			"docker_image":                   m.DockerImage,
			"driver":                         m.Driver,
			"driver_instance_pool_id":        m.DriverInstancePoolId,
			"driver_node_type_id":            m.DriverNodeTypeId,
			"enable_elastic_disk":            m.EnableElasticDisk,
			"enable_local_disk_encryption":   m.EnableLocalDiskEncryption,
			"executors":                      m.Executors,
			"gcp_attributes":                 m.GcpAttributes,
			"init_scripts":                   m.InitScripts,
			"instance_pool_id":               m.InstancePoolId,
			"is_single_node":                 m.IsSingleNode,
			"jdbc_port":                      m.JdbcPort,
			"kind":                           m.Kind,
			"last_restarted_time":            m.LastRestartedTime,
			"last_state_loss_time":           m.LastStateLossTime,
			"node_type_id":                   m.NodeTypeId,
			"num_workers":                    m.NumWorkers,
			"policy_id":                      m.PolicyId,
			"remote_disk_throughput":         m.RemoteDiskThroughput,
			"runtime_engine":                 m.RuntimeEngine,
			"single_user_name":               m.SingleUserName,
			"spark_conf":                     m.SparkConf,
			"spark_context_id":               m.SparkContextId,
			"spark_env_vars":                 m.SparkEnvVars,
			"spark_version":                  m.SparkVersion,
			"spec":                           m.Spec,
			"ssh_public_keys":                m.SshPublicKeys,
			"start_time":                     m.StartTime,
			"state":                          m.State,
			"state_message":                  m.StateMessage,
			"terminated_time":                m.TerminatedTime,
			"termination_reason":             m.TerminationReason,
			"total_initial_remote_disk_size": m.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 m.UseMlRuntime,
			"workload_type":                  m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale":               AutoScale{}.Type(ctx),
			"autotermination_minutes": types.Int64Type,
			"aws_attributes":          AwsAttributes{}.Type(ctx),
			"azure_attributes":        AzureAttributes{}.Type(ctx),
			"cluster_cores":           types.Float64Type,
			"cluster_id":              types.StringType,
			"cluster_log_conf":        ClusterLogConf{}.Type(ctx),
			"cluster_log_status":      LogSyncStatus{}.Type(ctx),
			"cluster_memory_mb":       types.Int64Type,
			"cluster_name":            types.StringType,
			"cluster_source":          types.StringType,
			"creator_user_name":       types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"default_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"docker_image":                 DockerImage{}.Type(ctx),
			"driver":                       SparkNode{}.Type(ctx),
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"executors": basetypes.ListType{
				ElemType: SparkNode{}.Type(ctx),
			},
			"gcp_attributes": GcpAttributes{}.Type(ctx),
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id":       types.StringType,
			"is_single_node":         types.BoolType,
			"jdbc_port":              types.Int64Type,
			"kind":                   types.StringType,
			"last_restarted_time":    types.Int64Type,
			"last_state_loss_time":   types.Int64Type,
			"node_type_id":           types.StringType,
			"num_workers":            types.Int64Type,
			"policy_id":              types.StringType,
			"remote_disk_throughput": types.Int64Type,
			"runtime_engine":         types.StringType,
			"single_user_name":       types.StringType,
			"spark_conf": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_context_id": types.Int64Type,
			"spark_env_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_version": types.StringType,
			"spec":          ClusterSpec{}.Type(ctx),
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"start_time":                     types.Int64Type,
			"state":                          types.StringType,
			"state_message":                  types.StringType,
			"terminated_time":                types.Int64Type,
			"termination_reason":             TerminationReason{}.Type(ctx),
			"total_initial_remote_disk_size": types.Int64Type,
			"use_ml_runtime":                 types.BoolType,
			"workload_type":                  WorkloadType{}.Type(ctx),
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in ClusterDetails as
// a AutoScale value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v AutoScale
	d := m.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoscale sets the value of the Autoscale field in ClusterDetails.
func (m *ClusterDetails) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	m.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in ClusterDetails as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v AwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterDetails.
func (m *ClusterDetails) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterDetails as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v AzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterDetails.
func (m *ClusterDetails) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterDetails as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if m.ClusterLogConf.IsNull() || m.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v ClusterLogConf
	d := m.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterDetails.
func (m *ClusterDetails) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	m.ClusterLogConf = vs
}

// GetClusterLogStatus returns the value of the ClusterLogStatus field in ClusterDetails as
// a LogSyncStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetClusterLogStatus(ctx context.Context) (LogSyncStatus, bool) {
	var e LogSyncStatus
	if m.ClusterLogStatus.IsNull() || m.ClusterLogStatus.IsUnknown() {
		return e, false
	}
	var v LogSyncStatus
	d := m.ClusterLogStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterLogStatus sets the value of the ClusterLogStatus field in ClusterDetails.
func (m *ClusterDetails) SetClusterLogStatus(ctx context.Context, v LogSyncStatus) {
	vs := v.ToObjectValue(ctx)
	m.ClusterLogStatus = vs
}

// GetCustomTags returns the value of the CustomTags field in ClusterDetails as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in ClusterDetails.
func (m *ClusterDetails) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in ClusterDetails as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
	if m.DefaultTags.IsNull() || m.DefaultTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.DefaultTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultTags sets the value of the DefaultTags field in ClusterDetails.
func (m *ClusterDetails) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DefaultTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterDetails as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if m.DockerImage.IsNull() || m.DockerImage.IsUnknown() {
		return e, false
	}
	var v DockerImage
	d := m.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDockerImage sets the value of the DockerImage field in ClusterDetails.
func (m *ClusterDetails) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	m.DockerImage = vs
}

// GetDriver returns the value of the Driver field in ClusterDetails as
// a SparkNode value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetDriver(ctx context.Context) (SparkNode, bool) {
	var e SparkNode
	if m.Driver.IsNull() || m.Driver.IsUnknown() {
		return e, false
	}
	var v SparkNode
	d := m.Driver.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDriver sets the value of the Driver field in ClusterDetails.
func (m *ClusterDetails) SetDriver(ctx context.Context, v SparkNode) {
	vs := v.ToObjectValue(ctx)
	m.Driver = vs
}

// GetExecutors returns the value of the Executors field in ClusterDetails as
// a slice of SparkNode values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetExecutors(ctx context.Context) ([]SparkNode, bool) {
	if m.Executors.IsNull() || m.Executors.IsUnknown() {
		return nil, false
	}
	var v []SparkNode
	d := m.Executors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExecutors sets the value of the Executors field in ClusterDetails.
func (m *ClusterDetails) SetExecutors(ctx context.Context, v []SparkNode) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["executors"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Executors = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterDetails as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v GcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterDetails.
func (m *ClusterDetails) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in ClusterDetails as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := m.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterDetails.
func (m *ClusterDetails) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterDetails as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkConf.IsNull() || m.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in ClusterDetails.
func (m *ClusterDetails) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterDetails as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkEnvVars.IsNull() || m.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterDetails.
func (m *ClusterDetails) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSpec returns the value of the Spec field in ClusterDetails as
// a ClusterSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetSpec(ctx context.Context) (ClusterSpec, bool) {
	var e ClusterSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v ClusterSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in ClusterDetails.
func (m *ClusterDetails) SetSpec(ctx context.Context, v ClusterSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterDetails as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if m.SshPublicKeys.IsNull() || m.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterDetails.
func (m *ClusterDetails) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetTerminationReason returns the value of the TerminationReason field in ClusterDetails as
// a TerminationReason value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetTerminationReason(ctx context.Context) (TerminationReason, bool) {
	var e TerminationReason
	if m.TerminationReason.IsNull() || m.TerminationReason.IsUnknown() {
		return e, false
	}
	var v TerminationReason
	d := m.TerminationReason.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTerminationReason sets the value of the TerminationReason field in ClusterDetails.
func (m *ClusterDetails) SetTerminationReason(ctx context.Context, v TerminationReason) {
	vs := v.ToObjectValue(ctx)
	m.TerminationReason = vs
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterDetails as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterDetails) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if m.WorkloadType.IsNull() || m.WorkloadType.IsUnknown() {
		return e, false
	}
	var v WorkloadType
	d := m.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterDetails.
func (m *ClusterDetails) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	m.WorkloadType = vs
}

type ClusterEvent struct {
	ClusterId types.String `tfsdk:"cluster_id"`

	DataPlaneEventDetails types.Object `tfsdk:"data_plane_event_details"`

	Details types.Object `tfsdk:"details"`
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	Timestamp types.Int64 `tfsdk:"timestamp"`

	Type_ types.String `tfsdk:"type"`
}

func (to *ClusterEvent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterEvent) {
	if !from.DataPlaneEventDetails.IsNull() && !from.DataPlaneEventDetails.IsUnknown() {
		if toDataPlaneEventDetails, ok := to.GetDataPlaneEventDetails(ctx); ok {
			if fromDataPlaneEventDetails, ok := from.GetDataPlaneEventDetails(ctx); ok {
				// Recursively sync the fields of DataPlaneEventDetails
				toDataPlaneEventDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromDataPlaneEventDetails)
				to.SetDataPlaneEventDetails(ctx, toDataPlaneEventDetails)
			}
		}
	}
	if !from.Details.IsNull() && !from.Details.IsUnknown() {
		if toDetails, ok := to.GetDetails(ctx); ok {
			if fromDetails, ok := from.GetDetails(ctx); ok {
				// Recursively sync the fields of Details
				toDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromDetails)
				to.SetDetails(ctx, toDetails)
			}
		}
	}
}

func (to *ClusterEvent) SyncFieldsDuringRead(ctx context.Context, from ClusterEvent) {
	if !from.DataPlaneEventDetails.IsNull() && !from.DataPlaneEventDetails.IsUnknown() {
		if toDataPlaneEventDetails, ok := to.GetDataPlaneEventDetails(ctx); ok {
			if fromDataPlaneEventDetails, ok := from.GetDataPlaneEventDetails(ctx); ok {
				toDataPlaneEventDetails.SyncFieldsDuringRead(ctx, fromDataPlaneEventDetails)
				to.SetDataPlaneEventDetails(ctx, toDataPlaneEventDetails)
			}
		}
	}
	if !from.Details.IsNull() && !from.Details.IsUnknown() {
		if toDetails, ok := to.GetDetails(ctx); ok {
			if fromDetails, ok := from.GetDetails(ctx); ok {
				toDetails.SyncFieldsDuringRead(ctx, fromDetails)
				to.SetDetails(ctx, toDetails)
			}
		}
	}
}

func (m ClusterEvent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["data_plane_event_details"] = attrs["data_plane_event_details"].SetOptional()
	attrs["details"] = attrs["details"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterEvent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterEvent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_plane_event_details": reflect.TypeOf(DataPlaneEventDetails{}),
		"details":                  reflect.TypeOf(EventDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterEvent
// only implements ToObjectValue() and Type().
func (m ClusterEvent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":               m.ClusterId,
			"data_plane_event_details": m.DataPlaneEventDetails,
			"details":                  m.Details,
			"timestamp":                m.Timestamp,
			"type":                     m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterEvent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":               types.StringType,
			"data_plane_event_details": DataPlaneEventDetails{}.Type(ctx),
			"details":                  EventDetails{}.Type(ctx),
			"timestamp":                types.Int64Type,
			"type":                     types.StringType,
		},
	}
}

// GetDataPlaneEventDetails returns the value of the DataPlaneEventDetails field in ClusterEvent as
// a DataPlaneEventDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterEvent) GetDataPlaneEventDetails(ctx context.Context) (DataPlaneEventDetails, bool) {
	var e DataPlaneEventDetails
	if m.DataPlaneEventDetails.IsNull() || m.DataPlaneEventDetails.IsUnknown() {
		return e, false
	}
	var v DataPlaneEventDetails
	d := m.DataPlaneEventDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataPlaneEventDetails sets the value of the DataPlaneEventDetails field in ClusterEvent.
func (m *ClusterEvent) SetDataPlaneEventDetails(ctx context.Context, v DataPlaneEventDetails) {
	vs := v.ToObjectValue(ctx)
	m.DataPlaneEventDetails = vs
}

// GetDetails returns the value of the Details field in ClusterEvent as
// a EventDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterEvent) GetDetails(ctx context.Context) (EventDetails, bool) {
	var e EventDetails
	if m.Details.IsNull() || m.Details.IsUnknown() {
		return e, false
	}
	var v EventDetails
	d := m.Details.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDetails sets the value of the Details field in ClusterEvent.
func (m *ClusterEvent) SetDetails(ctx context.Context, v EventDetails) {
	vs := v.ToObjectValue(ctx)
	m.Details = vs
}

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	ClusterId types.String `tfsdk:"cluster_id"`
	// Status of all libraries on the cluster.
	LibraryStatuses types.List `tfsdk:"library_statuses"`
}

func (to *ClusterLibraryStatuses) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterLibraryStatuses) {
	if !from.LibraryStatuses.IsNull() && !from.LibraryStatuses.IsUnknown() && to.LibraryStatuses.IsNull() && len(from.LibraryStatuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LibraryStatuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LibraryStatuses = from.LibraryStatuses
	}
}

func (to *ClusterLibraryStatuses) SyncFieldsDuringRead(ctx context.Context, from ClusterLibraryStatuses) {
	if !from.LibraryStatuses.IsNull() && !from.LibraryStatuses.IsUnknown() && to.LibraryStatuses.IsNull() && len(from.LibraryStatuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LibraryStatuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LibraryStatuses = from.LibraryStatuses
	}
}

func (m ClusterLibraryStatuses) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["library_statuses"] = attrs["library_statuses"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterLibraryStatuses.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterLibraryStatuses) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"library_statuses": reflect.TypeOf(LibraryFullStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterLibraryStatuses
// only implements ToObjectValue() and Type().
func (m ClusterLibraryStatuses) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":       m.ClusterId,
			"library_statuses": m.LibraryStatuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterLibraryStatuses) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"library_statuses": basetypes.ListType{
				ElemType: LibraryFullStatus{}.Type(ctx),
			},
		},
	}
}

// GetLibraryStatuses returns the value of the LibraryStatuses field in ClusterLibraryStatuses as
// a slice of LibraryFullStatus values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterLibraryStatuses) GetLibraryStatuses(ctx context.Context) ([]LibraryFullStatus, bool) {
	if m.LibraryStatuses.IsNull() || m.LibraryStatuses.IsUnknown() {
		return nil, false
	}
	var v []LibraryFullStatus
	d := m.LibraryStatuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraryStatuses sets the value of the LibraryStatuses field in ClusterLibraryStatuses.
func (m *ClusterLibraryStatuses) SetLibraryStatuses(ctx context.Context, v []LibraryFullStatus) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["library_statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LibraryStatuses = types.ListValueMust(t, vs)
}

// Cluster log delivery config
type ClusterLogConf struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs types.Object `tfsdk:"dbfs"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 types.Object `tfsdk:"s3"`
	// destination needs to be provided, e.g. `{ "volumes": { "destination":
	// "/Volumes/catalog/schema/volume/cluster_log" } }`
	Volumes types.Object `tfsdk:"volumes"`
}

func (to *ClusterLogConf) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterLogConf) {
	if !from.Dbfs.IsNull() && !from.Dbfs.IsUnknown() {
		if toDbfs, ok := to.GetDbfs(ctx); ok {
			if fromDbfs, ok := from.GetDbfs(ctx); ok {
				// Recursively sync the fields of Dbfs
				toDbfs.SyncFieldsDuringCreateOrUpdate(ctx, fromDbfs)
				to.SetDbfs(ctx, toDbfs)
			}
		}
	}
	if !from.S3.IsNull() && !from.S3.IsUnknown() {
		if toS3, ok := to.GetS3(ctx); ok {
			if fromS3, ok := from.GetS3(ctx); ok {
				// Recursively sync the fields of S3
				toS3.SyncFieldsDuringCreateOrUpdate(ctx, fromS3)
				to.SetS3(ctx, toS3)
			}
		}
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() {
		if toVolumes, ok := to.GetVolumes(ctx); ok {
			if fromVolumes, ok := from.GetVolumes(ctx); ok {
				// Recursively sync the fields of Volumes
				toVolumes.SyncFieldsDuringCreateOrUpdate(ctx, fromVolumes)
				to.SetVolumes(ctx, toVolumes)
			}
		}
	}
}

func (to *ClusterLogConf) SyncFieldsDuringRead(ctx context.Context, from ClusterLogConf) {
	if !from.Dbfs.IsNull() && !from.Dbfs.IsUnknown() {
		if toDbfs, ok := to.GetDbfs(ctx); ok {
			if fromDbfs, ok := from.GetDbfs(ctx); ok {
				toDbfs.SyncFieldsDuringRead(ctx, fromDbfs)
				to.SetDbfs(ctx, toDbfs)
			}
		}
	}
	if !from.S3.IsNull() && !from.S3.IsUnknown() {
		if toS3, ok := to.GetS3(ctx); ok {
			if fromS3, ok := from.GetS3(ctx); ok {
				toS3.SyncFieldsDuringRead(ctx, fromS3)
				to.SetS3(ctx, toS3)
			}
		}
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() {
		if toVolumes, ok := to.GetVolumes(ctx); ok {
			if fromVolumes, ok := from.GetVolumes(ctx); ok {
				toVolumes.SyncFieldsDuringRead(ctx, fromVolumes)
				to.SetVolumes(ctx, toVolumes)
			}
		}
	}
}

func (m ClusterLogConf) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dbfs"] = attrs["dbfs"].SetOptional()
	attrs["s3"] = attrs["s3"].SetOptional()
	attrs["volumes"] = attrs["volumes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterLogConf.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterLogConf) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbfs":    reflect.TypeOf(DbfsStorageInfo{}),
		"s3":      reflect.TypeOf(S3StorageInfo{}),
		"volumes": reflect.TypeOf(VolumesStorageInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterLogConf
// only implements ToObjectValue() and Type().
func (m ClusterLogConf) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbfs":    m.Dbfs,
			"s3":      m.S3,
			"volumes": m.Volumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterLogConf) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbfs":    DbfsStorageInfo{}.Type(ctx),
			"s3":      S3StorageInfo{}.Type(ctx),
			"volumes": VolumesStorageInfo{}.Type(ctx),
		},
	}
}

// GetDbfs returns the value of the Dbfs field in ClusterLogConf as
// a DbfsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterLogConf) GetDbfs(ctx context.Context) (DbfsStorageInfo, bool) {
	var e DbfsStorageInfo
	if m.Dbfs.IsNull() || m.Dbfs.IsUnknown() {
		return e, false
	}
	var v DbfsStorageInfo
	d := m.Dbfs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDbfs sets the value of the Dbfs field in ClusterLogConf.
func (m *ClusterLogConf) SetDbfs(ctx context.Context, v DbfsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Dbfs = vs
}

// GetS3 returns the value of the S3 field in ClusterLogConf as
// a S3StorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterLogConf) GetS3(ctx context.Context) (S3StorageInfo, bool) {
	var e S3StorageInfo
	if m.S3.IsNull() || m.S3.IsUnknown() {
		return e, false
	}
	var v S3StorageInfo
	d := m.S3.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetS3 sets the value of the S3 field in ClusterLogConf.
func (m *ClusterLogConf) SetS3(ctx context.Context, v S3StorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.S3 = vs
}

// GetVolumes returns the value of the Volumes field in ClusterLogConf as
// a VolumesStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterLogConf) GetVolumes(ctx context.Context) (VolumesStorageInfo, bool) {
	var e VolumesStorageInfo
	if m.Volumes.IsNull() || m.Volumes.IsUnknown() {
		return e, false
	}
	var v VolumesStorageInfo
	d := m.Volumes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumes sets the value of the Volumes field in ClusterLogConf.
func (m *ClusterLogConf) SetVolumes(ctx context.Context, v VolumesStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Volumes = vs
}

type ClusterPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ClusterPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *ClusterPermission) SyncFieldsDuringRead(ctx context.Context, from ClusterPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m ClusterPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermission
// only implements ToObjectValue() and Type().
func (m ClusterPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// GetInheritedFromObject returns the value of the InheritedFromObject field in ClusterPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in ClusterPermission.
func (m *ClusterPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type ClusterPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *ClusterPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ClusterPermissions) SyncFieldsDuringRead(ctx context.Context, from ClusterPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ClusterPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissions
// only implements ToObjectValue() and Type().
func (m ClusterPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ClusterAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ClusterPermissions as
// a slice of ClusterAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterPermissions) GetAccessControlList(ctx context.Context) ([]ClusterAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPermissions.
func (m *ClusterPermissions) SetAccessControlList(ctx context.Context, v []ClusterAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type ClusterPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ClusterPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPermissionsDescription) {
}

func (to *ClusterPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from ClusterPermissionsDescription) {
}

func (m ClusterPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissionsDescription
// only implements ToObjectValue() and Type().
func (m ClusterPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ClusterPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-"`
}

func (to *ClusterPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ClusterPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from ClusterPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ClusterPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissionsRequest
// only implements ToObjectValue() and Type().
func (m ClusterPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"cluster_id":          m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ClusterAccessControlRequest{}.Type(ctx),
			},
			"cluster_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ClusterPermissionsRequest as
// a slice of ClusterAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterPermissionsRequest) GetAccessControlList(ctx context.Context) ([]ClusterAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPermissionsRequest.
func (m *ClusterPermissionsRequest) SetAccessControlList(ctx context.Context, v []ClusterAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type ClusterPolicyAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *ClusterPolicyAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPolicyAccessControlRequest) {
}

func (to *ClusterPolicyAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from ClusterPolicyAccessControlRequest) {
}

func (m ClusterPolicyAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPolicyAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPolicyAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyAccessControlRequest
// only implements ToObjectValue() and Type().
func (m ClusterPolicyAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPolicyAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ClusterPolicyAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *ClusterPolicyAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPolicyAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *ClusterPolicyAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from ClusterPolicyAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m ClusterPolicyAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPolicyAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPolicyAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ClusterPolicyPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyAccessControlResponse
// only implements ToObjectValue() and Type().
func (m ClusterPolicyAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPolicyAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ClusterPolicyPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in ClusterPolicyAccessControlResponse as
// a slice of ClusterPolicyPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterPolicyAccessControlResponse) GetAllPermissions(ctx context.Context) ([]ClusterPolicyPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ClusterPolicyAccessControlResponse.
func (m *ClusterPolicyAccessControlResponse) SetAllPermissions(ctx context.Context, v []ClusterPolicyPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type ClusterPolicyPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ClusterPolicyPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPolicyPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *ClusterPolicyPermission) SyncFieldsDuringRead(ctx context.Context, from ClusterPolicyPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m ClusterPolicyPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPolicyPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPolicyPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermission
// only implements ToObjectValue() and Type().
func (m ClusterPolicyPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPolicyPermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// GetInheritedFromObject returns the value of the InheritedFromObject field in ClusterPolicyPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterPolicyPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in ClusterPolicyPermission.
func (m *ClusterPolicyPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type ClusterPolicyPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *ClusterPolicyPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPolicyPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ClusterPolicyPermissions) SyncFieldsDuringRead(ctx context.Context, from ClusterPolicyPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ClusterPolicyPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPolicyPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPolicyPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterPolicyAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissions
// only implements ToObjectValue() and Type().
func (m ClusterPolicyPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPolicyPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ClusterPolicyAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ClusterPolicyPermissions as
// a slice of ClusterPolicyAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterPolicyPermissions) GetAccessControlList(ctx context.Context) ([]ClusterPolicyAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPolicyPermissions.
func (m *ClusterPolicyPermissions) SetAccessControlList(ctx context.Context, v []ClusterPolicyAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type ClusterPolicyPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ClusterPolicyPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPolicyPermissionsDescription) {
}

func (to *ClusterPolicyPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from ClusterPolicyPermissionsDescription) {
}

func (m ClusterPolicyPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPolicyPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPolicyPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissionsDescription
// only implements ToObjectValue() and Type().
func (m ClusterPolicyPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPolicyPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ClusterPolicyPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-"`
}

func (to *ClusterPolicyPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterPolicyPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ClusterPolicyPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from ClusterPolicyPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ClusterPolicyPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["cluster_policy_id"] = attrs["cluster_policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPolicyPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterPolicyPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterPolicyAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissionsRequest
// only implements ToObjectValue() and Type().
func (m ClusterPolicyPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"cluster_policy_id":   m.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterPolicyPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ClusterPolicyAccessControlRequest{}.Type(ctx),
			},
			"cluster_policy_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ClusterPolicyPermissionsRequest as
// a slice of ClusterPolicyAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterPolicyPermissionsRequest) GetAccessControlList(ctx context.Context) ([]ClusterPolicyAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPolicyPermissionsRequest.
func (m *ClusterPolicyPermissionsRequest) SetAccessControlList(ctx context.Context, v []ClusterPolicyAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// Represents a change to the cluster settings required for the cluster to
// become compliant with its policy.
type ClusterSettingsChange struct {
	// The field where this change would be made.
	Field types.String `tfsdk:"field"`
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	NewValue types.String `tfsdk:"new_value"`
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	PreviousValue types.String `tfsdk:"previous_value"`
}

func (to *ClusterSettingsChange) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterSettingsChange) {
}

func (to *ClusterSettingsChange) SyncFieldsDuringRead(ctx context.Context, from ClusterSettingsChange) {
}

func (m ClusterSettingsChange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["field"] = attrs["field"].SetOptional()
	attrs["new_value"] = attrs["new_value"].SetOptional()
	attrs["previous_value"] = attrs["previous_value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterSettingsChange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterSettingsChange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSettingsChange
// only implements ToObjectValue() and Type().
func (m ClusterSettingsChange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"field":          m.Field,
			"new_value":      m.NewValue,
			"previous_value": m.PreviousValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterSettingsChange) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"field":          types.StringType,
			"new_value":      types.StringType,
			"previous_value": types.StringType,
		},
	}
}

type ClusterSize struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers"`
}

func (to *ClusterSize) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterSize) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				// Recursively sync the fields of Autoscale
				toAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
}

func (to *ClusterSize) SyncFieldsDuringRead(ctx context.Context, from ClusterSize) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				toAutoscale.SyncFieldsDuringRead(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
}

func (m ClusterSize) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["num_workers"] = attrs["num_workers"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterSize.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterSize) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale": reflect.TypeOf(AutoScale{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSize
// only implements ToObjectValue() and Type().
func (m ClusterSize) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":   m.Autoscale,
			"num_workers": m.NumWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterSize) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale":   AutoScale{}.Type(ctx),
			"num_workers": types.Int64Type,
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in ClusterSize as
// a AutoScale value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSize) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v AutoScale
	d := m.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoscale sets the value of the Autoscale field in ClusterSize.
func (m *ClusterSize) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	m.Autoscale = vs
}

// Contains a snapshot of the latest user specified settings that were used to
// create/edit the cluster.
type ClusterSpec struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.Object `tfsdk:"cluster_log_conf"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName types.String `tfsdk:"cluster_name"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags"`

	DataSecurityMode types.String `tfsdk:"data_security_mode"`
	// Custom docker image BYOC
	DockerImage types.Object `tfsdk:"docker_image"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode types.Bool `tfsdk:"is_single_node"`

	Kind types.String `tfsdk:"kind"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars types.Map `tfsdk:"spark_env_vars"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime types.Bool `tfsdk:"use_ml_runtime"`

	WorkloadType types.Object `tfsdk:"workload_type"`
}

func (to *ClusterSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterSpec) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				// Recursively sync the fields of Autoscale
				toAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				// Recursively sync the fields of ClusterLogConf
				toClusterLogConf.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				// Recursively sync the fields of DockerImage
				toDockerImage.SyncFieldsDuringCreateOrUpdate(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				// Recursively sync the fields of WorkloadType
				toWorkloadType.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (to *ClusterSpec) SyncFieldsDuringRead(ctx context.Context, from ClusterSpec) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				toAutoscale.SyncFieldsDuringRead(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				toClusterLogConf.SyncFieldsDuringRead(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				toDockerImage.SyncFieldsDuringRead(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				toWorkloadType.SyncFieldsDuringRead(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (m ClusterSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_policy_default_values"] = attrs["apply_policy_default_values"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["is_single_node"] = attrs["is_single_node"].SetOptional()
	attrs["kind"] = attrs["kind"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["num_workers"] = attrs["num_workers"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["runtime_engine"] = attrs["runtime_engine"].SetOptional()
	attrs["single_user_name"] = attrs["single_user_name"].SetOptional()
	attrs["spark_conf"] = attrs["spark_conf"].SetOptional()
	attrs["spark_env_vars"] = attrs["spark_env_vars"].SetOptional()
	attrs["spark_version"] = attrs["spark_version"].SetOptional()
	attrs["ssh_public_keys"] = attrs["ssh_public_keys"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()
	attrs["use_ml_runtime"] = attrs["use_ml_runtime"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(AutoScale{}),
		"aws_attributes":   reflect.TypeOf(AwsAttributes{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSpec
// only implements ToObjectValue() and Type().
func (m ClusterSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_policy_default_values":    m.ApplyPolicyDefaultValues,
			"autoscale":                      m.Autoscale,
			"autotermination_minutes":        m.AutoterminationMinutes,
			"aws_attributes":                 m.AwsAttributes,
			"azure_attributes":               m.AzureAttributes,
			"cluster_log_conf":               m.ClusterLogConf,
			"cluster_name":                   m.ClusterName,
			"custom_tags":                    m.CustomTags,
			"data_security_mode":             m.DataSecurityMode,
			"docker_image":                   m.DockerImage,
			"driver_instance_pool_id":        m.DriverInstancePoolId,
			"driver_node_type_id":            m.DriverNodeTypeId,
			"enable_elastic_disk":            m.EnableElasticDisk,
			"enable_local_disk_encryption":   m.EnableLocalDiskEncryption,
			"gcp_attributes":                 m.GcpAttributes,
			"init_scripts":                   m.InitScripts,
			"instance_pool_id":               m.InstancePoolId,
			"is_single_node":                 m.IsSingleNode,
			"kind":                           m.Kind,
			"node_type_id":                   m.NodeTypeId,
			"num_workers":                    m.NumWorkers,
			"policy_id":                      m.PolicyId,
			"remote_disk_throughput":         m.RemoteDiskThroughput,
			"runtime_engine":                 m.RuntimeEngine,
			"single_user_name":               m.SingleUserName,
			"spark_conf":                     m.SparkConf,
			"spark_env_vars":                 m.SparkEnvVars,
			"spark_version":                  m.SparkVersion,
			"ssh_public_keys":                m.SshPublicKeys,
			"total_initial_remote_disk_size": m.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 m.UseMlRuntime,
			"workload_type":                  m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale":                   AutoScale{}.Type(ctx),
			"autotermination_minutes":     types.Int64Type,
			"aws_attributes":              AwsAttributes{}.Type(ctx),
			"azure_attributes":            AzureAttributes{}.Type(ctx),
			"cluster_log_conf":            ClusterLogConf{}.Type(ctx),
			"cluster_name":                types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode":           types.StringType,
			"docker_image":                 DockerImage{}.Type(ctx),
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes":               GcpAttributes{}.Type(ctx),
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id":       types.StringType,
			"is_single_node":         types.BoolType,
			"kind":                   types.StringType,
			"node_type_id":           types.StringType,
			"num_workers":            types.Int64Type,
			"policy_id":              types.StringType,
			"remote_disk_throughput": types.Int64Type,
			"runtime_engine":         types.StringType,
			"single_user_name":       types.StringType,
			"spark_conf": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_env_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_version": types.StringType,
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"total_initial_remote_disk_size": types.Int64Type,
			"use_ml_runtime":                 types.BoolType,
			"workload_type":                  WorkloadType{}.Type(ctx),
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in ClusterSpec as
// a AutoScale value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v AutoScale
	d := m.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoscale sets the value of the Autoscale field in ClusterSpec.
func (m *ClusterSpec) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	m.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in ClusterSpec as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v AwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterSpec.
func (m *ClusterSpec) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterSpec as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v AzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterSpec.
func (m *ClusterSpec) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterSpec as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if m.ClusterLogConf.IsNull() || m.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v ClusterLogConf
	d := m.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterSpec.
func (m *ClusterSpec) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	m.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in ClusterSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in ClusterSpec.
func (m *ClusterSpec) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterSpec as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if m.DockerImage.IsNull() || m.DockerImage.IsUnknown() {
		return e, false
	}
	var v DockerImage
	d := m.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDockerImage sets the value of the DockerImage field in ClusterSpec.
func (m *ClusterSpec) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	m.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterSpec as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v GcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterSpec.
func (m *ClusterSpec) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in ClusterSpec as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := m.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterSpec.
func (m *ClusterSpec) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkConf.IsNull() || m.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in ClusterSpec.
func (m *ClusterSpec) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkEnvVars.IsNull() || m.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterSpec.
func (m *ClusterSpec) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterSpec as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if m.SshPublicKeys.IsNull() || m.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterSpec.
func (m *ClusterSpec) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterSpec as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterSpec) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if m.WorkloadType.IsNull() || m.WorkloadType.IsUnknown() {
		return e, false
	}
	var v WorkloadType
	d := m.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterSpec.
func (m *ClusterSpec) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	m.WorkloadType = vs
}

type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId types.String `tfsdk:"-"`
}

func (to *ClusterStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterStatus) {
}

func (to *ClusterStatus) SyncFieldsDuringRead(ctx context.Context, from ClusterStatus) {
}

func (m ClusterStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterStatus
// only implements ToObjectValue() and Type().
func (m ClusterStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type Command struct {
	// Running cluster id
	ClusterId types.String `tfsdk:"cluster_id"`
	// Executable code
	Command types.String `tfsdk:"command"`
	// Running context id
	ContextId types.String `tfsdk:"context_id"`

	Language types.String `tfsdk:"language"`
}

func (to *Command) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Command) {
}

func (to *Command) SyncFieldsDuringRead(ctx context.Context, from Command) {
}

func (m Command) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["command"] = attrs["command"].SetOptional()
	attrs["context_id"] = attrs["context_id"].SetOptional()
	attrs["language"] = attrs["language"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Command.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Command) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Command
// only implements ToObjectValue() and Type().
func (m Command) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
			"command":    m.Command,
			"context_id": m.ContextId,
			"language":   m.Language,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Command) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"command":    types.StringType,
			"context_id": types.StringType,
			"language":   types.StringType,
		},
	}
}

type CommandStatusRequest struct {
	ClusterId types.String `tfsdk:"-"`

	CommandId types.String `tfsdk:"-"`

	ContextId types.String `tfsdk:"-"`
}

func (to *CommandStatusRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CommandStatusRequest) {
}

func (to *CommandStatusRequest) SyncFieldsDuringRead(ctx context.Context, from CommandStatusRequest) {
}

func (m CommandStatusRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["context_id"] = attrs["context_id"].SetRequired()
	attrs["command_id"] = attrs["command_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CommandStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CommandStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommandStatusRequest
// only implements ToObjectValue() and Type().
func (m CommandStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
			"command_id": m.CommandId,
			"context_id": m.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CommandStatusRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"command_id": types.StringType,
			"context_id": types.StringType,
		},
	}
}

type CommandStatusResponse struct {
	Id types.String `tfsdk:"id"`

	Results types.Object `tfsdk:"results"`

	Status types.String `tfsdk:"status"`
}

func (to *CommandStatusResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CommandStatusResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() {
		if toResults, ok := to.GetResults(ctx); ok {
			if fromResults, ok := from.GetResults(ctx); ok {
				// Recursively sync the fields of Results
				toResults.SyncFieldsDuringCreateOrUpdate(ctx, fromResults)
				to.SetResults(ctx, toResults)
			}
		}
	}
}

func (to *CommandStatusResponse) SyncFieldsDuringRead(ctx context.Context, from CommandStatusResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() {
		if toResults, ok := to.GetResults(ctx); ok {
			if fromResults, ok := from.GetResults(ctx); ok {
				toResults.SyncFieldsDuringRead(ctx, fromResults)
				to.SetResults(ctx, toResults)
			}
		}
	}
}

func (m CommandStatusResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CommandStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CommandStatusResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Results{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommandStatusResponse
// only implements ToObjectValue() and Type().
func (m CommandStatusResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      m.Id,
			"results": m.Results,
			"status":  m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CommandStatusResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":      types.StringType,
			"results": Results{}.Type(ctx),
			"status":  types.StringType,
		},
	}
}

// GetResults returns the value of the Results field in CommandStatusResponse as
// a Results value.
// If the field is unknown or null, the boolean return value is false.
func (m *CommandStatusResponse) GetResults(ctx context.Context) (Results, bool) {
	var e Results
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return e, false
	}
	var v Results
	d := m.Results.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in CommandStatusResponse.
func (m *CommandStatusResponse) SetResults(ctx context.Context, v Results) {
	vs := v.ToObjectValue(ctx)
	m.Results = vs
}

type ContextStatusRequest struct {
	ClusterId types.String `tfsdk:"-"`

	ContextId types.String `tfsdk:"-"`
}

func (to *ContextStatusRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ContextStatusRequest) {
}

func (to *ContextStatusRequest) SyncFieldsDuringRead(ctx context.Context, from ContextStatusRequest) {
}

func (m ContextStatusRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["context_id"] = attrs["context_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ContextStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ContextStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContextStatusRequest
// only implements ToObjectValue() and Type().
func (m ContextStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
			"context_id": m.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ContextStatusRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"context_id": types.StringType,
		},
	}
}

type ContextStatusResponse struct {
	Id types.String `tfsdk:"id"`

	Status types.String `tfsdk:"status"`
}

func (to *ContextStatusResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ContextStatusResponse) {
}

func (to *ContextStatusResponse) SyncFieldsDuringRead(ctx context.Context, from ContextStatusResponse) {
}

func (m ContextStatusResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ContextStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ContextStatusResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContextStatusResponse
// only implements ToObjectValue() and Type().
func (m ContextStatusResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":     m.Id,
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ContextStatusResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":     types.StringType,
			"status": types.StringType,
		},
	}
}

type CreateCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	CloneFrom types.Object `tfsdk:"clone_from"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.Object `tfsdk:"cluster_log_conf"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName types.String `tfsdk:"cluster_name"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags"`

	DataSecurityMode types.String `tfsdk:"data_security_mode"`
	// Custom docker image BYOC
	DockerImage types.Object `tfsdk:"docker_image"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode types.Bool `tfsdk:"is_single_node"`

	Kind types.String `tfsdk:"kind"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars types.Map `tfsdk:"spark_env_vars"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime types.Bool `tfsdk:"use_ml_runtime"`

	WorkloadType types.Object `tfsdk:"workload_type"`
}

func (to *CreateCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCluster) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				// Recursively sync the fields of Autoscale
				toAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.CloneFrom.IsNull() && !from.CloneFrom.IsUnknown() {
		if toCloneFrom, ok := to.GetCloneFrom(ctx); ok {
			if fromCloneFrom, ok := from.GetCloneFrom(ctx); ok {
				// Recursively sync the fields of CloneFrom
				toCloneFrom.SyncFieldsDuringCreateOrUpdate(ctx, fromCloneFrom)
				to.SetCloneFrom(ctx, toCloneFrom)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				// Recursively sync the fields of ClusterLogConf
				toClusterLogConf.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				// Recursively sync the fields of DockerImage
				toDockerImage.SyncFieldsDuringCreateOrUpdate(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				// Recursively sync the fields of WorkloadType
				toWorkloadType.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (to *CreateCluster) SyncFieldsDuringRead(ctx context.Context, from CreateCluster) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				toAutoscale.SyncFieldsDuringRead(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.CloneFrom.IsNull() && !from.CloneFrom.IsUnknown() {
		if toCloneFrom, ok := to.GetCloneFrom(ctx); ok {
			if fromCloneFrom, ok := from.GetCloneFrom(ctx); ok {
				toCloneFrom.SyncFieldsDuringRead(ctx, fromCloneFrom)
				to.SetCloneFrom(ctx, toCloneFrom)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				toClusterLogConf.SyncFieldsDuringRead(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				toDockerImage.SyncFieldsDuringRead(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				toWorkloadType.SyncFieldsDuringRead(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (m CreateCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_policy_default_values"] = attrs["apply_policy_default_values"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["clone_from"] = attrs["clone_from"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["is_single_node"] = attrs["is_single_node"].SetOptional()
	attrs["kind"] = attrs["kind"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["num_workers"] = attrs["num_workers"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["runtime_engine"] = attrs["runtime_engine"].SetOptional()
	attrs["single_user_name"] = attrs["single_user_name"].SetOptional()
	attrs["spark_conf"] = attrs["spark_conf"].SetOptional()
	attrs["spark_env_vars"] = attrs["spark_env_vars"].SetOptional()
	attrs["spark_version"] = attrs["spark_version"].SetRequired()
	attrs["ssh_public_keys"] = attrs["ssh_public_keys"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()
	attrs["use_ml_runtime"] = attrs["use_ml_runtime"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(AutoScale{}),
		"aws_attributes":   reflect.TypeOf(AwsAttributes{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes{}),
		"clone_from":       reflect.TypeOf(CloneCluster{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCluster
// only implements ToObjectValue() and Type().
func (m CreateCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_policy_default_values":    m.ApplyPolicyDefaultValues,
			"autoscale":                      m.Autoscale,
			"autotermination_minutes":        m.AutoterminationMinutes,
			"aws_attributes":                 m.AwsAttributes,
			"azure_attributes":               m.AzureAttributes,
			"clone_from":                     m.CloneFrom,
			"cluster_log_conf":               m.ClusterLogConf,
			"cluster_name":                   m.ClusterName,
			"custom_tags":                    m.CustomTags,
			"data_security_mode":             m.DataSecurityMode,
			"docker_image":                   m.DockerImage,
			"driver_instance_pool_id":        m.DriverInstancePoolId,
			"driver_node_type_id":            m.DriverNodeTypeId,
			"enable_elastic_disk":            m.EnableElasticDisk,
			"enable_local_disk_encryption":   m.EnableLocalDiskEncryption,
			"gcp_attributes":                 m.GcpAttributes,
			"init_scripts":                   m.InitScripts,
			"instance_pool_id":               m.InstancePoolId,
			"is_single_node":                 m.IsSingleNode,
			"kind":                           m.Kind,
			"node_type_id":                   m.NodeTypeId,
			"num_workers":                    m.NumWorkers,
			"policy_id":                      m.PolicyId,
			"remote_disk_throughput":         m.RemoteDiskThroughput,
			"runtime_engine":                 m.RuntimeEngine,
			"single_user_name":               m.SingleUserName,
			"spark_conf":                     m.SparkConf,
			"spark_env_vars":                 m.SparkEnvVars,
			"spark_version":                  m.SparkVersion,
			"ssh_public_keys":                m.SshPublicKeys,
			"total_initial_remote_disk_size": m.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 m.UseMlRuntime,
			"workload_type":                  m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale":                   AutoScale{}.Type(ctx),
			"autotermination_minutes":     types.Int64Type,
			"aws_attributes":              AwsAttributes{}.Type(ctx),
			"azure_attributes":            AzureAttributes{}.Type(ctx),
			"clone_from":                  CloneCluster{}.Type(ctx),
			"cluster_log_conf":            ClusterLogConf{}.Type(ctx),
			"cluster_name":                types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode":           types.StringType,
			"docker_image":                 DockerImage{}.Type(ctx),
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes":               GcpAttributes{}.Type(ctx),
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id":       types.StringType,
			"is_single_node":         types.BoolType,
			"kind":                   types.StringType,
			"node_type_id":           types.StringType,
			"num_workers":            types.Int64Type,
			"policy_id":              types.StringType,
			"remote_disk_throughput": types.Int64Type,
			"runtime_engine":         types.StringType,
			"single_user_name":       types.StringType,
			"spark_conf": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_env_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_version": types.StringType,
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"total_initial_remote_disk_size": types.Int64Type,
			"use_ml_runtime":                 types.BoolType,
			"workload_type":                  WorkloadType{}.Type(ctx),
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in CreateCluster as
// a AutoScale value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v AutoScale
	d := m.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoscale sets the value of the Autoscale field in CreateCluster.
func (m *CreateCluster) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	m.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in CreateCluster as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v AwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in CreateCluster.
func (m *CreateCluster) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in CreateCluster as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v AzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in CreateCluster.
func (m *CreateCluster) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetCloneFrom returns the value of the CloneFrom field in CreateCluster as
// a CloneCluster value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetCloneFrom(ctx context.Context) (CloneCluster, bool) {
	var e CloneCluster
	if m.CloneFrom.IsNull() || m.CloneFrom.IsUnknown() {
		return e, false
	}
	var v CloneCluster
	d := m.CloneFrom.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCloneFrom sets the value of the CloneFrom field in CreateCluster.
func (m *CreateCluster) SetCloneFrom(ctx context.Context, v CloneCluster) {
	vs := v.ToObjectValue(ctx)
	m.CloneFrom = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in CreateCluster as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if m.ClusterLogConf.IsNull() || m.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v ClusterLogConf
	d := m.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in CreateCluster.
func (m *CreateCluster) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	m.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in CreateCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in CreateCluster.
func (m *CreateCluster) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in CreateCluster as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if m.DockerImage.IsNull() || m.DockerImage.IsUnknown() {
		return e, false
	}
	var v DockerImage
	d := m.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDockerImage sets the value of the DockerImage field in CreateCluster.
func (m *CreateCluster) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	m.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in CreateCluster as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v GcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in CreateCluster.
func (m *CreateCluster) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in CreateCluster as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := m.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in CreateCluster.
func (m *CreateCluster) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in CreateCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkConf.IsNull() || m.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in CreateCluster.
func (m *CreateCluster) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in CreateCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkEnvVars.IsNull() || m.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in CreateCluster.
func (m *CreateCluster) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in CreateCluster as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if m.SshPublicKeys.IsNull() || m.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in CreateCluster.
func (m *CreateCluster) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in CreateCluster as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCluster) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if m.WorkloadType.IsNull() || m.WorkloadType.IsUnknown() {
		return e, false
	}
	var v WorkloadType
	d := m.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkloadType sets the value of the WorkloadType field in CreateCluster.
func (m *CreateCluster) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	m.WorkloadType = vs
}

type CreateClusterResponse struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

func (to *CreateClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateClusterResponse) {
}

func (to *CreateClusterResponse) SyncFieldsDuringRead(ctx context.Context, from CreateClusterResponse) {
}

func (m CreateClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateClusterResponse
// only implements ToObjectValue() and Type().
func (m CreateClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type CreateContext struct {
	// Running cluster id
	ClusterId types.String `tfsdk:"cluster_id"`

	Language types.String `tfsdk:"language"`
}

func (to *CreateContext) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateContext) {
}

func (to *CreateContext) SyncFieldsDuringRead(ctx context.Context, from CreateContext) {
}

func (m CreateContext) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["language"] = attrs["language"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateContext.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateContext) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateContext
// only implements ToObjectValue() and Type().
func (m CreateContext) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
			"language":   m.Language,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateContext) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"language":   types.StringType,
		},
	}
}

type CreateDefaultBaseEnvironmentRequest struct {
	DefaultBaseEnvironment types.Object `tfsdk:"default_base_environment"`
	// A unique identifier for this request. A random UUID is recommended. This
	// request is only idempotent if a `request_id` is provided.
	RequestId types.String `tfsdk:"request_id"`
}

func (to *CreateDefaultBaseEnvironmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDefaultBaseEnvironmentRequest) {
	if !from.DefaultBaseEnvironment.IsNull() && !from.DefaultBaseEnvironment.IsUnknown() {
		if toDefaultBaseEnvironment, ok := to.GetDefaultBaseEnvironment(ctx); ok {
			if fromDefaultBaseEnvironment, ok := from.GetDefaultBaseEnvironment(ctx); ok {
				// Recursively sync the fields of DefaultBaseEnvironment
				toDefaultBaseEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultBaseEnvironment)
				to.SetDefaultBaseEnvironment(ctx, toDefaultBaseEnvironment)
			}
		}
	}
}

func (to *CreateDefaultBaseEnvironmentRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDefaultBaseEnvironmentRequest) {
	if !from.DefaultBaseEnvironment.IsNull() && !from.DefaultBaseEnvironment.IsUnknown() {
		if toDefaultBaseEnvironment, ok := to.GetDefaultBaseEnvironment(ctx); ok {
			if fromDefaultBaseEnvironment, ok := from.GetDefaultBaseEnvironment(ctx); ok {
				toDefaultBaseEnvironment.SyncFieldsDuringRead(ctx, fromDefaultBaseEnvironment)
				to.SetDefaultBaseEnvironment(ctx, toDefaultBaseEnvironment)
			}
		}
	}
}

func (m CreateDefaultBaseEnvironmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_base_environment"] = attrs["default_base_environment"].SetRequired()
	attrs["request_id"] = attrs["request_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDefaultBaseEnvironmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_base_environment": reflect.TypeOf(DefaultBaseEnvironment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDefaultBaseEnvironmentRequest
// only implements ToObjectValue() and Type().
func (m CreateDefaultBaseEnvironmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_base_environment": m.DefaultBaseEnvironment,
			"request_id":               m.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDefaultBaseEnvironmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_base_environment": DefaultBaseEnvironment{}.Type(ctx),
			"request_id":               types.StringType,
		},
	}
}

// GetDefaultBaseEnvironment returns the value of the DefaultBaseEnvironment field in CreateDefaultBaseEnvironmentRequest as
// a DefaultBaseEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDefaultBaseEnvironmentRequest) GetDefaultBaseEnvironment(ctx context.Context) (DefaultBaseEnvironment, bool) {
	var e DefaultBaseEnvironment
	if m.DefaultBaseEnvironment.IsNull() || m.DefaultBaseEnvironment.IsUnknown() {
		return e, false
	}
	var v DefaultBaseEnvironment
	d := m.DefaultBaseEnvironment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultBaseEnvironment sets the value of the DefaultBaseEnvironment field in CreateDefaultBaseEnvironmentRequest.
func (m *CreateDefaultBaseEnvironmentRequest) SetDefaultBaseEnvironment(ctx context.Context, v DefaultBaseEnvironment) {
	vs := v.ToObjectValue(ctx)
	m.DefaultBaseEnvironment = vs
}

type CreateInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags types.Map `tfsdk:"custom_tags"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec types.Object `tfsdk:"disk_spec"`
	// For pools with node type flexibility (Fleet-V2), whether auto generated
	// alternate node type ids are enabled. This field should not be true if
	// node_type_flexibility is set.
	EnableAutoAlternateNodeTypes types.Bool `tfsdk:"enable_auto_alternate_node_types"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances"`
	// For pools with node type flexibility (Fleet-V2), this object contains the
	// information about the alternate node type ids to use when attempting to
	// launch a cluster if the node type id is not available. This field should
	// not be set if enable_auto_alternate_node_types is true.
	NodeTypeFlexibility types.Object `tfsdk:"node_type_flexibility"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Custom Docker Image BYOC
	PreloadedDockerImages types.List `tfsdk:"preloaded_docker_images"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions types.List `tfsdk:"preloaded_spark_versions"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
}

func (to *CreateInstancePool) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateInstancePool) {
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.DiskSpec.IsNull() && !from.DiskSpec.IsUnknown() {
		if toDiskSpec, ok := to.GetDiskSpec(ctx); ok {
			if fromDiskSpec, ok := from.GetDiskSpec(ctx); ok {
				// Recursively sync the fields of DiskSpec
				toDiskSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDiskSpec)
				to.SetDiskSpec(ctx, toDiskSpec)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.NodeTypeFlexibility.IsNull() && !from.NodeTypeFlexibility.IsUnknown() {
		if toNodeTypeFlexibility, ok := to.GetNodeTypeFlexibility(ctx); ok {
			if fromNodeTypeFlexibility, ok := from.GetNodeTypeFlexibility(ctx); ok {
				// Recursively sync the fields of NodeTypeFlexibility
				toNodeTypeFlexibility.SyncFieldsDuringCreateOrUpdate(ctx, fromNodeTypeFlexibility)
				to.SetNodeTypeFlexibility(ctx, toNodeTypeFlexibility)
			}
		}
	}
	if !from.PreloadedDockerImages.IsNull() && !from.PreloadedDockerImages.IsUnknown() && to.PreloadedDockerImages.IsNull() && len(from.PreloadedDockerImages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedDockerImages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedDockerImages = from.PreloadedDockerImages
	}
	if !from.PreloadedSparkVersions.IsNull() && !from.PreloadedSparkVersions.IsUnknown() && to.PreloadedSparkVersions.IsNull() && len(from.PreloadedSparkVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedSparkVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedSparkVersions = from.PreloadedSparkVersions
	}
}

func (to *CreateInstancePool) SyncFieldsDuringRead(ctx context.Context, from CreateInstancePool) {
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.DiskSpec.IsNull() && !from.DiskSpec.IsUnknown() {
		if toDiskSpec, ok := to.GetDiskSpec(ctx); ok {
			if fromDiskSpec, ok := from.GetDiskSpec(ctx); ok {
				toDiskSpec.SyncFieldsDuringRead(ctx, fromDiskSpec)
				to.SetDiskSpec(ctx, toDiskSpec)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.NodeTypeFlexibility.IsNull() && !from.NodeTypeFlexibility.IsUnknown() {
		if toNodeTypeFlexibility, ok := to.GetNodeTypeFlexibility(ctx); ok {
			if fromNodeTypeFlexibility, ok := from.GetNodeTypeFlexibility(ctx); ok {
				toNodeTypeFlexibility.SyncFieldsDuringRead(ctx, fromNodeTypeFlexibility)
				to.SetNodeTypeFlexibility(ctx, toNodeTypeFlexibility)
			}
		}
	}
	if !from.PreloadedDockerImages.IsNull() && !from.PreloadedDockerImages.IsUnknown() && to.PreloadedDockerImages.IsNull() && len(from.PreloadedDockerImages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedDockerImages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedDockerImages = from.PreloadedDockerImages
	}
	if !from.PreloadedSparkVersions.IsNull() && !from.PreloadedSparkVersions.IsUnknown() && to.PreloadedSparkVersions.IsNull() && len(from.PreloadedSparkVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedSparkVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedSparkVersions = from.PreloadedSparkVersions
	}
}

func (m CreateInstancePool) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].SetOptional()
	attrs["enable_auto_alternate_node_types"] = attrs["enable_auto_alternate_node_types"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["idle_instance_autotermination_minutes"] = attrs["idle_instance_autotermination_minutes"].SetOptional()
	attrs["instance_pool_name"] = attrs["instance_pool_name"].SetRequired()
	attrs["max_capacity"] = attrs["max_capacity"].SetOptional()
	attrs["min_idle_instances"] = attrs["min_idle_instances"].SetOptional()
	attrs["node_type_flexibility"] = attrs["node_type_flexibility"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetRequired()
	attrs["preloaded_docker_images"] = attrs["preloaded_docker_images"].SetOptional()
	attrs["preloaded_spark_versions"] = attrs["preloaded_spark_versions"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateInstancePool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateInstancePool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes{}),
		"node_type_flexibility":    reflect.TypeOf(NodeTypeFlexibility{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstancePool
// only implements ToObjectValue() and Type().
func (m CreateInstancePool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        m.AwsAttributes,
			"azure_attributes":                      m.AzureAttributes,
			"custom_tags":                           m.CustomTags,
			"disk_spec":                             m.DiskSpec,
			"enable_auto_alternate_node_types":      m.EnableAutoAlternateNodeTypes,
			"enable_elastic_disk":                   m.EnableElasticDisk,
			"gcp_attributes":                        m.GcpAttributes,
			"idle_instance_autotermination_minutes": m.IdleInstanceAutoterminationMinutes,
			"instance_pool_name":                    m.InstancePoolName,
			"max_capacity":                          m.MaxCapacity,
			"min_idle_instances":                    m.MinIdleInstances,
			"node_type_flexibility":                 m.NodeTypeFlexibility,
			"node_type_id":                          m.NodeTypeId,
			"preloaded_docker_images":               m.PreloadedDockerImages,
			"preloaded_spark_versions":              m.PreloadedSparkVersions,
			"remote_disk_throughput":                m.RemoteDiskThroughput,
			"total_initial_remote_disk_size":        m.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes":   InstancePoolAwsAttributes{}.Type(ctx),
			"azure_attributes": InstancePoolAzureAttributes{}.Type(ctx),
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec":                             DiskSpec{}.Type(ctx),
			"enable_auto_alternate_node_types":      types.BoolType,
			"enable_elastic_disk":                   types.BoolType,
			"gcp_attributes":                        InstancePoolGcpAttributes{}.Type(ctx),
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
			"node_type_flexibility":                 NodeTypeFlexibility{}.Type(ctx),
			"node_type_id":                          types.StringType,
			"preloaded_docker_images": basetypes.ListType{
				ElemType: DockerImage{}.Type(ctx),
			},
			"preloaded_spark_versions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"remote_disk_throughput":         types.Int64Type,
			"total_initial_remote_disk_size": types.Int64Type,
		},
	}
}

// GetAwsAttributes returns the value of the AwsAttributes field in CreateInstancePool as
// a InstancePoolAwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstancePool) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes, bool) {
	var e InstancePoolAwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolAwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in CreateInstancePool.
func (m *CreateInstancePool) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in CreateInstancePool as
// a InstancePoolAzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstancePool) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes, bool) {
	var e InstancePoolAzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolAzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in CreateInstancePool.
func (m *CreateInstancePool) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetCustomTags returns the value of the CustomTags field in CreateInstancePool as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstancePool) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in CreateInstancePool.
func (m *CreateInstancePool) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in CreateInstancePool as
// a DiskSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstancePool) GetDiskSpec(ctx context.Context) (DiskSpec, bool) {
	var e DiskSpec
	if m.DiskSpec.IsNull() || m.DiskSpec.IsUnknown() {
		return e, false
	}
	var v DiskSpec
	d := m.DiskSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDiskSpec sets the value of the DiskSpec field in CreateInstancePool.
func (m *CreateInstancePool) SetDiskSpec(ctx context.Context, v DiskSpec) {
	vs := v.ToObjectValue(ctx)
	m.DiskSpec = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in CreateInstancePool as
// a InstancePoolGcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstancePool) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes, bool) {
	var e InstancePoolGcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolGcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in CreateInstancePool.
func (m *CreateInstancePool) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetNodeTypeFlexibility returns the value of the NodeTypeFlexibility field in CreateInstancePool as
// a NodeTypeFlexibility value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstancePool) GetNodeTypeFlexibility(ctx context.Context) (NodeTypeFlexibility, bool) {
	var e NodeTypeFlexibility
	if m.NodeTypeFlexibility.IsNull() || m.NodeTypeFlexibility.IsUnknown() {
		return e, false
	}
	var v NodeTypeFlexibility
	d := m.NodeTypeFlexibility.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeTypeFlexibility sets the value of the NodeTypeFlexibility field in CreateInstancePool.
func (m *CreateInstancePool) SetNodeTypeFlexibility(ctx context.Context, v NodeTypeFlexibility) {
	vs := v.ToObjectValue(ctx)
	m.NodeTypeFlexibility = vs
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in CreateInstancePool as
// a slice of DockerImage values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstancePool) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage, bool) {
	if m.PreloadedDockerImages.IsNull() || m.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage
	d := m.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in CreateInstancePool.
func (m *CreateInstancePool) SetPreloadedDockerImages(ctx context.Context, v []DockerImage) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in CreateInstancePool as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateInstancePool) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
	if m.PreloadedSparkVersions.IsNull() || m.PreloadedSparkVersions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.PreloadedSparkVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in CreateInstancePool.
func (m *CreateInstancePool) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
}

func (to *CreateInstancePoolResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateInstancePoolResponse) {
}

func (to *CreateInstancePoolResponse) SyncFieldsDuringRead(ctx context.Context, from CreateInstancePoolResponse) {
}

func (m CreateInstancePoolResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateInstancePoolResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateInstancePoolResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstancePoolResponse
// only implements ToObjectValue() and Type().
func (m CreateInstancePoolResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": m.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateInstancePoolResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type CreatePolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries types.List `tfsdk:"libraries"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId types.String `tfsdk:"policy_family_id"`
}

func (to *CreatePolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePolicy) {
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
}

func (to *CreatePolicy) SyncFieldsDuringRead(ctx context.Context, from CreatePolicy) {
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
}

func (m CreatePolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["definition"] = attrs["definition"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["libraries"] = attrs["libraries"].SetOptional()
	attrs["max_clusters_per_user"] = attrs["max_clusters_per_user"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["policy_family_definition_overrides"] = attrs["policy_family_definition_overrides"].SetOptional()
	attrs["policy_family_id"] = attrs["policy_family_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePolicy
// only implements ToObjectValue() and Type().
func (m CreatePolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"definition":                         m.Definition,
			"description":                        m.Description,
			"libraries":                          m.Libraries,
			"max_clusters_per_user":              m.MaxClustersPerUser,
			"name":                               m.Name,
			"policy_family_definition_overrides": m.PolicyFamilyDefinitionOverrides,
			"policy_family_id":                   m.PolicyFamilyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"definition":  types.StringType,
			"description": types.StringType,
			"libraries": basetypes.ListType{
				ElemType: Library{}.Type(ctx),
			},
			"max_clusters_per_user":              types.Int64Type,
			"name":                               types.StringType,
			"policy_family_definition_overrides": types.StringType,
			"policy_family_id":                   types.StringType,
		},
	}
}

// GetLibraries returns the value of the Libraries field in CreatePolicy as
// a slice of Library values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePolicy) GetLibraries(ctx context.Context) ([]Library, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in CreatePolicy.
func (m *CreatePolicy) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId types.String `tfsdk:"policy_id"`
}

func (to *CreatePolicyResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePolicyResponse) {
}

func (to *CreatePolicyResponse) SyncFieldsDuringRead(ctx context.Context, from CreatePolicyResponse) {
}

func (m CreatePolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy_id"] = attrs["policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePolicyResponse
// only implements ToObjectValue() and Type().
func (m CreatePolicyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePolicyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type CreateResponse struct {
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id"`
}

func (to *CreateResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateResponse) {
}

func (to *CreateResponse) SyncFieldsDuringRead(ctx context.Context, from CreateResponse) {
}

func (m CreateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["script_id"] = attrs["script_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateResponse
// only implements ToObjectValue() and Type().
func (m CreateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": m.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"script_id": types.StringType,
		},
	}
}

type Created struct {
	Id types.String `tfsdk:"id"`
}

func (to *Created) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Created) {
}

func (to *Created) SyncFieldsDuringRead(ctx context.Context, from Created) {
}

func (m Created) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Created.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Created) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Created
// only implements ToObjectValue() and Type().
func (m Created) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Created) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CustomPolicyTag struct {
	// The key of the tag. - Must be unique among all custom tags of the same
	// policy - Cannot be “budget-policy-name”, “budget-policy-id” or
	// "budget-policy-resolution-result" - these tags are preserved.
	Key types.String `tfsdk:"key"`
	// The value of the tag.
	Value types.String `tfsdk:"value"`
}

func (to *CustomPolicyTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomPolicyTag) {
}

func (to *CustomPolicyTag) SyncFieldsDuringRead(ctx context.Context, from CustomPolicyTag) {
}

func (m CustomPolicyTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomPolicyTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CustomPolicyTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomPolicyTag
// only implements ToObjectValue() and Type().
func (m CustomPolicyTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomPolicyTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type DataPlaneEventDetails struct {
	EventType types.String `tfsdk:"event_type"`

	ExecutorFailures types.Int64 `tfsdk:"executor_failures"`

	HostId types.String `tfsdk:"host_id"`

	Timestamp types.Int64 `tfsdk:"timestamp"`
}

func (to *DataPlaneEventDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataPlaneEventDetails) {
}

func (to *DataPlaneEventDetails) SyncFieldsDuringRead(ctx context.Context, from DataPlaneEventDetails) {
}

func (m DataPlaneEventDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["event_type"] = attrs["event_type"].SetOptional()
	attrs["executor_failures"] = attrs["executor_failures"].SetOptional()
	attrs["host_id"] = attrs["host_id"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataPlaneEventDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DataPlaneEventDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneEventDetails
// only implements ToObjectValue() and Type().
func (m DataPlaneEventDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"event_type":        m.EventType,
			"executor_failures": m.ExecutorFailures,
			"host_id":           m.HostId,
			"timestamp":         m.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataPlaneEventDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"event_type":        types.StringType,
			"executor_failures": types.Int64Type,
			"host_id":           types.StringType,
			"timestamp":         types.Int64Type,
		},
	}
}

// A storage location in DBFS
type DbfsStorageInfo struct {
	// dbfs destination, e.g. `dbfs:/my/path`
	Destination types.String `tfsdk:"destination"`
}

func (to *DbfsStorageInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DbfsStorageInfo) {
}

func (to *DbfsStorageInfo) SyncFieldsDuringRead(ctx context.Context, from DbfsStorageInfo) {
}

func (m DbfsStorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination"] = attrs["destination"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DbfsStorageInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DbfsStorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DbfsStorageInfo
// only implements ToObjectValue() and Type().
func (m DbfsStorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": m.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DbfsStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type DefaultBaseEnvironment struct {
	BaseEnvironmentCache types.List `tfsdk:"base_environment_cache"`

	BaseEnvironmentType types.String `tfsdk:"base_environment_type"`

	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp"`

	CreatorUserId types.Int64 `tfsdk:"creator_user_id"`
	// Note: we made `environment` non-internal because we need to expose its
	// `client` field. All other fields should be treated as internal.
	Environment types.Object `tfsdk:"environment"`

	Filepath types.String `tfsdk:"filepath"`

	Id types.String `tfsdk:"id"`

	IsDefault types.Bool `tfsdk:"is_default"`

	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`

	LastUpdatedUserId types.Int64 `tfsdk:"last_updated_user_id"`

	Message types.String `tfsdk:"message"`

	Name types.String `tfsdk:"name"`

	PrincipalIds types.List `tfsdk:"principal_ids"`

	Status types.String `tfsdk:"status"`
}

func (to *DefaultBaseEnvironment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DefaultBaseEnvironment) {
	if !from.BaseEnvironmentCache.IsNull() && !from.BaseEnvironmentCache.IsUnknown() && to.BaseEnvironmentCache.IsNull() && len(from.BaseEnvironmentCache.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for BaseEnvironmentCache, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.BaseEnvironmentCache = from.BaseEnvironmentCache
	}
	if !from.Environment.IsNull() && !from.Environment.IsUnknown() {
		if toEnvironment, ok := to.GetEnvironment(ctx); ok {
			if fromEnvironment, ok := from.GetEnvironment(ctx); ok {
				// Recursively sync the fields of Environment
				toEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromEnvironment)
				to.SetEnvironment(ctx, toEnvironment)
			}
		}
	}
	if !from.PrincipalIds.IsNull() && !from.PrincipalIds.IsUnknown() && to.PrincipalIds.IsNull() && len(from.PrincipalIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrincipalIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrincipalIds = from.PrincipalIds
	}
}

func (to *DefaultBaseEnvironment) SyncFieldsDuringRead(ctx context.Context, from DefaultBaseEnvironment) {
	if !from.BaseEnvironmentCache.IsNull() && !from.BaseEnvironmentCache.IsUnknown() && to.BaseEnvironmentCache.IsNull() && len(from.BaseEnvironmentCache.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for BaseEnvironmentCache, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.BaseEnvironmentCache = from.BaseEnvironmentCache
	}
	if !from.Environment.IsNull() && !from.Environment.IsUnknown() {
		if toEnvironment, ok := to.GetEnvironment(ctx); ok {
			if fromEnvironment, ok := from.GetEnvironment(ctx); ok {
				toEnvironment.SyncFieldsDuringRead(ctx, fromEnvironment)
				to.SetEnvironment(ctx, toEnvironment)
			}
		}
	}
	if !from.PrincipalIds.IsNull() && !from.PrincipalIds.IsUnknown() && to.PrincipalIds.IsNull() && len(from.PrincipalIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrincipalIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrincipalIds = from.PrincipalIds
	}
}

func (m DefaultBaseEnvironment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["base_environment_cache"] = attrs["base_environment_cache"].SetOptional()
	attrs["base_environment_type"] = attrs["base_environment_type"].SetOptional()
	attrs["created_timestamp"] = attrs["created_timestamp"].SetComputed()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetComputed()
	attrs["environment"] = attrs["environment"].SetOptional()
	attrs["filepath"] = attrs["filepath"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["is_default"] = attrs["is_default"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetComputed()
	attrs["last_updated_user_id"] = attrs["last_updated_user_id"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["principal_ids"] = attrs["principal_ids"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultBaseEnvironment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DefaultBaseEnvironment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"base_environment_cache": reflect.TypeOf(DefaultBaseEnvironmentCache{}),
		"environment":            reflect.TypeOf(Environment{}),
		"principal_ids":          reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultBaseEnvironment
// only implements ToObjectValue() and Type().
func (m DefaultBaseEnvironment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"base_environment_cache": m.BaseEnvironmentCache,
			"base_environment_type":  m.BaseEnvironmentType,
			"created_timestamp":      m.CreatedTimestamp,
			"creator_user_id":        m.CreatorUserId,
			"environment":            m.Environment,
			"filepath":               m.Filepath,
			"id":                     m.Id,
			"is_default":             m.IsDefault,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"last_updated_user_id":   m.LastUpdatedUserId,
			"message":                m.Message,
			"name":                   m.Name,
			"principal_ids":          m.PrincipalIds,
			"status":                 m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DefaultBaseEnvironment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_environment_cache": basetypes.ListType{
				ElemType: DefaultBaseEnvironmentCache{}.Type(ctx),
			},
			"base_environment_type":  types.StringType,
			"created_timestamp":      types.Int64Type,
			"creator_user_id":        types.Int64Type,
			"environment":            Environment{}.Type(ctx),
			"filepath":               types.StringType,
			"id":                     types.StringType,
			"is_default":             types.BoolType,
			"last_updated_timestamp": types.Int64Type,
			"last_updated_user_id":   types.Int64Type,
			"message":                types.StringType,
			"name":                   types.StringType,
			"principal_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"status": types.StringType,
		},
	}
}

// GetBaseEnvironmentCache returns the value of the BaseEnvironmentCache field in DefaultBaseEnvironment as
// a slice of DefaultBaseEnvironmentCache values.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultBaseEnvironment) GetBaseEnvironmentCache(ctx context.Context) ([]DefaultBaseEnvironmentCache, bool) {
	if m.BaseEnvironmentCache.IsNull() || m.BaseEnvironmentCache.IsUnknown() {
		return nil, false
	}
	var v []DefaultBaseEnvironmentCache
	d := m.BaseEnvironmentCache.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBaseEnvironmentCache sets the value of the BaseEnvironmentCache field in DefaultBaseEnvironment.
func (m *DefaultBaseEnvironment) SetBaseEnvironmentCache(ctx context.Context, v []DefaultBaseEnvironmentCache) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["base_environment_cache"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.BaseEnvironmentCache = types.ListValueMust(t, vs)
}

// GetEnvironment returns the value of the Environment field in DefaultBaseEnvironment as
// a Environment value.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultBaseEnvironment) GetEnvironment(ctx context.Context) (Environment, bool) {
	var e Environment
	if m.Environment.IsNull() || m.Environment.IsUnknown() {
		return e, false
	}
	var v Environment
	d := m.Environment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironment sets the value of the Environment field in DefaultBaseEnvironment.
func (m *DefaultBaseEnvironment) SetEnvironment(ctx context.Context, v Environment) {
	vs := v.ToObjectValue(ctx)
	m.Environment = vs
}

// GetPrincipalIds returns the value of the PrincipalIds field in DefaultBaseEnvironment as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultBaseEnvironment) GetPrincipalIds(ctx context.Context) ([]types.Int64, bool) {
	if m.PrincipalIds.IsNull() || m.PrincipalIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.PrincipalIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrincipalIds sets the value of the PrincipalIds field in DefaultBaseEnvironment.
func (m *DefaultBaseEnvironment) SetPrincipalIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["principal_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrincipalIds = types.ListValueMust(t, vs)
}

type DefaultBaseEnvironmentCache struct {
	IndefiniteMaterializedEnvironment types.Object `tfsdk:"indefinite_materialized_environment"`

	MaterializedEnvironment types.Object `tfsdk:"materialized_environment"`

	Message types.String `tfsdk:"message"`

	Status types.String `tfsdk:"status"`
}

func (to *DefaultBaseEnvironmentCache) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DefaultBaseEnvironmentCache) {
	if !from.IndefiniteMaterializedEnvironment.IsNull() && !from.IndefiniteMaterializedEnvironment.IsUnknown() {
		if toIndefiniteMaterializedEnvironment, ok := to.GetIndefiniteMaterializedEnvironment(ctx); ok {
			if fromIndefiniteMaterializedEnvironment, ok := from.GetIndefiniteMaterializedEnvironment(ctx); ok {
				// Recursively sync the fields of IndefiniteMaterializedEnvironment
				toIndefiniteMaterializedEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromIndefiniteMaterializedEnvironment)
				to.SetIndefiniteMaterializedEnvironment(ctx, toIndefiniteMaterializedEnvironment)
			}
		}
	}
	if !from.MaterializedEnvironment.IsNull() && !from.MaterializedEnvironment.IsUnknown() {
		if toMaterializedEnvironment, ok := to.GetMaterializedEnvironment(ctx); ok {
			if fromMaterializedEnvironment, ok := from.GetMaterializedEnvironment(ctx); ok {
				// Recursively sync the fields of MaterializedEnvironment
				toMaterializedEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromMaterializedEnvironment)
				to.SetMaterializedEnvironment(ctx, toMaterializedEnvironment)
			}
		}
	}
}

func (to *DefaultBaseEnvironmentCache) SyncFieldsDuringRead(ctx context.Context, from DefaultBaseEnvironmentCache) {
	if !from.IndefiniteMaterializedEnvironment.IsNull() && !from.IndefiniteMaterializedEnvironment.IsUnknown() {
		if toIndefiniteMaterializedEnvironment, ok := to.GetIndefiniteMaterializedEnvironment(ctx); ok {
			if fromIndefiniteMaterializedEnvironment, ok := from.GetIndefiniteMaterializedEnvironment(ctx); ok {
				toIndefiniteMaterializedEnvironment.SyncFieldsDuringRead(ctx, fromIndefiniteMaterializedEnvironment)
				to.SetIndefiniteMaterializedEnvironment(ctx, toIndefiniteMaterializedEnvironment)
			}
		}
	}
	if !from.MaterializedEnvironment.IsNull() && !from.MaterializedEnvironment.IsUnknown() {
		if toMaterializedEnvironment, ok := to.GetMaterializedEnvironment(ctx); ok {
			if fromMaterializedEnvironment, ok := from.GetMaterializedEnvironment(ctx); ok {
				toMaterializedEnvironment.SyncFieldsDuringRead(ctx, fromMaterializedEnvironment)
				to.SetMaterializedEnvironment(ctx, toMaterializedEnvironment)
			}
		}
	}
}

func (m DefaultBaseEnvironmentCache) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["indefinite_materialized_environment"] = attrs["indefinite_materialized_environment"].SetOptional()
	attrs["materialized_environment"] = attrs["materialized_environment"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultBaseEnvironmentCache.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DefaultBaseEnvironmentCache) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"indefinite_materialized_environment": reflect.TypeOf(MaterializedEnvironment{}),
		"materialized_environment":            reflect.TypeOf(MaterializedEnvironment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultBaseEnvironmentCache
// only implements ToObjectValue() and Type().
func (m DefaultBaseEnvironmentCache) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"indefinite_materialized_environment": m.IndefiniteMaterializedEnvironment,
			"materialized_environment":            m.MaterializedEnvironment,
			"message":                             m.Message,
			"status":                              m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DefaultBaseEnvironmentCache) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"indefinite_materialized_environment": MaterializedEnvironment{}.Type(ctx),
			"materialized_environment":            MaterializedEnvironment{}.Type(ctx),
			"message":                             types.StringType,
			"status":                              types.StringType,
		},
	}
}

// GetIndefiniteMaterializedEnvironment returns the value of the IndefiniteMaterializedEnvironment field in DefaultBaseEnvironmentCache as
// a MaterializedEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultBaseEnvironmentCache) GetIndefiniteMaterializedEnvironment(ctx context.Context) (MaterializedEnvironment, bool) {
	var e MaterializedEnvironment
	if m.IndefiniteMaterializedEnvironment.IsNull() || m.IndefiniteMaterializedEnvironment.IsUnknown() {
		return e, false
	}
	var v MaterializedEnvironment
	d := m.IndefiniteMaterializedEnvironment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIndefiniteMaterializedEnvironment sets the value of the IndefiniteMaterializedEnvironment field in DefaultBaseEnvironmentCache.
func (m *DefaultBaseEnvironmentCache) SetIndefiniteMaterializedEnvironment(ctx context.Context, v MaterializedEnvironment) {
	vs := v.ToObjectValue(ctx)
	m.IndefiniteMaterializedEnvironment = vs
}

// GetMaterializedEnvironment returns the value of the MaterializedEnvironment field in DefaultBaseEnvironmentCache as
// a MaterializedEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultBaseEnvironmentCache) GetMaterializedEnvironment(ctx context.Context) (MaterializedEnvironment, bool) {
	var e MaterializedEnvironment
	if m.MaterializedEnvironment.IsNull() || m.MaterializedEnvironment.IsUnknown() {
		return e, false
	}
	var v MaterializedEnvironment
	d := m.MaterializedEnvironment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMaterializedEnvironment sets the value of the MaterializedEnvironment field in DefaultBaseEnvironmentCache.
func (m *DefaultBaseEnvironmentCache) SetMaterializedEnvironment(ctx context.Context, v MaterializedEnvironment) {
	vs := v.ToObjectValue(ctx)
	m.MaterializedEnvironment = vs
}

type DeleteCluster struct {
	// The cluster to be terminated.
	ClusterId types.String `tfsdk:"cluster_id"`
}

func (to *DeleteCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCluster) {
}

func (to *DeleteCluster) SyncFieldsDuringRead(ctx context.Context, from DeleteCluster) {
}

func (m DeleteCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCluster
// only implements ToObjectValue() and Type().
func (m DeleteCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type DeleteClusterResponse struct {
}

func (to *DeleteClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteClusterResponse) {
}

func (to *DeleteClusterResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteClusterResponse) {
}

func (m DeleteClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteClusterResponse
// only implements ToObjectValue() and Type().
func (m DeleteClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteDefaultBaseEnvironmentRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteDefaultBaseEnvironmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultBaseEnvironmentRequest) {
}

func (to *DeleteDefaultBaseEnvironmentRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultBaseEnvironmentRequest) {
}

func (m DeleteDefaultBaseEnvironmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDefaultBaseEnvironmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultBaseEnvironmentRequest
// only implements ToObjectValue() and Type().
func (m DeleteDefaultBaseEnvironmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultBaseEnvironmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-"`
}

func (to *DeleteGlobalInitScriptRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGlobalInitScriptRequest) {
}

func (to *DeleteGlobalInitScriptRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteGlobalInitScriptRequest) {
}

func (m DeleteGlobalInitScriptRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["script_id"] = attrs["script_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGlobalInitScriptRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteGlobalInitScriptRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGlobalInitScriptRequest
// only implements ToObjectValue() and Type().
func (m DeleteGlobalInitScriptRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": m.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGlobalInitScriptRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"script_id": types.StringType,
		},
	}
}

type DeleteInstancePool struct {
	// The instance pool to be terminated.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
}

func (to *DeleteInstancePool) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteInstancePool) {
}

func (to *DeleteInstancePool) SyncFieldsDuringRead(ctx context.Context, from DeleteInstancePool) {
}

func (m DeleteInstancePool) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstancePool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteInstancePool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstancePool
// only implements ToObjectValue() and Type().
func (m DeleteInstancePool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": m.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type DeleteInstancePoolResponse struct {
}

func (to *DeleteInstancePoolResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteInstancePoolResponse) {
}

func (to *DeleteInstancePoolResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteInstancePoolResponse) {
}

func (m DeleteInstancePoolResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstancePoolResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteInstancePoolResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstancePoolResponse
// only implements ToObjectValue() and Type().
func (m DeleteInstancePoolResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteInstancePoolResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	PolicyId types.String `tfsdk:"policy_id"`
}

func (to *DeletePolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePolicy) {
}

func (to *DeletePolicy) SyncFieldsDuringRead(ctx context.Context, from DeletePolicy) {
}

func (m DeletePolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePolicy
// only implements ToObjectValue() and Type().
func (m DeletePolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type DeletePolicyResponse struct {
}

func (to *DeletePolicyResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePolicyResponse) {
}

func (to *DeletePolicyResponse) SyncFieldsDuringRead(ctx context.Context, from DeletePolicyResponse) {
}

func (m DeletePolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePolicyResponse
// only implements ToObjectValue() and Type().
func (m DeletePolicyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePolicyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteResponse struct {
}

func (to *DeleteResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteResponse) {
}

func (to *DeleteResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteResponse) {
}

func (m DeleteResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse
// only implements ToObjectValue() and Type().
func (m DeleteResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DestroyContext struct {
	ClusterId types.String `tfsdk:"cluster_id"`

	ContextId types.String `tfsdk:"context_id"`
}

func (to *DestroyContext) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DestroyContext) {
}

func (to *DestroyContext) SyncFieldsDuringRead(ctx context.Context, from DestroyContext) {
}

func (m DestroyContext) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["context_id"] = attrs["context_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DestroyContext.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DestroyContext) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DestroyContext
// only implements ToObjectValue() and Type().
func (m DestroyContext) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
			"context_id": m.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DestroyContext) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"context_id": types.StringType,
		},
	}
}

type DestroyResponse struct {
}

func (to *DestroyResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DestroyResponse) {
}

func (to *DestroyResponse) SyncFieldsDuringRead(ctx context.Context, from DestroyResponse) {
}

func (m DestroyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DestroyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DestroyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DestroyResponse
// only implements ToObjectValue() and Type().
func (m DestroyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DestroyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Describes the disks that are launched for each instance in the spark cluster.
// For example, if the cluster has 3 instances, each instance is configured to
// launch 2 disks, 100 GiB each, then Databricks will launch a total of 6 disks,
// 100 GiB each, for this cluster.
type DiskSpec struct {
	// The number of disks launched for each instance: - This feature is only
	// enabled for supported node types. - Users can choose up to the limit of
	// the disks supported by the node type. - For node types with no OS disk,
	// at least one disk must be specified; otherwise, cluster creation will
	// fail.
	//
	// If disks are attached, Databricks will configure Spark to use only the
	// disks for scratch storage, because heterogenously sized scratch devices
	// can lead to inefficient disk utilization. If no disks are attached,
	// Databricks will configure Spark to use instance store disks.
	//
	// Note: If disks are specified, then the Spark configuration
	// `spark.local.dir` will be overridden.
	//
	// Disks will be mounted at: - For AWS: `/ebs0`, `/ebs1`, and etc. - For
	// Azure: `/remote_volume0`, `/remote_volume1`, and etc.
	DiskCount types.Int64 `tfsdk:"disk_count"`

	DiskIops types.Int64 `tfsdk:"disk_iops"`
	// The size of each disk (in GiB) launched for each instance. Values must
	// fall into the supported range for a particular instance type.
	//
	// For AWS: - General Purpose SSD: 100 - 4096 GiB - Throughput Optimized
	// HDD: 500 - 4096 GiB
	//
	// For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS (HDD): 1-
	// 1023 GiB
	DiskSize types.Int64 `tfsdk:"disk_size"`

	DiskThroughput types.Int64 `tfsdk:"disk_throughput"`
	// The type of disks that will be launched with this cluster.
	DiskType types.Object `tfsdk:"disk_type"`
}

func (to *DiskSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DiskSpec) {
	if !from.DiskType.IsNull() && !from.DiskType.IsUnknown() {
		if toDiskType, ok := to.GetDiskType(ctx); ok {
			if fromDiskType, ok := from.GetDiskType(ctx); ok {
				// Recursively sync the fields of DiskType
				toDiskType.SyncFieldsDuringCreateOrUpdate(ctx, fromDiskType)
				to.SetDiskType(ctx, toDiskType)
			}
		}
	}
}

func (to *DiskSpec) SyncFieldsDuringRead(ctx context.Context, from DiskSpec) {
	if !from.DiskType.IsNull() && !from.DiskType.IsUnknown() {
		if toDiskType, ok := to.GetDiskType(ctx); ok {
			if fromDiskType, ok := from.GetDiskType(ctx); ok {
				toDiskType.SyncFieldsDuringRead(ctx, fromDiskType)
				to.SetDiskType(ctx, toDiskType)
			}
		}
	}
}

func (m DiskSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disk_count"] = attrs["disk_count"].SetOptional()
	attrs["disk_iops"] = attrs["disk_iops"].SetOptional()
	attrs["disk_size"] = attrs["disk_size"].SetOptional()
	attrs["disk_throughput"] = attrs["disk_throughput"].SetOptional()
	attrs["disk_type"] = attrs["disk_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DiskSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DiskSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disk_type": reflect.TypeOf(DiskType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DiskSpec
// only implements ToObjectValue() and Type().
func (m DiskSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disk_count":      m.DiskCount,
			"disk_iops":       m.DiskIops,
			"disk_size":       m.DiskSize,
			"disk_throughput": m.DiskThroughput,
			"disk_type":       m.DiskType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DiskSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disk_count":      types.Int64Type,
			"disk_iops":       types.Int64Type,
			"disk_size":       types.Int64Type,
			"disk_throughput": types.Int64Type,
			"disk_type":       DiskType{}.Type(ctx),
		},
	}
}

// GetDiskType returns the value of the DiskType field in DiskSpec as
// a DiskType value.
// If the field is unknown or null, the boolean return value is false.
func (m *DiskSpec) GetDiskType(ctx context.Context) (DiskType, bool) {
	var e DiskType
	if m.DiskType.IsNull() || m.DiskType.IsUnknown() {
		return e, false
	}
	var v DiskType
	d := m.DiskType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDiskType sets the value of the DiskType field in DiskSpec.
func (m *DiskSpec) SetDiskType(ctx context.Context, v DiskType) {
	vs := v.ToObjectValue(ctx)
	m.DiskType = vs
}

// Describes the disk type.
type DiskType struct {
	AzureDiskVolumeType types.String `tfsdk:"azure_disk_volume_type"`

	EbsVolumeType types.String `tfsdk:"ebs_volume_type"`
}

func (to *DiskType) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DiskType) {
}

func (to *DiskType) SyncFieldsDuringRead(ctx context.Context, from DiskType) {
}

func (m DiskType) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["azure_disk_volume_type"] = attrs["azure_disk_volume_type"].SetOptional()
	attrs["ebs_volume_type"] = attrs["ebs_volume_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DiskType.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DiskType) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DiskType
// only implements ToObjectValue() and Type().
func (m DiskType) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"azure_disk_volume_type": m.AzureDiskVolumeType,
			"ebs_volume_type":        m.EbsVolumeType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DiskType) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"azure_disk_volume_type": types.StringType,
			"ebs_volume_type":        types.StringType,
		},
	}
}

type DockerBasicAuth struct {
	// Password of the user
	Password types.String `tfsdk:"password"`
	// Name of the user
	Username types.String `tfsdk:"username"`
}

func (to *DockerBasicAuth) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DockerBasicAuth) {
}

func (to *DockerBasicAuth) SyncFieldsDuringRead(ctx context.Context, from DockerBasicAuth) {
}

func (m DockerBasicAuth) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["password"] = attrs["password"].SetOptional()
	attrs["username"] = attrs["username"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DockerBasicAuth.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DockerBasicAuth) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DockerBasicAuth
// only implements ToObjectValue() and Type().
func (m DockerBasicAuth) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"password": m.Password,
			"username": m.Username,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DockerBasicAuth) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"password": types.StringType,
			"username": types.StringType,
		},
	}
}

type DockerImage struct {
	// Basic auth with username and password
	BasicAuth types.Object `tfsdk:"basic_auth"`
	// URL of the docker image.
	Url types.String `tfsdk:"url"`
}

func (to *DockerImage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DockerImage) {
	if !from.BasicAuth.IsNull() && !from.BasicAuth.IsUnknown() {
		if toBasicAuth, ok := to.GetBasicAuth(ctx); ok {
			if fromBasicAuth, ok := from.GetBasicAuth(ctx); ok {
				// Recursively sync the fields of BasicAuth
				toBasicAuth.SyncFieldsDuringCreateOrUpdate(ctx, fromBasicAuth)
				to.SetBasicAuth(ctx, toBasicAuth)
			}
		}
	}
}

func (to *DockerImage) SyncFieldsDuringRead(ctx context.Context, from DockerImage) {
	if !from.BasicAuth.IsNull() && !from.BasicAuth.IsUnknown() {
		if toBasicAuth, ok := to.GetBasicAuth(ctx); ok {
			if fromBasicAuth, ok := from.GetBasicAuth(ctx); ok {
				toBasicAuth.SyncFieldsDuringRead(ctx, fromBasicAuth)
				to.SetBasicAuth(ctx, toBasicAuth)
			}
		}
	}
}

func (m DockerImage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["basic_auth"] = attrs["basic_auth"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DockerImage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DockerImage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"basic_auth": reflect.TypeOf(DockerBasicAuth{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DockerImage
// only implements ToObjectValue() and Type().
func (m DockerImage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"basic_auth": m.BasicAuth,
			"url":        m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DockerImage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"basic_auth": DockerBasicAuth{}.Type(ctx),
			"url":        types.StringType,
		},
	}
}

// GetBasicAuth returns the value of the BasicAuth field in DockerImage as
// a DockerBasicAuth value.
// If the field is unknown or null, the boolean return value is false.
func (m *DockerImage) GetBasicAuth(ctx context.Context) (DockerBasicAuth, bool) {
	var e DockerBasicAuth
	if m.BasicAuth.IsNull() || m.BasicAuth.IsUnknown() {
		return e, false
	}
	var v DockerBasicAuth
	d := m.BasicAuth.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBasicAuth sets the value of the BasicAuth field in DockerImage.
func (m *DockerImage) SetBasicAuth(ctx context.Context, v DockerBasicAuth) {
	vs := v.ToObjectValue(ctx)
	m.BasicAuth = vs
}

type EditCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// ID of the cluster
	ClusterId types.String `tfsdk:"cluster_id"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.Object `tfsdk:"cluster_log_conf"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName types.String `tfsdk:"cluster_name"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags"`

	DataSecurityMode types.String `tfsdk:"data_security_mode"`
	// Custom docker image BYOC
	DockerImage types.Object `tfsdk:"docker_image"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode types.Bool `tfsdk:"is_single_node"`

	Kind types.String `tfsdk:"kind"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars types.Map `tfsdk:"spark_env_vars"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime types.Bool `tfsdk:"use_ml_runtime"`

	WorkloadType types.Object `tfsdk:"workload_type"`
}

func (to *EditCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditCluster) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				// Recursively sync the fields of Autoscale
				toAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				// Recursively sync the fields of ClusterLogConf
				toClusterLogConf.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				// Recursively sync the fields of DockerImage
				toDockerImage.SyncFieldsDuringCreateOrUpdate(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				// Recursively sync the fields of WorkloadType
				toWorkloadType.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (to *EditCluster) SyncFieldsDuringRead(ctx context.Context, from EditCluster) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				toAutoscale.SyncFieldsDuringRead(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				toClusterLogConf.SyncFieldsDuringRead(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				toDockerImage.SyncFieldsDuringRead(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				toWorkloadType.SyncFieldsDuringRead(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (m EditCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_policy_default_values"] = attrs["apply_policy_default_values"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["is_single_node"] = attrs["is_single_node"].SetOptional()
	attrs["kind"] = attrs["kind"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["num_workers"] = attrs["num_workers"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["runtime_engine"] = attrs["runtime_engine"].SetOptional()
	attrs["single_user_name"] = attrs["single_user_name"].SetOptional()
	attrs["spark_conf"] = attrs["spark_conf"].SetOptional()
	attrs["spark_env_vars"] = attrs["spark_env_vars"].SetOptional()
	attrs["spark_version"] = attrs["spark_version"].SetRequired()
	attrs["ssh_public_keys"] = attrs["ssh_public_keys"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()
	attrs["use_ml_runtime"] = attrs["use_ml_runtime"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(AutoScale{}),
		"aws_attributes":   reflect.TypeOf(AwsAttributes{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditCluster
// only implements ToObjectValue() and Type().
func (m EditCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_policy_default_values":    m.ApplyPolicyDefaultValues,
			"autoscale":                      m.Autoscale,
			"autotermination_minutes":        m.AutoterminationMinutes,
			"aws_attributes":                 m.AwsAttributes,
			"azure_attributes":               m.AzureAttributes,
			"cluster_id":                     m.ClusterId,
			"cluster_log_conf":               m.ClusterLogConf,
			"cluster_name":                   m.ClusterName,
			"custom_tags":                    m.CustomTags,
			"data_security_mode":             m.DataSecurityMode,
			"docker_image":                   m.DockerImage,
			"driver_instance_pool_id":        m.DriverInstancePoolId,
			"driver_node_type_id":            m.DriverNodeTypeId,
			"enable_elastic_disk":            m.EnableElasticDisk,
			"enable_local_disk_encryption":   m.EnableLocalDiskEncryption,
			"gcp_attributes":                 m.GcpAttributes,
			"init_scripts":                   m.InitScripts,
			"instance_pool_id":               m.InstancePoolId,
			"is_single_node":                 m.IsSingleNode,
			"kind":                           m.Kind,
			"node_type_id":                   m.NodeTypeId,
			"num_workers":                    m.NumWorkers,
			"policy_id":                      m.PolicyId,
			"remote_disk_throughput":         m.RemoteDiskThroughput,
			"runtime_engine":                 m.RuntimeEngine,
			"single_user_name":               m.SingleUserName,
			"spark_conf":                     m.SparkConf,
			"spark_env_vars":                 m.SparkEnvVars,
			"spark_version":                  m.SparkVersion,
			"ssh_public_keys":                m.SshPublicKeys,
			"total_initial_remote_disk_size": m.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 m.UseMlRuntime,
			"workload_type":                  m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EditCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale":                   AutoScale{}.Type(ctx),
			"autotermination_minutes":     types.Int64Type,
			"aws_attributes":              AwsAttributes{}.Type(ctx),
			"azure_attributes":            AzureAttributes{}.Type(ctx),
			"cluster_id":                  types.StringType,
			"cluster_log_conf":            ClusterLogConf{}.Type(ctx),
			"cluster_name":                types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode":           types.StringType,
			"docker_image":                 DockerImage{}.Type(ctx),
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes":               GcpAttributes{}.Type(ctx),
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id":       types.StringType,
			"is_single_node":         types.BoolType,
			"kind":                   types.StringType,
			"node_type_id":           types.StringType,
			"num_workers":            types.Int64Type,
			"policy_id":              types.StringType,
			"remote_disk_throughput": types.Int64Type,
			"runtime_engine":         types.StringType,
			"single_user_name":       types.StringType,
			"spark_conf": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_env_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_version": types.StringType,
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"total_initial_remote_disk_size": types.Int64Type,
			"use_ml_runtime":                 types.BoolType,
			"workload_type":                  WorkloadType{}.Type(ctx),
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in EditCluster as
// a AutoScale value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v AutoScale
	d := m.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoscale sets the value of the Autoscale field in EditCluster.
func (m *EditCluster) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	m.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in EditCluster as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v AwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in EditCluster.
func (m *EditCluster) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in EditCluster as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v AzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in EditCluster.
func (m *EditCluster) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in EditCluster as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if m.ClusterLogConf.IsNull() || m.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v ClusterLogConf
	d := m.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in EditCluster.
func (m *EditCluster) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	m.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in EditCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EditCluster.
func (m *EditCluster) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in EditCluster as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if m.DockerImage.IsNull() || m.DockerImage.IsUnknown() {
		return e, false
	}
	var v DockerImage
	d := m.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDockerImage sets the value of the DockerImage field in EditCluster.
func (m *EditCluster) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	m.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in EditCluster as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v GcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in EditCluster.
func (m *EditCluster) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in EditCluster as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := m.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in EditCluster.
func (m *EditCluster) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in EditCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkConf.IsNull() || m.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in EditCluster.
func (m *EditCluster) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in EditCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkEnvVars.IsNull() || m.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in EditCluster.
func (m *EditCluster) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in EditCluster as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if m.SshPublicKeys.IsNull() || m.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in EditCluster.
func (m *EditCluster) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in EditCluster as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditCluster) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if m.WorkloadType.IsNull() || m.WorkloadType.IsUnknown() {
		return e, false
	}
	var v WorkloadType
	d := m.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkloadType sets the value of the WorkloadType field in EditCluster.
func (m *EditCluster) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	m.WorkloadType = vs
}

type EditClusterResponse struct {
}

func (to *EditClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditClusterResponse) {
}

func (to *EditClusterResponse) SyncFieldsDuringRead(ctx context.Context, from EditClusterResponse) {
}

func (m EditClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditClusterResponse
// only implements ToObjectValue() and Type().
func (m EditClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EditClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditInstancePool struct {
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags types.Map `tfsdk:"custom_tags"`
	// For pools with node type flexibility (Fleet-V2), whether auto generated
	// alternate node type ids are enabled. This field should not be true if
	// node_type_flexibility is set.
	EnableAutoAlternateNodeTypes types.Bool `tfsdk:"enable_auto_alternate_node_types"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes"`
	// Instance pool ID
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances"`
	// For pools with node type flexibility (Fleet-V2), this object contains the
	// information about the alternate node type ids to use when attempting to
	// launch a cluster if the node type id is not available. This field should
	// not be set if enable_auto_alternate_node_types is true.
	NodeTypeFlexibility types.Object `tfsdk:"node_type_flexibility"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
}

func (to *EditInstancePool) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditInstancePool) {
	if !from.NodeTypeFlexibility.IsNull() && !from.NodeTypeFlexibility.IsUnknown() {
		if toNodeTypeFlexibility, ok := to.GetNodeTypeFlexibility(ctx); ok {
			if fromNodeTypeFlexibility, ok := from.GetNodeTypeFlexibility(ctx); ok {
				// Recursively sync the fields of NodeTypeFlexibility
				toNodeTypeFlexibility.SyncFieldsDuringCreateOrUpdate(ctx, fromNodeTypeFlexibility)
				to.SetNodeTypeFlexibility(ctx, toNodeTypeFlexibility)
			}
		}
	}
}

func (to *EditInstancePool) SyncFieldsDuringRead(ctx context.Context, from EditInstancePool) {
	if !from.NodeTypeFlexibility.IsNull() && !from.NodeTypeFlexibility.IsUnknown() {
		if toNodeTypeFlexibility, ok := to.GetNodeTypeFlexibility(ctx); ok {
			if fromNodeTypeFlexibility, ok := from.GetNodeTypeFlexibility(ctx); ok {
				toNodeTypeFlexibility.SyncFieldsDuringRead(ctx, fromNodeTypeFlexibility)
				to.SetNodeTypeFlexibility(ctx, toNodeTypeFlexibility)
			}
		}
	}
}

func (m EditInstancePool) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["enable_auto_alternate_node_types"] = attrs["enable_auto_alternate_node_types"].SetOptional()
	attrs["idle_instance_autotermination_minutes"] = attrs["idle_instance_autotermination_minutes"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()
	attrs["instance_pool_name"] = attrs["instance_pool_name"].SetRequired()
	attrs["max_capacity"] = attrs["max_capacity"].SetOptional()
	attrs["min_idle_instances"] = attrs["min_idle_instances"].SetOptional()
	attrs["node_type_flexibility"] = attrs["node_type_flexibility"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetRequired()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditInstancePool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditInstancePool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags":           reflect.TypeOf(types.String{}),
		"node_type_flexibility": reflect.TypeOf(NodeTypeFlexibility{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditInstancePool
// only implements ToObjectValue() and Type().
func (m EditInstancePool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags":                           m.CustomTags,
			"enable_auto_alternate_node_types":      m.EnableAutoAlternateNodeTypes,
			"idle_instance_autotermination_minutes": m.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      m.InstancePoolId,
			"instance_pool_name":                    m.InstancePoolName,
			"max_capacity":                          m.MaxCapacity,
			"min_idle_instances":                    m.MinIdleInstances,
			"node_type_flexibility":                 m.NodeTypeFlexibility,
			"node_type_id":                          m.NodeTypeId,
			"remote_disk_throughput":                m.RemoteDiskThroughput,
			"total_initial_remote_disk_size":        m.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EditInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"enable_auto_alternate_node_types":      types.BoolType,
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_id":                      types.StringType,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
			"node_type_flexibility":                 NodeTypeFlexibility{}.Type(ctx),
			"node_type_id":                          types.StringType,
			"remote_disk_throughput":                types.Int64Type,
			"total_initial_remote_disk_size":        types.Int64Type,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in EditInstancePool as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditInstancePool) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EditInstancePool.
func (m *EditInstancePool) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetNodeTypeFlexibility returns the value of the NodeTypeFlexibility field in EditInstancePool as
// a NodeTypeFlexibility value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditInstancePool) GetNodeTypeFlexibility(ctx context.Context) (NodeTypeFlexibility, bool) {
	var e NodeTypeFlexibility
	if m.NodeTypeFlexibility.IsNull() || m.NodeTypeFlexibility.IsUnknown() {
		return e, false
	}
	var v NodeTypeFlexibility
	d := m.NodeTypeFlexibility.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeTypeFlexibility sets the value of the NodeTypeFlexibility field in EditInstancePool.
func (m *EditInstancePool) SetNodeTypeFlexibility(ctx context.Context, v NodeTypeFlexibility) {
	vs := v.ToObjectValue(ctx)
	m.NodeTypeFlexibility = vs
}

type EditInstancePoolResponse struct {
}

func (to *EditInstancePoolResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditInstancePoolResponse) {
}

func (to *EditInstancePoolResponse) SyncFieldsDuringRead(ctx context.Context, from EditInstancePoolResponse) {
}

func (m EditInstancePoolResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditInstancePoolResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditInstancePoolResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditInstancePoolResponse
// only implements ToObjectValue() and Type().
func (m EditInstancePoolResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EditInstancePoolResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditPolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries types.List `tfsdk:"libraries"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId types.String `tfsdk:"policy_family_id"`
	// The ID of the policy to update.
	PolicyId types.String `tfsdk:"policy_id"`
}

func (to *EditPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditPolicy) {
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
}

func (to *EditPolicy) SyncFieldsDuringRead(ctx context.Context, from EditPolicy) {
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
}

func (m EditPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["definition"] = attrs["definition"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["libraries"] = attrs["libraries"].SetOptional()
	attrs["max_clusters_per_user"] = attrs["max_clusters_per_user"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["policy_family_definition_overrides"] = attrs["policy_family_definition_overrides"].SetOptional()
	attrs["policy_family_id"] = attrs["policy_family_id"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPolicy
// only implements ToObjectValue() and Type().
func (m EditPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"definition":                         m.Definition,
			"description":                        m.Description,
			"libraries":                          m.Libraries,
			"max_clusters_per_user":              m.MaxClustersPerUser,
			"name":                               m.Name,
			"policy_family_definition_overrides": m.PolicyFamilyDefinitionOverrides,
			"policy_family_id":                   m.PolicyFamilyId,
			"policy_id":                          m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EditPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"definition":  types.StringType,
			"description": types.StringType,
			"libraries": basetypes.ListType{
				ElemType: Library{}.Type(ctx),
			},
			"max_clusters_per_user":              types.Int64Type,
			"name":                               types.StringType,
			"policy_family_definition_overrides": types.StringType,
			"policy_family_id":                   types.StringType,
			"policy_id":                          types.StringType,
		},
	}
}

// GetLibraries returns the value of the Libraries field in EditPolicy as
// a slice of Library values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPolicy) GetLibraries(ctx context.Context) ([]Library, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in EditPolicy.
func (m *EditPolicy) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

type EditPolicyResponse struct {
}

func (to *EditPolicyResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditPolicyResponse) {
}

func (to *EditPolicyResponse) SyncFieldsDuringRead(ctx context.Context, from EditPolicyResponse) {
}

func (m EditPolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditPolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPolicyResponse
// only implements ToObjectValue() and Type().
func (m EditPolicyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EditPolicyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditResponse struct {
}

func (to *EditResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditResponse) {
}

func (to *EditResponse) SyncFieldsDuringRead(ctx context.Context, from EditResponse) {
}

func (m EditResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditResponse
// only implements ToObjectValue() and Type().
func (m EditResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EditResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EnforceClusterComplianceRequest struct {
	// The ID of the cluster you want to enforce policy compliance on.
	ClusterId types.String `tfsdk:"cluster_id"`
	// If set, previews the changes that would be made to a cluster to enforce
	// compliance but does not update the cluster.
	ValidateOnly types.Bool `tfsdk:"validate_only"`
}

func (to *EnforceClusterComplianceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnforceClusterComplianceRequest) {
}

func (to *EnforceClusterComplianceRequest) SyncFieldsDuringRead(ctx context.Context, from EnforceClusterComplianceRequest) {
}

func (m EnforceClusterComplianceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["validate_only"] = attrs["validate_only"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforceClusterComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EnforceClusterComplianceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforceClusterComplianceRequest
// only implements ToObjectValue() and Type().
func (m EnforceClusterComplianceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":    m.ClusterId,
			"validate_only": m.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnforceClusterComplianceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":    types.StringType,
			"validate_only": types.BoolType,
		},
	}
}

type EnforceClusterComplianceResponse struct {
	// A list of changes that have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	Changes types.List `tfsdk:"changes"`
	// Whether any changes have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	HasChanges types.Bool `tfsdk:"has_changes"`
}

func (to *EnforceClusterComplianceResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnforceClusterComplianceResponse) {
	if !from.Changes.IsNull() && !from.Changes.IsUnknown() && to.Changes.IsNull() && len(from.Changes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Changes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Changes = from.Changes
	}
}

func (to *EnforceClusterComplianceResponse) SyncFieldsDuringRead(ctx context.Context, from EnforceClusterComplianceResponse) {
	if !from.Changes.IsNull() && !from.Changes.IsUnknown() && to.Changes.IsNull() && len(from.Changes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Changes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Changes = from.Changes
	}
}

func (m EnforceClusterComplianceResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["changes"] = attrs["changes"].SetOptional()
	attrs["has_changes"] = attrs["has_changes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforceClusterComplianceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EnforceClusterComplianceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(ClusterSettingsChange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforceClusterComplianceResponse
// only implements ToObjectValue() and Type().
func (m EnforceClusterComplianceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes":     m.Changes,
			"has_changes": m.HasChanges,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnforceClusterComplianceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: ClusterSettingsChange{}.Type(ctx),
			},
			"has_changes": types.BoolType,
		},
	}
}

// GetChanges returns the value of the Changes field in EnforceClusterComplianceResponse as
// a slice of ClusterSettingsChange values.
// If the field is unknown or null, the boolean return value is false.
func (m *EnforceClusterComplianceResponse) GetChanges(ctx context.Context) ([]ClusterSettingsChange, bool) {
	if m.Changes.IsNull() || m.Changes.IsUnknown() {
		return nil, false
	}
	var v []ClusterSettingsChange
	d := m.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in EnforceClusterComplianceResponse.
func (m *EnforceClusterComplianceResponse) SetChanges(ctx context.Context, v []ClusterSettingsChange) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Changes = types.ListValueMust(t, vs)
}

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and DLT's environment for classic
// and serverless pipelines. In this minimal environment spec, only pip
// dependencies are supported.
type Environment struct {
	// Use `environment_version` instead.
	Client types.String `tfsdk:"client"`
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a valid pip requirements file line per
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/.
	// Allowed dependencies include a requirement specifier, an archive URL, a
	// local project path (such as WSFS or UC Volumes in Databricks), or a VCS
	// project URL.
	Dependencies types.List `tfsdk:"dependencies"`
	// Required. Environment version used by the environment. Each version comes
	// with a specific Python version and a set of Python packages. The version
	// is a string, consisting of an integer.
	EnvironmentVersion types.String `tfsdk:"environment_version"`
	// Use `java_dependencies` instead.
	JarDependencies types.List `tfsdk:"jar_dependencies"`
	// List of jar dependencies, should be string representing volume paths. For
	// example: `/Volumes/path/to/test.jar`.
	JavaDependencies types.List `tfsdk:"java_dependencies"`
}

func (to *Environment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Environment) {
	if !from.Dependencies.IsNull() && !from.Dependencies.IsUnknown() && to.Dependencies.IsNull() && len(from.Dependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Dependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Dependencies = from.Dependencies
	}
	if !from.JarDependencies.IsNull() && !from.JarDependencies.IsUnknown() && to.JarDependencies.IsNull() && len(from.JarDependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for JarDependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.JarDependencies = from.JarDependencies
	}
	if !from.JavaDependencies.IsNull() && !from.JavaDependencies.IsUnknown() && to.JavaDependencies.IsNull() && len(from.JavaDependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for JavaDependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.JavaDependencies = from.JavaDependencies
	}
}

func (to *Environment) SyncFieldsDuringRead(ctx context.Context, from Environment) {
	if !from.Dependencies.IsNull() && !from.Dependencies.IsUnknown() && to.Dependencies.IsNull() && len(from.Dependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Dependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Dependencies = from.Dependencies
	}
	if !from.JarDependencies.IsNull() && !from.JarDependencies.IsUnknown() && to.JarDependencies.IsNull() && len(from.JarDependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for JarDependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.JarDependencies = from.JarDependencies
	}
	if !from.JavaDependencies.IsNull() && !from.JavaDependencies.IsUnknown() && to.JavaDependencies.IsNull() && len(from.JavaDependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for JavaDependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.JavaDependencies = from.JavaDependencies
	}
}

func (m Environment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["client"] = attrs["client"].SetOptional()
	attrs["dependencies"] = attrs["dependencies"].SetOptional()
	attrs["environment_version"] = attrs["environment_version"].SetOptional()
	attrs["jar_dependencies"] = attrs["jar_dependencies"].SetOptional()
	attrs["java_dependencies"] = attrs["java_dependencies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Environment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Environment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies":      reflect.TypeOf(types.String{}),
		"jar_dependencies":  reflect.TypeOf(types.String{}),
		"java_dependencies": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Environment
// only implements ToObjectValue() and Type().
func (m Environment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client":              m.Client,
			"dependencies":        m.Dependencies,
			"environment_version": m.EnvironmentVersion,
			"jar_dependencies":    m.JarDependencies,
			"java_dependencies":   m.JavaDependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Environment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"client": types.StringType,
			"dependencies": basetypes.ListType{
				ElemType: types.StringType,
			},
			"environment_version": types.StringType,
			"jar_dependencies": basetypes.ListType{
				ElemType: types.StringType,
			},
			"java_dependencies": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDependencies returns the value of the Dependencies field in Environment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Environment) GetDependencies(ctx context.Context) ([]types.String, bool) {
	if m.Dependencies.IsNull() || m.Dependencies.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Dependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependencies sets the value of the Dependencies field in Environment.
func (m *Environment) SetDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Dependencies = types.ListValueMust(t, vs)
}

// GetJarDependencies returns the value of the JarDependencies field in Environment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Environment) GetJarDependencies(ctx context.Context) ([]types.String, bool) {
	if m.JarDependencies.IsNull() || m.JarDependencies.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.JarDependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarDependencies sets the value of the JarDependencies field in Environment.
func (m *Environment) SetJarDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["jar_dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.JarDependencies = types.ListValueMust(t, vs)
}

// GetJavaDependencies returns the value of the JavaDependencies field in Environment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Environment) GetJavaDependencies(ctx context.Context) ([]types.String, bool) {
	if m.JavaDependencies.IsNull() || m.JavaDependencies.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.JavaDependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJavaDependencies sets the value of the JavaDependencies field in Environment.
func (m *Environment) SetJavaDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["java_dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.JavaDependencies = types.ListValueMust(t, vs)
}

type EventDetails struct {
	// * For created clusters, the attributes of the cluster. * For edited
	// clusters, the new attributes of the cluster.
	Attributes types.Object `tfsdk:"attributes"`
	// The cause of a change in target size.
	Cause types.String `tfsdk:"cause"`
	// The actual cluster size that was set in the cluster creation or edit.
	ClusterSize types.Object `tfsdk:"cluster_size"`
	// The current number of vCPUs in the cluster.
	CurrentNumVcpus types.Int64 `tfsdk:"current_num_vcpus"`
	// The current number of nodes in the cluster.
	CurrentNumWorkers types.Int64 `tfsdk:"current_num_workers"`

	DidNotExpandReason types.String `tfsdk:"did_not_expand_reason"`
	// Current disk size in bytes
	DiskSize types.Int64 `tfsdk:"disk_size"`
	// More details about the change in driver's state
	DriverStateMessage types.String `tfsdk:"driver_state_message"`
	// Whether or not a blocklisted node should be terminated. For
	// ClusterEventType NODE_BLACKLISTED.
	EnableTerminationForNodeBlocklisted types.Bool `tfsdk:"enable_termination_for_node_blocklisted"`

	FreeSpace types.Int64 `tfsdk:"free_space"`
	// List of global and cluster init scripts associated with this cluster
	// event.
	InitScripts types.Object `tfsdk:"init_scripts"`
	// Instance Id where the event originated from
	InstanceId types.String `tfsdk:"instance_id"`
	// Unique identifier of the specific job run associated with this cluster
	// event * For clusters created for jobs, this will be the same as the
	// cluster name
	JobRunName types.String `tfsdk:"job_run_name"`
	// The cluster attributes before a cluster was edited.
	PreviousAttributes types.Object `tfsdk:"previous_attributes"`
	// The size of the cluster before an edit or resize.
	PreviousClusterSize types.Object `tfsdk:"previous_cluster_size"`
	// Previous disk size in bytes
	PreviousDiskSize types.Int64 `tfsdk:"previous_disk_size"`
	// A termination reason: * On a TERMINATED event, this is the reason of the
	// termination. * On a RESIZE_COMPLETE event, this indicates the reason that
	// we failed to acquire some nodes.
	Reason types.Object `tfsdk:"reason"`
	// The targeted number of vCPUs in the cluster.
	TargetNumVcpus types.Int64 `tfsdk:"target_num_vcpus"`
	// The targeted number of nodes in the cluster.
	TargetNumWorkers types.Int64 `tfsdk:"target_num_workers"`
	// The user that caused the event to occur. (Empty if it was done by the
	// control plane.)
	User types.String `tfsdk:"user"`
}

func (to *EventDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EventDetails) {
	if !from.Attributes.IsNull() && !from.Attributes.IsUnknown() {
		if toAttributes, ok := to.GetAttributes(ctx); ok {
			if fromAttributes, ok := from.GetAttributes(ctx); ok {
				// Recursively sync the fields of Attributes
				toAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAttributes)
				to.SetAttributes(ctx, toAttributes)
			}
		}
	}
	if !from.ClusterSize.IsNull() && !from.ClusterSize.IsUnknown() {
		if toClusterSize, ok := to.GetClusterSize(ctx); ok {
			if fromClusterSize, ok := from.GetClusterSize(ctx); ok {
				// Recursively sync the fields of ClusterSize
				toClusterSize.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterSize)
				to.SetClusterSize(ctx, toClusterSize)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() {
		if toInitScripts, ok := to.GetInitScripts(ctx); ok {
			if fromInitScripts, ok := from.GetInitScripts(ctx); ok {
				// Recursively sync the fields of InitScripts
				toInitScripts.SyncFieldsDuringCreateOrUpdate(ctx, fromInitScripts)
				to.SetInitScripts(ctx, toInitScripts)
			}
		}
	}
	if !from.PreviousAttributes.IsNull() && !from.PreviousAttributes.IsUnknown() {
		if toPreviousAttributes, ok := to.GetPreviousAttributes(ctx); ok {
			if fromPreviousAttributes, ok := from.GetPreviousAttributes(ctx); ok {
				// Recursively sync the fields of PreviousAttributes
				toPreviousAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromPreviousAttributes)
				to.SetPreviousAttributes(ctx, toPreviousAttributes)
			}
		}
	}
	if !from.PreviousClusterSize.IsNull() && !from.PreviousClusterSize.IsUnknown() {
		if toPreviousClusterSize, ok := to.GetPreviousClusterSize(ctx); ok {
			if fromPreviousClusterSize, ok := from.GetPreviousClusterSize(ctx); ok {
				// Recursively sync the fields of PreviousClusterSize
				toPreviousClusterSize.SyncFieldsDuringCreateOrUpdate(ctx, fromPreviousClusterSize)
				to.SetPreviousClusterSize(ctx, toPreviousClusterSize)
			}
		}
	}
	if !from.Reason.IsNull() && !from.Reason.IsUnknown() {
		if toReason, ok := to.GetReason(ctx); ok {
			if fromReason, ok := from.GetReason(ctx); ok {
				// Recursively sync the fields of Reason
				toReason.SyncFieldsDuringCreateOrUpdate(ctx, fromReason)
				to.SetReason(ctx, toReason)
			}
		}
	}
}

func (to *EventDetails) SyncFieldsDuringRead(ctx context.Context, from EventDetails) {
	if !from.Attributes.IsNull() && !from.Attributes.IsUnknown() {
		if toAttributes, ok := to.GetAttributes(ctx); ok {
			if fromAttributes, ok := from.GetAttributes(ctx); ok {
				toAttributes.SyncFieldsDuringRead(ctx, fromAttributes)
				to.SetAttributes(ctx, toAttributes)
			}
		}
	}
	if !from.ClusterSize.IsNull() && !from.ClusterSize.IsUnknown() {
		if toClusterSize, ok := to.GetClusterSize(ctx); ok {
			if fromClusterSize, ok := from.GetClusterSize(ctx); ok {
				toClusterSize.SyncFieldsDuringRead(ctx, fromClusterSize)
				to.SetClusterSize(ctx, toClusterSize)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() {
		if toInitScripts, ok := to.GetInitScripts(ctx); ok {
			if fromInitScripts, ok := from.GetInitScripts(ctx); ok {
				toInitScripts.SyncFieldsDuringRead(ctx, fromInitScripts)
				to.SetInitScripts(ctx, toInitScripts)
			}
		}
	}
	if !from.PreviousAttributes.IsNull() && !from.PreviousAttributes.IsUnknown() {
		if toPreviousAttributes, ok := to.GetPreviousAttributes(ctx); ok {
			if fromPreviousAttributes, ok := from.GetPreviousAttributes(ctx); ok {
				toPreviousAttributes.SyncFieldsDuringRead(ctx, fromPreviousAttributes)
				to.SetPreviousAttributes(ctx, toPreviousAttributes)
			}
		}
	}
	if !from.PreviousClusterSize.IsNull() && !from.PreviousClusterSize.IsUnknown() {
		if toPreviousClusterSize, ok := to.GetPreviousClusterSize(ctx); ok {
			if fromPreviousClusterSize, ok := from.GetPreviousClusterSize(ctx); ok {
				toPreviousClusterSize.SyncFieldsDuringRead(ctx, fromPreviousClusterSize)
				to.SetPreviousClusterSize(ctx, toPreviousClusterSize)
			}
		}
	}
	if !from.Reason.IsNull() && !from.Reason.IsUnknown() {
		if toReason, ok := to.GetReason(ctx); ok {
			if fromReason, ok := from.GetReason(ctx); ok {
				toReason.SyncFieldsDuringRead(ctx, fromReason)
				to.SetReason(ctx, toReason)
			}
		}
	}
}

func (m EventDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["attributes"] = attrs["attributes"].SetOptional()
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["current_num_vcpus"] = attrs["current_num_vcpus"].SetOptional()
	attrs["current_num_workers"] = attrs["current_num_workers"].SetOptional()
	attrs["did_not_expand_reason"] = attrs["did_not_expand_reason"].SetOptional()
	attrs["disk_size"] = attrs["disk_size"].SetOptional()
	attrs["driver_state_message"] = attrs["driver_state_message"].SetOptional()
	attrs["enable_termination_for_node_blocklisted"] = attrs["enable_termination_for_node_blocklisted"].SetOptional()
	attrs["free_space"] = attrs["free_space"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["instance_id"] = attrs["instance_id"].SetOptional()
	attrs["job_run_name"] = attrs["job_run_name"].SetOptional()
	attrs["previous_attributes"] = attrs["previous_attributes"].SetOptional()
	attrs["previous_cluster_size"] = attrs["previous_cluster_size"].SetOptional()
	attrs["previous_disk_size"] = attrs["previous_disk_size"].SetOptional()
	attrs["reason"] = attrs["reason"].SetOptional()
	attrs["target_num_vcpus"] = attrs["target_num_vcpus"].SetOptional()
	attrs["target_num_workers"] = attrs["target_num_workers"].SetOptional()
	attrs["user"] = attrs["user"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EventDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EventDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attributes":            reflect.TypeOf(ClusterAttributes{}),
		"cluster_size":          reflect.TypeOf(ClusterSize{}),
		"init_scripts":          reflect.TypeOf(InitScriptEventDetails{}),
		"previous_attributes":   reflect.TypeOf(ClusterAttributes{}),
		"previous_cluster_size": reflect.TypeOf(ClusterSize{}),
		"reason":                reflect.TypeOf(TerminationReason{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EventDetails
// only implements ToObjectValue() and Type().
func (m EventDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":            m.Attributes,
			"cause":                 m.Cause,
			"cluster_size":          m.ClusterSize,
			"current_num_vcpus":     m.CurrentNumVcpus,
			"current_num_workers":   m.CurrentNumWorkers,
			"did_not_expand_reason": m.DidNotExpandReason,
			"disk_size":             m.DiskSize,
			"driver_state_message":  m.DriverStateMessage,
			"enable_termination_for_node_blocklisted": m.EnableTerminationForNodeBlocklisted,
			"free_space":            m.FreeSpace,
			"init_scripts":          m.InitScripts,
			"instance_id":           m.InstanceId,
			"job_run_name":          m.JobRunName,
			"previous_attributes":   m.PreviousAttributes,
			"previous_cluster_size": m.PreviousClusterSize,
			"previous_disk_size":    m.PreviousDiskSize,
			"reason":                m.Reason,
			"target_num_vcpus":      m.TargetNumVcpus,
			"target_num_workers":    m.TargetNumWorkers,
			"user":                  m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EventDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":            ClusterAttributes{}.Type(ctx),
			"cause":                 types.StringType,
			"cluster_size":          ClusterSize{}.Type(ctx),
			"current_num_vcpus":     types.Int64Type,
			"current_num_workers":   types.Int64Type,
			"did_not_expand_reason": types.StringType,
			"disk_size":             types.Int64Type,
			"driver_state_message":  types.StringType,
			"enable_termination_for_node_blocklisted": types.BoolType,
			"free_space":            types.Int64Type,
			"init_scripts":          InitScriptEventDetails{}.Type(ctx),
			"instance_id":           types.StringType,
			"job_run_name":          types.StringType,
			"previous_attributes":   ClusterAttributes{}.Type(ctx),
			"previous_cluster_size": ClusterSize{}.Type(ctx),
			"previous_disk_size":    types.Int64Type,
			"reason":                TerminationReason{}.Type(ctx),
			"target_num_vcpus":      types.Int64Type,
			"target_num_workers":    types.Int64Type,
			"user":                  types.StringType,
		},
	}
}

// GetAttributes returns the value of the Attributes field in EventDetails as
// a ClusterAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *EventDetails) GetAttributes(ctx context.Context) (ClusterAttributes, bool) {
	var e ClusterAttributes
	if m.Attributes.IsNull() || m.Attributes.IsUnknown() {
		return e, false
	}
	var v ClusterAttributes
	d := m.Attributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAttributes sets the value of the Attributes field in EventDetails.
func (m *EventDetails) SetAttributes(ctx context.Context, v ClusterAttributes) {
	vs := v.ToObjectValue(ctx)
	m.Attributes = vs
}

// GetClusterSize returns the value of the ClusterSize field in EventDetails as
// a ClusterSize value.
// If the field is unknown or null, the boolean return value is false.
func (m *EventDetails) GetClusterSize(ctx context.Context) (ClusterSize, bool) {
	var e ClusterSize
	if m.ClusterSize.IsNull() || m.ClusterSize.IsUnknown() {
		return e, false
	}
	var v ClusterSize
	d := m.ClusterSize.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterSize sets the value of the ClusterSize field in EventDetails.
func (m *EventDetails) SetClusterSize(ctx context.Context, v ClusterSize) {
	vs := v.ToObjectValue(ctx)
	m.ClusterSize = vs
}

// GetInitScripts returns the value of the InitScripts field in EventDetails as
// a InitScriptEventDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *EventDetails) GetInitScripts(ctx context.Context) (InitScriptEventDetails, bool) {
	var e InitScriptEventDetails
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return e, false
	}
	var v InitScriptEventDetails
	d := m.InitScripts.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in EventDetails.
func (m *EventDetails) SetInitScripts(ctx context.Context, v InitScriptEventDetails) {
	vs := v.ToObjectValue(ctx)
	m.InitScripts = vs
}

// GetPreviousAttributes returns the value of the PreviousAttributes field in EventDetails as
// a ClusterAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *EventDetails) GetPreviousAttributes(ctx context.Context) (ClusterAttributes, bool) {
	var e ClusterAttributes
	if m.PreviousAttributes.IsNull() || m.PreviousAttributes.IsUnknown() {
		return e, false
	}
	var v ClusterAttributes
	d := m.PreviousAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreviousAttributes sets the value of the PreviousAttributes field in EventDetails.
func (m *EventDetails) SetPreviousAttributes(ctx context.Context, v ClusterAttributes) {
	vs := v.ToObjectValue(ctx)
	m.PreviousAttributes = vs
}

// GetPreviousClusterSize returns the value of the PreviousClusterSize field in EventDetails as
// a ClusterSize value.
// If the field is unknown or null, the boolean return value is false.
func (m *EventDetails) GetPreviousClusterSize(ctx context.Context) (ClusterSize, bool) {
	var e ClusterSize
	if m.PreviousClusterSize.IsNull() || m.PreviousClusterSize.IsUnknown() {
		return e, false
	}
	var v ClusterSize
	d := m.PreviousClusterSize.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreviousClusterSize sets the value of the PreviousClusterSize field in EventDetails.
func (m *EventDetails) SetPreviousClusterSize(ctx context.Context, v ClusterSize) {
	vs := v.ToObjectValue(ctx)
	m.PreviousClusterSize = vs
}

// GetReason returns the value of the Reason field in EventDetails as
// a TerminationReason value.
// If the field is unknown or null, the boolean return value is false.
func (m *EventDetails) GetReason(ctx context.Context) (TerminationReason, bool) {
	var e TerminationReason
	if m.Reason.IsNull() || m.Reason.IsUnknown() {
		return e, false
	}
	var v TerminationReason
	d := m.Reason.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReason sets the value of the Reason field in EventDetails.
func (m *EventDetails) SetReason(ctx context.Context, v TerminationReason) {
	vs := v.ToObjectValue(ctx)
	m.Reason = vs
}

// Attributes set during cluster creation which are related to GCP.
type GcpAttributes struct {
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs, on-demand VMs, or preemptible VMs with a fallback
	// to on-demand VMs if the former is unavailable.
	Availability types.String `tfsdk:"availability"`
	// Boot disk size in GB
	BootDiskSize types.Int64 `tfsdk:"boot_disk_size"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. This value should be greater than 0, to make sure
	// the cluster driver node is placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand types.Int64 `tfsdk:"first_on_demand"`
	// If provided, the cluster will impersonate the google service account when
	// accessing gcloud services (like GCS). The google service account must
	// have previously been added to the Databricks environment by an account
	// administrator.
	GoogleServiceAccount types.String `tfsdk:"google_service_account"`
	// If provided, each node (workers and driver) in the cluster will have this
	// number of local SSDs attached. Each local SSD is 375GB in size. Refer to
	// [GCP documentation] for the supported number of local SSDs for each
	// instance type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount types.Int64 `tfsdk:"local_ssd_count"`
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs (when set to true) versus standard compute engine
	// VMs (when set to false; default). Note: Soon to be deprecated, use the
	// 'availability' field instead.
	UsePreemptibleExecutors types.Bool `tfsdk:"use_preemptible_executors"`
	// Identifier for the availability zone in which the cluster resides. This
	// can be one of the following: - "HA" => High availability, spread nodes
	// across availability zones for a Databricks deployment region [default]. -
	// "AUTO" => Databricks picks an availability zone to schedule the cluster
	// on. - A GCP availability zone => Pick One of the available zones for
	// (machine type + region) from
	// https://cloud.google.com/compute/docs/regions-zones.
	ZoneId types.String `tfsdk:"zone_id"`
}

func (to *GcpAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpAttributes) {
}

func (to *GcpAttributes) SyncFieldsDuringRead(ctx context.Context, from GcpAttributes) {
}

func (m GcpAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability"] = attrs["availability"].SetOptional()
	attrs["boot_disk_size"] = attrs["boot_disk_size"].SetOptional()
	attrs["first_on_demand"] = attrs["first_on_demand"].SetOptional()
	attrs["google_service_account"] = attrs["google_service_account"].SetOptional()
	attrs["local_ssd_count"] = attrs["local_ssd_count"].SetOptional()
	attrs["use_preemptible_executors"] = attrs["use_preemptible_executors"].SetOptional()
	attrs["zone_id"] = attrs["zone_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GcpAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpAttributes
// only implements ToObjectValue() and Type().
func (m GcpAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":              m.Availability,
			"boot_disk_size":            m.BootDiskSize,
			"first_on_demand":           m.FirstOnDemand,
			"google_service_account":    m.GoogleServiceAccount,
			"local_ssd_count":           m.LocalSsdCount,
			"use_preemptible_executors": m.UsePreemptibleExecutors,
			"zone_id":                   m.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GcpAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":              types.StringType,
			"boot_disk_size":            types.Int64Type,
			"first_on_demand":           types.Int64Type,
			"google_service_account":    types.StringType,
			"local_ssd_count":           types.Int64Type,
			"use_preemptible_executors": types.BoolType,
			"zone_id":                   types.StringType,
		},
	}
}

// A storage location in Google Cloud Platform's GCS
type GcsStorageInfo struct {
	// GCS destination/URI, e.g. `gs://my-bucket/some-prefix`
	Destination types.String `tfsdk:"destination"`
}

func (to *GcsStorageInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcsStorageInfo) {
}

func (to *GcsStorageInfo) SyncFieldsDuringRead(ctx context.Context, from GcsStorageInfo) {
}

func (m GcsStorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination"] = attrs["destination"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcsStorageInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GcsStorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcsStorageInfo
// only implements ToObjectValue() and Type().
func (m GcsStorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": m.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GcsStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type GetClusterComplianceRequest struct {
	// The ID of the cluster to get the compliance status
	ClusterId types.String `tfsdk:"-"`
}

func (to *GetClusterComplianceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterComplianceRequest) {
}

func (to *GetClusterComplianceRequest) SyncFieldsDuringRead(ctx context.Context, from GetClusterComplianceRequest) {
}

func (m GetClusterComplianceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterComplianceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterComplianceRequest
// only implements ToObjectValue() and Type().
func (m GetClusterComplianceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterComplianceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type GetClusterComplianceResponse struct {
	// Whether the cluster is compliant with its policy or not. Clusters could
	// be out of compliance if the policy was updated after the cluster was last
	// edited.
	IsCompliant types.Bool `tfsdk:"is_compliant"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	Violations types.Map `tfsdk:"violations"`
}

func (to *GetClusterComplianceResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterComplianceResponse) {
}

func (to *GetClusterComplianceResponse) SyncFieldsDuringRead(ctx context.Context, from GetClusterComplianceResponse) {
}

func (m GetClusterComplianceResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_compliant"] = attrs["is_compliant"].SetOptional()
	attrs["violations"] = attrs["violations"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterComplianceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterComplianceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterComplianceResponse
// only implements ToObjectValue() and Type().
func (m GetClusterComplianceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_compliant": m.IsCompliant,
			"violations":   m.Violations,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterComplianceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_compliant": types.BoolType,
			"violations": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetViolations returns the value of the Violations field in GetClusterComplianceResponse as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetClusterComplianceResponse) GetViolations(ctx context.Context) (map[string]types.String, bool) {
	if m.Violations.IsNull() || m.Violations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Violations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViolations sets the value of the Violations field in GetClusterComplianceResponse.
func (m *GetClusterComplianceResponse) SetViolations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["violations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Violations = types.MapValueMust(t, vs)
}

type GetClusterPermissionLevelsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-"`
}

func (to *GetClusterPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterPermissionLevelsRequest) {
}

func (to *GetClusterPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetClusterPermissionLevelsRequest) {
}

func (m GetClusterPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetClusterPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type GetClusterPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetClusterPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetClusterPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetClusterPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetClusterPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ClusterPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetClusterPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ClusterPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetClusterPermissionLevelsResponse as
// a slice of ClusterPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetClusterPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]ClusterPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ClusterPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetClusterPermissionLevelsResponse.
func (m *GetClusterPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []ClusterPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetClusterPermissionsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-"`
}

func (to *GetClusterPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterPermissionsRequest) {
}

func (to *GetClusterPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetClusterPermissionsRequest) {
}

func (m GetClusterPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetClusterPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type GetClusterPolicyPermissionLevelsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-"`
}

func (to *GetClusterPolicyPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterPolicyPermissionLevelsRequest) {
}

func (to *GetClusterPolicyPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetClusterPolicyPermissionLevelsRequest) {
}

func (m GetClusterPolicyPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_policy_id"] = attrs["cluster_policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPolicyPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterPolicyPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetClusterPolicyPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_policy_id": m.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterPolicyPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_policy_id": types.StringType,
		},
	}
}

type GetClusterPolicyPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetClusterPolicyPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterPolicyPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetClusterPolicyPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetClusterPolicyPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetClusterPolicyPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPolicyPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterPolicyPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ClusterPolicyPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetClusterPolicyPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterPolicyPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ClusterPolicyPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetClusterPolicyPermissionLevelsResponse as
// a slice of ClusterPolicyPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetClusterPolicyPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]ClusterPolicyPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetClusterPolicyPermissionLevelsResponse.
func (m *GetClusterPolicyPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []ClusterPolicyPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetClusterPolicyPermissionsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-"`
}

func (to *GetClusterPolicyPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterPolicyPermissionsRequest) {
}

func (to *GetClusterPolicyPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetClusterPolicyPermissionsRequest) {
}

func (m GetClusterPolicyPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_policy_id"] = attrs["cluster_policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPolicyPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterPolicyPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetClusterPolicyPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_policy_id": m.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterPolicyPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_policy_id": types.StringType,
		},
	}
}

type GetClusterPolicyRequest struct {
	// Canonical unique identifier for the Cluster Policy.
	PolicyId types.String `tfsdk:"-"`
}

func (to *GetClusterPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterPolicyRequest) {
}

func (to *GetClusterPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from GetClusterPolicyRequest) {
}

func (m GetClusterPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyRequest
// only implements ToObjectValue() and Type().
func (m GetClusterPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type GetClusterRequest struct {
	// The cluster about which to retrieve information.
	ClusterId types.String `tfsdk:"-"`
}

func (to *GetClusterRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetClusterRequest) {
}

func (to *GetClusterRequest) SyncFieldsDuringRead(ctx context.Context, from GetClusterRequest) {
}

func (m GetClusterRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetClusterRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterRequest
// only implements ToObjectValue() and Type().
func (m GetClusterRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetClusterRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type GetDefaultBaseEnvironmentRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetDefaultBaseEnvironmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDefaultBaseEnvironmentRequest) {
}

func (to *GetDefaultBaseEnvironmentRequest) SyncFieldsDuringRead(ctx context.Context, from GetDefaultBaseEnvironmentRequest) {
}

func (m GetDefaultBaseEnvironmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDefaultBaseEnvironmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDefaultBaseEnvironmentRequest
// only implements ToObjectValue() and Type().
func (m GetDefaultBaseEnvironmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDefaultBaseEnvironmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetEvents struct {
	// The ID of the cluster to retrieve events about.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The end time in epoch milliseconds. If empty, returns events up to the
	// current time.
	EndTime types.Int64 `tfsdk:"end_time"`
	// An optional set of event types to filter on. If empty, all event types
	// are returned.
	EventTypes types.List `tfsdk:"event_types"`
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The maximum number of events to include in a page of events. Defaults to
	// 50, and maximum allowed value is 500.
	Limit types.Int64 `tfsdk:"limit"`
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The offset in the result set. Defaults to 0 (no offset). When an offset
	// is specified and the results are requested in descending order, the
	// end_time field is required.
	Offset types.Int64 `tfsdk:"offset"`
	// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".
	Order types.String `tfsdk:"order"`
	// The maximum number of events to include in a page of events. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is empty or 0, the server will decide the number
	// of results to be returned. The field has to be in the range [0,500]. If
	// the value is outside the range, the server enforces 0 or 500.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of events respectively. If page_token
	// is empty, the first page is returned.
	PageToken types.String `tfsdk:"page_token"`
	// The start time in epoch milliseconds. If empty, returns events starting
	// from the beginning of time.
	StartTime types.Int64 `tfsdk:"start_time"`
}

func (to *GetEvents) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEvents) {
	if !from.EventTypes.IsNull() && !from.EventTypes.IsUnknown() && to.EventTypes.IsNull() && len(from.EventTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EventTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EventTypes = from.EventTypes
	}
}

func (to *GetEvents) SyncFieldsDuringRead(ctx context.Context, from GetEvents) {
	if !from.EventTypes.IsNull() && !from.EventTypes.IsUnknown() && to.EventTypes.IsNull() && len(from.EventTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EventTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EventTypes = from.EventTypes
	}
}

func (m GetEvents) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["end_time"] = attrs["end_time"].SetOptional()
	attrs["event_types"] = attrs["event_types"].SetOptional()
	attrs["limit"] = attrs["limit"].SetOptional()
	attrs["offset"] = attrs["offset"].SetOptional()
	attrs["order"] = attrs["order"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEvents.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEvents) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"event_types": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEvents
// only implements ToObjectValue() and Type().
func (m GetEvents) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":  m.ClusterId,
			"end_time":    m.EndTime,
			"event_types": m.EventTypes,
			"limit":       m.Limit,
			"offset":      m.Offset,
			"order":       m.Order,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
			"start_time":  m.StartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEvents) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"end_time":   types.Int64Type,
			"event_types": basetypes.ListType{
				ElemType: types.StringType,
			},
			"limit":      types.Int64Type,
			"offset":     types.Int64Type,
			"order":      types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"start_time": types.Int64Type,
		},
	}
}

// GetEventTypes returns the value of the EventTypes field in GetEvents as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetEvents) GetEventTypes(ctx context.Context) ([]types.String, bool) {
	if m.EventTypes.IsNull() || m.EventTypes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EventTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEventTypes sets the value of the EventTypes field in GetEvents.
func (m *GetEvents) SetEventTypes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["event_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EventTypes = types.ListValueMust(t, vs)
}

type GetEventsResponse struct {
	Events types.List `tfsdk:"events"`
	// Deprecated: use next_page_token or prev_page_token instead.
	//
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	NextPage types.Object `tfsdk:"next_page"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken types.String `tfsdk:"prev_page_token"`
	// Deprecated: Returns 0 when request uses page_token. Will start returning
	// zero when request uses offset/limit soon.
	//
	// The total number of events filtered by the start_time, end_time, and
	// event_types.
	TotalCount types.Int64 `tfsdk:"total_count"`
}

func (to *GetEventsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEventsResponse) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
	if !from.NextPage.IsNull() && !from.NextPage.IsUnknown() {
		if toNextPage, ok := to.GetNextPage(ctx); ok {
			if fromNextPage, ok := from.GetNextPage(ctx); ok {
				// Recursively sync the fields of NextPage
				toNextPage.SyncFieldsDuringCreateOrUpdate(ctx, fromNextPage)
				to.SetNextPage(ctx, toNextPage)
			}
		}
	}
}

func (to *GetEventsResponse) SyncFieldsDuringRead(ctx context.Context, from GetEventsResponse) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
	if !from.NextPage.IsNull() && !from.NextPage.IsUnknown() {
		if toNextPage, ok := to.GetNextPage(ctx); ok {
			if fromNextPage, ok := from.GetNextPage(ctx); ok {
				toNextPage.SyncFieldsDuringRead(ctx, fromNextPage)
				to.SetNextPage(ctx, toNextPage)
			}
		}
	}
}

func (m GetEventsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["events"] = attrs["events"].SetOptional()
	attrs["next_page"] = attrs["next_page"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["prev_page_token"] = attrs["prev_page_token"].SetOptional()
	attrs["total_count"] = attrs["total_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEventsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEventsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":    reflect.TypeOf(ClusterEvent{}),
		"next_page": reflect.TypeOf(GetEvents{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEventsResponse
// only implements ToObjectValue() and Type().
func (m GetEventsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"events":          m.Events,
			"next_page":       m.NextPage,
			"next_page_token": m.NextPageToken,
			"prev_page_token": m.PrevPageToken,
			"total_count":     m.TotalCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEventsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"events": basetypes.ListType{
				ElemType: ClusterEvent{}.Type(ctx),
			},
			"next_page":       GetEvents{}.Type(ctx),
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
			"total_count":     types.Int64Type,
		},
	}
}

// GetEvents returns the value of the Events field in GetEventsResponse as
// a slice of ClusterEvent values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetEventsResponse) GetEvents(ctx context.Context) ([]ClusterEvent, bool) {
	if m.Events.IsNull() || m.Events.IsUnknown() {
		return nil, false
	}
	var v []ClusterEvent
	d := m.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in GetEventsResponse.
func (m *GetEventsResponse) SetEvents(ctx context.Context, v []ClusterEvent) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

// GetNextPage returns the value of the NextPage field in GetEventsResponse as
// a GetEvents value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetEventsResponse) GetNextPage(ctx context.Context) (GetEvents, bool) {
	var e GetEvents
	if m.NextPage.IsNull() || m.NextPage.IsUnknown() {
		return e, false
	}
	var v GetEvents
	d := m.NextPage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNextPage sets the value of the NextPage field in GetEventsResponse.
func (m *GetEventsResponse) SetNextPage(ctx context.Context, v GetEvents) {
	vs := v.ToObjectValue(ctx)
	m.NextPage = vs
}

type GetGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-"`
}

func (to *GetGlobalInitScriptRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGlobalInitScriptRequest) {
}

func (to *GetGlobalInitScriptRequest) SyncFieldsDuringRead(ctx context.Context, from GetGlobalInitScriptRequest) {
}

func (m GetGlobalInitScriptRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["script_id"] = attrs["script_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGlobalInitScriptRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetGlobalInitScriptRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGlobalInitScriptRequest
// only implements ToObjectValue() and Type().
func (m GetGlobalInitScriptRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": m.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGlobalInitScriptRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"script_id": types.StringType,
		},
	}
}

type GetInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags types.Map `tfsdk:"custom_tags"`
	// Tags that are added by Databricks regardless of any ``custom_tags``,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
	DefaultTags types.Map `tfsdk:"default_tags"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec types.Object `tfsdk:"disk_spec"`
	// For pools with node type flexibility (Fleet-V2), whether auto generated
	// alternate node type ids are enabled. This field should not be true if
	// node_type_flexibility is set.
	EnableAutoAlternateNodeTypes types.Bool `tfsdk:"enable_auto_alternate_node_types"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes"`
	// Canonical unique identifier for the pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances"`
	// For pools with node type flexibility (Fleet-V2), this object contains the
	// information about the alternate node type ids to use when attempting to
	// launch a cluster if the node type id is not available. This field should
	// not be set if enable_auto_alternate_node_types is true.
	NodeTypeFlexibility types.Object `tfsdk:"node_type_flexibility"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Custom Docker Image BYOC
	PreloadedDockerImages types.List `tfsdk:"preloaded_docker_images"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions types.List `tfsdk:"preloaded_spark_versions"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// Current state of the instance pool.
	State types.String `tfsdk:"state"`
	// Usage statistics about the instance pool.
	Stats types.Object `tfsdk:"stats"`
	// Status of failed pending instances in the pool.
	Status types.Object `tfsdk:"status"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
}

func (to *GetInstancePool) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetInstancePool) {
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.DiskSpec.IsNull() && !from.DiskSpec.IsUnknown() {
		if toDiskSpec, ok := to.GetDiskSpec(ctx); ok {
			if fromDiskSpec, ok := from.GetDiskSpec(ctx); ok {
				// Recursively sync the fields of DiskSpec
				toDiskSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDiskSpec)
				to.SetDiskSpec(ctx, toDiskSpec)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.NodeTypeFlexibility.IsNull() && !from.NodeTypeFlexibility.IsUnknown() {
		if toNodeTypeFlexibility, ok := to.GetNodeTypeFlexibility(ctx); ok {
			if fromNodeTypeFlexibility, ok := from.GetNodeTypeFlexibility(ctx); ok {
				// Recursively sync the fields of NodeTypeFlexibility
				toNodeTypeFlexibility.SyncFieldsDuringCreateOrUpdate(ctx, fromNodeTypeFlexibility)
				to.SetNodeTypeFlexibility(ctx, toNodeTypeFlexibility)
			}
		}
	}
	if !from.PreloadedDockerImages.IsNull() && !from.PreloadedDockerImages.IsUnknown() && to.PreloadedDockerImages.IsNull() && len(from.PreloadedDockerImages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedDockerImages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedDockerImages = from.PreloadedDockerImages
	}
	if !from.PreloadedSparkVersions.IsNull() && !from.PreloadedSparkVersions.IsUnknown() && to.PreloadedSparkVersions.IsNull() && len(from.PreloadedSparkVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedSparkVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedSparkVersions = from.PreloadedSparkVersions
	}
	if !from.Stats.IsNull() && !from.Stats.IsUnknown() {
		if toStats, ok := to.GetStats(ctx); ok {
			if fromStats, ok := from.GetStats(ctx); ok {
				// Recursively sync the fields of Stats
				toStats.SyncFieldsDuringCreateOrUpdate(ctx, fromStats)
				to.SetStats(ctx, toStats)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *GetInstancePool) SyncFieldsDuringRead(ctx context.Context, from GetInstancePool) {
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.DiskSpec.IsNull() && !from.DiskSpec.IsUnknown() {
		if toDiskSpec, ok := to.GetDiskSpec(ctx); ok {
			if fromDiskSpec, ok := from.GetDiskSpec(ctx); ok {
				toDiskSpec.SyncFieldsDuringRead(ctx, fromDiskSpec)
				to.SetDiskSpec(ctx, toDiskSpec)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.NodeTypeFlexibility.IsNull() && !from.NodeTypeFlexibility.IsUnknown() {
		if toNodeTypeFlexibility, ok := to.GetNodeTypeFlexibility(ctx); ok {
			if fromNodeTypeFlexibility, ok := from.GetNodeTypeFlexibility(ctx); ok {
				toNodeTypeFlexibility.SyncFieldsDuringRead(ctx, fromNodeTypeFlexibility)
				to.SetNodeTypeFlexibility(ctx, toNodeTypeFlexibility)
			}
		}
	}
	if !from.PreloadedDockerImages.IsNull() && !from.PreloadedDockerImages.IsUnknown() && to.PreloadedDockerImages.IsNull() && len(from.PreloadedDockerImages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedDockerImages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedDockerImages = from.PreloadedDockerImages
	}
	if !from.PreloadedSparkVersions.IsNull() && !from.PreloadedSparkVersions.IsUnknown() && to.PreloadedSparkVersions.IsNull() && len(from.PreloadedSparkVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedSparkVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedSparkVersions = from.PreloadedSparkVersions
	}
	if !from.Stats.IsNull() && !from.Stats.IsUnknown() {
		if toStats, ok := to.GetStats(ctx); ok {
			if fromStats, ok := from.GetStats(ctx); ok {
				toStats.SyncFieldsDuringRead(ctx, fromStats)
				to.SetStats(ctx, toStats)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m GetInstancePool) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["default_tags"] = attrs["default_tags"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].SetOptional()
	attrs["enable_auto_alternate_node_types"] = attrs["enable_auto_alternate_node_types"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["idle_instance_autotermination_minutes"] = attrs["idle_instance_autotermination_minutes"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()
	attrs["instance_pool_name"] = attrs["instance_pool_name"].SetOptional()
	attrs["max_capacity"] = attrs["max_capacity"].SetOptional()
	attrs["min_idle_instances"] = attrs["min_idle_instances"].SetOptional()
	attrs["node_type_flexibility"] = attrs["node_type_flexibility"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["preloaded_docker_images"] = attrs["preloaded_docker_images"].SetOptional()
	attrs["preloaded_spark_versions"] = attrs["preloaded_spark_versions"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["stats"] = attrs["stats"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetInstancePool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetInstancePool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"default_tags":             reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes{}),
		"node_type_flexibility":    reflect.TypeOf(NodeTypeFlexibility{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
		"stats":                    reflect.TypeOf(InstancePoolStats{}),
		"status":                   reflect.TypeOf(InstancePoolStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePool
// only implements ToObjectValue() and Type().
func (m GetInstancePool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        m.AwsAttributes,
			"azure_attributes":                      m.AzureAttributes,
			"custom_tags":                           m.CustomTags,
			"default_tags":                          m.DefaultTags,
			"disk_spec":                             m.DiskSpec,
			"enable_auto_alternate_node_types":      m.EnableAutoAlternateNodeTypes,
			"enable_elastic_disk":                   m.EnableElasticDisk,
			"gcp_attributes":                        m.GcpAttributes,
			"idle_instance_autotermination_minutes": m.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      m.InstancePoolId,
			"instance_pool_name":                    m.InstancePoolName,
			"max_capacity":                          m.MaxCapacity,
			"min_idle_instances":                    m.MinIdleInstances,
			"node_type_flexibility":                 m.NodeTypeFlexibility,
			"node_type_id":                          m.NodeTypeId,
			"preloaded_docker_images":               m.PreloadedDockerImages,
			"preloaded_spark_versions":              m.PreloadedSparkVersions,
			"remote_disk_throughput":                m.RemoteDiskThroughput,
			"state":                                 m.State,
			"stats":                                 m.Stats,
			"status":                                m.Status,
			"total_initial_remote_disk_size":        m.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes":   InstancePoolAwsAttributes{}.Type(ctx),
			"azure_attributes": InstancePoolAzureAttributes{}.Type(ctx),
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"default_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec":                             DiskSpec{}.Type(ctx),
			"enable_auto_alternate_node_types":      types.BoolType,
			"enable_elastic_disk":                   types.BoolType,
			"gcp_attributes":                        InstancePoolGcpAttributes{}.Type(ctx),
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_id":                      types.StringType,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
			"node_type_flexibility":                 NodeTypeFlexibility{}.Type(ctx),
			"node_type_id":                          types.StringType,
			"preloaded_docker_images": basetypes.ListType{
				ElemType: DockerImage{}.Type(ctx),
			},
			"preloaded_spark_versions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"remote_disk_throughput":         types.Int64Type,
			"state":                          types.StringType,
			"stats":                          InstancePoolStats{}.Type(ctx),
			"status":                         InstancePoolStatus{}.Type(ctx),
			"total_initial_remote_disk_size": types.Int64Type,
		},
	}
}

// GetAwsAttributes returns the value of the AwsAttributes field in GetInstancePool as
// a InstancePoolAwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes, bool) {
	var e InstancePoolAwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolAwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in GetInstancePool.
func (m *GetInstancePool) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in GetInstancePool as
// a InstancePoolAzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes, bool) {
	var e InstancePoolAzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolAzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in GetInstancePool.
func (m *GetInstancePool) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetCustomTags returns the value of the CustomTags field in GetInstancePool as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in GetInstancePool.
func (m *GetInstancePool) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in GetInstancePool as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
	if m.DefaultTags.IsNull() || m.DefaultTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.DefaultTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultTags sets the value of the DefaultTags field in GetInstancePool.
func (m *GetInstancePool) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DefaultTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in GetInstancePool as
// a DiskSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetDiskSpec(ctx context.Context) (DiskSpec, bool) {
	var e DiskSpec
	if m.DiskSpec.IsNull() || m.DiskSpec.IsUnknown() {
		return e, false
	}
	var v DiskSpec
	d := m.DiskSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDiskSpec sets the value of the DiskSpec field in GetInstancePool.
func (m *GetInstancePool) SetDiskSpec(ctx context.Context, v DiskSpec) {
	vs := v.ToObjectValue(ctx)
	m.DiskSpec = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in GetInstancePool as
// a InstancePoolGcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes, bool) {
	var e InstancePoolGcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolGcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in GetInstancePool.
func (m *GetInstancePool) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetNodeTypeFlexibility returns the value of the NodeTypeFlexibility field in GetInstancePool as
// a NodeTypeFlexibility value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetNodeTypeFlexibility(ctx context.Context) (NodeTypeFlexibility, bool) {
	var e NodeTypeFlexibility
	if m.NodeTypeFlexibility.IsNull() || m.NodeTypeFlexibility.IsUnknown() {
		return e, false
	}
	var v NodeTypeFlexibility
	d := m.NodeTypeFlexibility.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeTypeFlexibility sets the value of the NodeTypeFlexibility field in GetInstancePool.
func (m *GetInstancePool) SetNodeTypeFlexibility(ctx context.Context, v NodeTypeFlexibility) {
	vs := v.ToObjectValue(ctx)
	m.NodeTypeFlexibility = vs
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in GetInstancePool as
// a slice of DockerImage values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage, bool) {
	if m.PreloadedDockerImages.IsNull() || m.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage
	d := m.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in GetInstancePool.
func (m *GetInstancePool) SetPreloadedDockerImages(ctx context.Context, v []DockerImage) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in GetInstancePool as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
	if m.PreloadedSparkVersions.IsNull() || m.PreloadedSparkVersions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.PreloadedSparkVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in GetInstancePool.
func (m *GetInstancePool) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

// GetStats returns the value of the Stats field in GetInstancePool as
// a InstancePoolStats value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetStats(ctx context.Context) (InstancePoolStats, bool) {
	var e InstancePoolStats
	if m.Stats.IsNull() || m.Stats.IsUnknown() {
		return e, false
	}
	var v InstancePoolStats
	d := m.Stats.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStats sets the value of the Stats field in GetInstancePool.
func (m *GetInstancePool) SetStats(ctx context.Context, v InstancePoolStats) {
	vs := v.ToObjectValue(ctx)
	m.Stats = vs
}

// GetStatus returns the value of the Status field in GetInstancePool as
// a InstancePoolStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePool) GetStatus(ctx context.Context) (InstancePoolStatus, bool) {
	var e InstancePoolStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v InstancePoolStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in GetInstancePool.
func (m *GetInstancePool) SetStatus(ctx context.Context, v InstancePoolStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

type GetInstancePoolPermissionLevelsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-"`
}

func (to *GetInstancePoolPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetInstancePoolPermissionLevelsRequest) {
}

func (to *GetInstancePoolPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetInstancePoolPermissionLevelsRequest) {
}

func (m GetInstancePoolPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetInstancePoolPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetInstancePoolPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetInstancePoolPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": m.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetInstancePoolPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type GetInstancePoolPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetInstancePoolPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetInstancePoolPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetInstancePoolPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetInstancePoolPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetInstancePoolPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetInstancePoolPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetInstancePoolPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(InstancePoolPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetInstancePoolPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetInstancePoolPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: InstancePoolPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetInstancePoolPermissionLevelsResponse as
// a slice of InstancePoolPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetInstancePoolPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]InstancePoolPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetInstancePoolPermissionLevelsResponse.
func (m *GetInstancePoolPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []InstancePoolPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetInstancePoolPermissionsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-"`
}

func (to *GetInstancePoolPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetInstancePoolPermissionsRequest) {
}

func (to *GetInstancePoolPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetInstancePoolPermissionsRequest) {
}

func (m GetInstancePoolPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetInstancePoolPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetInstancePoolPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetInstancePoolPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": m.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetInstancePoolPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type GetInstancePoolRequest struct {
	// The canonical unique identifier for the instance pool.
	InstancePoolId types.String `tfsdk:"-"`
}

func (to *GetInstancePoolRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetInstancePoolRequest) {
}

func (to *GetInstancePoolRequest) SyncFieldsDuringRead(ctx context.Context, from GetInstancePoolRequest) {
}

func (m GetInstancePoolRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetInstancePoolRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetInstancePoolRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolRequest
// only implements ToObjectValue() and Type().
func (m GetInstancePoolRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": m.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetInstancePoolRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type GetPolicyFamilyRequest struct {
	// The family ID about which to retrieve information.
	PolicyFamilyId types.String `tfsdk:"-"`
	// The version number for the family to fetch. Defaults to the latest
	// version.
	Version types.Int64 `tfsdk:"-"`
}

func (to *GetPolicyFamilyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPolicyFamilyRequest) {
}

func (to *GetPolicyFamilyRequest) SyncFieldsDuringRead(ctx context.Context, from GetPolicyFamilyRequest) {
}

func (m GetPolicyFamilyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy_family_id"] = attrs["policy_family_id"].SetRequired()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPolicyFamilyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetPolicyFamilyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPolicyFamilyRequest
// only implements ToObjectValue() and Type().
func (m GetPolicyFamilyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_family_id": m.PolicyFamilyId,
			"version":          m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPolicyFamilyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_family_id": types.StringType,
			"version":          types.Int64Type,
		},
	}
}

type GetSparkVersionsResponse struct {
	// All the available Spark versions.
	Versions types.List `tfsdk:"versions"`
}

func (to *GetSparkVersionsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSparkVersionsResponse) {
	if !from.Versions.IsNull() && !from.Versions.IsUnknown() && to.Versions.IsNull() && len(from.Versions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Versions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Versions = from.Versions
	}
}

func (to *GetSparkVersionsResponse) SyncFieldsDuringRead(ctx context.Context, from GetSparkVersionsResponse) {
	if !from.Versions.IsNull() && !from.Versions.IsUnknown() && to.Versions.IsNull() && len(from.Versions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Versions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Versions = from.Versions
	}
}

func (m GetSparkVersionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["versions"] = attrs["versions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSparkVersionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSparkVersionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"versions": reflect.TypeOf(SparkVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSparkVersionsResponse
// only implements ToObjectValue() and Type().
func (m GetSparkVersionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"versions": m.Versions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSparkVersionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"versions": basetypes.ListType{
				ElemType: SparkVersion{}.Type(ctx),
			},
		},
	}
}

// GetVersions returns the value of the Versions field in GetSparkVersionsResponse as
// a slice of SparkVersion values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetSparkVersionsResponse) GetVersions(ctx context.Context) ([]SparkVersion, bool) {
	if m.Versions.IsNull() || m.Versions.IsUnknown() {
		return nil, false
	}
	var v []SparkVersion
	d := m.Versions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVersions sets the value of the Versions field in GetSparkVersionsResponse.
func (m *GetSparkVersionsResponse) SetVersions(ctx context.Context, v []SparkVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Versions = types.ListValueMust(t, vs)
}

type GlobalInitScriptCreateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the script
	Name types.String `tfsdk:"name"`
	// The position of a global init script, where 0 represents the first script
	// to run, 1 is the second script to run, in ascending order.
	//
	// If you omit the numeric position for a new global init script, it
	// defaults to last position. It will run after all current scripts. Setting
	// any value greater than the position of the last script is equivalent to
	// the last position. Example: Take three existing scripts with positions 0,
	// 1, and 2. Any position of (3) or greater puts the script in the last
	// position. If an explicit position value conflicts with an existing script
	// value, your request succeeds, but the original script at that position
	// and all later scripts have their positions incremented by 1.
	Position types.Int64 `tfsdk:"position"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script"`
}

func (to *GlobalInitScriptCreateRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GlobalInitScriptCreateRequest) {
}

func (to *GlobalInitScriptCreateRequest) SyncFieldsDuringRead(ctx context.Context, from GlobalInitScriptCreateRequest) {
}

func (m GlobalInitScriptCreateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["script"] = attrs["script"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GlobalInitScriptCreateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GlobalInitScriptCreateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptCreateRequest
// only implements ToObjectValue() and Type().
func (m GlobalInitScriptCreateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":  m.Enabled,
			"name":     m.Name,
			"position": m.Position,
			"script":   m.Script,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GlobalInitScriptCreateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":  types.BoolType,
			"name":     types.StringType,
			"position": types.Int64Type,
			"script":   types.StringType,
		},
	}
}

type GlobalInitScriptDetails struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// The username of the user who created the script.
	CreatedBy types.String `tfsdk:"created_by"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the script
	Name types.String `tfsdk:"name"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position types.Int64 `tfsdk:"position"`
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// The username of the user who last updated the script
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (to *GlobalInitScriptDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GlobalInitScriptDetails) {
}

func (to *GlobalInitScriptDetails) SyncFieldsDuringRead(ctx context.Context, from GlobalInitScriptDetails) {
}

func (m GlobalInitScriptDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["script_id"] = attrs["script_id"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GlobalInitScriptDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GlobalInitScriptDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptDetails
// only implements ToObjectValue() and Type().
func (m GlobalInitScriptDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at": m.CreatedAt,
			"created_by": m.CreatedBy,
			"enabled":    m.Enabled,
			"name":       m.Name,
			"position":   m.Position,
			"script_id":  m.ScriptId,
			"updated_at": m.UpdatedAt,
			"updated_by": m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GlobalInitScriptDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"enabled":    types.BoolType,
			"name":       types.StringType,
			"position":   types.Int64Type,
			"script_id":  types.StringType,
			"updated_at": types.Int64Type,
			"updated_by": types.StringType,
		},
	}
}

type GlobalInitScriptDetailsWithContent struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// The username of the user who created the script.
	CreatedBy types.String `tfsdk:"created_by"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the script
	Name types.String `tfsdk:"name"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position types.Int64 `tfsdk:"position"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script"`
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// The username of the user who last updated the script
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (to *GlobalInitScriptDetailsWithContent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GlobalInitScriptDetailsWithContent) {
}

func (to *GlobalInitScriptDetailsWithContent) SyncFieldsDuringRead(ctx context.Context, from GlobalInitScriptDetailsWithContent) {
}

func (m GlobalInitScriptDetailsWithContent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["script"] = attrs["script"].SetOptional()
	attrs["script_id"] = attrs["script_id"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GlobalInitScriptDetailsWithContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GlobalInitScriptDetailsWithContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptDetailsWithContent
// only implements ToObjectValue() and Type().
func (m GlobalInitScriptDetailsWithContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at": m.CreatedAt,
			"created_by": m.CreatedBy,
			"enabled":    m.Enabled,
			"name":       m.Name,
			"position":   m.Position,
			"script":     m.Script,
			"script_id":  m.ScriptId,
			"updated_at": m.UpdatedAt,
			"updated_by": m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GlobalInitScriptDetailsWithContent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"enabled":    types.BoolType,
			"name":       types.StringType,
			"position":   types.Int64Type,
			"script":     types.StringType,
			"script_id":  types.StringType,
			"updated_at": types.Int64Type,
			"updated_by": types.StringType,
		},
	}
}

type GlobalInitScriptUpdateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the script
	Name types.String `tfsdk:"name"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order. To move the script to
	// run first, set its position to 0.
	//
	// To move the script to the end, set its position to any value greater or
	// equal to the position of the last script. Example, three existing scripts
	// with positions 0, 1, and 2. Any position value of 2 or greater puts the
	// script in the last position (2).
	//
	// If an explicit position value conflicts with an existing script, your
	// request succeeds, but the original script at that position and all later
	// scripts have their positions incremented by 1.
	Position types.Int64 `tfsdk:"position"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script"`
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-"`
}

func (to *GlobalInitScriptUpdateRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GlobalInitScriptUpdateRequest) {
}

func (to *GlobalInitScriptUpdateRequest) SyncFieldsDuringRead(ctx context.Context, from GlobalInitScriptUpdateRequest) {
}

func (m GlobalInitScriptUpdateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["script"] = attrs["script"].SetRequired()
	attrs["script_id"] = attrs["script_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GlobalInitScriptUpdateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GlobalInitScriptUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptUpdateRequest
// only implements ToObjectValue() and Type().
func (m GlobalInitScriptUpdateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":   m.Enabled,
			"name":      m.Name,
			"position":  m.Position,
			"script":    m.Script,
			"script_id": m.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GlobalInitScriptUpdateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":   types.BoolType,
			"name":      types.StringType,
			"position":  types.Int64Type,
			"script":    types.StringType,
			"script_id": types.StringType,
		},
	}
}

type InitScriptEventDetails struct {
	// The cluster scoped init scripts associated with this cluster event.
	Cluster types.List `tfsdk:"cluster"`
	// The global init scripts associated with this cluster event.
	Global types.List `tfsdk:"global"`
	// The private ip of the node we are reporting init script execution details
	// for (we will select the execution details from only one node rather than
	// reporting the execution details from every node to keep these event
	// details small)
	//
	// This should only be defined for the INIT_SCRIPTS_FINISHED event
	ReportedForNode types.String `tfsdk:"reported_for_node"`
}

func (to *InitScriptEventDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InitScriptEventDetails) {
	if !from.Cluster.IsNull() && !from.Cluster.IsUnknown() && to.Cluster.IsNull() && len(from.Cluster.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Cluster, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Cluster = from.Cluster
	}
	if !from.Global.IsNull() && !from.Global.IsUnknown() && to.Global.IsNull() && len(from.Global.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Global, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Global = from.Global
	}
}

func (to *InitScriptEventDetails) SyncFieldsDuringRead(ctx context.Context, from InitScriptEventDetails) {
	if !from.Cluster.IsNull() && !from.Cluster.IsUnknown() && to.Cluster.IsNull() && len(from.Cluster.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Cluster, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Cluster = from.Cluster
	}
	if !from.Global.IsNull() && !from.Global.IsUnknown() && to.Global.IsNull() && len(from.Global.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Global, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Global = from.Global
	}
}

func (m InitScriptEventDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster"] = attrs["cluster"].SetOptional()
	attrs["global"] = attrs["global"].SetOptional()
	attrs["reported_for_node"] = attrs["reported_for_node"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InitScriptEventDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InitScriptEventDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster": reflect.TypeOf(InitScriptInfoAndExecutionDetails{}),
		"global":  reflect.TypeOf(InitScriptInfoAndExecutionDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InitScriptEventDetails
// only implements ToObjectValue() and Type().
func (m InitScriptEventDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster":           m.Cluster,
			"global":            m.Global,
			"reported_for_node": m.ReportedForNode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InitScriptEventDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster": basetypes.ListType{
				ElemType: InitScriptInfoAndExecutionDetails{}.Type(ctx),
			},
			"global": basetypes.ListType{
				ElemType: InitScriptInfoAndExecutionDetails{}.Type(ctx),
			},
			"reported_for_node": types.StringType,
		},
	}
}

// GetCluster returns the value of the Cluster field in InitScriptEventDetails as
// a slice of InitScriptInfoAndExecutionDetails values.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptEventDetails) GetCluster(ctx context.Context) ([]InitScriptInfoAndExecutionDetails, bool) {
	if m.Cluster.IsNull() || m.Cluster.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfoAndExecutionDetails
	d := m.Cluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCluster sets the value of the Cluster field in InitScriptEventDetails.
func (m *InitScriptEventDetails) SetCluster(ctx context.Context, v []InitScriptInfoAndExecutionDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Cluster = types.ListValueMust(t, vs)
}

// GetGlobal returns the value of the Global field in InitScriptEventDetails as
// a slice of InitScriptInfoAndExecutionDetails values.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptEventDetails) GetGlobal(ctx context.Context) ([]InitScriptInfoAndExecutionDetails, bool) {
	if m.Global.IsNull() || m.Global.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfoAndExecutionDetails
	d := m.Global.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGlobal sets the value of the Global field in InitScriptEventDetails.
func (m *InitScriptEventDetails) SetGlobal(ctx context.Context, v []InitScriptInfoAndExecutionDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["global"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Global = types.ListValueMust(t, vs)
}

// Config for an individual init script Next ID: 11
type InitScriptInfo struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	Abfss types.Object `tfsdk:"abfss"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs types.Object `tfsdk:"dbfs"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File types.Object `tfsdk:"file"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs types.Object `tfsdk:"gcs"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 types.Object `tfsdk:"s3"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes types.Object `tfsdk:"volumes"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace types.Object `tfsdk:"workspace"`
}

func (to *InitScriptInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InitScriptInfo) {
	if !from.Abfss.IsNull() && !from.Abfss.IsUnknown() {
		if toAbfss, ok := to.GetAbfss(ctx); ok {
			if fromAbfss, ok := from.GetAbfss(ctx); ok {
				// Recursively sync the fields of Abfss
				toAbfss.SyncFieldsDuringCreateOrUpdate(ctx, fromAbfss)
				to.SetAbfss(ctx, toAbfss)
			}
		}
	}
	if !from.Dbfs.IsNull() && !from.Dbfs.IsUnknown() {
		if toDbfs, ok := to.GetDbfs(ctx); ok {
			if fromDbfs, ok := from.GetDbfs(ctx); ok {
				// Recursively sync the fields of Dbfs
				toDbfs.SyncFieldsDuringCreateOrUpdate(ctx, fromDbfs)
				to.SetDbfs(ctx, toDbfs)
			}
		}
	}
	if !from.File.IsNull() && !from.File.IsUnknown() {
		if toFile, ok := to.GetFile(ctx); ok {
			if fromFile, ok := from.GetFile(ctx); ok {
				// Recursively sync the fields of File
				toFile.SyncFieldsDuringCreateOrUpdate(ctx, fromFile)
				to.SetFile(ctx, toFile)
			}
		}
	}
	if !from.Gcs.IsNull() && !from.Gcs.IsUnknown() {
		if toGcs, ok := to.GetGcs(ctx); ok {
			if fromGcs, ok := from.GetGcs(ctx); ok {
				// Recursively sync the fields of Gcs
				toGcs.SyncFieldsDuringCreateOrUpdate(ctx, fromGcs)
				to.SetGcs(ctx, toGcs)
			}
		}
	}
	if !from.S3.IsNull() && !from.S3.IsUnknown() {
		if toS3, ok := to.GetS3(ctx); ok {
			if fromS3, ok := from.GetS3(ctx); ok {
				// Recursively sync the fields of S3
				toS3.SyncFieldsDuringCreateOrUpdate(ctx, fromS3)
				to.SetS3(ctx, toS3)
			}
		}
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() {
		if toVolumes, ok := to.GetVolumes(ctx); ok {
			if fromVolumes, ok := from.GetVolumes(ctx); ok {
				// Recursively sync the fields of Volumes
				toVolumes.SyncFieldsDuringCreateOrUpdate(ctx, fromVolumes)
				to.SetVolumes(ctx, toVolumes)
			}
		}
	}
	if !from.Workspace.IsNull() && !from.Workspace.IsUnknown() {
		if toWorkspace, ok := to.GetWorkspace(ctx); ok {
			if fromWorkspace, ok := from.GetWorkspace(ctx); ok {
				// Recursively sync the fields of Workspace
				toWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspace)
				to.SetWorkspace(ctx, toWorkspace)
			}
		}
	}
}

func (to *InitScriptInfo) SyncFieldsDuringRead(ctx context.Context, from InitScriptInfo) {
	if !from.Abfss.IsNull() && !from.Abfss.IsUnknown() {
		if toAbfss, ok := to.GetAbfss(ctx); ok {
			if fromAbfss, ok := from.GetAbfss(ctx); ok {
				toAbfss.SyncFieldsDuringRead(ctx, fromAbfss)
				to.SetAbfss(ctx, toAbfss)
			}
		}
	}
	if !from.Dbfs.IsNull() && !from.Dbfs.IsUnknown() {
		if toDbfs, ok := to.GetDbfs(ctx); ok {
			if fromDbfs, ok := from.GetDbfs(ctx); ok {
				toDbfs.SyncFieldsDuringRead(ctx, fromDbfs)
				to.SetDbfs(ctx, toDbfs)
			}
		}
	}
	if !from.File.IsNull() && !from.File.IsUnknown() {
		if toFile, ok := to.GetFile(ctx); ok {
			if fromFile, ok := from.GetFile(ctx); ok {
				toFile.SyncFieldsDuringRead(ctx, fromFile)
				to.SetFile(ctx, toFile)
			}
		}
	}
	if !from.Gcs.IsNull() && !from.Gcs.IsUnknown() {
		if toGcs, ok := to.GetGcs(ctx); ok {
			if fromGcs, ok := from.GetGcs(ctx); ok {
				toGcs.SyncFieldsDuringRead(ctx, fromGcs)
				to.SetGcs(ctx, toGcs)
			}
		}
	}
	if !from.S3.IsNull() && !from.S3.IsUnknown() {
		if toS3, ok := to.GetS3(ctx); ok {
			if fromS3, ok := from.GetS3(ctx); ok {
				toS3.SyncFieldsDuringRead(ctx, fromS3)
				to.SetS3(ctx, toS3)
			}
		}
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() {
		if toVolumes, ok := to.GetVolumes(ctx); ok {
			if fromVolumes, ok := from.GetVolumes(ctx); ok {
				toVolumes.SyncFieldsDuringRead(ctx, fromVolumes)
				to.SetVolumes(ctx, toVolumes)
			}
		}
	}
	if !from.Workspace.IsNull() && !from.Workspace.IsUnknown() {
		if toWorkspace, ok := to.GetWorkspace(ctx); ok {
			if fromWorkspace, ok := from.GetWorkspace(ctx); ok {
				toWorkspace.SyncFieldsDuringRead(ctx, fromWorkspace)
				to.SetWorkspace(ctx, toWorkspace)
			}
		}
	}
}

func (m InitScriptInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["abfss"] = attrs["abfss"].SetOptional()
	attrs["dbfs"] = attrs["dbfs"].SetOptional()
	attrs["file"] = attrs["file"].SetOptional()
	attrs["gcs"] = attrs["gcs"].SetOptional()
	attrs["s3"] = attrs["s3"].SetOptional()
	attrs["volumes"] = attrs["volumes"].SetOptional()
	attrs["workspace"] = attrs["workspace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InitScriptInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InitScriptInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"abfss":     reflect.TypeOf(Adlsgen2Info{}),
		"dbfs":      reflect.TypeOf(DbfsStorageInfo{}),
		"file":      reflect.TypeOf(LocalFileInfo{}),
		"gcs":       reflect.TypeOf(GcsStorageInfo{}),
		"s3":        reflect.TypeOf(S3StorageInfo{}),
		"volumes":   reflect.TypeOf(VolumesStorageInfo{}),
		"workspace": reflect.TypeOf(WorkspaceStorageInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InitScriptInfo
// only implements ToObjectValue() and Type().
func (m InitScriptInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"abfss":     m.Abfss,
			"dbfs":      m.Dbfs,
			"file":      m.File,
			"gcs":       m.Gcs,
			"s3":        m.S3,
			"volumes":   m.Volumes,
			"workspace": m.Workspace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InitScriptInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"abfss":     Adlsgen2Info{}.Type(ctx),
			"dbfs":      DbfsStorageInfo{}.Type(ctx),
			"file":      LocalFileInfo{}.Type(ctx),
			"gcs":       GcsStorageInfo{}.Type(ctx),
			"s3":        S3StorageInfo{}.Type(ctx),
			"volumes":   VolumesStorageInfo{}.Type(ctx),
			"workspace": WorkspaceStorageInfo{}.Type(ctx),
		},
	}
}

// GetAbfss returns the value of the Abfss field in InitScriptInfo as
// a Adlsgen2Info value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfo) GetAbfss(ctx context.Context) (Adlsgen2Info, bool) {
	var e Adlsgen2Info
	if m.Abfss.IsNull() || m.Abfss.IsUnknown() {
		return e, false
	}
	var v Adlsgen2Info
	d := m.Abfss.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAbfss sets the value of the Abfss field in InitScriptInfo.
func (m *InitScriptInfo) SetAbfss(ctx context.Context, v Adlsgen2Info) {
	vs := v.ToObjectValue(ctx)
	m.Abfss = vs
}

// GetDbfs returns the value of the Dbfs field in InitScriptInfo as
// a DbfsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfo) GetDbfs(ctx context.Context) (DbfsStorageInfo, bool) {
	var e DbfsStorageInfo
	if m.Dbfs.IsNull() || m.Dbfs.IsUnknown() {
		return e, false
	}
	var v DbfsStorageInfo
	d := m.Dbfs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDbfs sets the value of the Dbfs field in InitScriptInfo.
func (m *InitScriptInfo) SetDbfs(ctx context.Context, v DbfsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Dbfs = vs
}

// GetFile returns the value of the File field in InitScriptInfo as
// a LocalFileInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfo) GetFile(ctx context.Context) (LocalFileInfo, bool) {
	var e LocalFileInfo
	if m.File.IsNull() || m.File.IsUnknown() {
		return e, false
	}
	var v LocalFileInfo
	d := m.File.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFile sets the value of the File field in InitScriptInfo.
func (m *InitScriptInfo) SetFile(ctx context.Context, v LocalFileInfo) {
	vs := v.ToObjectValue(ctx)
	m.File = vs
}

// GetGcs returns the value of the Gcs field in InitScriptInfo as
// a GcsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfo) GetGcs(ctx context.Context) (GcsStorageInfo, bool) {
	var e GcsStorageInfo
	if m.Gcs.IsNull() || m.Gcs.IsUnknown() {
		return e, false
	}
	var v GcsStorageInfo
	d := m.Gcs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcs sets the value of the Gcs field in InitScriptInfo.
func (m *InitScriptInfo) SetGcs(ctx context.Context, v GcsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Gcs = vs
}

// GetS3 returns the value of the S3 field in InitScriptInfo as
// a S3StorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfo) GetS3(ctx context.Context) (S3StorageInfo, bool) {
	var e S3StorageInfo
	if m.S3.IsNull() || m.S3.IsUnknown() {
		return e, false
	}
	var v S3StorageInfo
	d := m.S3.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetS3 sets the value of the S3 field in InitScriptInfo.
func (m *InitScriptInfo) SetS3(ctx context.Context, v S3StorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.S3 = vs
}

// GetVolumes returns the value of the Volumes field in InitScriptInfo as
// a VolumesStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfo) GetVolumes(ctx context.Context) (VolumesStorageInfo, bool) {
	var e VolumesStorageInfo
	if m.Volumes.IsNull() || m.Volumes.IsUnknown() {
		return e, false
	}
	var v VolumesStorageInfo
	d := m.Volumes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumes sets the value of the Volumes field in InitScriptInfo.
func (m *InitScriptInfo) SetVolumes(ctx context.Context, v VolumesStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Volumes = vs
}

// GetWorkspace returns the value of the Workspace field in InitScriptInfo as
// a WorkspaceStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfo) GetWorkspace(ctx context.Context) (WorkspaceStorageInfo, bool) {
	var e WorkspaceStorageInfo
	if m.Workspace.IsNull() || m.Workspace.IsUnknown() {
		return e, false
	}
	var v WorkspaceStorageInfo
	d := m.Workspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspace sets the value of the Workspace field in InitScriptInfo.
func (m *InitScriptInfo) SetWorkspace(ctx context.Context, v WorkspaceStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Workspace = vs
}

type InitScriptInfoAndExecutionDetails struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	Abfss types.Object `tfsdk:"abfss"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs types.Object `tfsdk:"dbfs"`
	// Additional details regarding errors (such as a file not found message if
	// the status is FAILED_FETCH). This field should only be used to provide
	// *additional* information to the status field, not duplicate it.
	ErrorMessage types.String `tfsdk:"error_message"`
	// The number duration of the script execution in seconds
	ExecutionDurationSeconds types.Int64 `tfsdk:"execution_duration_seconds"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File types.Object `tfsdk:"file"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs types.Object `tfsdk:"gcs"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 types.Object `tfsdk:"s3"`
	// The current status of the script
	Status types.String `tfsdk:"status"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes types.Object `tfsdk:"volumes"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace types.Object `tfsdk:"workspace"`
}

func (to *InitScriptInfoAndExecutionDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InitScriptInfoAndExecutionDetails) {
	if !from.Abfss.IsNull() && !from.Abfss.IsUnknown() {
		if toAbfss, ok := to.GetAbfss(ctx); ok {
			if fromAbfss, ok := from.GetAbfss(ctx); ok {
				// Recursively sync the fields of Abfss
				toAbfss.SyncFieldsDuringCreateOrUpdate(ctx, fromAbfss)
				to.SetAbfss(ctx, toAbfss)
			}
		}
	}
	if !from.Dbfs.IsNull() && !from.Dbfs.IsUnknown() {
		if toDbfs, ok := to.GetDbfs(ctx); ok {
			if fromDbfs, ok := from.GetDbfs(ctx); ok {
				// Recursively sync the fields of Dbfs
				toDbfs.SyncFieldsDuringCreateOrUpdate(ctx, fromDbfs)
				to.SetDbfs(ctx, toDbfs)
			}
		}
	}
	if !from.File.IsNull() && !from.File.IsUnknown() {
		if toFile, ok := to.GetFile(ctx); ok {
			if fromFile, ok := from.GetFile(ctx); ok {
				// Recursively sync the fields of File
				toFile.SyncFieldsDuringCreateOrUpdate(ctx, fromFile)
				to.SetFile(ctx, toFile)
			}
		}
	}
	if !from.Gcs.IsNull() && !from.Gcs.IsUnknown() {
		if toGcs, ok := to.GetGcs(ctx); ok {
			if fromGcs, ok := from.GetGcs(ctx); ok {
				// Recursively sync the fields of Gcs
				toGcs.SyncFieldsDuringCreateOrUpdate(ctx, fromGcs)
				to.SetGcs(ctx, toGcs)
			}
		}
	}
	if !from.S3.IsNull() && !from.S3.IsUnknown() {
		if toS3, ok := to.GetS3(ctx); ok {
			if fromS3, ok := from.GetS3(ctx); ok {
				// Recursively sync the fields of S3
				toS3.SyncFieldsDuringCreateOrUpdate(ctx, fromS3)
				to.SetS3(ctx, toS3)
			}
		}
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() {
		if toVolumes, ok := to.GetVolumes(ctx); ok {
			if fromVolumes, ok := from.GetVolumes(ctx); ok {
				// Recursively sync the fields of Volumes
				toVolumes.SyncFieldsDuringCreateOrUpdate(ctx, fromVolumes)
				to.SetVolumes(ctx, toVolumes)
			}
		}
	}
	if !from.Workspace.IsNull() && !from.Workspace.IsUnknown() {
		if toWorkspace, ok := to.GetWorkspace(ctx); ok {
			if fromWorkspace, ok := from.GetWorkspace(ctx); ok {
				// Recursively sync the fields of Workspace
				toWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspace)
				to.SetWorkspace(ctx, toWorkspace)
			}
		}
	}
}

func (to *InitScriptInfoAndExecutionDetails) SyncFieldsDuringRead(ctx context.Context, from InitScriptInfoAndExecutionDetails) {
	if !from.Abfss.IsNull() && !from.Abfss.IsUnknown() {
		if toAbfss, ok := to.GetAbfss(ctx); ok {
			if fromAbfss, ok := from.GetAbfss(ctx); ok {
				toAbfss.SyncFieldsDuringRead(ctx, fromAbfss)
				to.SetAbfss(ctx, toAbfss)
			}
		}
	}
	if !from.Dbfs.IsNull() && !from.Dbfs.IsUnknown() {
		if toDbfs, ok := to.GetDbfs(ctx); ok {
			if fromDbfs, ok := from.GetDbfs(ctx); ok {
				toDbfs.SyncFieldsDuringRead(ctx, fromDbfs)
				to.SetDbfs(ctx, toDbfs)
			}
		}
	}
	if !from.File.IsNull() && !from.File.IsUnknown() {
		if toFile, ok := to.GetFile(ctx); ok {
			if fromFile, ok := from.GetFile(ctx); ok {
				toFile.SyncFieldsDuringRead(ctx, fromFile)
				to.SetFile(ctx, toFile)
			}
		}
	}
	if !from.Gcs.IsNull() && !from.Gcs.IsUnknown() {
		if toGcs, ok := to.GetGcs(ctx); ok {
			if fromGcs, ok := from.GetGcs(ctx); ok {
				toGcs.SyncFieldsDuringRead(ctx, fromGcs)
				to.SetGcs(ctx, toGcs)
			}
		}
	}
	if !from.S3.IsNull() && !from.S3.IsUnknown() {
		if toS3, ok := to.GetS3(ctx); ok {
			if fromS3, ok := from.GetS3(ctx); ok {
				toS3.SyncFieldsDuringRead(ctx, fromS3)
				to.SetS3(ctx, toS3)
			}
		}
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() {
		if toVolumes, ok := to.GetVolumes(ctx); ok {
			if fromVolumes, ok := from.GetVolumes(ctx); ok {
				toVolumes.SyncFieldsDuringRead(ctx, fromVolumes)
				to.SetVolumes(ctx, toVolumes)
			}
		}
	}
	if !from.Workspace.IsNull() && !from.Workspace.IsUnknown() {
		if toWorkspace, ok := to.GetWorkspace(ctx); ok {
			if fromWorkspace, ok := from.GetWorkspace(ctx); ok {
				toWorkspace.SyncFieldsDuringRead(ctx, fromWorkspace)
				to.SetWorkspace(ctx, toWorkspace)
			}
		}
	}
}

func (m InitScriptInfoAndExecutionDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["abfss"] = attrs["abfss"].SetOptional()
	attrs["dbfs"] = attrs["dbfs"].SetOptional()
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["execution_duration_seconds"] = attrs["execution_duration_seconds"].SetOptional()
	attrs["file"] = attrs["file"].SetOptional()
	attrs["gcs"] = attrs["gcs"].SetOptional()
	attrs["s3"] = attrs["s3"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["volumes"] = attrs["volumes"].SetOptional()
	attrs["workspace"] = attrs["workspace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InitScriptInfoAndExecutionDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InitScriptInfoAndExecutionDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"abfss":     reflect.TypeOf(Adlsgen2Info{}),
		"dbfs":      reflect.TypeOf(DbfsStorageInfo{}),
		"file":      reflect.TypeOf(LocalFileInfo{}),
		"gcs":       reflect.TypeOf(GcsStorageInfo{}),
		"s3":        reflect.TypeOf(S3StorageInfo{}),
		"volumes":   reflect.TypeOf(VolumesStorageInfo{}),
		"workspace": reflect.TypeOf(WorkspaceStorageInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InitScriptInfoAndExecutionDetails
// only implements ToObjectValue() and Type().
func (m InitScriptInfoAndExecutionDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"abfss":                      m.Abfss,
			"dbfs":                       m.Dbfs,
			"error_message":              m.ErrorMessage,
			"execution_duration_seconds": m.ExecutionDurationSeconds,
			"file":                       m.File,
			"gcs":                        m.Gcs,
			"s3":                         m.S3,
			"status":                     m.Status,
			"volumes":                    m.Volumes,
			"workspace":                  m.Workspace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InitScriptInfoAndExecutionDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"abfss":                      Adlsgen2Info{}.Type(ctx),
			"dbfs":                       DbfsStorageInfo{}.Type(ctx),
			"error_message":              types.StringType,
			"execution_duration_seconds": types.Int64Type,
			"file":                       LocalFileInfo{}.Type(ctx),
			"gcs":                        GcsStorageInfo{}.Type(ctx),
			"s3":                         S3StorageInfo{}.Type(ctx),
			"status":                     types.StringType,
			"volumes":                    VolumesStorageInfo{}.Type(ctx),
			"workspace":                  WorkspaceStorageInfo{}.Type(ctx),
		},
	}
}

// GetAbfss returns the value of the Abfss field in InitScriptInfoAndExecutionDetails as
// a Adlsgen2Info value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfoAndExecutionDetails) GetAbfss(ctx context.Context) (Adlsgen2Info, bool) {
	var e Adlsgen2Info
	if m.Abfss.IsNull() || m.Abfss.IsUnknown() {
		return e, false
	}
	var v Adlsgen2Info
	d := m.Abfss.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAbfss sets the value of the Abfss field in InitScriptInfoAndExecutionDetails.
func (m *InitScriptInfoAndExecutionDetails) SetAbfss(ctx context.Context, v Adlsgen2Info) {
	vs := v.ToObjectValue(ctx)
	m.Abfss = vs
}

// GetDbfs returns the value of the Dbfs field in InitScriptInfoAndExecutionDetails as
// a DbfsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfoAndExecutionDetails) GetDbfs(ctx context.Context) (DbfsStorageInfo, bool) {
	var e DbfsStorageInfo
	if m.Dbfs.IsNull() || m.Dbfs.IsUnknown() {
		return e, false
	}
	var v DbfsStorageInfo
	d := m.Dbfs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDbfs sets the value of the Dbfs field in InitScriptInfoAndExecutionDetails.
func (m *InitScriptInfoAndExecutionDetails) SetDbfs(ctx context.Context, v DbfsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Dbfs = vs
}

// GetFile returns the value of the File field in InitScriptInfoAndExecutionDetails as
// a LocalFileInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfoAndExecutionDetails) GetFile(ctx context.Context) (LocalFileInfo, bool) {
	var e LocalFileInfo
	if m.File.IsNull() || m.File.IsUnknown() {
		return e, false
	}
	var v LocalFileInfo
	d := m.File.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFile sets the value of the File field in InitScriptInfoAndExecutionDetails.
func (m *InitScriptInfoAndExecutionDetails) SetFile(ctx context.Context, v LocalFileInfo) {
	vs := v.ToObjectValue(ctx)
	m.File = vs
}

// GetGcs returns the value of the Gcs field in InitScriptInfoAndExecutionDetails as
// a GcsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfoAndExecutionDetails) GetGcs(ctx context.Context) (GcsStorageInfo, bool) {
	var e GcsStorageInfo
	if m.Gcs.IsNull() || m.Gcs.IsUnknown() {
		return e, false
	}
	var v GcsStorageInfo
	d := m.Gcs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcs sets the value of the Gcs field in InitScriptInfoAndExecutionDetails.
func (m *InitScriptInfoAndExecutionDetails) SetGcs(ctx context.Context, v GcsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Gcs = vs
}

// GetS3 returns the value of the S3 field in InitScriptInfoAndExecutionDetails as
// a S3StorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfoAndExecutionDetails) GetS3(ctx context.Context) (S3StorageInfo, bool) {
	var e S3StorageInfo
	if m.S3.IsNull() || m.S3.IsUnknown() {
		return e, false
	}
	var v S3StorageInfo
	d := m.S3.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetS3 sets the value of the S3 field in InitScriptInfoAndExecutionDetails.
func (m *InitScriptInfoAndExecutionDetails) SetS3(ctx context.Context, v S3StorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.S3 = vs
}

// GetVolumes returns the value of the Volumes field in InitScriptInfoAndExecutionDetails as
// a VolumesStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfoAndExecutionDetails) GetVolumes(ctx context.Context) (VolumesStorageInfo, bool) {
	var e VolumesStorageInfo
	if m.Volumes.IsNull() || m.Volumes.IsUnknown() {
		return e, false
	}
	var v VolumesStorageInfo
	d := m.Volumes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumes sets the value of the Volumes field in InitScriptInfoAndExecutionDetails.
func (m *InitScriptInfoAndExecutionDetails) SetVolumes(ctx context.Context, v VolumesStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Volumes = vs
}

// GetWorkspace returns the value of the Workspace field in InitScriptInfoAndExecutionDetails as
// a WorkspaceStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *InitScriptInfoAndExecutionDetails) GetWorkspace(ctx context.Context) (WorkspaceStorageInfo, bool) {
	var e WorkspaceStorageInfo
	if m.Workspace.IsNull() || m.Workspace.IsUnknown() {
		return e, false
	}
	var v WorkspaceStorageInfo
	d := m.Workspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspace sets the value of the Workspace field in InitScriptInfoAndExecutionDetails.
func (m *InitScriptInfoAndExecutionDetails) SetWorkspace(ctx context.Context, v WorkspaceStorageInfo) {
	vs := v.ToObjectValue(ctx)
	m.Workspace = vs
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The libraries to install.
	Libraries types.List `tfsdk:"libraries"`
}

func (to *InstallLibraries) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstallLibraries) {
}

func (to *InstallLibraries) SyncFieldsDuringRead(ctx context.Context, from InstallLibraries) {
}

func (m InstallLibraries) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["libraries"] = attrs["libraries"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstallLibraries.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstallLibraries) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallLibraries
// only implements ToObjectValue() and Type().
func (m InstallLibraries) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
			"libraries":  m.Libraries,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstallLibraries) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"libraries": basetypes.ListType{
				ElemType: Library{}.Type(ctx),
			},
		},
	}
}

// GetLibraries returns the value of the Libraries field in InstallLibraries as
// a slice of Library values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstallLibraries) GetLibraries(ctx context.Context) ([]Library, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in InstallLibraries.
func (m *InstallLibraries) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

type InstallLibrariesResponse struct {
}

func (to *InstallLibrariesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstallLibrariesResponse) {
}

func (to *InstallLibrariesResponse) SyncFieldsDuringRead(ctx context.Context, from InstallLibrariesResponse) {
}

func (m InstallLibrariesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstallLibrariesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstallLibrariesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallLibrariesResponse
// only implements ToObjectValue() and Type().
func (m InstallLibrariesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m InstallLibrariesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type InstancePoolAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *InstancePoolAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolAccessControlRequest) {
}

func (to *InstancePoolAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from InstancePoolAccessControlRequest) {
}

func (m InstancePoolAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAccessControlRequest
// only implements ToObjectValue() and Type().
func (m InstancePoolAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type InstancePoolAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *InstancePoolAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *InstancePoolAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from InstancePoolAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m InstancePoolAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(InstancePoolPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAccessControlResponse
// only implements ToObjectValue() and Type().
func (m InstancePoolAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: InstancePoolPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in InstancePoolAccessControlResponse as
// a slice of InstancePoolPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAccessControlResponse) GetAllPermissions(ctx context.Context) ([]InstancePoolPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in InstancePoolAccessControlResponse.
func (m *InstancePoolAccessControlResponse) SetAllPermissions(ctx context.Context, v []InstancePoolPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type InstancePoolAndStats struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags types.Map `tfsdk:"custom_tags"`
	// Tags that are added by Databricks regardless of any ``custom_tags``,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
	DefaultTags types.Map `tfsdk:"default_tags"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec types.Object `tfsdk:"disk_spec"`
	// For pools with node type flexibility (Fleet-V2), whether auto generated
	// alternate node type ids are enabled. This field should not be true if
	// node_type_flexibility is set.
	EnableAutoAlternateNodeTypes types.Bool `tfsdk:"enable_auto_alternate_node_types"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes"`
	// Canonical unique identifier for the pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances"`
	// For pools with node type flexibility (Fleet-V2), this object contains the
	// information about the alternate node type ids to use when attempting to
	// launch a cluster if the node type id is not available. This field should
	// not be set if enable_auto_alternate_node_types is true.
	NodeTypeFlexibility types.Object `tfsdk:"node_type_flexibility"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Custom Docker Image BYOC
	PreloadedDockerImages types.List `tfsdk:"preloaded_docker_images"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions types.List `tfsdk:"preloaded_spark_versions"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// Current state of the instance pool.
	State types.String `tfsdk:"state"`
	// Usage statistics about the instance pool.
	Stats types.Object `tfsdk:"stats"`
	// Status of failed pending instances in the pool.
	Status types.Object `tfsdk:"status"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
}

func (to *InstancePoolAndStats) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolAndStats) {
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.DiskSpec.IsNull() && !from.DiskSpec.IsUnknown() {
		if toDiskSpec, ok := to.GetDiskSpec(ctx); ok {
			if fromDiskSpec, ok := from.GetDiskSpec(ctx); ok {
				// Recursively sync the fields of DiskSpec
				toDiskSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDiskSpec)
				to.SetDiskSpec(ctx, toDiskSpec)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.NodeTypeFlexibility.IsNull() && !from.NodeTypeFlexibility.IsUnknown() {
		if toNodeTypeFlexibility, ok := to.GetNodeTypeFlexibility(ctx); ok {
			if fromNodeTypeFlexibility, ok := from.GetNodeTypeFlexibility(ctx); ok {
				// Recursively sync the fields of NodeTypeFlexibility
				toNodeTypeFlexibility.SyncFieldsDuringCreateOrUpdate(ctx, fromNodeTypeFlexibility)
				to.SetNodeTypeFlexibility(ctx, toNodeTypeFlexibility)
			}
		}
	}
	if !from.PreloadedDockerImages.IsNull() && !from.PreloadedDockerImages.IsUnknown() && to.PreloadedDockerImages.IsNull() && len(from.PreloadedDockerImages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedDockerImages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedDockerImages = from.PreloadedDockerImages
	}
	if !from.PreloadedSparkVersions.IsNull() && !from.PreloadedSparkVersions.IsUnknown() && to.PreloadedSparkVersions.IsNull() && len(from.PreloadedSparkVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedSparkVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedSparkVersions = from.PreloadedSparkVersions
	}
	if !from.Stats.IsNull() && !from.Stats.IsUnknown() {
		if toStats, ok := to.GetStats(ctx); ok {
			if fromStats, ok := from.GetStats(ctx); ok {
				// Recursively sync the fields of Stats
				toStats.SyncFieldsDuringCreateOrUpdate(ctx, fromStats)
				to.SetStats(ctx, toStats)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *InstancePoolAndStats) SyncFieldsDuringRead(ctx context.Context, from InstancePoolAndStats) {
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.DiskSpec.IsNull() && !from.DiskSpec.IsUnknown() {
		if toDiskSpec, ok := to.GetDiskSpec(ctx); ok {
			if fromDiskSpec, ok := from.GetDiskSpec(ctx); ok {
				toDiskSpec.SyncFieldsDuringRead(ctx, fromDiskSpec)
				to.SetDiskSpec(ctx, toDiskSpec)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.NodeTypeFlexibility.IsNull() && !from.NodeTypeFlexibility.IsUnknown() {
		if toNodeTypeFlexibility, ok := to.GetNodeTypeFlexibility(ctx); ok {
			if fromNodeTypeFlexibility, ok := from.GetNodeTypeFlexibility(ctx); ok {
				toNodeTypeFlexibility.SyncFieldsDuringRead(ctx, fromNodeTypeFlexibility)
				to.SetNodeTypeFlexibility(ctx, toNodeTypeFlexibility)
			}
		}
	}
	if !from.PreloadedDockerImages.IsNull() && !from.PreloadedDockerImages.IsUnknown() && to.PreloadedDockerImages.IsNull() && len(from.PreloadedDockerImages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedDockerImages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedDockerImages = from.PreloadedDockerImages
	}
	if !from.PreloadedSparkVersions.IsNull() && !from.PreloadedSparkVersions.IsUnknown() && to.PreloadedSparkVersions.IsNull() && len(from.PreloadedSparkVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PreloadedSparkVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PreloadedSparkVersions = from.PreloadedSparkVersions
	}
	if !from.Stats.IsNull() && !from.Stats.IsUnknown() {
		if toStats, ok := to.GetStats(ctx); ok {
			if fromStats, ok := from.GetStats(ctx); ok {
				toStats.SyncFieldsDuringRead(ctx, fromStats)
				to.SetStats(ctx, toStats)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m InstancePoolAndStats) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["default_tags"] = attrs["default_tags"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].SetOptional()
	attrs["enable_auto_alternate_node_types"] = attrs["enable_auto_alternate_node_types"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["idle_instance_autotermination_minutes"] = attrs["idle_instance_autotermination_minutes"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["instance_pool_name"] = attrs["instance_pool_name"].SetOptional()
	attrs["max_capacity"] = attrs["max_capacity"].SetOptional()
	attrs["min_idle_instances"] = attrs["min_idle_instances"].SetOptional()
	attrs["node_type_flexibility"] = attrs["node_type_flexibility"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["preloaded_docker_images"] = attrs["preloaded_docker_images"].SetOptional()
	attrs["preloaded_spark_versions"] = attrs["preloaded_spark_versions"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["stats"] = attrs["stats"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolAndStats.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolAndStats) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"default_tags":             reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes{}),
		"node_type_flexibility":    reflect.TypeOf(NodeTypeFlexibility{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
		"stats":                    reflect.TypeOf(InstancePoolStats{}),
		"status":                   reflect.TypeOf(InstancePoolStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAndStats
// only implements ToObjectValue() and Type().
func (m InstancePoolAndStats) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        m.AwsAttributes,
			"azure_attributes":                      m.AzureAttributes,
			"custom_tags":                           m.CustomTags,
			"default_tags":                          m.DefaultTags,
			"disk_spec":                             m.DiskSpec,
			"enable_auto_alternate_node_types":      m.EnableAutoAlternateNodeTypes,
			"enable_elastic_disk":                   m.EnableElasticDisk,
			"gcp_attributes":                        m.GcpAttributes,
			"idle_instance_autotermination_minutes": m.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      m.InstancePoolId,
			"instance_pool_name":                    m.InstancePoolName,
			"max_capacity":                          m.MaxCapacity,
			"min_idle_instances":                    m.MinIdleInstances,
			"node_type_flexibility":                 m.NodeTypeFlexibility,
			"node_type_id":                          m.NodeTypeId,
			"preloaded_docker_images":               m.PreloadedDockerImages,
			"preloaded_spark_versions":              m.PreloadedSparkVersions,
			"remote_disk_throughput":                m.RemoteDiskThroughput,
			"state":                                 m.State,
			"stats":                                 m.Stats,
			"status":                                m.Status,
			"total_initial_remote_disk_size":        m.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolAndStats) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes":   InstancePoolAwsAttributes{}.Type(ctx),
			"azure_attributes": InstancePoolAzureAttributes{}.Type(ctx),
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"default_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec":                             DiskSpec{}.Type(ctx),
			"enable_auto_alternate_node_types":      types.BoolType,
			"enable_elastic_disk":                   types.BoolType,
			"gcp_attributes":                        InstancePoolGcpAttributes{}.Type(ctx),
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_id":                      types.StringType,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
			"node_type_flexibility":                 NodeTypeFlexibility{}.Type(ctx),
			"node_type_id":                          types.StringType,
			"preloaded_docker_images": basetypes.ListType{
				ElemType: DockerImage{}.Type(ctx),
			},
			"preloaded_spark_versions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"remote_disk_throughput":         types.Int64Type,
			"state":                          types.StringType,
			"stats":                          InstancePoolStats{}.Type(ctx),
			"status":                         InstancePoolStatus{}.Type(ctx),
			"total_initial_remote_disk_size": types.Int64Type,
		},
	}
}

// GetAwsAttributes returns the value of the AwsAttributes field in InstancePoolAndStats as
// a InstancePoolAwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes, bool) {
	var e InstancePoolAwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolAwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in InstancePoolAndStats as
// a InstancePoolAzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes, bool) {
	var e InstancePoolAzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolAzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetCustomTags returns the value of the CustomTags field in InstancePoolAndStats as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in InstancePoolAndStats as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
	if m.DefaultTags.IsNull() || m.DefaultTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.DefaultTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultTags sets the value of the DefaultTags field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DefaultTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in InstancePoolAndStats as
// a DiskSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetDiskSpec(ctx context.Context) (DiskSpec, bool) {
	var e DiskSpec
	if m.DiskSpec.IsNull() || m.DiskSpec.IsUnknown() {
		return e, false
	}
	var v DiskSpec
	d := m.DiskSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDiskSpec sets the value of the DiskSpec field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetDiskSpec(ctx context.Context, v DiskSpec) {
	vs := v.ToObjectValue(ctx)
	m.DiskSpec = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in InstancePoolAndStats as
// a InstancePoolGcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes, bool) {
	var e InstancePoolGcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v InstancePoolGcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetNodeTypeFlexibility returns the value of the NodeTypeFlexibility field in InstancePoolAndStats as
// a NodeTypeFlexibility value.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetNodeTypeFlexibility(ctx context.Context) (NodeTypeFlexibility, bool) {
	var e NodeTypeFlexibility
	if m.NodeTypeFlexibility.IsNull() || m.NodeTypeFlexibility.IsUnknown() {
		return e, false
	}
	var v NodeTypeFlexibility
	d := m.NodeTypeFlexibility.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeTypeFlexibility sets the value of the NodeTypeFlexibility field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetNodeTypeFlexibility(ctx context.Context, v NodeTypeFlexibility) {
	vs := v.ToObjectValue(ctx)
	m.NodeTypeFlexibility = vs
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in InstancePoolAndStats as
// a slice of DockerImage values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage, bool) {
	if m.PreloadedDockerImages.IsNull() || m.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage
	d := m.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetPreloadedDockerImages(ctx context.Context, v []DockerImage) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in InstancePoolAndStats as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
	if m.PreloadedSparkVersions.IsNull() || m.PreloadedSparkVersions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.PreloadedSparkVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

// GetStats returns the value of the Stats field in InstancePoolAndStats as
// a InstancePoolStats value.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetStats(ctx context.Context) (InstancePoolStats, bool) {
	var e InstancePoolStats
	if m.Stats.IsNull() || m.Stats.IsUnknown() {
		return e, false
	}
	var v InstancePoolStats
	d := m.Stats.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStats sets the value of the Stats field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetStats(ctx context.Context, v InstancePoolStats) {
	vs := v.ToObjectValue(ctx)
	m.Stats = vs
}

// GetStatus returns the value of the Status field in InstancePoolAndStats as
// a InstancePoolStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolAndStats) GetStatus(ctx context.Context) (InstancePoolStatus, bool) {
	var e InstancePoolStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v InstancePoolStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in InstancePoolAndStats.
func (m *InstancePoolAndStats) SetStatus(ctx context.Context, v InstancePoolStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// Attributes set during instance pool creation which are related to Amazon Web
// Services.
type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	Availability types.String `tfsdk:"availability"`
	// All AWS instances belonging to the instance pool will have this instance
	// profile. If omitted, instances will initially be launched with the
	// workspace's default instance profile. If defined, clusters that use the
	// pool will inherit the instance profile, and must not specify their own
	// instance profile on cluster creation or update. If the pool does not
	// specify an instance profile, clusters using the pool may specify any
	// instance profile. The instance profile must have previously been added to
	// the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Calculates the bid price for AWS spot instances, as a percentage of the
	// corresponding instance type's on-demand price. For example, if this field
	// is set to 50, and the cluster needs a new `r3.xlarge` spot instance, then
	// the bid price is half of the price of on-demand `r3.xlarge` instances.
	// Similarly, if this field is set to 200, the bid price is twice the price
	// of on-demand `r3.xlarge` instances. If not specified, the default value
	// is 100. When spot instances are requested for this cluster, only spot
	// instances whose bid price percentage matches this field will be
	// considered. Note that, for safety, we enforce this field to be no more
	// than 10000.
	SpotBidPricePercent types.Int64 `tfsdk:"spot_bid_price_percent"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. The list of available zones as well as the default value
	// can be found by using the `List Zones` method.
	ZoneId types.String `tfsdk:"zone_id"`
}

func (to *InstancePoolAwsAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolAwsAttributes) {
}

func (to *InstancePoolAwsAttributes) SyncFieldsDuringRead(ctx context.Context, from InstancePoolAwsAttributes) {
}

func (m InstancePoolAwsAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability"] = attrs["availability"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["spot_bid_price_percent"] = attrs["spot_bid_price_percent"].SetOptional()
	attrs["zone_id"] = attrs["zone_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolAwsAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolAwsAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAwsAttributes
// only implements ToObjectValue() and Type().
func (m InstancePoolAwsAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":           m.Availability,
			"instance_profile_arn":   m.InstanceProfileArn,
			"spot_bid_price_percent": m.SpotBidPricePercent,
			"zone_id":                m.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolAwsAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":           types.StringType,
			"instance_profile_arn":   types.StringType,
			"spot_bid_price_percent": types.Int64Type,
			"zone_id":                types.StringType,
		},
	}
}

// Attributes set during instance pool creation which are related to Azure.
type InstancePoolAzureAttributes struct {
	// Availability type used for the spot nodes.
	Availability types.String `tfsdk:"availability"`
	// With variable pricing, you have option to set a max price, in US dollars
	// (USD) For example, the value 2 would be a max price of $2.00 USD per
	// hour. If you set the max price to be -1, the VM won't be evicted based on
	// price. The price for the VM will be the current price for spot or the
	// price for a standard VM, which ever is less, as long as there is capacity
	// and quota available.
	SpotBidMaxPrice types.Float64 `tfsdk:"spot_bid_max_price"`
}

func (to *InstancePoolAzureAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolAzureAttributes) {
}

func (to *InstancePoolAzureAttributes) SyncFieldsDuringRead(ctx context.Context, from InstancePoolAzureAttributes) {
}

func (m InstancePoolAzureAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability"] = attrs["availability"].SetOptional()
	attrs["spot_bid_max_price"] = attrs["spot_bid_max_price"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolAzureAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolAzureAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAzureAttributes
// only implements ToObjectValue() and Type().
func (m InstancePoolAzureAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":       m.Availability,
			"spot_bid_max_price": m.SpotBidMaxPrice,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolAzureAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":       types.StringType,
			"spot_bid_max_price": types.Float64Type,
		},
	}
}

// Attributes set during instance pool creation which are related to GCP.
type InstancePoolGcpAttributes struct {
	GcpAvailability types.String `tfsdk:"gcp_availability"`
	// If provided, each node in the instance pool will have this number of
	// local SSDs attached. Each local SSD is 375GB in size. Refer to [GCP
	// documentation] for the supported number of local SSDs for each instance
	// type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount types.Int64 `tfsdk:"local_ssd_count"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west1-a". The provided
	// availability zone must be in the same region as the Databricks workspace.
	// For example, "us-west1-a" is not a valid zone id if the Databricks
	// workspace resides in the "us-east1" region. This is an optional field at
	// instance pool creation, and if not specified, a default zone will be
	// used.
	//
	// This field can be one of the following: - "HA" => High availability,
	// spread nodes across availability zones for a Databricks deployment region
	// - A GCP availability zone => Pick One of the available zones for (machine
	// type + region) from https://cloud.google.com/compute/docs/regions-zones
	// (e.g. "us-west1-a").
	//
	// If empty, Databricks picks an availability zone to schedule the cluster
	// on.
	ZoneId types.String `tfsdk:"zone_id"`
}

func (to *InstancePoolGcpAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolGcpAttributes) {
}

func (to *InstancePoolGcpAttributes) SyncFieldsDuringRead(ctx context.Context, from InstancePoolGcpAttributes) {
}

func (m InstancePoolGcpAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_availability"] = attrs["gcp_availability"].SetOptional()
	attrs["local_ssd_count"] = attrs["local_ssd_count"].SetOptional()
	attrs["zone_id"] = attrs["zone_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolGcpAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolGcpAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolGcpAttributes
// only implements ToObjectValue() and Type().
func (m InstancePoolGcpAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_availability": m.GcpAvailability,
			"local_ssd_count":  m.LocalSsdCount,
			"zone_id":          m.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolGcpAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_availability": types.StringType,
			"local_ssd_count":  types.Int64Type,
			"zone_id":          types.StringType,
		},
	}
}

type InstancePoolPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *InstancePoolPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *InstancePoolPermission) SyncFieldsDuringRead(ctx context.Context, from InstancePoolPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m InstancePoolPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermission
// only implements ToObjectValue() and Type().
func (m InstancePoolPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolPermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// GetInheritedFromObject returns the value of the InheritedFromObject field in InstancePoolPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in InstancePoolPermission.
func (m *InstancePoolPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type InstancePoolPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *InstancePoolPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *InstancePoolPermissions) SyncFieldsDuringRead(ctx context.Context, from InstancePoolPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m InstancePoolPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(InstancePoolAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissions
// only implements ToObjectValue() and Type().
func (m InstancePoolPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: InstancePoolAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in InstancePoolPermissions as
// a slice of InstancePoolAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolPermissions) GetAccessControlList(ctx context.Context) ([]InstancePoolAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in InstancePoolPermissions.
func (m *InstancePoolPermissions) SetAccessControlList(ctx context.Context, v []InstancePoolAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type InstancePoolPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *InstancePoolPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolPermissionsDescription) {
}

func (to *InstancePoolPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from InstancePoolPermissionsDescription) {
}

func (m InstancePoolPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissionsDescription
// only implements ToObjectValue() and Type().
func (m InstancePoolPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type InstancePoolPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-"`
}

func (to *InstancePoolPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *InstancePoolPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from InstancePoolPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m InstancePoolPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(InstancePoolAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissionsRequest
// only implements ToObjectValue() and Type().
func (m InstancePoolPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"instance_pool_id":    m.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: InstancePoolAccessControlRequest{}.Type(ctx),
			},
			"instance_pool_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in InstancePoolPermissionsRequest as
// a slice of InstancePoolAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolPermissionsRequest) GetAccessControlList(ctx context.Context) ([]InstancePoolAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in InstancePoolPermissionsRequest.
func (m *InstancePoolPermissionsRequest) SetAccessControlList(ctx context.Context, v []InstancePoolAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type InstancePoolStats struct {
	// Number of active instances in the pool that are NOT part of a cluster.
	IdleCount types.Int64 `tfsdk:"idle_count"`
	// Number of pending instances in the pool that are NOT part of a cluster.
	PendingIdleCount types.Int64 `tfsdk:"pending_idle_count"`
	// Number of pending instances in the pool that are part of a cluster.
	PendingUsedCount types.Int64 `tfsdk:"pending_used_count"`
	// Number of active instances in the pool that are part of a cluster.
	UsedCount types.Int64 `tfsdk:"used_count"`
}

func (to *InstancePoolStats) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolStats) {
}

func (to *InstancePoolStats) SyncFieldsDuringRead(ctx context.Context, from InstancePoolStats) {
}

func (m InstancePoolStats) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["idle_count"] = attrs["idle_count"].SetOptional()
	attrs["pending_idle_count"] = attrs["pending_idle_count"].SetOptional()
	attrs["pending_used_count"] = attrs["pending_used_count"].SetOptional()
	attrs["used_count"] = attrs["used_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolStats.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolStats) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolStats
// only implements ToObjectValue() and Type().
func (m InstancePoolStats) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"idle_count":         m.IdleCount,
			"pending_idle_count": m.PendingIdleCount,
			"pending_used_count": m.PendingUsedCount,
			"used_count":         m.UsedCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolStats) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"idle_count":         types.Int64Type,
			"pending_idle_count": types.Int64Type,
			"pending_used_count": types.Int64Type,
			"used_count":         types.Int64Type,
		},
	}
}

type InstancePoolStatus struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	PendingInstanceErrors types.List `tfsdk:"pending_instance_errors"`
}

func (to *InstancePoolStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstancePoolStatus) {
	if !from.PendingInstanceErrors.IsNull() && !from.PendingInstanceErrors.IsUnknown() && to.PendingInstanceErrors.IsNull() && len(from.PendingInstanceErrors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PendingInstanceErrors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PendingInstanceErrors = from.PendingInstanceErrors
	}
}

func (to *InstancePoolStatus) SyncFieldsDuringRead(ctx context.Context, from InstancePoolStatus) {
	if !from.PendingInstanceErrors.IsNull() && !from.PendingInstanceErrors.IsUnknown() && to.PendingInstanceErrors.IsNull() && len(from.PendingInstanceErrors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PendingInstanceErrors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PendingInstanceErrors = from.PendingInstanceErrors
	}
}

func (m InstancePoolStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pending_instance_errors"] = attrs["pending_instance_errors"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstancePoolStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pending_instance_errors": reflect.TypeOf(PendingInstanceError{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolStatus
// only implements ToObjectValue() and Type().
func (m InstancePoolStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pending_instance_errors": m.PendingInstanceErrors,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstancePoolStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pending_instance_errors": basetypes.ListType{
				ElemType: PendingInstanceError{}.Type(ctx),
			},
		},
	}
}

// GetPendingInstanceErrors returns the value of the PendingInstanceErrors field in InstancePoolStatus as
// a slice of PendingInstanceError values.
// If the field is unknown or null, the boolean return value is false.
func (m *InstancePoolStatus) GetPendingInstanceErrors(ctx context.Context) ([]PendingInstanceError, bool) {
	if m.PendingInstanceErrors.IsNull() || m.PendingInstanceErrors.IsUnknown() {
		return nil, false
	}
	var v []PendingInstanceError
	d := m.PendingInstanceErrors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPendingInstanceErrors sets the value of the PendingInstanceErrors field in InstancePoolStatus.
func (m *InstancePoolStatus) SetPendingInstanceErrors(ctx context.Context, v []PendingInstanceError) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pending_instance_errors"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PendingInstanceErrors = types.ListValueMust(t, vs)
}

type InstanceProfile struct {
	// The AWS IAM role ARN of the role associated with the instance profile.
	// This field is required if your role name and instance profile name do not
	// match and you want to use the instance profile with [Databricks SQL
	// Serverless].
	//
	// Otherwise, this field is optional.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	IamRoleArn types.String `tfsdk:"iam_role_arn"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile types.Bool `tfsdk:"is_meta_instance_profile"`
}

func (to *InstanceProfile) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InstanceProfile) {
}

func (to *InstanceProfile) SyncFieldsDuringRead(ctx context.Context, from InstanceProfile) {
}

func (m InstanceProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["iam_role_arn"] = attrs["iam_role_arn"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetRequired()
	attrs["is_meta_instance_profile"] = attrs["is_meta_instance_profile"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstanceProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InstanceProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstanceProfile
// only implements ToObjectValue() and Type().
func (m InstanceProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"iam_role_arn":             m.IamRoleArn,
			"instance_profile_arn":     m.InstanceProfileArn,
			"is_meta_instance_profile": m.IsMetaInstanceProfile,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InstanceProfile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"iam_role_arn":             types.StringType,
			"instance_profile_arn":     types.StringType,
			"is_meta_instance_profile": types.BoolType,
		},
	}
}

type Library struct {
	// Specification of a CRAN library to be installed as part of the library
	Cran types.Object `tfsdk:"cran"`
	// Deprecated. URI of the egg library to install. Installing Python egg
	// files is deprecated and is not supported in Databricks Runtime 14.0 and
	// above.
	Egg types.String `tfsdk:"egg"`
	// URI of the JAR library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "jar":
	// "/Workspace/path/to/library.jar" }`, `{ "jar" :
	// "/Volumes/path/to/library.jar" }` or `{ "jar":
	// "s3://my-bucket/library.jar" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Jar types.String `tfsdk:"jar"`
	// Specification of a maven library to be installed. For example: `{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }`
	Maven types.Object `tfsdk:"maven"`
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	Pypi types.Object `tfsdk:"pypi"`
	// URI of the requirements.txt file to install. Only Workspace paths and
	// Unity Catalog Volumes paths are supported. For example: `{
	// "requirements": "/Workspace/path/to/requirements.txt" }` or `{
	// "requirements" : "/Volumes/path/to/requirements.txt" }`
	Requirements types.String `tfsdk:"requirements"`
	// URI of the wheel library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "whl":
	// "/Workspace/path/to/library.whl" }`, `{ "whl" :
	// "/Volumes/path/to/library.whl" }` or `{ "whl":
	// "s3://my-bucket/library.whl" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Whl types.String `tfsdk:"whl"`
}

func (to *Library) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Library) {
	if !from.Cran.IsNull() && !from.Cran.IsUnknown() {
		if toCran, ok := to.GetCran(ctx); ok {
			if fromCran, ok := from.GetCran(ctx); ok {
				// Recursively sync the fields of Cran
				toCran.SyncFieldsDuringCreateOrUpdate(ctx, fromCran)
				to.SetCran(ctx, toCran)
			}
		}
	}
	if !from.Maven.IsNull() && !from.Maven.IsUnknown() {
		if toMaven, ok := to.GetMaven(ctx); ok {
			if fromMaven, ok := from.GetMaven(ctx); ok {
				// Recursively sync the fields of Maven
				toMaven.SyncFieldsDuringCreateOrUpdate(ctx, fromMaven)
				to.SetMaven(ctx, toMaven)
			}
		}
	}
	if !from.Pypi.IsNull() && !from.Pypi.IsUnknown() {
		if toPypi, ok := to.GetPypi(ctx); ok {
			if fromPypi, ok := from.GetPypi(ctx); ok {
				// Recursively sync the fields of Pypi
				toPypi.SyncFieldsDuringCreateOrUpdate(ctx, fromPypi)
				to.SetPypi(ctx, toPypi)
			}
		}
	}
}

func (to *Library) SyncFieldsDuringRead(ctx context.Context, from Library) {
	if !from.Cran.IsNull() && !from.Cran.IsUnknown() {
		if toCran, ok := to.GetCran(ctx); ok {
			if fromCran, ok := from.GetCran(ctx); ok {
				toCran.SyncFieldsDuringRead(ctx, fromCran)
				to.SetCran(ctx, toCran)
			}
		}
	}
	if !from.Maven.IsNull() && !from.Maven.IsUnknown() {
		if toMaven, ok := to.GetMaven(ctx); ok {
			if fromMaven, ok := from.GetMaven(ctx); ok {
				toMaven.SyncFieldsDuringRead(ctx, fromMaven)
				to.SetMaven(ctx, toMaven)
			}
		}
	}
	if !from.Pypi.IsNull() && !from.Pypi.IsUnknown() {
		if toPypi, ok := to.GetPypi(ctx); ok {
			if fromPypi, ok := from.GetPypi(ctx); ok {
				toPypi.SyncFieldsDuringRead(ctx, fromPypi)
				to.SetPypi(ctx, toPypi)
			}
		}
	}
}

func (m Library) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cran"] = attrs["cran"].SetOptional()
	attrs["egg"] = attrs["egg"].SetOptional()
	attrs["jar"] = attrs["jar"].SetOptional()
	attrs["maven"] = attrs["maven"].SetOptional()
	attrs["pypi"] = attrs["pypi"].SetOptional()
	attrs["requirements"] = attrs["requirements"].SetOptional()
	attrs["whl"] = attrs["whl"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Library.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Library) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cran":  reflect.TypeOf(RCranLibrary{}),
		"maven": reflect.TypeOf(MavenLibrary{}),
		"pypi":  reflect.TypeOf(PythonPyPiLibrary{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Library
// only implements ToObjectValue() and Type().
func (m Library) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cran":         m.Cran,
			"egg":          m.Egg,
			"jar":          m.Jar,
			"maven":        m.Maven,
			"pypi":         m.Pypi,
			"requirements": m.Requirements,
			"whl":          m.Whl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Library) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cran":         RCranLibrary{}.Type(ctx),
			"egg":          types.StringType,
			"jar":          types.StringType,
			"maven":        MavenLibrary{}.Type(ctx),
			"pypi":         PythonPyPiLibrary{}.Type(ctx),
			"requirements": types.StringType,
			"whl":          types.StringType,
		},
	}
}

// GetCran returns the value of the Cran field in Library as
// a RCranLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (m *Library) GetCran(ctx context.Context) (RCranLibrary, bool) {
	var e RCranLibrary
	if m.Cran.IsNull() || m.Cran.IsUnknown() {
		return e, false
	}
	var v RCranLibrary
	d := m.Cran.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCran sets the value of the Cran field in Library.
func (m *Library) SetCran(ctx context.Context, v RCranLibrary) {
	vs := v.ToObjectValue(ctx)
	m.Cran = vs
}

// GetMaven returns the value of the Maven field in Library as
// a MavenLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (m *Library) GetMaven(ctx context.Context) (MavenLibrary, bool) {
	var e MavenLibrary
	if m.Maven.IsNull() || m.Maven.IsUnknown() {
		return e, false
	}
	var v MavenLibrary
	d := m.Maven.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMaven sets the value of the Maven field in Library.
func (m *Library) SetMaven(ctx context.Context, v MavenLibrary) {
	vs := v.ToObjectValue(ctx)
	m.Maven = vs
}

// GetPypi returns the value of the Pypi field in Library as
// a PythonPyPiLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (m *Library) GetPypi(ctx context.Context) (PythonPyPiLibrary, bool) {
	var e PythonPyPiLibrary
	if m.Pypi.IsNull() || m.Pypi.IsUnknown() {
		return e, false
	}
	var v PythonPyPiLibrary
	d := m.Pypi.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPypi sets the value of the Pypi field in Library.
func (m *Library) SetPypi(ctx context.Context, v PythonPyPiLibrary) {
	vs := v.ToObjectValue(ctx)
	m.Pypi = vs
}

// The status of the library on a specific cluster.
type LibraryFullStatus struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	IsLibraryForAllClusters types.Bool `tfsdk:"is_library_for_all_clusters"`
	// Unique identifier for the library.
	Library types.Object `tfsdk:"library"`
	// All the info and warning messages that have occurred so far for this
	// library.
	Messages types.List `tfsdk:"messages"`
	// Status of installing the library on the cluster.
	Status types.String `tfsdk:"status"`
}

func (to *LibraryFullStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LibraryFullStatus) {
	if !from.Library.IsNull() && !from.Library.IsUnknown() {
		if toLibrary, ok := to.GetLibrary(ctx); ok {
			if fromLibrary, ok := from.GetLibrary(ctx); ok {
				// Recursively sync the fields of Library
				toLibrary.SyncFieldsDuringCreateOrUpdate(ctx, fromLibrary)
				to.SetLibrary(ctx, toLibrary)
			}
		}
	}
	if !from.Messages.IsNull() && !from.Messages.IsUnknown() && to.Messages.IsNull() && len(from.Messages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Messages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Messages = from.Messages
	}
}

func (to *LibraryFullStatus) SyncFieldsDuringRead(ctx context.Context, from LibraryFullStatus) {
	if !from.Library.IsNull() && !from.Library.IsUnknown() {
		if toLibrary, ok := to.GetLibrary(ctx); ok {
			if fromLibrary, ok := from.GetLibrary(ctx); ok {
				toLibrary.SyncFieldsDuringRead(ctx, fromLibrary)
				to.SetLibrary(ctx, toLibrary)
			}
		}
	}
	if !from.Messages.IsNull() && !from.Messages.IsUnknown() && to.Messages.IsNull() && len(from.Messages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Messages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Messages = from.Messages
	}
}

func (m LibraryFullStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_library_for_all_clusters"] = attrs["is_library_for_all_clusters"].SetOptional()
	attrs["library"] = attrs["library"].SetOptional()
	attrs["messages"] = attrs["messages"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LibraryFullStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LibraryFullStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"library":  reflect.TypeOf(Library{}),
		"messages": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LibraryFullStatus
// only implements ToObjectValue() and Type().
func (m LibraryFullStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_library_for_all_clusters": m.IsLibraryForAllClusters,
			"library":                     m.Library,
			"messages":                    m.Messages,
			"status":                      m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LibraryFullStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_library_for_all_clusters": types.BoolType,
			"library":                     Library{}.Type(ctx),
			"messages": basetypes.ListType{
				ElemType: types.StringType,
			},
			"status": types.StringType,
		},
	}
}

// GetLibrary returns the value of the Library field in LibraryFullStatus as
// a Library value.
// If the field is unknown or null, the boolean return value is false.
func (m *LibraryFullStatus) GetLibrary(ctx context.Context) (Library, bool) {
	var e Library
	if m.Library.IsNull() || m.Library.IsUnknown() {
		return e, false
	}
	var v Library
	d := m.Library.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibrary sets the value of the Library field in LibraryFullStatus.
func (m *LibraryFullStatus) SetLibrary(ctx context.Context, v Library) {
	vs := v.ToObjectValue(ctx)
	m.Library = vs
}

// GetMessages returns the value of the Messages field in LibraryFullStatus as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *LibraryFullStatus) GetMessages(ctx context.Context) ([]types.String, bool) {
	if m.Messages.IsNull() || m.Messages.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Messages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMessages sets the value of the Messages field in LibraryFullStatus.
func (m *LibraryFullStatus) SetMessages(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Messages = types.ListValueMust(t, vs)
}

type ListAllClusterLibraryStatuses struct {
}

func (to *ListAllClusterLibraryStatuses) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllClusterLibraryStatuses) {
}

func (to *ListAllClusterLibraryStatuses) SyncFieldsDuringRead(ctx context.Context, from ListAllClusterLibraryStatuses) {
}

func (m ListAllClusterLibraryStatuses) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllClusterLibraryStatuses.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAllClusterLibraryStatuses) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllClusterLibraryStatuses
// only implements ToObjectValue() and Type().
func (m ListAllClusterLibraryStatuses) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListAllClusterLibraryStatuses) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListAllClusterLibraryStatusesResponse struct {
	// A list of cluster statuses.
	Statuses types.List `tfsdk:"statuses"`
}

func (to *ListAllClusterLibraryStatusesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAllClusterLibraryStatusesResponse) {
	if !from.Statuses.IsNull() && !from.Statuses.IsUnknown() && to.Statuses.IsNull() && len(from.Statuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Statuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Statuses = from.Statuses
	}
}

func (to *ListAllClusterLibraryStatusesResponse) SyncFieldsDuringRead(ctx context.Context, from ListAllClusterLibraryStatusesResponse) {
	if !from.Statuses.IsNull() && !from.Statuses.IsUnknown() && to.Statuses.IsNull() && len(from.Statuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Statuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Statuses = from.Statuses
	}
}

func (m ListAllClusterLibraryStatusesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["statuses"] = attrs["statuses"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllClusterLibraryStatusesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAllClusterLibraryStatusesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statuses": reflect.TypeOf(ClusterLibraryStatuses{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllClusterLibraryStatusesResponse
// only implements ToObjectValue() and Type().
func (m ListAllClusterLibraryStatusesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statuses": m.Statuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAllClusterLibraryStatusesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statuses": basetypes.ListType{
				ElemType: ClusterLibraryStatuses{}.Type(ctx),
			},
		},
	}
}

// GetStatuses returns the value of the Statuses field in ListAllClusterLibraryStatusesResponse as
// a slice of ClusterLibraryStatuses values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListAllClusterLibraryStatusesResponse) GetStatuses(ctx context.Context) ([]ClusterLibraryStatuses, bool) {
	if m.Statuses.IsNull() || m.Statuses.IsUnknown() {
		return nil, false
	}
	var v []ClusterLibraryStatuses
	d := m.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in ListAllClusterLibraryStatusesResponse.
func (m *ListAllClusterLibraryStatusesResponse) SetStatuses(ctx context.Context, v []ClusterLibraryStatuses) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Statuses = types.ListValueMust(t, vs)
}

type ListAvailableZonesResponse struct {
	// The availability zone if no ``zone_id`` is provided in the cluster
	// creation request.
	DefaultZone types.String `tfsdk:"default_zone"`
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	Zones types.List `tfsdk:"zones"`
}

func (to *ListAvailableZonesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAvailableZonesResponse) {
	if !from.Zones.IsNull() && !from.Zones.IsUnknown() && to.Zones.IsNull() && len(from.Zones.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Zones, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Zones = from.Zones
	}
}

func (to *ListAvailableZonesResponse) SyncFieldsDuringRead(ctx context.Context, from ListAvailableZonesResponse) {
	if !from.Zones.IsNull() && !from.Zones.IsUnknown() && to.Zones.IsNull() && len(from.Zones.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Zones, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Zones = from.Zones
	}
}

func (m ListAvailableZonesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_zone"] = attrs["default_zone"].SetOptional()
	attrs["zones"] = attrs["zones"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAvailableZonesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAvailableZonesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"zones": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAvailableZonesResponse
// only implements ToObjectValue() and Type().
func (m ListAvailableZonesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_zone": m.DefaultZone,
			"zones":        m.Zones,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAvailableZonesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_zone": types.StringType,
			"zones": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetZones returns the value of the Zones field in ListAvailableZonesResponse as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListAvailableZonesResponse) GetZones(ctx context.Context) ([]types.String, bool) {
	if m.Zones.IsNull() || m.Zones.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Zones.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetZones sets the value of the Zones field in ListAvailableZonesResponse.
func (m *ListAvailableZonesResponse) SetZones(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["zones"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Zones = types.ListValueMust(t, vs)
}

type ListClusterCompliancesRequest struct {
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token that can be used to navigate to the next page or previous
	// page as returned by `next_page_token` or `prev_page_token`.
	PageToken types.String `tfsdk:"-"`
	// Canonical unique identifier for the cluster policy.
	PolicyId types.String `tfsdk:"-"`
}

func (to *ListClusterCompliancesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListClusterCompliancesRequest) {
}

func (to *ListClusterCompliancesRequest) SyncFieldsDuringRead(ctx context.Context, from ListClusterCompliancesRequest) {
}

func (m ListClusterCompliancesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy_id"] = attrs["policy_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClusterCompliancesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListClusterCompliancesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterCompliancesRequest
// only implements ToObjectValue() and Type().
func (m ListClusterCompliancesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"policy_id":  m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListClusterCompliancesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"policy_id":  types.StringType,
		},
	}
}

type ListClusterCompliancesResponse struct {
	// A list of clusters and their policy compliance statuses.
	Clusters types.List `tfsdk:"clusters"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken types.String `tfsdk:"prev_page_token"`
}

func (to *ListClusterCompliancesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListClusterCompliancesResponse) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
}

func (to *ListClusterCompliancesResponse) SyncFieldsDuringRead(ctx context.Context, from ListClusterCompliancesResponse) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
}

func (m ListClusterCompliancesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clusters"] = attrs["clusters"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["prev_page_token"] = attrs["prev_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClusterCompliancesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListClusterCompliancesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters": reflect.TypeOf(ClusterCompliance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterCompliancesResponse
// only implements ToObjectValue() and Type().
func (m ListClusterCompliancesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusters":        m.Clusters,
			"next_page_token": m.NextPageToken,
			"prev_page_token": m.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListClusterCompliancesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusters": basetypes.ListType{
				ElemType: ClusterCompliance{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in ListClusterCompliancesResponse as
// a slice of ClusterCompliance values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListClusterCompliancesResponse) GetClusters(ctx context.Context) ([]ClusterCompliance, bool) {
	if m.Clusters.IsNull() || m.Clusters.IsUnknown() {
		return nil, false
	}
	var v []ClusterCompliance
	d := m.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in ListClusterCompliancesResponse.
func (m *ListClusterCompliancesResponse) SetClusters(ctx context.Context, v []ClusterCompliance) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Clusters = types.ListValueMust(t, vs)
}

type ListClusterPoliciesRequest struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	SortColumn types.String `tfsdk:"-"`
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	SortOrder types.String `tfsdk:"-"`
}

func (to *ListClusterPoliciesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListClusterPoliciesRequest) {
}

func (to *ListClusterPoliciesRequest) SyncFieldsDuringRead(ctx context.Context, from ListClusterPoliciesRequest) {
}

func (m ListClusterPoliciesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sort_order"] = attrs["sort_order"].SetOptional()
	attrs["sort_column"] = attrs["sort_column"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClusterPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListClusterPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterPoliciesRequest
// only implements ToObjectValue() and Type().
func (m ListClusterPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sort_column": m.SortColumn,
			"sort_order":  m.SortOrder,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListClusterPoliciesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sort_column": types.StringType,
			"sort_order":  types.StringType,
		},
	}
}

type ListClustersFilterBy struct {
	// The source of cluster creation.
	ClusterSources types.List `tfsdk:"cluster_sources"`
	// The current state of the clusters.
	ClusterStates types.List `tfsdk:"cluster_states"`
	// Whether the clusters are pinned or not.
	IsPinned types.Bool `tfsdk:"is_pinned"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
}

func (to *ListClustersFilterBy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListClustersFilterBy) {
	if !from.ClusterSources.IsNull() && !from.ClusterSources.IsUnknown() && to.ClusterSources.IsNull() && len(from.ClusterSources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ClusterSources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ClusterSources = from.ClusterSources
	}
	if !from.ClusterStates.IsNull() && !from.ClusterStates.IsUnknown() && to.ClusterStates.IsNull() && len(from.ClusterStates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ClusterStates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ClusterStates = from.ClusterStates
	}
}

func (to *ListClustersFilterBy) SyncFieldsDuringRead(ctx context.Context, from ListClustersFilterBy) {
	if !from.ClusterSources.IsNull() && !from.ClusterSources.IsUnknown() && to.ClusterSources.IsNull() && len(from.ClusterSources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ClusterSources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ClusterSources = from.ClusterSources
	}
	if !from.ClusterStates.IsNull() && !from.ClusterStates.IsUnknown() && to.ClusterStates.IsNull() && len(from.ClusterStates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ClusterStates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ClusterStates = from.ClusterStates
	}
}

func (m ListClustersFilterBy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_sources"] = attrs["cluster_sources"].SetOptional()
	attrs["cluster_states"] = attrs["cluster_states"].SetOptional()
	attrs["is_pinned"] = attrs["is_pinned"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClustersFilterBy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListClustersFilterBy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_sources": reflect.TypeOf(types.String{}),
		"cluster_states":  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersFilterBy
// only implements ToObjectValue() and Type().
func (m ListClustersFilterBy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_sources": m.ClusterSources,
			"cluster_states":  m.ClusterStates,
			"is_pinned":       m.IsPinned,
			"policy_id":       m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListClustersFilterBy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_sources": basetypes.ListType{
				ElemType: types.StringType,
			},
			"cluster_states": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_pinned": types.BoolType,
			"policy_id": types.StringType,
		},
	}
}

// GetClusterSources returns the value of the ClusterSources field in ListClustersFilterBy as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListClustersFilterBy) GetClusterSources(ctx context.Context) ([]types.String, bool) {
	if m.ClusterSources.IsNull() || m.ClusterSources.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ClusterSources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterSources sets the value of the ClusterSources field in ListClustersFilterBy.
func (m *ListClustersFilterBy) SetClusterSources(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_sources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ClusterSources = types.ListValueMust(t, vs)
}

// GetClusterStates returns the value of the ClusterStates field in ListClustersFilterBy as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListClustersFilterBy) GetClusterStates(ctx context.Context) ([]types.String, bool) {
	if m.ClusterStates.IsNull() || m.ClusterStates.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ClusterStates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterStates sets the value of the ClusterStates field in ListClustersFilterBy.
func (m *ListClustersFilterBy) SetClusterStates(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_states"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ClusterStates = types.ListValueMust(t, vs)
}

type ListClustersRequest struct {
	// Filters to apply to the list of clusters.
	FilterBy types.Object `tfsdk:"-"`
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize types.Int64 `tfsdk:"-"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of clusters respectively.
	PageToken types.String `tfsdk:"-"`
	// Sort the list of clusters by a specific criteria.
	SortBy types.Object `tfsdk:"-"`
}

func (to *ListClustersRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListClustersRequest) {
	if !from.FilterBy.IsNull() && !from.FilterBy.IsUnknown() {
		if toFilterBy, ok := to.GetFilterBy(ctx); ok {
			if fromFilterBy, ok := from.GetFilterBy(ctx); ok {
				// Recursively sync the fields of FilterBy
				toFilterBy.SyncFieldsDuringCreateOrUpdate(ctx, fromFilterBy)
				to.SetFilterBy(ctx, toFilterBy)
			}
		}
	}
	if !from.SortBy.IsNull() && !from.SortBy.IsUnknown() {
		if toSortBy, ok := to.GetSortBy(ctx); ok {
			if fromSortBy, ok := from.GetSortBy(ctx); ok {
				// Recursively sync the fields of SortBy
				toSortBy.SyncFieldsDuringCreateOrUpdate(ctx, fromSortBy)
				to.SetSortBy(ctx, toSortBy)
			}
		}
	}
}

func (to *ListClustersRequest) SyncFieldsDuringRead(ctx context.Context, from ListClustersRequest) {
	if !from.FilterBy.IsNull() && !from.FilterBy.IsUnknown() {
		if toFilterBy, ok := to.GetFilterBy(ctx); ok {
			if fromFilterBy, ok := from.GetFilterBy(ctx); ok {
				toFilterBy.SyncFieldsDuringRead(ctx, fromFilterBy)
				to.SetFilterBy(ctx, toFilterBy)
			}
		}
	}
	if !from.SortBy.IsNull() && !from.SortBy.IsUnknown() {
		if toSortBy, ok := to.GetSortBy(ctx); ok {
			if fromSortBy, ok := from.GetSortBy(ctx); ok {
				toSortBy.SyncFieldsDuringRead(ctx, fromSortBy)
				to.SetSortBy(ctx, toSortBy)
			}
		}
	}
}

func (m ListClustersRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter_by"] = attrs["filter_by"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["sort_by"] = attrs["sort_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClustersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListClustersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter_by": reflect.TypeOf(ListClustersFilterBy{}),
		"sort_by":   reflect.TypeOf(ListClustersSortBy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersRequest
// only implements ToObjectValue() and Type().
func (m ListClustersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_by":  m.FilterBy,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"sort_by":    m.SortBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListClustersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_by":  ListClustersFilterBy{}.Type(ctx),
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"sort_by":    ListClustersSortBy{}.Type(ctx),
		},
	}
}

// GetFilterBy returns the value of the FilterBy field in ListClustersRequest as
// a ListClustersFilterBy value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListClustersRequest) GetFilterBy(ctx context.Context) (ListClustersFilterBy, bool) {
	var e ListClustersFilterBy
	if m.FilterBy.IsNull() || m.FilterBy.IsUnknown() {
		return e, false
	}
	var v ListClustersFilterBy
	d := m.FilterBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilterBy sets the value of the FilterBy field in ListClustersRequest.
func (m *ListClustersRequest) SetFilterBy(ctx context.Context, v ListClustersFilterBy) {
	vs := v.ToObjectValue(ctx)
	m.FilterBy = vs
}

// GetSortBy returns the value of the SortBy field in ListClustersRequest as
// a ListClustersSortBy value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListClustersRequest) GetSortBy(ctx context.Context) (ListClustersSortBy, bool) {
	var e ListClustersSortBy
	if m.SortBy.IsNull() || m.SortBy.IsUnknown() {
		return e, false
	}
	var v ListClustersSortBy
	d := m.SortBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSortBy sets the value of the SortBy field in ListClustersRequest.
func (m *ListClustersRequest) SetSortBy(ctx context.Context, v ListClustersSortBy) {
	vs := v.ToObjectValue(ctx)
	m.SortBy = vs
}

type ListClustersResponse struct {
	Clusters types.List `tfsdk:"clusters"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken types.String `tfsdk:"prev_page_token"`
}

func (to *ListClustersResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListClustersResponse) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
}

func (to *ListClustersResponse) SyncFieldsDuringRead(ctx context.Context, from ListClustersResponse) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
}

func (m ListClustersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clusters"] = attrs["clusters"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["prev_page_token"] = attrs["prev_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClustersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListClustersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters": reflect.TypeOf(ClusterDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersResponse
// only implements ToObjectValue() and Type().
func (m ListClustersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusters":        m.Clusters,
			"next_page_token": m.NextPageToken,
			"prev_page_token": m.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListClustersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusters": basetypes.ListType{
				ElemType: ClusterDetails{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in ListClustersResponse as
// a slice of ClusterDetails values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListClustersResponse) GetClusters(ctx context.Context) ([]ClusterDetails, bool) {
	if m.Clusters.IsNull() || m.Clusters.IsUnknown() {
		return nil, false
	}
	var v []ClusterDetails
	d := m.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in ListClustersResponse.
func (m *ListClustersResponse) SetClusters(ctx context.Context, v []ClusterDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Clusters = types.ListValueMust(t, vs)
}

type ListClustersSortBy struct {
	// The direction to sort by.
	Direction types.String `tfsdk:"direction"`
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	Field types.String `tfsdk:"field"`
}

func (to *ListClustersSortBy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListClustersSortBy) {
}

func (to *ListClustersSortBy) SyncFieldsDuringRead(ctx context.Context, from ListClustersSortBy) {
}

func (m ListClustersSortBy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["direction"] = attrs["direction"].SetOptional()
	attrs["field"] = attrs["field"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClustersSortBy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListClustersSortBy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersSortBy
// only implements ToObjectValue() and Type().
func (m ListClustersSortBy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direction": m.Direction,
			"field":     m.Field,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListClustersSortBy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direction": types.StringType,
			"field":     types.StringType,
		},
	}
}

type ListDefaultBaseEnvironmentsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListDefaultBaseEnvironmentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDefaultBaseEnvironmentsRequest) {
}

func (to *ListDefaultBaseEnvironmentsRequest) SyncFieldsDuringRead(ctx context.Context, from ListDefaultBaseEnvironmentsRequest) {
}

func (m ListDefaultBaseEnvironmentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDefaultBaseEnvironmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDefaultBaseEnvironmentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDefaultBaseEnvironmentsRequest
// only implements ToObjectValue() and Type().
func (m ListDefaultBaseEnvironmentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDefaultBaseEnvironmentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListDefaultBaseEnvironmentsResponse struct {
	DefaultBaseEnvironments types.List `tfsdk:"default_base_environments"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDefaultBaseEnvironmentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDefaultBaseEnvironmentsResponse) {
	if !from.DefaultBaseEnvironments.IsNull() && !from.DefaultBaseEnvironments.IsUnknown() && to.DefaultBaseEnvironments.IsNull() && len(from.DefaultBaseEnvironments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DefaultBaseEnvironments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DefaultBaseEnvironments = from.DefaultBaseEnvironments
	}
}

func (to *ListDefaultBaseEnvironmentsResponse) SyncFieldsDuringRead(ctx context.Context, from ListDefaultBaseEnvironmentsResponse) {
	if !from.DefaultBaseEnvironments.IsNull() && !from.DefaultBaseEnvironments.IsUnknown() && to.DefaultBaseEnvironments.IsNull() && len(from.DefaultBaseEnvironments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DefaultBaseEnvironments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DefaultBaseEnvironments = from.DefaultBaseEnvironments
	}
}

func (m ListDefaultBaseEnvironmentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_base_environments"] = attrs["default_base_environments"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDefaultBaseEnvironmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDefaultBaseEnvironmentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_base_environments": reflect.TypeOf(DefaultBaseEnvironment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDefaultBaseEnvironmentsResponse
// only implements ToObjectValue() and Type().
func (m ListDefaultBaseEnvironmentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_base_environments": m.DefaultBaseEnvironments,
			"next_page_token":           m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDefaultBaseEnvironmentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_base_environments": basetypes.ListType{
				ElemType: DefaultBaseEnvironment{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDefaultBaseEnvironments returns the value of the DefaultBaseEnvironments field in ListDefaultBaseEnvironmentsResponse as
// a slice of DefaultBaseEnvironment values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDefaultBaseEnvironmentsResponse) GetDefaultBaseEnvironments(ctx context.Context) ([]DefaultBaseEnvironment, bool) {
	if m.DefaultBaseEnvironments.IsNull() || m.DefaultBaseEnvironments.IsUnknown() {
		return nil, false
	}
	var v []DefaultBaseEnvironment
	d := m.DefaultBaseEnvironments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultBaseEnvironments sets the value of the DefaultBaseEnvironments field in ListDefaultBaseEnvironmentsResponse.
func (m *ListDefaultBaseEnvironmentsResponse) SetDefaultBaseEnvironments(ctx context.Context, v []DefaultBaseEnvironment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_base_environments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DefaultBaseEnvironments = types.ListValueMust(t, vs)
}

type ListGlobalInitScriptsRequest struct {
}

func (to *ListGlobalInitScriptsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGlobalInitScriptsRequest) {
}

func (to *ListGlobalInitScriptsRequest) SyncFieldsDuringRead(ctx context.Context, from ListGlobalInitScriptsRequest) {
}

func (m ListGlobalInitScriptsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGlobalInitScriptsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGlobalInitScriptsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGlobalInitScriptsRequest
// only implements ToObjectValue() and Type().
func (m ListGlobalInitScriptsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListGlobalInitScriptsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListGlobalInitScriptsResponse struct {
	Scripts types.List `tfsdk:"scripts"`
}

func (to *ListGlobalInitScriptsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGlobalInitScriptsResponse) {
	if !from.Scripts.IsNull() && !from.Scripts.IsUnknown() && to.Scripts.IsNull() && len(from.Scripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scripts = from.Scripts
	}
}

func (to *ListGlobalInitScriptsResponse) SyncFieldsDuringRead(ctx context.Context, from ListGlobalInitScriptsResponse) {
	if !from.Scripts.IsNull() && !from.Scripts.IsUnknown() && to.Scripts.IsNull() && len(from.Scripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scripts = from.Scripts
	}
}

func (m ListGlobalInitScriptsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["scripts"] = attrs["scripts"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGlobalInitScriptsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGlobalInitScriptsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scripts": reflect.TypeOf(GlobalInitScriptDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGlobalInitScriptsResponse
// only implements ToObjectValue() and Type().
func (m ListGlobalInitScriptsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scripts": m.Scripts,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGlobalInitScriptsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scripts": basetypes.ListType{
				ElemType: GlobalInitScriptDetails{}.Type(ctx),
			},
		},
	}
}

// GetScripts returns the value of the Scripts field in ListGlobalInitScriptsResponse as
// a slice of GlobalInitScriptDetails values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListGlobalInitScriptsResponse) GetScripts(ctx context.Context) ([]GlobalInitScriptDetails, bool) {
	if m.Scripts.IsNull() || m.Scripts.IsUnknown() {
		return nil, false
	}
	var v []GlobalInitScriptDetails
	d := m.Scripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScripts sets the value of the Scripts field in ListGlobalInitScriptsResponse.
func (m *ListGlobalInitScriptsResponse) SetScripts(ctx context.Context, v []GlobalInitScriptDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scripts = types.ListValueMust(t, vs)
}

type ListInstancePools struct {
	InstancePools types.List `tfsdk:"instance_pools"`
}

func (to *ListInstancePools) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListInstancePools) {
	if !from.InstancePools.IsNull() && !from.InstancePools.IsUnknown() && to.InstancePools.IsNull() && len(from.InstancePools.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InstancePools, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InstancePools = from.InstancePools
	}
}

func (to *ListInstancePools) SyncFieldsDuringRead(ctx context.Context, from ListInstancePools) {
	if !from.InstancePools.IsNull() && !from.InstancePools.IsUnknown() && to.InstancePools.IsNull() && len(from.InstancePools.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InstancePools, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InstancePools = from.InstancePools
	}
}

func (m ListInstancePools) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_pools"] = attrs["instance_pools"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstancePools.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListInstancePools) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"instance_pools": reflect.TypeOf(InstancePoolAndStats{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstancePools
// only implements ToObjectValue() and Type().
func (m ListInstancePools) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pools": m.InstancePools,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListInstancePools) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pools": basetypes.ListType{
				ElemType: InstancePoolAndStats{}.Type(ctx),
			},
		},
	}
}

// GetInstancePools returns the value of the InstancePools field in ListInstancePools as
// a slice of InstancePoolAndStats values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListInstancePools) GetInstancePools(ctx context.Context) ([]InstancePoolAndStats, bool) {
	if m.InstancePools.IsNull() || m.InstancePools.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAndStats
	d := m.InstancePools.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstancePools sets the value of the InstancePools field in ListInstancePools.
func (m *ListInstancePools) SetInstancePools(ctx context.Context, v []InstancePoolAndStats) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_pools"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InstancePools = types.ListValueMust(t, vs)
}

type ListInstancePoolsRequest struct {
}

func (to *ListInstancePoolsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListInstancePoolsRequest) {
}

func (to *ListInstancePoolsRequest) SyncFieldsDuringRead(ctx context.Context, from ListInstancePoolsRequest) {
}

func (m ListInstancePoolsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstancePoolsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListInstancePoolsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstancePoolsRequest
// only implements ToObjectValue() and Type().
func (m ListInstancePoolsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListInstancePoolsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListInstanceProfilesRequest struct {
}

func (to *ListInstanceProfilesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListInstanceProfilesRequest) {
}

func (to *ListInstanceProfilesRequest) SyncFieldsDuringRead(ctx context.Context, from ListInstanceProfilesRequest) {
}

func (m ListInstanceProfilesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstanceProfilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListInstanceProfilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstanceProfilesRequest
// only implements ToObjectValue() and Type().
func (m ListInstanceProfilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListInstanceProfilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListInstanceProfilesResponse struct {
	// A list of instance profiles that the user can access.
	InstanceProfiles types.List `tfsdk:"instance_profiles"`
}

func (to *ListInstanceProfilesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListInstanceProfilesResponse) {
	if !from.InstanceProfiles.IsNull() && !from.InstanceProfiles.IsUnknown() && to.InstanceProfiles.IsNull() && len(from.InstanceProfiles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InstanceProfiles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InstanceProfiles = from.InstanceProfiles
	}
}

func (to *ListInstanceProfilesResponse) SyncFieldsDuringRead(ctx context.Context, from ListInstanceProfilesResponse) {
	if !from.InstanceProfiles.IsNull() && !from.InstanceProfiles.IsUnknown() && to.InstanceProfiles.IsNull() && len(from.InstanceProfiles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InstanceProfiles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InstanceProfiles = from.InstanceProfiles
	}
}

func (m ListInstanceProfilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_profiles"] = attrs["instance_profiles"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstanceProfilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListInstanceProfilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"instance_profiles": reflect.TypeOf(InstanceProfile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstanceProfilesResponse
// only implements ToObjectValue() and Type().
func (m ListInstanceProfilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_profiles": m.InstanceProfiles,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListInstanceProfilesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_profiles": basetypes.ListType{
				ElemType: InstanceProfile{}.Type(ctx),
			},
		},
	}
}

// GetInstanceProfiles returns the value of the InstanceProfiles field in ListInstanceProfilesResponse as
// a slice of InstanceProfile values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListInstanceProfilesResponse) GetInstanceProfiles(ctx context.Context) ([]InstanceProfile, bool) {
	if m.InstanceProfiles.IsNull() || m.InstanceProfiles.IsUnknown() {
		return nil, false
	}
	var v []InstanceProfile
	d := m.InstanceProfiles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstanceProfiles sets the value of the InstanceProfiles field in ListInstanceProfilesResponse.
func (m *ListInstanceProfilesResponse) SetInstanceProfiles(ctx context.Context, v []InstanceProfile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_profiles"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InstanceProfiles = types.ListValueMust(t, vs)
}

type ListNodeTypesRequest struct {
}

func (to *ListNodeTypesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNodeTypesRequest) {
}

func (to *ListNodeTypesRequest) SyncFieldsDuringRead(ctx context.Context, from ListNodeTypesRequest) {
}

func (m ListNodeTypesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNodeTypesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNodeTypesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNodeTypesRequest
// only implements ToObjectValue() and Type().
func (m ListNodeTypesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListNodeTypesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNodeTypesResponse struct {
	// The list of available Spark node types.
	NodeTypes types.List `tfsdk:"node_types"`
}

func (to *ListNodeTypesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNodeTypesResponse) {
	if !from.NodeTypes.IsNull() && !from.NodeTypes.IsUnknown() && to.NodeTypes.IsNull() && len(from.NodeTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for NodeTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.NodeTypes = from.NodeTypes
	}
}

func (to *ListNodeTypesResponse) SyncFieldsDuringRead(ctx context.Context, from ListNodeTypesResponse) {
	if !from.NodeTypes.IsNull() && !from.NodeTypes.IsUnknown() && to.NodeTypes.IsNull() && len(from.NodeTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for NodeTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.NodeTypes = from.NodeTypes
	}
}

func (m ListNodeTypesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["node_types"] = attrs["node_types"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNodeTypesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNodeTypesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_types": reflect.TypeOf(NodeType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNodeTypesResponse
// only implements ToObjectValue() and Type().
func (m ListNodeTypesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"node_types": m.NodeTypes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNodeTypesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"node_types": basetypes.ListType{
				ElemType: NodeType{}.Type(ctx),
			},
		},
	}
}

// GetNodeTypes returns the value of the NodeTypes field in ListNodeTypesResponse as
// a slice of NodeType values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListNodeTypesResponse) GetNodeTypes(ctx context.Context) ([]NodeType, bool) {
	if m.NodeTypes.IsNull() || m.NodeTypes.IsUnknown() {
		return nil, false
	}
	var v []NodeType
	d := m.NodeTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeTypes sets the value of the NodeTypes field in ListNodeTypesResponse.
func (m *ListNodeTypesResponse) SetNodeTypes(ctx context.Context, v []NodeType) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["node_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.NodeTypes = types.ListValueMust(t, vs)
}

type ListPoliciesResponse struct {
	// List of policies.
	Policies types.List `tfsdk:"policies"`
}

func (to *ListPoliciesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPoliciesResponse) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (to *ListPoliciesResponse) SyncFieldsDuringRead(ctx context.Context, from ListPoliciesResponse) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (m ListPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policies"] = attrs["policies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPoliciesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(Policy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPoliciesResponse
// only implements ToObjectValue() and Type().
func (m ListPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policies": m.Policies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPoliciesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policies": basetypes.ListType{
				ElemType: Policy{}.Type(ctx),
			},
		},
	}
}

// GetPolicies returns the value of the Policies field in ListPoliciesResponse as
// a slice of Policy values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListPoliciesResponse) GetPolicies(ctx context.Context) ([]Policy, bool) {
	if m.Policies.IsNull() || m.Policies.IsUnknown() {
		return nil, false
	}
	var v []Policy
	d := m.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListPoliciesResponse.
func (m *ListPoliciesResponse) SetPolicies(ctx context.Context, v []Policy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Policies = types.ListValueMust(t, vs)
}

type ListPolicyFamiliesRequest struct {
	// Maximum number of policy families to return.
	MaxResults types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListPolicyFamiliesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPolicyFamiliesRequest) {
}

func (to *ListPolicyFamiliesRequest) SyncFieldsDuringRead(ctx context.Context, from ListPolicyFamiliesRequest) {
}

func (m ListPolicyFamiliesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPolicyFamiliesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPolicyFamiliesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPolicyFamiliesRequest
// only implements ToObjectValue() and Type().
func (m ListPolicyFamiliesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPolicyFamiliesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListPolicyFamiliesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of policy families.
	PolicyFamilies types.List `tfsdk:"policy_families"`
}

func (to *ListPolicyFamiliesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPolicyFamiliesResponse) {
	if !from.PolicyFamilies.IsNull() && !from.PolicyFamilies.IsUnknown() && to.PolicyFamilies.IsNull() && len(from.PolicyFamilies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PolicyFamilies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PolicyFamilies = from.PolicyFamilies
	}
}

func (to *ListPolicyFamiliesResponse) SyncFieldsDuringRead(ctx context.Context, from ListPolicyFamiliesResponse) {
	if !from.PolicyFamilies.IsNull() && !from.PolicyFamilies.IsUnknown() && to.PolicyFamilies.IsNull() && len(from.PolicyFamilies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PolicyFamilies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PolicyFamilies = from.PolicyFamilies
	}
}

func (m ListPolicyFamiliesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["policy_families"] = attrs["policy_families"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPolicyFamiliesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPolicyFamiliesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy_families": reflect.TypeOf(PolicyFamily{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPolicyFamiliesResponse
// only implements ToObjectValue() and Type().
func (m ListPolicyFamiliesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"policy_families": m.PolicyFamilies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPolicyFamiliesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"policy_families": basetypes.ListType{
				ElemType: PolicyFamily{}.Type(ctx),
			},
		},
	}
}

// GetPolicyFamilies returns the value of the PolicyFamilies field in ListPolicyFamiliesResponse as
// a slice of PolicyFamily values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListPolicyFamiliesResponse) GetPolicyFamilies(ctx context.Context) ([]PolicyFamily, bool) {
	if m.PolicyFamilies.IsNull() || m.PolicyFamilies.IsUnknown() {
		return nil, false
	}
	var v []PolicyFamily
	d := m.PolicyFamilies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicyFamilies sets the value of the PolicyFamilies field in ListPolicyFamiliesResponse.
func (m *ListPolicyFamiliesResponse) SetPolicyFamilies(ctx context.Context, v []PolicyFamily) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy_families"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PolicyFamilies = types.ListValueMust(t, vs)
}

type ListZonesRequest struct {
}

func (to *ListZonesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListZonesRequest) {
}

func (to *ListZonesRequest) SyncFieldsDuringRead(ctx context.Context, from ListZonesRequest) {
}

func (m ListZonesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListZonesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListZonesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListZonesRequest
// only implements ToObjectValue() and Type().
func (m ListZonesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListZonesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LocalFileInfo struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	Destination types.String `tfsdk:"destination"`
}

func (to *LocalFileInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LocalFileInfo) {
}

func (to *LocalFileInfo) SyncFieldsDuringRead(ctx context.Context, from LocalFileInfo) {
}

func (m LocalFileInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination"] = attrs["destination"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LocalFileInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LocalFileInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LocalFileInfo
// only implements ToObjectValue() and Type().
func (m LocalFileInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": m.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LocalFileInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type LogAnalyticsInfo struct {
	LogAnalyticsPrimaryKey types.String `tfsdk:"log_analytics_primary_key"`

	LogAnalyticsWorkspaceId types.String `tfsdk:"log_analytics_workspace_id"`
}

func (to *LogAnalyticsInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogAnalyticsInfo) {
}

func (to *LogAnalyticsInfo) SyncFieldsDuringRead(ctx context.Context, from LogAnalyticsInfo) {
}

func (m LogAnalyticsInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_analytics_primary_key"] = attrs["log_analytics_primary_key"].SetOptional()
	attrs["log_analytics_workspace_id"] = attrs["log_analytics_workspace_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogAnalyticsInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogAnalyticsInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogAnalyticsInfo
// only implements ToObjectValue() and Type().
func (m LogAnalyticsInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_analytics_primary_key":  m.LogAnalyticsPrimaryKey,
			"log_analytics_workspace_id": m.LogAnalyticsWorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogAnalyticsInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_analytics_primary_key":  types.StringType,
			"log_analytics_workspace_id": types.StringType,
		},
	}
}

// The log delivery status
type LogSyncStatus struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	LastAttempted types.Int64 `tfsdk:"last_attempted"`
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	LastException types.String `tfsdk:"last_exception"`
}

func (to *LogSyncStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogSyncStatus) {
}

func (to *LogSyncStatus) SyncFieldsDuringRead(ctx context.Context, from LogSyncStatus) {
}

func (m LogSyncStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_attempted"] = attrs["last_attempted"].SetOptional()
	attrs["last_exception"] = attrs["last_exception"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogSyncStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogSyncStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogSyncStatus
// only implements ToObjectValue() and Type().
func (m LogSyncStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_attempted": m.LastAttempted,
			"last_exception": m.LastException,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogSyncStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_attempted": types.Int64Type,
			"last_exception": types.StringType,
		},
	}
}

// Materialized Environment information enables environment sharing and reuse
// via Environment Caching during library installations. Currently this feature
// is only supported for Python libraries.
//
// - If the env cache entry in LMv2 DB doesn't exist or invalid, library
// installations and environment materialization will occur. A new Materialized
// Environment metadata will be sent from DP upon successful library
// installations and env materialization, and is persisted into database by
// LMv2. - If the env cache entry in LMv2 DB is valid, the Materialized
// Environment will be sent to DP by LMv2, and DP will restore the cached
// environment from a store instead of reinstalling libraries from scratch.
//
// If changed, also update
// estore/namespaces/defaultbaseenvironments/latest.proto with new version
type MaterializedEnvironment struct {
	// The timestamp (in epoch milliseconds) when the materialized env is
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
}

func (to *MaterializedEnvironment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MaterializedEnvironment) {
}

func (to *MaterializedEnvironment) SyncFieldsDuringRead(ctx context.Context, from MaterializedEnvironment) {
}

func (m MaterializedEnvironment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MaterializedEnvironment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MaterializedEnvironment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MaterializedEnvironment
// only implements ToObjectValue() and Type().
func (m MaterializedEnvironment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_updated_timestamp": m.LastUpdatedTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MaterializedEnvironment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_updated_timestamp": types.Int64Type,
		},
	}
}

type MavenLibrary struct {
	// Gradle-style maven coordinates. For example: "org.jsoup:jsoup:1.7.2".
	Coordinates types.String `tfsdk:"coordinates"`
	// List of dependences to exclude. For example: `["slf4j:slf4j",
	// "*:hadoop-client"]`.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	Exclusions types.List `tfsdk:"exclusions"`
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	Repo types.String `tfsdk:"repo"`
}

func (to *MavenLibrary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MavenLibrary) {
	if !from.Exclusions.IsNull() && !from.Exclusions.IsUnknown() && to.Exclusions.IsNull() && len(from.Exclusions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exclusions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exclusions = from.Exclusions
	}
}

func (to *MavenLibrary) SyncFieldsDuringRead(ctx context.Context, from MavenLibrary) {
	if !from.Exclusions.IsNull() && !from.Exclusions.IsUnknown() && to.Exclusions.IsNull() && len(from.Exclusions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exclusions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exclusions = from.Exclusions
	}
}

func (m MavenLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["coordinates"] = attrs["coordinates"].SetRequired()
	attrs["exclusions"] = attrs["exclusions"].SetOptional()
	attrs["repo"] = attrs["repo"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MavenLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MavenLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclusions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MavenLibrary
// only implements ToObjectValue() and Type().
func (m MavenLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"coordinates": m.Coordinates,
			"exclusions":  m.Exclusions,
			"repo":        m.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MavenLibrary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"coordinates": types.StringType,
			"exclusions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"repo": types.StringType,
		},
	}
}

// GetExclusions returns the value of the Exclusions field in MavenLibrary as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *MavenLibrary) GetExclusions(ctx context.Context) ([]types.String, bool) {
	if m.Exclusions.IsNull() || m.Exclusions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Exclusions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExclusions sets the value of the Exclusions field in MavenLibrary.
func (m *MavenLibrary) SetExclusions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exclusions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Exclusions = types.ListValueMust(t, vs)
}

// This structure embodies the machine type that hosts spark containers Note:
// this should be an internal data structure for now It is defined in proto in
// case we want to send it over the wire in the future (which is likely)
type NodeInstanceType struct {
	// Unique identifier across instance types
	InstanceTypeId types.String `tfsdk:"instance_type_id"`
	// Size of the individual local disks attached to this instance (i.e. per
	// local disk).
	LocalDiskSizeGb types.Int64 `tfsdk:"local_disk_size_gb"`
	// Number of local disks that are present on this instance.
	LocalDisks types.Int64 `tfsdk:"local_disks"`
	// Size of the individual local nvme disks attached to this instance (i.e.
	// per local disk).
	LocalNvmeDiskSizeGb types.Int64 `tfsdk:"local_nvme_disk_size_gb"`
	// Number of local nvme disks that are present on this instance.
	LocalNvmeDisks types.Int64 `tfsdk:"local_nvme_disks"`
}

func (to *NodeInstanceType) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NodeInstanceType) {
}

func (to *NodeInstanceType) SyncFieldsDuringRead(ctx context.Context, from NodeInstanceType) {
}

func (m NodeInstanceType) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_type_id"] = attrs["instance_type_id"].SetRequired()
	attrs["local_disk_size_gb"] = attrs["local_disk_size_gb"].SetOptional()
	attrs["local_disks"] = attrs["local_disks"].SetOptional()
	attrs["local_nvme_disk_size_gb"] = attrs["local_nvme_disk_size_gb"].SetOptional()
	attrs["local_nvme_disks"] = attrs["local_nvme_disks"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NodeInstanceType.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NodeInstanceType) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NodeInstanceType
// only implements ToObjectValue() and Type().
func (m NodeInstanceType) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_type_id":        m.InstanceTypeId,
			"local_disk_size_gb":      m.LocalDiskSizeGb,
			"local_disks":             m.LocalDisks,
			"local_nvme_disk_size_gb": m.LocalNvmeDiskSizeGb,
			"local_nvme_disks":        m.LocalNvmeDisks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NodeInstanceType) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_type_id":        types.StringType,
			"local_disk_size_gb":      types.Int64Type,
			"local_disks":             types.Int64Type,
			"local_nvme_disk_size_gb": types.Int64Type,
			"local_nvme_disks":        types.Int64Type,
		},
	}
}

// A description of a Spark node type including both the dimensions of the node
// and the instance type on which it will be hosted.
type NodeType struct {
	// A descriptive category for this node type. Examples include "Memory
	// Optimized" and "Compute Optimized".
	Category types.String `tfsdk:"category"`
	// A string description associated with this node type, e.g., "r3.xlarge".
	Description types.String `tfsdk:"description"`
	// An optional hint at the display order of node types in the UI. Within a
	// node type category, lowest numbers come first.
	DisplayOrder types.Int64 `tfsdk:"display_order"`
	// An identifier for the type of hardware that this node runs on, e.g.,
	// "r3.2xlarge" in AWS.
	InstanceTypeId types.String `tfsdk:"instance_type_id"`
	// Whether the node type is deprecated. Non-deprecated node types offer
	// greater performance.
	IsDeprecated types.Bool `tfsdk:"is_deprecated"`
	// AWS specific, whether this instance supports encryption in transit, used
	// for hipaa and pci workloads.
	IsEncryptedInTransit types.Bool `tfsdk:"is_encrypted_in_transit"`
	// Whether this is an Arm-based instance.
	IsGraviton types.Bool `tfsdk:"is_graviton"`
	// Whether this node is hidden from presentation in the UI.
	IsHidden types.Bool `tfsdk:"is_hidden"`
	// Whether this node comes with IO cache enabled by default.
	IsIoCacheEnabled types.Bool `tfsdk:"is_io_cache_enabled"`
	// Memory (in MB) available for this node type.
	MemoryMb types.Int64 `tfsdk:"memory_mb"`
	// A collection of node type info reported by the cloud provider
	NodeInfo types.Object `tfsdk:"node_info"`
	// The NodeInstanceType object corresponding to instance_type_id
	NodeInstanceType types.Object `tfsdk:"node_instance_type"`
	// Unique identifier for this node type.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Number of CPU cores available for this node type. Note that this can be
	// fractional, e.g., 2.5 cores, if the the number of cores on a machine
	// instance is not divisible by the number of Spark nodes on that machine.
	NumCores types.Float64 `tfsdk:"num_cores"`
	// Number of GPUs available for this node type.
	NumGpus types.Int64 `tfsdk:"num_gpus"`

	PhotonDriverCapable types.Bool `tfsdk:"photon_driver_capable"`

	PhotonWorkerCapable types.Bool `tfsdk:"photon_worker_capable"`
	// Whether this node type support cluster tags.
	SupportClusterTags types.Bool `tfsdk:"support_cluster_tags"`
	// Whether this node type support EBS volumes. EBS volumes is disabled for
	// node types that we could place multiple corresponding containers on the
	// same hosting instance.
	SupportEbsVolumes types.Bool `tfsdk:"support_ebs_volumes"`
	// Whether this node type supports port forwarding.
	SupportPortForwarding types.Bool `tfsdk:"support_port_forwarding"`
}

func (to *NodeType) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NodeType) {
	if !from.NodeInfo.IsNull() && !from.NodeInfo.IsUnknown() {
		if toNodeInfo, ok := to.GetNodeInfo(ctx); ok {
			if fromNodeInfo, ok := from.GetNodeInfo(ctx); ok {
				// Recursively sync the fields of NodeInfo
				toNodeInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromNodeInfo)
				to.SetNodeInfo(ctx, toNodeInfo)
			}
		}
	}
	if !from.NodeInstanceType.IsNull() && !from.NodeInstanceType.IsUnknown() {
		if toNodeInstanceType, ok := to.GetNodeInstanceType(ctx); ok {
			if fromNodeInstanceType, ok := from.GetNodeInstanceType(ctx); ok {
				// Recursively sync the fields of NodeInstanceType
				toNodeInstanceType.SyncFieldsDuringCreateOrUpdate(ctx, fromNodeInstanceType)
				to.SetNodeInstanceType(ctx, toNodeInstanceType)
			}
		}
	}
}

func (to *NodeType) SyncFieldsDuringRead(ctx context.Context, from NodeType) {
	if !from.NodeInfo.IsNull() && !from.NodeInfo.IsUnknown() {
		if toNodeInfo, ok := to.GetNodeInfo(ctx); ok {
			if fromNodeInfo, ok := from.GetNodeInfo(ctx); ok {
				toNodeInfo.SyncFieldsDuringRead(ctx, fromNodeInfo)
				to.SetNodeInfo(ctx, toNodeInfo)
			}
		}
	}
	if !from.NodeInstanceType.IsNull() && !from.NodeInstanceType.IsUnknown() {
		if toNodeInstanceType, ok := to.GetNodeInstanceType(ctx); ok {
			if fromNodeInstanceType, ok := from.GetNodeInstanceType(ctx); ok {
				toNodeInstanceType.SyncFieldsDuringRead(ctx, fromNodeInstanceType)
				to.SetNodeInstanceType(ctx, toNodeInstanceType)
			}
		}
	}
}

func (m NodeType) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["category"] = attrs["category"].SetRequired()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["display_order"] = attrs["display_order"].SetOptional()
	attrs["instance_type_id"] = attrs["instance_type_id"].SetRequired()
	attrs["is_deprecated"] = attrs["is_deprecated"].SetOptional()
	attrs["is_encrypted_in_transit"] = attrs["is_encrypted_in_transit"].SetOptional()
	attrs["is_graviton"] = attrs["is_graviton"].SetOptional()
	attrs["is_hidden"] = attrs["is_hidden"].SetOptional()
	attrs["is_io_cache_enabled"] = attrs["is_io_cache_enabled"].SetOptional()
	attrs["memory_mb"] = attrs["memory_mb"].SetRequired()
	attrs["node_info"] = attrs["node_info"].SetOptional()
	attrs["node_instance_type"] = attrs["node_instance_type"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetRequired()
	attrs["num_cores"] = attrs["num_cores"].SetRequired()
	attrs["num_gpus"] = attrs["num_gpus"].SetOptional()
	attrs["photon_driver_capable"] = attrs["photon_driver_capable"].SetOptional()
	attrs["photon_worker_capable"] = attrs["photon_worker_capable"].SetOptional()
	attrs["support_cluster_tags"] = attrs["support_cluster_tags"].SetOptional()
	attrs["support_ebs_volumes"] = attrs["support_ebs_volumes"].SetOptional()
	attrs["support_port_forwarding"] = attrs["support_port_forwarding"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NodeType.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NodeType) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_info":          reflect.TypeOf(CloudProviderNodeInfo{}),
		"node_instance_type": reflect.TypeOf(NodeInstanceType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NodeType
// only implements ToObjectValue() and Type().
func (m NodeType) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"category":                m.Category,
			"description":             m.Description,
			"display_order":           m.DisplayOrder,
			"instance_type_id":        m.InstanceTypeId,
			"is_deprecated":           m.IsDeprecated,
			"is_encrypted_in_transit": m.IsEncryptedInTransit,
			"is_graviton":             m.IsGraviton,
			"is_hidden":               m.IsHidden,
			"is_io_cache_enabled":     m.IsIoCacheEnabled,
			"memory_mb":               m.MemoryMb,
			"node_info":               m.NodeInfo,
			"node_instance_type":      m.NodeInstanceType,
			"node_type_id":            m.NodeTypeId,
			"num_cores":               m.NumCores,
			"num_gpus":                m.NumGpus,
			"photon_driver_capable":   m.PhotonDriverCapable,
			"photon_worker_capable":   m.PhotonWorkerCapable,
			"support_cluster_tags":    m.SupportClusterTags,
			"support_ebs_volumes":     m.SupportEbsVolumes,
			"support_port_forwarding": m.SupportPortForwarding,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NodeType) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"category":                types.StringType,
			"description":             types.StringType,
			"display_order":           types.Int64Type,
			"instance_type_id":        types.StringType,
			"is_deprecated":           types.BoolType,
			"is_encrypted_in_transit": types.BoolType,
			"is_graviton":             types.BoolType,
			"is_hidden":               types.BoolType,
			"is_io_cache_enabled":     types.BoolType,
			"memory_mb":               types.Int64Type,
			"node_info":               CloudProviderNodeInfo{}.Type(ctx),
			"node_instance_type":      NodeInstanceType{}.Type(ctx),
			"node_type_id":            types.StringType,
			"num_cores":               types.Float64Type,
			"num_gpus":                types.Int64Type,
			"photon_driver_capable":   types.BoolType,
			"photon_worker_capable":   types.BoolType,
			"support_cluster_tags":    types.BoolType,
			"support_ebs_volumes":     types.BoolType,
			"support_port_forwarding": types.BoolType,
		},
	}
}

// GetNodeInfo returns the value of the NodeInfo field in NodeType as
// a CloudProviderNodeInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *NodeType) GetNodeInfo(ctx context.Context) (CloudProviderNodeInfo, bool) {
	var e CloudProviderNodeInfo
	if m.NodeInfo.IsNull() || m.NodeInfo.IsUnknown() {
		return e, false
	}
	var v CloudProviderNodeInfo
	d := m.NodeInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeInfo sets the value of the NodeInfo field in NodeType.
func (m *NodeType) SetNodeInfo(ctx context.Context, v CloudProviderNodeInfo) {
	vs := v.ToObjectValue(ctx)
	m.NodeInfo = vs
}

// GetNodeInstanceType returns the value of the NodeInstanceType field in NodeType as
// a NodeInstanceType value.
// If the field is unknown or null, the boolean return value is false.
func (m *NodeType) GetNodeInstanceType(ctx context.Context) (NodeInstanceType, bool) {
	var e NodeInstanceType
	if m.NodeInstanceType.IsNull() || m.NodeInstanceType.IsUnknown() {
		return e, false
	}
	var v NodeInstanceType
	d := m.NodeInstanceType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeInstanceType sets the value of the NodeInstanceType field in NodeType.
func (m *NodeType) SetNodeInstanceType(ctx context.Context, v NodeInstanceType) {
	vs := v.ToObjectValue(ctx)
	m.NodeInstanceType = vs
}

// For Fleet-V2 using classic clusters, this object contains the information
// about the alternate node type ids to use when attempting to launch a cluster.
// It can be used with both the driver and worker node types.
type NodeTypeFlexibility struct {
}

func (to *NodeTypeFlexibility) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NodeTypeFlexibility) {
}

func (to *NodeTypeFlexibility) SyncFieldsDuringRead(ctx context.Context, from NodeTypeFlexibility) {
}

func (m NodeTypeFlexibility) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NodeTypeFlexibility.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NodeTypeFlexibility) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NodeTypeFlexibility
// only implements ToObjectValue() and Type().
func (m NodeTypeFlexibility) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m NodeTypeFlexibility) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Error message of a failed pending instances
type PendingInstanceError struct {
	InstanceId types.String `tfsdk:"instance_id"`

	Message types.String `tfsdk:"message"`
}

func (to *PendingInstanceError) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PendingInstanceError) {
}

func (to *PendingInstanceError) SyncFieldsDuringRead(ctx context.Context, from PendingInstanceError) {
}

func (m PendingInstanceError) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_id"] = attrs["instance_id"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PendingInstanceError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PendingInstanceError) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PendingInstanceError
// only implements ToObjectValue() and Type().
func (m PendingInstanceError) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_id": m.InstanceId,
			"message":     m.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PendingInstanceError) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_id": types.StringType,
			"message":     types.StringType,
		},
	}
}

type PermanentDeleteCluster struct {
	// The cluster to be deleted.
	ClusterId types.String `tfsdk:"cluster_id"`
}

func (to *PermanentDeleteCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PermanentDeleteCluster) {
}

func (to *PermanentDeleteCluster) SyncFieldsDuringRead(ctx context.Context, from PermanentDeleteCluster) {
}

func (m PermanentDeleteCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermanentDeleteCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PermanentDeleteCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermanentDeleteCluster
// only implements ToObjectValue() and Type().
func (m PermanentDeleteCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PermanentDeleteCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type PermanentDeleteClusterResponse struct {
}

func (to *PermanentDeleteClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PermanentDeleteClusterResponse) {
}

func (to *PermanentDeleteClusterResponse) SyncFieldsDuringRead(ctx context.Context, from PermanentDeleteClusterResponse) {
}

func (m PermanentDeleteClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermanentDeleteClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PermanentDeleteClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermanentDeleteClusterResponse
// only implements ToObjectValue() and Type().
func (m PermanentDeleteClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m PermanentDeleteClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type PinCluster struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

func (to *PinCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PinCluster) {
}

func (to *PinCluster) SyncFieldsDuringRead(ctx context.Context, from PinCluster) {
}

func (m PinCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PinCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PinCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PinCluster
// only implements ToObjectValue() and Type().
func (m PinCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PinCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type PinClusterResponse struct {
}

func (to *PinClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PinClusterResponse) {
}

func (to *PinClusterResponse) SyncFieldsDuringRead(ctx context.Context, from PinClusterResponse) {
}

func (m PinClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PinClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PinClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PinClusterResponse
// only implements ToObjectValue() and Type().
func (m PinClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m PinClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Describes a Cluster Policy entity.
type Policy struct {
	// Creation time. The timestamp (in millisecond) when this Cluster Policy
	// was created.
	CreatedAtTimestamp types.Int64 `tfsdk:"created_at_timestamp"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description"`
	// If true, policy is a default policy created and managed by Databricks.
	// Default policies cannot be deleted, and their policy families cannot be
	// changed.
	IsDefault types.Bool `tfsdk:"is_default"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries types.List `tfsdk:"libraries"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId types.String `tfsdk:"policy_family_id"`
	// Canonical unique identifier for the Cluster Policy.
	PolicyId types.String `tfsdk:"policy_id"`
}

func (to *Policy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Policy) {
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
}

func (to *Policy) SyncFieldsDuringRead(ctx context.Context, from Policy) {
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
}

func (m Policy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at_timestamp"] = attrs["created_at_timestamp"].SetOptional()
	attrs["creator_user_name"] = attrs["creator_user_name"].SetOptional()
	attrs["definition"] = attrs["definition"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["is_default"] = attrs["is_default"].SetOptional()
	attrs["libraries"] = attrs["libraries"].SetOptional()
	attrs["max_clusters_per_user"] = attrs["max_clusters_per_user"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["policy_family_definition_overrides"] = attrs["policy_family_definition_overrides"].SetOptional()
	attrs["policy_family_id"] = attrs["policy_family_id"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Policy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Policy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Policy
// only implements ToObjectValue() and Type().
func (m Policy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at_timestamp":               m.CreatedAtTimestamp,
			"creator_user_name":                  m.CreatorUserName,
			"definition":                         m.Definition,
			"description":                        m.Description,
			"is_default":                         m.IsDefault,
			"libraries":                          m.Libraries,
			"max_clusters_per_user":              m.MaxClustersPerUser,
			"name":                               m.Name,
			"policy_family_definition_overrides": m.PolicyFamilyDefinitionOverrides,
			"policy_family_id":                   m.PolicyFamilyId,
			"policy_id":                          m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Policy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at_timestamp": types.Int64Type,
			"creator_user_name":    types.StringType,
			"definition":           types.StringType,
			"description":          types.StringType,
			"is_default":           types.BoolType,
			"libraries": basetypes.ListType{
				ElemType: Library{}.Type(ctx),
			},
			"max_clusters_per_user":              types.Int64Type,
			"name":                               types.StringType,
			"policy_family_definition_overrides": types.StringType,
			"policy_family_id":                   types.StringType,
			"policy_id":                          types.StringType,
		},
	}
}

// GetLibraries returns the value of the Libraries field in Policy as
// a slice of Library values.
// If the field is unknown or null, the boolean return value is false.
func (m *Policy) GetLibraries(ctx context.Context) ([]Library, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in Policy.
func (m *Policy) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

type PolicyFamily struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition"`
	// Human-readable description of the purpose of the policy family.
	Description types.String `tfsdk:"description"`
	// Name of the policy family.
	Name types.String `tfsdk:"name"`
	// Unique identifier for the policy family.
	PolicyFamilyId types.String `tfsdk:"policy_family_id"`
}

func (to *PolicyFamily) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PolicyFamily) {
}

func (to *PolicyFamily) SyncFieldsDuringRead(ctx context.Context, from PolicyFamily) {
}

func (m PolicyFamily) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["definition"] = attrs["definition"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["policy_family_id"] = attrs["policy_family_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PolicyFamily.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PolicyFamily) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PolicyFamily
// only implements ToObjectValue() and Type().
func (m PolicyFamily) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"definition":       m.Definition,
			"description":      m.Description,
			"name":             m.Name,
			"policy_family_id": m.PolicyFamilyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PolicyFamily) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"definition":       types.StringType,
			"description":      types.StringType,
			"name":             types.StringType,
			"policy_family_id": types.StringType,
		},
	}
}

type PythonPyPiLibrary struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	Package types.String `tfsdk:"package"`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo types.String `tfsdk:"repo"`
}

func (to *PythonPyPiLibrary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PythonPyPiLibrary) {
}

func (to *PythonPyPiLibrary) SyncFieldsDuringRead(ctx context.Context, from PythonPyPiLibrary) {
}

func (m PythonPyPiLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["package"] = attrs["package"].SetRequired()
	attrs["repo"] = attrs["repo"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PythonPyPiLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PythonPyPiLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PythonPyPiLibrary
// only implements ToObjectValue() and Type().
func (m PythonPyPiLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"package": m.Package,
			"repo":    m.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PythonPyPiLibrary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"package": types.StringType,
			"repo":    types.StringType,
		},
	}
}

type RCranLibrary struct {
	// The name of the CRAN package to install.
	Package types.String `tfsdk:"package"`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo types.String `tfsdk:"repo"`
}

func (to *RCranLibrary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RCranLibrary) {
}

func (to *RCranLibrary) SyncFieldsDuringRead(ctx context.Context, from RCranLibrary) {
}

func (m RCranLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["package"] = attrs["package"].SetRequired()
	attrs["repo"] = attrs["repo"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RCranLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RCranLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RCranLibrary
// only implements ToObjectValue() and Type().
func (m RCranLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"package": m.Package,
			"repo":    m.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RCranLibrary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"package": types.StringType,
			"repo":    types.StringType,
		},
	}
}

type RefreshDefaultBaseEnvironmentsRequest struct {
	Ids types.List `tfsdk:"ids"`
}

func (to *RefreshDefaultBaseEnvironmentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RefreshDefaultBaseEnvironmentsRequest) {
}

func (to *RefreshDefaultBaseEnvironmentsRequest) SyncFieldsDuringRead(ctx context.Context, from RefreshDefaultBaseEnvironmentsRequest) {
}

func (m RefreshDefaultBaseEnvironmentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ids"] = attrs["ids"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RefreshDefaultBaseEnvironmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RefreshDefaultBaseEnvironmentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RefreshDefaultBaseEnvironmentsRequest
// only implements ToObjectValue() and Type().
func (m RefreshDefaultBaseEnvironmentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ids": m.Ids,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RefreshDefaultBaseEnvironmentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetIds returns the value of the Ids field in RefreshDefaultBaseEnvironmentsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RefreshDefaultBaseEnvironmentsRequest) GetIds(ctx context.Context) ([]types.String, bool) {
	if m.Ids.IsNull() || m.Ids.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Ids.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIds sets the value of the Ids field in RefreshDefaultBaseEnvironmentsRequest.
func (m *RefreshDefaultBaseEnvironmentsRequest) SetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Ids = types.ListValueMust(t, vs)
}

type RefreshDefaultBaseEnvironmentsResponse struct {
}

func (to *RefreshDefaultBaseEnvironmentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RefreshDefaultBaseEnvironmentsResponse) {
}

func (to *RefreshDefaultBaseEnvironmentsResponse) SyncFieldsDuringRead(ctx context.Context, from RefreshDefaultBaseEnvironmentsResponse) {
}

func (m RefreshDefaultBaseEnvironmentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RefreshDefaultBaseEnvironmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RefreshDefaultBaseEnvironmentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RefreshDefaultBaseEnvironmentsResponse
// only implements ToObjectValue() and Type().
func (m RefreshDefaultBaseEnvironmentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RefreshDefaultBaseEnvironmentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RemoveInstanceProfile struct {
	// The ARN of the instance profile to remove. This field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
}

func (to *RemoveInstanceProfile) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveInstanceProfile) {
}

func (to *RemoveInstanceProfile) SyncFieldsDuringRead(ctx context.Context, from RemoveInstanceProfile) {
}

func (m RemoveInstanceProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveInstanceProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RemoveInstanceProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveInstanceProfile
// only implements ToObjectValue() and Type().
func (m RemoveInstanceProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_profile_arn": m.InstanceProfileArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RemoveInstanceProfile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_profile_arn": types.StringType,
		},
	}
}

type RemoveResponse struct {
}

func (to *RemoveResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveResponse) {
}

func (to *RemoveResponse) SyncFieldsDuringRead(ctx context.Context, from RemoveResponse) {
}

func (m RemoveResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RemoveResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveResponse
// only implements ToObjectValue() and Type().
func (m RemoveResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RemoveResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ResizeCluster struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale"`
	// The cluster to be resized.
	ClusterId types.String `tfsdk:"cluster_id"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers"`
}

func (to *ResizeCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResizeCluster) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				// Recursively sync the fields of Autoscale
				toAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
}

func (to *ResizeCluster) SyncFieldsDuringRead(ctx context.Context, from ResizeCluster) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				toAutoscale.SyncFieldsDuringRead(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
}

func (m ResizeCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["num_workers"] = attrs["num_workers"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResizeCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResizeCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale": reflect.TypeOf(AutoScale{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResizeCluster
// only implements ToObjectValue() and Type().
func (m ResizeCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":   m.Autoscale,
			"cluster_id":  m.ClusterId,
			"num_workers": m.NumWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResizeCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale":   AutoScale{}.Type(ctx),
			"cluster_id":  types.StringType,
			"num_workers": types.Int64Type,
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in ResizeCluster as
// a AutoScale value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResizeCluster) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v AutoScale
	d := m.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoscale sets the value of the Autoscale field in ResizeCluster.
func (m *ResizeCluster) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	m.Autoscale = vs
}

type ResizeClusterResponse struct {
}

func (to *ResizeClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResizeClusterResponse) {
}

func (to *ResizeClusterResponse) SyncFieldsDuringRead(ctx context.Context, from ResizeClusterResponse) {
}

func (m ResizeClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResizeClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResizeClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResizeClusterResponse
// only implements ToObjectValue() and Type().
func (m ResizeClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ResizeClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestartCluster struct {
	// The cluster to be started.
	ClusterId types.String `tfsdk:"cluster_id"`

	RestartUser types.String `tfsdk:"restart_user"`
}

func (to *RestartCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestartCluster) {
}

func (to *RestartCluster) SyncFieldsDuringRead(ctx context.Context, from RestartCluster) {
}

func (m RestartCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["restart_user"] = attrs["restart_user"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestartCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartCluster
// only implements ToObjectValue() and Type().
func (m RestartCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":   m.ClusterId,
			"restart_user": m.RestartUser,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestartCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":   types.StringType,
			"restart_user": types.StringType,
		},
	}
}

type RestartClusterResponse struct {
}

func (to *RestartClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestartClusterResponse) {
}

func (to *RestartClusterResponse) SyncFieldsDuringRead(ctx context.Context, from RestartClusterResponse) {
}

func (m RestartClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestartClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartClusterResponse
// only implements ToObjectValue() and Type().
func (m RestartClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RestartClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Results struct {
	// The cause of the error
	Cause types.String `tfsdk:"cause"`

	Data types.Object `tfsdk:"data"`
	// The image filename
	FileName types.String `tfsdk:"file_name"`

	FileNames types.List `tfsdk:"file_names"`
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	IsJsonSchema types.Bool `tfsdk:"is_json_schema"`
	// internal field used by SDK
	Pos types.Int64 `tfsdk:"pos"`

	ResultType types.String `tfsdk:"result_type"`
	// The table schema
	Schema types.List `tfsdk:"schema"`
	// The summary of the error
	Summary types.String `tfsdk:"summary"`
	// true if partial results are returned.
	Truncated types.Bool `tfsdk:"truncated"`
}

func (to *Results) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Results) {
	if !from.FileNames.IsNull() && !from.FileNames.IsUnknown() && to.FileNames.IsNull() && len(from.FileNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FileNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FileNames = from.FileNames
	}
	if !from.Schema.IsNull() && !from.Schema.IsUnknown() && to.Schema.IsNull() && len(from.Schema.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Schema, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Schema = from.Schema
	}
}

func (to *Results) SyncFieldsDuringRead(ctx context.Context, from Results) {
	if !from.FileNames.IsNull() && !from.FileNames.IsUnknown() && to.FileNames.IsNull() && len(from.FileNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FileNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FileNames = from.FileNames
	}
	if !from.Schema.IsNull() && !from.Schema.IsUnknown() && to.Schema.IsNull() && len(from.Schema.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Schema, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Schema = from.Schema
	}
}

func (m Results) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["data"] = attrs["data"].SetOptional()
	attrs["file_name"] = attrs["file_name"].SetOptional()
	attrs["file_names"] = attrs["file_names"].SetOptional()
	attrs["is_json_schema"] = attrs["is_json_schema"].SetOptional()
	attrs["pos"] = attrs["pos"].SetOptional()
	attrs["result_type"] = attrs["result_type"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["summary"] = attrs["summary"].SetOptional()
	attrs["truncated"] = attrs["truncated"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Results.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Results) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_names": reflect.TypeOf(types.String{}),
		"schema":     reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Results
// only implements ToObjectValue() and Type().
func (m Results) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cause":          m.Cause,
			"data":           m.Data,
			"file_name":      m.FileName,
			"file_names":     m.FileNames,
			"is_json_schema": m.IsJsonSchema,
			"pos":            m.Pos,
			"result_type":    m.ResultType,
			"schema":         m.Schema,
			"summary":        m.Summary,
			"truncated":      m.Truncated,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Results) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":     types.StringType,
			"data":      types.ObjectType{},
			"file_name": types.StringType,
			"file_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_json_schema": types.BoolType,
			"pos":            types.Int64Type,
			"result_type":    types.StringType,
			"schema": basetypes.ListType{
				ElemType: basetypes.MapType{
					ElemType: types.ObjectType{},
				},
			},
			"summary":   types.StringType,
			"truncated": types.BoolType,
		},
	}
}

// GetFileNames returns the value of the FileNames field in Results as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Results) GetFileNames(ctx context.Context) ([]types.String, bool) {
	if m.FileNames.IsNull() || m.FileNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.FileNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileNames sets the value of the FileNames field in Results.
func (m *Results) SetFileNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["file_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FileNames = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in Results as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *Results) GetSchema(ctx context.Context) ([]types.Object, bool) {
	if m.Schema.IsNull() || m.Schema.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Schema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchema sets the value of the Schema field in Results.
func (m *Results) SetSchema(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Schema = types.ListValueMust(t, vs)
}

// A storage location in Amazon S3
type S3StorageInfo struct {
	// (Optional) Set canned access control list for the logs, e.g.
	// `bucket-owner-full-control`. If `canned_cal` is set, please make sure the
	// cluster iam role has `s3:PutObjectAcl` permission on the destination
	// bucket and prefix. The full list of possible canned acl can be found at
	// http://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl.
	// Please also note that by default only the object owner gets full
	// controls. If you are using cross account role for writing data, you may
	// want to set `bucket-owner-full-control` to make bucket owner able to read
	// the logs.
	CannedAcl types.String `tfsdk:"canned_acl"`
	// S3 destination, e.g. `s3://my-bucket/some-prefix` Note that logs will be
	// delivered using cluster iam role, please make sure you set cluster iam
	// role and the role has write access to the destination. Please also note
	// that you cannot use AWS keys to deliver logs.
	Destination types.String `tfsdk:"destination"`
	// (Optional) Flag to enable server side encryption, `false` by default.
	EnableEncryption types.Bool `tfsdk:"enable_encryption"`
	// (Optional) The encryption type, it could be `sse-s3` or `sse-kms`. It
	// will be used only when encryption is enabled and the default type is
	// `sse-s3`.
	EncryptionType types.String `tfsdk:"encryption_type"`
	// S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`. Either region or
	// endpoint needs to be set. If both are set, endpoint will be used.
	Endpoint types.String `tfsdk:"endpoint"`
	// (Optional) Kms key which will be used if encryption is enabled and
	// encryption type is set to `sse-kms`.
	KmsKey types.String `tfsdk:"kms_key"`
	// S3 region, e.g. `us-west-2`. Either region or endpoint needs to be set.
	// If both are set, endpoint will be used.
	Region types.String `tfsdk:"region"`
}

func (to *S3StorageInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from S3StorageInfo) {
}

func (to *S3StorageInfo) SyncFieldsDuringRead(ctx context.Context, from S3StorageInfo) {
}

func (m S3StorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["canned_acl"] = attrs["canned_acl"].SetOptional()
	attrs["destination"] = attrs["destination"].SetRequired()
	attrs["enable_encryption"] = attrs["enable_encryption"].SetOptional()
	attrs["encryption_type"] = attrs["encryption_type"].SetOptional()
	attrs["endpoint"] = attrs["endpoint"].SetOptional()
	attrs["kms_key"] = attrs["kms_key"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in S3StorageInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m S3StorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, S3StorageInfo
// only implements ToObjectValue() and Type().
func (m S3StorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"canned_acl":        m.CannedAcl,
			"destination":       m.Destination,
			"enable_encryption": m.EnableEncryption,
			"encryption_type":   m.EncryptionType,
			"endpoint":          m.Endpoint,
			"kms_key":           m.KmsKey,
			"region":            m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m S3StorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"canned_acl":        types.StringType,
			"destination":       types.StringType,
			"enable_encryption": types.BoolType,
			"encryption_type":   types.StringType,
			"endpoint":          types.StringType,
			"kms_key":           types.StringType,
			"region":            types.StringType,
		},
	}
}

// Describes a specific Spark driver or executor.
type SparkNode struct {
	// The private IP address of the host instance.
	HostPrivateIp types.String `tfsdk:"host_private_ip"`
	// Globally unique identifier for the host instance from the cloud provider.
	InstanceId types.String `tfsdk:"instance_id"`
	// Attributes specific to AWS for a Spark node.
	NodeAwsAttributes types.Object `tfsdk:"node_aws_attributes"`
	// Globally unique identifier for this node.
	NodeId types.String `tfsdk:"node_id"`
	// Private IP address (typically a 10.x.x.x address) of the Spark node. Note
	// that this is different from the private IP address of the host instance.
	PrivateIp types.String `tfsdk:"private_ip"`
	// Public DNS address of this node. This address can be used to access the
	// Spark JDBC server on the driver node. To communicate with the JDBC
	// server, traffic must be manually authorized by adding security group
	// rules to the "worker-unmanaged" security group via the AWS console.
	PublicDns types.String `tfsdk:"public_dns"`
	// The timestamp (in millisecond) when the Spark node is launched.
	StartTimestamp types.Int64 `tfsdk:"start_timestamp"`
}

func (to *SparkNode) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SparkNode) {
	if !from.NodeAwsAttributes.IsNull() && !from.NodeAwsAttributes.IsUnknown() {
		if toNodeAwsAttributes, ok := to.GetNodeAwsAttributes(ctx); ok {
			if fromNodeAwsAttributes, ok := from.GetNodeAwsAttributes(ctx); ok {
				// Recursively sync the fields of NodeAwsAttributes
				toNodeAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromNodeAwsAttributes)
				to.SetNodeAwsAttributes(ctx, toNodeAwsAttributes)
			}
		}
	}
}

func (to *SparkNode) SyncFieldsDuringRead(ctx context.Context, from SparkNode) {
	if !from.NodeAwsAttributes.IsNull() && !from.NodeAwsAttributes.IsUnknown() {
		if toNodeAwsAttributes, ok := to.GetNodeAwsAttributes(ctx); ok {
			if fromNodeAwsAttributes, ok := from.GetNodeAwsAttributes(ctx); ok {
				toNodeAwsAttributes.SyncFieldsDuringRead(ctx, fromNodeAwsAttributes)
				to.SetNodeAwsAttributes(ctx, toNodeAwsAttributes)
			}
		}
	}
}

func (m SparkNode) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["host_private_ip"] = attrs["host_private_ip"].SetOptional()
	attrs["instance_id"] = attrs["instance_id"].SetOptional()
	attrs["node_aws_attributes"] = attrs["node_aws_attributes"].SetOptional()
	attrs["node_id"] = attrs["node_id"].SetOptional()
	attrs["private_ip"] = attrs["private_ip"].SetOptional()
	attrs["public_dns"] = attrs["public_dns"].SetOptional()
	attrs["start_timestamp"] = attrs["start_timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkNode.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SparkNode) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_aws_attributes": reflect.TypeOf(SparkNodeAwsAttributes{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkNode
// only implements ToObjectValue() and Type().
func (m SparkNode) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"host_private_ip":     m.HostPrivateIp,
			"instance_id":         m.InstanceId,
			"node_aws_attributes": m.NodeAwsAttributes,
			"node_id":             m.NodeId,
			"private_ip":          m.PrivateIp,
			"public_dns":          m.PublicDns,
			"start_timestamp":     m.StartTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SparkNode) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"host_private_ip":     types.StringType,
			"instance_id":         types.StringType,
			"node_aws_attributes": SparkNodeAwsAttributes{}.Type(ctx),
			"node_id":             types.StringType,
			"private_ip":          types.StringType,
			"public_dns":          types.StringType,
			"start_timestamp":     types.Int64Type,
		},
	}
}

// GetNodeAwsAttributes returns the value of the NodeAwsAttributes field in SparkNode as
// a SparkNodeAwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *SparkNode) GetNodeAwsAttributes(ctx context.Context) (SparkNodeAwsAttributes, bool) {
	var e SparkNodeAwsAttributes
	if m.NodeAwsAttributes.IsNull() || m.NodeAwsAttributes.IsUnknown() {
		return e, false
	}
	var v SparkNodeAwsAttributes
	d := m.NodeAwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeAwsAttributes sets the value of the NodeAwsAttributes field in SparkNode.
func (m *SparkNode) SetNodeAwsAttributes(ctx context.Context, v SparkNodeAwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.NodeAwsAttributes = vs
}

// Attributes specific to AWS for a Spark node.
type SparkNodeAwsAttributes struct {
	// Whether this node is on an Amazon spot instance.
	IsSpot types.Bool `tfsdk:"is_spot"`
}

func (to *SparkNodeAwsAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SparkNodeAwsAttributes) {
}

func (to *SparkNodeAwsAttributes) SyncFieldsDuringRead(ctx context.Context, from SparkNodeAwsAttributes) {
}

func (m SparkNodeAwsAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_spot"] = attrs["is_spot"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkNodeAwsAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SparkNodeAwsAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkNodeAwsAttributes
// only implements ToObjectValue() and Type().
func (m SparkNodeAwsAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_spot": m.IsSpot,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SparkNodeAwsAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_spot": types.BoolType,
		},
	}
}

type SparkVersion struct {
	// Spark version key, for example "2.1.x-scala2.11". This is the value which
	// should be provided as the "spark_version" when creating a new cluster.
	// Note that the exact Spark version may change over time for a "wildcard"
	// version (i.e., "2.1.x-scala2.11" is a "wildcard" version) with minor bug
	// fixes.
	Key types.String `tfsdk:"key"`
	// A descriptive name for this Spark version, for example "Spark 2.1".
	Name types.String `tfsdk:"name"`
}

func (to *SparkVersion) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SparkVersion) {
}

func (to *SparkVersion) SyncFieldsDuringRead(ctx context.Context, from SparkVersion) {
}

func (m SparkVersion) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkVersion.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SparkVersion) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkVersion
// only implements ToObjectValue() and Type().
func (m SparkVersion) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":  m.Key,
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SparkVersion) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":  types.StringType,
			"name": types.StringType,
		},
	}
}

type SparkVersionsRequest struct {
}

func (to *SparkVersionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SparkVersionsRequest) {
}

func (to *SparkVersionsRequest) SyncFieldsDuringRead(ctx context.Context, from SparkVersionsRequest) {
}

func (m SparkVersionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SparkVersionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkVersionsRequest
// only implements ToObjectValue() and Type().
func (m SparkVersionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SparkVersionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type StartCluster struct {
	// The cluster to be started.
	ClusterId types.String `tfsdk:"cluster_id"`
}

func (to *StartCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartCluster) {
}

func (to *StartCluster) SyncFieldsDuringRead(ctx context.Context, from StartCluster) {
}

func (m StartCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StartCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartCluster
// only implements ToObjectValue() and Type().
func (m StartCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StartCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type StartClusterResponse struct {
}

func (to *StartClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartClusterResponse) {
}

func (to *StartClusterResponse) SyncFieldsDuringRead(ctx context.Context, from StartClusterResponse) {
}

func (m StartClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StartClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartClusterResponse
// only implements ToObjectValue() and Type().
func (m StartClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m StartClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code types.String `tfsdk:"code"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters types.Map `tfsdk:"parameters"`
	// type of the termination
	Type_ types.String `tfsdk:"type"`
}

func (to *TerminationReason) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TerminationReason) {
}

func (to *TerminationReason) SyncFieldsDuringRead(ctx context.Context, from TerminationReason) {
}

func (m TerminationReason) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["code"] = attrs["code"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TerminationReason.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TerminationReason) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TerminationReason
// only implements ToObjectValue() and Type().
func (m TerminationReason) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"code":       m.Code,
			"parameters": m.Parameters,
			"type":       m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TerminationReason) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"code": types.StringType,
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"type": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in TerminationReason as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TerminationReason) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in TerminationReason.
func (m *TerminationReason) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.MapValueMust(t, vs)
}

type UninstallLibraries struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The libraries to uninstall.
	Libraries types.List `tfsdk:"libraries"`
}

func (to *UninstallLibraries) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UninstallLibraries) {
}

func (to *UninstallLibraries) SyncFieldsDuringRead(ctx context.Context, from UninstallLibraries) {
}

func (m UninstallLibraries) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["libraries"] = attrs["libraries"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UninstallLibraries.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UninstallLibraries) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UninstallLibraries
// only implements ToObjectValue() and Type().
func (m UninstallLibraries) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
			"libraries":  m.Libraries,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UninstallLibraries) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"libraries": basetypes.ListType{
				ElemType: Library{}.Type(ctx),
			},
		},
	}
}

// GetLibraries returns the value of the Libraries field in UninstallLibraries as
// a slice of Library values.
// If the field is unknown or null, the boolean return value is false.
func (m *UninstallLibraries) GetLibraries(ctx context.Context) ([]Library, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in UninstallLibraries.
func (m *UninstallLibraries) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

type UninstallLibrariesResponse struct {
}

func (to *UninstallLibrariesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UninstallLibrariesResponse) {
}

func (to *UninstallLibrariesResponse) SyncFieldsDuringRead(ctx context.Context, from UninstallLibrariesResponse) {
}

func (m UninstallLibrariesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UninstallLibrariesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UninstallLibrariesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UninstallLibrariesResponse
// only implements ToObjectValue() and Type().
func (m UninstallLibrariesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UninstallLibrariesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UnpinCluster struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

func (to *UnpinCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UnpinCluster) {
}

func (to *UnpinCluster) SyncFieldsDuringRead(ctx context.Context, from UnpinCluster) {
}

func (m UnpinCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpinCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UnpinCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpinCluster
// only implements ToObjectValue() and Type().
func (m UnpinCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": m.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UnpinCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type UnpinClusterResponse struct {
}

func (to *UnpinClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UnpinClusterResponse) {
}

func (to *UnpinClusterResponse) SyncFieldsDuringRead(ctx context.Context, from UnpinClusterResponse) {
}

func (m UnpinClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpinClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UnpinClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpinClusterResponse
// only implements ToObjectValue() and Type().
func (m UnpinClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UnpinClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateCluster struct {
	// The cluster to be updated.
	Cluster types.Object `tfsdk:"cluster"`
	// ID of the cluster.
	ClusterId types.String `tfsdk:"cluster_id"`
	// Used to specify which cluster attributes and size fields to update. See
	// https://google.aip.dev/161 for more details.
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"update_mask"`
}

func (to *UpdateCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCluster) {
	if !from.Cluster.IsNull() && !from.Cluster.IsUnknown() {
		if toCluster, ok := to.GetCluster(ctx); ok {
			if fromCluster, ok := from.GetCluster(ctx); ok {
				// Recursively sync the fields of Cluster
				toCluster.SyncFieldsDuringCreateOrUpdate(ctx, fromCluster)
				to.SetCluster(ctx, toCluster)
			}
		}
	}
}

func (to *UpdateCluster) SyncFieldsDuringRead(ctx context.Context, from UpdateCluster) {
	if !from.Cluster.IsNull() && !from.Cluster.IsUnknown() {
		if toCluster, ok := to.GetCluster(ctx); ok {
			if fromCluster, ok := from.GetCluster(ctx); ok {
				toCluster.SyncFieldsDuringRead(ctx, fromCluster)
				to.SetCluster(ctx, toCluster)
			}
		}
	}
}

func (m UpdateCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster"] = attrs["cluster"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster": reflect.TypeOf(UpdateClusterResource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCluster
// only implements ToObjectValue() and Type().
func (m UpdateCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster":     m.Cluster,
			"cluster_id":  m.ClusterId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster":     UpdateClusterResource{}.Type(ctx),
			"cluster_id":  types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetCluster returns the value of the Cluster field in UpdateCluster as
// a UpdateClusterResource value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCluster) GetCluster(ctx context.Context) (UpdateClusterResource, bool) {
	var e UpdateClusterResource
	if m.Cluster.IsNull() || m.Cluster.IsUnknown() {
		return e, false
	}
	var v UpdateClusterResource
	d := m.Cluster.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCluster sets the value of the Cluster field in UpdateCluster.
func (m *UpdateCluster) SetCluster(ctx context.Context, v UpdateClusterResource) {
	vs := v.ToObjectValue(ctx)
	m.Cluster = vs
}

type UpdateClusterResource struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.Object `tfsdk:"cluster_log_conf"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName types.String `tfsdk:"cluster_name"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags"`

	DataSecurityMode types.String `tfsdk:"data_security_mode"`
	// Custom docker image BYOC
	DockerImage types.Object `tfsdk:"docker_image"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode types.Bool `tfsdk:"is_single_node"`

	Kind types.String `tfsdk:"kind"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput types.Int64 `tfsdk:"remote_disk_throughput"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars types.Map `tfsdk:"spark_env_vars"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime types.Bool `tfsdk:"use_ml_runtime"`

	WorkloadType types.Object `tfsdk:"workload_type"`
}

func (to *UpdateClusterResource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateClusterResource) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				// Recursively sync the fields of Autoscale
				toAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				// Recursively sync the fields of ClusterLogConf
				toClusterLogConf.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				// Recursively sync the fields of DockerImage
				toDockerImage.SyncFieldsDuringCreateOrUpdate(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				// Recursively sync the fields of WorkloadType
				toWorkloadType.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (to *UpdateClusterResource) SyncFieldsDuringRead(ctx context.Context, from UpdateClusterResource) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				toAutoscale.SyncFieldsDuringRead(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				toClusterLogConf.SyncFieldsDuringRead(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.DockerImage.IsNull() && !from.DockerImage.IsUnknown() {
		if toDockerImage, ok := to.GetDockerImage(ctx); ok {
			if fromDockerImage, ok := from.GetDockerImage(ctx); ok {
				toDockerImage.SyncFieldsDuringRead(ctx, fromDockerImage)
				to.SetDockerImage(ctx, toDockerImage)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
	if !from.WorkloadType.IsNull() && !from.WorkloadType.IsUnknown() {
		if toWorkloadType, ok := to.GetWorkloadType(ctx); ok {
			if fromWorkloadType, ok := from.GetWorkloadType(ctx); ok {
				toWorkloadType.SyncFieldsDuringRead(ctx, fromWorkloadType)
				to.SetWorkloadType(ctx, toWorkloadType)
			}
		}
	}
}

func (m UpdateClusterResource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["is_single_node"] = attrs["is_single_node"].SetOptional()
	attrs["kind"] = attrs["kind"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["num_workers"] = attrs["num_workers"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["runtime_engine"] = attrs["runtime_engine"].SetOptional()
	attrs["single_user_name"] = attrs["single_user_name"].SetOptional()
	attrs["spark_conf"] = attrs["spark_conf"].SetOptional()
	attrs["spark_env_vars"] = attrs["spark_env_vars"].SetOptional()
	attrs["spark_version"] = attrs["spark_version"].SetOptional()
	attrs["ssh_public_keys"] = attrs["ssh_public_keys"].SetOptional()
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()
	attrs["use_ml_runtime"] = attrs["use_ml_runtime"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateClusterResource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateClusterResource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(AutoScale{}),
		"aws_attributes":   reflect.TypeOf(AwsAttributes{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateClusterResource
// only implements ToObjectValue() and Type().
func (m UpdateClusterResource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":                      m.Autoscale,
			"autotermination_minutes":        m.AutoterminationMinutes,
			"aws_attributes":                 m.AwsAttributes,
			"azure_attributes":               m.AzureAttributes,
			"cluster_log_conf":               m.ClusterLogConf,
			"cluster_name":                   m.ClusterName,
			"custom_tags":                    m.CustomTags,
			"data_security_mode":             m.DataSecurityMode,
			"docker_image":                   m.DockerImage,
			"driver_instance_pool_id":        m.DriverInstancePoolId,
			"driver_node_type_id":            m.DriverNodeTypeId,
			"enable_elastic_disk":            m.EnableElasticDisk,
			"enable_local_disk_encryption":   m.EnableLocalDiskEncryption,
			"gcp_attributes":                 m.GcpAttributes,
			"init_scripts":                   m.InitScripts,
			"instance_pool_id":               m.InstancePoolId,
			"is_single_node":                 m.IsSingleNode,
			"kind":                           m.Kind,
			"node_type_id":                   m.NodeTypeId,
			"num_workers":                    m.NumWorkers,
			"policy_id":                      m.PolicyId,
			"remote_disk_throughput":         m.RemoteDiskThroughput,
			"runtime_engine":                 m.RuntimeEngine,
			"single_user_name":               m.SingleUserName,
			"spark_conf":                     m.SparkConf,
			"spark_env_vars":                 m.SparkEnvVars,
			"spark_version":                  m.SparkVersion,
			"ssh_public_keys":                m.SshPublicKeys,
			"total_initial_remote_disk_size": m.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 m.UseMlRuntime,
			"workload_type":                  m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateClusterResource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale":               AutoScale{}.Type(ctx),
			"autotermination_minutes": types.Int64Type,
			"aws_attributes":          AwsAttributes{}.Type(ctx),
			"azure_attributes":        AzureAttributes{}.Type(ctx),
			"cluster_log_conf":        ClusterLogConf{}.Type(ctx),
			"cluster_name":            types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode":           types.StringType,
			"docker_image":                 DockerImage{}.Type(ctx),
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes":               GcpAttributes{}.Type(ctx),
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id":       types.StringType,
			"is_single_node":         types.BoolType,
			"kind":                   types.StringType,
			"node_type_id":           types.StringType,
			"num_workers":            types.Int64Type,
			"policy_id":              types.StringType,
			"remote_disk_throughput": types.Int64Type,
			"runtime_engine":         types.StringType,
			"single_user_name":       types.StringType,
			"spark_conf": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_env_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_version": types.StringType,
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"total_initial_remote_disk_size": types.Int64Type,
			"use_ml_runtime":                 types.BoolType,
			"workload_type":                  WorkloadType{}.Type(ctx),
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in UpdateClusterResource as
// a AutoScale value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v AutoScale
	d := m.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoscale sets the value of the Autoscale field in UpdateClusterResource.
func (m *UpdateClusterResource) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	m.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in UpdateClusterResource as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v AwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in UpdateClusterResource.
func (m *UpdateClusterResource) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in UpdateClusterResource as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v AzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in UpdateClusterResource.
func (m *UpdateClusterResource) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in UpdateClusterResource as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if m.ClusterLogConf.IsNull() || m.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v ClusterLogConf
	d := m.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in UpdateClusterResource.
func (m *UpdateClusterResource) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	m.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in UpdateClusterResource as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in UpdateClusterResource.
func (m *UpdateClusterResource) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in UpdateClusterResource as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if m.DockerImage.IsNull() || m.DockerImage.IsUnknown() {
		return e, false
	}
	var v DockerImage
	d := m.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDockerImage sets the value of the DockerImage field in UpdateClusterResource.
func (m *UpdateClusterResource) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	m.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in UpdateClusterResource as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v GcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in UpdateClusterResource.
func (m *UpdateClusterResource) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in UpdateClusterResource as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := m.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in UpdateClusterResource.
func (m *UpdateClusterResource) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in UpdateClusterResource as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkConf.IsNull() || m.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in UpdateClusterResource.
func (m *UpdateClusterResource) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in UpdateClusterResource as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkEnvVars.IsNull() || m.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in UpdateClusterResource.
func (m *UpdateClusterResource) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in UpdateClusterResource as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if m.SshPublicKeys.IsNull() || m.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in UpdateClusterResource.
func (m *UpdateClusterResource) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in UpdateClusterResource as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateClusterResource) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if m.WorkloadType.IsNull() || m.WorkloadType.IsUnknown() {
		return e, false
	}
	var v WorkloadType
	d := m.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkloadType sets the value of the WorkloadType field in UpdateClusterResource.
func (m *UpdateClusterResource) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	m.WorkloadType = vs
}

type UpdateClusterResponse struct {
}

func (to *UpdateClusterResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateClusterResponse) {
}

func (to *UpdateClusterResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateClusterResponse) {
}

func (m UpdateClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateClusterResponse
// only implements ToObjectValue() and Type().
func (m UpdateClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateDefaultBaseEnvironmentRequest struct {
	DefaultBaseEnvironment types.Object `tfsdk:"default_base_environment"`

	Id types.String `tfsdk:"-"`
}

func (to *UpdateDefaultBaseEnvironmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDefaultBaseEnvironmentRequest) {
	if !from.DefaultBaseEnvironment.IsNull() && !from.DefaultBaseEnvironment.IsUnknown() {
		if toDefaultBaseEnvironment, ok := to.GetDefaultBaseEnvironment(ctx); ok {
			if fromDefaultBaseEnvironment, ok := from.GetDefaultBaseEnvironment(ctx); ok {
				// Recursively sync the fields of DefaultBaseEnvironment
				toDefaultBaseEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultBaseEnvironment)
				to.SetDefaultBaseEnvironment(ctx, toDefaultBaseEnvironment)
			}
		}
	}
}

func (to *UpdateDefaultBaseEnvironmentRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDefaultBaseEnvironmentRequest) {
	if !from.DefaultBaseEnvironment.IsNull() && !from.DefaultBaseEnvironment.IsUnknown() {
		if toDefaultBaseEnvironment, ok := to.GetDefaultBaseEnvironment(ctx); ok {
			if fromDefaultBaseEnvironment, ok := from.GetDefaultBaseEnvironment(ctx); ok {
				toDefaultBaseEnvironment.SyncFieldsDuringRead(ctx, fromDefaultBaseEnvironment)
				to.SetDefaultBaseEnvironment(ctx, toDefaultBaseEnvironment)
			}
		}
	}
}

func (m UpdateDefaultBaseEnvironmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_base_environment"] = attrs["default_base_environment"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDefaultBaseEnvironmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_base_environment": reflect.TypeOf(DefaultBaseEnvironment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultBaseEnvironmentRequest
// only implements ToObjectValue() and Type().
func (m UpdateDefaultBaseEnvironmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_base_environment": m.DefaultBaseEnvironment,
			"id":                       m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDefaultBaseEnvironmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_base_environment": DefaultBaseEnvironment{}.Type(ctx),
			"id":                       types.StringType,
		},
	}
}

// GetDefaultBaseEnvironment returns the value of the DefaultBaseEnvironment field in UpdateDefaultBaseEnvironmentRequest as
// a DefaultBaseEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDefaultBaseEnvironmentRequest) GetDefaultBaseEnvironment(ctx context.Context) (DefaultBaseEnvironment, bool) {
	var e DefaultBaseEnvironment
	if m.DefaultBaseEnvironment.IsNull() || m.DefaultBaseEnvironment.IsUnknown() {
		return e, false
	}
	var v DefaultBaseEnvironment
	d := m.DefaultBaseEnvironment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultBaseEnvironment sets the value of the DefaultBaseEnvironment field in UpdateDefaultBaseEnvironmentRequest.
func (m *UpdateDefaultBaseEnvironmentRequest) SetDefaultBaseEnvironment(ctx context.Context, v DefaultBaseEnvironment) {
	vs := v.ToObjectValue(ctx)
	m.DefaultBaseEnvironment = vs
}

type UpdateDefaultDefaultBaseEnvironmentRequest struct {
	BaseEnvironmentType types.String `tfsdk:"base_environment_type"`

	Id types.String `tfsdk:"id"`
}

func (to *UpdateDefaultDefaultBaseEnvironmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDefaultDefaultBaseEnvironmentRequest) {
}

func (to *UpdateDefaultDefaultBaseEnvironmentRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDefaultDefaultBaseEnvironmentRequest) {
}

func (m UpdateDefaultDefaultBaseEnvironmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["base_environment_type"] = attrs["base_environment_type"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDefaultDefaultBaseEnvironmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultDefaultBaseEnvironmentRequest
// only implements ToObjectValue() and Type().
func (m UpdateDefaultDefaultBaseEnvironmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"base_environment_type": m.BaseEnvironmentType,
			"id":                    m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDefaultDefaultBaseEnvironmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_environment_type": types.StringType,
			"id":                    types.StringType,
		},
	}
}

type UpdateResponse struct {
}

func (to *UpdateResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateResponse) {
}

func (to *UpdateResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateResponse) {
}

func (m UpdateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse
// only implements ToObjectValue() and Type().
func (m UpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A storage location back by UC Volumes.
type VolumesStorageInfo struct {
	// UC Volumes destination, e.g.
	// `/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh` or
	// `dbfs:/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh`
	Destination types.String `tfsdk:"destination"`
}

func (to *VolumesStorageInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VolumesStorageInfo) {
}

func (to *VolumesStorageInfo) SyncFieldsDuringRead(ctx context.Context, from VolumesStorageInfo) {
}

func (m VolumesStorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination"] = attrs["destination"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VolumesStorageInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VolumesStorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VolumesStorageInfo
// only implements ToObjectValue() and Type().
func (m VolumesStorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": m.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VolumesStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

// Cluster Attributes showing for clusters workload types.
type WorkloadType struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	Clients types.Object `tfsdk:"clients"`
}

func (to *WorkloadType) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkloadType) {
	if !from.Clients.IsNull() && !from.Clients.IsUnknown() {
		if toClients, ok := to.GetClients(ctx); ok {
			if fromClients, ok := from.GetClients(ctx); ok {
				// Recursively sync the fields of Clients
				toClients.SyncFieldsDuringCreateOrUpdate(ctx, fromClients)
				to.SetClients(ctx, toClients)
			}
		}
	}
}

func (to *WorkloadType) SyncFieldsDuringRead(ctx context.Context, from WorkloadType) {
	if !from.Clients.IsNull() && !from.Clients.IsUnknown() {
		if toClients, ok := to.GetClients(ctx); ok {
			if fromClients, ok := from.GetClients(ctx); ok {
				toClients.SyncFieldsDuringRead(ctx, fromClients)
				to.SetClients(ctx, toClients)
			}
		}
	}
}

func (m WorkloadType) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clients"] = attrs["clients"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkloadType.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkloadType) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clients": reflect.TypeOf(ClientsTypes{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkloadType
// only implements ToObjectValue() and Type().
func (m WorkloadType) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clients": m.Clients,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkloadType) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clients": ClientsTypes{}.Type(ctx),
		},
	}
}

// GetClients returns the value of the Clients field in WorkloadType as
// a ClientsTypes value.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkloadType) GetClients(ctx context.Context) (ClientsTypes, bool) {
	var e ClientsTypes
	if m.Clients.IsNull() || m.Clients.IsUnknown() {
		return e, false
	}
	var v ClientsTypes
	d := m.Clients.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClients sets the value of the Clients field in WorkloadType.
func (m *WorkloadType) SetClients(ctx context.Context, v ClientsTypes) {
	vs := v.ToObjectValue(ctx)
	m.Clients = vs
}

// A storage location in Workspace Filesystem (WSFS)
type WorkspaceStorageInfo struct {
	// wsfs destination, e.g. `workspace:/cluster-init-scripts/setup-datadog.sh`
	Destination types.String `tfsdk:"destination"`
}

func (to *WorkspaceStorageInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceStorageInfo) {
}

func (to *WorkspaceStorageInfo) SyncFieldsDuringRead(ctx context.Context, from WorkspaceStorageInfo) {
}

func (m WorkspaceStorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination"] = attrs["destination"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceStorageInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceStorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceStorageInfo
// only implements ToObjectValue() and Type().
func (m WorkspaceStorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": m.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}
