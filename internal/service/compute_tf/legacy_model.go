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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AddInstanceProfile_SdkV2 struct {
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
func (a AddInstanceProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddInstanceProfile_SdkV2
// only implements ToObjectValue() and Type().
func (o AddInstanceProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AddInstanceProfile_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"iam_role_arn":             types.StringType,
			"instance_profile_arn":     types.StringType,
			"is_meta_instance_profile": types.BoolType,
			"skip_validation":          types.BoolType,
		},
	}
}

type AddResponse_SdkV2 struct {
}

func (newState *AddResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan AddResponse_SdkV2) {
}

func (newState *AddResponse_SdkV2) SyncFieldsDuringRead(existingState AddResponse_SdkV2) {
}

func (c AddResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o AddResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o AddResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A storage location in Adls Gen2
type Adlsgen2Info_SdkV2 struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	Destination types.String `tfsdk:"destination"`
}

func (newState *Adlsgen2Info_SdkV2) SyncFieldsDuringCreateOrUpdate(plan Adlsgen2Info_SdkV2) {
}

func (newState *Adlsgen2Info_SdkV2) SyncFieldsDuringRead(existingState Adlsgen2Info_SdkV2) {
}

func (c Adlsgen2Info_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Adlsgen2Info_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Adlsgen2Info_SdkV2
// only implements ToObjectValue() and Type().
func (o Adlsgen2Info_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Adlsgen2Info_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type AutoScale_SdkV2 struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. Note that `max_workers` must be strictly greater than
	// `min_workers`.
	MaxWorkers types.Int64 `tfsdk:"max_workers"`
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers types.Int64 `tfsdk:"min_workers"`
}

func (newState *AutoScale_SdkV2) SyncFieldsDuringCreateOrUpdate(plan AutoScale_SdkV2) {
}

func (newState *AutoScale_SdkV2) SyncFieldsDuringRead(existingState AutoScale_SdkV2) {
}

func (c AutoScale_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AutoScale_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoScale_SdkV2
// only implements ToObjectValue() and Type().
func (o AutoScale_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_workers": o.MaxWorkers,
			"min_workers": o.MinWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AutoScale_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_workers": types.Int64Type,
			"min_workers": types.Int64Type,
		},
	}
}

// Attributes set during cluster creation which are related to Amazon Web
// Services.
type AwsAttributes_SdkV2 struct {
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

func (newState *AwsAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan AwsAttributes_SdkV2) {
}

func (newState *AwsAttributes_SdkV2) SyncFieldsDuringRead(existingState AwsAttributes_SdkV2) {
}

func (c AwsAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AwsAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AwsAttributes_SdkV2) Type(ctx context.Context) attr.Type {
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
type AzureAttributes_SdkV2 struct {
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
	LogAnalyticsInfo types.List `tfsdk:"log_analytics_info"`
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	SpotBidMaxPrice types.Float64 `tfsdk:"spot_bid_max_price"`
}

func (newState *AzureAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan AzureAttributes_SdkV2) {
}

func (newState *AzureAttributes_SdkV2) SyncFieldsDuringRead(existingState AzureAttributes_SdkV2) {
}

func (c AzureAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability"] = attrs["availability"].SetOptional()
	attrs["first_on_demand"] = attrs["first_on_demand"].SetOptional()
	attrs["log_analytics_info"] = attrs["log_analytics_info"].SetOptional()
	attrs["log_analytics_info"] = attrs["log_analytics_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a AzureAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"log_analytics_info": reflect.TypeOf(LogAnalyticsInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AzureAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":    types.StringType,
			"first_on_demand": types.Int64Type,
			"log_analytics_info": basetypes.ListType{
				ElemType: LogAnalyticsInfo_SdkV2{}.Type(ctx),
			},
			"spot_bid_max_price": types.Float64Type,
		},
	}
}

// GetLogAnalyticsInfo returns the value of the LogAnalyticsInfo field in AzureAttributes_SdkV2 as
// a LogAnalyticsInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AzureAttributes_SdkV2) GetLogAnalyticsInfo(ctx context.Context) (LogAnalyticsInfo_SdkV2, bool) {
	var e LogAnalyticsInfo_SdkV2
	if o.LogAnalyticsInfo.IsNull() || o.LogAnalyticsInfo.IsUnknown() {
		return e, false
	}
	var v []LogAnalyticsInfo_SdkV2
	d := o.LogAnalyticsInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogAnalyticsInfo sets the value of the LogAnalyticsInfo field in AzureAttributes_SdkV2.
func (o *AzureAttributes_SdkV2) SetLogAnalyticsInfo(ctx context.Context, v LogAnalyticsInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_analytics_info"]
	o.LogAnalyticsInfo = types.ListValueMust(t, vs)
}

type CancelCommand_SdkV2 struct {
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
func (a CancelCommand_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelCommand_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelCommand_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"commandId": o.CommandId,
			"contextId": o.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelCommand_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"commandId": types.StringType,
			"contextId": types.StringType,
		},
	}
}

type CancelResponse_SdkV2 struct {
}

func (newState *CancelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CancelResponse_SdkV2) {
}

func (newState *CancelResponse_SdkV2) SyncFieldsDuringRead(existingState CancelResponse_SdkV2) {
}

func (c CancelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CancelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ChangeClusterOwner_SdkV2 struct {
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
func (a ChangeClusterOwner_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChangeClusterOwner_SdkV2
// only implements ToObjectValue() and Type().
func (o ChangeClusterOwner_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":     o.ClusterId,
			"owner_username": o.OwnerUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ChangeClusterOwner_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":     types.StringType,
			"owner_username": types.StringType,
		},
	}
}

type ChangeClusterOwnerResponse_SdkV2 struct {
}

func (newState *ChangeClusterOwnerResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ChangeClusterOwnerResponse_SdkV2) {
}

func (newState *ChangeClusterOwnerResponse_SdkV2) SyncFieldsDuringRead(existingState ChangeClusterOwnerResponse_SdkV2) {
}

func (c ChangeClusterOwnerResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChangeClusterOwnerResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ChangeClusterOwnerResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChangeClusterOwnerResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ChangeClusterOwnerResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ClientsTypes_SdkV2 struct {
	// With jobs set, the cluster can be used for jobs
	Jobs types.Bool `tfsdk:"jobs"`
	// With notebooks set, this cluster can be used for notebooks
	Notebooks types.Bool `tfsdk:"notebooks"`
}

func (newState *ClientsTypes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClientsTypes_SdkV2) {
}

func (newState *ClientsTypes_SdkV2) SyncFieldsDuringRead(existingState ClientsTypes_SdkV2) {
}

func (c ClientsTypes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClientsTypes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClientsTypes_SdkV2
// only implements ToObjectValue() and Type().
func (o ClientsTypes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"jobs":      o.Jobs,
			"notebooks": o.Notebooks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClientsTypes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"jobs":      types.BoolType,
			"notebooks": types.BoolType,
		},
	}
}

type CloneCluster_SdkV2 struct {
	// The cluster that is being cloned.
	SourceClusterId types.String `tfsdk:"source_cluster_id"`
}

func (newState *CloneCluster_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CloneCluster_SdkV2) {
}

func (newState *CloneCluster_SdkV2) SyncFieldsDuringRead(existingState CloneCluster_SdkV2) {
}

func (c CloneCluster_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CloneCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloneCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o CloneCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"source_cluster_id": o.SourceClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloneCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"source_cluster_id": types.StringType,
		},
	}
}

type CloudProviderNodeInfo_SdkV2 struct {
	// Status as reported by the cloud provider
	Status types.List `tfsdk:"status"`
}

func (newState *CloudProviderNodeInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CloudProviderNodeInfo_SdkV2) {
}

func (newState *CloudProviderNodeInfo_SdkV2) SyncFieldsDuringRead(existingState CloudProviderNodeInfo_SdkV2) {
}

func (c CloudProviderNodeInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CloudProviderNodeInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"status": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudProviderNodeInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CloudProviderNodeInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloudProviderNodeInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetStatus returns the value of the Status field in CloudProviderNodeInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CloudProviderNodeInfo_SdkV2) GetStatus(ctx context.Context) ([]types.String, bool) {
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

// SetStatus sets the value of the Status field in CloudProviderNodeInfo_SdkV2.
func (o *CloudProviderNodeInfo_SdkV2) SetStatus(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Status = types.ListValueMust(t, vs)
}

type ClusterAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *ClusterAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterAccessControlRequest_SdkV2) {
}

func (newState *ClusterAccessControlRequest_SdkV2) SyncFieldsDuringRead(existingState ClusterAccessControlRequest_SdkV2) {
}

func (c ClusterAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ClusterAccessControlResponse_SdkV2 struct {
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

func (newState *ClusterAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterAccessControlResponse_SdkV2) {
}

func (newState *ClusterAccessControlResponse_SdkV2) SyncFieldsDuringRead(existingState ClusterAccessControlResponse_SdkV2) {
}

func (c ClusterAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ClusterPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ClusterPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in ClusterAccessControlResponse_SdkV2 as
// a slice of ClusterPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]ClusterPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ClusterPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ClusterAccessControlResponse_SdkV2.
func (o *ClusterAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []ClusterPermission_SdkV2) {
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
type ClusterAttributes_SdkV2 struct {
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf"`
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
	DockerImage types.List `tfsdk:"docker_image"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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

	WorkloadType types.List `tfsdk:"workload_type"`
}

func (newState *ClusterAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterAttributes_SdkV2) {
}

func (newState *ClusterAttributes_SdkV2) SyncFieldsDuringRead(existingState ClusterAttributes_SdkV2) {
}

func (c ClusterAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
	attrs["workload_type"] = attrs["workload_type"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":   reflect.TypeOf(AwsAttributes_SdkV2{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes_SdkV2{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf_SdkV2{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage_SdkV2{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes_SdkV2{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo_SdkV2{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes_SdkV2{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf_SdkV2{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes_SdkV2{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo_SdkV2{}.Type(ctx),
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAwsAttributes returns the value of the AwsAttributes field in ClusterAttributes_SdkV2 as
// a AwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetAwsAttributes(ctx context.Context) (AwsAttributes_SdkV2, bool) {
	var e AwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetAwsAttributes(ctx context.Context, v AwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterAttributes_SdkV2 as
// a AzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetAzureAttributes(ctx context.Context) (AzureAttributes_SdkV2, bool) {
	var e AzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetAzureAttributes(ctx context.Context, v AzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterAttributes_SdkV2 as
// a ClusterLogConf_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetClusterLogConf(ctx context.Context) (ClusterLogConf_SdkV2, bool) {
	var e ClusterLogConf_SdkV2
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf_SdkV2
	d := o.ClusterLogConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetClusterLogConf(ctx context.Context, v ClusterLogConf_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_conf"]
	o.ClusterLogConf = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in ClusterAttributes_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterAttributes_SdkV2 as
// a DockerImage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetDockerImage(ctx context.Context) (DockerImage_SdkV2, bool) {
	var e DockerImage_SdkV2
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage_SdkV2
	d := o.DockerImage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetDockerImage(ctx context.Context, v DockerImage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["docker_image"]
	o.DockerImage = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterAttributes_SdkV2 as
// a GcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetGcpAttributes(ctx context.Context) (GcpAttributes_SdkV2, bool) {
	var e GcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetGcpAttributes(ctx context.Context, v GcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in ClusterAttributes_SdkV2 as
// a slice of InitScriptInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetInitScripts(ctx context.Context) ([]InitScriptInfo_SdkV2, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo_SdkV2
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetInitScripts(ctx context.Context, v []InitScriptInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterAttributes_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkConf sets the value of the SparkConf field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterAttributes_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterAttributes_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
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

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterAttributes_SdkV2 as
// a WorkloadType_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAttributes_SdkV2) GetWorkloadType(ctx context.Context) (WorkloadType_SdkV2, bool) {
	var e WorkloadType_SdkV2
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType_SdkV2
	d := o.WorkloadType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterAttributes_SdkV2.
func (o *ClusterAttributes_SdkV2) SetWorkloadType(ctx context.Context, v WorkloadType_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workload_type"]
	o.WorkloadType = types.ListValueMust(t, vs)
}

type ClusterCompliance_SdkV2 struct {
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

func (newState *ClusterCompliance_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterCompliance_SdkV2) {
}

func (newState *ClusterCompliance_SdkV2) SyncFieldsDuringRead(existingState ClusterCompliance_SdkV2) {
}

func (c ClusterCompliance_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterCompliance_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterCompliance_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterCompliance_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":   o.ClusterId,
			"is_compliant": o.IsCompliant,
			"violations":   o.Violations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterCompliance_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetViolations returns the value of the Violations field in ClusterCompliance_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterCompliance_SdkV2) GetViolations(ctx context.Context) (map[string]types.String, bool) {
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

// SetViolations sets the value of the Violations field in ClusterCompliance_SdkV2.
func (o *ClusterCompliance_SdkV2) SetViolations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["violations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Violations = types.MapValueMust(t, vs)
}

// Describes all of the metadata about a single Spark cluster in Databricks.
type ClusterDetails_SdkV2 struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
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
	ClusterLogConf types.List `tfsdk:"cluster_log_conf"`
	// Cluster log delivery status.
	ClusterLogStatus types.List `tfsdk:"cluster_log_status"`
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
	DockerImage types.List `tfsdk:"docker_image"`
	// Node on which the Spark driver resides. The driver node contains the
	// Spark master and the Databricks application that manages the per-notebook
	// Spark REPLs.
	Driver types.List `tfsdk:"driver"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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
	Spec types.List `tfsdk:"spec"`
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
	TerminationReason types.List `tfsdk:"termination_reason"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime types.Bool `tfsdk:"use_ml_runtime"`

	WorkloadType types.List `tfsdk:"workload_type"`
}

func (newState *ClusterDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterDetails_SdkV2) {
}

func (newState *ClusterDetails_SdkV2) SyncFieldsDuringRead(existingState ClusterDetails_SdkV2) {
}

func (c ClusterDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_cores"] = attrs["cluster_cores"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_log_status"] = attrs["cluster_log_status"].SetOptional()
	attrs["cluster_log_status"] = attrs["cluster_log_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_memory_mb"] = attrs["cluster_memory_mb"].SetOptional()
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["cluster_source"] = attrs["cluster_source"].SetOptional()
	attrs["creator_user_name"] = attrs["creator_user_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["default_tags"] = attrs["default_tags"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["driver"] = attrs["driver"].SetOptional()
	attrs["driver"] = attrs["driver"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["executors"] = attrs["executors"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["ssh_public_keys"] = attrs["ssh_public_keys"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["state_message"] = attrs["state_message"].SetOptional()
	attrs["terminated_time"] = attrs["terminated_time"].SetOptional()
	attrs["termination_reason"] = attrs["termination_reason"].SetOptional()
	attrs["termination_reason"] = attrs["termination_reason"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["total_initial_remote_disk_size"] = attrs["total_initial_remote_disk_size"].SetOptional()
	attrs["use_ml_runtime"] = attrs["use_ml_runtime"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":          reflect.TypeOf(AutoScale_SdkV2{}),
		"aws_attributes":     reflect.TypeOf(AwsAttributes_SdkV2{}),
		"azure_attributes":   reflect.TypeOf(AzureAttributes_SdkV2{}),
		"cluster_log_conf":   reflect.TypeOf(ClusterLogConf_SdkV2{}),
		"cluster_log_status": reflect.TypeOf(LogSyncStatus_SdkV2{}),
		"custom_tags":        reflect.TypeOf(types.String{}),
		"default_tags":       reflect.TypeOf(types.String{}),
		"docker_image":       reflect.TypeOf(DockerImage_SdkV2{}),
		"driver":             reflect.TypeOf(SparkNode_SdkV2{}),
		"executors":          reflect.TypeOf(SparkNode_SdkV2{}),
		"gcp_attributes":     reflect.TypeOf(GcpAttributes_SdkV2{}),
		"init_scripts":       reflect.TypeOf(InitScriptInfo_SdkV2{}),
		"spark_conf":         reflect.TypeOf(types.String{}),
		"spark_env_vars":     reflect.TypeOf(types.String{}),
		"spec":               reflect.TypeOf(ClusterSpec_SdkV2{}),
		"ssh_public_keys":    reflect.TypeOf(types.String{}),
		"termination_reason": reflect.TypeOf(TerminationReason_SdkV2{}),
		"workload_type":      reflect.TypeOf(WorkloadType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale": basetypes.ListType{
				ElemType: AutoScale_SdkV2{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes_SdkV2{}.Type(ctx),
			},
			"cluster_cores": types.Float64Type,
			"cluster_id":    types.StringType,
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf_SdkV2{}.Type(ctx),
			},
			"cluster_log_status": basetypes.ListType{
				ElemType: LogSyncStatus_SdkV2{}.Type(ctx),
			},
			"cluster_memory_mb": types.Int64Type,
			"cluster_name":      types.StringType,
			"cluster_source":    types.StringType,
			"creator_user_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"default_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"docker_image": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"driver": basetypes.ListType{
				ElemType: SparkNode_SdkV2{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"executors": basetypes.ListType{
				ElemType: SparkNode_SdkV2{}.Type(ctx),
			},
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes_SdkV2{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo_SdkV2{}.Type(ctx),
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
			"spec": basetypes.ListType{
				ElemType: ClusterSpec_SdkV2{}.Type(ctx),
			},
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"start_time":      types.Int64Type,
			"state":           types.StringType,
			"state_message":   types.StringType,
			"terminated_time": types.Int64Type,
			"termination_reason": basetypes.ListType{
				ElemType: TerminationReason_SdkV2{}.Type(ctx),
			},
			"total_initial_remote_disk_size": types.Int64Type,
			"use_ml_runtime":                 types.BoolType,
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in ClusterDetails_SdkV2 as
// a AutoScale_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetAutoscale(ctx context.Context) (AutoScale_SdkV2, bool) {
	var e AutoScale_SdkV2
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale_SdkV2
	d := o.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetAutoscale(ctx context.Context, v AutoScale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	o.Autoscale = types.ListValueMust(t, vs)
}

// GetAwsAttributes returns the value of the AwsAttributes field in ClusterDetails_SdkV2 as
// a AwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetAwsAttributes(ctx context.Context) (AwsAttributes_SdkV2, bool) {
	var e AwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetAwsAttributes(ctx context.Context, v AwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterDetails_SdkV2 as
// a AzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetAzureAttributes(ctx context.Context) (AzureAttributes_SdkV2, bool) {
	var e AzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetAzureAttributes(ctx context.Context, v AzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterDetails_SdkV2 as
// a ClusterLogConf_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetClusterLogConf(ctx context.Context) (ClusterLogConf_SdkV2, bool) {
	var e ClusterLogConf_SdkV2
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf_SdkV2
	d := o.ClusterLogConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetClusterLogConf(ctx context.Context, v ClusterLogConf_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_conf"]
	o.ClusterLogConf = types.ListValueMust(t, vs)
}

// GetClusterLogStatus returns the value of the ClusterLogStatus field in ClusterDetails_SdkV2 as
// a LogSyncStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetClusterLogStatus(ctx context.Context) (LogSyncStatus_SdkV2, bool) {
	var e LogSyncStatus_SdkV2
	if o.ClusterLogStatus.IsNull() || o.ClusterLogStatus.IsUnknown() {
		return e, false
	}
	var v []LogSyncStatus_SdkV2
	d := o.ClusterLogStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogStatus sets the value of the ClusterLogStatus field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetClusterLogStatus(ctx context.Context, v LogSyncStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_status"]
	o.ClusterLogStatus = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in ClusterDetails_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in ClusterDetails_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetDefaultTags sets the value of the DefaultTags field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DefaultTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterDetails_SdkV2 as
// a DockerImage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetDockerImage(ctx context.Context) (DockerImage_SdkV2, bool) {
	var e DockerImage_SdkV2
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage_SdkV2
	d := o.DockerImage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetDockerImage(ctx context.Context, v DockerImage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["docker_image"]
	o.DockerImage = types.ListValueMust(t, vs)
}

// GetDriver returns the value of the Driver field in ClusterDetails_SdkV2 as
// a SparkNode_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetDriver(ctx context.Context) (SparkNode_SdkV2, bool) {
	var e SparkNode_SdkV2
	if o.Driver.IsNull() || o.Driver.IsUnknown() {
		return e, false
	}
	var v []SparkNode_SdkV2
	d := o.Driver.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDriver sets the value of the Driver field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetDriver(ctx context.Context, v SparkNode_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["driver"]
	o.Driver = types.ListValueMust(t, vs)
}

// GetExecutors returns the value of the Executors field in ClusterDetails_SdkV2 as
// a slice of SparkNode_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetExecutors(ctx context.Context) ([]SparkNode_SdkV2, bool) {
	if o.Executors.IsNull() || o.Executors.IsUnknown() {
		return nil, false
	}
	var v []SparkNode_SdkV2
	d := o.Executors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExecutors sets the value of the Executors field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetExecutors(ctx context.Context, v []SparkNode_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["executors"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Executors = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterDetails_SdkV2 as
// a GcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetGcpAttributes(ctx context.Context) (GcpAttributes_SdkV2, bool) {
	var e GcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetGcpAttributes(ctx context.Context, v GcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in ClusterDetails_SdkV2 as
// a slice of InitScriptInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetInitScripts(ctx context.Context) ([]InitScriptInfo_SdkV2, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo_SdkV2
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetInitScripts(ctx context.Context, v []InitScriptInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterDetails_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkConf sets the value of the SparkConf field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterDetails_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSpec returns the value of the Spec field in ClusterDetails_SdkV2 as
// a ClusterSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetSpec(ctx context.Context) (ClusterSpec_SdkV2, bool) {
	var e ClusterSpec_SdkV2
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []ClusterSpec_SdkV2
	d := o.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetSpec(ctx context.Context, v ClusterSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	o.Spec = types.ListValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterDetails_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
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

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetTerminationReason returns the value of the TerminationReason field in ClusterDetails_SdkV2 as
// a TerminationReason_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetTerminationReason(ctx context.Context) (TerminationReason_SdkV2, bool) {
	var e TerminationReason_SdkV2
	if o.TerminationReason.IsNull() || o.TerminationReason.IsUnknown() {
		return e, false
	}
	var v []TerminationReason_SdkV2
	d := o.TerminationReason.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTerminationReason sets the value of the TerminationReason field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetTerminationReason(ctx context.Context, v TerminationReason_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["termination_reason"]
	o.TerminationReason = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterDetails_SdkV2 as
// a WorkloadType_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterDetails_SdkV2) GetWorkloadType(ctx context.Context) (WorkloadType_SdkV2, bool) {
	var e WorkloadType_SdkV2
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType_SdkV2
	d := o.WorkloadType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterDetails_SdkV2.
func (o *ClusterDetails_SdkV2) SetWorkloadType(ctx context.Context, v WorkloadType_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workload_type"]
	o.WorkloadType = types.ListValueMust(t, vs)
}

type ClusterEvent_SdkV2 struct {
	ClusterId types.String `tfsdk:"cluster_id"`

	DataPlaneEventDetails types.List `tfsdk:"data_plane_event_details"`

	Details types.List `tfsdk:"details"`
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	Timestamp types.Int64 `tfsdk:"timestamp"`

	Type_ types.String `tfsdk:"type"`
}

func (newState *ClusterEvent_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterEvent_SdkV2) {
}

func (newState *ClusterEvent_SdkV2) SyncFieldsDuringRead(existingState ClusterEvent_SdkV2) {
}

func (c ClusterEvent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetRequired()
	attrs["data_plane_event_details"] = attrs["data_plane_event_details"].SetOptional()
	attrs["data_plane_event_details"] = attrs["data_plane_event_details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["details"] = attrs["details"].SetOptional()
	attrs["details"] = attrs["details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a ClusterEvent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_plane_event_details": reflect.TypeOf(DataPlaneEventDetails_SdkV2{}),
		"details":                  reflect.TypeOf(EventDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterEvent_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterEvent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterEvent_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"data_plane_event_details": basetypes.ListType{
				ElemType: DataPlaneEventDetails_SdkV2{}.Type(ctx),
			},
			"details": basetypes.ListType{
				ElemType: EventDetails_SdkV2{}.Type(ctx),
			},
			"timestamp": types.Int64Type,
			"type":      types.StringType,
		},
	}
}

// GetDataPlaneEventDetails returns the value of the DataPlaneEventDetails field in ClusterEvent_SdkV2 as
// a DataPlaneEventDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterEvent_SdkV2) GetDataPlaneEventDetails(ctx context.Context) (DataPlaneEventDetails_SdkV2, bool) {
	var e DataPlaneEventDetails_SdkV2
	if o.DataPlaneEventDetails.IsNull() || o.DataPlaneEventDetails.IsUnknown() {
		return e, false
	}
	var v []DataPlaneEventDetails_SdkV2
	d := o.DataPlaneEventDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataPlaneEventDetails sets the value of the DataPlaneEventDetails field in ClusterEvent_SdkV2.
func (o *ClusterEvent_SdkV2) SetDataPlaneEventDetails(ctx context.Context, v DataPlaneEventDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_plane_event_details"]
	o.DataPlaneEventDetails = types.ListValueMust(t, vs)
}

// GetDetails returns the value of the Details field in ClusterEvent_SdkV2 as
// a EventDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterEvent_SdkV2) GetDetails(ctx context.Context) (EventDetails_SdkV2, bool) {
	var e EventDetails_SdkV2
	if o.Details.IsNull() || o.Details.IsUnknown() {
		return e, false
	}
	var v []EventDetails_SdkV2
	d := o.Details.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDetails sets the value of the Details field in ClusterEvent_SdkV2.
func (o *ClusterEvent_SdkV2) SetDetails(ctx context.Context, v EventDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["details"]
	o.Details = types.ListValueMust(t, vs)
}

type ClusterLibraryStatuses_SdkV2 struct {
	// Unique identifier for the cluster.
	ClusterId types.String `tfsdk:"cluster_id"`
	// Status of all libraries on the cluster.
	LibraryStatuses types.List `tfsdk:"library_statuses"`
}

func (newState *ClusterLibraryStatuses_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterLibraryStatuses_SdkV2) {
}

func (newState *ClusterLibraryStatuses_SdkV2) SyncFieldsDuringRead(existingState ClusterLibraryStatuses_SdkV2) {
}

func (c ClusterLibraryStatuses_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterLibraryStatuses_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"library_statuses": reflect.TypeOf(LibraryFullStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterLibraryStatuses_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterLibraryStatuses_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":       o.ClusterId,
			"library_statuses": o.LibraryStatuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterLibraryStatuses_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"library_statuses": basetypes.ListType{
				ElemType: LibraryFullStatus_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetLibraryStatuses returns the value of the LibraryStatuses field in ClusterLibraryStatuses_SdkV2 as
// a slice of LibraryFullStatus_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterLibraryStatuses_SdkV2) GetLibraryStatuses(ctx context.Context) ([]LibraryFullStatus_SdkV2, bool) {
	if o.LibraryStatuses.IsNull() || o.LibraryStatuses.IsUnknown() {
		return nil, false
	}
	var v []LibraryFullStatus_SdkV2
	d := o.LibraryStatuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraryStatuses sets the value of the LibraryStatuses field in ClusterLibraryStatuses_SdkV2.
func (o *ClusterLibraryStatuses_SdkV2) SetLibraryStatuses(ctx context.Context, v []LibraryFullStatus_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["library_statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LibraryStatuses = types.ListValueMust(t, vs)
}

// Cluster log delivery config
type ClusterLogConf_SdkV2 struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs types.List `tfsdk:"dbfs"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 types.List `tfsdk:"s3"`
	// destination needs to be provided, e.g. `{ "volumes": { "destination":
	// "/Volumes/catalog/schema/volume/cluster_log" } }`
	Volumes types.List `tfsdk:"volumes"`
}

func (newState *ClusterLogConf_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterLogConf_SdkV2) {
}

func (newState *ClusterLogConf_SdkV2) SyncFieldsDuringRead(existingState ClusterLogConf_SdkV2) {
}

func (c ClusterLogConf_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dbfs"] = attrs["dbfs"].SetOptional()
	attrs["dbfs"] = attrs["dbfs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["s3"] = attrs["s3"].SetOptional()
	attrs["s3"] = attrs["s3"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["volumes"] = attrs["volumes"].SetOptional()
	attrs["volumes"] = attrs["volumes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterLogConf.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterLogConf_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbfs":    reflect.TypeOf(DbfsStorageInfo_SdkV2{}),
		"s3":      reflect.TypeOf(S3StorageInfo_SdkV2{}),
		"volumes": reflect.TypeOf(VolumesStorageInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterLogConf_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterLogConf_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbfs":    o.Dbfs,
			"s3":      o.S3,
			"volumes": o.Volumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterLogConf_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbfs": basetypes.ListType{
				ElemType: DbfsStorageInfo_SdkV2{}.Type(ctx),
			},
			"s3": basetypes.ListType{
				ElemType: S3StorageInfo_SdkV2{}.Type(ctx),
			},
			"volumes": basetypes.ListType{
				ElemType: VolumesStorageInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDbfs returns the value of the Dbfs field in ClusterLogConf_SdkV2 as
// a DbfsStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterLogConf_SdkV2) GetDbfs(ctx context.Context) (DbfsStorageInfo_SdkV2, bool) {
	var e DbfsStorageInfo_SdkV2
	if o.Dbfs.IsNull() || o.Dbfs.IsUnknown() {
		return e, false
	}
	var v []DbfsStorageInfo_SdkV2
	d := o.Dbfs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbfs sets the value of the Dbfs field in ClusterLogConf_SdkV2.
func (o *ClusterLogConf_SdkV2) SetDbfs(ctx context.Context, v DbfsStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbfs"]
	o.Dbfs = types.ListValueMust(t, vs)
}

// GetS3 returns the value of the S3 field in ClusterLogConf_SdkV2 as
// a S3StorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterLogConf_SdkV2) GetS3(ctx context.Context) (S3StorageInfo_SdkV2, bool) {
	var e S3StorageInfo_SdkV2
	if o.S3.IsNull() || o.S3.IsUnknown() {
		return e, false
	}
	var v []S3StorageInfo_SdkV2
	d := o.S3.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetS3 sets the value of the S3 field in ClusterLogConf_SdkV2.
func (o *ClusterLogConf_SdkV2) SetS3(ctx context.Context, v S3StorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["s3"]
	o.S3 = types.ListValueMust(t, vs)
}

// GetVolumes returns the value of the Volumes field in ClusterLogConf_SdkV2 as
// a VolumesStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterLogConf_SdkV2) GetVolumes(ctx context.Context) (VolumesStorageInfo_SdkV2, bool) {
	var e VolumesStorageInfo_SdkV2
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return e, false
	}
	var v []VolumesStorageInfo_SdkV2
	d := o.Volumes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumes sets the value of the Volumes field in ClusterLogConf_SdkV2.
func (o *ClusterLogConf_SdkV2) SetVolumes(ctx context.Context, v VolumesStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["volumes"]
	o.Volumes = types.ListValueMust(t, vs)
}

type ClusterPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ClusterPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterPermission_SdkV2) {
}

func (newState *ClusterPermission_SdkV2) SyncFieldsDuringRead(existingState ClusterPermission_SdkV2) {
}

func (c ClusterPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in ClusterPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in ClusterPermission_SdkV2.
func (o *ClusterPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type ClusterPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *ClusterPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterPermissions_SdkV2) {
}

func (newState *ClusterPermissions_SdkV2) SyncFieldsDuringRead(existingState ClusterPermissions_SdkV2) {
}

func (c ClusterPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ClusterAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ClusterPermissions_SdkV2 as
// a slice of ClusterAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]ClusterAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPermissions_SdkV2.
func (o *ClusterPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []ClusterAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ClusterPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ClusterPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterPermissionsDescription_SdkV2) {
}

func (newState *ClusterPermissionsDescription_SdkV2) SyncFieldsDuringRead(existingState ClusterPermissionsDescription_SdkV2) {
}

func (c ClusterPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ClusterPermissionsRequest_SdkV2 struct {
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
func (a ClusterPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"cluster_id":          o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ClusterAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"cluster_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ClusterPermissionsRequest_SdkV2 as
// a slice of ClusterAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]ClusterAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPermissionsRequest_SdkV2.
func (o *ClusterPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []ClusterAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ClusterPolicyAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *ClusterPolicyAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterPolicyAccessControlRequest_SdkV2) {
}

func (newState *ClusterPolicyAccessControlRequest_SdkV2) SyncFieldsDuringRead(existingState ClusterPolicyAccessControlRequest_SdkV2) {
}

func (c ClusterPolicyAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPolicyAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterPolicyAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ClusterPolicyAccessControlResponse_SdkV2 struct {
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

func (newState *ClusterPolicyAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterPolicyAccessControlResponse_SdkV2) {
}

func (newState *ClusterPolicyAccessControlResponse_SdkV2) SyncFieldsDuringRead(existingState ClusterPolicyAccessControlResponse_SdkV2) {
}

func (c ClusterPolicyAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ClusterPolicyPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPolicyAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterPolicyAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ClusterPolicyPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in ClusterPolicyAccessControlResponse_SdkV2 as
// a slice of ClusterPolicyPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterPolicyAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]ClusterPolicyPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ClusterPolicyAccessControlResponse_SdkV2.
func (o *ClusterPolicyAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []ClusterPolicyPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type ClusterPolicyPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ClusterPolicyPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterPolicyPermission_SdkV2) {
}

func (newState *ClusterPolicyPermission_SdkV2) SyncFieldsDuringRead(existingState ClusterPolicyPermission_SdkV2) {
}

func (c ClusterPolicyPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPolicyPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in ClusterPolicyPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterPolicyPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in ClusterPolicyPermission_SdkV2.
func (o *ClusterPolicyPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type ClusterPolicyPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *ClusterPolicyPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterPolicyPermissions_SdkV2) {
}

func (newState *ClusterPolicyPermissions_SdkV2) SyncFieldsDuringRead(existingState ClusterPolicyPermissions_SdkV2) {
}

func (c ClusterPolicyPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterPolicyAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPolicyPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ClusterPolicyAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ClusterPolicyPermissions_SdkV2 as
// a slice of ClusterPolicyAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterPolicyPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]ClusterPolicyAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPolicyPermissions_SdkV2.
func (o *ClusterPolicyPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []ClusterPolicyAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ClusterPolicyPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ClusterPolicyPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterPolicyPermissionsDescription_SdkV2) {
}

func (newState *ClusterPolicyPermissionsDescription_SdkV2) SyncFieldsDuringRead(existingState ClusterPolicyPermissionsDescription_SdkV2) {
}

func (c ClusterPolicyPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterPolicyPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPolicyPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ClusterPolicyPermissionsRequest_SdkV2 struct {
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
func (a ClusterPolicyPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ClusterPolicyAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterPolicyPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterPolicyPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"cluster_policy_id":   o.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ClusterPolicyAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"cluster_policy_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ClusterPolicyPermissionsRequest_SdkV2 as
// a slice of ClusterPolicyAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterPolicyPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]ClusterPolicyAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ClusterPolicyPermissionsRequest_SdkV2.
func (o *ClusterPolicyPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []ClusterPolicyAccessControlRequest_SdkV2) {
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
type ClusterSettingsChange_SdkV2 struct {
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

func (newState *ClusterSettingsChange_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterSettingsChange_SdkV2) {
}

func (newState *ClusterSettingsChange_SdkV2) SyncFieldsDuringRead(existingState ClusterSettingsChange_SdkV2) {
}

func (c ClusterSettingsChange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterSettingsChange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSettingsChange_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterSettingsChange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"field":          o.Field,
			"new_value":      o.NewValue,
			"previous_value": o.PreviousValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterSettingsChange_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"field":          types.StringType,
			"new_value":      types.StringType,
			"previous_value": types.StringType,
		},
	}
}

type ClusterSize_SdkV2 struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale"`
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

func (newState *ClusterSize_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterSize_SdkV2) {
}

func (newState *ClusterSize_SdkV2) SyncFieldsDuringRead(existingState ClusterSize_SdkV2) {
}

func (c ClusterSize_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a ClusterSize_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale": reflect.TypeOf(AutoScale_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSize_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterSize_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":   o.Autoscale,
			"num_workers": o.NumWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterSize_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale": basetypes.ListType{
				ElemType: AutoScale_SdkV2{}.Type(ctx),
			},
			"num_workers": types.Int64Type,
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in ClusterSize_SdkV2 as
// a AutoScale_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSize_SdkV2) GetAutoscale(ctx context.Context) (AutoScale_SdkV2, bool) {
	var e AutoScale_SdkV2
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale_SdkV2
	d := o.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in ClusterSize_SdkV2.
func (o *ClusterSize_SdkV2) SetAutoscale(ctx context.Context, v AutoScale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	o.Autoscale = types.ListValueMust(t, vs)
}

// Contains a snapshot of the latest user specified settings that were used to
// create/edit the cluster.
type ClusterSpec_SdkV2 struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf"`
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
	DockerImage types.List `tfsdk:"docker_image"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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

	WorkloadType types.List `tfsdk:"workload_type"`
}

func (newState *ClusterSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ClusterSpec_SdkV2) {
}

func (newState *ClusterSpec_SdkV2) SyncFieldsDuringRead(existingState ClusterSpec_SdkV2) {
}

func (c ClusterSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_policy_default_values"] = attrs["apply_policy_default_values"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
	attrs["workload_type"] = attrs["workload_type"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(AutoScale_SdkV2{}),
		"aws_attributes":   reflect.TypeOf(AwsAttributes_SdkV2{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes_SdkV2{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf_SdkV2{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage_SdkV2{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes_SdkV2{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo_SdkV2{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale": basetypes.ListType{
				ElemType: AutoScale_SdkV2{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes_SdkV2{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf_SdkV2{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes_SdkV2{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo_SdkV2{}.Type(ctx),
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in ClusterSpec_SdkV2 as
// a AutoScale_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetAutoscale(ctx context.Context) (AutoScale_SdkV2, bool) {
	var e AutoScale_SdkV2
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale_SdkV2
	d := o.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetAutoscale(ctx context.Context, v AutoScale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	o.Autoscale = types.ListValueMust(t, vs)
}

// GetAwsAttributes returns the value of the AwsAttributes field in ClusterSpec_SdkV2 as
// a AwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetAwsAttributes(ctx context.Context) (AwsAttributes_SdkV2, bool) {
	var e AwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetAwsAttributes(ctx context.Context, v AwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in ClusterSpec_SdkV2 as
// a AzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetAzureAttributes(ctx context.Context) (AzureAttributes_SdkV2, bool) {
	var e AzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetAzureAttributes(ctx context.Context, v AzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetClusterLogConf returns the value of the ClusterLogConf field in ClusterSpec_SdkV2 as
// a ClusterLogConf_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetClusterLogConf(ctx context.Context) (ClusterLogConf_SdkV2, bool) {
	var e ClusterLogConf_SdkV2
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf_SdkV2
	d := o.ClusterLogConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetClusterLogConf(ctx context.Context, v ClusterLogConf_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_conf"]
	o.ClusterLogConf = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in ClusterSpec_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in ClusterSpec_SdkV2 as
// a DockerImage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetDockerImage(ctx context.Context) (DockerImage_SdkV2, bool) {
	var e DockerImage_SdkV2
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage_SdkV2
	d := o.DockerImage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetDockerImage(ctx context.Context, v DockerImage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["docker_image"]
	o.DockerImage = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in ClusterSpec_SdkV2 as
// a GcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetGcpAttributes(ctx context.Context) (GcpAttributes_SdkV2, bool) {
	var e GcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetGcpAttributes(ctx context.Context, v GcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in ClusterSpec_SdkV2 as
// a slice of InitScriptInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetInitScripts(ctx context.Context) ([]InitScriptInfo_SdkV2, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo_SdkV2
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetInitScripts(ctx context.Context, v []InitScriptInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in ClusterSpec_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkConf sets the value of the SparkConf field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in ClusterSpec_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkEnvVars sets the value of the SparkEnvVars field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in ClusterSpec_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
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

// SetSshPublicKeys sets the value of the SshPublicKeys field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in ClusterSpec_SdkV2 as
// a WorkloadType_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetWorkloadType(ctx context.Context) (WorkloadType_SdkV2, bool) {
	var e WorkloadType_SdkV2
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType_SdkV2
	d := o.WorkloadType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetWorkloadType(ctx context.Context, v WorkloadType_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workload_type"]
	o.WorkloadType = types.ListValueMust(t, vs)
}

type ClusterStatus_SdkV2 struct {
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
func (a ClusterStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type Command_SdkV2 struct {
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
func (a Command_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Command_SdkV2
// only implements ToObjectValue() and Type().
func (o Command_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Command_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"command":   types.StringType,
			"contextId": types.StringType,
			"language":  types.StringType,
		},
	}
}

type CommandStatusRequest_SdkV2 struct {
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
func (a CommandStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommandStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CommandStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"commandId": o.CommandId,
			"contextId": o.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CommandStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"commandId": types.StringType,
			"contextId": types.StringType,
		},
	}
}

type CommandStatusResponse_SdkV2 struct {
	Id types.String `tfsdk:"id"`

	Results types.List `tfsdk:"results"`

	Status types.String `tfsdk:"status"`
}

func (newState *CommandStatusResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CommandStatusResponse_SdkV2) {
}

func (newState *CommandStatusResponse_SdkV2) SyncFieldsDuringRead(existingState CommandStatusResponse_SdkV2) {
}

func (c CommandStatusResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()
	attrs["results"] = attrs["results"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a CommandStatusResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Results_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommandStatusResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CommandStatusResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":      o.Id,
			"results": o.Results,
			"status":  o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CommandStatusResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"results": basetypes.ListType{
				ElemType: Results_SdkV2{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// GetResults returns the value of the Results field in CommandStatusResponse_SdkV2 as
// a Results_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CommandStatusResponse_SdkV2) GetResults(ctx context.Context) (Results_SdkV2, bool) {
	var e Results_SdkV2
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return e, false
	}
	var v []Results_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResults sets the value of the Results field in CommandStatusResponse_SdkV2.
func (o *CommandStatusResponse_SdkV2) SetResults(ctx context.Context, v Results_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	o.Results = types.ListValueMust(t, vs)
}

type ContextStatusRequest_SdkV2 struct {
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
func (a ContextStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContextStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ContextStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"contextId": o.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ContextStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"contextId": types.StringType,
		},
	}
}

type ContextStatusResponse_SdkV2 struct {
	Id types.String `tfsdk:"id"`

	Status types.String `tfsdk:"status"`
}

func (newState *ContextStatusResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ContextStatusResponse_SdkV2) {
}

func (newState *ContextStatusResponse_SdkV2) SyncFieldsDuringRead(existingState ContextStatusResponse_SdkV2) {
}

func (c ContextStatusResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ContextStatusResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContextStatusResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ContextStatusResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":     o.Id,
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ContextStatusResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":     types.StringType,
			"status": types.StringType,
		},
	}
}

type CreateCluster_SdkV2 struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	CloneFrom types.List `tfsdk:"clone_from"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf"`
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
	DockerImage types.List `tfsdk:"docker_image"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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

	WorkloadType types.List `tfsdk:"workload_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(AutoScale_SdkV2{}),
		"aws_attributes":   reflect.TypeOf(AwsAttributes_SdkV2{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes_SdkV2{}),
		"clone_from":       reflect.TypeOf(CloneCluster_SdkV2{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf_SdkV2{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage_SdkV2{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes_SdkV2{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo_SdkV2{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale": basetypes.ListType{
				ElemType: AutoScale_SdkV2{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes_SdkV2{}.Type(ctx),
			},
			"clone_from": basetypes.ListType{
				ElemType: CloneCluster_SdkV2{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf_SdkV2{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes_SdkV2{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo_SdkV2{}.Type(ctx),
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in CreateCluster_SdkV2 as
// a AutoScale_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetAutoscale(ctx context.Context) (AutoScale_SdkV2, bool) {
	var e AutoScale_SdkV2
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale_SdkV2
	d := o.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetAutoscale(ctx context.Context, v AutoScale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	o.Autoscale = types.ListValueMust(t, vs)
}

// GetAwsAttributes returns the value of the AwsAttributes field in CreateCluster_SdkV2 as
// a AwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetAwsAttributes(ctx context.Context) (AwsAttributes_SdkV2, bool) {
	var e AwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetAwsAttributes(ctx context.Context, v AwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in CreateCluster_SdkV2 as
// a AzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetAzureAttributes(ctx context.Context) (AzureAttributes_SdkV2, bool) {
	var e AzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetAzureAttributes(ctx context.Context, v AzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetCloneFrom returns the value of the CloneFrom field in CreateCluster_SdkV2 as
// a CloneCluster_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetCloneFrom(ctx context.Context) (CloneCluster_SdkV2, bool) {
	var e CloneCluster_SdkV2
	if o.CloneFrom.IsNull() || o.CloneFrom.IsUnknown() {
		return e, false
	}
	var v []CloneCluster_SdkV2
	d := o.CloneFrom.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloneFrom sets the value of the CloneFrom field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetCloneFrom(ctx context.Context, v CloneCluster_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clone_from"]
	o.CloneFrom = types.ListValueMust(t, vs)
}

// GetClusterLogConf returns the value of the ClusterLogConf field in CreateCluster_SdkV2 as
// a ClusterLogConf_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetClusterLogConf(ctx context.Context) (ClusterLogConf_SdkV2, bool) {
	var e ClusterLogConf_SdkV2
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf_SdkV2
	d := o.ClusterLogConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetClusterLogConf(ctx context.Context, v ClusterLogConf_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_conf"]
	o.ClusterLogConf = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in CreateCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in CreateCluster_SdkV2 as
// a DockerImage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetDockerImage(ctx context.Context) (DockerImage_SdkV2, bool) {
	var e DockerImage_SdkV2
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage_SdkV2
	d := o.DockerImage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetDockerImage(ctx context.Context, v DockerImage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["docker_image"]
	o.DockerImage = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in CreateCluster_SdkV2 as
// a GcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetGcpAttributes(ctx context.Context) (GcpAttributes_SdkV2, bool) {
	var e GcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetGcpAttributes(ctx context.Context, v GcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in CreateCluster_SdkV2 as
// a slice of InitScriptInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetInitScripts(ctx context.Context) ([]InitScriptInfo_SdkV2, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo_SdkV2
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetInitScripts(ctx context.Context, v []InitScriptInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in CreateCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkConf sets the value of the SparkConf field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in CreateCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkEnvVars sets the value of the SparkEnvVars field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in CreateCluster_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
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

// SetSshPublicKeys sets the value of the SshPublicKeys field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in CreateCluster_SdkV2 as
// a WorkloadType_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCluster_SdkV2) GetWorkloadType(ctx context.Context) (WorkloadType_SdkV2, bool) {
	var e WorkloadType_SdkV2
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType_SdkV2
	d := o.WorkloadType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in CreateCluster_SdkV2.
func (o *CreateCluster_SdkV2) SetWorkloadType(ctx context.Context, v WorkloadType_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workload_type"]
	o.WorkloadType = types.ListValueMust(t, vs)
}

type CreateClusterResponse_SdkV2 struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

func (newState *CreateClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CreateClusterResponse_SdkV2) {
}

func (newState *CreateClusterResponse_SdkV2) SyncFieldsDuringRead(existingState CreateClusterResponse_SdkV2) {
}

func (c CreateClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type CreateContext_SdkV2 struct {
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
func (a CreateContext_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateContext_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateContext_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"language":  o.Language,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateContext_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"language":  types.StringType,
		},
	}
}

type CreateDefaultBaseEnvironmentRequest_SdkV2 struct {
	DefaultBaseEnvironment types.List `tfsdk:"default_base_environment"`
	// A unique identifier for this request. A random UUID is recommended. This
	// request is only idempotent if a `request_id` is provided.
	RequestId types.String `tfsdk:"request_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDefaultBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_base_environment": reflect.TypeOf(DefaultBaseEnvironment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDefaultBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateDefaultBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_base_environment": o.DefaultBaseEnvironment,
			"request_id":               o.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDefaultBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_base_environment": basetypes.ListType{
				ElemType: DefaultBaseEnvironment_SdkV2{}.Type(ctx),
			},
			"request_id": types.StringType,
		},
	}
}

// GetDefaultBaseEnvironment returns the value of the DefaultBaseEnvironment field in CreateDefaultBaseEnvironmentRequest_SdkV2 as
// a DefaultBaseEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDefaultBaseEnvironmentRequest_SdkV2) GetDefaultBaseEnvironment(ctx context.Context) (DefaultBaseEnvironment_SdkV2, bool) {
	var e DefaultBaseEnvironment_SdkV2
	if o.DefaultBaseEnvironment.IsNull() || o.DefaultBaseEnvironment.IsUnknown() {
		return e, false
	}
	var v []DefaultBaseEnvironment_SdkV2
	d := o.DefaultBaseEnvironment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultBaseEnvironment sets the value of the DefaultBaseEnvironment field in CreateDefaultBaseEnvironmentRequest_SdkV2.
func (o *CreateDefaultBaseEnvironmentRequest_SdkV2) SetDefaultBaseEnvironment(ctx context.Context, v DefaultBaseEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_base_environment"]
	o.DefaultBaseEnvironment = types.ListValueMust(t, vs)
}

type CreateInstancePool_SdkV2 struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags types.Map `tfsdk:"custom_tags"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec types.List `tfsdk:"disk_spec"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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
	NodeTypeFlexibility types.List `tfsdk:"node_type_flexibility"`
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
func (a CreateInstancePool_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes_SdkV2{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes_SdkV2{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec_SdkV2{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes_SdkV2{}),
		"node_type_flexibility":    reflect.TypeOf(NodeTypeFlexibility_SdkV2{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage_SdkV2{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstancePool_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateInstancePool_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        o.AwsAttributes,
			"azure_attributes":                      o.AzureAttributes,
			"custom_tags":                           o.CustomTags,
			"disk_spec":                             o.DiskSpec,
			"enable_auto_alternate_node_types":      o.EnableAutoAlternateNodeTypes,
			"enable_elastic_disk":                   o.EnableElasticDisk,
			"gcp_attributes":                        o.GcpAttributes,
			"idle_instance_autotermination_minutes": o.IdleInstanceAutoterminationMinutes,
			"instance_pool_name":                    o.InstancePoolName,
			"max_capacity":                          o.MaxCapacity,
			"min_idle_instances":                    o.MinIdleInstances,
			"node_type_flexibility":                 o.NodeTypeFlexibility,
			"node_type_id":                          o.NodeTypeId,
			"preloaded_docker_images":               o.PreloadedDockerImages,
			"preloaded_spark_versions":              o.PreloadedSparkVersions,
			"remote_disk_throughput":                o.RemoteDiskThroughput,
			"total_initial_remote_disk_size":        o.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateInstancePool_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes": basetypes.ListType{
				ElemType: InstancePoolAwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: InstancePoolAzureAttributes_SdkV2{}.Type(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec": basetypes.ListType{
				ElemType: DiskSpec_SdkV2{}.Type(ctx),
			},
			"enable_auto_alternate_node_types": types.BoolType,
			"enable_elastic_disk":              types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: InstancePoolGcpAttributes_SdkV2{}.Type(ctx),
			},
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
			"node_type_flexibility": basetypes.ListType{
				ElemType: NodeTypeFlexibility_SdkV2{}.Type(ctx),
			},
			"node_type_id": types.StringType,
			"preloaded_docker_images": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"preloaded_spark_versions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"remote_disk_throughput":         types.Int64Type,
			"total_initial_remote_disk_size": types.Int64Type,
		},
	}
}

// GetAwsAttributes returns the value of the AwsAttributes field in CreateInstancePool_SdkV2 as
// a InstancePoolAwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool_SdkV2) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes_SdkV2, bool) {
	var e InstancePoolAwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in CreateInstancePool_SdkV2.
func (o *CreateInstancePool_SdkV2) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in CreateInstancePool_SdkV2 as
// a InstancePoolAzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool_SdkV2) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes_SdkV2, bool) {
	var e InstancePoolAzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in CreateInstancePool_SdkV2.
func (o *CreateInstancePool_SdkV2) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in CreateInstancePool_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in CreateInstancePool_SdkV2.
func (o *CreateInstancePool_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in CreateInstancePool_SdkV2 as
// a DiskSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool_SdkV2) GetDiskSpec(ctx context.Context) (DiskSpec_SdkV2, bool) {
	var e DiskSpec_SdkV2
	if o.DiskSpec.IsNull() || o.DiskSpec.IsUnknown() {
		return e, false
	}
	var v []DiskSpec_SdkV2
	d := o.DiskSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDiskSpec sets the value of the DiskSpec field in CreateInstancePool_SdkV2.
func (o *CreateInstancePool_SdkV2) SetDiskSpec(ctx context.Context, v DiskSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["disk_spec"]
	o.DiskSpec = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in CreateInstancePool_SdkV2 as
// a InstancePoolGcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool_SdkV2) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes_SdkV2, bool) {
	var e InstancePoolGcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolGcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in CreateInstancePool_SdkV2.
func (o *CreateInstancePool_SdkV2) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetNodeTypeFlexibility returns the value of the NodeTypeFlexibility field in CreateInstancePool_SdkV2 as
// a NodeTypeFlexibility_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool_SdkV2) GetNodeTypeFlexibility(ctx context.Context) (NodeTypeFlexibility_SdkV2, bool) {
	var e NodeTypeFlexibility_SdkV2
	if o.NodeTypeFlexibility.IsNull() || o.NodeTypeFlexibility.IsUnknown() {
		return e, false
	}
	var v []NodeTypeFlexibility_SdkV2
	d := o.NodeTypeFlexibility.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeTypeFlexibility sets the value of the NodeTypeFlexibility field in CreateInstancePool_SdkV2.
func (o *CreateInstancePool_SdkV2) SetNodeTypeFlexibility(ctx context.Context, v NodeTypeFlexibility_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_type_flexibility"]
	o.NodeTypeFlexibility = types.ListValueMust(t, vs)
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in CreateInstancePool_SdkV2 as
// a slice of DockerImage_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool_SdkV2) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage_SdkV2, bool) {
	if o.PreloadedDockerImages.IsNull() || o.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage_SdkV2
	d := o.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in CreateInstancePool_SdkV2.
func (o *CreateInstancePool_SdkV2) SetPreloadedDockerImages(ctx context.Context, v []DockerImage_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in CreateInstancePool_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateInstancePool_SdkV2) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
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

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in CreateInstancePool_SdkV2.
func (o *CreateInstancePool_SdkV2) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

type CreateInstancePoolResponse_SdkV2 struct {
	// The ID of the created instance pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
}

func (newState *CreateInstancePoolResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CreateInstancePoolResponse_SdkV2) {
}

func (newState *CreateInstancePoolResponse_SdkV2) SyncFieldsDuringRead(existingState CreateInstancePoolResponse_SdkV2) {
}

func (c CreateInstancePoolResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateInstancePoolResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateInstancePoolResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateInstancePoolResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateInstancePoolResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type CreatePolicy_SdkV2 struct {
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
func (a CreatePolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o CreatePolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreatePolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"definition":  types.StringType,
			"description": types.StringType,
			"libraries": basetypes.ListType{
				ElemType: Library_SdkV2{}.Type(ctx),
			},
			"max_clusters_per_user":              types.Int64Type,
			"name":                               types.StringType,
			"policy_family_definition_overrides": types.StringType,
			"policy_family_id":                   types.StringType,
		},
	}
}

// GetLibraries returns the value of the Libraries field in CreatePolicy_SdkV2 as
// a slice of Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePolicy_SdkV2) GetLibraries(ctx context.Context) ([]Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in CreatePolicy_SdkV2.
func (o *CreatePolicy_SdkV2) SetLibraries(ctx context.Context, v []Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type CreatePolicyResponse_SdkV2 struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId types.String `tfsdk:"policy_id"`
}

func (newState *CreatePolicyResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CreatePolicyResponse_SdkV2) {
}

func (newState *CreatePolicyResponse_SdkV2) SyncFieldsDuringRead(existingState CreatePolicyResponse_SdkV2) {
}

func (c CreatePolicyResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreatePolicyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePolicyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreatePolicyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePolicyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type CreateResponse_SdkV2 struct {
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id"`
}

func (newState *CreateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CreateResponse_SdkV2) {
}

func (newState *CreateResponse_SdkV2) SyncFieldsDuringRead(existingState CreateResponse_SdkV2) {
}

func (c CreateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": o.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"script_id": types.StringType,
		},
	}
}

type Created_SdkV2 struct {
	Id types.String `tfsdk:"id"`
}

func (newState *Created_SdkV2) SyncFieldsDuringCreateOrUpdate(plan Created_SdkV2) {
}

func (newState *Created_SdkV2) SyncFieldsDuringRead(existingState Created_SdkV2) {
}

func (c Created_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Created_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Created_SdkV2
// only implements ToObjectValue() and Type().
func (o Created_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Created_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CustomPolicyTag_SdkV2 struct {
	// The key of the tag. - Must be unique among all custom tags of the same
	// policy - Cannot be “budget-policy-name”, “budget-policy-id” or
	// "budget-policy-resolution-result" - these tags are preserved.
	Key types.String `tfsdk:"key"`
	// The value of the tag.
	Value types.String `tfsdk:"value"`
}

func (newState *CustomPolicyTag_SdkV2) SyncFieldsDuringCreateOrUpdate(plan CustomPolicyTag_SdkV2) {
}

func (newState *CustomPolicyTag_SdkV2) SyncFieldsDuringRead(existingState CustomPolicyTag_SdkV2) {
}

func (c CustomPolicyTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CustomPolicyTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomPolicyTag_SdkV2
// only implements ToObjectValue() and Type().
func (o CustomPolicyTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomPolicyTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type DataPlaneEventDetails_SdkV2 struct {
	EventType types.String `tfsdk:"event_type"`

	ExecutorFailures types.Int64 `tfsdk:"executor_failures"`

	HostId types.String `tfsdk:"host_id"`

	Timestamp types.Int64 `tfsdk:"timestamp"`
}

func (newState *DataPlaneEventDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DataPlaneEventDetails_SdkV2) {
}

func (newState *DataPlaneEventDetails_SdkV2) SyncFieldsDuringRead(existingState DataPlaneEventDetails_SdkV2) {
}

func (c DataPlaneEventDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DataPlaneEventDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneEventDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o DataPlaneEventDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DataPlaneEventDetails_SdkV2) Type(ctx context.Context) attr.Type {
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
type DbfsStorageInfo_SdkV2 struct {
	// dbfs destination, e.g. `dbfs:/my/path`
	Destination types.String `tfsdk:"destination"`
}

func (newState *DbfsStorageInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DbfsStorageInfo_SdkV2) {
}

func (newState *DbfsStorageInfo_SdkV2) SyncFieldsDuringRead(existingState DbfsStorageInfo_SdkV2) {
}

func (c DbfsStorageInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DbfsStorageInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DbfsStorageInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o DbfsStorageInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DbfsStorageInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type DefaultBaseEnvironment_SdkV2 struct {
	BaseEnvironmentCache types.List `tfsdk:"base_environment_cache"`

	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp"`

	CreatorUserId types.Int64 `tfsdk:"creator_user_id"`
	// Note: we made `environment` non-internal because we need to expose its
	// `client` field. All other fields should be treated as internal.
	Environment types.List `tfsdk:"environment"`

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

func (newState *DefaultBaseEnvironment_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DefaultBaseEnvironment_SdkV2) {
}

func (newState *DefaultBaseEnvironment_SdkV2) SyncFieldsDuringRead(existingState DefaultBaseEnvironment_SdkV2) {
}

func (c DefaultBaseEnvironment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["base_environment_cache"] = attrs["base_environment_cache"].SetOptional()
	attrs["created_timestamp"] = attrs["created_timestamp"].SetOptional()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetOptional()
	attrs["environment"] = attrs["environment"].SetOptional()
	attrs["environment"] = attrs["environment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["filepath"] = attrs["filepath"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["is_default"] = attrs["is_default"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["last_updated_user_id"] = attrs["last_updated_user_id"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["principal_ids"] = attrs["principal_ids"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultBaseEnvironment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DefaultBaseEnvironment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"base_environment_cache": reflect.TypeOf(DefaultBaseEnvironmentCache_SdkV2{}),
		"environment":            reflect.TypeOf(Environment_SdkV2{}),
		"principal_ids":          reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultBaseEnvironment_SdkV2
// only implements ToObjectValue() and Type().
func (o DefaultBaseEnvironment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"base_environment_cache": o.BaseEnvironmentCache,
			"created_timestamp":      o.CreatedTimestamp,
			"creator_user_id":        o.CreatorUserId,
			"environment":            o.Environment,
			"filepath":               o.Filepath,
			"id":                     o.Id,
			"is_default":             o.IsDefault,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"last_updated_user_id":   o.LastUpdatedUserId,
			"message":                o.Message,
			"name":                   o.Name,
			"principal_ids":          o.PrincipalIds,
			"status":                 o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DefaultBaseEnvironment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_environment_cache": basetypes.ListType{
				ElemType: DefaultBaseEnvironmentCache_SdkV2{}.Type(ctx),
			},
			"created_timestamp": types.Int64Type,
			"creator_user_id":   types.Int64Type,
			"environment": basetypes.ListType{
				ElemType: Environment_SdkV2{}.Type(ctx),
			},
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

// GetBaseEnvironmentCache returns the value of the BaseEnvironmentCache field in DefaultBaseEnvironment_SdkV2 as
// a slice of DefaultBaseEnvironmentCache_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DefaultBaseEnvironment_SdkV2) GetBaseEnvironmentCache(ctx context.Context) ([]DefaultBaseEnvironmentCache_SdkV2, bool) {
	if o.BaseEnvironmentCache.IsNull() || o.BaseEnvironmentCache.IsUnknown() {
		return nil, false
	}
	var v []DefaultBaseEnvironmentCache_SdkV2
	d := o.BaseEnvironmentCache.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBaseEnvironmentCache sets the value of the BaseEnvironmentCache field in DefaultBaseEnvironment_SdkV2.
func (o *DefaultBaseEnvironment_SdkV2) SetBaseEnvironmentCache(ctx context.Context, v []DefaultBaseEnvironmentCache_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["base_environment_cache"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.BaseEnvironmentCache = types.ListValueMust(t, vs)
}

// GetEnvironment returns the value of the Environment field in DefaultBaseEnvironment_SdkV2 as
// a Environment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *DefaultBaseEnvironment_SdkV2) GetEnvironment(ctx context.Context) (Environment_SdkV2, bool) {
	var e Environment_SdkV2
	if o.Environment.IsNull() || o.Environment.IsUnknown() {
		return e, false
	}
	var v []Environment_SdkV2
	d := o.Environment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnvironment sets the value of the Environment field in DefaultBaseEnvironment_SdkV2.
func (o *DefaultBaseEnvironment_SdkV2) SetEnvironment(ctx context.Context, v Environment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	o.Environment = types.ListValueMust(t, vs)
}

// GetPrincipalIds returns the value of the PrincipalIds field in DefaultBaseEnvironment_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DefaultBaseEnvironment_SdkV2) GetPrincipalIds(ctx context.Context) ([]types.Int64, bool) {
	if o.PrincipalIds.IsNull() || o.PrincipalIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.PrincipalIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrincipalIds sets the value of the PrincipalIds field in DefaultBaseEnvironment_SdkV2.
func (o *DefaultBaseEnvironment_SdkV2) SetPrincipalIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["principal_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrincipalIds = types.ListValueMust(t, vs)
}

type DefaultBaseEnvironmentCache_SdkV2 struct {
	MaterializedEnvironment types.List `tfsdk:"materialized_environment"`

	Message types.String `tfsdk:"message"`

	Status types.String `tfsdk:"status"`
}

func (newState *DefaultBaseEnvironmentCache_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DefaultBaseEnvironmentCache_SdkV2) {
}

func (newState *DefaultBaseEnvironmentCache_SdkV2) SyncFieldsDuringRead(existingState DefaultBaseEnvironmentCache_SdkV2) {
}

func (c DefaultBaseEnvironmentCache_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["materialized_environment"] = attrs["materialized_environment"].SetOptional()
	attrs["materialized_environment"] = attrs["materialized_environment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a DefaultBaseEnvironmentCache_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"materialized_environment": reflect.TypeOf(MaterializedEnvironment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultBaseEnvironmentCache_SdkV2
// only implements ToObjectValue() and Type().
func (o DefaultBaseEnvironmentCache_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"materialized_environment": o.MaterializedEnvironment,
			"message":                  o.Message,
			"status":                   o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DefaultBaseEnvironmentCache_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"materialized_environment": basetypes.ListType{
				ElemType: MaterializedEnvironment_SdkV2{}.Type(ctx),
			},
			"message": types.StringType,
			"status":  types.StringType,
		},
	}
}

// GetMaterializedEnvironment returns the value of the MaterializedEnvironment field in DefaultBaseEnvironmentCache_SdkV2 as
// a MaterializedEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *DefaultBaseEnvironmentCache_SdkV2) GetMaterializedEnvironment(ctx context.Context) (MaterializedEnvironment_SdkV2, bool) {
	var e MaterializedEnvironment_SdkV2
	if o.MaterializedEnvironment.IsNull() || o.MaterializedEnvironment.IsUnknown() {
		return e, false
	}
	var v []MaterializedEnvironment_SdkV2
	d := o.MaterializedEnvironment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMaterializedEnvironment sets the value of the MaterializedEnvironment field in DefaultBaseEnvironmentCache_SdkV2.
func (o *DefaultBaseEnvironmentCache_SdkV2) SetMaterializedEnvironment(ctx context.Context, v MaterializedEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["materialized_environment"]
	o.MaterializedEnvironment = types.ListValueMust(t, vs)
}

type DeleteCluster_SdkV2 struct {
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
func (a DeleteCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type DeleteClusterResponse_SdkV2 struct {
}

func (newState *DeleteClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DeleteClusterResponse_SdkV2) {
}

func (newState *DeleteClusterResponse_SdkV2) SyncFieldsDuringRead(existingState DeleteClusterResponse_SdkV2) {
}

func (c DeleteClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteDefaultBaseEnvironmentRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDefaultBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDefaultBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDefaultBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteGlobalInitScriptRequest_SdkV2 struct {
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
func (a DeleteGlobalInitScriptRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGlobalInitScriptRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteGlobalInitScriptRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": o.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteGlobalInitScriptRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"script_id": types.StringType,
		},
	}
}

type DeleteInstancePool_SdkV2 struct {
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
func (a DeleteInstancePool_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstancePool_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteInstancePool_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstancePool_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type DeleteInstancePoolResponse_SdkV2 struct {
}

func (newState *DeleteInstancePoolResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DeleteInstancePoolResponse_SdkV2) {
}

func (newState *DeleteInstancePoolResponse_SdkV2) SyncFieldsDuringRead(existingState DeleteInstancePoolResponse_SdkV2) {
}

func (c DeleteInstancePoolResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteInstancePoolResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteInstancePoolResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteInstancePoolResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteInstancePoolResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeletePolicy_SdkV2 struct {
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
func (a DeletePolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type DeletePolicyResponse_SdkV2 struct {
}

func (newState *DeletePolicyResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DeletePolicyResponse_SdkV2) {
}

func (newState *DeletePolicyResponse_SdkV2) SyncFieldsDuringRead(existingState DeletePolicyResponse_SdkV2) {
}

func (c DeletePolicyResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePolicyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePolicyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePolicyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePolicyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteResponse_SdkV2 struct {
}

func (newState *DeleteResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DeleteResponse_SdkV2) {
}

func (newState *DeleteResponse_SdkV2) SyncFieldsDuringRead(existingState DeleteResponse_SdkV2) {
}

func (c DeleteResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DestroyContext_SdkV2 struct {
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
func (a DestroyContext_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DestroyContext_SdkV2
// only implements ToObjectValue() and Type().
func (o DestroyContext_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusterId": o.ClusterId,
			"contextId": o.ContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DestroyContext_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusterId": types.StringType,
			"contextId": types.StringType,
		},
	}
}

type DestroyResponse_SdkV2 struct {
}

func (newState *DestroyResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DestroyResponse_SdkV2) {
}

func (newState *DestroyResponse_SdkV2) SyncFieldsDuringRead(existingState DestroyResponse_SdkV2) {
}

func (c DestroyResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DestroyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DestroyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DestroyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DestroyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DestroyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Describes the disks that are launched for each instance in the spark cluster.
// For example, if the cluster has 3 instances, each instance is configured to
// launch 2 disks, 100 GiB each, then Databricks will launch a total of 6 disks,
// 100 GiB each, for this cluster.
type DiskSpec_SdkV2 struct {
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
	DiskType types.List `tfsdk:"disk_type"`
}

func (newState *DiskSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DiskSpec_SdkV2) {
}

func (newState *DiskSpec_SdkV2) SyncFieldsDuringRead(existingState DiskSpec_SdkV2) {
}

func (c DiskSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disk_count"] = attrs["disk_count"].SetOptional()
	attrs["disk_iops"] = attrs["disk_iops"].SetOptional()
	attrs["disk_size"] = attrs["disk_size"].SetOptional()
	attrs["disk_throughput"] = attrs["disk_throughput"].SetOptional()
	attrs["disk_type"] = attrs["disk_type"].SetOptional()
	attrs["disk_type"] = attrs["disk_type"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DiskSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DiskSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disk_type": reflect.TypeOf(DiskType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DiskSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o DiskSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DiskSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disk_count":      types.Int64Type,
			"disk_iops":       types.Int64Type,
			"disk_size":       types.Int64Type,
			"disk_throughput": types.Int64Type,
			"disk_type": basetypes.ListType{
				ElemType: DiskType_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDiskType returns the value of the DiskType field in DiskSpec_SdkV2 as
// a DiskType_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *DiskSpec_SdkV2) GetDiskType(ctx context.Context) (DiskType_SdkV2, bool) {
	var e DiskType_SdkV2
	if o.DiskType.IsNull() || o.DiskType.IsUnknown() {
		return e, false
	}
	var v []DiskType_SdkV2
	d := o.DiskType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDiskType sets the value of the DiskType field in DiskSpec_SdkV2.
func (o *DiskSpec_SdkV2) SetDiskType(ctx context.Context, v DiskType_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["disk_type"]
	o.DiskType = types.ListValueMust(t, vs)
}

// Describes the disk type.
type DiskType_SdkV2 struct {
	AzureDiskVolumeType types.String `tfsdk:"azure_disk_volume_type"`

	EbsVolumeType types.String `tfsdk:"ebs_volume_type"`
}

func (newState *DiskType_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DiskType_SdkV2) {
}

func (newState *DiskType_SdkV2) SyncFieldsDuringRead(existingState DiskType_SdkV2) {
}

func (c DiskType_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DiskType_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DiskType_SdkV2
// only implements ToObjectValue() and Type().
func (o DiskType_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"azure_disk_volume_type": o.AzureDiskVolumeType,
			"ebs_volume_type":        o.EbsVolumeType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DiskType_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"azure_disk_volume_type": types.StringType,
			"ebs_volume_type":        types.StringType,
		},
	}
}

type DockerBasicAuth_SdkV2 struct {
	// Password of the user
	Password types.String `tfsdk:"password"`
	// Name of the user
	Username types.String `tfsdk:"username"`
}

func (newState *DockerBasicAuth_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DockerBasicAuth_SdkV2) {
}

func (newState *DockerBasicAuth_SdkV2) SyncFieldsDuringRead(existingState DockerBasicAuth_SdkV2) {
}

func (c DockerBasicAuth_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DockerBasicAuth_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DockerBasicAuth_SdkV2
// only implements ToObjectValue() and Type().
func (o DockerBasicAuth_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"password": o.Password,
			"username": o.Username,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DockerBasicAuth_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"password": types.StringType,
			"username": types.StringType,
		},
	}
}

type DockerImage_SdkV2 struct {
	// Basic auth with username and password
	BasicAuth types.List `tfsdk:"basic_auth"`
	// URL of the docker image.
	Url types.String `tfsdk:"url"`
}

func (newState *DockerImage_SdkV2) SyncFieldsDuringCreateOrUpdate(plan DockerImage_SdkV2) {
}

func (newState *DockerImage_SdkV2) SyncFieldsDuringRead(existingState DockerImage_SdkV2) {
}

func (c DockerImage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["basic_auth"] = attrs["basic_auth"].SetOptional()
	attrs["basic_auth"] = attrs["basic_auth"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a DockerImage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"basic_auth": reflect.TypeOf(DockerBasicAuth_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DockerImage_SdkV2
// only implements ToObjectValue() and Type().
func (o DockerImage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"basic_auth": o.BasicAuth,
			"url":        o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DockerImage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"basic_auth": basetypes.ListType{
				ElemType: DockerBasicAuth_SdkV2{}.Type(ctx),
			},
			"url": types.StringType,
		},
	}
}

// GetBasicAuth returns the value of the BasicAuth field in DockerImage_SdkV2 as
// a DockerBasicAuth_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *DockerImage_SdkV2) GetBasicAuth(ctx context.Context) (DockerBasicAuth_SdkV2, bool) {
	var e DockerBasicAuth_SdkV2
	if o.BasicAuth.IsNull() || o.BasicAuth.IsUnknown() {
		return e, false
	}
	var v []DockerBasicAuth_SdkV2
	d := o.BasicAuth.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBasicAuth sets the value of the BasicAuth field in DockerImage_SdkV2.
func (o *DockerImage_SdkV2) SetBasicAuth(ctx context.Context, v DockerBasicAuth_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["basic_auth"]
	o.BasicAuth = types.ListValueMust(t, vs)
}

type EditCluster_SdkV2 struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
	// ID of the cluster
	ClusterId types.String `tfsdk:"cluster_id"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf"`
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
	DockerImage types.List `tfsdk:"docker_image"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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

	WorkloadType types.List `tfsdk:"workload_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(AutoScale_SdkV2{}),
		"aws_attributes":   reflect.TypeOf(AwsAttributes_SdkV2{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes_SdkV2{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf_SdkV2{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage_SdkV2{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes_SdkV2{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo_SdkV2{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o EditCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EditCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale": basetypes.ListType{
				ElemType: AutoScale_SdkV2{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes_SdkV2{}.Type(ctx),
			},
			"cluster_id": types.StringType,
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf_SdkV2{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes_SdkV2{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo_SdkV2{}.Type(ctx),
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in EditCluster_SdkV2 as
// a AutoScale_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetAutoscale(ctx context.Context) (AutoScale_SdkV2, bool) {
	var e AutoScale_SdkV2
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale_SdkV2
	d := o.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetAutoscale(ctx context.Context, v AutoScale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	o.Autoscale = types.ListValueMust(t, vs)
}

// GetAwsAttributes returns the value of the AwsAttributes field in EditCluster_SdkV2 as
// a AwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetAwsAttributes(ctx context.Context) (AwsAttributes_SdkV2, bool) {
	var e AwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetAwsAttributes(ctx context.Context, v AwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in EditCluster_SdkV2 as
// a AzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetAzureAttributes(ctx context.Context) (AzureAttributes_SdkV2, bool) {
	var e AzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetAzureAttributes(ctx context.Context, v AzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetClusterLogConf returns the value of the ClusterLogConf field in EditCluster_SdkV2 as
// a ClusterLogConf_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetClusterLogConf(ctx context.Context) (ClusterLogConf_SdkV2, bool) {
	var e ClusterLogConf_SdkV2
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf_SdkV2
	d := o.ClusterLogConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetClusterLogConf(ctx context.Context, v ClusterLogConf_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_conf"]
	o.ClusterLogConf = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in EditCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in EditCluster_SdkV2 as
// a DockerImage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetDockerImage(ctx context.Context) (DockerImage_SdkV2, bool) {
	var e DockerImage_SdkV2
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage_SdkV2
	d := o.DockerImage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetDockerImage(ctx context.Context, v DockerImage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["docker_image"]
	o.DockerImage = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in EditCluster_SdkV2 as
// a GcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetGcpAttributes(ctx context.Context) (GcpAttributes_SdkV2, bool) {
	var e GcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetGcpAttributes(ctx context.Context, v GcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in EditCluster_SdkV2 as
// a slice of InitScriptInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetInitScripts(ctx context.Context) ([]InitScriptInfo_SdkV2, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo_SdkV2
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetInitScripts(ctx context.Context, v []InitScriptInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in EditCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkConf sets the value of the SparkConf field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in EditCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkEnvVars sets the value of the SparkEnvVars field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in EditCluster_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
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

// SetSshPublicKeys sets the value of the SshPublicKeys field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in EditCluster_SdkV2 as
// a WorkloadType_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditCluster_SdkV2) GetWorkloadType(ctx context.Context) (WorkloadType_SdkV2, bool) {
	var e WorkloadType_SdkV2
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType_SdkV2
	d := o.WorkloadType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in EditCluster_SdkV2.
func (o *EditCluster_SdkV2) SetWorkloadType(ctx context.Context, v WorkloadType_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workload_type"]
	o.WorkloadType = types.ListValueMust(t, vs)
}

type EditClusterResponse_SdkV2 struct {
}

func (newState *EditClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan EditClusterResponse_SdkV2) {
}

func (newState *EditClusterResponse_SdkV2) SyncFieldsDuringRead(existingState EditClusterResponse_SdkV2) {
}

func (c EditClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EditClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditInstancePool_SdkV2 struct {
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
	NodeTypeFlexibility types.List `tfsdk:"node_type_flexibility"`
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
func (a EditInstancePool_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags":           reflect.TypeOf(types.String{}),
		"node_type_flexibility": reflect.TypeOf(NodeTypeFlexibility_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditInstancePool_SdkV2
// only implements ToObjectValue() and Type().
func (o EditInstancePool_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags":                           o.CustomTags,
			"enable_auto_alternate_node_types":      o.EnableAutoAlternateNodeTypes,
			"idle_instance_autotermination_minutes": o.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      o.InstancePoolId,
			"instance_pool_name":                    o.InstancePoolName,
			"max_capacity":                          o.MaxCapacity,
			"min_idle_instances":                    o.MinIdleInstances,
			"node_type_flexibility":                 o.NodeTypeFlexibility,
			"node_type_id":                          o.NodeTypeId,
			"remote_disk_throughput":                o.RemoteDiskThroughput,
			"total_initial_remote_disk_size":        o.TotalInitialRemoteDiskSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EditInstancePool_SdkV2) Type(ctx context.Context) attr.Type {
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
			"node_type_flexibility": basetypes.ListType{
				ElemType: NodeTypeFlexibility_SdkV2{}.Type(ctx),
			},
			"node_type_id":                   types.StringType,
			"remote_disk_throughput":         types.Int64Type,
			"total_initial_remote_disk_size": types.Int64Type,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in EditInstancePool_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditInstancePool_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in EditInstancePool_SdkV2.
func (o *EditInstancePool_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetNodeTypeFlexibility returns the value of the NodeTypeFlexibility field in EditInstancePool_SdkV2 as
// a NodeTypeFlexibility_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditInstancePool_SdkV2) GetNodeTypeFlexibility(ctx context.Context) (NodeTypeFlexibility_SdkV2, bool) {
	var e NodeTypeFlexibility_SdkV2
	if o.NodeTypeFlexibility.IsNull() || o.NodeTypeFlexibility.IsUnknown() {
		return e, false
	}
	var v []NodeTypeFlexibility_SdkV2
	d := o.NodeTypeFlexibility.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeTypeFlexibility sets the value of the NodeTypeFlexibility field in EditInstancePool_SdkV2.
func (o *EditInstancePool_SdkV2) SetNodeTypeFlexibility(ctx context.Context, v NodeTypeFlexibility_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_type_flexibility"]
	o.NodeTypeFlexibility = types.ListValueMust(t, vs)
}

type EditInstancePoolResponse_SdkV2 struct {
}

func (newState *EditInstancePoolResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan EditInstancePoolResponse_SdkV2) {
}

func (newState *EditInstancePoolResponse_SdkV2) SyncFieldsDuringRead(existingState EditInstancePoolResponse_SdkV2) {
}

func (c EditInstancePoolResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditInstancePoolResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditInstancePoolResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditInstancePoolResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EditInstancePoolResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditInstancePoolResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditPolicy_SdkV2 struct {
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
func (a EditPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o EditPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EditPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"definition":  types.StringType,
			"description": types.StringType,
			"libraries": basetypes.ListType{
				ElemType: Library_SdkV2{}.Type(ctx),
			},
			"max_clusters_per_user":              types.Int64Type,
			"name":                               types.StringType,
			"policy_family_definition_overrides": types.StringType,
			"policy_family_id":                   types.StringType,
			"policy_id":                          types.StringType,
		},
	}
}

// GetLibraries returns the value of the Libraries field in EditPolicy_SdkV2 as
// a slice of Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPolicy_SdkV2) GetLibraries(ctx context.Context) ([]Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in EditPolicy_SdkV2.
func (o *EditPolicy_SdkV2) SetLibraries(ctx context.Context, v []Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type EditPolicyResponse_SdkV2 struct {
}

func (newState *EditPolicyResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan EditPolicyResponse_SdkV2) {
}

func (newState *EditPolicyResponse_SdkV2) SyncFieldsDuringRead(existingState EditPolicyResponse_SdkV2) {
}

func (c EditPolicyResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditPolicyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPolicyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EditPolicyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditPolicyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditResponse_SdkV2 struct {
}

func (newState *EditResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan EditResponse_SdkV2) {
}

func (newState *EditResponse_SdkV2) SyncFieldsDuringRead(existingState EditResponse_SdkV2) {
}

func (c EditResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EditResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EnforceClusterComplianceRequest_SdkV2 struct {
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
func (a EnforceClusterComplianceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforceClusterComplianceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o EnforceClusterComplianceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":    o.ClusterId,
			"validate_only": o.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":    types.StringType,
			"validate_only": types.BoolType,
		},
	}
}

type EnforceClusterComplianceResponse_SdkV2 struct {
	// A list of changes that have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	Changes types.List `tfsdk:"changes"`
	// Whether any changes have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	HasChanges types.Bool `tfsdk:"has_changes"`
}

func (newState *EnforceClusterComplianceResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan EnforceClusterComplianceResponse_SdkV2) {
}

func (newState *EnforceClusterComplianceResponse_SdkV2) SyncFieldsDuringRead(existingState EnforceClusterComplianceResponse_SdkV2) {
}

func (c EnforceClusterComplianceResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnforceClusterComplianceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(ClusterSettingsChange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforceClusterComplianceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EnforceClusterComplianceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes":     o.Changes,
			"has_changes": o.HasChanges,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: ClusterSettingsChange_SdkV2{}.Type(ctx),
			},
			"has_changes": types.BoolType,
		},
	}
}

// GetChanges returns the value of the Changes field in EnforceClusterComplianceResponse_SdkV2 as
// a slice of ClusterSettingsChange_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EnforceClusterComplianceResponse_SdkV2) GetChanges(ctx context.Context) ([]ClusterSettingsChange_SdkV2, bool) {
	if o.Changes.IsNull() || o.Changes.IsUnknown() {
		return nil, false
	}
	var v []ClusterSettingsChange_SdkV2
	d := o.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in EnforceClusterComplianceResponse_SdkV2.
func (o *EnforceClusterComplianceResponse_SdkV2) SetChanges(ctx context.Context, v []ClusterSettingsChange_SdkV2) {
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
type Environment_SdkV2 struct {
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

func (newState *Environment_SdkV2) SyncFieldsDuringCreateOrUpdate(plan Environment_SdkV2) {
}

func (newState *Environment_SdkV2) SyncFieldsDuringRead(existingState Environment_SdkV2) {
}

func (c Environment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Environment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies":     reflect.TypeOf(types.String{}),
		"jar_dependencies": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Environment_SdkV2
// only implements ToObjectValue() and Type().
func (o Environment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Environment_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetDependencies returns the value of the Dependencies field in Environment_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Environment_SdkV2) GetDependencies(ctx context.Context) ([]types.String, bool) {
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

// SetDependencies sets the value of the Dependencies field in Environment_SdkV2.
func (o *Environment_SdkV2) SetDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dependencies = types.ListValueMust(t, vs)
}

// GetJarDependencies returns the value of the JarDependencies field in Environment_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Environment_SdkV2) GetJarDependencies(ctx context.Context) ([]types.String, bool) {
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

// SetJarDependencies sets the value of the JarDependencies field in Environment_SdkV2.
func (o *Environment_SdkV2) SetJarDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["jar_dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JarDependencies = types.ListValueMust(t, vs)
}

type EventDetails_SdkV2 struct {
	// * For created clusters, the attributes of the cluster. * For edited
	// clusters, the new attributes of the cluster.
	Attributes types.List `tfsdk:"attributes"`
	// The cause of a change in target size.
	Cause types.String `tfsdk:"cause"`
	// The actual cluster size that was set in the cluster creation or edit.
	ClusterSize types.List `tfsdk:"cluster_size"`
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
	InitScripts types.List `tfsdk:"init_scripts"`
	// Instance Id where the event originated from
	InstanceId types.String `tfsdk:"instance_id"`
	// Unique identifier of the specific job run associated with this cluster
	// event * For clusters created for jobs, this will be the same as the
	// cluster name
	JobRunName types.String `tfsdk:"job_run_name"`
	// The cluster attributes before a cluster was edited.
	PreviousAttributes types.List `tfsdk:"previous_attributes"`
	// The size of the cluster before an edit or resize.
	PreviousClusterSize types.List `tfsdk:"previous_cluster_size"`
	// Previous disk size in bytes
	PreviousDiskSize types.Int64 `tfsdk:"previous_disk_size"`
	// A termination reason: * On a TERMINATED event, this is the reason of the
	// termination. * On a RESIZE_COMPLETE event, this indicates the reason that
	// we failed to acquire some nodes.
	Reason types.List `tfsdk:"reason"`
	// The targeted number of vCPUs in the cluster.
	TargetNumVcpus types.Int64 `tfsdk:"target_num_vcpus"`
	// The targeted number of nodes in the cluster.
	TargetNumWorkers types.Int64 `tfsdk:"target_num_workers"`
	// The user that caused the event to occur. (Empty if it was done by the
	// control plane.)
	User types.String `tfsdk:"user"`
}

func (newState *EventDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(plan EventDetails_SdkV2) {
}

func (newState *EventDetails_SdkV2) SyncFieldsDuringRead(existingState EventDetails_SdkV2) {
}

func (c EventDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["attributes"] = attrs["attributes"].SetOptional()
	attrs["attributes"] = attrs["attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["cluster_size"] = attrs["cluster_size"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["current_num_vcpus"] = attrs["current_num_vcpus"].SetOptional()
	attrs["current_num_workers"] = attrs["current_num_workers"].SetOptional()
	attrs["did_not_expand_reason"] = attrs["did_not_expand_reason"].SetOptional()
	attrs["disk_size"] = attrs["disk_size"].SetOptional()
	attrs["driver_state_message"] = attrs["driver_state_message"].SetOptional()
	attrs["enable_termination_for_node_blocklisted"] = attrs["enable_termination_for_node_blocklisted"].SetOptional()
	attrs["free_space"] = attrs["free_space"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["init_scripts"] = attrs["init_scripts"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["instance_id"] = attrs["instance_id"].SetOptional()
	attrs["job_run_name"] = attrs["job_run_name"].SetOptional()
	attrs["previous_attributes"] = attrs["previous_attributes"].SetOptional()
	attrs["previous_attributes"] = attrs["previous_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["previous_cluster_size"] = attrs["previous_cluster_size"].SetOptional()
	attrs["previous_cluster_size"] = attrs["previous_cluster_size"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["previous_disk_size"] = attrs["previous_disk_size"].SetOptional()
	attrs["reason"] = attrs["reason"].SetOptional()
	attrs["reason"] = attrs["reason"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a EventDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attributes":            reflect.TypeOf(ClusterAttributes_SdkV2{}),
		"cluster_size":          reflect.TypeOf(ClusterSize_SdkV2{}),
		"init_scripts":          reflect.TypeOf(InitScriptEventDetails_SdkV2{}),
		"previous_attributes":   reflect.TypeOf(ClusterAttributes_SdkV2{}),
		"previous_cluster_size": reflect.TypeOf(ClusterSize_SdkV2{}),
		"reason":                reflect.TypeOf(TerminationReason_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EventDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o EventDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EventDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes": basetypes.ListType{
				ElemType: ClusterAttributes_SdkV2{}.Type(ctx),
			},
			"cause": types.StringType,
			"cluster_size": basetypes.ListType{
				ElemType: ClusterSize_SdkV2{}.Type(ctx),
			},
			"current_num_vcpus":                       types.Int64Type,
			"current_num_workers":                     types.Int64Type,
			"did_not_expand_reason":                   types.StringType,
			"disk_size":                               types.Int64Type,
			"driver_state_message":                    types.StringType,
			"enable_termination_for_node_blocklisted": types.BoolType,
			"free_space":                              types.Int64Type,
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptEventDetails_SdkV2{}.Type(ctx),
			},
			"instance_id":  types.StringType,
			"job_run_name": types.StringType,
			"previous_attributes": basetypes.ListType{
				ElemType: ClusterAttributes_SdkV2{}.Type(ctx),
			},
			"previous_cluster_size": basetypes.ListType{
				ElemType: ClusterSize_SdkV2{}.Type(ctx),
			},
			"previous_disk_size": types.Int64Type,
			"reason": basetypes.ListType{
				ElemType: TerminationReason_SdkV2{}.Type(ctx),
			},
			"target_num_vcpus":   types.Int64Type,
			"target_num_workers": types.Int64Type,
			"user":               types.StringType,
		},
	}
}

// GetAttributes returns the value of the Attributes field in EventDetails_SdkV2 as
// a ClusterAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails_SdkV2) GetAttributes(ctx context.Context) (ClusterAttributes_SdkV2, bool) {
	var e ClusterAttributes_SdkV2
	if o.Attributes.IsNull() || o.Attributes.IsUnknown() {
		return e, false
	}
	var v []ClusterAttributes_SdkV2
	d := o.Attributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAttributes sets the value of the Attributes field in EventDetails_SdkV2.
func (o *EventDetails_SdkV2) SetAttributes(ctx context.Context, v ClusterAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["attributes"]
	o.Attributes = types.ListValueMust(t, vs)
}

// GetClusterSize returns the value of the ClusterSize field in EventDetails_SdkV2 as
// a ClusterSize_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails_SdkV2) GetClusterSize(ctx context.Context) (ClusterSize_SdkV2, bool) {
	var e ClusterSize_SdkV2
	if o.ClusterSize.IsNull() || o.ClusterSize.IsUnknown() {
		return e, false
	}
	var v []ClusterSize_SdkV2
	d := o.ClusterSize.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterSize sets the value of the ClusterSize field in EventDetails_SdkV2.
func (o *EventDetails_SdkV2) SetClusterSize(ctx context.Context, v ClusterSize_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_size"]
	o.ClusterSize = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in EventDetails_SdkV2 as
// a InitScriptEventDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails_SdkV2) GetInitScripts(ctx context.Context) (InitScriptEventDetails_SdkV2, bool) {
	var e InitScriptEventDetails_SdkV2
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return e, false
	}
	var v []InitScriptEventDetails_SdkV2
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitScripts sets the value of the InitScripts field in EventDetails_SdkV2.
func (o *EventDetails_SdkV2) SetInitScripts(ctx context.Context, v InitScriptEventDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetPreviousAttributes returns the value of the PreviousAttributes field in EventDetails_SdkV2 as
// a ClusterAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails_SdkV2) GetPreviousAttributes(ctx context.Context) (ClusterAttributes_SdkV2, bool) {
	var e ClusterAttributes_SdkV2
	if o.PreviousAttributes.IsNull() || o.PreviousAttributes.IsUnknown() {
		return e, false
	}
	var v []ClusterAttributes_SdkV2
	d := o.PreviousAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPreviousAttributes sets the value of the PreviousAttributes field in EventDetails_SdkV2.
func (o *EventDetails_SdkV2) SetPreviousAttributes(ctx context.Context, v ClusterAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["previous_attributes"]
	o.PreviousAttributes = types.ListValueMust(t, vs)
}

// GetPreviousClusterSize returns the value of the PreviousClusterSize field in EventDetails_SdkV2 as
// a ClusterSize_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails_SdkV2) GetPreviousClusterSize(ctx context.Context) (ClusterSize_SdkV2, bool) {
	var e ClusterSize_SdkV2
	if o.PreviousClusterSize.IsNull() || o.PreviousClusterSize.IsUnknown() {
		return e, false
	}
	var v []ClusterSize_SdkV2
	d := o.PreviousClusterSize.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPreviousClusterSize sets the value of the PreviousClusterSize field in EventDetails_SdkV2.
func (o *EventDetails_SdkV2) SetPreviousClusterSize(ctx context.Context, v ClusterSize_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["previous_cluster_size"]
	o.PreviousClusterSize = types.ListValueMust(t, vs)
}

// GetReason returns the value of the Reason field in EventDetails_SdkV2 as
// a TerminationReason_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EventDetails_SdkV2) GetReason(ctx context.Context) (TerminationReason_SdkV2, bool) {
	var e TerminationReason_SdkV2
	if o.Reason.IsNull() || o.Reason.IsUnknown() {
		return e, false
	}
	var v []TerminationReason_SdkV2
	d := o.Reason.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetReason sets the value of the Reason field in EventDetails_SdkV2.
func (o *EventDetails_SdkV2) SetReason(ctx context.Context, v TerminationReason_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["reason"]
	o.Reason = types.ListValueMust(t, vs)
}

// Attributes set during cluster creation which are related to GCP.
type GcpAttributes_SdkV2 struct {
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

func (newState *GcpAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GcpAttributes_SdkV2) {
}

func (newState *GcpAttributes_SdkV2) SyncFieldsDuringRead(existingState GcpAttributes_SdkV2) {
}

func (c GcpAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GcpAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GcpAttributes_SdkV2) Type(ctx context.Context) attr.Type {
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
type GcsStorageInfo_SdkV2 struct {
	// GCS destination/URI, e.g. `gs://my-bucket/some-prefix`
	Destination types.String `tfsdk:"destination"`
}

func (newState *GcsStorageInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GcsStorageInfo_SdkV2) {
}

func (newState *GcsStorageInfo_SdkV2) SyncFieldsDuringRead(existingState GcsStorageInfo_SdkV2) {
}

func (c GcsStorageInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GcsStorageInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcsStorageInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o GcsStorageInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcsStorageInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type GetClusterComplianceRequest_SdkV2 struct {
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
func (a GetClusterComplianceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterComplianceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterComplianceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterComplianceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type GetClusterComplianceResponse_SdkV2 struct {
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

func (newState *GetClusterComplianceResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GetClusterComplianceResponse_SdkV2) {
}

func (newState *GetClusterComplianceResponse_SdkV2) SyncFieldsDuringRead(existingState GetClusterComplianceResponse_SdkV2) {
}

func (c GetClusterComplianceResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetClusterComplianceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterComplianceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterComplianceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_compliant": o.IsCompliant,
			"violations":   o.Violations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterComplianceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_compliant": types.BoolType,
			"violations": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetViolations returns the value of the Violations field in GetClusterComplianceResponse_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetClusterComplianceResponse_SdkV2) GetViolations(ctx context.Context) (map[string]types.String, bool) {
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

// SetViolations sets the value of the Violations field in GetClusterComplianceResponse_SdkV2.
func (o *GetClusterComplianceResponse_SdkV2) SetViolations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["violations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Violations = types.MapValueMust(t, vs)
}

type GetClusterPermissionLevelsRequest_SdkV2 struct {
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
func (a GetClusterPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type GetClusterPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetClusterPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GetClusterPermissionLevelsResponse_SdkV2) {
}

func (newState *GetClusterPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(existingState GetClusterPermissionLevelsResponse_SdkV2) {
}

func (c GetClusterPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetClusterPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ClusterPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ClusterPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetClusterPermissionLevelsResponse_SdkV2 as
// a slice of ClusterPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetClusterPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]ClusterPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ClusterPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetClusterPermissionLevelsResponse_SdkV2.
func (o *GetClusterPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []ClusterPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetClusterPermissionsRequest_SdkV2 struct {
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
func (a GetClusterPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type GetClusterPolicyPermissionLevelsRequest_SdkV2 struct {
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
func (a GetClusterPolicyPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterPolicyPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_policy_id": o.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_policy_id": types.StringType,
		},
	}
}

type GetClusterPolicyPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetClusterPolicyPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GetClusterPolicyPermissionLevelsResponse_SdkV2) {
}

func (newState *GetClusterPolicyPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(existingState GetClusterPolicyPermissionLevelsResponse_SdkV2) {
}

func (c GetClusterPolicyPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetClusterPolicyPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ClusterPolicyPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterPolicyPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ClusterPolicyPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetClusterPolicyPermissionLevelsResponse_SdkV2 as
// a slice of ClusterPolicyPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetClusterPolicyPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]ClusterPolicyPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ClusterPolicyPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetClusterPolicyPermissionLevelsResponse_SdkV2.
func (o *GetClusterPolicyPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []ClusterPolicyPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetClusterPolicyPermissionsRequest_SdkV2 struct {
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
func (a GetClusterPolicyPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterPolicyPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_policy_id": o.ClusterPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_policy_id": types.StringType,
		},
	}
}

type GetClusterPolicyRequest_SdkV2 struct {
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
func (a GetClusterPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type GetClusterRequest_SdkV2 struct {
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
func (a GetClusterRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetClusterRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetClusterRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type GetEvents_SdkV2 struct {
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

func (newState *GetEvents_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GetEvents_SdkV2) {
}

func (newState *GetEvents_SdkV2) SyncFieldsDuringRead(existingState GetEvents_SdkV2) {
}

func (c GetEvents_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetEvents_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"event_types": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEvents_SdkV2
// only implements ToObjectValue() and Type().
func (o GetEvents_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetEvents_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetEventTypes returns the value of the EventTypes field in GetEvents_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetEvents_SdkV2) GetEventTypes(ctx context.Context) ([]types.String, bool) {
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

// SetEventTypes sets the value of the EventTypes field in GetEvents_SdkV2.
func (o *GetEvents_SdkV2) SetEventTypes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["event_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EventTypes = types.ListValueMust(t, vs)
}

type GetEventsResponse_SdkV2 struct {
	Events types.List `tfsdk:"events"`
	// Deprecated: use next_page_token or prev_page_token instead.
	//
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	NextPage types.List `tfsdk:"next_page"`
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

func (newState *GetEventsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GetEventsResponse_SdkV2) {
}

func (newState *GetEventsResponse_SdkV2) SyncFieldsDuringRead(existingState GetEventsResponse_SdkV2) {
}

func (c GetEventsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["events"] = attrs["events"].SetOptional()
	attrs["next_page"] = attrs["next_page"].SetOptional()
	attrs["next_page"] = attrs["next_page"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a GetEventsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":    reflect.TypeOf(ClusterEvent_SdkV2{}),
		"next_page": reflect.TypeOf(GetEvents_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEventsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetEventsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetEventsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"events": basetypes.ListType{
				ElemType: ClusterEvent_SdkV2{}.Type(ctx),
			},
			"next_page": basetypes.ListType{
				ElemType: GetEvents_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
			"total_count":     types.Int64Type,
		},
	}
}

// GetEvents returns the value of the Events field in GetEventsResponse_SdkV2 as
// a slice of ClusterEvent_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetEventsResponse_SdkV2) GetEvents(ctx context.Context) ([]ClusterEvent_SdkV2, bool) {
	if o.Events.IsNull() || o.Events.IsUnknown() {
		return nil, false
	}
	var v []ClusterEvent_SdkV2
	d := o.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in GetEventsResponse_SdkV2.
func (o *GetEventsResponse_SdkV2) SetEvents(ctx context.Context, v []ClusterEvent_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// GetNextPage returns the value of the NextPage field in GetEventsResponse_SdkV2 as
// a GetEvents_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetEventsResponse_SdkV2) GetNextPage(ctx context.Context) (GetEvents_SdkV2, bool) {
	var e GetEvents_SdkV2
	if o.NextPage.IsNull() || o.NextPage.IsUnknown() {
		return e, false
	}
	var v []GetEvents_SdkV2
	d := o.NextPage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNextPage sets the value of the NextPage field in GetEventsResponse_SdkV2.
func (o *GetEventsResponse_SdkV2) SetNextPage(ctx context.Context, v GetEvents_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["next_page"]
	o.NextPage = types.ListValueMust(t, vs)
}

type GetGlobalInitScriptRequest_SdkV2 struct {
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
func (a GetGlobalInitScriptRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGlobalInitScriptRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetGlobalInitScriptRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"script_id": o.ScriptId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetGlobalInitScriptRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"script_id": types.StringType,
		},
	}
}

type GetInstancePool_SdkV2 struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
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
	DiskSpec types.List `tfsdk:"disk_spec"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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
	NodeTypeFlexibility types.List `tfsdk:"node_type_flexibility"`
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
	Stats types.List `tfsdk:"stats"`
	// Status of failed pending instances in the pool.
	Status types.List `tfsdk:"status"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
}

func (newState *GetInstancePool_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GetInstancePool_SdkV2) {
}

func (newState *GetInstancePool_SdkV2) SyncFieldsDuringRead(existingState GetInstancePool_SdkV2) {
}

func (c GetInstancePool_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["default_tags"] = attrs["default_tags"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["enable_auto_alternate_node_types"] = attrs["enable_auto_alternate_node_types"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["idle_instance_autotermination_minutes"] = attrs["idle_instance_autotermination_minutes"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetRequired()
	attrs["instance_pool_name"] = attrs["instance_pool_name"].SetOptional()
	attrs["max_capacity"] = attrs["max_capacity"].SetOptional()
	attrs["min_idle_instances"] = attrs["min_idle_instances"].SetOptional()
	attrs["node_type_flexibility"] = attrs["node_type_flexibility"].SetOptional()
	attrs["node_type_flexibility"] = attrs["node_type_flexibility"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["preloaded_docker_images"] = attrs["preloaded_docker_images"].SetOptional()
	attrs["preloaded_spark_versions"] = attrs["preloaded_spark_versions"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["stats"] = attrs["stats"].SetOptional()
	attrs["stats"] = attrs["stats"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a GetInstancePool_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes_SdkV2{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes_SdkV2{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"default_tags":             reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec_SdkV2{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes_SdkV2{}),
		"node_type_flexibility":    reflect.TypeOf(NodeTypeFlexibility_SdkV2{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage_SdkV2{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
		"stats":                    reflect.TypeOf(InstancePoolStats_SdkV2{}),
		"status":                   reflect.TypeOf(InstancePoolStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePool_SdkV2
// only implements ToObjectValue() and Type().
func (o GetInstancePool_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        o.AwsAttributes,
			"azure_attributes":                      o.AzureAttributes,
			"custom_tags":                           o.CustomTags,
			"default_tags":                          o.DefaultTags,
			"disk_spec":                             o.DiskSpec,
			"enable_auto_alternate_node_types":      o.EnableAutoAlternateNodeTypes,
			"enable_elastic_disk":                   o.EnableElasticDisk,
			"gcp_attributes":                        o.GcpAttributes,
			"idle_instance_autotermination_minutes": o.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      o.InstancePoolId,
			"instance_pool_name":                    o.InstancePoolName,
			"max_capacity":                          o.MaxCapacity,
			"min_idle_instances":                    o.MinIdleInstances,
			"node_type_flexibility":                 o.NodeTypeFlexibility,
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
func (o GetInstancePool_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes": basetypes.ListType{
				ElemType: InstancePoolAwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: InstancePoolAzureAttributes_SdkV2{}.Type(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"default_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec": basetypes.ListType{
				ElemType: DiskSpec_SdkV2{}.Type(ctx),
			},
			"enable_auto_alternate_node_types": types.BoolType,
			"enable_elastic_disk":              types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: InstancePoolGcpAttributes_SdkV2{}.Type(ctx),
			},
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_id":                      types.StringType,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
			"node_type_flexibility": basetypes.ListType{
				ElemType: NodeTypeFlexibility_SdkV2{}.Type(ctx),
			},
			"node_type_id": types.StringType,
			"preloaded_docker_images": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"preloaded_spark_versions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"remote_disk_throughput": types.Int64Type,
			"state":                  types.StringType,
			"stats": basetypes.ListType{
				ElemType: InstancePoolStats_SdkV2{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: InstancePoolStatus_SdkV2{}.Type(ctx),
			},
			"total_initial_remote_disk_size": types.Int64Type,
		},
	}
}

// GetAwsAttributes returns the value of the AwsAttributes field in GetInstancePool_SdkV2 as
// a InstancePoolAwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes_SdkV2, bool) {
	var e InstancePoolAwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in GetInstancePool_SdkV2 as
// a InstancePoolAzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes_SdkV2, bool) {
	var e InstancePoolAzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in GetInstancePool_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in GetInstancePool_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetDefaultTags sets the value of the DefaultTags field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DefaultTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in GetInstancePool_SdkV2 as
// a DiskSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetDiskSpec(ctx context.Context) (DiskSpec_SdkV2, bool) {
	var e DiskSpec_SdkV2
	if o.DiskSpec.IsNull() || o.DiskSpec.IsUnknown() {
		return e, false
	}
	var v []DiskSpec_SdkV2
	d := o.DiskSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDiskSpec sets the value of the DiskSpec field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetDiskSpec(ctx context.Context, v DiskSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["disk_spec"]
	o.DiskSpec = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in GetInstancePool_SdkV2 as
// a InstancePoolGcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes_SdkV2, bool) {
	var e InstancePoolGcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolGcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetNodeTypeFlexibility returns the value of the NodeTypeFlexibility field in GetInstancePool_SdkV2 as
// a NodeTypeFlexibility_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetNodeTypeFlexibility(ctx context.Context) (NodeTypeFlexibility_SdkV2, bool) {
	var e NodeTypeFlexibility_SdkV2
	if o.NodeTypeFlexibility.IsNull() || o.NodeTypeFlexibility.IsUnknown() {
		return e, false
	}
	var v []NodeTypeFlexibility_SdkV2
	d := o.NodeTypeFlexibility.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeTypeFlexibility sets the value of the NodeTypeFlexibility field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetNodeTypeFlexibility(ctx context.Context, v NodeTypeFlexibility_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_type_flexibility"]
	o.NodeTypeFlexibility = types.ListValueMust(t, vs)
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in GetInstancePool_SdkV2 as
// a slice of DockerImage_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage_SdkV2, bool) {
	if o.PreloadedDockerImages.IsNull() || o.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage_SdkV2
	d := o.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetPreloadedDockerImages(ctx context.Context, v []DockerImage_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in GetInstancePool_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
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

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

// GetStats returns the value of the Stats field in GetInstancePool_SdkV2 as
// a InstancePoolStats_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetStats(ctx context.Context) (InstancePoolStats_SdkV2, bool) {
	var e InstancePoolStats_SdkV2
	if o.Stats.IsNull() || o.Stats.IsUnknown() {
		return e, false
	}
	var v []InstancePoolStats_SdkV2
	d := o.Stats.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStats sets the value of the Stats field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetStats(ctx context.Context, v InstancePoolStats_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stats"]
	o.Stats = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in GetInstancePool_SdkV2 as
// a InstancePoolStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePool_SdkV2) GetStatus(ctx context.Context) (InstancePoolStatus_SdkV2, bool) {
	var e InstancePoolStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []InstancePoolStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in GetInstancePool_SdkV2.
func (o *GetInstancePool_SdkV2) SetStatus(ctx context.Context, v InstancePoolStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

type GetInstancePoolPermissionLevelsRequest_SdkV2 struct {
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
func (a GetInstancePoolPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetInstancePoolPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type GetInstancePoolPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetInstancePoolPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GetInstancePoolPermissionLevelsResponse_SdkV2) {
}

func (newState *GetInstancePoolPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(existingState GetInstancePoolPermissionLevelsResponse_SdkV2) {
}

func (c GetInstancePoolPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetInstancePoolPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(InstancePoolPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetInstancePoolPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: InstancePoolPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetInstancePoolPermissionLevelsResponse_SdkV2 as
// a slice of InstancePoolPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetInstancePoolPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]InstancePoolPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetInstancePoolPermissionLevelsResponse_SdkV2.
func (o *GetInstancePoolPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []InstancePoolPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetInstancePoolPermissionsRequest_SdkV2 struct {
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
func (a GetInstancePoolPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetInstancePoolPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type GetInstancePoolRequest_SdkV2 struct {
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
func (a GetInstancePoolRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetInstancePoolRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetInstancePoolRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pool_id": o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

type GetPolicyFamilyRequest_SdkV2 struct {
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
func (a GetPolicyFamilyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPolicyFamilyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPolicyFamilyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_family_id": o.PolicyFamilyId,
			"version":          o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPolicyFamilyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_family_id": types.StringType,
			"version":          types.Int64Type,
		},
	}
}

type GetSparkVersionsResponse_SdkV2 struct {
	// All the available Spark versions.
	Versions types.List `tfsdk:"versions"`
}

func (newState *GetSparkVersionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GetSparkVersionsResponse_SdkV2) {
}

func (newState *GetSparkVersionsResponse_SdkV2) SyncFieldsDuringRead(existingState GetSparkVersionsResponse_SdkV2) {
}

func (c GetSparkVersionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetSparkVersionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"versions": reflect.TypeOf(SparkVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSparkVersionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetSparkVersionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"versions": o.Versions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSparkVersionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"versions": basetypes.ListType{
				ElemType: SparkVersion_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetVersions returns the value of the Versions field in GetSparkVersionsResponse_SdkV2 as
// a slice of SparkVersion_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetSparkVersionsResponse_SdkV2) GetVersions(ctx context.Context) ([]SparkVersion_SdkV2, bool) {
	if o.Versions.IsNull() || o.Versions.IsUnknown() {
		return nil, false
	}
	var v []SparkVersion_SdkV2
	d := o.Versions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVersions sets the value of the Versions field in GetSparkVersionsResponse_SdkV2.
func (o *GetSparkVersionsResponse_SdkV2) SetVersions(ctx context.Context, v []SparkVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Versions = types.ListValueMust(t, vs)
}

type GlobalInitScriptCreateRequest_SdkV2 struct {
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
func (a GlobalInitScriptCreateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptCreateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GlobalInitScriptCreateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GlobalInitScriptCreateRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":  types.BoolType,
			"name":     types.StringType,
			"position": types.Int64Type,
			"script":   types.StringType,
		},
	}
}

type GlobalInitScriptDetails_SdkV2 struct {
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

func (newState *GlobalInitScriptDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GlobalInitScriptDetails_SdkV2) {
}

func (newState *GlobalInitScriptDetails_SdkV2) SyncFieldsDuringRead(existingState GlobalInitScriptDetails_SdkV2) {
}

func (c GlobalInitScriptDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GlobalInitScriptDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o GlobalInitScriptDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GlobalInitScriptDetails_SdkV2) Type(ctx context.Context) attr.Type {
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

type GlobalInitScriptDetailsWithContent_SdkV2 struct {
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

func (newState *GlobalInitScriptDetailsWithContent_SdkV2) SyncFieldsDuringCreateOrUpdate(plan GlobalInitScriptDetailsWithContent_SdkV2) {
}

func (newState *GlobalInitScriptDetailsWithContent_SdkV2) SyncFieldsDuringRead(existingState GlobalInitScriptDetailsWithContent_SdkV2) {
}

func (c GlobalInitScriptDetailsWithContent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GlobalInitScriptDetailsWithContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptDetailsWithContent_SdkV2
// only implements ToObjectValue() and Type().
func (o GlobalInitScriptDetailsWithContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GlobalInitScriptDetailsWithContent_SdkV2) Type(ctx context.Context) attr.Type {
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

type GlobalInitScriptUpdateRequest_SdkV2 struct {
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
func (a GlobalInitScriptUpdateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GlobalInitScriptUpdateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GlobalInitScriptUpdateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GlobalInitScriptUpdateRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type InitScriptEventDetails_SdkV2 struct {
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

func (newState *InitScriptEventDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InitScriptEventDetails_SdkV2) {
}

func (newState *InitScriptEventDetails_SdkV2) SyncFieldsDuringRead(existingState InitScriptEventDetails_SdkV2) {
}

func (c InitScriptEventDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InitScriptEventDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster": reflect.TypeOf(InitScriptInfoAndExecutionDetails_SdkV2{}),
		"global":  reflect.TypeOf(InitScriptInfoAndExecutionDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InitScriptEventDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o InitScriptEventDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster":           o.Cluster,
			"global":            o.Global,
			"reported_for_node": o.ReportedForNode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InitScriptEventDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster": basetypes.ListType{
				ElemType: InitScriptInfoAndExecutionDetails_SdkV2{}.Type(ctx),
			},
			"global": basetypes.ListType{
				ElemType: InitScriptInfoAndExecutionDetails_SdkV2{}.Type(ctx),
			},
			"reported_for_node": types.StringType,
		},
	}
}

// GetCluster returns the value of the Cluster field in InitScriptEventDetails_SdkV2 as
// a slice of InitScriptInfoAndExecutionDetails_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptEventDetails_SdkV2) GetCluster(ctx context.Context) ([]InitScriptInfoAndExecutionDetails_SdkV2, bool) {
	if o.Cluster.IsNull() || o.Cluster.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfoAndExecutionDetails_SdkV2
	d := o.Cluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCluster sets the value of the Cluster field in InitScriptEventDetails_SdkV2.
func (o *InitScriptEventDetails_SdkV2) SetCluster(ctx context.Context, v []InitScriptInfoAndExecutionDetails_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Cluster = types.ListValueMust(t, vs)
}

// GetGlobal returns the value of the Global field in InitScriptEventDetails_SdkV2 as
// a slice of InitScriptInfoAndExecutionDetails_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptEventDetails_SdkV2) GetGlobal(ctx context.Context) ([]InitScriptInfoAndExecutionDetails_SdkV2, bool) {
	if o.Global.IsNull() || o.Global.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfoAndExecutionDetails_SdkV2
	d := o.Global.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGlobal sets the value of the Global field in InitScriptEventDetails_SdkV2.
func (o *InitScriptEventDetails_SdkV2) SetGlobal(ctx context.Context, v []InitScriptInfoAndExecutionDetails_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["global"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Global = types.ListValueMust(t, vs)
}

// Config for an individual init script Next ID: 11
type InitScriptInfo_SdkV2 struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	Abfss types.List `tfsdk:"abfss"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs types.List `tfsdk:"dbfs"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File types.List `tfsdk:"file"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs types.List `tfsdk:"gcs"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 types.List `tfsdk:"s3"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes types.List `tfsdk:"volumes"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace types.List `tfsdk:"workspace"`
}

func (newState *InitScriptInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InitScriptInfo_SdkV2) {
}

func (newState *InitScriptInfo_SdkV2) SyncFieldsDuringRead(existingState InitScriptInfo_SdkV2) {
}

func (c InitScriptInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["abfss"] = attrs["abfss"].SetOptional()
	attrs["abfss"] = attrs["abfss"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["dbfs"] = attrs["dbfs"].SetOptional()
	attrs["dbfs"] = attrs["dbfs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["file"] = attrs["file"].SetOptional()
	attrs["file"] = attrs["file"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["gcs"] = attrs["gcs"].SetOptional()
	attrs["gcs"] = attrs["gcs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["s3"] = attrs["s3"].SetOptional()
	attrs["s3"] = attrs["s3"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["volumes"] = attrs["volumes"].SetOptional()
	attrs["volumes"] = attrs["volumes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["workspace"] = attrs["workspace"].SetOptional()
	attrs["workspace"] = attrs["workspace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InitScriptInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InitScriptInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"abfss":     reflect.TypeOf(Adlsgen2Info_SdkV2{}),
		"dbfs":      reflect.TypeOf(DbfsStorageInfo_SdkV2{}),
		"file":      reflect.TypeOf(LocalFileInfo_SdkV2{}),
		"gcs":       reflect.TypeOf(GcsStorageInfo_SdkV2{}),
		"s3":        reflect.TypeOf(S3StorageInfo_SdkV2{}),
		"volumes":   reflect.TypeOf(VolumesStorageInfo_SdkV2{}),
		"workspace": reflect.TypeOf(WorkspaceStorageInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InitScriptInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o InitScriptInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o InitScriptInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"abfss": basetypes.ListType{
				ElemType: Adlsgen2Info_SdkV2{}.Type(ctx),
			},
			"dbfs": basetypes.ListType{
				ElemType: DbfsStorageInfo_SdkV2{}.Type(ctx),
			},
			"file": basetypes.ListType{
				ElemType: LocalFileInfo_SdkV2{}.Type(ctx),
			},
			"gcs": basetypes.ListType{
				ElemType: GcsStorageInfo_SdkV2{}.Type(ctx),
			},
			"s3": basetypes.ListType{
				ElemType: S3StorageInfo_SdkV2{}.Type(ctx),
			},
			"volumes": basetypes.ListType{
				ElemType: VolumesStorageInfo_SdkV2{}.Type(ctx),
			},
			"workspace": basetypes.ListType{
				ElemType: WorkspaceStorageInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAbfss returns the value of the Abfss field in InitScriptInfo_SdkV2 as
// a Adlsgen2Info_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo_SdkV2) GetAbfss(ctx context.Context) (Adlsgen2Info_SdkV2, bool) {
	var e Adlsgen2Info_SdkV2
	if o.Abfss.IsNull() || o.Abfss.IsUnknown() {
		return e, false
	}
	var v []Adlsgen2Info_SdkV2
	d := o.Abfss.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAbfss sets the value of the Abfss field in InitScriptInfo_SdkV2.
func (o *InitScriptInfo_SdkV2) SetAbfss(ctx context.Context, v Adlsgen2Info_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["abfss"]
	o.Abfss = types.ListValueMust(t, vs)
}

// GetDbfs returns the value of the Dbfs field in InitScriptInfo_SdkV2 as
// a DbfsStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo_SdkV2) GetDbfs(ctx context.Context) (DbfsStorageInfo_SdkV2, bool) {
	var e DbfsStorageInfo_SdkV2
	if o.Dbfs.IsNull() || o.Dbfs.IsUnknown() {
		return e, false
	}
	var v []DbfsStorageInfo_SdkV2
	d := o.Dbfs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbfs sets the value of the Dbfs field in InitScriptInfo_SdkV2.
func (o *InitScriptInfo_SdkV2) SetDbfs(ctx context.Context, v DbfsStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbfs"]
	o.Dbfs = types.ListValueMust(t, vs)
}

// GetFile returns the value of the File field in InitScriptInfo_SdkV2 as
// a LocalFileInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo_SdkV2) GetFile(ctx context.Context) (LocalFileInfo_SdkV2, bool) {
	var e LocalFileInfo_SdkV2
	if o.File.IsNull() || o.File.IsUnknown() {
		return e, false
	}
	var v []LocalFileInfo_SdkV2
	d := o.File.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFile sets the value of the File field in InitScriptInfo_SdkV2.
func (o *InitScriptInfo_SdkV2) SetFile(ctx context.Context, v LocalFileInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file"]
	o.File = types.ListValueMust(t, vs)
}

// GetGcs returns the value of the Gcs field in InitScriptInfo_SdkV2 as
// a GcsStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo_SdkV2) GetGcs(ctx context.Context) (GcsStorageInfo_SdkV2, bool) {
	var e GcsStorageInfo_SdkV2
	if o.Gcs.IsNull() || o.Gcs.IsUnknown() {
		return e, false
	}
	var v []GcsStorageInfo_SdkV2
	d := o.Gcs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcs sets the value of the Gcs field in InitScriptInfo_SdkV2.
func (o *InitScriptInfo_SdkV2) SetGcs(ctx context.Context, v GcsStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcs"]
	o.Gcs = types.ListValueMust(t, vs)
}

// GetS3 returns the value of the S3 field in InitScriptInfo_SdkV2 as
// a S3StorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo_SdkV2) GetS3(ctx context.Context) (S3StorageInfo_SdkV2, bool) {
	var e S3StorageInfo_SdkV2
	if o.S3.IsNull() || o.S3.IsUnknown() {
		return e, false
	}
	var v []S3StorageInfo_SdkV2
	d := o.S3.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetS3 sets the value of the S3 field in InitScriptInfo_SdkV2.
func (o *InitScriptInfo_SdkV2) SetS3(ctx context.Context, v S3StorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["s3"]
	o.S3 = types.ListValueMust(t, vs)
}

// GetVolumes returns the value of the Volumes field in InitScriptInfo_SdkV2 as
// a VolumesStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo_SdkV2) GetVolumes(ctx context.Context) (VolumesStorageInfo_SdkV2, bool) {
	var e VolumesStorageInfo_SdkV2
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return e, false
	}
	var v []VolumesStorageInfo_SdkV2
	d := o.Volumes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumes sets the value of the Volumes field in InitScriptInfo_SdkV2.
func (o *InitScriptInfo_SdkV2) SetVolumes(ctx context.Context, v VolumesStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["volumes"]
	o.Volumes = types.ListValueMust(t, vs)
}

// GetWorkspace returns the value of the Workspace field in InitScriptInfo_SdkV2 as
// a WorkspaceStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfo_SdkV2) GetWorkspace(ctx context.Context) (WorkspaceStorageInfo_SdkV2, bool) {
	var e WorkspaceStorageInfo_SdkV2
	if o.Workspace.IsNull() || o.Workspace.IsUnknown() {
		return e, false
	}
	var v []WorkspaceStorageInfo_SdkV2
	d := o.Workspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspace sets the value of the Workspace field in InitScriptInfo_SdkV2.
func (o *InitScriptInfo_SdkV2) SetWorkspace(ctx context.Context, v WorkspaceStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace"]
	o.Workspace = types.ListValueMust(t, vs)
}

type InitScriptInfoAndExecutionDetails_SdkV2 struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	Abfss types.List `tfsdk:"abfss"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs types.List `tfsdk:"dbfs"`
	// Additional details regarding errors (such as a file not found message if
	// the status is FAILED_FETCH). This field should only be used to provide
	// *additional* information to the status field, not duplicate it.
	ErrorMessage types.String `tfsdk:"error_message"`
	// The number duration of the script execution in seconds
	ExecutionDurationSeconds types.Int64 `tfsdk:"execution_duration_seconds"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File types.List `tfsdk:"file"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs types.List `tfsdk:"gcs"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 types.List `tfsdk:"s3"`
	// The current status of the script
	Status types.String `tfsdk:"status"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes types.List `tfsdk:"volumes"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace types.List `tfsdk:"workspace"`
}

func (newState *InitScriptInfoAndExecutionDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InitScriptInfoAndExecutionDetails_SdkV2) {
}

func (newState *InitScriptInfoAndExecutionDetails_SdkV2) SyncFieldsDuringRead(existingState InitScriptInfoAndExecutionDetails_SdkV2) {
}

func (c InitScriptInfoAndExecutionDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["abfss"] = attrs["abfss"].SetOptional()
	attrs["abfss"] = attrs["abfss"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["dbfs"] = attrs["dbfs"].SetOptional()
	attrs["dbfs"] = attrs["dbfs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["execution_duration_seconds"] = attrs["execution_duration_seconds"].SetOptional()
	attrs["file"] = attrs["file"].SetOptional()
	attrs["file"] = attrs["file"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["gcs"] = attrs["gcs"].SetOptional()
	attrs["gcs"] = attrs["gcs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["s3"] = attrs["s3"].SetOptional()
	attrs["s3"] = attrs["s3"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetOptional()
	attrs["volumes"] = attrs["volumes"].SetOptional()
	attrs["volumes"] = attrs["volumes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["workspace"] = attrs["workspace"].SetOptional()
	attrs["workspace"] = attrs["workspace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InitScriptInfoAndExecutionDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InitScriptInfoAndExecutionDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"abfss":     reflect.TypeOf(Adlsgen2Info_SdkV2{}),
		"dbfs":      reflect.TypeOf(DbfsStorageInfo_SdkV2{}),
		"file":      reflect.TypeOf(LocalFileInfo_SdkV2{}),
		"gcs":       reflect.TypeOf(GcsStorageInfo_SdkV2{}),
		"s3":        reflect.TypeOf(S3StorageInfo_SdkV2{}),
		"volumes":   reflect.TypeOf(VolumesStorageInfo_SdkV2{}),
		"workspace": reflect.TypeOf(WorkspaceStorageInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InitScriptInfoAndExecutionDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o InitScriptInfoAndExecutionDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o InitScriptInfoAndExecutionDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"abfss": basetypes.ListType{
				ElemType: Adlsgen2Info_SdkV2{}.Type(ctx),
			},
			"dbfs": basetypes.ListType{
				ElemType: DbfsStorageInfo_SdkV2{}.Type(ctx),
			},
			"error_message":              types.StringType,
			"execution_duration_seconds": types.Int64Type,
			"file": basetypes.ListType{
				ElemType: LocalFileInfo_SdkV2{}.Type(ctx),
			},
			"gcs": basetypes.ListType{
				ElemType: GcsStorageInfo_SdkV2{}.Type(ctx),
			},
			"s3": basetypes.ListType{
				ElemType: S3StorageInfo_SdkV2{}.Type(ctx),
			},
			"status": types.StringType,
			"volumes": basetypes.ListType{
				ElemType: VolumesStorageInfo_SdkV2{}.Type(ctx),
			},
			"workspace": basetypes.ListType{
				ElemType: WorkspaceStorageInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAbfss returns the value of the Abfss field in InitScriptInfoAndExecutionDetails_SdkV2 as
// a Adlsgen2Info_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) GetAbfss(ctx context.Context) (Adlsgen2Info_SdkV2, bool) {
	var e Adlsgen2Info_SdkV2
	if o.Abfss.IsNull() || o.Abfss.IsUnknown() {
		return e, false
	}
	var v []Adlsgen2Info_SdkV2
	d := o.Abfss.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAbfss sets the value of the Abfss field in InitScriptInfoAndExecutionDetails_SdkV2.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) SetAbfss(ctx context.Context, v Adlsgen2Info_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["abfss"]
	o.Abfss = types.ListValueMust(t, vs)
}

// GetDbfs returns the value of the Dbfs field in InitScriptInfoAndExecutionDetails_SdkV2 as
// a DbfsStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) GetDbfs(ctx context.Context) (DbfsStorageInfo_SdkV2, bool) {
	var e DbfsStorageInfo_SdkV2
	if o.Dbfs.IsNull() || o.Dbfs.IsUnknown() {
		return e, false
	}
	var v []DbfsStorageInfo_SdkV2
	d := o.Dbfs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbfs sets the value of the Dbfs field in InitScriptInfoAndExecutionDetails_SdkV2.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) SetDbfs(ctx context.Context, v DbfsStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbfs"]
	o.Dbfs = types.ListValueMust(t, vs)
}

// GetFile returns the value of the File field in InitScriptInfoAndExecutionDetails_SdkV2 as
// a LocalFileInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) GetFile(ctx context.Context) (LocalFileInfo_SdkV2, bool) {
	var e LocalFileInfo_SdkV2
	if o.File.IsNull() || o.File.IsUnknown() {
		return e, false
	}
	var v []LocalFileInfo_SdkV2
	d := o.File.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFile sets the value of the File field in InitScriptInfoAndExecutionDetails_SdkV2.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) SetFile(ctx context.Context, v LocalFileInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file"]
	o.File = types.ListValueMust(t, vs)
}

// GetGcs returns the value of the Gcs field in InitScriptInfoAndExecutionDetails_SdkV2 as
// a GcsStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) GetGcs(ctx context.Context) (GcsStorageInfo_SdkV2, bool) {
	var e GcsStorageInfo_SdkV2
	if o.Gcs.IsNull() || o.Gcs.IsUnknown() {
		return e, false
	}
	var v []GcsStorageInfo_SdkV2
	d := o.Gcs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcs sets the value of the Gcs field in InitScriptInfoAndExecutionDetails_SdkV2.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) SetGcs(ctx context.Context, v GcsStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcs"]
	o.Gcs = types.ListValueMust(t, vs)
}

// GetS3 returns the value of the S3 field in InitScriptInfoAndExecutionDetails_SdkV2 as
// a S3StorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) GetS3(ctx context.Context) (S3StorageInfo_SdkV2, bool) {
	var e S3StorageInfo_SdkV2
	if o.S3.IsNull() || o.S3.IsUnknown() {
		return e, false
	}
	var v []S3StorageInfo_SdkV2
	d := o.S3.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetS3 sets the value of the S3 field in InitScriptInfoAndExecutionDetails_SdkV2.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) SetS3(ctx context.Context, v S3StorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["s3"]
	o.S3 = types.ListValueMust(t, vs)
}

// GetVolumes returns the value of the Volumes field in InitScriptInfoAndExecutionDetails_SdkV2 as
// a VolumesStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) GetVolumes(ctx context.Context) (VolumesStorageInfo_SdkV2, bool) {
	var e VolumesStorageInfo_SdkV2
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return e, false
	}
	var v []VolumesStorageInfo_SdkV2
	d := o.Volumes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumes sets the value of the Volumes field in InitScriptInfoAndExecutionDetails_SdkV2.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) SetVolumes(ctx context.Context, v VolumesStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["volumes"]
	o.Volumes = types.ListValueMust(t, vs)
}

// GetWorkspace returns the value of the Workspace field in InitScriptInfoAndExecutionDetails_SdkV2 as
// a WorkspaceStorageInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) GetWorkspace(ctx context.Context) (WorkspaceStorageInfo_SdkV2, bool) {
	var e WorkspaceStorageInfo_SdkV2
	if o.Workspace.IsNull() || o.Workspace.IsUnknown() {
		return e, false
	}
	var v []WorkspaceStorageInfo_SdkV2
	d := o.Workspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspace sets the value of the Workspace field in InitScriptInfoAndExecutionDetails_SdkV2.
func (o *InitScriptInfoAndExecutionDetails_SdkV2) SetWorkspace(ctx context.Context, v WorkspaceStorageInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace"]
	o.Workspace = types.ListValueMust(t, vs)
}

type InstallLibraries_SdkV2 struct {
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
func (a InstallLibraries_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallLibraries_SdkV2
// only implements ToObjectValue() and Type().
func (o InstallLibraries_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
			"libraries":  o.Libraries,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstallLibraries_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"libraries": basetypes.ListType{
				ElemType: Library_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetLibraries returns the value of the Libraries field in InstallLibraries_SdkV2 as
// a slice of Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstallLibraries_SdkV2) GetLibraries(ctx context.Context) ([]Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in InstallLibraries_SdkV2.
func (o *InstallLibraries_SdkV2) SetLibraries(ctx context.Context, v []Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type InstallLibrariesResponse_SdkV2 struct {
}

func (newState *InstallLibrariesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstallLibrariesResponse_SdkV2) {
}

func (newState *InstallLibrariesResponse_SdkV2) SyncFieldsDuringRead(existingState InstallLibrariesResponse_SdkV2) {
}

func (c InstallLibrariesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InstallLibrariesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InstallLibrariesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstallLibrariesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o InstallLibrariesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type InstancePoolAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *InstancePoolAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolAccessControlRequest_SdkV2) {
}

func (newState *InstancePoolAccessControlRequest_SdkV2) SyncFieldsDuringRead(existingState InstancePoolAccessControlRequest_SdkV2) {
}

func (c InstancePoolAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o InstancePoolAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type InstancePoolAccessControlResponse_SdkV2 struct {
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

func (newState *InstancePoolAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolAccessControlResponse_SdkV2) {
}

func (newState *InstancePoolAccessControlResponse_SdkV2) SyncFieldsDuringRead(existingState InstancePoolAccessControlResponse_SdkV2) {
}

func (c InstancePoolAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(InstancePoolPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o InstancePoolAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: InstancePoolPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in InstancePoolAccessControlResponse_SdkV2 as
// a slice of InstancePoolPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]InstancePoolPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in InstancePoolAccessControlResponse_SdkV2.
func (o *InstancePoolAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []InstancePoolPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type InstancePoolAndStats_SdkV2 struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
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
	DiskSpec types.List `tfsdk:"disk_spec"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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
	NodeTypeFlexibility types.List `tfsdk:"node_type_flexibility"`
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
	Stats types.List `tfsdk:"stats"`
	// Status of failed pending instances in the pool.
	Status types.List `tfsdk:"status"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize types.Int64 `tfsdk:"total_initial_remote_disk_size"`
}

func (newState *InstancePoolAndStats_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolAndStats_SdkV2) {
}

func (newState *InstancePoolAndStats_SdkV2) SyncFieldsDuringRead(existingState InstancePoolAndStats_SdkV2) {
}

func (c InstancePoolAndStats_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["default_tags"] = attrs["default_tags"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].SetOptional()
	attrs["disk_spec"] = attrs["disk_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["enable_auto_alternate_node_types"] = attrs["enable_auto_alternate_node_types"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["idle_instance_autotermination_minutes"] = attrs["idle_instance_autotermination_minutes"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["instance_pool_name"] = attrs["instance_pool_name"].SetOptional()
	attrs["max_capacity"] = attrs["max_capacity"].SetOptional()
	attrs["min_idle_instances"] = attrs["min_idle_instances"].SetOptional()
	attrs["node_type_flexibility"] = attrs["node_type_flexibility"].SetOptional()
	attrs["node_type_flexibility"] = attrs["node_type_flexibility"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["preloaded_docker_images"] = attrs["preloaded_docker_images"].SetOptional()
	attrs["preloaded_spark_versions"] = attrs["preloaded_spark_versions"].SetOptional()
	attrs["remote_disk_throughput"] = attrs["remote_disk_throughput"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["stats"] = attrs["stats"].SetOptional()
	attrs["stats"] = attrs["stats"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a InstancePoolAndStats_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_attributes":           reflect.TypeOf(InstancePoolAwsAttributes_SdkV2{}),
		"azure_attributes":         reflect.TypeOf(InstancePoolAzureAttributes_SdkV2{}),
		"custom_tags":              reflect.TypeOf(types.String{}),
		"default_tags":             reflect.TypeOf(types.String{}),
		"disk_spec":                reflect.TypeOf(DiskSpec_SdkV2{}),
		"gcp_attributes":           reflect.TypeOf(InstancePoolGcpAttributes_SdkV2{}),
		"node_type_flexibility":    reflect.TypeOf(NodeTypeFlexibility_SdkV2{}),
		"preloaded_docker_images":  reflect.TypeOf(DockerImage_SdkV2{}),
		"preloaded_spark_versions": reflect.TypeOf(types.String{}),
		"stats":                    reflect.TypeOf(InstancePoolStats_SdkV2{}),
		"status":                   reflect.TypeOf(InstancePoolStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAndStats_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolAndStats_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_attributes":                        o.AwsAttributes,
			"azure_attributes":                      o.AzureAttributes,
			"custom_tags":                           o.CustomTags,
			"default_tags":                          o.DefaultTags,
			"disk_spec":                             o.DiskSpec,
			"enable_auto_alternate_node_types":      o.EnableAutoAlternateNodeTypes,
			"enable_elastic_disk":                   o.EnableElasticDisk,
			"gcp_attributes":                        o.GcpAttributes,
			"idle_instance_autotermination_minutes": o.IdleInstanceAutoterminationMinutes,
			"instance_pool_id":                      o.InstancePoolId,
			"instance_pool_name":                    o.InstancePoolName,
			"max_capacity":                          o.MaxCapacity,
			"min_idle_instances":                    o.MinIdleInstances,
			"node_type_flexibility":                 o.NodeTypeFlexibility,
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
func (o InstancePoolAndStats_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes": basetypes.ListType{
				ElemType: InstancePoolAwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: InstancePoolAzureAttributes_SdkV2{}.Type(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"default_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec": basetypes.ListType{
				ElemType: DiskSpec_SdkV2{}.Type(ctx),
			},
			"enable_auto_alternate_node_types": types.BoolType,
			"enable_elastic_disk":              types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: InstancePoolGcpAttributes_SdkV2{}.Type(ctx),
			},
			"idle_instance_autotermination_minutes": types.Int64Type,
			"instance_pool_id":                      types.StringType,
			"instance_pool_name":                    types.StringType,
			"max_capacity":                          types.Int64Type,
			"min_idle_instances":                    types.Int64Type,
			"node_type_flexibility": basetypes.ListType{
				ElemType: NodeTypeFlexibility_SdkV2{}.Type(ctx),
			},
			"node_type_id": types.StringType,
			"preloaded_docker_images": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"preloaded_spark_versions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"remote_disk_throughput": types.Int64Type,
			"state":                  types.StringType,
			"stats": basetypes.ListType{
				ElemType: InstancePoolStats_SdkV2{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: InstancePoolStatus_SdkV2{}.Type(ctx),
			},
			"total_initial_remote_disk_size": types.Int64Type,
		},
	}
}

// GetAwsAttributes returns the value of the AwsAttributes field in InstancePoolAndStats_SdkV2 as
// a InstancePoolAwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetAwsAttributes(ctx context.Context) (InstancePoolAwsAttributes_SdkV2, bool) {
	var e InstancePoolAwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetAwsAttributes(ctx context.Context, v InstancePoolAwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in InstancePoolAndStats_SdkV2 as
// a InstancePoolAzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetAzureAttributes(ctx context.Context) (InstancePoolAzureAttributes_SdkV2, bool) {
	var e InstancePoolAzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolAzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetAzureAttributes(ctx context.Context, v InstancePoolAzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in InstancePoolAndStats_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDefaultTags returns the value of the DefaultTags field in InstancePoolAndStats_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetDefaultTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetDefaultTags sets the value of the DefaultTags field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetDefaultTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DefaultTags = types.MapValueMust(t, vs)
}

// GetDiskSpec returns the value of the DiskSpec field in InstancePoolAndStats_SdkV2 as
// a DiskSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetDiskSpec(ctx context.Context) (DiskSpec_SdkV2, bool) {
	var e DiskSpec_SdkV2
	if o.DiskSpec.IsNull() || o.DiskSpec.IsUnknown() {
		return e, false
	}
	var v []DiskSpec_SdkV2
	d := o.DiskSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDiskSpec sets the value of the DiskSpec field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetDiskSpec(ctx context.Context, v DiskSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["disk_spec"]
	o.DiskSpec = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in InstancePoolAndStats_SdkV2 as
// a InstancePoolGcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetGcpAttributes(ctx context.Context) (InstancePoolGcpAttributes_SdkV2, bool) {
	var e InstancePoolGcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []InstancePoolGcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetGcpAttributes(ctx context.Context, v InstancePoolGcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetNodeTypeFlexibility returns the value of the NodeTypeFlexibility field in InstancePoolAndStats_SdkV2 as
// a NodeTypeFlexibility_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetNodeTypeFlexibility(ctx context.Context) (NodeTypeFlexibility_SdkV2, bool) {
	var e NodeTypeFlexibility_SdkV2
	if o.NodeTypeFlexibility.IsNull() || o.NodeTypeFlexibility.IsUnknown() {
		return e, false
	}
	var v []NodeTypeFlexibility_SdkV2
	d := o.NodeTypeFlexibility.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeTypeFlexibility sets the value of the NodeTypeFlexibility field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetNodeTypeFlexibility(ctx context.Context, v NodeTypeFlexibility_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_type_flexibility"]
	o.NodeTypeFlexibility = types.ListValueMust(t, vs)
}

// GetPreloadedDockerImages returns the value of the PreloadedDockerImages field in InstancePoolAndStats_SdkV2 as
// a slice of DockerImage_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetPreloadedDockerImages(ctx context.Context) ([]DockerImage_SdkV2, bool) {
	if o.PreloadedDockerImages.IsNull() || o.PreloadedDockerImages.IsUnknown() {
		return nil, false
	}
	var v []DockerImage_SdkV2
	d := o.PreloadedDockerImages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPreloadedDockerImages sets the value of the PreloadedDockerImages field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetPreloadedDockerImages(ctx context.Context, v []DockerImage_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_docker_images"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedDockerImages = types.ListValueMust(t, vs)
}

// GetPreloadedSparkVersions returns the value of the PreloadedSparkVersions field in InstancePoolAndStats_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetPreloadedSparkVersions(ctx context.Context) ([]types.String, bool) {
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

// SetPreloadedSparkVersions sets the value of the PreloadedSparkVersions field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetPreloadedSparkVersions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["preloaded_spark_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PreloadedSparkVersions = types.ListValueMust(t, vs)
}

// GetStats returns the value of the Stats field in InstancePoolAndStats_SdkV2 as
// a InstancePoolStats_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetStats(ctx context.Context) (InstancePoolStats_SdkV2, bool) {
	var e InstancePoolStats_SdkV2
	if o.Stats.IsNull() || o.Stats.IsUnknown() {
		return e, false
	}
	var v []InstancePoolStats_SdkV2
	d := o.Stats.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStats sets the value of the Stats field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetStats(ctx context.Context, v InstancePoolStats_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stats"]
	o.Stats = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in InstancePoolAndStats_SdkV2 as
// a InstancePoolStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolAndStats_SdkV2) GetStatus(ctx context.Context) (InstancePoolStatus_SdkV2, bool) {
	var e InstancePoolStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []InstancePoolStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in InstancePoolAndStats_SdkV2.
func (o *InstancePoolAndStats_SdkV2) SetStatus(ctx context.Context, v InstancePoolStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

// Attributes set during instance pool creation which are related to Amazon Web
// Services.
type InstancePoolAwsAttributes_SdkV2 struct {
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

func (newState *InstancePoolAwsAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolAwsAttributes_SdkV2) {
}

func (newState *InstancePoolAwsAttributes_SdkV2) SyncFieldsDuringRead(existingState InstancePoolAwsAttributes_SdkV2) {
}

func (c InstancePoolAwsAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolAwsAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAwsAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolAwsAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":           o.Availability,
			"spot_bid_price_percent": o.SpotBidPricePercent,
			"zone_id":                o.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolAwsAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":           types.StringType,
			"spot_bid_price_percent": types.Int64Type,
			"zone_id":                types.StringType,
		},
	}
}

// Attributes set during instance pool creation which are related to Azure.
type InstancePoolAzureAttributes_SdkV2 struct {
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

func (newState *InstancePoolAzureAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolAzureAttributes_SdkV2) {
}

func (newState *InstancePoolAzureAttributes_SdkV2) SyncFieldsDuringRead(existingState InstancePoolAzureAttributes_SdkV2) {
}

func (c InstancePoolAzureAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolAzureAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolAzureAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolAzureAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability":       o.Availability,
			"spot_bid_max_price": o.SpotBidMaxPrice,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolAzureAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":       types.StringType,
			"spot_bid_max_price": types.Float64Type,
		},
	}
}

// Attributes set during instance pool creation which are related to GCP.
type InstancePoolGcpAttributes_SdkV2 struct {
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

func (newState *InstancePoolGcpAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolGcpAttributes_SdkV2) {
}

func (newState *InstancePoolGcpAttributes_SdkV2) SyncFieldsDuringRead(existingState InstancePoolGcpAttributes_SdkV2) {
}

func (c InstancePoolGcpAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolGcpAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolGcpAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolGcpAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_availability": o.GcpAvailability,
			"local_ssd_count":  o.LocalSsdCount,
			"zone_id":          o.ZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolGcpAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_availability": types.StringType,
			"local_ssd_count":  types.Int64Type,
			"zone_id":          types.StringType,
		},
	}
}

type InstancePoolPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *InstancePoolPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolPermission_SdkV2) {
}

func (newState *InstancePoolPermission_SdkV2) SyncFieldsDuringRead(existingState InstancePoolPermission_SdkV2) {
}

func (c InstancePoolPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in InstancePoolPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in InstancePoolPermission_SdkV2.
func (o *InstancePoolPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type InstancePoolPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *InstancePoolPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolPermissions_SdkV2) {
}

func (newState *InstancePoolPermissions_SdkV2) SyncFieldsDuringRead(existingState InstancePoolPermissions_SdkV2) {
}

func (c InstancePoolPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(InstancePoolAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: InstancePoolAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in InstancePoolPermissions_SdkV2 as
// a slice of InstancePoolAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]InstancePoolAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in InstancePoolPermissions_SdkV2.
func (o *InstancePoolPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []InstancePoolAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type InstancePoolPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *InstancePoolPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolPermissionsDescription_SdkV2) {
}

func (newState *InstancePoolPermissionsDescription_SdkV2) SyncFieldsDuringRead(existingState InstancePoolPermissionsDescription_SdkV2) {
}

func (c InstancePoolPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type InstancePoolPermissionsRequest_SdkV2 struct {
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
func (a InstancePoolPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(InstancePoolAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"instance_pool_id":    o.InstancePoolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: InstancePoolAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"instance_pool_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in InstancePoolPermissionsRequest_SdkV2 as
// a slice of InstancePoolAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]InstancePoolAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in InstancePoolPermissionsRequest_SdkV2.
func (o *InstancePoolPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []InstancePoolAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type InstancePoolStats_SdkV2 struct {
	// Number of active instances in the pool that are NOT part of a cluster.
	IdleCount types.Int64 `tfsdk:"idle_count"`
	// Number of pending instances in the pool that are NOT part of a cluster.
	PendingIdleCount types.Int64 `tfsdk:"pending_idle_count"`
	// Number of pending instances in the pool that are part of a cluster.
	PendingUsedCount types.Int64 `tfsdk:"pending_used_count"`
	// Number of active instances in the pool that are part of a cluster.
	UsedCount types.Int64 `tfsdk:"used_count"`
}

func (newState *InstancePoolStats_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolStats_SdkV2) {
}

func (newState *InstancePoolStats_SdkV2) SyncFieldsDuringRead(existingState InstancePoolStats_SdkV2) {
}

func (c InstancePoolStats_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolStats_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolStats_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolStats_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o InstancePoolStats_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"idle_count":         types.Int64Type,
			"pending_idle_count": types.Int64Type,
			"pending_used_count": types.Int64Type,
			"used_count":         types.Int64Type,
		},
	}
}

type InstancePoolStatus_SdkV2 struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	PendingInstanceErrors types.List `tfsdk:"pending_instance_errors"`
}

func (newState *InstancePoolStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstancePoolStatus_SdkV2) {
}

func (newState *InstancePoolStatus_SdkV2) SyncFieldsDuringRead(existingState InstancePoolStatus_SdkV2) {
}

func (c InstancePoolStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstancePoolStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pending_instance_errors": reflect.TypeOf(PendingInstanceError_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstancePoolStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o InstancePoolStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pending_instance_errors": o.PendingInstanceErrors,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pending_instance_errors": basetypes.ListType{
				ElemType: PendingInstanceError_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPendingInstanceErrors returns the value of the PendingInstanceErrors field in InstancePoolStatus_SdkV2 as
// a slice of PendingInstanceError_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *InstancePoolStatus_SdkV2) GetPendingInstanceErrors(ctx context.Context) ([]PendingInstanceError_SdkV2, bool) {
	if o.PendingInstanceErrors.IsNull() || o.PendingInstanceErrors.IsUnknown() {
		return nil, false
	}
	var v []PendingInstanceError_SdkV2
	d := o.PendingInstanceErrors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPendingInstanceErrors sets the value of the PendingInstanceErrors field in InstancePoolStatus_SdkV2.
func (o *InstancePoolStatus_SdkV2) SetPendingInstanceErrors(ctx context.Context, v []PendingInstanceError_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pending_instance_errors"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PendingInstanceErrors = types.ListValueMust(t, vs)
}

type InstanceProfile_SdkV2 struct {
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

func (newState *InstanceProfile_SdkV2) SyncFieldsDuringCreateOrUpdate(plan InstanceProfile_SdkV2) {
}

func (newState *InstanceProfile_SdkV2) SyncFieldsDuringRead(existingState InstanceProfile_SdkV2) {
}

func (c InstanceProfile_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InstanceProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InstanceProfile_SdkV2
// only implements ToObjectValue() and Type().
func (o InstanceProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"iam_role_arn":             o.IamRoleArn,
			"instance_profile_arn":     o.InstanceProfileArn,
			"is_meta_instance_profile": o.IsMetaInstanceProfile,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InstanceProfile_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"iam_role_arn":             types.StringType,
			"instance_profile_arn":     types.StringType,
			"is_meta_instance_profile": types.BoolType,
		},
	}
}

type Library_SdkV2 struct {
	// Specification of a CRAN library to be installed as part of the library
	Cran types.List `tfsdk:"cran"`
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
	Maven types.List `tfsdk:"maven"`
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	Pypi types.List `tfsdk:"pypi"`
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

func (newState *Library_SdkV2) SyncFieldsDuringCreateOrUpdate(plan Library_SdkV2) {
}

func (newState *Library_SdkV2) SyncFieldsDuringRead(existingState Library_SdkV2) {
}

func (c Library_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cran"] = attrs["cran"].SetOptional()
	attrs["cran"] = attrs["cran"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["egg"] = attrs["egg"].SetOptional()
	attrs["jar"] = attrs["jar"].SetOptional()
	attrs["maven"] = attrs["maven"].SetOptional()
	attrs["maven"] = attrs["maven"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["pypi"] = attrs["pypi"].SetOptional()
	attrs["pypi"] = attrs["pypi"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a Library_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cran":  reflect.TypeOf(RCranLibrary_SdkV2{}),
		"maven": reflect.TypeOf(MavenLibrary_SdkV2{}),
		"pypi":  reflect.TypeOf(PythonPyPiLibrary_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Library_SdkV2
// only implements ToObjectValue() and Type().
func (o Library_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Library_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cran": basetypes.ListType{
				ElemType: RCranLibrary_SdkV2{}.Type(ctx),
			},
			"egg": types.StringType,
			"jar": types.StringType,
			"maven": basetypes.ListType{
				ElemType: MavenLibrary_SdkV2{}.Type(ctx),
			},
			"pypi": basetypes.ListType{
				ElemType: PythonPyPiLibrary_SdkV2{}.Type(ctx),
			},
			"requirements": types.StringType,
			"whl":          types.StringType,
		},
	}
}

// GetCran returns the value of the Cran field in Library_SdkV2 as
// a RCranLibrary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Library_SdkV2) GetCran(ctx context.Context) (RCranLibrary_SdkV2, bool) {
	var e RCranLibrary_SdkV2
	if o.Cran.IsNull() || o.Cran.IsUnknown() {
		return e, false
	}
	var v []RCranLibrary_SdkV2
	d := o.Cran.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCran sets the value of the Cran field in Library_SdkV2.
func (o *Library_SdkV2) SetCran(ctx context.Context, v RCranLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cran"]
	o.Cran = types.ListValueMust(t, vs)
}

// GetMaven returns the value of the Maven field in Library_SdkV2 as
// a MavenLibrary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Library_SdkV2) GetMaven(ctx context.Context) (MavenLibrary_SdkV2, bool) {
	var e MavenLibrary_SdkV2
	if o.Maven.IsNull() || o.Maven.IsUnknown() {
		return e, false
	}
	var v []MavenLibrary_SdkV2
	d := o.Maven.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMaven sets the value of the Maven field in Library_SdkV2.
func (o *Library_SdkV2) SetMaven(ctx context.Context, v MavenLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["maven"]
	o.Maven = types.ListValueMust(t, vs)
}

// GetPypi returns the value of the Pypi field in Library_SdkV2 as
// a PythonPyPiLibrary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Library_SdkV2) GetPypi(ctx context.Context) (PythonPyPiLibrary_SdkV2, bool) {
	var e PythonPyPiLibrary_SdkV2
	if o.Pypi.IsNull() || o.Pypi.IsUnknown() {
		return e, false
	}
	var v []PythonPyPiLibrary_SdkV2
	d := o.Pypi.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPypi sets the value of the Pypi field in Library_SdkV2.
func (o *Library_SdkV2) SetPypi(ctx context.Context, v PythonPyPiLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pypi"]
	o.Pypi = types.ListValueMust(t, vs)
}

// The status of the library on a specific cluster.
type LibraryFullStatus_SdkV2 struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	IsLibraryForAllClusters types.Bool `tfsdk:"is_library_for_all_clusters"`
	// Unique identifier for the library.
	Library types.List `tfsdk:"library"`
	// All the info and warning messages that have occurred so far for this
	// library.
	Messages types.List `tfsdk:"messages"`
	// Status of installing the library on the cluster.
	Status types.String `tfsdk:"status"`
}

func (newState *LibraryFullStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(plan LibraryFullStatus_SdkV2) {
}

func (newState *LibraryFullStatus_SdkV2) SyncFieldsDuringRead(existingState LibraryFullStatus_SdkV2) {
}

func (c LibraryFullStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_library_for_all_clusters"] = attrs["is_library_for_all_clusters"].SetOptional()
	attrs["library"] = attrs["library"].SetOptional()
	attrs["library"] = attrs["library"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a LibraryFullStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"library":  reflect.TypeOf(Library_SdkV2{}),
		"messages": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LibraryFullStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o LibraryFullStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LibraryFullStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_library_for_all_clusters": types.BoolType,
			"library": basetypes.ListType{
				ElemType: Library_SdkV2{}.Type(ctx),
			},
			"messages": basetypes.ListType{
				ElemType: types.StringType,
			},
			"status": types.StringType,
		},
	}
}

// GetLibrary returns the value of the Library field in LibraryFullStatus_SdkV2 as
// a Library_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LibraryFullStatus_SdkV2) GetLibrary(ctx context.Context) (Library_SdkV2, bool) {
	var e Library_SdkV2
	if o.Library.IsNull() || o.Library.IsUnknown() {
		return e, false
	}
	var v []Library_SdkV2
	d := o.Library.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLibrary sets the value of the Library field in LibraryFullStatus_SdkV2.
func (o *LibraryFullStatus_SdkV2) SetLibrary(ctx context.Context, v Library_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["library"]
	o.Library = types.ListValueMust(t, vs)
}

// GetMessages returns the value of the Messages field in LibraryFullStatus_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *LibraryFullStatus_SdkV2) GetMessages(ctx context.Context) ([]types.String, bool) {
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

// SetMessages sets the value of the Messages field in LibraryFullStatus_SdkV2.
func (o *LibraryFullStatus_SdkV2) SetMessages(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Messages = types.ListValueMust(t, vs)
}

type ListAllClusterLibraryStatuses_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAllClusterLibraryStatuses.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAllClusterLibraryStatuses_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllClusterLibraryStatuses_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAllClusterLibraryStatuses_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatuses_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListAllClusterLibraryStatusesResponse_SdkV2 struct {
	// A list of cluster statuses.
	Statuses types.List `tfsdk:"statuses"`
}

func (newState *ListAllClusterLibraryStatusesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListAllClusterLibraryStatusesResponse_SdkV2) {
}

func (newState *ListAllClusterLibraryStatusesResponse_SdkV2) SyncFieldsDuringRead(existingState ListAllClusterLibraryStatusesResponse_SdkV2) {
}

func (c ListAllClusterLibraryStatusesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAllClusterLibraryStatusesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statuses": reflect.TypeOf(ClusterLibraryStatuses_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAllClusterLibraryStatusesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAllClusterLibraryStatusesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statuses": o.Statuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatusesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statuses": basetypes.ListType{
				ElemType: ClusterLibraryStatuses_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStatuses returns the value of the Statuses field in ListAllClusterLibraryStatusesResponse_SdkV2 as
// a slice of ClusterLibraryStatuses_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAllClusterLibraryStatusesResponse_SdkV2) GetStatuses(ctx context.Context) ([]ClusterLibraryStatuses_SdkV2, bool) {
	if o.Statuses.IsNull() || o.Statuses.IsUnknown() {
		return nil, false
	}
	var v []ClusterLibraryStatuses_SdkV2
	d := o.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in ListAllClusterLibraryStatusesResponse_SdkV2.
func (o *ListAllClusterLibraryStatusesResponse_SdkV2) SetStatuses(ctx context.Context, v []ClusterLibraryStatuses_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Statuses = types.ListValueMust(t, vs)
}

type ListAvailableZonesResponse_SdkV2 struct {
	// The availability zone if no ``zone_id`` is provided in the cluster
	// creation request.
	DefaultZone types.String `tfsdk:"default_zone"`
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	Zones types.List `tfsdk:"zones"`
}

func (newState *ListAvailableZonesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListAvailableZonesResponse_SdkV2) {
}

func (newState *ListAvailableZonesResponse_SdkV2) SyncFieldsDuringRead(existingState ListAvailableZonesResponse_SdkV2) {
}

func (c ListAvailableZonesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAvailableZonesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"zones": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAvailableZonesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAvailableZonesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_zone": o.DefaultZone,
			"zones":        o.Zones,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAvailableZonesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_zone": types.StringType,
			"zones": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetZones returns the value of the Zones field in ListAvailableZonesResponse_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAvailableZonesResponse_SdkV2) GetZones(ctx context.Context) ([]types.String, bool) {
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

// SetZones sets the value of the Zones field in ListAvailableZonesResponse_SdkV2.
func (o *ListAvailableZonesResponse_SdkV2) SetZones(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["zones"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Zones = types.ListValueMust(t, vs)
}

type ListClusterCompliancesRequest_SdkV2 struct {
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
func (a ListClusterCompliancesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterCompliancesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListClusterCompliancesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
			"policy_id":  o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClusterCompliancesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"policy_id":  types.StringType,
		},
	}
}

type ListClusterCompliancesResponse_SdkV2 struct {
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

func (newState *ListClusterCompliancesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListClusterCompliancesResponse_SdkV2) {
}

func (newState *ListClusterCompliancesResponse_SdkV2) SyncFieldsDuringRead(existingState ListClusterCompliancesResponse_SdkV2) {
}

func (c ListClusterCompliancesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListClusterCompliancesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters": reflect.TypeOf(ClusterCompliance_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterCompliancesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListClusterCompliancesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusters":        o.Clusters,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClusterCompliancesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusters": basetypes.ListType{
				ElemType: ClusterCompliance_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in ListClusterCompliancesResponse_SdkV2 as
// a slice of ClusterCompliance_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListClusterCompliancesResponse_SdkV2) GetClusters(ctx context.Context) ([]ClusterCompliance_SdkV2, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []ClusterCompliance_SdkV2
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in ListClusterCompliancesResponse_SdkV2.
func (o *ListClusterCompliancesResponse_SdkV2) SetClusters(ctx context.Context, v []ClusterCompliance_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

type ListClusterPoliciesRequest_SdkV2 struct {
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
func (a ListClusterPoliciesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClusterPoliciesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListClusterPoliciesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sort_column": o.SortColumn,
			"sort_order":  o.SortOrder,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClusterPoliciesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sort_column": types.StringType,
			"sort_order":  types.StringType,
		},
	}
}

type ListClustersFilterBy_SdkV2 struct {
	// The source of cluster creation.
	ClusterSources types.List `tfsdk:"cluster_sources"`
	// The current state of the clusters.
	ClusterStates types.List `tfsdk:"cluster_states"`
	// Whether the clusters are pinned or not.
	IsPinned types.Bool `tfsdk:"is_pinned"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
}

func (newState *ListClustersFilterBy_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListClustersFilterBy_SdkV2) {
}

func (newState *ListClustersFilterBy_SdkV2) SyncFieldsDuringRead(existingState ListClustersFilterBy_SdkV2) {
}

func (c ListClustersFilterBy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListClustersFilterBy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_sources": reflect.TypeOf(types.String{}),
		"cluster_states":  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersFilterBy_SdkV2
// only implements ToObjectValue() and Type().
func (o ListClustersFilterBy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListClustersFilterBy_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetClusterSources returns the value of the ClusterSources field in ListClustersFilterBy_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListClustersFilterBy_SdkV2) GetClusterSources(ctx context.Context) ([]types.String, bool) {
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

// SetClusterSources sets the value of the ClusterSources field in ListClustersFilterBy_SdkV2.
func (o *ListClustersFilterBy_SdkV2) SetClusterSources(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_sources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ClusterSources = types.ListValueMust(t, vs)
}

// GetClusterStates returns the value of the ClusterStates field in ListClustersFilterBy_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListClustersFilterBy_SdkV2) GetClusterStates(ctx context.Context) ([]types.String, bool) {
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

// SetClusterStates sets the value of the ClusterStates field in ListClustersFilterBy_SdkV2.
func (o *ListClustersFilterBy_SdkV2) SetClusterStates(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_states"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ClusterStates = types.ListValueMust(t, vs)
}

type ListClustersRequest_SdkV2 struct {
	// Filters to apply to the list of clusters.
	FilterBy types.List `tfsdk:"-"`
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize types.Int64 `tfsdk:"-"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of clusters respectively.
	PageToken types.String `tfsdk:"-"`
	// Sort the list of clusters by a specific criteria.
	SortBy types.List `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListClustersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListClustersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter_by": reflect.TypeOf(ListClustersFilterBy_SdkV2{}),
		"sort_by":   reflect.TypeOf(ListClustersSortBy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListClustersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListClustersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_by": basetypes.ListType{
				ElemType: ListClustersFilterBy_SdkV2{}.Type(ctx),
			},
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"sort_by": basetypes.ListType{
				ElemType: ListClustersSortBy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFilterBy returns the value of the FilterBy field in ListClustersRequest_SdkV2 as
// a ListClustersFilterBy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListClustersRequest_SdkV2) GetFilterBy(ctx context.Context) (ListClustersFilterBy_SdkV2, bool) {
	var e ListClustersFilterBy_SdkV2
	if o.FilterBy.IsNull() || o.FilterBy.IsUnknown() {
		return e, false
	}
	var v []ListClustersFilterBy_SdkV2
	d := o.FilterBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilterBy sets the value of the FilterBy field in ListClustersRequest_SdkV2.
func (o *ListClustersRequest_SdkV2) SetFilterBy(ctx context.Context, v ListClustersFilterBy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter_by"]
	o.FilterBy = types.ListValueMust(t, vs)
}

// GetSortBy returns the value of the SortBy field in ListClustersRequest_SdkV2 as
// a ListClustersSortBy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListClustersRequest_SdkV2) GetSortBy(ctx context.Context) (ListClustersSortBy_SdkV2, bool) {
	var e ListClustersSortBy_SdkV2
	if o.SortBy.IsNull() || o.SortBy.IsUnknown() {
		return e, false
	}
	var v []ListClustersSortBy_SdkV2
	d := o.SortBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSortBy sets the value of the SortBy field in ListClustersRequest_SdkV2.
func (o *ListClustersRequest_SdkV2) SetSortBy(ctx context.Context, v ListClustersSortBy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sort_by"]
	o.SortBy = types.ListValueMust(t, vs)
}

type ListClustersResponse_SdkV2 struct {
	Clusters types.List `tfsdk:"clusters"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken types.String `tfsdk:"prev_page_token"`
}

func (newState *ListClustersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListClustersResponse_SdkV2) {
}

func (newState *ListClustersResponse_SdkV2) SyncFieldsDuringRead(existingState ListClustersResponse_SdkV2) {
}

func (c ListClustersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListClustersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters": reflect.TypeOf(ClusterDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListClustersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clusters":        o.Clusters,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClustersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clusters": basetypes.ListType{
				ElemType: ClusterDetails_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in ListClustersResponse_SdkV2 as
// a slice of ClusterDetails_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListClustersResponse_SdkV2) GetClusters(ctx context.Context) ([]ClusterDetails_SdkV2, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []ClusterDetails_SdkV2
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in ListClustersResponse_SdkV2.
func (o *ListClustersResponse_SdkV2) SetClusters(ctx context.Context, v []ClusterDetails_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

type ListClustersSortBy_SdkV2 struct {
	// The direction to sort by.
	Direction types.String `tfsdk:"direction"`
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	Field types.String `tfsdk:"field"`
}

func (newState *ListClustersSortBy_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListClustersSortBy_SdkV2) {
}

func (newState *ListClustersSortBy_SdkV2) SyncFieldsDuringRead(existingState ListClustersSortBy_SdkV2) {
}

func (c ListClustersSortBy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListClustersSortBy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListClustersSortBy_SdkV2
// only implements ToObjectValue() and Type().
func (o ListClustersSortBy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direction": o.Direction,
			"field":     o.Field,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListClustersSortBy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direction": types.StringType,
			"field":     types.StringType,
		},
	}
}

type ListDefaultBaseEnvironmentsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDefaultBaseEnvironmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDefaultBaseEnvironmentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDefaultBaseEnvironmentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDefaultBaseEnvironmentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDefaultBaseEnvironmentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListDefaultBaseEnvironmentsResponse_SdkV2 struct {
	DefaultBaseEnvironments types.List `tfsdk:"default_base_environments"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListDefaultBaseEnvironmentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListDefaultBaseEnvironmentsResponse_SdkV2) {
}

func (newState *ListDefaultBaseEnvironmentsResponse_SdkV2) SyncFieldsDuringRead(existingState ListDefaultBaseEnvironmentsResponse_SdkV2) {
}

func (c ListDefaultBaseEnvironmentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDefaultBaseEnvironmentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_base_environments": reflect.TypeOf(DefaultBaseEnvironment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDefaultBaseEnvironmentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDefaultBaseEnvironmentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_base_environments": o.DefaultBaseEnvironments,
			"next_page_token":           o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDefaultBaseEnvironmentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_base_environments": basetypes.ListType{
				ElemType: DefaultBaseEnvironment_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDefaultBaseEnvironments returns the value of the DefaultBaseEnvironments field in ListDefaultBaseEnvironmentsResponse_SdkV2 as
// a slice of DefaultBaseEnvironment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListDefaultBaseEnvironmentsResponse_SdkV2) GetDefaultBaseEnvironments(ctx context.Context) ([]DefaultBaseEnvironment_SdkV2, bool) {
	if o.DefaultBaseEnvironments.IsNull() || o.DefaultBaseEnvironments.IsUnknown() {
		return nil, false
	}
	var v []DefaultBaseEnvironment_SdkV2
	d := o.DefaultBaseEnvironments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultBaseEnvironments sets the value of the DefaultBaseEnvironments field in ListDefaultBaseEnvironmentsResponse_SdkV2.
func (o *ListDefaultBaseEnvironmentsResponse_SdkV2) SetDefaultBaseEnvironments(ctx context.Context, v []DefaultBaseEnvironment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_base_environments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DefaultBaseEnvironments = types.ListValueMust(t, vs)
}

type ListGlobalInitScriptsRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGlobalInitScriptsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListGlobalInitScriptsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGlobalInitScriptsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListGlobalInitScriptsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListGlobalInitScriptsResponse_SdkV2 struct {
	Scripts types.List `tfsdk:"scripts"`
}

func (newState *ListGlobalInitScriptsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListGlobalInitScriptsResponse_SdkV2) {
}

func (newState *ListGlobalInitScriptsResponse_SdkV2) SyncFieldsDuringRead(existingState ListGlobalInitScriptsResponse_SdkV2) {
}

func (c ListGlobalInitScriptsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListGlobalInitScriptsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scripts": reflect.TypeOf(GlobalInitScriptDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGlobalInitScriptsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListGlobalInitScriptsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scripts": o.Scripts,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scripts": basetypes.ListType{
				ElemType: GlobalInitScriptDetails_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetScripts returns the value of the Scripts field in ListGlobalInitScriptsResponse_SdkV2 as
// a slice of GlobalInitScriptDetails_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListGlobalInitScriptsResponse_SdkV2) GetScripts(ctx context.Context) ([]GlobalInitScriptDetails_SdkV2, bool) {
	if o.Scripts.IsNull() || o.Scripts.IsUnknown() {
		return nil, false
	}
	var v []GlobalInitScriptDetails_SdkV2
	d := o.Scripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScripts sets the value of the Scripts field in ListGlobalInitScriptsResponse_SdkV2.
func (o *ListGlobalInitScriptsResponse_SdkV2) SetScripts(ctx context.Context, v []GlobalInitScriptDetails_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scripts = types.ListValueMust(t, vs)
}

type ListInstancePools_SdkV2 struct {
	InstancePools types.List `tfsdk:"instance_pools"`
}

func (newState *ListInstancePools_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListInstancePools_SdkV2) {
}

func (newState *ListInstancePools_SdkV2) SyncFieldsDuringRead(existingState ListInstancePools_SdkV2) {
}

func (c ListInstancePools_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListInstancePools_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"instance_pools": reflect.TypeOf(InstancePoolAndStats_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstancePools_SdkV2
// only implements ToObjectValue() and Type().
func (o ListInstancePools_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_pools": o.InstancePools,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstancePools_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pools": basetypes.ListType{
				ElemType: InstancePoolAndStats_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetInstancePools returns the value of the InstancePools field in ListInstancePools_SdkV2 as
// a slice of InstancePoolAndStats_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListInstancePools_SdkV2) GetInstancePools(ctx context.Context) ([]InstancePoolAndStats_SdkV2, bool) {
	if o.InstancePools.IsNull() || o.InstancePools.IsUnknown() {
		return nil, false
	}
	var v []InstancePoolAndStats_SdkV2
	d := o.InstancePools.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstancePools sets the value of the InstancePools field in ListInstancePools_SdkV2.
func (o *ListInstancePools_SdkV2) SetInstancePools(ctx context.Context, v []InstancePoolAndStats_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_pools"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InstancePools = types.ListValueMust(t, vs)
}

type ListInstancePoolsRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstancePoolsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListInstancePoolsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstancePoolsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListInstancePoolsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstancePoolsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListInstanceProfilesRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListInstanceProfilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListInstanceProfilesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstanceProfilesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListInstanceProfilesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstanceProfilesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListInstanceProfilesResponse_SdkV2 struct {
	// A list of instance profiles that the user can access.
	InstanceProfiles types.List `tfsdk:"instance_profiles"`
}

func (newState *ListInstanceProfilesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListInstanceProfilesResponse_SdkV2) {
}

func (newState *ListInstanceProfilesResponse_SdkV2) SyncFieldsDuringRead(existingState ListInstanceProfilesResponse_SdkV2) {
}

func (c ListInstanceProfilesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListInstanceProfilesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"instance_profiles": reflect.TypeOf(InstanceProfile_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListInstanceProfilesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListInstanceProfilesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_profiles": o.InstanceProfiles,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListInstanceProfilesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_profiles": basetypes.ListType{
				ElemType: InstanceProfile_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetInstanceProfiles returns the value of the InstanceProfiles field in ListInstanceProfilesResponse_SdkV2 as
// a slice of InstanceProfile_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListInstanceProfilesResponse_SdkV2) GetInstanceProfiles(ctx context.Context) ([]InstanceProfile_SdkV2, bool) {
	if o.InstanceProfiles.IsNull() || o.InstanceProfiles.IsUnknown() {
		return nil, false
	}
	var v []InstanceProfile_SdkV2
	d := o.InstanceProfiles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstanceProfiles sets the value of the InstanceProfiles field in ListInstanceProfilesResponse_SdkV2.
func (o *ListInstanceProfilesResponse_SdkV2) SetInstanceProfiles(ctx context.Context, v []InstanceProfile_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_profiles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InstanceProfiles = types.ListValueMust(t, vs)
}

type ListNodeTypesRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNodeTypesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNodeTypesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNodeTypesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNodeTypesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListNodeTypesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNodeTypesResponse_SdkV2 struct {
	// The list of available Spark node types.
	NodeTypes types.List `tfsdk:"node_types"`
}

func (newState *ListNodeTypesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListNodeTypesResponse_SdkV2) {
}

func (newState *ListNodeTypesResponse_SdkV2) SyncFieldsDuringRead(existingState ListNodeTypesResponse_SdkV2) {
}

func (c ListNodeTypesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNodeTypesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_types": reflect.TypeOf(NodeType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNodeTypesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNodeTypesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"node_types": o.NodeTypes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNodeTypesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"node_types": basetypes.ListType{
				ElemType: NodeType_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetNodeTypes returns the value of the NodeTypes field in ListNodeTypesResponse_SdkV2 as
// a slice of NodeType_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListNodeTypesResponse_SdkV2) GetNodeTypes(ctx context.Context) ([]NodeType_SdkV2, bool) {
	if o.NodeTypes.IsNull() || o.NodeTypes.IsUnknown() {
		return nil, false
	}
	var v []NodeType_SdkV2
	d := o.NodeTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNodeTypes sets the value of the NodeTypes field in ListNodeTypesResponse_SdkV2.
func (o *ListNodeTypesResponse_SdkV2) SetNodeTypes(ctx context.Context, v []NodeType_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NodeTypes = types.ListValueMust(t, vs)
}

type ListPoliciesResponse_SdkV2 struct {
	// List of policies.
	Policies types.List `tfsdk:"policies"`
}

func (newState *ListPoliciesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListPoliciesResponse_SdkV2) {
}

func (newState *ListPoliciesResponse_SdkV2) SyncFieldsDuringRead(existingState ListPoliciesResponse_SdkV2) {
}

func (c ListPoliciesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListPoliciesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(Policy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPoliciesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPoliciesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policies": o.Policies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPoliciesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policies": basetypes.ListType{
				ElemType: Policy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPolicies returns the value of the Policies field in ListPoliciesResponse_SdkV2 as
// a slice of Policy_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPoliciesResponse_SdkV2) GetPolicies(ctx context.Context) ([]Policy_SdkV2, bool) {
	if o.Policies.IsNull() || o.Policies.IsUnknown() {
		return nil, false
	}
	var v []Policy_SdkV2
	d := o.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListPoliciesResponse_SdkV2.
func (o *ListPoliciesResponse_SdkV2) SetPolicies(ctx context.Context, v []Policy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Policies = types.ListValueMust(t, vs)
}

type ListPolicyFamiliesRequest_SdkV2 struct {
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
func (a ListPolicyFamiliesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPolicyFamiliesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPolicyFamiliesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListPolicyFamiliesResponse_SdkV2 struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of policy families.
	PolicyFamilies types.List `tfsdk:"policy_families"`
}

func (newState *ListPolicyFamiliesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListPolicyFamiliesResponse_SdkV2) {
}

func (newState *ListPolicyFamiliesResponse_SdkV2) SyncFieldsDuringRead(existingState ListPolicyFamiliesResponse_SdkV2) {
}

func (c ListPolicyFamiliesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListPolicyFamiliesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy_families": reflect.TypeOf(PolicyFamily_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPolicyFamiliesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPolicyFamiliesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"policy_families": o.PolicyFamilies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"policy_families": basetypes.ListType{
				ElemType: PolicyFamily_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPolicyFamilies returns the value of the PolicyFamilies field in ListPolicyFamiliesResponse_SdkV2 as
// a slice of PolicyFamily_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPolicyFamiliesResponse_SdkV2) GetPolicyFamilies(ctx context.Context) ([]PolicyFamily_SdkV2, bool) {
	if o.PolicyFamilies.IsNull() || o.PolicyFamilies.IsUnknown() {
		return nil, false
	}
	var v []PolicyFamily_SdkV2
	d := o.PolicyFamilies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicyFamilies sets the value of the PolicyFamilies field in ListPolicyFamiliesResponse_SdkV2.
func (o *ListPolicyFamiliesResponse_SdkV2) SetPolicyFamilies(ctx context.Context, v []PolicyFamily_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["policy_families"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PolicyFamilies = types.ListValueMust(t, vs)
}

type ListZonesRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListZonesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListZonesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListZonesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListZonesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListZonesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LocalFileInfo_SdkV2 struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	Destination types.String `tfsdk:"destination"`
}

func (newState *LocalFileInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan LocalFileInfo_SdkV2) {
}

func (newState *LocalFileInfo_SdkV2) SyncFieldsDuringRead(existingState LocalFileInfo_SdkV2) {
}

func (c LocalFileInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LocalFileInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LocalFileInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o LocalFileInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LocalFileInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type LogAnalyticsInfo_SdkV2 struct {
	LogAnalyticsPrimaryKey types.String `tfsdk:"log_analytics_primary_key"`

	LogAnalyticsWorkspaceId types.String `tfsdk:"log_analytics_workspace_id"`
}

func (newState *LogAnalyticsInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan LogAnalyticsInfo_SdkV2) {
}

func (newState *LogAnalyticsInfo_SdkV2) SyncFieldsDuringRead(existingState LogAnalyticsInfo_SdkV2) {
}

func (c LogAnalyticsInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogAnalyticsInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogAnalyticsInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o LogAnalyticsInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_analytics_primary_key":  o.LogAnalyticsPrimaryKey,
			"log_analytics_workspace_id": o.LogAnalyticsWorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogAnalyticsInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_analytics_primary_key":  types.StringType,
			"log_analytics_workspace_id": types.StringType,
		},
	}
}

// The log delivery status
type LogSyncStatus_SdkV2 struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	LastAttempted types.Int64 `tfsdk:"last_attempted"`
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	LastException types.String `tfsdk:"last_exception"`
}

func (newState *LogSyncStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(plan LogSyncStatus_SdkV2) {
}

func (newState *LogSyncStatus_SdkV2) SyncFieldsDuringRead(existingState LogSyncStatus_SdkV2) {
}

func (c LogSyncStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogSyncStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogSyncStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o LogSyncStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_attempted": o.LastAttempted,
			"last_exception": o.LastException,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogSyncStatus_SdkV2) Type(ctx context.Context) attr.Type {
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
type MaterializedEnvironment_SdkV2 struct {
	// The timestamp (in epoch milliseconds) when the materialized env is
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
}

func (newState *MaterializedEnvironment_SdkV2) SyncFieldsDuringCreateOrUpdate(plan MaterializedEnvironment_SdkV2) {
}

func (newState *MaterializedEnvironment_SdkV2) SyncFieldsDuringRead(existingState MaterializedEnvironment_SdkV2) {
}

func (c MaterializedEnvironment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MaterializedEnvironment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MaterializedEnvironment_SdkV2
// only implements ToObjectValue() and Type().
func (o MaterializedEnvironment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_updated_timestamp": o.LastUpdatedTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MaterializedEnvironment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_updated_timestamp": types.Int64Type,
		},
	}
}

type MavenLibrary_SdkV2 struct {
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

func (newState *MavenLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(plan MavenLibrary_SdkV2) {
}

func (newState *MavenLibrary_SdkV2) SyncFieldsDuringRead(existingState MavenLibrary_SdkV2) {
}

func (c MavenLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MavenLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclusions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MavenLibrary_SdkV2
// only implements ToObjectValue() and Type().
func (o MavenLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"coordinates": o.Coordinates,
			"exclusions":  o.Exclusions,
			"repo":        o.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MavenLibrary_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetExclusions returns the value of the Exclusions field in MavenLibrary_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MavenLibrary_SdkV2) GetExclusions(ctx context.Context) ([]types.String, bool) {
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

// SetExclusions sets the value of the Exclusions field in MavenLibrary_SdkV2.
func (o *MavenLibrary_SdkV2) SetExclusions(ctx context.Context, v []types.String) {
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
type NodeInstanceType_SdkV2 struct {
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

func (newState *NodeInstanceType_SdkV2) SyncFieldsDuringCreateOrUpdate(plan NodeInstanceType_SdkV2) {
}

func (newState *NodeInstanceType_SdkV2) SyncFieldsDuringRead(existingState NodeInstanceType_SdkV2) {
}

func (c NodeInstanceType_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NodeInstanceType_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NodeInstanceType_SdkV2
// only implements ToObjectValue() and Type().
func (o NodeInstanceType_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o NodeInstanceType_SdkV2) Type(ctx context.Context) attr.Type {
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
type NodeType_SdkV2 struct {
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
	NodeInfo types.List `tfsdk:"node_info"`
	// The NodeInstanceType object corresponding to instance_type_id
	NodeInstanceType types.List `tfsdk:"node_instance_type"`
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

func (newState *NodeType_SdkV2) SyncFieldsDuringCreateOrUpdate(plan NodeType_SdkV2) {
}

func (newState *NodeType_SdkV2) SyncFieldsDuringRead(existingState NodeType_SdkV2) {
}

func (c NodeType_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["node_info"] = attrs["node_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["node_instance_type"] = attrs["node_instance_type"].SetOptional()
	attrs["node_instance_type"] = attrs["node_instance_type"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a NodeType_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_info":          reflect.TypeOf(CloudProviderNodeInfo_SdkV2{}),
		"node_instance_type": reflect.TypeOf(NodeInstanceType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NodeType_SdkV2
// only implements ToObjectValue() and Type().
func (o NodeType_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o NodeType_SdkV2) Type(ctx context.Context) attr.Type {
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
			"node_info": basetypes.ListType{
				ElemType: CloudProviderNodeInfo_SdkV2{}.Type(ctx),
			},
			"node_instance_type": basetypes.ListType{
				ElemType: NodeInstanceType_SdkV2{}.Type(ctx),
			},
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

// GetNodeInfo returns the value of the NodeInfo field in NodeType_SdkV2 as
// a CloudProviderNodeInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *NodeType_SdkV2) GetNodeInfo(ctx context.Context) (CloudProviderNodeInfo_SdkV2, bool) {
	var e CloudProviderNodeInfo_SdkV2
	if o.NodeInfo.IsNull() || o.NodeInfo.IsUnknown() {
		return e, false
	}
	var v []CloudProviderNodeInfo_SdkV2
	d := o.NodeInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeInfo sets the value of the NodeInfo field in NodeType_SdkV2.
func (o *NodeType_SdkV2) SetNodeInfo(ctx context.Context, v CloudProviderNodeInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_info"]
	o.NodeInfo = types.ListValueMust(t, vs)
}

// GetNodeInstanceType returns the value of the NodeInstanceType field in NodeType_SdkV2 as
// a NodeInstanceType_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *NodeType_SdkV2) GetNodeInstanceType(ctx context.Context) (NodeInstanceType_SdkV2, bool) {
	var e NodeInstanceType_SdkV2
	if o.NodeInstanceType.IsNull() || o.NodeInstanceType.IsUnknown() {
		return e, false
	}
	var v []NodeInstanceType_SdkV2
	d := o.NodeInstanceType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeInstanceType sets the value of the NodeInstanceType field in NodeType_SdkV2.
func (o *NodeType_SdkV2) SetNodeInstanceType(ctx context.Context, v NodeInstanceType_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_instance_type"]
	o.NodeInstanceType = types.ListValueMust(t, vs)
}

// For Fleet-V2 using classic clusters, this object contains the information
// about the alternate node type ids to use when attempting to launch a cluster.
// It can be used with both the driver and worker node types.
type NodeTypeFlexibility_SdkV2 struct {
}

func (newState *NodeTypeFlexibility_SdkV2) SyncFieldsDuringCreateOrUpdate(plan NodeTypeFlexibility_SdkV2) {
}

func (newState *NodeTypeFlexibility_SdkV2) SyncFieldsDuringRead(existingState NodeTypeFlexibility_SdkV2) {
}

func (c NodeTypeFlexibility_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NodeTypeFlexibility.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NodeTypeFlexibility_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NodeTypeFlexibility_SdkV2
// only implements ToObjectValue() and Type().
func (o NodeTypeFlexibility_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o NodeTypeFlexibility_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Error message of a failed pending instances
type PendingInstanceError_SdkV2 struct {
	InstanceId types.String `tfsdk:"instance_id"`

	Message types.String `tfsdk:"message"`
}

func (newState *PendingInstanceError_SdkV2) SyncFieldsDuringCreateOrUpdate(plan PendingInstanceError_SdkV2) {
}

func (newState *PendingInstanceError_SdkV2) SyncFieldsDuringRead(existingState PendingInstanceError_SdkV2) {
}

func (c PendingInstanceError_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PendingInstanceError_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PendingInstanceError_SdkV2
// only implements ToObjectValue() and Type().
func (o PendingInstanceError_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_id": o.InstanceId,
			"message":     o.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PendingInstanceError_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_id": types.StringType,
			"message":     types.StringType,
		},
	}
}

type PermanentDeleteCluster_SdkV2 struct {
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
func (a PermanentDeleteCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermanentDeleteCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o PermanentDeleteCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermanentDeleteCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type PermanentDeleteClusterResponse_SdkV2 struct {
}

func (newState *PermanentDeleteClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan PermanentDeleteClusterResponse_SdkV2) {
}

func (newState *PermanentDeleteClusterResponse_SdkV2) SyncFieldsDuringRead(existingState PermanentDeleteClusterResponse_SdkV2) {
}

func (c PermanentDeleteClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermanentDeleteClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermanentDeleteClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermanentDeleteClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PermanentDeleteClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type PinCluster_SdkV2 struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PinCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PinCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PinCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o PinCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PinCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type PinClusterResponse_SdkV2 struct {
}

func (newState *PinClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan PinClusterResponse_SdkV2) {
}

func (newState *PinClusterResponse_SdkV2) SyncFieldsDuringRead(existingState PinClusterResponse_SdkV2) {
}

func (c PinClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PinClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PinClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PinClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PinClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PinClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Describes a Cluster Policy entity.
type Policy_SdkV2 struct {
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

func (newState *Policy_SdkV2) SyncFieldsDuringCreateOrUpdate(plan Policy_SdkV2) {
}

func (newState *Policy_SdkV2) SyncFieldsDuringRead(existingState Policy_SdkV2) {
}

func (c Policy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Policy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Policy_SdkV2
// only implements ToObjectValue() and Type().
func (o Policy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Policy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at_timestamp": types.Int64Type,
			"creator_user_name":    types.StringType,
			"definition":           types.StringType,
			"description":          types.StringType,
			"is_default":           types.BoolType,
			"libraries": basetypes.ListType{
				ElemType: Library_SdkV2{}.Type(ctx),
			},
			"max_clusters_per_user":              types.Int64Type,
			"name":                               types.StringType,
			"policy_family_definition_overrides": types.StringType,
			"policy_family_id":                   types.StringType,
			"policy_id":                          types.StringType,
		},
	}
}

// GetLibraries returns the value of the Libraries field in Policy_SdkV2 as
// a slice of Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Policy_SdkV2) GetLibraries(ctx context.Context) ([]Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in Policy_SdkV2.
func (o *Policy_SdkV2) SetLibraries(ctx context.Context, v []Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type PolicyFamily_SdkV2 struct {
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

func (newState *PolicyFamily_SdkV2) SyncFieldsDuringCreateOrUpdate(plan PolicyFamily_SdkV2) {
}

func (newState *PolicyFamily_SdkV2) SyncFieldsDuringRead(existingState PolicyFamily_SdkV2) {
}

func (c PolicyFamily_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PolicyFamily_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PolicyFamily_SdkV2
// only implements ToObjectValue() and Type().
func (o PolicyFamily_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PolicyFamily_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"definition":       types.StringType,
			"description":      types.StringType,
			"name":             types.StringType,
			"policy_family_id": types.StringType,
		},
	}
}

type PythonPyPiLibrary_SdkV2 struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	Package types.String `tfsdk:"package"`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo types.String `tfsdk:"repo"`
}

func (newState *PythonPyPiLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(plan PythonPyPiLibrary_SdkV2) {
}

func (newState *PythonPyPiLibrary_SdkV2) SyncFieldsDuringRead(existingState PythonPyPiLibrary_SdkV2) {
}

func (c PythonPyPiLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PythonPyPiLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PythonPyPiLibrary_SdkV2
// only implements ToObjectValue() and Type().
func (o PythonPyPiLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"package": o.Package,
			"repo":    o.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PythonPyPiLibrary_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"package": types.StringType,
			"repo":    types.StringType,
		},
	}
}

type RCranLibrary_SdkV2 struct {
	// The name of the CRAN package to install.
	Package types.String `tfsdk:"package"`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo types.String `tfsdk:"repo"`
}

func (newState *RCranLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(plan RCranLibrary_SdkV2) {
}

func (newState *RCranLibrary_SdkV2) SyncFieldsDuringRead(existingState RCranLibrary_SdkV2) {
}

func (c RCranLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RCranLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RCranLibrary_SdkV2
// only implements ToObjectValue() and Type().
func (o RCranLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"package": o.Package,
			"repo":    o.Repo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RCranLibrary_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"package": types.StringType,
			"repo":    types.StringType,
		},
	}
}

type RefreshDefaultBaseEnvironmentsRequest_SdkV2 struct {
	Ids types.List `tfsdk:"ids"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RefreshDefaultBaseEnvironmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RefreshDefaultBaseEnvironmentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RefreshDefaultBaseEnvironmentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RefreshDefaultBaseEnvironmentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ids": o.Ids,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RefreshDefaultBaseEnvironmentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetIds returns the value of the Ids field in RefreshDefaultBaseEnvironmentsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RefreshDefaultBaseEnvironmentsRequest_SdkV2) GetIds(ctx context.Context) ([]types.String, bool) {
	if o.Ids.IsNull() || o.Ids.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Ids.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIds sets the value of the Ids field in RefreshDefaultBaseEnvironmentsRequest_SdkV2.
func (o *RefreshDefaultBaseEnvironmentsRequest_SdkV2) SetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Ids = types.ListValueMust(t, vs)
}

type RefreshDefaultBaseEnvironmentsResponse_SdkV2 struct {
}

func (newState *RefreshDefaultBaseEnvironmentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan RefreshDefaultBaseEnvironmentsResponse_SdkV2) {
}

func (newState *RefreshDefaultBaseEnvironmentsResponse_SdkV2) SyncFieldsDuringRead(existingState RefreshDefaultBaseEnvironmentsResponse_SdkV2) {
}

func (c RefreshDefaultBaseEnvironmentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RefreshDefaultBaseEnvironmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RefreshDefaultBaseEnvironmentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RefreshDefaultBaseEnvironmentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RefreshDefaultBaseEnvironmentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RefreshDefaultBaseEnvironmentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RemoveInstanceProfile_SdkV2 struct {
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
func (a RemoveInstanceProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveInstanceProfile_SdkV2
// only implements ToObjectValue() and Type().
func (o RemoveInstanceProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_profile_arn": o.InstanceProfileArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RemoveInstanceProfile_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_profile_arn": types.StringType,
		},
	}
}

type RemoveResponse_SdkV2 struct {
}

func (newState *RemoveResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan RemoveResponse_SdkV2) {
}

func (newState *RemoveResponse_SdkV2) SyncFieldsDuringRead(existingState RemoveResponse_SdkV2) {
}

func (c RemoveResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RemoveResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RemoveResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RemoveResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ResizeCluster_SdkV2 struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale"`
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
func (a ResizeCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale": reflect.TypeOf(AutoScale_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResizeCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o ResizeCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscale":   o.Autoscale,
			"cluster_id":  o.ClusterId,
			"num_workers": o.NumWorkers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResizeCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale": basetypes.ListType{
				ElemType: AutoScale_SdkV2{}.Type(ctx),
			},
			"cluster_id":  types.StringType,
			"num_workers": types.Int64Type,
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in ResizeCluster_SdkV2 as
// a AutoScale_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResizeCluster_SdkV2) GetAutoscale(ctx context.Context) (AutoScale_SdkV2, bool) {
	var e AutoScale_SdkV2
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale_SdkV2
	d := o.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in ResizeCluster_SdkV2.
func (o *ResizeCluster_SdkV2) SetAutoscale(ctx context.Context, v AutoScale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	o.Autoscale = types.ListValueMust(t, vs)
}

type ResizeClusterResponse_SdkV2 struct {
}

func (newState *ResizeClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ResizeClusterResponse_SdkV2) {
}

func (newState *ResizeClusterResponse_SdkV2) SyncFieldsDuringRead(existingState ResizeClusterResponse_SdkV2) {
}

func (c ResizeClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResizeClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResizeClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResizeClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ResizeClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ResizeClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestartCluster_SdkV2 struct {
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
func (a RestartCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o RestartCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":   o.ClusterId,
			"restart_user": o.RestartUser,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestartCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":   types.StringType,
			"restart_user": types.StringType,
		},
	}
}

type RestartClusterResponse_SdkV2 struct {
}

func (newState *RestartClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan RestartClusterResponse_SdkV2) {
}

func (newState *RestartClusterResponse_SdkV2) SyncFieldsDuringRead(existingState RestartClusterResponse_SdkV2) {
}

func (c RestartClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestartClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RestartClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RestartClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Results_SdkV2 struct {
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

func (newState *Results_SdkV2) SyncFieldsDuringCreateOrUpdate(plan Results_SdkV2) {
}

func (newState *Results_SdkV2) SyncFieldsDuringRead(existingState Results_SdkV2) {
}

func (c Results_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Results_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fileNames": reflect.TypeOf(types.String{}),
		"schema":    reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Results_SdkV2
// only implements ToObjectValue() and Type().
func (o Results_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Results_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetFileNames returns the value of the FileNames field in Results_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Results_SdkV2) GetFileNames(ctx context.Context) ([]types.String, bool) {
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

// SetFileNames sets the value of the FileNames field in Results_SdkV2.
func (o *Results_SdkV2) SetFileNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["fileNames"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FileNames = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in Results_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *Results_SdkV2) GetSchema(ctx context.Context) ([]types.Object, bool) {
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

// SetSchema sets the value of the Schema field in Results_SdkV2.
func (o *Results_SdkV2) SetSchema(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schema = types.ListValueMust(t, vs)
}

// A storage location in Amazon S3
type S3StorageInfo_SdkV2 struct {
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

func (newState *S3StorageInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan S3StorageInfo_SdkV2) {
}

func (newState *S3StorageInfo_SdkV2) SyncFieldsDuringRead(existingState S3StorageInfo_SdkV2) {
}

func (c S3StorageInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a S3StorageInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, S3StorageInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o S3StorageInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o S3StorageInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
type SparkNode_SdkV2 struct {
	// The private IP address of the host instance.
	HostPrivateIp types.String `tfsdk:"host_private_ip"`
	// Globally unique identifier for the host instance from the cloud provider.
	InstanceId types.String `tfsdk:"instance_id"`
	// Attributes specific to AWS for a Spark node.
	NodeAwsAttributes types.List `tfsdk:"node_aws_attributes"`
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

func (newState *SparkNode_SdkV2) SyncFieldsDuringCreateOrUpdate(plan SparkNode_SdkV2) {
}

func (newState *SparkNode_SdkV2) SyncFieldsDuringRead(existingState SparkNode_SdkV2) {
}

func (c SparkNode_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["host_private_ip"] = attrs["host_private_ip"].SetOptional()
	attrs["instance_id"] = attrs["instance_id"].SetOptional()
	attrs["node_aws_attributes"] = attrs["node_aws_attributes"].SetOptional()
	attrs["node_aws_attributes"] = attrs["node_aws_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a SparkNode_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"node_aws_attributes": reflect.TypeOf(SparkNodeAwsAttributes_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkNode_SdkV2
// only implements ToObjectValue() and Type().
func (o SparkNode_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SparkNode_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"host_private_ip": types.StringType,
			"instance_id":     types.StringType,
			"node_aws_attributes": basetypes.ListType{
				ElemType: SparkNodeAwsAttributes_SdkV2{}.Type(ctx),
			},
			"node_id":         types.StringType,
			"private_ip":      types.StringType,
			"public_dns":      types.StringType,
			"start_timestamp": types.Int64Type,
		},
	}
}

// GetNodeAwsAttributes returns the value of the NodeAwsAttributes field in SparkNode_SdkV2 as
// a SparkNodeAwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SparkNode_SdkV2) GetNodeAwsAttributes(ctx context.Context) (SparkNodeAwsAttributes_SdkV2, bool) {
	var e SparkNodeAwsAttributes_SdkV2
	if o.NodeAwsAttributes.IsNull() || o.NodeAwsAttributes.IsUnknown() {
		return e, false
	}
	var v []SparkNodeAwsAttributes_SdkV2
	d := o.NodeAwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNodeAwsAttributes sets the value of the NodeAwsAttributes field in SparkNode_SdkV2.
func (o *SparkNode_SdkV2) SetNodeAwsAttributes(ctx context.Context, v SparkNodeAwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["node_aws_attributes"]
	o.NodeAwsAttributes = types.ListValueMust(t, vs)
}

// Attributes specific to AWS for a Spark node.
type SparkNodeAwsAttributes_SdkV2 struct {
	// Whether this node is on an Amazon spot instance.
	IsSpot types.Bool `tfsdk:"is_spot"`
}

func (newState *SparkNodeAwsAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(plan SparkNodeAwsAttributes_SdkV2) {
}

func (newState *SparkNodeAwsAttributes_SdkV2) SyncFieldsDuringRead(existingState SparkNodeAwsAttributes_SdkV2) {
}

func (c SparkNodeAwsAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SparkNodeAwsAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkNodeAwsAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (o SparkNodeAwsAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_spot": o.IsSpot,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparkNodeAwsAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_spot": types.BoolType,
		},
	}
}

type SparkVersion_SdkV2 struct {
	// Spark version key, for example "2.1.x-scala2.11". This is the value which
	// should be provided as the "spark_version" when creating a new cluster.
	// Note that the exact Spark version may change over time for a "wildcard"
	// version (i.e., "2.1.x-scala2.11" is a "wildcard" version) with minor bug
	// fixes.
	Key types.String `tfsdk:"key"`
	// A descriptive name for this Spark version, for example "Spark 2.1".
	Name types.String `tfsdk:"name"`
}

func (newState *SparkVersion_SdkV2) SyncFieldsDuringCreateOrUpdate(plan SparkVersion_SdkV2) {
}

func (newState *SparkVersion_SdkV2) SyncFieldsDuringRead(existingState SparkVersion_SdkV2) {
}

func (c SparkVersion_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SparkVersion_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkVersion_SdkV2
// only implements ToObjectValue() and Type().
func (o SparkVersion_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":  o.Key,
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparkVersion_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":  types.StringType,
			"name": types.StringType,
		},
	}
}

type SparkVersionsRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparkVersionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkVersionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SparkVersionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SparkVersionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type StartCluster_SdkV2 struct {
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
func (a StartCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o StartCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type StartClusterResponse_SdkV2 struct {
}

func (newState *StartClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan StartClusterResponse_SdkV2) {
}

func (newState *StartClusterResponse_SdkV2) SyncFieldsDuringRead(existingState StartClusterResponse_SdkV2) {
}

func (c StartClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o StartClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o StartClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type TerminationReason_SdkV2 struct {
	// status code indicating why the cluster was terminated
	Code types.String `tfsdk:"code"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters types.Map `tfsdk:"parameters"`
	// type of the termination
	Type_ types.String `tfsdk:"type"`
}

func (newState *TerminationReason_SdkV2) SyncFieldsDuringCreateOrUpdate(plan TerminationReason_SdkV2) {
}

func (newState *TerminationReason_SdkV2) SyncFieldsDuringRead(existingState TerminationReason_SdkV2) {
}

func (c TerminationReason_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TerminationReason_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TerminationReason_SdkV2
// only implements ToObjectValue() and Type().
func (o TerminationReason_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"code":       o.Code,
			"parameters": o.Parameters,
			"type":       o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TerminationReason_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetParameters returns the value of the Parameters field in TerminationReason_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TerminationReason_SdkV2) GetParameters(ctx context.Context) (map[string]types.String, bool) {
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

// SetParameters sets the value of the Parameters field in TerminationReason_SdkV2.
func (o *TerminationReason_SdkV2) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.MapValueMust(t, vs)
}

type UninstallLibraries_SdkV2 struct {
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
func (a UninstallLibraries_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"libraries": reflect.TypeOf(Library_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UninstallLibraries_SdkV2
// only implements ToObjectValue() and Type().
func (o UninstallLibraries_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
			"libraries":  o.Libraries,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UninstallLibraries_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"libraries": basetypes.ListType{
				ElemType: Library_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetLibraries returns the value of the Libraries field in UninstallLibraries_SdkV2 as
// a slice of Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UninstallLibraries_SdkV2) GetLibraries(ctx context.Context) ([]Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in UninstallLibraries_SdkV2.
func (o *UninstallLibraries_SdkV2) SetLibraries(ctx context.Context, v []Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

type UninstallLibrariesResponse_SdkV2 struct {
}

func (newState *UninstallLibrariesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan UninstallLibrariesResponse_SdkV2) {
}

func (newState *UninstallLibrariesResponse_SdkV2) SyncFieldsDuringRead(existingState UninstallLibrariesResponse_SdkV2) {
}

func (c UninstallLibrariesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UninstallLibrariesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UninstallLibrariesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UninstallLibrariesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UninstallLibrariesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UnpinCluster_SdkV2 struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpinCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpinCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpinCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o UnpinCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id": o.ClusterId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UnpinCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

type UnpinClusterResponse_SdkV2 struct {
}

func (newState *UnpinClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan UnpinClusterResponse_SdkV2) {
}

func (newState *UnpinClusterResponse_SdkV2) SyncFieldsDuringRead(existingState UnpinClusterResponse_SdkV2) {
}

func (c UnpinClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpinClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpinClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpinClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UnpinClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UnpinClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateCluster_SdkV2 struct {
	// The cluster to be updated.
	Cluster types.List `tfsdk:"cluster"`
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
func (a UpdateCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster": reflect.TypeOf(UpdateClusterResource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster":     o.Cluster,
			"cluster_id":  o.ClusterId,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster": basetypes.ListType{
				ElemType: UpdateClusterResource_SdkV2{}.Type(ctx),
			},
			"cluster_id":  types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetCluster returns the value of the Cluster field in UpdateCluster_SdkV2 as
// a UpdateClusterResource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCluster_SdkV2) GetCluster(ctx context.Context) (UpdateClusterResource_SdkV2, bool) {
	var e UpdateClusterResource_SdkV2
	if o.Cluster.IsNull() || o.Cluster.IsUnknown() {
		return e, false
	}
	var v []UpdateClusterResource_SdkV2
	d := o.Cluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCluster sets the value of the Cluster field in UpdateCluster_SdkV2.
func (o *UpdateCluster_SdkV2) SetCluster(ctx context.Context, v UpdateClusterResource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster"]
	o.Cluster = types.ListValueMust(t, vs)
}

type UpdateClusterResource_SdkV2 struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf"`
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
	DockerImage types.List `tfsdk:"docker_image"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
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

	WorkloadType types.List `tfsdk:"workload_type"`
}

func (newState *UpdateClusterResource_SdkV2) SyncFieldsDuringCreateOrUpdate(plan UpdateClusterResource_SdkV2) {
}

func (newState *UpdateClusterResource_SdkV2) SyncFieldsDuringRead(existingState UpdateClusterResource_SdkV2) {
}

func (c UpdateClusterResource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["autotermination_minutes"] = attrs["autotermination_minutes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_name"] = attrs["cluster_name"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["data_security_mode"] = attrs["data_security_mode"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].SetOptional()
	attrs["docker_image"] = attrs["docker_image"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_elastic_disk"] = attrs["enable_elastic_disk"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
	attrs["workload_type"] = attrs["workload_type"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateClusterResource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateClusterResource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(AutoScale_SdkV2{}),
		"aws_attributes":   reflect.TypeOf(AwsAttributes_SdkV2{}),
		"azure_attributes": reflect.TypeOf(AzureAttributes_SdkV2{}),
		"cluster_log_conf": reflect.TypeOf(ClusterLogConf_SdkV2{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"docker_image":     reflect.TypeOf(DockerImage_SdkV2{}),
		"gcp_attributes":   reflect.TypeOf(GcpAttributes_SdkV2{}),
		"init_scripts":     reflect.TypeOf(InitScriptInfo_SdkV2{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
		"workload_type":    reflect.TypeOf(WorkloadType_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateClusterResource_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateClusterResource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateClusterResource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale": basetypes.ListType{
				ElemType: AutoScale_SdkV2{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes_SdkV2{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf_SdkV2{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage_SdkV2{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes_SdkV2{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo_SdkV2{}.Type(ctx),
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in UpdateClusterResource_SdkV2 as
// a AutoScale_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetAutoscale(ctx context.Context) (AutoScale_SdkV2, bool) {
	var e AutoScale_SdkV2
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []AutoScale_SdkV2
	d := o.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetAutoscale(ctx context.Context, v AutoScale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	o.Autoscale = types.ListValueMust(t, vs)
}

// GetAwsAttributes returns the value of the AwsAttributes field in UpdateClusterResource_SdkV2 as
// a AwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetAwsAttributes(ctx context.Context) (AwsAttributes_SdkV2, bool) {
	var e AwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []AwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetAwsAttributes(ctx context.Context, v AwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in UpdateClusterResource_SdkV2 as
// a AzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetAzureAttributes(ctx context.Context) (AzureAttributes_SdkV2, bool) {
	var e AzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []AzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetAzureAttributes(ctx context.Context, v AzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetClusterLogConf returns the value of the ClusterLogConf field in UpdateClusterResource_SdkV2 as
// a ClusterLogConf_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetClusterLogConf(ctx context.Context) (ClusterLogConf_SdkV2, bool) {
	var e ClusterLogConf_SdkV2
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []ClusterLogConf_SdkV2
	d := o.ClusterLogConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetClusterLogConf(ctx context.Context, v ClusterLogConf_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_conf"]
	o.ClusterLogConf = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in UpdateClusterResource_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetDockerImage returns the value of the DockerImage field in UpdateClusterResource_SdkV2 as
// a DockerImage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetDockerImage(ctx context.Context) (DockerImage_SdkV2, bool) {
	var e DockerImage_SdkV2
	if o.DockerImage.IsNull() || o.DockerImage.IsUnknown() {
		return e, false
	}
	var v []DockerImage_SdkV2
	d := o.DockerImage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDockerImage sets the value of the DockerImage field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetDockerImage(ctx context.Context, v DockerImage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["docker_image"]
	o.DockerImage = types.ListValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in UpdateClusterResource_SdkV2 as
// a GcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetGcpAttributes(ctx context.Context) (GcpAttributes_SdkV2, bool) {
	var e GcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []GcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetGcpAttributes(ctx context.Context, v GcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in UpdateClusterResource_SdkV2 as
// a slice of InitScriptInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetInitScripts(ctx context.Context) ([]InitScriptInfo_SdkV2, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []InitScriptInfo_SdkV2
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetInitScripts(ctx context.Context, v []InitScriptInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in UpdateClusterResource_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkConf sets the value of the SparkConf field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in UpdateClusterResource_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkEnvVars sets the value of the SparkEnvVars field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in UpdateClusterResource_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
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

// SetSshPublicKeys sets the value of the SshPublicKeys field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

// GetWorkloadType returns the value of the WorkloadType field in UpdateClusterResource_SdkV2 as
// a WorkloadType_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateClusterResource_SdkV2) GetWorkloadType(ctx context.Context) (WorkloadType_SdkV2, bool) {
	var e WorkloadType_SdkV2
	if o.WorkloadType.IsNull() || o.WorkloadType.IsUnknown() {
		return e, false
	}
	var v []WorkloadType_SdkV2
	d := o.WorkloadType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkloadType sets the value of the WorkloadType field in UpdateClusterResource_SdkV2.
func (o *UpdateClusterResource_SdkV2) SetWorkloadType(ctx context.Context, v WorkloadType_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workload_type"]
	o.WorkloadType = types.ListValueMust(t, vs)
}

type UpdateClusterResponse_SdkV2 struct {
}

func (newState *UpdateClusterResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan UpdateClusterResponse_SdkV2) {
}

func (newState *UpdateClusterResponse_SdkV2) SyncFieldsDuringRead(existingState UpdateClusterResponse_SdkV2) {
}

func (c UpdateClusterResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateClusterResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateClusterResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateClusterResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateClusterResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateClusterResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateDefaultBaseEnvironmentRequest_SdkV2 struct {
	DefaultBaseEnvironment types.List `tfsdk:"default_base_environment"`

	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDefaultBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_base_environment": reflect.TypeOf(DefaultBaseEnvironment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateDefaultBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_base_environment": o.DefaultBaseEnvironment,
			"id":                       o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDefaultBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_base_environment": basetypes.ListType{
				ElemType: DefaultBaseEnvironment_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
		},
	}
}

// GetDefaultBaseEnvironment returns the value of the DefaultBaseEnvironment field in UpdateDefaultBaseEnvironmentRequest_SdkV2 as
// a DefaultBaseEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateDefaultBaseEnvironmentRequest_SdkV2) GetDefaultBaseEnvironment(ctx context.Context) (DefaultBaseEnvironment_SdkV2, bool) {
	var e DefaultBaseEnvironment_SdkV2
	if o.DefaultBaseEnvironment.IsNull() || o.DefaultBaseEnvironment.IsUnknown() {
		return e, false
	}
	var v []DefaultBaseEnvironment_SdkV2
	d := o.DefaultBaseEnvironment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultBaseEnvironment sets the value of the DefaultBaseEnvironment field in UpdateDefaultBaseEnvironmentRequest_SdkV2.
func (o *UpdateDefaultBaseEnvironmentRequest_SdkV2) SetDefaultBaseEnvironment(ctx context.Context, v DefaultBaseEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_base_environment"]
	o.DefaultBaseEnvironment = types.ListValueMust(t, vs)
}

type UpdateDefaultDefaultBaseEnvironmentRequest_SdkV2 struct {
	Id types.String `tfsdk:"id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultDefaultBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDefaultDefaultBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultDefaultBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateDefaultDefaultBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDefaultDefaultBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type UpdateResponse_SdkV2 struct {
}

func (newState *UpdateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan UpdateResponse_SdkV2) {
}

func (newState *UpdateResponse_SdkV2) SyncFieldsDuringRead(existingState UpdateResponse_SdkV2) {
}

func (c UpdateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A storage location back by UC Volumes.
type VolumesStorageInfo_SdkV2 struct {
	// UC Volumes destination, e.g.
	// `/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh` or
	// `dbfs:/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh`
	Destination types.String `tfsdk:"destination"`
}

func (newState *VolumesStorageInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan VolumesStorageInfo_SdkV2) {
}

func (newState *VolumesStorageInfo_SdkV2) SyncFieldsDuringRead(existingState VolumesStorageInfo_SdkV2) {
}

func (c VolumesStorageInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a VolumesStorageInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VolumesStorageInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o VolumesStorageInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VolumesStorageInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

// Cluster Attributes showing for clusters workload types.
type WorkloadType_SdkV2 struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	Clients types.List `tfsdk:"clients"`
}

func (newState *WorkloadType_SdkV2) SyncFieldsDuringCreateOrUpdate(plan WorkloadType_SdkV2) {
}

func (newState *WorkloadType_SdkV2) SyncFieldsDuringRead(existingState WorkloadType_SdkV2) {
}

func (c WorkloadType_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clients"] = attrs["clients"].SetRequired()
	attrs["clients"] = attrs["clients"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkloadType.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkloadType_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clients": reflect.TypeOf(ClientsTypes_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkloadType_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkloadType_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clients": o.Clients,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkloadType_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clients": basetypes.ListType{
				ElemType: ClientsTypes_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetClients returns the value of the Clients field in WorkloadType_SdkV2 as
// a ClientsTypes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkloadType_SdkV2) GetClients(ctx context.Context) (ClientsTypes_SdkV2, bool) {
	var e ClientsTypes_SdkV2
	if o.Clients.IsNull() || o.Clients.IsUnknown() {
		return e, false
	}
	var v []ClientsTypes_SdkV2
	d := o.Clients.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClients sets the value of the Clients field in WorkloadType_SdkV2.
func (o *WorkloadType_SdkV2) SetClients(ctx context.Context, v ClientsTypes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clients"]
	o.Clients = types.ListValueMust(t, vs)
}

// A storage location in Workspace Filesystem (WSFS)
type WorkspaceStorageInfo_SdkV2 struct {
	// wsfs destination, e.g. `workspace:/cluster-init-scripts/setup-datadog.sh`
	Destination types.String `tfsdk:"destination"`
}

func (newState *WorkspaceStorageInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(plan WorkspaceStorageInfo_SdkV2) {
}

func (newState *WorkspaceStorageInfo_SdkV2) SyncFieldsDuringRead(existingState WorkspaceStorageInfo_SdkV2) {
}

func (c WorkspaceStorageInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceStorageInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceStorageInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceStorageInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}
