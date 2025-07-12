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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddInstanceProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddInstanceProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddInstanceProfile
// only implements ToObjectValue() and Type().
func (o AddInstanceProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"iam_role_arn":             o.IamRoleArn,
			"instance_profile_arn":     o.InstanceProfileArn,
			"is_meta_instance_profile": o.IsMetaInstanceProfile,
			"skip_validation":          o.SkipValidation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AddInstanceProfile) Type(ctx context.Context) attr.Type {
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

func (newState *AddResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddResponse) {
}

func (newState *AddResponse) SyncEffectiveFieldsDuringRead(existingState AddResponse) {
}

func (c AddResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddResponse
// only implements ToObjectValue() and Type().
func (o AddResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o AddResponse) Type(ctx context.Context) attr.Type {
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

func (newState *Adlsgen2Info) SyncEffectiveFieldsDuringCreateOrUpdate(plan Adlsgen2Info) {
}

func (newState *Adlsgen2Info) SyncEffectiveFieldsDuringRead(existingState Adlsgen2Info) {
}

func (c Adlsgen2Info) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Adlsgen2Info) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Adlsgen2Info
// only implements ToObjectValue() and Type().
func (o Adlsgen2Info) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Adlsgen2Info) Type(ctx context.Context) attr.Type {
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

func (newState *AutoScale) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoScale) {
}

func (newState *AutoScale) SyncEffectiveFieldsDuringRead(existingState AutoScale) {
}

func (c AutoScale) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AutoScale) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoScale
// only implements ToObjectValue() and Type().
func (o AutoScale) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_workers": o.MaxWorkers,
			"min_workers": o.MinWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AutoScale) Type(ctx context.Context) attr.Type {
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

func (newState *AwsAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsAttributes) {
}

func (newState *AwsAttributes) SyncEffectiveFieldsDuringRead(existingState AwsAttributes) {
}

func (c AwsAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AwsAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsAttributes
// only implements ToObjectValue() and Type().
func (o AwsAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":           o.Availability,
			"ebs_volume_count":       o.EbsVolumeCount,
			"ebs_volume_iops":        o.EbsVolumeIops,
			"ebs_volume_size":        o.EbsVolumeSize,
			"ebs_volume_throughput":  o.EbsVolumeThroughput,
			"ebs_volume_type":        o.EbsVolumeType,
			"first_on_demand":        o.FirstOnDemand,
			"instance_profile_arn":   o.InstanceProfileArn,
			"spot_bid_price_percent": o.SpotBidPricePercent,
			"zone_id":                o.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsAttributes) Type(ctx context.Context) attr.Type {
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

func (newState *AzureAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureAttributes) {
}

func (newState *AzureAttributes) SyncEffectiveFieldsDuringRead(existingState AzureAttributes) {
}

func (c AzureAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AzureAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_analytics_info": reflect.TypeOf(LogAnalyticsInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureAttributes
// only implements ToObjectValue() and Type().
func (o AzureAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":       o.Availability,
			"first_on_demand":    o.FirstOnDemand,
			"log_analytics_info": o.LogAnalyticsInfo,
			"spot_bid_max_price": o.SpotBidMaxPrice,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureAttributes) Type(ctx context.Context) attr.Type {
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
func (o *AzureAttributes) GetLogAnalyticsInfo(ctx context.Context) (LogAnalyticsInfo, bool) {
	var e LogAnalyticsInfo
	if o.LogAnalyticsInfo.IsNull() || o.LogAnalyticsInfo.IsUnknown() {
		return e, false
	}
	var v []LogAnalyticsInfo
	d := o.LogAnalyticsInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogAnalyticsInfo sets the value of the LogAnalyticsInfo field in AzureAttributes.
func (o *AzureAttributes) SetLogAnalyticsInfo(ctx context.Context, v LogAnalyticsInfo) {
	vs := v.ToObjectValue(ctx)
	o.LogAnalyticsInfo = vs
}

type CancelCommand struct {
	ClusterId types.String `tfsdk:"clusterId"`

	CommandId types.String `tfsdk:"commandId"`

	ContextId types.String `tfsdk:"contextId"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelCommand.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelCommand) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelCommand
// only implements ToObjectValue() and Type().
func (o CancelCommand) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"commandId": o.CommandId,
			"contextId": o.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelCommand) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"commandId": types.StringType,
			"contextId": types.StringType,
		},
	}
}

type CancelResponse struct {
}

func (newState *CancelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelResponse) {
}

func (newState *CancelResponse) SyncEffectiveFieldsDuringRead(existingState CancelResponse) {
}

func (c CancelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelResponse
// only implements ToObjectValue() and Type().
func (o CancelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CancelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ChangeClusterOwner struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	// New owner of the cluster_id after this RPC.
	OwnerUsername types.String `tfsdk:"owner_username"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChangeClusterOwner.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ChangeClusterOwner) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChangeClusterOwner
// only implements ToObjectValue() and Type().
func (o ChangeClusterOwner) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":     o.ClusterId,
			"owner_username": o.OwnerUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ChangeClusterOwner) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":     types.StringType,
			"owner_username": types.StringType,
		},
	}
}

type ChangeClusterOwnerResponse struct {
}

func (newState *ChangeClusterOwnerResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ChangeClusterOwnerResponse) {
}

func (newState *ChangeClusterOwnerResponse) SyncEffectiveFieldsDuringRead(existingState ChangeClusterOwnerResponse) {
}

func (c ChangeClusterOwnerResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChangeClusterOwnerResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ChangeClusterOwnerResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChangeClusterOwnerResponse
// only implements ToObjectValue() and Type().
func (o ChangeClusterOwnerResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse) Type(ctx context.Context) attr.Type {
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

func (newState *ClientsTypes) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClientsTypes) {
}

func (newState *ClientsTypes) SyncEffectiveFieldsDuringRead(existingState ClientsTypes) {
}

func (c ClientsTypes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClientsTypes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClientsTypes
// only implements ToObjectValue() and Type().
func (o ClientsTypes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"jobs":      o.Jobs,
			"notebooks": o.Notebooks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClientsTypes) Type(ctx context.Context) attr.Type {
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

func (newState *CloneCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloneCluster) {
}

func (newState *CloneCluster) SyncEffectiveFieldsDuringRead(existingState CloneCluster) {
}

func (c CloneCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CloneCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloneCluster
// only implements ToObjectValue() and Type().
func (o CloneCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"source_cluster_id": o.SourceClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloneCluster) Type(ctx context.Context) attr.Type {
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

func (newState *CloudProviderNodeInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudProviderNodeInfo) {
}

func (newState *CloudProviderNodeInfo) SyncEffectiveFieldsDuringRead(existingState CloudProviderNodeInfo) {
}

func (c CloudProviderNodeInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CloudProviderNodeInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"status": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudProviderNodeInfo
// only implements ToObjectValue() and Type().
func (o CloudProviderNodeInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloudProviderNodeInfo) Type(ctx context.Context) attr.Type {
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
func (o *CloudProviderNodeInfo) GetStatus(ctx context.Context) ([]types.String, bool) {
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in CloudProviderNodeInfo.
func (o *CloudProviderNodeInfo) SetStatus(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Status = types.ListValueMust(t, vs)
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

func (newState *ClusterAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAccessControlRequest) {
}

func (newState *ClusterAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState ClusterAccessControlRequest) {
}

func (c ClusterAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAccessControlRequest
// only implements ToObjectValue() and Type().
func (o ClusterAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ClusterAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAccessControlResponse) {
}

func (newState *ClusterAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState ClusterAccessControlResponse) {
}

func (c ClusterAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ClusterPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAccessControlResponse
// only implements ToObjectValue() and Type().
func (o ClusterAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (o *ClusterAccessControlResponse) GetAllPermissions(ctx context.Context) ([]ClusterPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ClusterPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ClusterAccessControlResponse.
func (o *ClusterAccessControlResponse) SetAllPermissions(ctx context.Context, v []ClusterPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
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

func (newState *ClusterAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAttributes) {
}

func (newState *ClusterAttributes) SyncEffectiveFieldsDuringRead(existingState ClusterAttributes) {
}

func (c ClusterAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o ClusterAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autotermination_minutes":        o.AutoterminationMinutes,
			"aws_attributes":                 o.AwsAttributes,
			"azure_attributes":               o.AzureAttributes,
			"cluster_log_conf":               o.ClusterLogConf,
			"cluster_name":                   o.ClusterName,
			"custom_tags":                    o.CustomTags,
			"data_security_mode":             o.DataSecurityMode,
			"docker_image":                   o.DockerImage,
			"driver_instance_pool_id":        o.DriverInstancePoolId,
			"driver_node_type_id":            o.DriverNodeTypeId,
			"enable_elastic_disk":            o.EnableElasticDisk,
			"enable_local_disk_encryption":   o.EnableLocalDiskEncryption,
			"gcp_attributes":                 o.GcpAttributes,
			"init_scripts":                   o.InitScripts,
			"instance_pool_id":               o.InstancePoolId,
			"is_single_node":                 o.IsSingleNode,
			"kind":                           o.Kind,
			"node_type_id":                   o.NodeTypeId,
			"policy_id":                      o.PolicyId,
			"remote_disk_throughput":         o.RemoteDiskThroughput,
			"runtime_engine":                 o.RuntimeEngine,
			"single_user_name":               o.SingleUserName,
			"spark_conf":                     o.SparkConf,
			"spark_env_vars":                 o.SparkEnvVars,
			"spark_version":                  o.SparkVersion,
			"ssh_public_keys":                o.SshPublicKeys,
			"total_initial_remote_disk_size": o.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 o.UseMlRuntime,
			"workload_type":                  o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAttributes) Type(ctx context.Context) attr.Type {
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
func (o *ClusterAttributes) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterAttributes.
func (o *ClusterAttributes) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterAttributes as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterAttributes.
func (o *ClusterAttributes) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterAttributes as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf
	d := o.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterAttributes.
func (o *ClusterAttributes) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	o.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in ClusterAttributes as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in ClusterAttributes.
func (o *ClusterAttributes) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterAttributes as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage
	d := o.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in ClusterAttributes.
func (o *ClusterAttributes) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	o.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterAttributes as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterAttributes.
func (o *ClusterAttributes) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in ClusterAttributes as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterAttributes.
func (o *ClusterAttributes) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterAttributes as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkConf.IsNull() || o.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in ClusterAttributes.
func (o *ClusterAttributes) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterAttributes as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkEnvVars.IsNull() || o.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterAttributes.
func (o *ClusterAttributes) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterAttributes as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if o.SshPublicKeys.IsNull() || o.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterAttributes.
func (o *ClusterAttributes) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterAttributes as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType
	d := o.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterAttributes.
func (o *ClusterAttributes) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	o.WorkloadType = vs
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

func (newState *ClusterCompliance) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterCompliance) {
}

func (newState *ClusterCompliance) SyncEffectiveFieldsDuringRead(existingState ClusterCompliance) {
}

func (c ClusterCompliance) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterCompliance) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterCompliance
// only implements ToObjectValue() and Type().
func (o ClusterCompliance) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":   o.ClusterId,
			"is_compliant": o.IsCompliant,
			"violations":   o.Violations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterCompliance) Type(ctx context.Context) attr.Type {
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
func (o *ClusterCompliance) GetViolations(ctx context.Context) (map[string]types.String, bool) {
	if o.Violations.IsNull() || o.Violations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Violations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViolations sets the value of the Violations field in ClusterCompliance.
func (o *ClusterCompliance) SetViolations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["violations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Violations = types.MapValueMust(t, vs)
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

func (newState *ClusterDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterDetails) {
}

func (newState *ClusterDetails) SyncEffectiveFieldsDuringRead(existingState ClusterDetails) {
}

func (c ClusterDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o ClusterDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":                      o.Autoscale,
			"autotermination_minutes":        o.AutoterminationMinutes,
			"aws_attributes":                 o.AwsAttributes,
			"azure_attributes":               o.AzureAttributes,
			"cluster_cores":                  o.ClusterCores,
			"cluster_id":                     o.ClusterId,
			"cluster_log_conf":               o.ClusterLogConf,
			"cluster_log_status":             o.ClusterLogStatus,
			"cluster_memory_mb":              o.ClusterMemoryMb,
			"cluster_name":                   o.ClusterName,
			"cluster_source":                 o.ClusterSource,
			"creator_user_name":              o.CreatorUserName,
			"custom_tags":                    o.CustomTags,
			"data_security_mode":             o.DataSecurityMode,
			"default_tags":                   o.DefaultTags,
			"docker_image":                   o.DockerImage,
			"driver":                         o.Driver,
			"driver_instance_pool_id":        o.DriverInstancePoolId,
			"driver_node_type_id":            o.DriverNodeTypeId,
			"enable_elastic_disk":            o.EnableElasticDisk,
			"enable_local_disk_encryption":   o.EnableLocalDiskEncryption,
			"executors":                      o.Executors,
			"gcp_attributes":                 o.GcpAttributes,
			"init_scripts":                   o.InitScripts,
			"instance_pool_id":               o.InstancePoolId,
			"is_single_node":                 o.IsSingleNode,
			"jdbc_port":                      o.JdbcPort,
			"kind":                           o.Kind,
			"last_restarted_time":            o.LastRestartedTime,
			"last_state_loss_time":           o.LastStateLossTime,
			"node_type_id":                   o.NodeTypeId,
			"num_workers":                    o.NumWorkers,
			"policy_id":                      o.PolicyId,
			"remote_disk_throughput":         o.RemoteDiskThroughput,
			"runtime_engine":                 o.RuntimeEngine,
			"single_user_name":               o.SingleUserName,
			"spark_conf":                     o.SparkConf,
			"spark_context_id":               o.SparkContextId,
			"spark_env_vars":                 o.SparkEnvVars,
			"spark_version":                  o.SparkVersion,
			"spec":                           o.Spec,
			"ssh_public_keys":                o.SshPublicKeys,
			"start_time":                     o.StartTime,
			"state":                          o.State,
			"state_message":                  o.StateMessage,
			"terminated_time":                o.TerminatedTime,
			"termination_reason":             o.TerminationReason,
			"total_initial_remote_disk_size": o.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 o.UseMlRuntime,
			"workload_type":                  o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterDetails) Type(ctx context.Context) attr.Type {
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
func (o *ClusterDetails) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in ClusterDetails.
func (o *ClusterDetails) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	o.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in ClusterDetails as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterDetails.
func (o *ClusterDetails) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterDetails as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterDetails.
func (o *ClusterDetails) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterDetails as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf
	d := o.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterDetails.
func (o *ClusterDetails) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	o.ClusterLogConf = vs
}

// GetClusterLogStatus returns the value of the ClusterLogStatus field in ClusterDetails as
// a LogSyncStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetClusterLogStatus(ctx context.Context) (LogSyncStatus, bool) {
	var e LogSyncStatus
	if o.ClusterLogStatus.IsNull() || o.ClusterLogStatus.IsUnknown() {
		return e, false
	}
	var v []LogSyncStatus
	d := o.ClusterLogStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogStatus sets the value of the ClusterLogStatus field in ClusterDetails.
func (o *ClusterDetails) SetClusterLogStatus(ctx context.Context, v LogSyncStatus) {
	vs := v.ToObjectValue(ctx)
	o.ClusterLogStatus = vs
}

// GetCustomTags returns the value of the CustomTags field in ClusterDetails as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in ClusterDetails.
func (o *ClusterDetails) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in ClusterDetails as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
	if o.DefaultTags.IsNull() || o.DefaultTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.DefaultTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultTags sets the value of the DefaultTags field in ClusterDetails.
func (o *ClusterDetails) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DefaultTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterDetails as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage
	d := o.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in ClusterDetails.
func (o *ClusterDetails) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	o.DockerImage = vs
}

// GetDriver returns the value of the Driver field in ClusterDetails as
// a SparkNode value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetDriver(ctx context.Context) (SparkNode, bool) {
	var e SparkNode
	if o.Driver.IsNull() || o.Driver.IsUnknown() {
		return e, false
	}
	var v []SparkNode
	d := o.Driver.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDriver sets the value of the Driver field in ClusterDetails.
func (o *ClusterDetails) SetDriver(ctx context.Context, v SparkNode) {
	vs := v.ToObjectValue(ctx)
	o.Driver = vs
}

// GetExecutors returns the value of the Executors field in ClusterDetails as
// a slice of SparkNode values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetExecutors(ctx context.Context) ([]SparkNode, bool) {
	if o.Executors.IsNull() || o.Executors.IsUnknown() {
		return nil, false
	}
	var v []SparkNode
	d := o.Executors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExecutors sets the value of the Executors field in ClusterDetails.
func (o *ClusterDetails) SetExecutors(ctx context.Context, v []SparkNode) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["executors"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Executors = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterDetails as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterDetails.
func (o *ClusterDetails) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in ClusterDetails as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterDetails.
func (o *ClusterDetails) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterDetails as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkConf.IsNull() || o.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in ClusterDetails.
func (o *ClusterDetails) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterDetails as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkEnvVars.IsNull() || o.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterDetails.
func (o *ClusterDetails) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSpec returns the value of the Spec field in ClusterDetails as
// a ClusterSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetSpec(ctx context.Context) (ClusterSpec, bool) {
	var e ClusterSpec
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []ClusterSpec
	d := o.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in ClusterDetails.
func (o *ClusterDetails) SetSpec(ctx context.Context, v ClusterSpec) {
	vs := v.ToObjectValue(ctx)
	o.Spec = vs
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterDetails as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if o.SshPublicKeys.IsNull() || o.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterDetails.
func (o *ClusterDetails) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetTerminationReason returns the value of the TerminationReason field in ClusterDetails as
// a TerminationReason value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetTerminationReason(ctx context.Context) (TerminationReason, bool) {
	var e TerminationReason
	if o.TerminationReason.IsNull() || o.TerminationReason.IsUnknown() {
		return e, false
	}
	var v []TerminationReason
	d := o.TerminationReason.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTerminationReason sets the value of the TerminationReason field in ClusterDetails.
func (o *ClusterDetails) SetTerminationReason(ctx context.Context, v TerminationReason) {
	vs := v.ToObjectValue(ctx)
	o.TerminationReason = vs
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterDetails as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType
	d := o.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterDetails.
func (o *ClusterDetails) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	o.WorkloadType = vs
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

func (newState *ClusterEvent) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterEvent) {
}

func (newState *ClusterEvent) SyncEffectiveFieldsDuringRead(existingState ClusterEvent) {
}

func (c ClusterEvent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterEvent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_plane_event_details": reflect.TypeOf(DataPlaneEventDetails{}),
		"details":                  reflect.TypeOf(EventDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterEvent
// only implements ToObjectValue() and Type().
func (o ClusterEvent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":               o.ClusterId,
			"data_plane_event_details": o.DataPlaneEventDetails,
			"details":                  o.Details,
			"timestamp":                o.Timestamp,
			"type":                     o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterEvent) Type(ctx context.Context) attr.Type {
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
func (o *ClusterEvent) GetDataPlaneEventDetails(ctx context.Context) (DataPlaneEventDetails, bool) {
	var e DataPlaneEventDetails
	if o.DataPlaneEventDetails.IsNull() || o.DataPlaneEventDetails.IsUnknown() {
		return e, false
	}
	var v []DataPlaneEventDetails
	d := o.DataPlaneEventDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataPlaneEventDetails sets the value of the DataPlaneEventDetails field in ClusterEvent.
func (o *ClusterEvent) SetDataPlaneEventDetails(ctx context.Context, v DataPlaneEventDetails) {
	vs := v.ToObjectValue(ctx)
	o.DataPlaneEventDetails = vs
}

// GetDetails returns the value of the Details field in ClusterEvent as
// a EventDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterEvent) GetDetails(ctx context.Context) (EventDetails, bool) {
	var e EventDetails
	if o.Details.IsNull() || o.Details.IsUnknown() {
		return e, false
	}
	var v []EventDetails
	d := o.Details.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDetails sets the value of the Details field in ClusterEvent.
func (o *ClusterEvent) SetDetails(ctx context.Context, v EventDetails) {
	vs := v.ToObjectValue(ctx)
	o.Details = vs
}

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	ClusterId types.String `tfsdk:"cluster_id"`
	// Status of all libraries on the cluster.
	LibraryStatuses types.List `tfsdk:"library_statuses"`
}

func (newState *ClusterLibraryStatuses) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterLibraryStatuses) {
}

func (newState *ClusterLibraryStatuses) SyncEffectiveFieldsDuringRead(existingState ClusterLibraryStatuses) {
}

func (c ClusterLibraryStatuses) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterLibraryStatuses) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"library_statuses": reflect.TypeOf(LibraryFullStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterLibraryStatuses
// only implements ToObjectValue() and Type().
func (o ClusterLibraryStatuses) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":       o.ClusterId,
			"library_statuses": o.LibraryStatuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterLibraryStatuses) Type(ctx context.Context) attr.Type {
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
func (o *ClusterLibraryStatuses) GetLibraryStatuses(ctx context.Context) ([]LibraryFullStatus, bool) {
	if o.LibraryStatuses.IsNull() || o.LibraryStatuses.IsUnknown() {
		return nil, false
	}
	var v []LibraryFullStatus
	d := o.LibraryStatuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraryStatuses sets the value of the LibraryStatuses field in ClusterLibraryStatuses.
func (o *ClusterLibraryStatuses) SetLibraryStatuses(ctx context.Context, v []LibraryFullStatus) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["library_statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LibraryStatuses = types.ListValueMust(t, vs)
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

func (newState *ClusterLogConf) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterLogConf) {
}

func (newState *ClusterLogConf) SyncEffectiveFieldsDuringRead(existingState ClusterLogConf) {
}

func (c ClusterLogConf) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterLogConf) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbfs":    reflect.TypeOf(DbfsStorageInfo{}),
		"s3":      reflect.TypeOf(S3StorageInfo{}),
		"volumes": reflect.TypeOf(VolumesStorageInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterLogConf
// only implements ToObjectValue() and Type().
func (o ClusterLogConf) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbfs":    o.Dbfs,
			"s3":      o.S3,
			"volumes": o.Volumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterLogConf) Type(ctx context.Context) attr.Type {
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
func (o *ClusterLogConf) GetDbfs(ctx context.Context) (DbfsStorageInfo, bool) {
	var e DbfsStorageInfo
	if o.Dbfs.IsNull() || o.Dbfs.IsUnknown() {
		return e, false
	}
	var v []DbfsStorageInfo
	d := o.Dbfs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbfs sets the value of the Dbfs field in ClusterLogConf.
func (o *ClusterLogConf) SetDbfs(ctx context.Context, v DbfsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Dbfs = vs
}

// GetS3 returns the value of the S3 field in ClusterLogConf as
// a S3StorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterLogConf) GetS3(ctx context.Context) (S3StorageInfo, bool) {
	var e S3StorageInfo
	if o.S3.IsNull() || o.S3.IsUnknown() {
		return e, false
	}
	var v []S3StorageInfo
	d := o.S3.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetS3 sets the value of the S3 field in ClusterLogConf.
func (o *ClusterLogConf) SetS3(ctx context.Context, v S3StorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.S3 = vs
}

// GetVolumes returns the value of the Volumes field in ClusterLogConf as
// a VolumesStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterLogConf) GetVolumes(ctx context.Context) (VolumesStorageInfo, bool) {
	var e VolumesStorageInfo
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return e, false
	}
	var v []VolumesStorageInfo
	d := o.Volumes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumes sets the value of the Volumes field in ClusterLogConf.
func (o *ClusterLogConf) SetVolumes(ctx context.Context, v VolumesStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Volumes = vs
}

type ClusterPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ClusterPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPermission) {
}

func (newState *ClusterPermission) SyncEffectiveFieldsDuringRead(existingState ClusterPermission) {
}

func (c ClusterPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermission
// only implements ToObjectValue() and Type().
func (o ClusterPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPermission) Type(ctx context.Context) attr.Type {
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
func (o *ClusterPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in ClusterPermission.
func (o *ClusterPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type ClusterPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *ClusterPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPermissions) {
}

func (newState *ClusterPermissions) SyncEffectiveFieldsDuringRead(existingState ClusterPermissions) {
}

func (c ClusterPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissions
// only implements ToObjectValue() and Type().
func (o ClusterPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPermissions) Type(ctx context.Context) attr.Type {
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
func (o *ClusterPermissions) GetAccessControlList(ctx context.Context) ([]ClusterAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPermissions.
func (o *ClusterPermissions) SetAccessControlList(ctx context.Context, v []ClusterAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ClusterPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ClusterPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPermissionsDescription) {
}

func (newState *ClusterPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState ClusterPermissionsDescription) {
}

func (c ClusterPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissionsDescription
// only implements ToObjectValue() and Type().
func (o ClusterPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPermissionsDescription) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissionsRequest
// only implements ToObjectValue() and Type().
func (o ClusterPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"cluster_id":          o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *ClusterPermissionsRequest) GetAccessControlList(ctx context.Context) ([]ClusterAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPermissionsRequest.
func (o *ClusterPermissionsRequest) SetAccessControlList(ctx context.Context, v []ClusterAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
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

func (newState *ClusterPolicyAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyAccessControlRequest) {
}

func (newState *ClusterPolicyAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyAccessControlRequest) {
}

func (c ClusterPolicyAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyAccessControlRequest
// only implements ToObjectValue() and Type().
func (o ClusterPolicyAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ClusterPolicyAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyAccessControlResponse) {
}

func (newState *ClusterPolicyAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyAccessControlResponse) {
}

func (c ClusterPolicyAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ClusterPolicyPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyAccessControlResponse
// only implements ToObjectValue() and Type().
func (o ClusterPolicyAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (o *ClusterPolicyAccessControlResponse) GetAllPermissions(ctx context.Context) ([]ClusterPolicyPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ClusterPolicyAccessControlResponse.
func (o *ClusterPolicyAccessControlResponse) SetAllPermissions(ctx context.Context, v []ClusterPolicyPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type ClusterPolicyPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ClusterPolicyPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyPermission) {
}

func (newState *ClusterPolicyPermission) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyPermission) {
}

func (c ClusterPolicyPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermission
// only implements ToObjectValue() and Type().
func (o ClusterPolicyPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyPermission) Type(ctx context.Context) attr.Type {
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
func (o *ClusterPolicyPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in ClusterPolicyPermission.
func (o *ClusterPolicyPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type ClusterPolicyPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *ClusterPolicyPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyPermissions) {
}

func (newState *ClusterPolicyPermissions) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyPermissions) {
}

func (c ClusterPolicyPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterPolicyAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissions
// only implements ToObjectValue() and Type().
func (o ClusterPolicyPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissions) Type(ctx context.Context) attr.Type {
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
func (o *ClusterPolicyPermissions) GetAccessControlList(ctx context.Context) ([]ClusterPolicyAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPolicyPermissions.
func (o *ClusterPolicyPermissions) SetAccessControlList(ctx context.Context, v []ClusterPolicyAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ClusterPolicyPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ClusterPolicyPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyPermissionsDescription) {
}

func (newState *ClusterPolicyPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyPermissionsDescription) {
}

func (c ClusterPolicyPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissionsDescription
// only implements ToObjectValue() and Type().
func (o ClusterPolicyPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsDescription) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterPolicyPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterPolicyPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterPolicyAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissionsRequest
// only implements ToObjectValue() and Type().
func (o ClusterPolicyPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"cluster_policy_id":   o.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *ClusterPolicyPermissionsRequest) GetAccessControlList(ctx context.Context) ([]ClusterPolicyAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPolicyPermissionsRequest.
func (o *ClusterPolicyPermissionsRequest) SetAccessControlList(ctx context.Context, v []ClusterPolicyAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
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

func (newState *ClusterSettingsChange) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterSettingsChange) {
}

func (newState *ClusterSettingsChange) SyncEffectiveFieldsDuringRead(existingState ClusterSettingsChange) {
}

func (c ClusterSettingsChange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterSettingsChange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSettingsChange
// only implements ToObjectValue() and Type().
func (o ClusterSettingsChange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"field":          o.Field,
			"new_value":      o.NewValue,
			"previous_value": o.PreviousValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterSettingsChange) Type(ctx context.Context) attr.Type {
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

func (newState *ClusterSize) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterSize) {
}

func (newState *ClusterSize) SyncEffectiveFieldsDuringRead(existingState ClusterSize) {
}

func (c ClusterSize) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterSize) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale": reflect.TypeOf(AutoScale{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSize
// only implements ToObjectValue() and Type().
func (o ClusterSize) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":   o.Autoscale,
			"num_workers": o.NumWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterSize) Type(ctx context.Context) attr.Type {
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
func (o *ClusterSize) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in ClusterSize.
func (o *ClusterSize) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	o.Autoscale = vs
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

func (newState *ClusterSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterSpec) {
}

func (newState *ClusterSpec) SyncEffectiveFieldsDuringRead(existingState ClusterSpec) {
}

func (c ClusterSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o ClusterSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_policy_default_values":    o.ApplyPolicyDefaultValues,
			"autoscale":                      o.Autoscale,
			"autotermination_minutes":        o.AutoterminationMinutes,
			"aws_attributes":                 o.AwsAttributes,
			"azure_attributes":               o.AzureAttributes,
			"cluster_log_conf":               o.ClusterLogConf,
			"cluster_name":                   o.ClusterName,
			"custom_tags":                    o.CustomTags,
			"data_security_mode":             o.DataSecurityMode,
			"docker_image":                   o.DockerImage,
			"driver_instance_pool_id":        o.DriverInstancePoolId,
			"driver_node_type_id":            o.DriverNodeTypeId,
			"enable_elastic_disk":            o.EnableElasticDisk,
			"enable_local_disk_encryption":   o.EnableLocalDiskEncryption,
			"gcp_attributes":                 o.GcpAttributes,
			"init_scripts":                   o.InitScripts,
			"instance_pool_id":               o.InstancePoolId,
			"is_single_node":                 o.IsSingleNode,
			"kind":                           o.Kind,
			"node_type_id":                   o.NodeTypeId,
			"num_workers":                    o.NumWorkers,
			"policy_id":                      o.PolicyId,
			"remote_disk_throughput":         o.RemoteDiskThroughput,
			"runtime_engine":                 o.RuntimeEngine,
			"single_user_name":               o.SingleUserName,
			"spark_conf":                     o.SparkConf,
			"spark_env_vars":                 o.SparkEnvVars,
			"spark_version":                  o.SparkVersion,
			"ssh_public_keys":                o.SshPublicKeys,
			"total_initial_remote_disk_size": o.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 o.UseMlRuntime,
			"workload_type":                  o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterSpec) Type(ctx context.Context) attr.Type {
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
func (o *ClusterSpec) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in ClusterSpec.
func (o *ClusterSpec) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	o.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in ClusterSpec as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterSpec.
func (o *ClusterSpec) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterSpec as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterSpec.
func (o *ClusterSpec) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterSpec as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf
	d := o.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterSpec.
func (o *ClusterSpec) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	o.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in ClusterSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in ClusterSpec.
func (o *ClusterSpec) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterSpec as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage
	d := o.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in ClusterSpec.
func (o *ClusterSpec) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	o.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterSpec as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterSpec.
func (o *ClusterSpec) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in ClusterSpec as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterSpec.
func (o *ClusterSpec) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkConf.IsNull() || o.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in ClusterSpec.
func (o *ClusterSpec) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkEnvVars.IsNull() || o.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterSpec.
func (o *ClusterSpec) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterSpec as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if o.SshPublicKeys.IsNull() || o.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterSpec.
func (o *ClusterSpec) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterSpec as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType
	d := o.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterSpec.
func (o *ClusterSpec) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	o.WorkloadType = vs
}

type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterStatus
// only implements ToObjectValue() and Type().
func (o ClusterStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type Command struct {
	// Running cluster id
	ClusterId types.String `tfsdk:"clusterId"`
	// Executable code
	Command types.String `tfsdk:"command"`
	// Running context id
	ContextId types.String `tfsdk:"contextId"`

	Language types.String `tfsdk:"language"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Command.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Command) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Command
// only implements ToObjectValue() and Type().
func (o Command) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"command":   o.Command,
			"contextId": o.ContextId,
			"language":  o.Language,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Command) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"command":   types.StringType,
			"contextId": types.StringType,
			"language":  types.StringType,
		},
	}
}

type CommandStatusRequest struct {
	ClusterId types.String `tfsdk:"-"`

	CommandId types.String `tfsdk:"-"`

	ContextId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CommandStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CommandStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommandStatusRequest
// only implements ToObjectValue() and Type().
func (o CommandStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"commandId": o.CommandId,
			"contextId": o.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CommandStatusRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"commandId": types.StringType,
			"contextId": types.StringType,
		},
	}
}

type CommandStatusResponse struct {
	Id types.String `tfsdk:"id"`

	Results types.Object `tfsdk:"results"`

	Status types.String `tfsdk:"status"`
}

func (newState *CommandStatusResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CommandStatusResponse) {
}

func (newState *CommandStatusResponse) SyncEffectiveFieldsDuringRead(existingState CommandStatusResponse) {
}

func (c CommandStatusResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CommandStatusResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Results{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommandStatusResponse
// only implements ToObjectValue() and Type().
func (o CommandStatusResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      o.Id,
			"results": o.Results,
			"status":  o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CommandStatusResponse) Type(ctx context.Context) attr.Type {
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
func (o *CommandStatusResponse) GetResults(ctx context.Context) (Results, bool) {
	var e Results
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return e, false
	}
	var v []Results
	d := o.Results.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResults sets the value of the Results field in CommandStatusResponse.
func (o *CommandStatusResponse) SetResults(ctx context.Context, v Results) {
	vs := v.ToObjectValue(ctx)
	o.Results = vs
}

type ContextStatusRequest struct {
	ClusterId types.String `tfsdk:"-"`

	ContextId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ContextStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ContextStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContextStatusRequest
// only implements ToObjectValue() and Type().
func (o ContextStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"contextId": o.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ContextStatusRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"contextId": types.StringType,
		},
	}
}

type ContextStatusResponse struct {
	Id types.String `tfsdk:"id"`

	Status types.String `tfsdk:"status"`
}

func (newState *ContextStatusResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContextStatusResponse) {
}

func (newState *ContextStatusResponse) SyncEffectiveFieldsDuringRead(existingState ContextStatusResponse) {
}

func (c ContextStatusResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ContextStatusResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContextStatusResponse
// only implements ToObjectValue() and Type().
func (o ContextStatusResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":     o.Id,
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ContextStatusResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o CreateCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_policy_default_values":    o.ApplyPolicyDefaultValues,
			"autoscale":                      o.Autoscale,
			"autotermination_minutes":        o.AutoterminationMinutes,
			"aws_attributes":                 o.AwsAttributes,
			"azure_attributes":               o.AzureAttributes,
			"clone_from":                     o.CloneFrom,
			"cluster_log_conf":               o.ClusterLogConf,
			"cluster_name":                   o.ClusterName,
			"custom_tags":                    o.CustomTags,
			"data_security_mode":             o.DataSecurityMode,
			"docker_image":                   o.DockerImage,
			"driver_instance_pool_id":        o.DriverInstancePoolId,
			"driver_node_type_id":            o.DriverNodeTypeId,
			"enable_elastic_disk":            o.EnableElasticDisk,
			"enable_local_disk_encryption":   o.EnableLocalDiskEncryption,
			"gcp_attributes":                 o.GcpAttributes,
			"init_scripts":                   o.InitScripts,
			"instance_pool_id":               o.InstancePoolId,
			"is_single_node":                 o.IsSingleNode,
			"kind":                           o.Kind,
			"node_type_id":                   o.NodeTypeId,
			"num_workers":                    o.NumWorkers,
			"policy_id":                      o.PolicyId,
			"remote_disk_throughput":         o.RemoteDiskThroughput,
			"runtime_engine":                 o.RuntimeEngine,
			"single_user_name":               o.SingleUserName,
			"spark_conf":                     o.SparkConf,
			"spark_env_vars":                 o.SparkEnvVars,
			"spark_version":                  o.SparkVersion,
			"ssh_public_keys":                o.SshPublicKeys,
			"total_initial_remote_disk_size": o.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 o.UseMlRuntime,
			"workload_type":                  o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCluster) Type(ctx context.Context) attr.Type {
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
func (o *CreateCluster) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in CreateCluster.
func (o *CreateCluster) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	o.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in CreateCluster as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in CreateCluster.
func (o *CreateCluster) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in CreateCluster as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in CreateCluster.
func (o *CreateCluster) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetCloneFrom returns the value of the CloneFrom field in CreateCluster as
// a CloneCluster value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetCloneFrom(ctx context.Context) (CloneCluster, bool) {
	var e CloneCluster
	if o.CloneFrom.IsNull() || o.CloneFrom.IsUnknown() {
		return e, false
	}
	var v []CloneCluster
	d := o.CloneFrom.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloneFrom sets the value of the CloneFrom field in CreateCluster.
func (o *CreateCluster) SetCloneFrom(ctx context.Context, v CloneCluster) {
	vs := v.ToObjectValue(ctx)
	o.CloneFrom = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in CreateCluster as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf
	d := o.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in CreateCluster.
func (o *CreateCluster) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	o.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in CreateCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in CreateCluster.
func (o *CreateCluster) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in CreateCluster as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage
	d := o.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in CreateCluster.
func (o *CreateCluster) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	o.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in CreateCluster as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in CreateCluster.
func (o *CreateCluster) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in CreateCluster as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in CreateCluster.
func (o *CreateCluster) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in CreateCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkConf.IsNull() || o.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in CreateCluster.
func (o *CreateCluster) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in CreateCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkEnvVars.IsNull() || o.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in CreateCluster.
func (o *CreateCluster) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in CreateCluster as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if o.SshPublicKeys.IsNull() || o.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in CreateCluster.
func (o *CreateCluster) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in CreateCluster as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType
	d := o.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in CreateCluster.
func (o *CreateCluster) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	o.WorkloadType = vs
}

type CreateClusterResponse struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

func (newState *CreateClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateClusterResponse) {
}

func (newState *CreateClusterResponse) SyncEffectiveFieldsDuringRead(existingState CreateClusterResponse) {
}

func (c CreateClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateClusterResponse
// only implements ToObjectValue() and Type().
func (o CreateClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type CreateContext struct {
	// Running cluster id
	ClusterId types.String `tfsdk:"clusterId"`

	Language types.String `tfsdk:"language"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateContext.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateContext) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateContext
// only implements ToObjectValue() and Type().
func (o CreateContext) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"language":  o.Language,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateContext) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"language":  types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateInstancePool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateInstancePool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstancePool
// only implements ToObjectValue() and Type().
func (o CreateInstancePool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        o.AwsAttributes,
			"azure_attributes":                      o.AzureAttributes,
			"custom_tags":                           o.CustomTags,
			"disk_spec":                             o.DiskSpec,
			"enable_elastic_disk":                   o.EnableElasticDisk,
			"gcp_attributes":                        o.GcpAttributes,
			"idle_instance_autotermination_minutes": o.IdleInstanceAutoterminationMinutes,
			"instance_pool_name":                    o.InstancePoolName,
			"max_capacity":                          o.MaxCapacity,
			"min_idle_instances":                    o.MinIdleInstances,
			"node_type_id":                          o.NodeTypeId,
			"preloaded_docker_images":               o.PreloadedDockerImages,
			"preloaded_spark_versions":              o.PreloadedSparkVersions,
			"remote_disk_throughput":                o.RemoteDiskThroughput,
			"total_initial_remote_disk_size":        o.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes":   InstancePoolAwsAttributes{}.Type(ctx),
			"azure_attributes": InstancePoolAzureAttributes{}.Type(ctx),
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec":                             DiskSpec{}.Type(ctx),
			"enable_elastic_disk":                   types.BoolType,
			"gcp_attributes":                        InstancePoolGcpAttributes{}.Type(ctx),
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
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
func (o *CreateInstancePool) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes, bool) {
	var e InstancePoolAwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in CreateInstancePool.
func (o *CreateInstancePool) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in CreateInstancePool as
// a InstancePoolAzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes, bool) {
	var e InstancePoolAzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in CreateInstancePool.
func (o *CreateInstancePool) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetCustomTags returns the value of the CustomTags field in CreateInstancePool as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in CreateInstancePool.
func (o *CreateInstancePool) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in CreateInstancePool as
// a DiskSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool) GetDiskSpec(ctx context.Context) (DiskSpec, bool) {
	var e DiskSpec
	if o.DiskSpec.IsNull() || o.DiskSpec.IsUnknown() {
		return e, false
	}
	var v []DiskSpec
	d := o.DiskSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDiskSpec sets the value of the DiskSpec field in CreateInstancePool.
func (o *CreateInstancePool) SetDiskSpec(ctx context.Context, v DiskSpec) {
	vs := v.ToObjectValue(ctx)
	o.DiskSpec = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in CreateInstancePool as
// a InstancePoolGcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes, bool) {
	var e InstancePoolGcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolGcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in CreateInstancePool.
func (o *CreateInstancePool) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in CreateInstancePool as
// a slice of DockerImage values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage, bool) {
	if o.PreloadedDockerImages.IsNull() || o.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage
	d := o.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in CreateInstancePool.
func (o *CreateInstancePool) SetPreloadedDockerImages(ctx context.Context, v []DockerImage) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in CreateInstancePool as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
	if o.PreloadedSparkVersions.IsNull() || o.PreloadedSparkVersions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PreloadedSparkVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in CreateInstancePool.
func (o *CreateInstancePool) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
}

func (newState *CreateInstancePoolResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateInstancePoolResponse) {
}

func (newState *CreateInstancePoolResponse) SyncEffectiveFieldsDuringRead(existingState CreateInstancePoolResponse) {
}

func (c CreateInstancePoolResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateInstancePoolResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstancePoolResponse
// only implements ToObjectValue() and Type().
func (o CreateInstancePoolResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateInstancePoolResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePolicy
// only implements ToObjectValue() and Type().
func (o CreatePolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"definition":                         o.Definition,
			"description":                        o.Description,
			"libraries":                          o.Libraries,
			"max_clusters_per_user":              o.MaxClustersPerUser,
			"name":                               o.Name,
			"policy_family_definition_overrides": o.PolicyFamilyDefinitionOverrides,
			"policy_family_id":                   o.PolicyFamilyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePolicy) Type(ctx context.Context) attr.Type {
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
func (o *CreatePolicy) GetLibraries(ctx context.Context) ([]Library, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in CreatePolicy.
func (o *CreatePolicy) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId types.String `tfsdk:"policy_id"`
}

func (newState *CreatePolicyResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePolicyResponse) {
}

func (newState *CreatePolicyResponse) SyncEffectiveFieldsDuringRead(existingState CreatePolicyResponse) {
}

func (c CreatePolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreatePolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePolicyResponse
// only implements ToObjectValue() and Type().
func (o CreatePolicyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePolicyResponse) Type(ctx context.Context) attr.Type {
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

func (newState *CreateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateResponse) {
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringRead(existingState CreateResponse) {
}

func (c CreateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateResponse
// only implements ToObjectValue() and Type().
func (o CreateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": o.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"script_id": types.StringType,
		},
	}
}

type Created struct {
	Id types.String `tfsdk:"id"`
}

func (newState *Created) SyncEffectiveFieldsDuringCreateOrUpdate(plan Created) {
}

func (newState *Created) SyncEffectiveFieldsDuringRead(existingState Created) {
}

func (c Created) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Created) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Created
// only implements ToObjectValue() and Type().
func (o Created) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Created) Type(ctx context.Context) attr.Type {
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

func (newState *CustomPolicyTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan CustomPolicyTag) {
}

func (newState *CustomPolicyTag) SyncEffectiveFieldsDuringRead(existingState CustomPolicyTag) {
}

func (c CustomPolicyTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CustomPolicyTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomPolicyTag
// only implements ToObjectValue() and Type().
func (o CustomPolicyTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomPolicyTag) Type(ctx context.Context) attr.Type {
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

func (newState *DataPlaneEventDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataPlaneEventDetails) {
}

func (newState *DataPlaneEventDetails) SyncEffectiveFieldsDuringRead(existingState DataPlaneEventDetails) {
}

func (c DataPlaneEventDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DataPlaneEventDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneEventDetails
// only implements ToObjectValue() and Type().
func (o DataPlaneEventDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"event_type":        o.EventType,
			"executor_failures": o.ExecutorFailures,
			"host_id":           o.HostId,
			"timestamp":         o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataPlaneEventDetails) Type(ctx context.Context) attr.Type {
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

func (newState *DbfsStorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan DbfsStorageInfo) {
}

func (newState *DbfsStorageInfo) SyncEffectiveFieldsDuringRead(existingState DbfsStorageInfo) {
}

func (c DbfsStorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DbfsStorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DbfsStorageInfo
// only implements ToObjectValue() and Type().
func (o DbfsStorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DbfsStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type DeleteCluster struct {
	// The cluster to be terminated.
	ClusterId types.String `tfsdk:"cluster_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCluster
// only implements ToObjectValue() and Type().
func (o DeleteCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type DeleteClusterResponse struct {
}

func (newState *DeleteClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteClusterResponse) {
}

func (newState *DeleteClusterResponse) SyncEffectiveFieldsDuringRead(existingState DeleteClusterResponse) {
}

func (c DeleteClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteClusterResponse
// only implements ToObjectValue() and Type().
func (o DeleteClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGlobalInitScriptRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteGlobalInitScriptRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGlobalInitScriptRequest
// only implements ToObjectValue() and Type().
func (o DeleteGlobalInitScriptRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": o.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteGlobalInitScriptRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstancePool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteInstancePool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstancePool
// only implements ToObjectValue() and Type().
func (o DeleteInstancePool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type DeleteInstancePoolResponse struct {
}

func (newState *DeleteInstancePoolResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteInstancePoolResponse) {
}

func (newState *DeleteInstancePoolResponse) SyncEffectiveFieldsDuringRead(existingState DeleteInstancePoolResponse) {
}

func (c DeleteInstancePoolResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstancePoolResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteInstancePoolResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstancePoolResponse
// only implements ToObjectValue() and Type().
func (o DeleteInstancePoolResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	PolicyId types.String `tfsdk:"policy_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePolicy
// only implements ToObjectValue() and Type().
func (o DeletePolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type DeletePolicyResponse struct {
}

func (newState *DeletePolicyResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePolicyResponse) {
}

func (newState *DeletePolicyResponse) SyncEffectiveFieldsDuringRead(existingState DeletePolicyResponse) {
}

func (c DeletePolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePolicyResponse
// only implements ToObjectValue() and Type().
func (o DeletePolicyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePolicyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

func (c DeleteResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse
// only implements ToObjectValue() and Type().
func (o DeleteResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DestroyContext struct {
	ClusterId types.String `tfsdk:"clusterId"`

	ContextId types.String `tfsdk:"contextId"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DestroyContext.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DestroyContext) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DestroyContext
// only implements ToObjectValue() and Type().
func (o DestroyContext) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"contextId": o.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DestroyContext) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"contextId": types.StringType,
		},
	}
}

type DestroyResponse struct {
}

func (newState *DestroyResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DestroyResponse) {
}

func (newState *DestroyResponse) SyncEffectiveFieldsDuringRead(existingState DestroyResponse) {
}

func (c DestroyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DestroyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DestroyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DestroyResponse
// only implements ToObjectValue() and Type().
func (o DestroyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DestroyResponse) Type(ctx context.Context) attr.Type {
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

func (newState *DiskSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan DiskSpec) {
}

func (newState *DiskSpec) SyncEffectiveFieldsDuringRead(existingState DiskSpec) {
}

func (c DiskSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DiskSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disk_type": reflect.TypeOf(DiskType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DiskSpec
// only implements ToObjectValue() and Type().
func (o DiskSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disk_count":      o.DiskCount,
			"disk_iops":       o.DiskIops,
			"disk_size":       o.DiskSize,
			"disk_throughput": o.DiskThroughput,
			"disk_type":       o.DiskType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DiskSpec) Type(ctx context.Context) attr.Type {
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
func (o *DiskSpec) GetDiskType(ctx context.Context) (DiskType, bool) {
	var e DiskType
	if o.DiskType.IsNull() || o.DiskType.IsUnknown() {
		return e, false
	}
	var v []DiskType
	d := o.DiskType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDiskType sets the value of the DiskType field in DiskSpec.
func (o *DiskSpec) SetDiskType(ctx context.Context, v DiskType) {
	vs := v.ToObjectValue(ctx)
	o.DiskType = vs
}

// Describes the disk type.
type DiskType struct {
	AzureDiskVolumeType types.String `tfsdk:"azure_disk_volume_type"`

	EbsVolumeType types.String `tfsdk:"ebs_volume_type"`
}

func (newState *DiskType) SyncEffectiveFieldsDuringCreateOrUpdate(plan DiskType) {
}

func (newState *DiskType) SyncEffectiveFieldsDuringRead(existingState DiskType) {
}

func (c DiskType) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DiskType) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DiskType
// only implements ToObjectValue() and Type().
func (o DiskType) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"azure_disk_volume_type": o.AzureDiskVolumeType,
			"ebs_volume_type":        o.EbsVolumeType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DiskType) Type(ctx context.Context) attr.Type {
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

func (newState *DockerBasicAuth) SyncEffectiveFieldsDuringCreateOrUpdate(plan DockerBasicAuth) {
}

func (newState *DockerBasicAuth) SyncEffectiveFieldsDuringRead(existingState DockerBasicAuth) {
}

func (c DockerBasicAuth) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DockerBasicAuth) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DockerBasicAuth
// only implements ToObjectValue() and Type().
func (o DockerBasicAuth) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"password": o.Password,
			"username": o.Username,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DockerBasicAuth) Type(ctx context.Context) attr.Type {
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

func (newState *DockerImage) SyncEffectiveFieldsDuringCreateOrUpdate(plan DockerImage) {
}

func (newState *DockerImage) SyncEffectiveFieldsDuringRead(existingState DockerImage) {
}

func (c DockerImage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DockerImage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"basic_auth": reflect.TypeOf(DockerBasicAuth{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DockerImage
// only implements ToObjectValue() and Type().
func (o DockerImage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"basic_auth": o.BasicAuth,
			"url":        o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DockerImage) Type(ctx context.Context) attr.Type {
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
func (o *DockerImage) GetBasicAuth(ctx context.Context) (DockerBasicAuth, bool) {
	var e DockerBasicAuth
	if o.BasicAuth.IsNull() || o.BasicAuth.IsUnknown() {
		return e, false
	}
	var v []DockerBasicAuth
	d := o.BasicAuth.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBasicAuth sets the value of the BasicAuth field in DockerImage.
func (o *DockerImage) SetBasicAuth(ctx context.Context, v DockerBasicAuth) {
	vs := v.ToObjectValue(ctx)
	o.BasicAuth = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o EditCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_policy_default_values":    o.ApplyPolicyDefaultValues,
			"autoscale":                      o.Autoscale,
			"autotermination_minutes":        o.AutoterminationMinutes,
			"aws_attributes":                 o.AwsAttributes,
			"azure_attributes":               o.AzureAttributes,
			"cluster_id":                     o.ClusterId,
			"cluster_log_conf":               o.ClusterLogConf,
			"cluster_name":                   o.ClusterName,
			"custom_tags":                    o.CustomTags,
			"data_security_mode":             o.DataSecurityMode,
			"docker_image":                   o.DockerImage,
			"driver_instance_pool_id":        o.DriverInstancePoolId,
			"driver_node_type_id":            o.DriverNodeTypeId,
			"enable_elastic_disk":            o.EnableElasticDisk,
			"enable_local_disk_encryption":   o.EnableLocalDiskEncryption,
			"gcp_attributes":                 o.GcpAttributes,
			"init_scripts":                   o.InitScripts,
			"instance_pool_id":               o.InstancePoolId,
			"is_single_node":                 o.IsSingleNode,
			"kind":                           o.Kind,
			"node_type_id":                   o.NodeTypeId,
			"num_workers":                    o.NumWorkers,
			"policy_id":                      o.PolicyId,
			"remote_disk_throughput":         o.RemoteDiskThroughput,
			"runtime_engine":                 o.RuntimeEngine,
			"single_user_name":               o.SingleUserName,
			"spark_conf":                     o.SparkConf,
			"spark_env_vars":                 o.SparkEnvVars,
			"spark_version":                  o.SparkVersion,
			"ssh_public_keys":                o.SshPublicKeys,
			"total_initial_remote_disk_size": o.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 o.UseMlRuntime,
			"workload_type":                  o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EditCluster) Type(ctx context.Context) attr.Type {
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
func (o *EditCluster) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in EditCluster.
func (o *EditCluster) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	o.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in EditCluster as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in EditCluster.
func (o *EditCluster) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in EditCluster as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in EditCluster.
func (o *EditCluster) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in EditCluster as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf
	d := o.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in EditCluster.
func (o *EditCluster) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	o.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in EditCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EditCluster.
func (o *EditCluster) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in EditCluster as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage
	d := o.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in EditCluster.
func (o *EditCluster) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	o.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in EditCluster as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in EditCluster.
func (o *EditCluster) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in EditCluster as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in EditCluster.
func (o *EditCluster) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in EditCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkConf.IsNull() || o.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in EditCluster.
func (o *EditCluster) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in EditCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkEnvVars.IsNull() || o.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in EditCluster.
func (o *EditCluster) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in EditCluster as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if o.SshPublicKeys.IsNull() || o.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in EditCluster.
func (o *EditCluster) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in EditCluster as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType
	d := o.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in EditCluster.
func (o *EditCluster) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	o.WorkloadType = vs
}

type EditClusterResponse struct {
}

func (newState *EditClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditClusterResponse) {
}

func (newState *EditClusterResponse) SyncEffectiveFieldsDuringRead(existingState EditClusterResponse) {
}

func (c EditClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditClusterResponse
// only implements ToObjectValue() and Type().
func (o EditClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditClusterResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditInstancePool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditInstancePool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditInstancePool
// only implements ToObjectValue() and Type().
func (o EditInstancePool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags":                           o.CustomTags,
			"idle_instance_autotermination_minutes": o.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      o.InstancePoolId,
			"instance_pool_name":                    o.InstancePoolName,
			"max_capacity":                          o.MaxCapacity,
			"min_idle_instances":                    o.MinIdleInstances,
			"node_type_id":                          o.NodeTypeId,
			"remote_disk_throughput":                o.RemoteDiskThroughput,
			"total_initial_remote_disk_size":        o.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EditInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_id":                      types.StringType,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
			"node_type_id":                          types.StringType,
			"remote_disk_throughput":                types.Int64Type,
			"total_initial_remote_disk_size":        types.Int64Type,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in EditInstancePool as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditInstancePool) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EditInstancePool.
func (o *EditInstancePool) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

type EditInstancePoolResponse struct {
}

func (newState *EditInstancePoolResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditInstancePoolResponse) {
}

func (newState *EditInstancePoolResponse) SyncEffectiveFieldsDuringRead(existingState EditInstancePoolResponse) {
}

func (c EditInstancePoolResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditInstancePoolResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditInstancePoolResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditInstancePoolResponse
// only implements ToObjectValue() and Type().
func (o EditInstancePoolResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditInstancePoolResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPolicy
// only implements ToObjectValue() and Type().
func (o EditPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"definition":                         o.Definition,
			"description":                        o.Description,
			"libraries":                          o.Libraries,
			"max_clusters_per_user":              o.MaxClustersPerUser,
			"name":                               o.Name,
			"policy_family_definition_overrides": o.PolicyFamilyDefinitionOverrides,
			"policy_family_id":                   o.PolicyFamilyId,
			"policy_id":                          o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EditPolicy) Type(ctx context.Context) attr.Type {
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
func (o *EditPolicy) GetLibraries(ctx context.Context) ([]Library, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in EditPolicy.
func (o *EditPolicy) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type EditPolicyResponse struct {
}

func (newState *EditPolicyResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPolicyResponse) {
}

func (newState *EditPolicyResponse) SyncEffectiveFieldsDuringRead(existingState EditPolicyResponse) {
}

func (c EditPolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditPolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPolicyResponse
// only implements ToObjectValue() and Type().
func (o EditPolicyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditPolicyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditResponse struct {
}

func (newState *EditResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditResponse) {
}

func (newState *EditResponse) SyncEffectiveFieldsDuringRead(existingState EditResponse) {
}

func (c EditResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditResponse
// only implements ToObjectValue() and Type().
func (o EditResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforceClusterComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnforceClusterComplianceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforceClusterComplianceRequest
// only implements ToObjectValue() and Type().
func (o EnforceClusterComplianceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":    o.ClusterId,
			"validate_only": o.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceRequest) Type(ctx context.Context) attr.Type {
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

func (newState *EnforceClusterComplianceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforceClusterComplianceResponse) {
}

func (newState *EnforceClusterComplianceResponse) SyncEffectiveFieldsDuringRead(existingState EnforceClusterComplianceResponse) {
}

func (c EnforceClusterComplianceResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnforceClusterComplianceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(ClusterSettingsChange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforceClusterComplianceResponse
// only implements ToObjectValue() and Type().
func (o EnforceClusterComplianceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes":     o.Changes,
			"has_changes": o.HasChanges,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceResponse) Type(ctx context.Context) attr.Type {
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
func (o *EnforceClusterComplianceResponse) GetChanges(ctx context.Context) ([]ClusterSettingsChange, bool) {
	if o.Changes.IsNull() || o.Changes.IsUnknown() {
		return nil, false
	}
	var v []ClusterSettingsChange
	d := o.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in EnforceClusterComplianceResponse.
func (o *EnforceClusterComplianceResponse) SetChanges(ctx context.Context, v []ClusterSettingsChange) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Changes = types.ListValueMust(t, vs)
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
	// List of jar dependencies, should be string representing volume paths. For
	// example: `/Volumes/path/to/test.jar`.
	JarDependencies types.List `tfsdk:"jar_dependencies"`
}

func (newState *Environment) SyncEffectiveFieldsDuringCreateOrUpdate(plan Environment) {
}

func (newState *Environment) SyncEffectiveFieldsDuringRead(existingState Environment) {
}

func (c Environment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["client"] = attrs["client"].SetOptional()
	attrs["dependencies"] = attrs["dependencies"].SetOptional()
	attrs["environment_version"] = attrs["environment_version"].SetOptional()
	attrs["jar_dependencies"] = attrs["jar_dependencies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Environment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Environment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies":     reflect.TypeOf(types.String{}),
		"jar_dependencies": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Environment
// only implements ToObjectValue() and Type().
func (o Environment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client":              o.Client,
			"dependencies":        o.Dependencies,
			"environment_version": o.EnvironmentVersion,
			"jar_dependencies":    o.JarDependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Environment) Type(ctx context.Context) attr.Type {
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
		},
	}
}

// GetDependencies returns the value of the Dependencies field in Environment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Environment) GetDependencies(ctx context.Context) ([]types.String, bool) {
	if o.Dependencies.IsNull() || o.Dependencies.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Dependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependencies sets the value of the Dependencies field in Environment.
func (o *Environment) SetDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dependencies = types.ListValueMust(t, vs)
}

// GetJarDependencies returns the value of the JarDependencies field in Environment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Environment) GetJarDependencies(ctx context.Context) ([]types.String, bool) {
	if o.JarDependencies.IsNull() || o.JarDependencies.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.JarDependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarDependencies sets the value of the JarDependencies field in Environment.
func (o *Environment) SetJarDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["jar_dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JarDependencies = types.ListValueMust(t, vs)
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

func (newState *EventDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan EventDetails) {
}

func (newState *EventDetails) SyncEffectiveFieldsDuringRead(existingState EventDetails) {
}

func (c EventDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EventDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o EventDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":            o.Attributes,
			"cause":                 o.Cause,
			"cluster_size":          o.ClusterSize,
			"current_num_vcpus":     o.CurrentNumVcpus,
			"current_num_workers":   o.CurrentNumWorkers,
			"did_not_expand_reason": o.DidNotExpandReason,
			"disk_size":             o.DiskSize,
			"driver_state_message":  o.DriverStateMessage,
			"enable_termination_for_node_blocklisted": o.EnableTerminationForNodeBlocklisted,
			"free_space":            o.FreeSpace,
			"init_scripts":          o.InitScripts,
			"instance_id":           o.InstanceId,
			"job_run_name":          o.JobRunName,
			"previous_attributes":   o.PreviousAttributes,
			"previous_cluster_size": o.PreviousClusterSize,
			"previous_disk_size":    o.PreviousDiskSize,
			"reason":                o.Reason,
			"target_num_vcpus":      o.TargetNumVcpus,
			"target_num_workers":    o.TargetNumWorkers,
			"user":                  o.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EventDetails) Type(ctx context.Context) attr.Type {
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
func (o *EventDetails) GetAttributes(ctx context.Context) (ClusterAttributes, bool) {
	var e ClusterAttributes
	if o.Attributes.IsNull() || o.Attributes.IsUnknown() {
		return e, false
	}
	var v []ClusterAttributes
	d := o.Attributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAttributes sets the value of the Attributes field in EventDetails.
func (o *EventDetails) SetAttributes(ctx context.Context, v ClusterAttributes) {
	vs := v.ToObjectValue(ctx)
	o.Attributes = vs
}

// GetClusterSize returns the value of the ClusterSize field in EventDetails as
// a ClusterSize value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails) GetClusterSize(ctx context.Context) (ClusterSize, bool) {
	var e ClusterSize
	if o.ClusterSize.IsNull() || o.ClusterSize.IsUnknown() {
		return e, false
	}
	var v []ClusterSize
	d := o.ClusterSize.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterSize sets the value of the ClusterSize field in EventDetails.
func (o *EventDetails) SetClusterSize(ctx context.Context, v ClusterSize) {
	vs := v.ToObjectValue(ctx)
	o.ClusterSize = vs
}

// GetInitScripts returns the value of the InitScripts field in EventDetails as
// a InitScriptEventDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails) GetInitScripts(ctx context.Context) (InitScriptEventDetails, bool) {
	var e InitScriptEventDetails
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return e, false
	}
	var v []InitScriptEventDetails
	d := o.InitScripts.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitScripts sets the value of the InitScripts field in EventDetails.
func (o *EventDetails) SetInitScripts(ctx context.Context, v InitScriptEventDetails) {
	vs := v.ToObjectValue(ctx)
	o.InitScripts = vs
}

// GetPreviousAttributes returns the value of the PreviousAttributes field in EventDetails as
// a ClusterAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails) GetPreviousAttributes(ctx context.Context) (ClusterAttributes, bool) {
	var e ClusterAttributes
	if o.PreviousAttributes.IsNull() || o.PreviousAttributes.IsUnknown() {
		return e, false
	}
	var v []ClusterAttributes
	d := o.PreviousAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPreviousAttributes sets the value of the PreviousAttributes field in EventDetails.
func (o *EventDetails) SetPreviousAttributes(ctx context.Context, v ClusterAttributes) {
	vs := v.ToObjectValue(ctx)
	o.PreviousAttributes = vs
}

// GetPreviousClusterSize returns the value of the PreviousClusterSize field in EventDetails as
// a ClusterSize value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails) GetPreviousClusterSize(ctx context.Context) (ClusterSize, bool) {
	var e ClusterSize
	if o.PreviousClusterSize.IsNull() || o.PreviousClusterSize.IsUnknown() {
		return e, false
	}
	var v []ClusterSize
	d := o.PreviousClusterSize.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPreviousClusterSize sets the value of the PreviousClusterSize field in EventDetails.
func (o *EventDetails) SetPreviousClusterSize(ctx context.Context, v ClusterSize) {
	vs := v.ToObjectValue(ctx)
	o.PreviousClusterSize = vs
}

// GetReason returns the value of the Reason field in EventDetails as
// a TerminationReason value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails) GetReason(ctx context.Context) (TerminationReason, bool) {
	var e TerminationReason
	if o.Reason.IsNull() || o.Reason.IsUnknown() {
		return e, false
	}
	var v []TerminationReason
	d := o.Reason.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetReason sets the value of the Reason field in EventDetails.
func (o *EventDetails) SetReason(ctx context.Context, v TerminationReason) {
	vs := v.ToObjectValue(ctx)
	o.Reason = vs
}

// Attributes set during cluster creation which are related to GCP.
type GcpAttributes struct {
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs, on-demand VMs, or preemptible VMs with a fallback
	// to on-demand VMs if the former is unavailable.
	Availability types.String `tfsdk:"availability"`
	// Boot disk size in GB
	BootDiskSize types.Int64 `tfsdk:"boot_disk_size"`
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

func (newState *GcpAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpAttributes) {
}

func (newState *GcpAttributes) SyncEffectiveFieldsDuringRead(existingState GcpAttributes) {
}

func (c GcpAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability"] = attrs["availability"].SetOptional()
	attrs["boot_disk_size"] = attrs["boot_disk_size"].SetOptional()
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
func (a GcpAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpAttributes
// only implements ToObjectValue() and Type().
func (o GcpAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":              o.Availability,
			"boot_disk_size":            o.BootDiskSize,
			"google_service_account":    o.GoogleServiceAccount,
			"local_ssd_count":           o.LocalSsdCount,
			"use_preemptible_executors": o.UsePreemptibleExecutors,
			"zone_id":                   o.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":              types.StringType,
			"boot_disk_size":            types.Int64Type,
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

func (newState *GcsStorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcsStorageInfo) {
}

func (newState *GcsStorageInfo) SyncEffectiveFieldsDuringRead(existingState GcsStorageInfo) {
}

func (c GcsStorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GcsStorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcsStorageInfo
// only implements ToObjectValue() and Type().
func (o GcsStorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcsStorageInfo) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetClusterComplianceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterComplianceRequest
// only implements ToObjectValue() and Type().
func (o GetClusterComplianceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterComplianceRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetClusterComplianceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterComplianceResponse) {
}

func (newState *GetClusterComplianceResponse) SyncEffectiveFieldsDuringRead(existingState GetClusterComplianceResponse) {
}

func (c GetClusterComplianceResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetClusterComplianceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterComplianceResponse
// only implements ToObjectValue() and Type().
func (o GetClusterComplianceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_compliant": o.IsCompliant,
			"violations":   o.Violations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterComplianceResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetClusterComplianceResponse) GetViolations(ctx context.Context) (map[string]types.String, bool) {
	if o.Violations.IsNull() || o.Violations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Violations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViolations sets the value of the Violations field in GetClusterComplianceResponse.
func (o *GetClusterComplianceResponse) SetViolations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["violations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Violations = types.MapValueMust(t, vs)
}

type GetClusterPermissionLevelsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetClusterPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetClusterPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetClusterPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPermissionLevelsResponse) {
}

func (newState *GetClusterPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetClusterPermissionLevelsResponse) {
}

func (c GetClusterPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetClusterPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ClusterPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetClusterPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetClusterPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]ClusterPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ClusterPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetClusterPermissionLevelsResponse.
func (o *GetClusterPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []ClusterPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetClusterPermissionsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetClusterPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetClusterPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPolicyPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetClusterPolicyPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetClusterPolicyPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_policy_id": o.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetClusterPolicyPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPolicyPermissionLevelsResponse) {
}

func (newState *GetClusterPolicyPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetClusterPolicyPermissionLevelsResponse) {
}

func (c GetClusterPolicyPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetClusterPolicyPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ClusterPolicyPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetClusterPolicyPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetClusterPolicyPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]ClusterPolicyPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetClusterPolicyPermissionLevelsResponse.
func (o *GetClusterPolicyPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []ClusterPolicyPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetClusterPolicyPermissionsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPolicyPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetClusterPolicyPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetClusterPolicyPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_policy_id": o.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetClusterPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyRequest
// only implements ToObjectValue() and Type().
func (o GetClusterPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetClusterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetClusterRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterRequest
// only implements ToObjectValue() and Type().
func (o GetClusterRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
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

func (newState *GetEvents) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEvents) {
}

func (newState *GetEvents) SyncEffectiveFieldsDuringRead(existingState GetEvents) {
}

func (c GetEvents) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetEvents) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"event_types": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEvents
// only implements ToObjectValue() and Type().
func (o GetEvents) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":  o.ClusterId,
			"end_time":    o.EndTime,
			"event_types": o.EventTypes,
			"limit":       o.Limit,
			"offset":      o.Offset,
			"order":       o.Order,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
			"start_time":  o.StartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEvents) Type(ctx context.Context) attr.Type {
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
func (o *GetEvents) GetEventTypes(ctx context.Context) ([]types.String, bool) {
	if o.EventTypes.IsNull() || o.EventTypes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.EventTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEventTypes sets the value of the EventTypes field in GetEvents.
func (o *GetEvents) SetEventTypes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["event_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EventTypes = types.ListValueMust(t, vs)
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

func (newState *GetEventsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEventsResponse) {
}

func (newState *GetEventsResponse) SyncEffectiveFieldsDuringRead(existingState GetEventsResponse) {
}

func (c GetEventsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetEventsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":    reflect.TypeOf(ClusterEvent{}),
		"next_page": reflect.TypeOf(GetEvents{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEventsResponse
// only implements ToObjectValue() and Type().
func (o GetEventsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"events":          o.Events,
			"next_page":       o.NextPage,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
			"total_count":     o.TotalCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEventsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetEventsResponse) GetEvents(ctx context.Context) ([]ClusterEvent, bool) {
	if o.Events.IsNull() || o.Events.IsUnknown() {
		return nil, false
	}
	var v []ClusterEvent
	d := o.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in GetEventsResponse.
func (o *GetEventsResponse) SetEvents(ctx context.Context, v []ClusterEvent) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// GetNextPage returns the value of the NextPage field in GetEventsResponse as
// a GetEvents value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetEventsResponse) GetNextPage(ctx context.Context) (GetEvents, bool) {
	var e GetEvents
	if o.NextPage.IsNull() || o.NextPage.IsUnknown() {
		return e, false
	}
	var v []GetEvents
	d := o.NextPage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNextPage sets the value of the NextPage field in GetEventsResponse.
func (o *GetEventsResponse) SetNextPage(ctx context.Context, v GetEvents) {
	vs := v.ToObjectValue(ctx)
	o.NextPage = vs
}

type GetGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGlobalInitScriptRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetGlobalInitScriptRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGlobalInitScriptRequest
// only implements ToObjectValue() and Type().
func (o GetGlobalInitScriptRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": o.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetGlobalInitScriptRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetInstancePool) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetInstancePool) {
}

func (newState *GetInstancePool) SyncEffectiveFieldsDuringRead(existingState GetInstancePool) {
}

func (c GetInstancePool) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["default_tags"] = attrs["default_tags"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["idle_instance_autotermination_minutes"] = attrs["idle_instance_autotermination_minutes"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()
	attrs["instance_pool_name"] = attrs["instance_pool_name"].SetOptional()
	attrs["max_capacity"] = attrs["max_capacity"].SetOptional()
	attrs["min_idle_instances"] = attrs["min_idle_instances"].SetOptional()
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
func (a GetInstancePool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"default_tags":             reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
		"stats":                    reflect.TypeOf(InstancePoolStats{}),
		"status":                   reflect.TypeOf(InstancePoolStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePool
// only implements ToObjectValue() and Type().
func (o GetInstancePool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        o.AwsAttributes,
			"azure_attributes":                      o.AzureAttributes,
			"custom_tags":                           o.CustomTags,
			"default_tags":                          o.DefaultTags,
			"disk_spec":                             o.DiskSpec,
			"enable_elastic_disk":                   o.EnableElasticDisk,
			"gcp_attributes":                        o.GcpAttributes,
			"idle_instance_autotermination_minutes": o.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      o.InstancePoolId,
			"instance_pool_name":                    o.InstancePoolName,
			"max_capacity":                          o.MaxCapacity,
			"min_idle_instances":                    o.MinIdleInstances,
			"node_type_id":                          o.NodeTypeId,
			"preloaded_docker_images":               o.PreloadedDockerImages,
			"preloaded_spark_versions":              o.PreloadedSparkVersions,
			"remote_disk_throughput":                o.RemoteDiskThroughput,
			"state":                                 o.State,
			"stats":                                 o.Stats,
			"status":                                o.Status,
			"total_initial_remote_disk_size":        o.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePool) Type(ctx context.Context) attr.Type {
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
			"enable_elastic_disk":                   types.BoolType,
			"gcp_attributes":                        InstancePoolGcpAttributes{}.Type(ctx),
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_id":                      types.StringType,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
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
func (o *GetInstancePool) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes, bool) {
	var e InstancePoolAwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in GetInstancePool.
func (o *GetInstancePool) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in GetInstancePool as
// a InstancePoolAzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes, bool) {
	var e InstancePoolAzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in GetInstancePool.
func (o *GetInstancePool) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetCustomTags returns the value of the CustomTags field in GetInstancePool as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in GetInstancePool.
func (o *GetInstancePool) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in GetInstancePool as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
	if o.DefaultTags.IsNull() || o.DefaultTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.DefaultTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultTags sets the value of the DefaultTags field in GetInstancePool.
func (o *GetInstancePool) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DefaultTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in GetInstancePool as
// a DiskSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetDiskSpec(ctx context.Context) (DiskSpec, bool) {
	var e DiskSpec
	if o.DiskSpec.IsNull() || o.DiskSpec.IsUnknown() {
		return e, false
	}
	var v []DiskSpec
	d := o.DiskSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDiskSpec sets the value of the DiskSpec field in GetInstancePool.
func (o *GetInstancePool) SetDiskSpec(ctx context.Context, v DiskSpec) {
	vs := v.ToObjectValue(ctx)
	o.DiskSpec = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in GetInstancePool as
// a InstancePoolGcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes, bool) {
	var e InstancePoolGcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolGcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in GetInstancePool.
func (o *GetInstancePool) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in GetInstancePool as
// a slice of DockerImage values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage, bool) {
	if o.PreloadedDockerImages.IsNull() || o.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage
	d := o.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in GetInstancePool.
func (o *GetInstancePool) SetPreloadedDockerImages(ctx context.Context, v []DockerImage) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in GetInstancePool as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
	if o.PreloadedSparkVersions.IsNull() || o.PreloadedSparkVersions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PreloadedSparkVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in GetInstancePool.
func (o *GetInstancePool) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

// GetStats returns the value of the Stats field in GetInstancePool as
// a InstancePoolStats value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetStats(ctx context.Context) (InstancePoolStats, bool) {
	var e InstancePoolStats
	if o.Stats.IsNull() || o.Stats.IsUnknown() {
		return e, false
	}
	var v []InstancePoolStats
	d := o.Stats.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStats sets the value of the Stats field in GetInstancePool.
func (o *GetInstancePool) SetStats(ctx context.Context, v InstancePoolStats) {
	vs := v.ToObjectValue(ctx)
	o.Stats = vs
}

// GetStatus returns the value of the Status field in GetInstancePool as
// a InstancePoolStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool) GetStatus(ctx context.Context) (InstancePoolStatus, bool) {
	var e InstancePoolStatus
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []InstancePoolStatus
	d := o.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in GetInstancePool.
func (o *GetInstancePool) SetStatus(ctx context.Context, v InstancePoolStatus) {
	vs := v.ToObjectValue(ctx)
	o.Status = vs
}

type GetInstancePoolPermissionLevelsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetInstancePoolPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetInstancePoolPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetInstancePoolPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetInstancePoolPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetInstancePoolPermissionLevelsResponse) {
}

func (newState *GetInstancePoolPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetInstancePoolPermissionLevelsResponse) {
}

func (c GetInstancePoolPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetInstancePoolPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(InstancePoolPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetInstancePoolPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetInstancePoolPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]InstancePoolPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetInstancePoolPermissionLevelsResponse.
func (o *GetInstancePoolPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []InstancePoolPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetInstancePoolPermissionsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetInstancePoolPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetInstancePoolPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetInstancePoolPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetInstancePoolRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetInstancePoolRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolRequest
// only implements ToObjectValue() and Type().
func (o GetInstancePoolRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPolicyFamilyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPolicyFamilyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPolicyFamilyRequest
// only implements ToObjectValue() and Type().
func (o GetPolicyFamilyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_family_id": o.PolicyFamilyId,
			"version":          o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPolicyFamilyRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetSparkVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSparkVersionsResponse) {
}

func (newState *GetSparkVersionsResponse) SyncEffectiveFieldsDuringRead(existingState GetSparkVersionsResponse) {
}

func (c GetSparkVersionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetSparkVersionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"versions": reflect.TypeOf(SparkVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSparkVersionsResponse
// only implements ToObjectValue() and Type().
func (o GetSparkVersionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"versions": o.Versions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSparkVersionsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetSparkVersionsResponse) GetVersions(ctx context.Context) ([]SparkVersion, bool) {
	if o.Versions.IsNull() || o.Versions.IsUnknown() {
		return nil, false
	}
	var v []SparkVersion
	d := o.Versions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVersions sets the value of the Versions field in GetSparkVersionsResponse.
func (o *GetSparkVersionsResponse) SetVersions(ctx context.Context, v []SparkVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Versions = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GlobalInitScriptCreateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GlobalInitScriptCreateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptCreateRequest
// only implements ToObjectValue() and Type().
func (o GlobalInitScriptCreateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":  o.Enabled,
			"name":     o.Name,
			"position": o.Position,
			"script":   o.Script,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GlobalInitScriptCreateRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GlobalInitScriptDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan GlobalInitScriptDetails) {
}

func (newState *GlobalInitScriptDetails) SyncEffectiveFieldsDuringRead(existingState GlobalInitScriptDetails) {
}

func (c GlobalInitScriptDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GlobalInitScriptDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptDetails
// only implements ToObjectValue() and Type().
func (o GlobalInitScriptDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at": o.CreatedAt,
			"created_by": o.CreatedBy,
			"enabled":    o.Enabled,
			"name":       o.Name,
			"position":   o.Position,
			"script_id":  o.ScriptId,
			"updated_at": o.UpdatedAt,
			"updated_by": o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetails) Type(ctx context.Context) attr.Type {
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

func (newState *GlobalInitScriptDetailsWithContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan GlobalInitScriptDetailsWithContent) {
}

func (newState *GlobalInitScriptDetailsWithContent) SyncEffectiveFieldsDuringRead(existingState GlobalInitScriptDetailsWithContent) {
}

func (c GlobalInitScriptDetailsWithContent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GlobalInitScriptDetailsWithContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptDetailsWithContent
// only implements ToObjectValue() and Type().
func (o GlobalInitScriptDetailsWithContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at": o.CreatedAt,
			"created_by": o.CreatedBy,
			"enabled":    o.Enabled,
			"name":       o.Name,
			"position":   o.Position,
			"script":     o.Script,
			"script_id":  o.ScriptId,
			"updated_at": o.UpdatedAt,
			"updated_by": o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetailsWithContent) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GlobalInitScriptUpdateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GlobalInitScriptUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptUpdateRequest
// only implements ToObjectValue() and Type().
func (o GlobalInitScriptUpdateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":   o.Enabled,
			"name":      o.Name,
			"position":  o.Position,
			"script":    o.Script,
			"script_id": o.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GlobalInitScriptUpdateRequest) Type(ctx context.Context) attr.Type {
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

func (newState *InitScriptEventDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan InitScriptEventDetails) {
}

func (newState *InitScriptEventDetails) SyncEffectiveFieldsDuringRead(existingState InitScriptEventDetails) {
}

func (c InitScriptEventDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InitScriptEventDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster": reflect.TypeOf(InitScriptInfoAndExecutionDetails{}),
		"global":  reflect.TypeOf(InitScriptInfoAndExecutionDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InitScriptEventDetails
// only implements ToObjectValue() and Type().
func (o InitScriptEventDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster":           o.Cluster,
			"global":            o.Global,
			"reported_for_node": o.ReportedForNode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InitScriptEventDetails) Type(ctx context.Context) attr.Type {
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
func (o *InitScriptEventDetails) GetCluster(ctx context.Context) ([]InitScriptInfoAndExecutionDetails, bool) {
	if o.Cluster.IsNull() || o.Cluster.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfoAndExecutionDetails
	d := o.Cluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCluster sets the value of the Cluster field in InitScriptEventDetails.
func (o *InitScriptEventDetails) SetCluster(ctx context.Context, v []InitScriptInfoAndExecutionDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Cluster = types.ListValueMust(t, vs)
}

// GetGlobal returns the value of the Global field in InitScriptEventDetails as
// a slice of InitScriptInfoAndExecutionDetails values.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptEventDetails) GetGlobal(ctx context.Context) ([]InitScriptInfoAndExecutionDetails, bool) {
	if o.Global.IsNull() || o.Global.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfoAndExecutionDetails
	d := o.Global.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGlobal sets the value of the Global field in InitScriptEventDetails.
func (o *InitScriptEventDetails) SetGlobal(ctx context.Context, v []InitScriptInfoAndExecutionDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["global"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Global = types.ListValueMust(t, vs)
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

func (newState *InitScriptInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan InitScriptInfo) {
}

func (newState *InitScriptInfo) SyncEffectiveFieldsDuringRead(existingState InitScriptInfo) {
}

func (c InitScriptInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InitScriptInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o InitScriptInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"abfss":     o.Abfss,
			"dbfs":      o.Dbfs,
			"file":      o.File,
			"gcs":       o.Gcs,
			"s3":        o.S3,
			"volumes":   o.Volumes,
			"workspace": o.Workspace,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InitScriptInfo) Type(ctx context.Context) attr.Type {
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
func (o *InitScriptInfo) GetAbfss(ctx context.Context) (Adlsgen2Info, bool) {
	var e Adlsgen2Info
	if o.Abfss.IsNull() || o.Abfss.IsUnknown() {
		return e, false
	}
	var v []Adlsgen2Info
	d := o.Abfss.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAbfss sets the value of the Abfss field in InitScriptInfo.
func (o *InitScriptInfo) SetAbfss(ctx context.Context, v Adlsgen2Info) {
	vs := v.ToObjectValue(ctx)
	o.Abfss = vs
}

// GetDbfs returns the value of the Dbfs field in InitScriptInfo as
// a DbfsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo) GetDbfs(ctx context.Context) (DbfsStorageInfo, bool) {
	var e DbfsStorageInfo
	if o.Dbfs.IsNull() || o.Dbfs.IsUnknown() {
		return e, false
	}
	var v []DbfsStorageInfo
	d := o.Dbfs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbfs sets the value of the Dbfs field in InitScriptInfo.
func (o *InitScriptInfo) SetDbfs(ctx context.Context, v DbfsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Dbfs = vs
}

// GetFile returns the value of the File field in InitScriptInfo as
// a LocalFileInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo) GetFile(ctx context.Context) (LocalFileInfo, bool) {
	var e LocalFileInfo
	if o.File.IsNull() || o.File.IsUnknown() {
		return e, false
	}
	var v []LocalFileInfo
	d := o.File.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFile sets the value of the File field in InitScriptInfo.
func (o *InitScriptInfo) SetFile(ctx context.Context, v LocalFileInfo) {
	vs := v.ToObjectValue(ctx)
	o.File = vs
}

// GetGcs returns the value of the Gcs field in InitScriptInfo as
// a GcsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo) GetGcs(ctx context.Context) (GcsStorageInfo, bool) {
	var e GcsStorageInfo
	if o.Gcs.IsNull() || o.Gcs.IsUnknown() {
		return e, false
	}
	var v []GcsStorageInfo
	d := o.Gcs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcs sets the value of the Gcs field in InitScriptInfo.
func (o *InitScriptInfo) SetGcs(ctx context.Context, v GcsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Gcs = vs
}

// GetS3 returns the value of the S3 field in InitScriptInfo as
// a S3StorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo) GetS3(ctx context.Context) (S3StorageInfo, bool) {
	var e S3StorageInfo
	if o.S3.IsNull() || o.S3.IsUnknown() {
		return e, false
	}
	var v []S3StorageInfo
	d := o.S3.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetS3 sets the value of the S3 field in InitScriptInfo.
func (o *InitScriptInfo) SetS3(ctx context.Context, v S3StorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.S3 = vs
}

// GetVolumes returns the value of the Volumes field in InitScriptInfo as
// a VolumesStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo) GetVolumes(ctx context.Context) (VolumesStorageInfo, bool) {
	var e VolumesStorageInfo
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return e, false
	}
	var v []VolumesStorageInfo
	d := o.Volumes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumes sets the value of the Volumes field in InitScriptInfo.
func (o *InitScriptInfo) SetVolumes(ctx context.Context, v VolumesStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Volumes = vs
}

// GetWorkspace returns the value of the Workspace field in InitScriptInfo as
// a WorkspaceStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo) GetWorkspace(ctx context.Context) (WorkspaceStorageInfo, bool) {
	var e WorkspaceStorageInfo
	if o.Workspace.IsNull() || o.Workspace.IsUnknown() {
		return e, false
	}
	var v []WorkspaceStorageInfo
	d := o.Workspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspace sets the value of the Workspace field in InitScriptInfo.
func (o *InitScriptInfo) SetWorkspace(ctx context.Context, v WorkspaceStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Workspace = vs
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

func (newState *InitScriptInfoAndExecutionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan InitScriptInfoAndExecutionDetails) {
}

func (newState *InitScriptInfoAndExecutionDetails) SyncEffectiveFieldsDuringRead(existingState InitScriptInfoAndExecutionDetails) {
}

func (c InitScriptInfoAndExecutionDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InitScriptInfoAndExecutionDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o InitScriptInfoAndExecutionDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"abfss":                      o.Abfss,
			"dbfs":                       o.Dbfs,
			"error_message":              o.ErrorMessage,
			"execution_duration_seconds": o.ExecutionDurationSeconds,
			"file":                       o.File,
			"gcs":                        o.Gcs,
			"s3":                         o.S3,
			"status":                     o.Status,
			"volumes":                    o.Volumes,
			"workspace":                  o.Workspace,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InitScriptInfoAndExecutionDetails) Type(ctx context.Context) attr.Type {
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
func (o *InitScriptInfoAndExecutionDetails) GetAbfss(ctx context.Context) (Adlsgen2Info, bool) {
	var e Adlsgen2Info
	if o.Abfss.IsNull() || o.Abfss.IsUnknown() {
		return e, false
	}
	var v []Adlsgen2Info
	d := o.Abfss.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAbfss sets the value of the Abfss field in InitScriptInfoAndExecutionDetails.
func (o *InitScriptInfoAndExecutionDetails) SetAbfss(ctx context.Context, v Adlsgen2Info) {
	vs := v.ToObjectValue(ctx)
	o.Abfss = vs
}

// GetDbfs returns the value of the Dbfs field in InitScriptInfoAndExecutionDetails as
// a DbfsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails) GetDbfs(ctx context.Context) (DbfsStorageInfo, bool) {
	var e DbfsStorageInfo
	if o.Dbfs.IsNull() || o.Dbfs.IsUnknown() {
		return e, false
	}
	var v []DbfsStorageInfo
	d := o.Dbfs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbfs sets the value of the Dbfs field in InitScriptInfoAndExecutionDetails.
func (o *InitScriptInfoAndExecutionDetails) SetDbfs(ctx context.Context, v DbfsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Dbfs = vs
}

// GetFile returns the value of the File field in InitScriptInfoAndExecutionDetails as
// a LocalFileInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails) GetFile(ctx context.Context) (LocalFileInfo, bool) {
	var e LocalFileInfo
	if o.File.IsNull() || o.File.IsUnknown() {
		return e, false
	}
	var v []LocalFileInfo
	d := o.File.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFile sets the value of the File field in InitScriptInfoAndExecutionDetails.
func (o *InitScriptInfoAndExecutionDetails) SetFile(ctx context.Context, v LocalFileInfo) {
	vs := v.ToObjectValue(ctx)
	o.File = vs
}

// GetGcs returns the value of the Gcs field in InitScriptInfoAndExecutionDetails as
// a GcsStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails) GetGcs(ctx context.Context) (GcsStorageInfo, bool) {
	var e GcsStorageInfo
	if o.Gcs.IsNull() || o.Gcs.IsUnknown() {
		return e, false
	}
	var v []GcsStorageInfo
	d := o.Gcs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcs sets the value of the Gcs field in InitScriptInfoAndExecutionDetails.
func (o *InitScriptInfoAndExecutionDetails) SetGcs(ctx context.Context, v GcsStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Gcs = vs
}

// GetS3 returns the value of the S3 field in InitScriptInfoAndExecutionDetails as
// a S3StorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails) GetS3(ctx context.Context) (S3StorageInfo, bool) {
	var e S3StorageInfo
	if o.S3.IsNull() || o.S3.IsUnknown() {
		return e, false
	}
	var v []S3StorageInfo
	d := o.S3.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetS3 sets the value of the S3 field in InitScriptInfoAndExecutionDetails.
func (o *InitScriptInfoAndExecutionDetails) SetS3(ctx context.Context, v S3StorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.S3 = vs
}

// GetVolumes returns the value of the Volumes field in InitScriptInfoAndExecutionDetails as
// a VolumesStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails) GetVolumes(ctx context.Context) (VolumesStorageInfo, bool) {
	var e VolumesStorageInfo
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return e, false
	}
	var v []VolumesStorageInfo
	d := o.Volumes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumes sets the value of the Volumes field in InitScriptInfoAndExecutionDetails.
func (o *InitScriptInfoAndExecutionDetails) SetVolumes(ctx context.Context, v VolumesStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Volumes = vs
}

// GetWorkspace returns the value of the Workspace field in InitScriptInfoAndExecutionDetails as
// a WorkspaceStorageInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails) GetWorkspace(ctx context.Context) (WorkspaceStorageInfo, bool) {
	var e WorkspaceStorageInfo
	if o.Workspace.IsNull() || o.Workspace.IsUnknown() {
		return e, false
	}
	var v []WorkspaceStorageInfo
	d := o.Workspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspace sets the value of the Workspace field in InitScriptInfoAndExecutionDetails.
func (o *InitScriptInfoAndExecutionDetails) SetWorkspace(ctx context.Context, v WorkspaceStorageInfo) {
	vs := v.ToObjectValue(ctx)
	o.Workspace = vs
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The libraries to install.
	Libraries types.List `tfsdk:"libraries"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstallLibraries.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InstallLibraries) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallLibraries
// only implements ToObjectValue() and Type().
func (o InstallLibraries) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
			"libraries":  o.Libraries,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstallLibraries) Type(ctx context.Context) attr.Type {
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
func (o *InstallLibraries) GetLibraries(ctx context.Context) ([]Library, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in InstallLibraries.
func (o *InstallLibraries) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type InstallLibrariesResponse struct {
}

func (newState *InstallLibrariesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstallLibrariesResponse) {
}

func (newState *InstallLibrariesResponse) SyncEffectiveFieldsDuringRead(existingState InstallLibrariesResponse) {
}

func (c InstallLibrariesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstallLibrariesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InstallLibrariesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallLibrariesResponse
// only implements ToObjectValue() and Type().
func (o InstallLibrariesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse) Type(ctx context.Context) attr.Type {
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

func (newState *InstancePoolAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAccessControlRequest) {
}

func (newState *InstancePoolAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState InstancePoolAccessControlRequest) {
}

func (c InstancePoolAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAccessControlRequest
// only implements ToObjectValue() and Type().
func (o InstancePoolAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (newState *InstancePoolAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAccessControlResponse) {
}

func (newState *InstancePoolAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState InstancePoolAccessControlResponse) {
}

func (c InstancePoolAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(InstancePoolPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAccessControlResponse
// only implements ToObjectValue() and Type().
func (o InstancePoolAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (o *InstancePoolAccessControlResponse) GetAllPermissions(ctx context.Context) ([]InstancePoolPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in InstancePoolAccessControlResponse.
func (o *InstancePoolAccessControlResponse) SetAllPermissions(ctx context.Context, v []InstancePoolPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
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

func (newState *InstancePoolAndStats) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAndStats) {
}

func (newState *InstancePoolAndStats) SyncEffectiveFieldsDuringRead(existingState InstancePoolAndStats) {
}

func (c InstancePoolAndStats) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["default_tags"] = attrs["default_tags"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["idle_instance_autotermination_minutes"] = attrs["idle_instance_autotermination_minutes"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["instance_pool_name"] = attrs["instance_pool_name"].SetOptional()
	attrs["max_capacity"] = attrs["max_capacity"].SetOptional()
	attrs["min_idle_instances"] = attrs["min_idle_instances"].SetOptional()
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
func (a InstancePoolAndStats) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"default_tags":             reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
		"stats":                    reflect.TypeOf(InstancePoolStats{}),
		"status":                   reflect.TypeOf(InstancePoolStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAndStats
// only implements ToObjectValue() and Type().
func (o InstancePoolAndStats) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        o.AwsAttributes,
			"azure_attributes":                      o.AzureAttributes,
			"custom_tags":                           o.CustomTags,
			"default_tags":                          o.DefaultTags,
			"disk_spec":                             o.DiskSpec,
			"enable_elastic_disk":                   o.EnableElasticDisk,
			"gcp_attributes":                        o.GcpAttributes,
			"idle_instance_autotermination_minutes": o.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      o.InstancePoolId,
			"instance_pool_name":                    o.InstancePoolName,
			"max_capacity":                          o.MaxCapacity,
			"min_idle_instances":                    o.MinIdleInstances,
			"node_type_id":                          o.NodeTypeId,
			"preloaded_docker_images":               o.PreloadedDockerImages,
			"preloaded_spark_versions":              o.PreloadedSparkVersions,
			"remote_disk_throughput":                o.RemoteDiskThroughput,
			"state":                                 o.State,
			"stats":                                 o.Stats,
			"status":                                o.Status,
			"total_initial_remote_disk_size":        o.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolAndStats) Type(ctx context.Context) attr.Type {
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
			"enable_elastic_disk":                   types.BoolType,
			"gcp_attributes":                        InstancePoolGcpAttributes{}.Type(ctx),
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_id":                      types.StringType,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
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
func (o *InstancePoolAndStats) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes, bool) {
	var e InstancePoolAwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in InstancePoolAndStats as
// a InstancePoolAzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes, bool) {
	var e InstancePoolAzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetCustomTags returns the value of the CustomTags field in InstancePoolAndStats as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in InstancePoolAndStats as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
	if o.DefaultTags.IsNull() || o.DefaultTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.DefaultTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultTags sets the value of the DefaultTags field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DefaultTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in InstancePoolAndStats as
// a DiskSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetDiskSpec(ctx context.Context) (DiskSpec, bool) {
	var e DiskSpec
	if o.DiskSpec.IsNull() || o.DiskSpec.IsUnknown() {
		return e, false
	}
	var v []DiskSpec
	d := o.DiskSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDiskSpec sets the value of the DiskSpec field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetDiskSpec(ctx context.Context, v DiskSpec) {
	vs := v.ToObjectValue(ctx)
	o.DiskSpec = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in InstancePoolAndStats as
// a InstancePoolGcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes, bool) {
	var e InstancePoolGcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolGcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in InstancePoolAndStats as
// a slice of DockerImage values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage, bool) {
	if o.PreloadedDockerImages.IsNull() || o.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage
	d := o.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetPreloadedDockerImages(ctx context.Context, v []DockerImage) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in InstancePoolAndStats as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
	if o.PreloadedSparkVersions.IsNull() || o.PreloadedSparkVersions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PreloadedSparkVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

// GetStats returns the value of the Stats field in InstancePoolAndStats as
// a InstancePoolStats value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetStats(ctx context.Context) (InstancePoolStats, bool) {
	var e InstancePoolStats
	if o.Stats.IsNull() || o.Stats.IsUnknown() {
		return e, false
	}
	var v []InstancePoolStats
	d := o.Stats.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStats sets the value of the Stats field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetStats(ctx context.Context, v InstancePoolStats) {
	vs := v.ToObjectValue(ctx)
	o.Stats = vs
}

// GetStatus returns the value of the Status field in InstancePoolAndStats as
// a InstancePoolStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats) GetStatus(ctx context.Context) (InstancePoolStatus, bool) {
	var e InstancePoolStatus
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []InstancePoolStatus
	d := o.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in InstancePoolAndStats.
func (o *InstancePoolAndStats) SetStatus(ctx context.Context, v InstancePoolStatus) {
	vs := v.ToObjectValue(ctx)
	o.Status = vs
}

// Attributes set during instance pool creation which are related to Amazon Web
// Services.
type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	Availability types.String `tfsdk:"availability"`
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

func (newState *InstancePoolAwsAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAwsAttributes) {
}

func (newState *InstancePoolAwsAttributes) SyncEffectiveFieldsDuringRead(existingState InstancePoolAwsAttributes) {
}

func (c InstancePoolAwsAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability"] = attrs["availability"].SetOptional()
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
func (a InstancePoolAwsAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAwsAttributes
// only implements ToObjectValue() and Type().
func (o InstancePoolAwsAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":           o.Availability,
			"spot_bid_price_percent": o.SpotBidPricePercent,
			"zone_id":                o.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolAwsAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":           types.StringType,
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

func (newState *InstancePoolAzureAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAzureAttributes) {
}

func (newState *InstancePoolAzureAttributes) SyncEffectiveFieldsDuringRead(existingState InstancePoolAzureAttributes) {
}

func (c InstancePoolAzureAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolAzureAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAzureAttributes
// only implements ToObjectValue() and Type().
func (o InstancePoolAzureAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":       o.Availability,
			"spot_bid_max_price": o.SpotBidMaxPrice,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolAzureAttributes) Type(ctx context.Context) attr.Type {
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

func (newState *InstancePoolGcpAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolGcpAttributes) {
}

func (newState *InstancePoolGcpAttributes) SyncEffectiveFieldsDuringRead(existingState InstancePoolGcpAttributes) {
}

func (c InstancePoolGcpAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolGcpAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolGcpAttributes
// only implements ToObjectValue() and Type().
func (o InstancePoolGcpAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_availability": o.GcpAvailability,
			"local_ssd_count":  o.LocalSsdCount,
			"zone_id":          o.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolGcpAttributes) Type(ctx context.Context) attr.Type {
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

func (newState *InstancePoolPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolPermission) {
}

func (newState *InstancePoolPermission) SyncEffectiveFieldsDuringRead(existingState InstancePoolPermission) {
}

func (c InstancePoolPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermission
// only implements ToObjectValue() and Type().
func (o InstancePoolPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolPermission) Type(ctx context.Context) attr.Type {
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
func (o *InstancePoolPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in InstancePoolPermission.
func (o *InstancePoolPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type InstancePoolPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *InstancePoolPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolPermissions) {
}

func (newState *InstancePoolPermissions) SyncEffectiveFieldsDuringRead(existingState InstancePoolPermissions) {
}

func (c InstancePoolPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(InstancePoolAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissions
// only implements ToObjectValue() and Type().
func (o InstancePoolPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolPermissions) Type(ctx context.Context) attr.Type {
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
func (o *InstancePoolPermissions) GetAccessControlList(ctx context.Context) ([]InstancePoolAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in InstancePoolPermissions.
func (o *InstancePoolPermissions) SetAccessControlList(ctx context.Context, v []InstancePoolAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type InstancePoolPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *InstancePoolPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolPermissionsDescription) {
}

func (newState *InstancePoolPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState InstancePoolPermissionsDescription) {
}

func (c InstancePoolPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissionsDescription
// only implements ToObjectValue() and Type().
func (o InstancePoolPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsDescription) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstancePoolPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InstancePoolPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(InstancePoolAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissionsRequest
// only implements ToObjectValue() and Type().
func (o InstancePoolPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"instance_pool_id":    o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *InstancePoolPermissionsRequest) GetAccessControlList(ctx context.Context) ([]InstancePoolAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in InstancePoolPermissionsRequest.
func (o *InstancePoolPermissionsRequest) SetAccessControlList(ctx context.Context, v []InstancePoolAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
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

func (newState *InstancePoolStats) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolStats) {
}

func (newState *InstancePoolStats) SyncEffectiveFieldsDuringRead(existingState InstancePoolStats) {
}

func (c InstancePoolStats) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolStats) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolStats
// only implements ToObjectValue() and Type().
func (o InstancePoolStats) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"idle_count":         o.IdleCount,
			"pending_idle_count": o.PendingIdleCount,
			"pending_used_count": o.PendingUsedCount,
			"used_count":         o.UsedCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolStats) Type(ctx context.Context) attr.Type {
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

func (newState *InstancePoolStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolStatus) {
}

func (newState *InstancePoolStatus) SyncEffectiveFieldsDuringRead(existingState InstancePoolStatus) {
}

func (c InstancePoolStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pending_instance_errors": reflect.TypeOf(PendingInstanceError{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolStatus
// only implements ToObjectValue() and Type().
func (o InstancePoolStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pending_instance_errors": o.PendingInstanceErrors,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolStatus) Type(ctx context.Context) attr.Type {
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
func (o *InstancePoolStatus) GetPendingInstanceErrors(ctx context.Context) ([]PendingInstanceError, bool) {
	if o.PendingInstanceErrors.IsNull() || o.PendingInstanceErrors.IsUnknown() {
		return nil, false
	}
	var v []PendingInstanceError
	d := o.PendingInstanceErrors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPendingInstanceErrors sets the value of the PendingInstanceErrors field in InstancePoolStatus.
func (o *InstancePoolStatus) SetPendingInstanceErrors(ctx context.Context, v []PendingInstanceError) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pending_instance_errors"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PendingInstanceErrors = types.ListValueMust(t, vs)
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

func (newState *InstanceProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstanceProfile) {
}

func (newState *InstanceProfile) SyncEffectiveFieldsDuringRead(existingState InstanceProfile) {
}

func (c InstanceProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstanceProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstanceProfile
// only implements ToObjectValue() and Type().
func (o InstanceProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"iam_role_arn":             o.IamRoleArn,
			"instance_profile_arn":     o.InstanceProfileArn,
			"is_meta_instance_profile": o.IsMetaInstanceProfile,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstanceProfile) Type(ctx context.Context) attr.Type {
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

func (newState *Library) SyncEffectiveFieldsDuringCreateOrUpdate(plan Library) {
}

func (newState *Library) SyncEffectiveFieldsDuringRead(existingState Library) {
}

func (c Library) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Library) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cran":  reflect.TypeOf(RCranLibrary{}),
		"maven": reflect.TypeOf(MavenLibrary{}),
		"pypi":  reflect.TypeOf(PythonPyPiLibrary{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Library
// only implements ToObjectValue() and Type().
func (o Library) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cran":         o.Cran,
			"egg":          o.Egg,
			"jar":          o.Jar,
			"maven":        o.Maven,
			"pypi":         o.Pypi,
			"requirements": o.Requirements,
			"whl":          o.Whl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Library) Type(ctx context.Context) attr.Type {
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
func (o *Library) GetCran(ctx context.Context) (RCranLibrary, bool) {
	var e RCranLibrary
	if o.Cran.IsNull() || o.Cran.IsUnknown() {
		return e, false
	}
	var v []RCranLibrary
	d := o.Cran.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCran sets the value of the Cran field in Library.
func (o *Library) SetCran(ctx context.Context, v RCranLibrary) {
	vs := v.ToObjectValue(ctx)
	o.Cran = vs
}

// GetMaven returns the value of the Maven field in Library as
// a MavenLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (o *Library) GetMaven(ctx context.Context) (MavenLibrary, bool) {
	var e MavenLibrary
	if o.Maven.IsNull() || o.Maven.IsUnknown() {
		return e, false
	}
	var v []MavenLibrary
	d := o.Maven.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMaven sets the value of the Maven field in Library.
func (o *Library) SetMaven(ctx context.Context, v MavenLibrary) {
	vs := v.ToObjectValue(ctx)
	o.Maven = vs
}

// GetPypi returns the value of the Pypi field in Library as
// a PythonPyPiLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (o *Library) GetPypi(ctx context.Context) (PythonPyPiLibrary, bool) {
	var e PythonPyPiLibrary
	if o.Pypi.IsNull() || o.Pypi.IsUnknown() {
		return e, false
	}
	var v []PythonPyPiLibrary
	d := o.Pypi.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPypi sets the value of the Pypi field in Library.
func (o *Library) SetPypi(ctx context.Context, v PythonPyPiLibrary) {
	vs := v.ToObjectValue(ctx)
	o.Pypi = vs
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

func (newState *LibraryFullStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan LibraryFullStatus) {
}

func (newState *LibraryFullStatus) SyncEffectiveFieldsDuringRead(existingState LibraryFullStatus) {
}

func (c LibraryFullStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LibraryFullStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"library":  reflect.TypeOf(Library{}),
		"messages": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LibraryFullStatus
// only implements ToObjectValue() and Type().
func (o LibraryFullStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_library_for_all_clusters": o.IsLibraryForAllClusters,
			"library":                     o.Library,
			"messages":                    o.Messages,
			"status":                      o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LibraryFullStatus) Type(ctx context.Context) attr.Type {
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
func (o *LibraryFullStatus) GetLibrary(ctx context.Context) (Library, bool) {
	var e Library
	if o.Library.IsNull() || o.Library.IsUnknown() {
		return e, false
	}
	var v []Library
	d := o.Library.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLibrary sets the value of the Library field in LibraryFullStatus.
func (o *LibraryFullStatus) SetLibrary(ctx context.Context, v Library) {
	vs := v.ToObjectValue(ctx)
	o.Library = vs
}

// GetMessages returns the value of the Messages field in LibraryFullStatus as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *LibraryFullStatus) GetMessages(ctx context.Context) ([]types.String, bool) {
	if o.Messages.IsNull() || o.Messages.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Messages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMessages sets the value of the Messages field in LibraryFullStatus.
func (o *LibraryFullStatus) SetMessages(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Messages = types.ListValueMust(t, vs)
}

type ListAllClusterLibraryStatuses struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllClusterLibraryStatuses.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllClusterLibraryStatuses) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllClusterLibraryStatuses
// only implements ToObjectValue() and Type().
func (o ListAllClusterLibraryStatuses) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatuses) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListAllClusterLibraryStatusesResponse struct {
	// A list of cluster statuses.
	Statuses types.List `tfsdk:"statuses"`
}

func (newState *ListAllClusterLibraryStatusesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllClusterLibraryStatusesResponse) {
}

func (newState *ListAllClusterLibraryStatusesResponse) SyncEffectiveFieldsDuringRead(existingState ListAllClusterLibraryStatusesResponse) {
}

func (c ListAllClusterLibraryStatusesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAllClusterLibraryStatusesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statuses": reflect.TypeOf(ClusterLibraryStatuses{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllClusterLibraryStatusesResponse
// only implements ToObjectValue() and Type().
func (o ListAllClusterLibraryStatusesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statuses": o.Statuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatusesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListAllClusterLibraryStatusesResponse) GetStatuses(ctx context.Context) ([]ClusterLibraryStatuses, bool) {
	if o.Statuses.IsNull() || o.Statuses.IsUnknown() {
		return nil, false
	}
	var v []ClusterLibraryStatuses
	d := o.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in ListAllClusterLibraryStatusesResponse.
func (o *ListAllClusterLibraryStatusesResponse) SetStatuses(ctx context.Context, v []ClusterLibraryStatuses) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Statuses = types.ListValueMust(t, vs)
}

type ListAvailableZonesResponse struct {
	// The availability zone if no ``zone_id`` is provided in the cluster
	// creation request.
	DefaultZone types.String `tfsdk:"default_zone"`
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	Zones types.List `tfsdk:"zones"`
}

func (newState *ListAvailableZonesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAvailableZonesResponse) {
}

func (newState *ListAvailableZonesResponse) SyncEffectiveFieldsDuringRead(existingState ListAvailableZonesResponse) {
}

func (c ListAvailableZonesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAvailableZonesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"zones": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAvailableZonesResponse
// only implements ToObjectValue() and Type().
func (o ListAvailableZonesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_zone": o.DefaultZone,
			"zones":        o.Zones,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAvailableZonesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListAvailableZonesResponse) GetZones(ctx context.Context) ([]types.String, bool) {
	if o.Zones.IsNull() || o.Zones.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Zones.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetZones sets the value of the Zones field in ListAvailableZonesResponse.
func (o *ListAvailableZonesResponse) SetZones(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["zones"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Zones = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClusterCompliancesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListClusterCompliancesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterCompliancesRequest
// only implements ToObjectValue() and Type().
func (o ListClusterCompliancesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
			"policy_id":  o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClusterCompliancesRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListClusterCompliancesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClusterCompliancesResponse) {
}

func (newState *ListClusterCompliancesResponse) SyncEffectiveFieldsDuringRead(existingState ListClusterCompliancesResponse) {
}

func (c ListClusterCompliancesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListClusterCompliancesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters": reflect.TypeOf(ClusterCompliance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterCompliancesResponse
// only implements ToObjectValue() and Type().
func (o ListClusterCompliancesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusters":        o.Clusters,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClusterCompliancesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListClusterCompliancesResponse) GetClusters(ctx context.Context) ([]ClusterCompliance, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []ClusterCompliance
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in ListClusterCompliancesResponse.
func (o *ListClusterCompliancesResponse) SetClusters(ctx context.Context, v []ClusterCompliance) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClusterPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListClusterPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterPoliciesRequest
// only implements ToObjectValue() and Type().
func (o ListClusterPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sort_column": o.SortColumn,
			"sort_order":  o.SortOrder,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClusterPoliciesRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListClustersFilterBy) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClustersFilterBy) {
}

func (newState *ListClustersFilterBy) SyncEffectiveFieldsDuringRead(existingState ListClustersFilterBy) {
}

func (c ListClustersFilterBy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListClustersFilterBy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_sources": reflect.TypeOf(types.String{}),
		"cluster_states":  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersFilterBy
// only implements ToObjectValue() and Type().
func (o ListClustersFilterBy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_sources": o.ClusterSources,
			"cluster_states":  o.ClusterStates,
			"is_pinned":       o.IsPinned,
			"policy_id":       o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClustersFilterBy) Type(ctx context.Context) attr.Type {
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
func (o *ListClustersFilterBy) GetClusterSources(ctx context.Context) ([]types.String, bool) {
	if o.ClusterSources.IsNull() || o.ClusterSources.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ClusterSources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterSources sets the value of the ClusterSources field in ListClustersFilterBy.
func (o *ListClustersFilterBy) SetClusterSources(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_sources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ClusterSources = types.ListValueMust(t, vs)
}

// GetClusterStates returns the value of the ClusterStates field in ListClustersFilterBy as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListClustersFilterBy) GetClusterStates(ctx context.Context) ([]types.String, bool) {
	if o.ClusterStates.IsNull() || o.ClusterStates.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ClusterStates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterStates sets the value of the ClusterStates field in ListClustersFilterBy.
func (o *ListClustersFilterBy) SetClusterStates(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_states"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ClusterStates = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClustersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListClustersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter_by": reflect.TypeOf(ListClustersFilterBy{}),
		"sort_by":   reflect.TypeOf(ListClustersSortBy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersRequest
// only implements ToObjectValue() and Type().
func (o ListClustersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_by":  o.FilterBy,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
			"sort_by":    o.SortBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClustersRequest) Type(ctx context.Context) attr.Type {
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
func (o *ListClustersRequest) GetFilterBy(ctx context.Context) (ListClustersFilterBy, bool) {
	var e ListClustersFilterBy
	if o.FilterBy.IsNull() || o.FilterBy.IsUnknown() {
		return e, false
	}
	var v []ListClustersFilterBy
	d := o.FilterBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilterBy sets the value of the FilterBy field in ListClustersRequest.
func (o *ListClustersRequest) SetFilterBy(ctx context.Context, v ListClustersFilterBy) {
	vs := v.ToObjectValue(ctx)
	o.FilterBy = vs
}

// GetSortBy returns the value of the SortBy field in ListClustersRequest as
// a ListClustersSortBy value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListClustersRequest) GetSortBy(ctx context.Context) (ListClustersSortBy, bool) {
	var e ListClustersSortBy
	if o.SortBy.IsNull() || o.SortBy.IsUnknown() {
		return e, false
	}
	var v []ListClustersSortBy
	d := o.SortBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSortBy sets the value of the SortBy field in ListClustersRequest.
func (o *ListClustersRequest) SetSortBy(ctx context.Context, v ListClustersSortBy) {
	vs := v.ToObjectValue(ctx)
	o.SortBy = vs
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

func (newState *ListClustersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClustersResponse) {
}

func (newState *ListClustersResponse) SyncEffectiveFieldsDuringRead(existingState ListClustersResponse) {
}

func (c ListClustersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListClustersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters": reflect.TypeOf(ClusterDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersResponse
// only implements ToObjectValue() and Type().
func (o ListClustersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusters":        o.Clusters,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClustersResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListClustersResponse) GetClusters(ctx context.Context) ([]ClusterDetails, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []ClusterDetails
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in ListClustersResponse.
func (o *ListClustersResponse) SetClusters(ctx context.Context, v []ClusterDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

type ListClustersSortBy struct {
	// The direction to sort by.
	Direction types.String `tfsdk:"direction"`
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	Field types.String `tfsdk:"field"`
}

func (newState *ListClustersSortBy) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClustersSortBy) {
}

func (newState *ListClustersSortBy) SyncEffectiveFieldsDuringRead(existingState ListClustersSortBy) {
}

func (c ListClustersSortBy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListClustersSortBy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersSortBy
// only implements ToObjectValue() and Type().
func (o ListClustersSortBy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direction": o.Direction,
			"field":     o.Field,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClustersSortBy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direction": types.StringType,
			"field":     types.StringType,
		},
	}
}

type ListGlobalInitScriptsRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGlobalInitScriptsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListGlobalInitScriptsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGlobalInitScriptsRequest
// only implements ToObjectValue() and Type().
func (o ListGlobalInitScriptsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListGlobalInitScriptsResponse struct {
	Scripts types.List `tfsdk:"scripts"`
}

func (newState *ListGlobalInitScriptsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListGlobalInitScriptsResponse) {
}

func (newState *ListGlobalInitScriptsResponse) SyncEffectiveFieldsDuringRead(existingState ListGlobalInitScriptsResponse) {
}

func (c ListGlobalInitScriptsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListGlobalInitScriptsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scripts": reflect.TypeOf(GlobalInitScriptDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGlobalInitScriptsResponse
// only implements ToObjectValue() and Type().
func (o ListGlobalInitScriptsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scripts": o.Scripts,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListGlobalInitScriptsResponse) GetScripts(ctx context.Context) ([]GlobalInitScriptDetails, bool) {
	if o.Scripts.IsNull() || o.Scripts.IsUnknown() {
		return nil, false
	}
	var v []GlobalInitScriptDetails
	d := o.Scripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScripts sets the value of the Scripts field in ListGlobalInitScriptsResponse.
func (o *ListGlobalInitScriptsResponse) SetScripts(ctx context.Context, v []GlobalInitScriptDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scripts = types.ListValueMust(t, vs)
}

type ListInstancePools struct {
	InstancePools types.List `tfsdk:"instance_pools"`
}

func (newState *ListInstancePools) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListInstancePools) {
}

func (newState *ListInstancePools) SyncEffectiveFieldsDuringRead(existingState ListInstancePools) {
}

func (c ListInstancePools) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListInstancePools) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"instance_pools": reflect.TypeOf(InstancePoolAndStats{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstancePools
// only implements ToObjectValue() and Type().
func (o ListInstancePools) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pools": o.InstancePools,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstancePools) Type(ctx context.Context) attr.Type {
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
func (o *ListInstancePools) GetInstancePools(ctx context.Context) ([]InstancePoolAndStats, bool) {
	if o.InstancePools.IsNull() || o.InstancePools.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAndStats
	d := o.InstancePools.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstancePools sets the value of the InstancePools field in ListInstancePools.
func (o *ListInstancePools) SetInstancePools(ctx context.Context, v []InstancePoolAndStats) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_pools"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InstancePools = types.ListValueMust(t, vs)
}

type ListInstancePoolsRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstancePoolsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListInstancePoolsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstancePoolsRequest
// only implements ToObjectValue() and Type().
func (o ListInstancePoolsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstancePoolsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListInstanceProfilesRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstanceProfilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListInstanceProfilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstanceProfilesRequest
// only implements ToObjectValue() and Type().
func (o ListInstanceProfilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstanceProfilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListInstanceProfilesResponse struct {
	// A list of instance profiles that the user can access.
	InstanceProfiles types.List `tfsdk:"instance_profiles"`
}

func (newState *ListInstanceProfilesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListInstanceProfilesResponse) {
}

func (newState *ListInstanceProfilesResponse) SyncEffectiveFieldsDuringRead(existingState ListInstanceProfilesResponse) {
}

func (c ListInstanceProfilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListInstanceProfilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"instance_profiles": reflect.TypeOf(InstanceProfile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstanceProfilesResponse
// only implements ToObjectValue() and Type().
func (o ListInstanceProfilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_profiles": o.InstanceProfiles,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstanceProfilesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListInstanceProfilesResponse) GetInstanceProfiles(ctx context.Context) ([]InstanceProfile, bool) {
	if o.InstanceProfiles.IsNull() || o.InstanceProfiles.IsUnknown() {
		return nil, false
	}
	var v []InstanceProfile
	d := o.InstanceProfiles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstanceProfiles sets the value of the InstanceProfiles field in ListInstanceProfilesResponse.
func (o *ListInstanceProfilesResponse) SetInstanceProfiles(ctx context.Context, v []InstanceProfile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_profiles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InstanceProfiles = types.ListValueMust(t, vs)
}

type ListNodeTypesRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNodeTypesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNodeTypesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNodeTypesRequest
// only implements ToObjectValue() and Type().
func (o ListNodeTypesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListNodeTypesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNodeTypesResponse struct {
	// The list of available Spark node types.
	NodeTypes types.List `tfsdk:"node_types"`
}

func (newState *ListNodeTypesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNodeTypesResponse) {
}

func (newState *ListNodeTypesResponse) SyncEffectiveFieldsDuringRead(existingState ListNodeTypesResponse) {
}

func (c ListNodeTypesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNodeTypesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_types": reflect.TypeOf(NodeType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNodeTypesResponse
// only implements ToObjectValue() and Type().
func (o ListNodeTypesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"node_types": o.NodeTypes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNodeTypesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListNodeTypesResponse) GetNodeTypes(ctx context.Context) ([]NodeType, bool) {
	if o.NodeTypes.IsNull() || o.NodeTypes.IsUnknown() {
		return nil, false
	}
	var v []NodeType
	d := o.NodeTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeTypes sets the value of the NodeTypes field in ListNodeTypesResponse.
func (o *ListNodeTypesResponse) SetNodeTypes(ctx context.Context, v []NodeType) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NodeTypes = types.ListValueMust(t, vs)
}

type ListPoliciesResponse struct {
	// List of policies.
	Policies types.List `tfsdk:"policies"`
}

func (newState *ListPoliciesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPoliciesResponse) {
}

func (newState *ListPoliciesResponse) SyncEffectiveFieldsDuringRead(existingState ListPoliciesResponse) {
}

func (c ListPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(Policy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPoliciesResponse
// only implements ToObjectValue() and Type().
func (o ListPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policies": o.Policies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPoliciesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListPoliciesResponse) GetPolicies(ctx context.Context) ([]Policy, bool) {
	if o.Policies.IsNull() || o.Policies.IsUnknown() {
		return nil, false
	}
	var v []Policy
	d := o.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListPoliciesResponse.
func (o *ListPoliciesResponse) SetPolicies(ctx context.Context, v []Policy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Policies = types.ListValueMust(t, vs)
}

type ListPolicyFamiliesRequest struct {
	// Maximum number of policy families to return.
	MaxResults types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPolicyFamiliesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPolicyFamiliesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPolicyFamiliesRequest
// only implements ToObjectValue() and Type().
func (o ListPolicyFamiliesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListPolicyFamiliesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPolicyFamiliesResponse) {
}

func (newState *ListPolicyFamiliesResponse) SyncEffectiveFieldsDuringRead(existingState ListPolicyFamiliesResponse) {
}

func (c ListPolicyFamiliesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListPolicyFamiliesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy_families": reflect.TypeOf(PolicyFamily{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPolicyFamiliesResponse
// only implements ToObjectValue() and Type().
func (o ListPolicyFamiliesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"policy_families": o.PolicyFamilies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListPolicyFamiliesResponse) GetPolicyFamilies(ctx context.Context) ([]PolicyFamily, bool) {
	if o.PolicyFamilies.IsNull() || o.PolicyFamilies.IsUnknown() {
		return nil, false
	}
	var v []PolicyFamily
	d := o.PolicyFamilies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicyFamilies sets the value of the PolicyFamilies field in ListPolicyFamiliesResponse.
func (o *ListPolicyFamiliesResponse) SetPolicyFamilies(ctx context.Context, v []PolicyFamily) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["policy_families"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PolicyFamilies = types.ListValueMust(t, vs)
}

type ListZonesRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListZonesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListZonesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListZonesRequest
// only implements ToObjectValue() and Type().
func (o ListZonesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListZonesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LocalFileInfo struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	Destination types.String `tfsdk:"destination"`
}

func (newState *LocalFileInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan LocalFileInfo) {
}

func (newState *LocalFileInfo) SyncEffectiveFieldsDuringRead(existingState LocalFileInfo) {
}

func (c LocalFileInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LocalFileInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LocalFileInfo
// only implements ToObjectValue() and Type().
func (o LocalFileInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LocalFileInfo) Type(ctx context.Context) attr.Type {
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

func (newState *LogAnalyticsInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogAnalyticsInfo) {
}

func (newState *LogAnalyticsInfo) SyncEffectiveFieldsDuringRead(existingState LogAnalyticsInfo) {
}

func (c LogAnalyticsInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogAnalyticsInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogAnalyticsInfo
// only implements ToObjectValue() and Type().
func (o LogAnalyticsInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_analytics_primary_key":  o.LogAnalyticsPrimaryKey,
			"log_analytics_workspace_id": o.LogAnalyticsWorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogAnalyticsInfo) Type(ctx context.Context) attr.Type {
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

func (newState *LogSyncStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogSyncStatus) {
}

func (newState *LogSyncStatus) SyncEffectiveFieldsDuringRead(existingState LogSyncStatus) {
}

func (c LogSyncStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogSyncStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogSyncStatus
// only implements ToObjectValue() and Type().
func (o LogSyncStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_attempted": o.LastAttempted,
			"last_exception": o.LastException,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogSyncStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_attempted": types.Int64Type,
			"last_exception": types.StringType,
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

func (newState *MavenLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan MavenLibrary) {
}

func (newState *MavenLibrary) SyncEffectiveFieldsDuringRead(existingState MavenLibrary) {
}

func (c MavenLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MavenLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclusions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MavenLibrary
// only implements ToObjectValue() and Type().
func (o MavenLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"coordinates": o.Coordinates,
			"exclusions":  o.Exclusions,
			"repo":        o.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MavenLibrary) Type(ctx context.Context) attr.Type {
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
func (o *MavenLibrary) GetExclusions(ctx context.Context) ([]types.String, bool) {
	if o.Exclusions.IsNull() || o.Exclusions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Exclusions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExclusions sets the value of the Exclusions field in MavenLibrary.
func (o *MavenLibrary) SetExclusions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exclusions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Exclusions = types.ListValueMust(t, vs)
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

func (newState *NodeInstanceType) SyncEffectiveFieldsDuringCreateOrUpdate(plan NodeInstanceType) {
}

func (newState *NodeInstanceType) SyncEffectiveFieldsDuringRead(existingState NodeInstanceType) {
}

func (c NodeInstanceType) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NodeInstanceType) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NodeInstanceType
// only implements ToObjectValue() and Type().
func (o NodeInstanceType) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_type_id":        o.InstanceTypeId,
			"local_disk_size_gb":      o.LocalDiskSizeGb,
			"local_disks":             o.LocalDisks,
			"local_nvme_disk_size_gb": o.LocalNvmeDiskSizeGb,
			"local_nvme_disks":        o.LocalNvmeDisks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NodeInstanceType) Type(ctx context.Context) attr.Type {
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

func (newState *NodeType) SyncEffectiveFieldsDuringCreateOrUpdate(plan NodeType) {
}

func (newState *NodeType) SyncEffectiveFieldsDuringRead(existingState NodeType) {
}

func (c NodeType) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NodeType) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_info":          reflect.TypeOf(CloudProviderNodeInfo{}),
		"node_instance_type": reflect.TypeOf(NodeInstanceType{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NodeType
// only implements ToObjectValue() and Type().
func (o NodeType) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"category":                o.Category,
			"description":             o.Description,
			"display_order":           o.DisplayOrder,
			"instance_type_id":        o.InstanceTypeId,
			"is_deprecated":           o.IsDeprecated,
			"is_encrypted_in_transit": o.IsEncryptedInTransit,
			"is_graviton":             o.IsGraviton,
			"is_hidden":               o.IsHidden,
			"is_io_cache_enabled":     o.IsIoCacheEnabled,
			"memory_mb":               o.MemoryMb,
			"node_info":               o.NodeInfo,
			"node_instance_type":      o.NodeInstanceType,
			"node_type_id":            o.NodeTypeId,
			"num_cores":               o.NumCores,
			"num_gpus":                o.NumGpus,
			"photon_driver_capable":   o.PhotonDriverCapable,
			"photon_worker_capable":   o.PhotonWorkerCapable,
			"support_cluster_tags":    o.SupportClusterTags,
			"support_ebs_volumes":     o.SupportEbsVolumes,
			"support_port_forwarding": o.SupportPortForwarding,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NodeType) Type(ctx context.Context) attr.Type {
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
func (o *NodeType) GetNodeInfo(ctx context.Context) (CloudProviderNodeInfo, bool) {
	var e CloudProviderNodeInfo
	if o.NodeInfo.IsNull() || o.NodeInfo.IsUnknown() {
		return e, false
	}
	var v []CloudProviderNodeInfo
	d := o.NodeInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeInfo sets the value of the NodeInfo field in NodeType.
func (o *NodeType) SetNodeInfo(ctx context.Context, v CloudProviderNodeInfo) {
	vs := v.ToObjectValue(ctx)
	o.NodeInfo = vs
}

// GetNodeInstanceType returns the value of the NodeInstanceType field in NodeType as
// a NodeInstanceType value.
// If the field is unknown or null, the boolean return value is false.
func (o *NodeType) GetNodeInstanceType(ctx context.Context) (NodeInstanceType, bool) {
	var e NodeInstanceType
	if o.NodeInstanceType.IsNull() || o.NodeInstanceType.IsUnknown() {
		return e, false
	}
	var v []NodeInstanceType
	d := o.NodeInstanceType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeInstanceType sets the value of the NodeInstanceType field in NodeType.
func (o *NodeType) SetNodeInstanceType(ctx context.Context, v NodeInstanceType) {
	vs := v.ToObjectValue(ctx)
	o.NodeInstanceType = vs
}

// Error message of a failed pending instances
type PendingInstanceError struct {
	InstanceId types.String `tfsdk:"instance_id"`

	Message types.String `tfsdk:"message"`
}

func (newState *PendingInstanceError) SyncEffectiveFieldsDuringCreateOrUpdate(plan PendingInstanceError) {
}

func (newState *PendingInstanceError) SyncEffectiveFieldsDuringRead(existingState PendingInstanceError) {
}

func (c PendingInstanceError) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PendingInstanceError) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PendingInstanceError
// only implements ToObjectValue() and Type().
func (o PendingInstanceError) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_id": o.InstanceId,
			"message":     o.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PendingInstanceError) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermanentDeleteCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermanentDeleteCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermanentDeleteCluster
// only implements ToObjectValue() and Type().
func (o PermanentDeleteCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermanentDeleteCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type PermanentDeleteClusterResponse struct {
}

func (newState *PermanentDeleteClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermanentDeleteClusterResponse) {
}

func (newState *PermanentDeleteClusterResponse) SyncEffectiveFieldsDuringRead(existingState PermanentDeleteClusterResponse) {
}

func (c PermanentDeleteClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermanentDeleteClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermanentDeleteClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermanentDeleteClusterResponse
// only implements ToObjectValue() and Type().
func (o PermanentDeleteClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type PinCluster struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PinCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PinCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PinCluster
// only implements ToObjectValue() and Type().
func (o PinCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PinCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type PinClusterResponse struct {
}

func (newState *PinClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PinClusterResponse) {
}

func (newState *PinClusterResponse) SyncEffectiveFieldsDuringRead(existingState PinClusterResponse) {
}

func (c PinClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PinClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PinClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PinClusterResponse
// only implements ToObjectValue() and Type().
func (o PinClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PinClusterResponse) Type(ctx context.Context) attr.Type {
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

func (newState *Policy) SyncEffectiveFieldsDuringCreateOrUpdate(plan Policy) {
}

func (newState *Policy) SyncEffectiveFieldsDuringRead(existingState Policy) {
}

func (c Policy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Policy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Policy
// only implements ToObjectValue() and Type().
func (o Policy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at_timestamp":               o.CreatedAtTimestamp,
			"creator_user_name":                  o.CreatorUserName,
			"definition":                         o.Definition,
			"description":                        o.Description,
			"is_default":                         o.IsDefault,
			"libraries":                          o.Libraries,
			"max_clusters_per_user":              o.MaxClustersPerUser,
			"name":                               o.Name,
			"policy_family_definition_overrides": o.PolicyFamilyDefinitionOverrides,
			"policy_family_id":                   o.PolicyFamilyId,
			"policy_id":                          o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Policy) Type(ctx context.Context) attr.Type {
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
func (o *Policy) GetLibraries(ctx context.Context) ([]Library, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in Policy.
func (o *Policy) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
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

func (newState *PolicyFamily) SyncEffectiveFieldsDuringCreateOrUpdate(plan PolicyFamily) {
}

func (newState *PolicyFamily) SyncEffectiveFieldsDuringRead(existingState PolicyFamily) {
}

func (c PolicyFamily) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PolicyFamily) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PolicyFamily
// only implements ToObjectValue() and Type().
func (o PolicyFamily) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"definition":       o.Definition,
			"description":      o.Description,
			"name":             o.Name,
			"policy_family_id": o.PolicyFamilyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PolicyFamily) Type(ctx context.Context) attr.Type {
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

func (newState *PythonPyPiLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan PythonPyPiLibrary) {
}

func (newState *PythonPyPiLibrary) SyncEffectiveFieldsDuringRead(existingState PythonPyPiLibrary) {
}

func (c PythonPyPiLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PythonPyPiLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PythonPyPiLibrary
// only implements ToObjectValue() and Type().
func (o PythonPyPiLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"package": o.Package,
			"repo":    o.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PythonPyPiLibrary) Type(ctx context.Context) attr.Type {
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

func (newState *RCranLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan RCranLibrary) {
}

func (newState *RCranLibrary) SyncEffectiveFieldsDuringRead(existingState RCranLibrary) {
}

func (c RCranLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RCranLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RCranLibrary
// only implements ToObjectValue() and Type().
func (o RCranLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"package": o.Package,
			"repo":    o.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RCranLibrary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"package": types.StringType,
			"repo":    types.StringType,
		},
	}
}

type RemoveInstanceProfile struct {
	// The ARN of the instance profile to remove. This field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveInstanceProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RemoveInstanceProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveInstanceProfile
// only implements ToObjectValue() and Type().
func (o RemoveInstanceProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_profile_arn": o.InstanceProfileArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RemoveInstanceProfile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_profile_arn": types.StringType,
		},
	}
}

type RemoveResponse struct {
}

func (newState *RemoveResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RemoveResponse) {
}

func (newState *RemoveResponse) SyncEffectiveFieldsDuringRead(existingState RemoveResponse) {
}

func (c RemoveResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RemoveResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveResponse
// only implements ToObjectValue() and Type().
func (o RemoveResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RemoveResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResizeCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResizeCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale": reflect.TypeOf(AutoScale{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResizeCluster
// only implements ToObjectValue() and Type().
func (o ResizeCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":   o.Autoscale,
			"cluster_id":  o.ClusterId,
			"num_workers": o.NumWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResizeCluster) Type(ctx context.Context) attr.Type {
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
func (o *ResizeCluster) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in ResizeCluster.
func (o *ResizeCluster) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	o.Autoscale = vs
}

type ResizeClusterResponse struct {
}

func (newState *ResizeClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResizeClusterResponse) {
}

func (newState *ResizeClusterResponse) SyncEffectiveFieldsDuringRead(existingState ResizeClusterResponse) {
}

func (c ResizeClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResizeClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResizeClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResizeClusterResponse
// only implements ToObjectValue() and Type().
func (o ResizeClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ResizeClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestartCluster struct {
	// The cluster to be started.
	ClusterId types.String `tfsdk:"cluster_id"`

	RestartUser types.String `tfsdk:"restart_user"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestartCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartCluster
// only implements ToObjectValue() and Type().
func (o RestartCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":   o.ClusterId,
			"restart_user": o.RestartUser,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestartCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":   types.StringType,
			"restart_user": types.StringType,
		},
	}
}

type RestartClusterResponse struct {
}

func (newState *RestartClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestartClusterResponse) {
}

func (newState *RestartClusterResponse) SyncEffectiveFieldsDuringRead(existingState RestartClusterResponse) {
}

func (c RestartClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestartClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartClusterResponse
// only implements ToObjectValue() and Type().
func (o RestartClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RestartClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Results struct {
	// The cause of the error
	Cause types.String `tfsdk:"cause"`

	Data types.Object `tfsdk:"data"`
	// The image filename
	FileName types.String `tfsdk:"fileName"`

	FileNames types.List `tfsdk:"fileNames"`
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	IsJsonSchema types.Bool `tfsdk:"isJsonSchema"`
	// internal field used by SDK
	Pos types.Int64 `tfsdk:"pos"`

	ResultType types.String `tfsdk:"resultType"`
	// The table schema
	Schema types.List `tfsdk:"schema"`
	// The summary of the error
	Summary types.String `tfsdk:"summary"`
	// true if partial results are returned.
	Truncated types.Bool `tfsdk:"truncated"`
}

func (newState *Results) SyncEffectiveFieldsDuringCreateOrUpdate(plan Results) {
}

func (newState *Results) SyncEffectiveFieldsDuringRead(existingState Results) {
}

func (c Results) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["data"] = attrs["data"].SetOptional()
	attrs["fileName"] = attrs["fileName"].SetOptional()
	attrs["fileNames"] = attrs["fileNames"].SetOptional()
	attrs["isJsonSchema"] = attrs["isJsonSchema"].SetOptional()
	attrs["pos"] = attrs["pos"].SetOptional()
	attrs["resultType"] = attrs["resultType"].SetOptional()
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
func (a Results) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fileNames": reflect.TypeOf(types.String{}),
		"schema":    reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Results
// only implements ToObjectValue() and Type().
func (o Results) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cause":        o.Cause,
			"data":         o.Data,
			"fileName":     o.FileName,
			"fileNames":    o.FileNames,
			"isJsonSchema": o.IsJsonSchema,
			"pos":          o.Pos,
			"resultType":   o.ResultType,
			"schema":       o.Schema,
			"summary":      o.Summary,
			"truncated":    o.Truncated,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Results) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":    types.StringType,
			"data":     types.ObjectType{},
			"fileName": types.StringType,
			"fileNames": basetypes.ListType{
				ElemType: types.StringType,
			},
			"isJsonSchema": types.BoolType,
			"pos":          types.Int64Type,
			"resultType":   types.StringType,
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
func (o *Results) GetFileNames(ctx context.Context) ([]types.String, bool) {
	if o.FileNames.IsNull() || o.FileNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.FileNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileNames sets the value of the FileNames field in Results.
func (o *Results) SetFileNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["fileNames"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FileNames = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in Results as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *Results) GetSchema(ctx context.Context) ([]types.Object, bool) {
	if o.Schema.IsNull() || o.Schema.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := o.Schema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchema sets the value of the Schema field in Results.
func (o *Results) SetSchema(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schema = types.ListValueMust(t, vs)
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

func (newState *S3StorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan S3StorageInfo) {
}

func (newState *S3StorageInfo) SyncEffectiveFieldsDuringRead(existingState S3StorageInfo) {
}

func (c S3StorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a S3StorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, S3StorageInfo
// only implements ToObjectValue() and Type().
func (o S3StorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"canned_acl":        o.CannedAcl,
			"destination":       o.Destination,
			"enable_encryption": o.EnableEncryption,
			"encryption_type":   o.EncryptionType,
			"endpoint":          o.Endpoint,
			"kms_key":           o.KmsKey,
			"region":            o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o S3StorageInfo) Type(ctx context.Context) attr.Type {
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

func (newState *SparkNode) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkNode) {
}

func (newState *SparkNode) SyncEffectiveFieldsDuringRead(existingState SparkNode) {
}

func (c SparkNode) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SparkNode) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_aws_attributes": reflect.TypeOf(SparkNodeAwsAttributes{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkNode
// only implements ToObjectValue() and Type().
func (o SparkNode) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"host_private_ip":     o.HostPrivateIp,
			"instance_id":         o.InstanceId,
			"node_aws_attributes": o.NodeAwsAttributes,
			"node_id":             o.NodeId,
			"private_ip":          o.PrivateIp,
			"public_dns":          o.PublicDns,
			"start_timestamp":     o.StartTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparkNode) Type(ctx context.Context) attr.Type {
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
func (o *SparkNode) GetNodeAwsAttributes(ctx context.Context) (SparkNodeAwsAttributes, bool) {
	var e SparkNodeAwsAttributes
	if o.NodeAwsAttributes.IsNull() || o.NodeAwsAttributes.IsUnknown() {
		return e, false
	}
	var v []SparkNodeAwsAttributes
	d := o.NodeAwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeAwsAttributes sets the value of the NodeAwsAttributes field in SparkNode.
func (o *SparkNode) SetNodeAwsAttributes(ctx context.Context, v SparkNodeAwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.NodeAwsAttributes = vs
}

// Attributes specific to AWS for a Spark node.
type SparkNodeAwsAttributes struct {
	// Whether this node is on an Amazon spot instance.
	IsSpot types.Bool `tfsdk:"is_spot"`
}

func (newState *SparkNodeAwsAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkNodeAwsAttributes) {
}

func (newState *SparkNodeAwsAttributes) SyncEffectiveFieldsDuringRead(existingState SparkNodeAwsAttributes) {
}

func (c SparkNodeAwsAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SparkNodeAwsAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkNodeAwsAttributes
// only implements ToObjectValue() and Type().
func (o SparkNodeAwsAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_spot": o.IsSpot,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparkNodeAwsAttributes) Type(ctx context.Context) attr.Type {
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

func (newState *SparkVersion) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkVersion) {
}

func (newState *SparkVersion) SyncEffectiveFieldsDuringRead(existingState SparkVersion) {
}

func (c SparkVersion) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SparkVersion) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkVersion
// only implements ToObjectValue() and Type().
func (o SparkVersion) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":  o.Key,
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparkVersion) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":  types.StringType,
			"name": types.StringType,
		},
	}
}

type SparkVersionsRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparkVersionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkVersionsRequest
// only implements ToObjectValue() and Type().
func (o SparkVersionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SparkVersionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type StartCluster struct {
	// The cluster to be started.
	ClusterId types.String `tfsdk:"cluster_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartCluster
// only implements ToObjectValue() and Type().
func (o StartCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type StartClusterResponse struct {
}

func (newState *StartClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartClusterResponse) {
}

func (newState *StartClusterResponse) SyncEffectiveFieldsDuringRead(existingState StartClusterResponse) {
}

func (c StartClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartClusterResponse
// only implements ToObjectValue() and Type().
func (o StartClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o StartClusterResponse) Type(ctx context.Context) attr.Type {
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

func (newState *TerminationReason) SyncEffectiveFieldsDuringCreateOrUpdate(plan TerminationReason) {
}

func (newState *TerminationReason) SyncEffectiveFieldsDuringRead(existingState TerminationReason) {
}

func (c TerminationReason) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TerminationReason) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TerminationReason
// only implements ToObjectValue() and Type().
func (o TerminationReason) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"code":       o.Code,
			"parameters": o.Parameters,
			"type":       o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TerminationReason) Type(ctx context.Context) attr.Type {
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
func (o *TerminationReason) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in TerminationReason.
func (o *TerminationReason) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.MapValueMust(t, vs)
}

type UninstallLibraries struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The libraries to uninstall.
	Libraries types.List `tfsdk:"libraries"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UninstallLibraries.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UninstallLibraries) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UninstallLibraries
// only implements ToObjectValue() and Type().
func (o UninstallLibraries) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
			"libraries":  o.Libraries,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UninstallLibraries) Type(ctx context.Context) attr.Type {
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
func (o *UninstallLibraries) GetLibraries(ctx context.Context) ([]Library, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in UninstallLibraries.
func (o *UninstallLibraries) SetLibraries(ctx context.Context, v []Library) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type UninstallLibrariesResponse struct {
}

func (newState *UninstallLibrariesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UninstallLibrariesResponse) {
}

func (newState *UninstallLibrariesResponse) SyncEffectiveFieldsDuringRead(existingState UninstallLibrariesResponse) {
}

func (c UninstallLibrariesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UninstallLibrariesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UninstallLibrariesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UninstallLibrariesResponse
// only implements ToObjectValue() and Type().
func (o UninstallLibrariesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UnpinCluster struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpinCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpinCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpinCluster
// only implements ToObjectValue() and Type().
func (o UnpinCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UnpinCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type UnpinClusterResponse struct {
}

func (newState *UnpinClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpinClusterResponse) {
}

func (newState *UnpinClusterResponse) SyncEffectiveFieldsDuringRead(existingState UnpinClusterResponse) {
}

func (c UnpinClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpinClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpinClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpinClusterResponse
// only implements ToObjectValue() and Type().
func (o UnpinClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UnpinClusterResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster": reflect.TypeOf(UpdateClusterResource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCluster
// only implements ToObjectValue() and Type().
func (o UpdateCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster":     o.Cluster,
			"cluster_id":  o.ClusterId,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCluster) Type(ctx context.Context) attr.Type {
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
func (o *UpdateCluster) GetCluster(ctx context.Context) (UpdateClusterResource, bool) {
	var e UpdateClusterResource
	if o.Cluster.IsNull() || o.Cluster.IsUnknown() {
		return e, false
	}
	var v []UpdateClusterResource
	d := o.Cluster.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCluster sets the value of the Cluster field in UpdateCluster.
func (o *UpdateCluster) SetCluster(ctx context.Context, v UpdateClusterResource) {
	vs := v.ToObjectValue(ctx)
	o.Cluster = vs
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

func (newState *UpdateClusterResource) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateClusterResource) {
}

func (newState *UpdateClusterResource) SyncEffectiveFieldsDuringRead(existingState UpdateClusterResource) {
}

func (c UpdateClusterResource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateClusterResource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o UpdateClusterResource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":                      o.Autoscale,
			"autotermination_minutes":        o.AutoterminationMinutes,
			"aws_attributes":                 o.AwsAttributes,
			"azure_attributes":               o.AzureAttributes,
			"cluster_log_conf":               o.ClusterLogConf,
			"cluster_name":                   o.ClusterName,
			"custom_tags":                    o.CustomTags,
			"data_security_mode":             o.DataSecurityMode,
			"docker_image":                   o.DockerImage,
			"driver_instance_pool_id":        o.DriverInstancePoolId,
			"driver_node_type_id":            o.DriverNodeTypeId,
			"enable_elastic_disk":            o.EnableElasticDisk,
			"enable_local_disk_encryption":   o.EnableLocalDiskEncryption,
			"gcp_attributes":                 o.GcpAttributes,
			"init_scripts":                   o.InitScripts,
			"instance_pool_id":               o.InstancePoolId,
			"is_single_node":                 o.IsSingleNode,
			"kind":                           o.Kind,
			"node_type_id":                   o.NodeTypeId,
			"num_workers":                    o.NumWorkers,
			"policy_id":                      o.PolicyId,
			"remote_disk_throughput":         o.RemoteDiskThroughput,
			"runtime_engine":                 o.RuntimeEngine,
			"single_user_name":               o.SingleUserName,
			"spark_conf":                     o.SparkConf,
			"spark_env_vars":                 o.SparkEnvVars,
			"spark_version":                  o.SparkVersion,
			"ssh_public_keys":                o.SshPublicKeys,
			"total_initial_remote_disk_size": o.TotalInitialRemoteDiskSize,
			"use_ml_runtime":                 o.UseMlRuntime,
			"workload_type":                  o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateClusterResource) Type(ctx context.Context) attr.Type {
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
func (o *UpdateClusterResource) GetAutoscale(ctx context.Context) (AutoScale, bool) {
	var e AutoScale
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in UpdateClusterResource.
func (o *UpdateClusterResource) SetAutoscale(ctx context.Context, v AutoScale) {
	vs := v.ToObjectValue(ctx)
	o.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in UpdateClusterResource as
// a AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetAwsAttributes(ctx context.Context) (AwsAttributes, bool) {
	var e AwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in UpdateClusterResource.
func (o *UpdateClusterResource) SetAwsAttributes(ctx context.Context, v AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in UpdateClusterResource as
// a AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetAzureAttributes(ctx context.Context) (AzureAttributes, bool) {
	var e AzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in UpdateClusterResource.
func (o *UpdateClusterResource) SetAzureAttributes(ctx context.Context, v AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in UpdateClusterResource as
// a ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetClusterLogConf(ctx context.Context) (ClusterLogConf, bool) {
	var e ClusterLogConf
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf
	d := o.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in UpdateClusterResource.
func (o *UpdateClusterResource) SetClusterLogConf(ctx context.Context, v ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	o.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in UpdateClusterResource as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in UpdateClusterResource.
func (o *UpdateClusterResource) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in UpdateClusterResource as
// a DockerImage value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetDockerImage(ctx context.Context) (DockerImage, bool) {
	var e DockerImage
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage
	d := o.DockerImage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in UpdateClusterResource.
func (o *UpdateClusterResource) SetDockerImage(ctx context.Context, v DockerImage) {
	vs := v.ToObjectValue(ctx)
	o.DockerImage = vs
}

// GetGcpAttributes returns the value of the GcpAttributes field in UpdateClusterResource as
// a GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetGcpAttributes(ctx context.Context) (GcpAttributes, bool) {
	var e GcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in UpdateClusterResource.
func (o *UpdateClusterResource) SetGcpAttributes(ctx context.Context, v GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in UpdateClusterResource as
// a slice of InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetInitScripts(ctx context.Context) ([]InitScriptInfo, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in UpdateClusterResource.
func (o *UpdateClusterResource) SetInitScripts(ctx context.Context, v []InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in UpdateClusterResource as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkConf.IsNull() || o.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in UpdateClusterResource.
func (o *UpdateClusterResource) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in UpdateClusterResource as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkEnvVars.IsNull() || o.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in UpdateClusterResource.
func (o *UpdateClusterResource) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in UpdateClusterResource as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if o.SshPublicKeys.IsNull() || o.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in UpdateClusterResource.
func (o *UpdateClusterResource) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in UpdateClusterResource as
// a WorkloadType value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource) GetWorkloadType(ctx context.Context) (WorkloadType, bool) {
	var e WorkloadType
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType
	d := o.WorkloadType.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in UpdateClusterResource.
func (o *UpdateClusterResource) SetWorkloadType(ctx context.Context, v WorkloadType) {
	vs := v.ToObjectValue(ctx)
	o.WorkloadType = vs
}

type UpdateClusterResponse struct {
}

func (newState *UpdateClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateClusterResponse) {
}

func (newState *UpdateClusterResponse) SyncEffectiveFieldsDuringRead(existingState UpdateClusterResponse) {
}

func (c UpdateClusterResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateClusterResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateClusterResponse
// only implements ToObjectValue() and Type().
func (o UpdateClusterResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

func (c UpdateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse
// only implements ToObjectValue() and Type().
func (o UpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse) Type(ctx context.Context) attr.Type {
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

func (newState *VolumesStorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan VolumesStorageInfo) {
}

func (newState *VolumesStorageInfo) SyncEffectiveFieldsDuringRead(existingState VolumesStorageInfo) {
}

func (c VolumesStorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a VolumesStorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VolumesStorageInfo
// only implements ToObjectValue() and Type().
func (o VolumesStorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VolumesStorageInfo) Type(ctx context.Context) attr.Type {
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

func (newState *WorkloadType) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkloadType) {
}

func (newState *WorkloadType) SyncEffectiveFieldsDuringRead(existingState WorkloadType) {
}

func (c WorkloadType) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkloadType) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clients": reflect.TypeOf(ClientsTypes{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkloadType
// only implements ToObjectValue() and Type().
func (o WorkloadType) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clients": o.Clients,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkloadType) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clients": ClientsTypes{}.Type(ctx),
		},
	}
}

// GetClients returns the value of the Clients field in WorkloadType as
// a ClientsTypes value.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkloadType) GetClients(ctx context.Context) (ClientsTypes, bool) {
	var e ClientsTypes
	if o.Clients.IsNull() || o.Clients.IsUnknown() {
		return e, false
	}
	var v []ClientsTypes
	d := o.Clients.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClients sets the value of the Clients field in WorkloadType.
func (o *WorkloadType) SetClients(ctx context.Context, v ClientsTypes) {
	vs := v.ToObjectValue(ctx)
	o.Clients = vs
}

// A storage location in Workspace Filesystem (WSFS)
type WorkspaceStorageInfo struct {
	// wsfs destination, e.g. `workspace:/cluster-init-scripts/setup-datadog.sh`
	Destination types.String `tfsdk:"destination"`
}

func (newState *WorkspaceStorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceStorageInfo) {
}

func (newState *WorkspaceStorageInfo) SyncEffectiveFieldsDuringRead(existingState WorkspaceStorageInfo) {
}

func (c WorkspaceStorageInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceStorageInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceStorageInfo
// only implements ToObjectValue() and Type().
func (o WorkspaceStorageInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}
