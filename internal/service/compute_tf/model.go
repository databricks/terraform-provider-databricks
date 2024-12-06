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
	"fmt"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
	IamRoleArn types.String `tfsdk:"iam_role_arn" tf:"optional"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:""`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile types.Bool `tfsdk:"is_meta_instance_profile" tf:"optional"`
	// By default, Databricks validates that it has sufficient permissions to
	// launch instances with the instance profile. This validation uses AWS
	// dry-run mode for the RunInstances API. If validation fails with an error
	// message that does not indicate an IAM related permission issue, (e.g.
	// “Your requested instance type is not supported in your requested
	// availability zone”), you can pass this flag to skip the validation and
	// forcibly add the instance profile.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

func (newState *AddInstanceProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddInstanceProfile) {
}

func (newState *AddInstanceProfile) SyncEffectiveFieldsDuringRead(existingState AddInstanceProfile) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = AddInstanceProfile{}

// Equal implements basetypes.ObjectValuable.
func (o AddInstanceProfile) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o AddInstanceProfile) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o AddInstanceProfile) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o AddInstanceProfile) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o AddInstanceProfile) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o AddInstanceProfile) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = AddResponse{}

// Equal implements basetypes.ObjectValuable.
func (o AddResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o AddResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o AddResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o AddResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o AddResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o AddResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o AddResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Adlsgen2Info struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	Destination types.String `tfsdk:"destination" tf:""`
}

func (newState *Adlsgen2Info) SyncEffectiveFieldsDuringCreateOrUpdate(plan Adlsgen2Info) {
}

func (newState *Adlsgen2Info) SyncEffectiveFieldsDuringRead(existingState Adlsgen2Info) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Adlsgen2Info{}

// Equal implements basetypes.ObjectValuable.
func (o Adlsgen2Info) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Adlsgen2Info) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Adlsgen2Info) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Adlsgen2Info) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Adlsgen2Info) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Adlsgen2Info) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	MaxWorkers types.Int64 `tfsdk:"max_workers" tf:"optional"`
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers types.Int64 `tfsdk:"min_workers" tf:"optional"`
}

func (newState *AutoScale) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoScale) {
}

func (newState *AutoScale) SyncEffectiveFieldsDuringRead(existingState AutoScale) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = AutoScale{}

// Equal implements basetypes.ObjectValuable.
func (o AutoScale) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o AutoScale) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o AutoScale) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o AutoScale) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o AutoScale) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o AutoScale) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type AwsAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones.
	//
	// Note: If `first_on_demand` is zero, this availability type will be used
	// for the entire cluster.
	Availability types.String `tfsdk:"availability" tf:"optional"`
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
	EbsVolumeCount types.Int64 `tfsdk:"ebs_volume_count" tf:"optional"`
	// If using gp3 volumes, what IOPS to use for the disk. If this is not set,
	// the maximum performance of a gp2 volume with the same volume size will be
	// used.
	EbsVolumeIops types.Int64 `tfsdk:"ebs_volume_iops" tf:"optional"`
	// The size of each EBS volume (in GiB) launched for each instance. For
	// general purpose SSD, this value must be within the range 100 - 4096. For
	// throughput optimized HDD, this value must be within the range 500 - 4096.
	EbsVolumeSize types.Int64 `tfsdk:"ebs_volume_size" tf:"optional"`
	// If using gp3 volumes, what throughput to use for the disk. If this is not
	// set, the maximum performance of a gp2 volume with the same volume size
	// will be used.
	EbsVolumeThroughput types.Int64 `tfsdk:"ebs_volume_throughput" tf:"optional"`
	// The type of EBS volumes that will be launched with this cluster.
	EbsVolumeType types.String `tfsdk:"ebs_volume_type" tf:"optional"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. If this value is greater than 0, the cluster driver
	// node in particular will be placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand types.Int64 `tfsdk:"first_on_demand" tf:"optional"`
	// Nodes for this cluster will only be placed on AWS instances with this
	// instance profile. If ommitted, nodes will be placed on instances without
	// an IAM instance profile. The instance profile must have previously been
	// added to the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	//
	// If this field is ommitted, we will pull in the default from the conf if
	// it exists.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
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
	//
	// The default value and documentation here should be kept consistent with
	// CommonConf.defaultSpotBidPricePercent and
	// CommonConf.maxSpotBidPricePercent.
	SpotBidPricePercent types.Int64 `tfsdk:"spot_bid_price_percent" tf:"optional"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. If the zone specified is "auto", will try to place cluster
	// in a zone with high availability, and will retry placement in a different
	// AZ if there is not enough capacity. The list of available zones as well
	// as the default value can be found by using the `List Zones` method.
	ZoneId types.String `tfsdk:"zone_id" tf:"optional"`
}

func (newState *AwsAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsAttributes) {
}

func (newState *AwsAttributes) SyncEffectiveFieldsDuringRead(existingState AwsAttributes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = AwsAttributes{}

// Equal implements basetypes.ObjectValuable.
func (o AwsAttributes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o AwsAttributes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o AwsAttributes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o AwsAttributes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o AwsAttributes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o AwsAttributes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type AzureAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones. Note: If `first_on_demand` is zero (which only
	// happens on pool clusters), this availability type will be used for the
	// entire cluster.
	Availability types.String `tfsdk:"availability" tf:"optional"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. This value should be greater than 0, to make sure
	// the cluster driver node is placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand types.Int64 `tfsdk:"first_on_demand" tf:"optional"`
	// Defines values necessary to configure and run Azure Log Analytics agent
	LogAnalyticsInfo types.List `tfsdk:"log_analytics_info" tf:"optional,object"`
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	SpotBidMaxPrice types.Float64 `tfsdk:"spot_bid_max_price" tf:"optional"`
}

func (newState *AzureAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureAttributes) {
}

func (newState *AzureAttributes) SyncEffectiveFieldsDuringRead(existingState AzureAttributes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = AzureAttributes{}

// Equal implements basetypes.ObjectValuable.
func (o AzureAttributes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o AzureAttributes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o AzureAttributes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o AzureAttributes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o AzureAttributes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o AzureAttributes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o AzureAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability":    types.StringType,
			"first_on_demand": types.Int64Type,
			"log_analytics_info": basetypes.ListType{
				ElemType: LogAnalyticsInfo{}.Type(ctx),
			},
			"spot_bid_max_price": types.Float64Type,
		},
	}
}

type CancelCommand struct {
	ClusterId types.String `tfsdk:"clusterId" tf:"optional"`

	CommandId types.String `tfsdk:"commandId" tf:"optional"`

	ContextId types.String `tfsdk:"contextId" tf:"optional"`
}

func (newState *CancelCommand) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelCommand) {
}

func (newState *CancelCommand) SyncEffectiveFieldsDuringRead(existingState CancelCommand) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CancelCommand{}

// Equal implements basetypes.ObjectValuable.
func (o CancelCommand) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CancelCommand) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CancelCommand) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CancelCommand) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CancelCommand) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CancelCommand) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CancelResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CancelResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CancelResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CancelResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CancelResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CancelResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CancelResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CancelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ChangeClusterOwner struct {
	// <needs content added>
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// New owner of the cluster_id after this RPC.
	OwnerUsername types.String `tfsdk:"owner_username" tf:""`
}

func (newState *ChangeClusterOwner) SyncEffectiveFieldsDuringCreateOrUpdate(plan ChangeClusterOwner) {
}

func (newState *ChangeClusterOwner) SyncEffectiveFieldsDuringRead(existingState ChangeClusterOwner) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ChangeClusterOwner{}

// Equal implements basetypes.ObjectValuable.
func (o ChangeClusterOwner) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ChangeClusterOwner) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ChangeClusterOwner) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ChangeClusterOwner) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ChangeClusterOwner) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ChangeClusterOwner) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ChangeClusterOwnerResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ChangeClusterOwnerResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ClientsTypes struct {
	// With jobs set, the cluster can be used for jobs
	Jobs types.Bool `tfsdk:"jobs" tf:"optional"`
	// With notebooks set, this cluster can be used for notebooks
	Notebooks types.Bool `tfsdk:"notebooks" tf:"optional"`
}

func (newState *ClientsTypes) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClientsTypes) {
}

func (newState *ClientsTypes) SyncEffectiveFieldsDuringRead(existingState ClientsTypes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClientsTypes{}

// Equal implements basetypes.ObjectValuable.
func (o ClientsTypes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClientsTypes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClientsTypes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClientsTypes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClientsTypes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClientsTypes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	SourceClusterId types.String `tfsdk:"source_cluster_id" tf:""`
}

func (newState *CloneCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloneCluster) {
}

func (newState *CloneCluster) SyncEffectiveFieldsDuringRead(existingState CloneCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CloneCluster{}

// Equal implements basetypes.ObjectValuable.
func (o CloneCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CloneCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CloneCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CloneCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CloneCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CloneCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Status types.List `tfsdk:"status" tf:"optional"`
}

func (newState *CloudProviderNodeInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudProviderNodeInfo) {
}

func (newState *CloudProviderNodeInfo) SyncEffectiveFieldsDuringRead(existingState CloudProviderNodeInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CloudProviderNodeInfo{}

// Equal implements basetypes.ObjectValuable.
func (o CloudProviderNodeInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CloudProviderNodeInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CloudProviderNodeInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CloudProviderNodeInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CloudProviderNodeInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CloudProviderNodeInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ClusterAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAccessControlRequest) {
}

func (newState *ClusterAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState ClusterAccessControlRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterAccessControlRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterAccessControlRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterAccessControlRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterAccessControlRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterAccessControlRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterAccessControlRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterAccessControlRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ClusterAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAccessControlResponse) {
}

func (newState *ClusterAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState ClusterAccessControlResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterAccessControlResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterAccessControlResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterAccessControlResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterAccessControlResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterAccessControlResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterAccessControlResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterAccessControlResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterAttributes struct {
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf" tf:"optional,object"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode types.String `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage types.List `tfsdk:"docker_image" tf:"optional,object"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf" tf:"optional"`
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
	SparkEnvVars types.Map `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:""`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType types.List `tfsdk:"workload_type" tf:"optional,object"`
}

func (newState *ClusterAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAttributes) {
}

func (newState *ClusterAttributes) SyncEffectiveFieldsDuringRead(existingState ClusterAttributes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterAttributes{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterAttributes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterAttributes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterAttributes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterAttributes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterAttributes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterAttributes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id": types.StringType,
			"node_type_id":     types.StringType,
			"policy_id":        types.StringType,
			"runtime_engine":   types.StringType,
			"single_user_name": types.StringType,
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType{}.Type(ctx),
			},
		},
	}
}

type ClusterCompliance struct {
	// Canonical unique identifier for a cluster.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// Whether this cluster is in compliance with the latest version of its
	// policy.
	IsCompliant types.Bool `tfsdk:"is_compliant" tf:"optional"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	Violations types.Map `tfsdk:"violations" tf:"optional"`
}

func (newState *ClusterCompliance) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterCompliance) {
}

func (newState *ClusterCompliance) SyncEffectiveFieldsDuringRead(existingState ClusterCompliance) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterCompliance{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterCompliance) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterCompliance) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterCompliance) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterCompliance) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterCompliance) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterCompliance) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterDetails struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// Number of CPU cores available for this cluster. Note that this can be
	// fractional, e.g. 7.5 cores, since certain node types are configured to
	// share cores between Spark nodes on the same instance.
	ClusterCores types.Float64 `tfsdk:"cluster_cores" tf:"optional"`
	// Canonical identifier for the cluster. This id is retained during cluster
	// restarts and resizes, while each new cluster has a globally unique id.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf" tf:"optional,object"`
	// Cluster log delivery status.
	ClusterLogStatus types.List `tfsdk:"cluster_log_status" tf:"optional,object"`
	// Total amount of cluster memory, in megabytes
	ClusterMemoryMb types.Int64 `tfsdk:"cluster_memory_mb" tf:"optional"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Determines whether the cluster was created by a user through the UI,
	// created by the Databricks Jobs Scheduler, or through an API request. This
	// is the same as cluster_creator, but read only.
	ClusterSource types.String `tfsdk:"cluster_source" tf:"optional"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode types.String `tfsdk:"data_security_mode" tf:"optional"`
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
	DefaultTags types.Map `tfsdk:"default_tags" tf:"optional"`

	DockerImage types.List `tfsdk:"docker_image" tf:"optional,object"`
	// Node on which the Spark driver resides. The driver node contains the
	// Spark master and the Databricks application that manages the per-notebook
	// Spark REPLs.
	Driver types.List `tfsdk:"driver" tf:"optional,object"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Nodes on which the Spark executors reside.
	Executors types.List `tfsdk:"executors" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// Port on which Spark JDBC server is listening, in the driver nod. No
	// service will be listeningon on this port in executor nodes.
	JdbcPort types.Int64 `tfsdk:"jdbc_port" tf:"optional"`
	// the timestamp that the cluster was started/restarted
	LastRestartedTime types.Int64 `tfsdk:"last_restarted_time" tf:"optional"`
	// Time when the cluster driver last lost its state (due to a restart or
	// driver failure).
	LastStateLossTime types.Int64 `tfsdk:"last_state_loss_time" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
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
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf" tf:"optional"`
	// A canonical SparkContext identifier. This value *does* change when the
	// Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
	// globally unique identifier over all Spark contexts.
	SparkContextId types.Int64 `tfsdk:"spark_context_id" tf:"optional"`
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
	SparkEnvVars types.Map `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:"optional"`
	// `spec` contains a snapshot of the field values that were used to create
	// or edit this cluster. The contents of `spec` can be used in the body of a
	// create cluster request. This field might not be populated for older
	// clusters. Note: not included in the response of the ListClusters API.
	Spec types.List `tfsdk:"spec" tf:"optional,object"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys" tf:"optional"`
	// Time (in epoch milliseconds) when the cluster creation request was
	// received (when the cluster entered a `PENDING` state).
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Current state of the cluster.
	State types.String `tfsdk:"state" tf:"optional"`
	// A message associated with the most recent state transition (e.g., the
	// reason why the cluster entered a `TERMINATED` state).
	StateMessage types.String `tfsdk:"state_message" tf:"optional"`
	// Time (in epoch milliseconds) when the cluster was terminated, if
	// applicable.
	TerminatedTime types.Int64 `tfsdk:"terminated_time" tf:"optional"`
	// Information about why the cluster was terminated. This field only appears
	// when the cluster is in a `TERMINATING` or `TERMINATED` state.
	TerminationReason types.List `tfsdk:"termination_reason" tf:"optional,object"`

	WorkloadType types.List `tfsdk:"workload_type" tf:"optional,object"`
}

func (newState *ClusterDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterDetails) {
}

func (newState *ClusterDetails) SyncEffectiveFieldsDuringRead(existingState ClusterDetails) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterDetails{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterDetails) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterDetails) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterDetails) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterDetails) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterDetails) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterDetails) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ClusterDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale": basetypes.ListType{
				ElemType: AutoScale{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes{}.Type(ctx),
			},
			"cluster_cores": types.Float64Type,
			"cluster_id":    types.StringType,
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf{}.Type(ctx),
			},
			"cluster_log_status": basetypes.ListType{
				ElemType: LogSyncStatus{}.Type(ctx),
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
				ElemType: DockerImage{}.Type(ctx),
			},
			"driver": basetypes.ListType{
				ElemType: SparkNode{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"executors": basetypes.ListType{
				ElemType: SparkNode{}.Type(ctx),
			},
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id":     types.StringType,
			"jdbc_port":            types.Int64Type,
			"last_restarted_time":  types.Int64Type,
			"last_state_loss_time": types.Int64Type,
			"node_type_id":         types.StringType,
			"num_workers":          types.Int64Type,
			"policy_id":            types.StringType,
			"runtime_engine":       types.StringType,
			"single_user_name":     types.StringType,
			"spark_conf": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_context_id": types.Int64Type,
			"spark_env_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_version": types.StringType,
			"spec": basetypes.ListType{
				ElemType: ClusterSpec{}.Type(ctx),
			},
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"start_time":      types.Int64Type,
			"state":           types.StringType,
			"state_message":   types.StringType,
			"terminated_time": types.Int64Type,
			"termination_reason": basetypes.ListType{
				ElemType: TerminationReason{}.Type(ctx),
			},
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType{}.Type(ctx),
			},
		},
	}
}

type ClusterEvent struct {
	// <needs content added>
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// <needs content added>
	DataPlaneEventDetails types.List `tfsdk:"data_plane_event_details" tf:"optional,object"`
	// <needs content added>
	Details types.List `tfsdk:"details" tf:"optional,object"`
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	Timestamp types.Int64 `tfsdk:"timestamp" tf:"optional"`

	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *ClusterEvent) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterEvent) {
}

func (newState *ClusterEvent) SyncEffectiveFieldsDuringRead(existingState ClusterEvent) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterEvent{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterEvent) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterEvent) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterEvent) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterEvent) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterEvent) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterEvent) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ClusterEvent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
			"data_plane_event_details": basetypes.ListType{
				ElemType: DataPlaneEventDetails{}.Type(ctx),
			},
			"details": basetypes.ListType{
				ElemType: EventDetails{}.Type(ctx),
			},
			"timestamp": types.Int64Type,
			"type":      types.StringType,
		},
	}
}

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// Status of all libraries on the cluster.
	LibraryStatuses types.List `tfsdk:"library_statuses" tf:"optional"`
}

func (newState *ClusterLibraryStatuses) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterLibraryStatuses) {
}

func (newState *ClusterLibraryStatuses) SyncEffectiveFieldsDuringRead(existingState ClusterLibraryStatuses) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterLibraryStatuses{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterLibraryStatuses) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterLibraryStatuses) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterLibraryStatuses) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterLibraryStatuses) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterLibraryStatuses) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterLibraryStatuses) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterLogConf struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs types.List `tfsdk:"dbfs" tf:"optional,object"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 types.List `tfsdk:"s3" tf:"optional,object"`
}

func (newState *ClusterLogConf) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterLogConf) {
}

func (newState *ClusterLogConf) SyncEffectiveFieldsDuringRead(existingState ClusterLogConf) {
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
		"dbfs": reflect.TypeOf(DbfsStorageInfo{}),
		"s3":   reflect.TypeOf(S3StorageInfo{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterLogConf{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterLogConf) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterLogConf) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterLogConf) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterLogConf) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterLogConf) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterLogConf) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ClusterLogConf) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbfs": basetypes.ListType{
				ElemType: DbfsStorageInfo{}.Type(ctx),
			},
			"s3": basetypes.ListType{
				ElemType: S3StorageInfo{}.Type(ctx),
			},
		},
	}
}

type ClusterPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ClusterPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPermission) {
}

func (newState *ClusterPermission) SyncEffectiveFieldsDuringRead(existingState ClusterPermission) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPermission{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPermission) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPermission) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPermission) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPermission) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPermission) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPermission) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ClusterPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPermissions) {
}

func (newState *ClusterPermissions) SyncEffectiveFieldsDuringRead(existingState ClusterPermissions) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPermissions{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ClusterPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPermissionsDescription) {
}

func (newState *ClusterPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState ClusterPermissionsDescription) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPermissionsDescription{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPermissionsDescription) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPermissionsDescription) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPermissionsDescription) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPermissionsDescription) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPermissionsDescription) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPermissionsDescription) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-"`
}

func (newState *ClusterPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPermissionsRequest) {
}

func (newState *ClusterPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState ClusterPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterPolicyAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ClusterPolicyAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyAccessControlRequest) {
}

func (newState *ClusterPolicyAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyAccessControlRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPolicyAccessControlRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ClusterPolicyAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyAccessControlResponse) {
}

func (newState *ClusterPolicyAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyAccessControlResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPolicyAccessControlResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPolicyAccessControlResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterPolicyPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ClusterPolicyPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyPermission) {
}

func (newState *ClusterPolicyPermission) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyPermission) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPolicyPermission{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPolicyPermission) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPolicyPermission) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPolicyPermission) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPolicyPermission) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPolicyPermission) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPolicyPermission) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterPolicyPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ClusterPolicyPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyPermissions) {
}

func (newState *ClusterPolicyPermissions) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyPermissions) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPolicyPermissions{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ClusterPolicyPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ClusterPolicyPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyPermissionsDescription) {
}

func (newState *ClusterPolicyPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyPermissionsDescription) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPolicyPermissionsDescription{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsDescription) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsDescription) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsDescription) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsDescription) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsDescription) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsDescription) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-"`
}

func (newState *ClusterPolicyPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterPolicyPermissionsRequest) {
}

func (newState *ClusterPolicyPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState ClusterPolicyPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterPolicyPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterPolicyPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Represents a change to the cluster settings required for the cluster to
// become compliant with its policy.
type ClusterSettingsChange struct {
	// The field where this change would be made.
	Field types.String `tfsdk:"field" tf:"optional"`
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	NewValue types.String `tfsdk:"new_value" tf:"optional"`
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	PreviousValue types.String `tfsdk:"previous_value" tf:"optional"`
}

func (newState *ClusterSettingsChange) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterSettingsChange) {
}

func (newState *ClusterSettingsChange) SyncEffectiveFieldsDuringRead(existingState ClusterSettingsChange) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterSettingsChange{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterSettingsChange) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterSettingsChange) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterSettingsChange) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterSettingsChange) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterSettingsChange) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterSettingsChange) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
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
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
}

func (newState *ClusterSize) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterSize) {
}

func (newState *ClusterSize) SyncEffectiveFieldsDuringRead(existingState ClusterSize) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterSize{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterSize) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterSize) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterSize) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterSize) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterSize) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterSize) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ClusterSize) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale": basetypes.ListType{
				ElemType: AutoScale{}.Type(ctx),
			},
			"num_workers": types.Int64Type,
		},
	}
}

type ClusterSpec struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf" tf:"optional,object"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode types.String `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage types.List `tfsdk:"docker_image" tf:"optional,object"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
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
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf" tf:"optional"`
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
	SparkEnvVars types.Map `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:"optional"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType types.List `tfsdk:"workload_type" tf:"optional,object"`
}

func (newState *ClusterSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterSpec) {
}

func (newState *ClusterSpec) SyncEffectiveFieldsDuringRead(existingState ClusterSpec) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterSpec{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterSpec) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterSpec) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterSpec) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterSpec) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterSpec) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterSpec) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ClusterSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale": basetypes.ListType{
				ElemType: AutoScale{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id": types.StringType,
			"node_type_id":     types.StringType,
			"num_workers":      types.Int64Type,
			"policy_id":        types.StringType,
			"runtime_engine":   types.StringType,
			"single_user_name": types.StringType,
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType{}.Type(ctx),
			},
		},
	}
}

// Get status
type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId types.String `tfsdk:"-"`
}

func (newState *ClusterStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterStatus) {
}

func (newState *ClusterStatus) SyncEffectiveFieldsDuringRead(existingState ClusterStatus) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ClusterStatus{}

// Equal implements basetypes.ObjectValuable.
func (o ClusterStatus) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ClusterStatus) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ClusterStatus) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ClusterStatus) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ClusterStatus) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ClusterStatus) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ClusterId types.String `tfsdk:"clusterId" tf:"optional"`
	// Executable code
	Command types.String `tfsdk:"command" tf:"optional"`
	// Running context id
	ContextId types.String `tfsdk:"contextId" tf:"optional"`

	Language types.String `tfsdk:"language" tf:"optional"`
}

func (newState *Command) SyncEffectiveFieldsDuringCreateOrUpdate(plan Command) {
}

func (newState *Command) SyncEffectiveFieldsDuringRead(existingState Command) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Command{}

// Equal implements basetypes.ObjectValuable.
func (o Command) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Command) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Command) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Command) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Command) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Command) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get command info
type CommandStatusRequest struct {
	ClusterId types.String `tfsdk:"-"`

	CommandId types.String `tfsdk:"-"`

	ContextId types.String `tfsdk:"-"`
}

func (newState *CommandStatusRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CommandStatusRequest) {
}

func (newState *CommandStatusRequest) SyncEffectiveFieldsDuringRead(existingState CommandStatusRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CommandStatusRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CommandStatusRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CommandStatusRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CommandStatusRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CommandStatusRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CommandStatusRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CommandStatusRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Id types.String `tfsdk:"id" tf:"optional"`

	Results types.List `tfsdk:"results" tf:"optional,object"`

	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *CommandStatusResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CommandStatusResponse) {
}

func (newState *CommandStatusResponse) SyncEffectiveFieldsDuringRead(existingState CommandStatusResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CommandStatusResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CommandStatusResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CommandStatusResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CommandStatusResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CommandStatusResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CommandStatusResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CommandStatusResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CommandStatusResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"results": basetypes.ListType{
				ElemType: Results{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// Get status
type ContextStatusRequest struct {
	ClusterId types.String `tfsdk:"-"`

	ContextId types.String `tfsdk:"-"`
}

func (newState *ContextStatusRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContextStatusRequest) {
}

func (newState *ContextStatusRequest) SyncEffectiveFieldsDuringRead(existingState ContextStatusRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ContextStatusRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ContextStatusRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ContextStatusRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ContextStatusRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ContextStatusRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ContextStatusRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ContextStatusRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Id types.String `tfsdk:"id" tf:"optional"`

	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *ContextStatusResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContextStatusResponse) {
}

func (newState *ContextStatusResponse) SyncEffectiveFieldsDuringRead(existingState ContextStatusResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ContextStatusResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ContextStatusResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ContextStatusResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ContextStatusResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ContextStatusResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ContextStatusResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ContextStatusResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	CloneFrom types.List `tfsdk:"clone_from" tf:"optional,object"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf" tf:"optional,object"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode types.String `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage types.List `tfsdk:"docker_image" tf:"optional,object"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
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
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf" tf:"optional"`
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
	SparkEnvVars types.Map `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:""`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType types.List `tfsdk:"workload_type" tf:"optional,object"`
}

func (newState *CreateCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCluster) {
}

func (newState *CreateCluster) SyncEffectiveFieldsDuringRead(existingState CreateCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateCluster{}

// Equal implements basetypes.ObjectValuable.
func (o CreateCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale": basetypes.ListType{
				ElemType: AutoScale{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes{}.Type(ctx),
			},
			"clone_from": basetypes.ListType{
				ElemType: CloneCluster{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id": types.StringType,
			"node_type_id":     types.StringType,
			"num_workers":      types.Int64Type,
			"policy_id":        types.StringType,
			"runtime_engine":   types.StringType,
			"single_user_name": types.StringType,
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType{}.Type(ctx),
			},
		},
	}
}

type CreateClusterResponse struct {
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
}

func (newState *CreateClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateClusterResponse) {
}

func (newState *CreateClusterResponse) SyncEffectiveFieldsDuringRead(existingState CreateClusterResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ClusterId types.String `tfsdk:"clusterId" tf:"optional"`

	Language types.String `tfsdk:"language" tf:"optional"`
}

func (newState *CreateContext) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateContext) {
}

func (newState *CreateContext) SyncEffectiveFieldsDuringRead(existingState CreateContext) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateContext{}

// Equal implements basetypes.ObjectValuable.
func (o CreateContext) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateContext) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateContext) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateContext) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateContext) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateContext) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec types.List `tfsdk:"disk_spec" tf:"optional,object"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes" tf:"optional"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name" tf:""`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity" tf:"optional"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:""`
	// Custom Docker Image BYOC
	PreloadedDockerImages types.List `tfsdk:"preloaded_docker_images" tf:"optional"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions types.List `tfsdk:"preloaded_spark_versions" tf:"optional"`
}

func (newState *CreateInstancePool) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateInstancePool) {
}

func (newState *CreateInstancePool) SyncEffectiveFieldsDuringRead(existingState CreateInstancePool) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateInstancePool{}

// Equal implements basetypes.ObjectValuable.
func (o CreateInstancePool) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateInstancePool) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateInstancePool) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateInstancePool) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateInstancePool) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateInstancePool) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes": basetypes.ListType{
				ElemType: InstancePoolAwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: InstancePoolAzureAttributes{}.Type(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec": basetypes.ListType{
				ElemType: DiskSpec{}.Type(ctx),
			},
			"enable_elastic_disk": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: InstancePoolGcpAttributes{}.Type(ctx),
			},
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
		},
	}
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
}

func (newState *CreateInstancePoolResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateInstancePoolResponse) {
}

func (newState *CreateInstancePoolResponse) SyncEffectiveFieldsDuringRead(existingState CreateInstancePoolResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateInstancePoolResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateInstancePoolResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateInstancePoolResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateInstancePoolResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateInstancePoolResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateInstancePoolResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateInstancePoolResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Definition types.String `tfsdk:"definition" tf:"optional"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description" tf:"optional"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries types.List `tfsdk:"libraries" tf:"optional"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user" tf:"optional"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides" tf:"optional"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId types.String `tfsdk:"policy_family_id" tf:"optional"`
}

func (newState *CreatePolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePolicy) {
}

func (newState *CreatePolicy) SyncEffectiveFieldsDuringRead(existingState CreatePolicy) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreatePolicy{}

// Equal implements basetypes.ObjectValuable.
func (o CreatePolicy) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreatePolicy) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreatePolicy) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreatePolicy) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreatePolicy) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreatePolicy) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
}

func (newState *CreatePolicyResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePolicyResponse) {
}

func (newState *CreatePolicyResponse) SyncEffectiveFieldsDuringRead(existingState CreatePolicyResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreatePolicyResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreatePolicyResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreatePolicyResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreatePolicyResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreatePolicyResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreatePolicyResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreatePolicyResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ScriptId types.String `tfsdk:"script_id" tf:"optional"`
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateResponse) {
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringRead(existingState CreateResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *Created) SyncEffectiveFieldsDuringCreateOrUpdate(plan Created) {
}

func (newState *Created) SyncEffectiveFieldsDuringRead(existingState Created) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Created{}

// Equal implements basetypes.ObjectValuable.
func (o Created) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Created) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Created) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Created) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Created) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Created) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Created) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DataPlaneEventDetails struct {
	// <needs content added>
	EventType types.String `tfsdk:"event_type" tf:"optional"`
	// <needs content added>
	ExecutorFailures types.Int64 `tfsdk:"executor_failures" tf:"optional"`
	// <needs content added>
	HostId types.String `tfsdk:"host_id" tf:"optional"`
	// <needs content added>
	Timestamp types.Int64 `tfsdk:"timestamp" tf:"optional"`
}

func (newState *DataPlaneEventDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataPlaneEventDetails) {
}

func (newState *DataPlaneEventDetails) SyncEffectiveFieldsDuringRead(existingState DataPlaneEventDetails) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DataPlaneEventDetails{}

// Equal implements basetypes.ObjectValuable.
func (o DataPlaneEventDetails) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DataPlaneEventDetails) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DataPlaneEventDetails) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DataPlaneEventDetails) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DataPlaneEventDetails) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DataPlaneEventDetails) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type DbfsStorageInfo struct {
	// dbfs destination, e.g. `dbfs:/my/path`
	Destination types.String `tfsdk:"destination" tf:""`
}

func (newState *DbfsStorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan DbfsStorageInfo) {
}

func (newState *DbfsStorageInfo) SyncEffectiveFieldsDuringRead(existingState DbfsStorageInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DbfsStorageInfo{}

// Equal implements basetypes.ObjectValuable.
func (o DbfsStorageInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DbfsStorageInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DbfsStorageInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DbfsStorageInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DbfsStorageInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DbfsStorageInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

func (newState *DeleteCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCluster) {
}

func (newState *DeleteCluster) SyncEffectiveFieldsDuringRead(existingState DeleteCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteCluster{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete init script
type DeleteGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-"`
}

func (newState *DeleteGlobalInitScriptRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteGlobalInitScriptRequest) {
}

func (newState *DeleteGlobalInitScriptRequest) SyncEffectiveFieldsDuringRead(existingState DeleteGlobalInitScriptRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteGlobalInitScriptRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteGlobalInitScriptRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteGlobalInitScriptRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteGlobalInitScriptRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteGlobalInitScriptRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteGlobalInitScriptRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteGlobalInitScriptRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:""`
}

func (newState *DeleteInstancePool) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteInstancePool) {
}

func (newState *DeleteInstancePool) SyncEffectiveFieldsDuringRead(existingState DeleteInstancePool) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteInstancePool{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteInstancePool) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteInstancePool) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteInstancePool) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteInstancePool) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteInstancePool) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteInstancePool) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteInstancePoolResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteInstancePoolResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	PolicyId types.String `tfsdk:"policy_id" tf:""`
}

func (newState *DeletePolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePolicy) {
}

func (newState *DeletePolicy) SyncEffectiveFieldsDuringRead(existingState DeletePolicy) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeletePolicy{}

// Equal implements basetypes.ObjectValuable.
func (o DeletePolicy) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeletePolicy) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeletePolicy) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeletePolicy) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeletePolicy) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeletePolicy) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeletePolicyResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeletePolicyResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeletePolicyResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeletePolicyResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeletePolicyResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeletePolicyResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeletePolicyResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DestroyContext struct {
	ClusterId types.String `tfsdk:"clusterId" tf:""`

	ContextId types.String `tfsdk:"contextId" tf:""`
}

func (newState *DestroyContext) SyncEffectiveFieldsDuringCreateOrUpdate(plan DestroyContext) {
}

func (newState *DestroyContext) SyncEffectiveFieldsDuringRead(existingState DestroyContext) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DestroyContext{}

// Equal implements basetypes.ObjectValuable.
func (o DestroyContext) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DestroyContext) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DestroyContext) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DestroyContext) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DestroyContext) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DestroyContext) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DestroyResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DestroyResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DestroyResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DestroyResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DestroyResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DestroyResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DestroyResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DestroyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

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
	DiskCount types.Int64 `tfsdk:"disk_count" tf:"optional"`

	DiskIops types.Int64 `tfsdk:"disk_iops" tf:"optional"`
	// The size of each disk (in GiB) launched for each instance. Values must
	// fall into the supported range for a particular instance type.
	//
	// For AWS: - General Purpose SSD: 100 - 4096 GiB - Throughput Optimized
	// HDD: 500 - 4096 GiB
	//
	// For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS (HDD): 1-
	// 1023 GiB
	DiskSize types.Int64 `tfsdk:"disk_size" tf:"optional"`

	DiskThroughput types.Int64 `tfsdk:"disk_throughput" tf:"optional"`
	// The type of disks that will be launched with this cluster.
	DiskType types.List `tfsdk:"disk_type" tf:"optional,object"`
}

func (newState *DiskSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan DiskSpec) {
}

func (newState *DiskSpec) SyncEffectiveFieldsDuringRead(existingState DiskSpec) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DiskSpec{}

// Equal implements basetypes.ObjectValuable.
func (o DiskSpec) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DiskSpec) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DiskSpec) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DiskSpec) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DiskSpec) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DiskSpec) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DiskSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disk_count":      types.Int64Type,
			"disk_iops":       types.Int64Type,
			"disk_size":       types.Int64Type,
			"disk_throughput": types.Int64Type,
			"disk_type": basetypes.ListType{
				ElemType: DiskType{}.Type(ctx),
			},
		},
	}
}

type DiskType struct {
	AzureDiskVolumeType types.String `tfsdk:"azure_disk_volume_type" tf:"optional"`

	EbsVolumeType types.String `tfsdk:"ebs_volume_type" tf:"optional"`
}

func (newState *DiskType) SyncEffectiveFieldsDuringCreateOrUpdate(plan DiskType) {
}

func (newState *DiskType) SyncEffectiveFieldsDuringRead(existingState DiskType) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DiskType{}

// Equal implements basetypes.ObjectValuable.
func (o DiskType) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DiskType) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DiskType) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DiskType) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DiskType) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DiskType) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Password types.String `tfsdk:"password" tf:"optional"`
	// Name of the user
	Username types.String `tfsdk:"username" tf:"optional"`
}

func (newState *DockerBasicAuth) SyncEffectiveFieldsDuringCreateOrUpdate(plan DockerBasicAuth) {
}

func (newState *DockerBasicAuth) SyncEffectiveFieldsDuringRead(existingState DockerBasicAuth) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DockerBasicAuth{}

// Equal implements basetypes.ObjectValuable.
func (o DockerBasicAuth) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DockerBasicAuth) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DockerBasicAuth) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DockerBasicAuth) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DockerBasicAuth) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DockerBasicAuth) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	BasicAuth types.List `tfsdk:"basic_auth" tf:"optional,object"`
	// URL of the docker image.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *DockerImage) SyncEffectiveFieldsDuringCreateOrUpdate(plan DockerImage) {
}

func (newState *DockerImage) SyncEffectiveFieldsDuringRead(existingState DockerImage) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DockerImage{}

// Equal implements basetypes.ObjectValuable.
func (o DockerImage) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DockerImage) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DockerImage) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DockerImage) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DockerImage) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DockerImage) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DockerImage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"basic_auth": basetypes.ListType{
				ElemType: DockerBasicAuth{}.Type(ctx),
			},
			"url": types.StringType,
		},
	}
}

type EditCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// ID of the cluster
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf" tf:"optional,object"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode types.String `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage types.List `tfsdk:"docker_image" tf:"optional,object"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
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
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf" tf:"optional"`
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
	SparkEnvVars types.Map `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:""`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType types.List `tfsdk:"workload_type" tf:"optional,object"`
}

func (newState *EditCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditCluster) {
}

func (newState *EditCluster) SyncEffectiveFieldsDuringRead(existingState EditCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EditCluster{}

// Equal implements basetypes.ObjectValuable.
func (o EditCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EditCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EditCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EditCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EditCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EditCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o EditCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale": basetypes.ListType{
				ElemType: AutoScale{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes{}.Type(ctx),
			},
			"cluster_id": types.StringType,
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id": types.StringType,
			"node_type_id":     types.StringType,
			"num_workers":      types.Int64Type,
			"policy_id":        types.StringType,
			"runtime_engine":   types.StringType,
			"single_user_name": types.StringType,
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType{}.Type(ctx),
			},
		},
	}
}

type EditClusterResponse struct {
}

func (newState *EditClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditClusterResponse) {
}

func (newState *EditClusterResponse) SyncEffectiveFieldsDuringRead(existingState EditClusterResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EditClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o EditClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EditClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EditClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EditClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EditClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EditClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes" tf:"optional"`
	// Instance pool ID
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:""`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name" tf:""`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity" tf:"optional"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:""`
}

func (newState *EditInstancePool) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditInstancePool) {
}

func (newState *EditInstancePool) SyncEffectiveFieldsDuringRead(existingState EditInstancePool) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EditInstancePool{}

// Equal implements basetypes.ObjectValuable.
func (o EditInstancePool) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EditInstancePool) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EditInstancePool) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EditInstancePool) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EditInstancePool) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EditInstancePool) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
		},
	}
}

type EditInstancePoolResponse struct {
}

func (newState *EditInstancePoolResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditInstancePoolResponse) {
}

func (newState *EditInstancePoolResponse) SyncEffectiveFieldsDuringRead(existingState EditInstancePoolResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EditInstancePoolResponse{}

// Equal implements basetypes.ObjectValuable.
func (o EditInstancePoolResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EditInstancePoolResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EditInstancePoolResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EditInstancePoolResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EditInstancePoolResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EditInstancePoolResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Definition types.String `tfsdk:"definition" tf:"optional"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description" tf:"optional"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries types.List `tfsdk:"libraries" tf:"optional"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user" tf:"optional"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides" tf:"optional"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId types.String `tfsdk:"policy_family_id" tf:"optional"`
	// The ID of the policy to update.
	PolicyId types.String `tfsdk:"policy_id" tf:""`
}

func (newState *EditPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPolicy) {
}

func (newState *EditPolicy) SyncEffectiveFieldsDuringRead(existingState EditPolicy) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EditPolicy{}

// Equal implements basetypes.ObjectValuable.
func (o EditPolicy) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EditPolicy) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EditPolicy) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EditPolicy) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EditPolicy) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EditPolicy) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type EditPolicyResponse struct {
}

func (newState *EditPolicyResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPolicyResponse) {
}

func (newState *EditPolicyResponse) SyncEffectiveFieldsDuringRead(existingState EditPolicyResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EditPolicyResponse{}

// Equal implements basetypes.ObjectValuable.
func (o EditPolicyResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EditPolicyResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EditPolicyResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EditPolicyResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EditPolicyResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EditPolicyResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EditResponse{}

// Equal implements basetypes.ObjectValuable.
func (o EditResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EditResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EditResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EditResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EditResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EditResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o EditResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EnforceClusterComplianceRequest struct {
	// The ID of the cluster you want to enforce policy compliance on.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// If set, previews the changes that would be made to a cluster to enforce
	// compliance but does not update the cluster.
	ValidateOnly types.Bool `tfsdk:"validate_only" tf:"optional"`
}

func (newState *EnforceClusterComplianceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforceClusterComplianceRequest) {
}

func (newState *EnforceClusterComplianceRequest) SyncEffectiveFieldsDuringRead(existingState EnforceClusterComplianceRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EnforceClusterComplianceRequest{}

// Equal implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Changes types.List `tfsdk:"changes" tf:"optional"`
	// Whether any changes have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	HasChanges types.Bool `tfsdk:"has_changes" tf:"optional"`
}

func (newState *EnforceClusterComplianceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforceClusterComplianceResponse) {
}

func (newState *EnforceClusterComplianceResponse) SyncEffectiveFieldsDuringRead(existingState EnforceClusterComplianceResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EnforceClusterComplianceResponse{}

// Equal implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EnforceClusterComplianceResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// The environment entity used to preserve serverless environment side panel and
// jobs' environment for non-notebook task. In this minimal environment spec,
// only pip dependencies are supported.
type Environment struct {
	// Client version used by the environment The client is the user-facing
	// environment of the runtime. Each client comes with a specific set of
	// pre-installed libraries. The version is a string, consisting of the major
	// client version.
	Client types.String `tfsdk:"client" tf:""`
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a pip requirement file line
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/ Allowed
	// dependency could be <requirement specifier>, <archive url/path>, <local
	// project path>(WSFS or Volumes in Databricks), <vcs project url> E.g.
	// dependencies: ["foo==0.0.1", "-r /Workspace/test/requirements.txt"]
	Dependencies types.List `tfsdk:"dependencies" tf:"optional"`
}

func (newState *Environment) SyncEffectiveFieldsDuringCreateOrUpdate(plan Environment) {
}

func (newState *Environment) SyncEffectiveFieldsDuringRead(existingState Environment) {
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
		"dependencies": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Environment{}

// Equal implements basetypes.ObjectValuable.
func (o Environment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Environment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Environment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Environment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Environment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Environment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Environment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"client": types.StringType,
			"dependencies": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type EventDetails struct {
	// * For created clusters, the attributes of the cluster. * For edited
	// clusters, the new attributes of the cluster.
	Attributes types.List `tfsdk:"attributes" tf:"optional,object"`
	// The cause of a change in target size.
	Cause types.String `tfsdk:"cause" tf:"optional"`
	// The actual cluster size that was set in the cluster creation or edit.
	ClusterSize types.List `tfsdk:"cluster_size" tf:"optional,object"`
	// The current number of vCPUs in the cluster.
	CurrentNumVcpus types.Int64 `tfsdk:"current_num_vcpus" tf:"optional"`
	// The current number of nodes in the cluster.
	CurrentNumWorkers types.Int64 `tfsdk:"current_num_workers" tf:"optional"`
	// <needs content added>
	DidNotExpandReason types.String `tfsdk:"did_not_expand_reason" tf:"optional"`
	// Current disk size in bytes
	DiskSize types.Int64 `tfsdk:"disk_size" tf:"optional"`
	// More details about the change in driver's state
	DriverStateMessage types.String `tfsdk:"driver_state_message" tf:"optional"`
	// Whether or not a blocklisted node should be terminated. For
	// ClusterEventType NODE_BLACKLISTED.
	EnableTerminationForNodeBlocklisted types.Bool `tfsdk:"enable_termination_for_node_blocklisted" tf:"optional"`
	// <needs content added>
	FreeSpace types.Int64 `tfsdk:"free_space" tf:"optional"`
	// List of global and cluster init scripts associated with this cluster
	// event.
	InitScripts types.List `tfsdk:"init_scripts" tf:"optional,object"`
	// Instance Id where the event originated from
	InstanceId types.String `tfsdk:"instance_id" tf:"optional"`
	// Unique identifier of the specific job run associated with this cluster
	// event * For clusters created for jobs, this will be the same as the
	// cluster name
	JobRunName types.String `tfsdk:"job_run_name" tf:"optional"`
	// The cluster attributes before a cluster was edited.
	PreviousAttributes types.List `tfsdk:"previous_attributes" tf:"optional,object"`
	// The size of the cluster before an edit or resize.
	PreviousClusterSize types.List `tfsdk:"previous_cluster_size" tf:"optional,object"`
	// Previous disk size in bytes
	PreviousDiskSize types.Int64 `tfsdk:"previous_disk_size" tf:"optional"`
	// A termination reason: * On a TERMINATED event, this is the reason of the
	// termination. * On a RESIZE_COMPLETE event, this indicates the reason that
	// we failed to acquire some nodes.
	Reason types.List `tfsdk:"reason" tf:"optional,object"`
	// The targeted number of vCPUs in the cluster.
	TargetNumVcpus types.Int64 `tfsdk:"target_num_vcpus" tf:"optional"`
	// The targeted number of nodes in the cluster.
	TargetNumWorkers types.Int64 `tfsdk:"target_num_workers" tf:"optional"`
	// The user that caused the event to occur. (Empty if it was done by the
	// control plane.)
	User types.String `tfsdk:"user" tf:"optional"`
}

func (newState *EventDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan EventDetails) {
}

func (newState *EventDetails) SyncEffectiveFieldsDuringRead(existingState EventDetails) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = EventDetails{}

// Equal implements basetypes.ObjectValuable.
func (o EventDetails) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o EventDetails) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o EventDetails) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o EventDetails) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o EventDetails) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o EventDetails) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o EventDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes": basetypes.ListType{
				ElemType: ClusterAttributes{}.Type(ctx),
			},
			"cause": types.StringType,
			"cluster_size": basetypes.ListType{
				ElemType: ClusterSize{}.Type(ctx),
			},
			"current_num_vcpus":                       types.Int64Type,
			"current_num_workers":                     types.Int64Type,
			"did_not_expand_reason":                   types.StringType,
			"disk_size":                               types.Int64Type,
			"driver_state_message":                    types.StringType,
			"enable_termination_for_node_blocklisted": types.BoolType,
			"free_space":                              types.Int64Type,
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptEventDetails{}.Type(ctx),
			},
			"instance_id":  types.StringType,
			"job_run_name": types.StringType,
			"previous_attributes": basetypes.ListType{
				ElemType: ClusterAttributes{}.Type(ctx),
			},
			"previous_cluster_size": basetypes.ListType{
				ElemType: ClusterSize{}.Type(ctx),
			},
			"previous_disk_size": types.Int64Type,
			"reason": basetypes.ListType{
				ElemType: TerminationReason{}.Type(ctx),
			},
			"target_num_vcpus":   types.Int64Type,
			"target_num_workers": types.Int64Type,
			"user":               types.StringType,
		},
	}
}

type GcpAttributes struct {
	// This field determines whether the instance pool will contain preemptible
	// VMs, on-demand VMs, or preemptible VMs with a fallback to on-demand VMs
	// if the former is unavailable.
	Availability types.String `tfsdk:"availability" tf:"optional"`
	// boot disk size in GB
	BootDiskSize types.Int64 `tfsdk:"boot_disk_size" tf:"optional"`
	// If provided, the cluster will impersonate the google service account when
	// accessing gcloud services (like GCS). The google service account must
	// have previously been added to the Databricks environment by an account
	// administrator.
	GoogleServiceAccount types.String `tfsdk:"google_service_account" tf:"optional"`
	// If provided, each node (workers and driver) in the cluster will have this
	// number of local SSDs attached. Each local SSD is 375GB in size. Refer to
	// [GCP documentation] for the supported number of local SSDs for each
	// instance type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount types.Int64 `tfsdk:"local_ssd_count" tf:"optional"`
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs (when set to true) versus standard compute engine
	// VMs (when set to false; default). Note: Soon to be deprecated, use the
	// availability field instead.
	UsePreemptibleExecutors types.Bool `tfsdk:"use_preemptible_executors" tf:"optional"`
	// Identifier for the availability zone in which the cluster resides. This
	// can be one of the following: - "HA" => High availability, spread nodes
	// across availability zones for a Databricks deployment region [default] -
	// "AUTO" => Databricks picks an availability zone to schedule the cluster
	// on. - A GCP availability zone => Pick One of the available zones for
	// (machine type + region) from
	// https://cloud.google.com/compute/docs/regions-zones.
	ZoneId types.String `tfsdk:"zone_id" tf:"optional"`
}

func (newState *GcpAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpAttributes) {
}

func (newState *GcpAttributes) SyncEffectiveFieldsDuringRead(existingState GcpAttributes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GcpAttributes{}

// Equal implements basetypes.ObjectValuable.
func (o GcpAttributes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GcpAttributes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GcpAttributes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GcpAttributes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GcpAttributes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GcpAttributes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type GcsStorageInfo struct {
	// GCS destination/URI, e.g. `gs://my-bucket/some-prefix`
	Destination types.String `tfsdk:"destination" tf:""`
}

func (newState *GcsStorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcsStorageInfo) {
}

func (newState *GcsStorageInfo) SyncEffectiveFieldsDuringRead(existingState GcsStorageInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GcsStorageInfo{}

// Equal implements basetypes.ObjectValuable.
func (o GcsStorageInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GcsStorageInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GcsStorageInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GcsStorageInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GcsStorageInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GcsStorageInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GcsStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

// Get cluster policy compliance
type GetClusterComplianceRequest struct {
	// The ID of the cluster to get the compliance status
	ClusterId types.String `tfsdk:"-"`
}

func (newState *GetClusterComplianceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterComplianceRequest) {
}

func (newState *GetClusterComplianceRequest) SyncEffectiveFieldsDuringRead(existingState GetClusterComplianceRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterComplianceRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterComplianceRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterComplianceRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterComplianceRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterComplianceRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterComplianceRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterComplianceRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	IsCompliant types.Bool `tfsdk:"is_compliant" tf:"optional"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	Violations types.Map `tfsdk:"violations" tf:"optional"`
}

func (newState *GetClusterComplianceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterComplianceResponse) {
}

func (newState *GetClusterComplianceResponse) SyncEffectiveFieldsDuringRead(existingState GetClusterComplianceResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterComplianceResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterComplianceResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterComplianceResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterComplianceResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterComplianceResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterComplianceResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterComplianceResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get cluster permission levels
type GetClusterPermissionLevelsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-"`
}

func (newState *GetClusterPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPermissionLevelsRequest) {
}

func (newState *GetClusterPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetClusterPermissionLevelsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterPermissionLevelsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetClusterPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPermissionLevelsResponse) {
}

func (newState *GetClusterPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetClusterPermissionLevelsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterPermissionLevelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterPermissionLevelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get cluster permissions
type GetClusterPermissionsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-"`
}

func (newState *GetClusterPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPermissionsRequest) {
}

func (newState *GetClusterPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetClusterPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id": types.StringType,
		},
	}
}

// Get cluster policy permission levels
type GetClusterPolicyPermissionLevelsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-"`
}

func (newState *GetClusterPolicyPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPolicyPermissionLevelsRequest) {
}

func (newState *GetClusterPolicyPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetClusterPolicyPermissionLevelsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterPolicyPermissionLevelsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetClusterPolicyPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPolicyPermissionLevelsResponse) {
}

func (newState *GetClusterPolicyPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetClusterPolicyPermissionLevelsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterPolicyPermissionLevelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionLevelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get cluster policy permissions
type GetClusterPolicyPermissionsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-"`
}

func (newState *GetClusterPolicyPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPolicyPermissionsRequest) {
}

func (newState *GetClusterPolicyPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetClusterPolicyPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterPolicyPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_policy_id": types.StringType,
		},
	}
}

// Get a cluster policy
type GetClusterPolicyRequest struct {
	// Canonical unique identifier for the Cluster Policy.
	PolicyId types.String `tfsdk:"-"`
}

func (newState *GetClusterPolicyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterPolicyRequest) {
}

func (newState *GetClusterPolicyRequest) SyncEffectiveFieldsDuringRead(existingState GetClusterPolicyRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterPolicyRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetClusterPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

// Get cluster info
type GetClusterRequest struct {
	// The cluster about which to retrieve information.
	ClusterId types.String `tfsdk:"-"`
}

func (newState *GetClusterRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetClusterRequest) {
}

func (newState *GetClusterRequest) SyncEffectiveFieldsDuringRead(existingState GetClusterRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetClusterRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetClusterRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetClusterRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetClusterRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetClusterRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetClusterRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetClusterRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// The end time in epoch milliseconds. If empty, returns events up to the
	// current time.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// An optional set of event types to filter on. If empty, all event types
	// are returned.
	EventTypes types.List `tfsdk:"event_types" tf:"optional"`
	// The maximum number of events to include in a page of events. Defaults to
	// 50, and maximum allowed value is 500.
	Limit types.Int64 `tfsdk:"limit" tf:"optional"`
	// The offset in the result set. Defaults to 0 (no offset). When an offset
	// is specified and the results are requested in descending order, the
	// end_time field is required.
	Offset types.Int64 `tfsdk:"offset" tf:"optional"`
	// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".
	Order types.String `tfsdk:"order" tf:"optional"`
	// The start time in epoch milliseconds. If empty, returns events starting
	// from the beginning of time.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
}

func (newState *GetEvents) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEvents) {
}

func (newState *GetEvents) SyncEffectiveFieldsDuringRead(existingState GetEvents) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetEvents{}

// Equal implements basetypes.ObjectValuable.
func (o GetEvents) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetEvents) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetEvents) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetEvents) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetEvents) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetEvents) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
			"start_time": types.Int64Type,
		},
	}
}

type GetEventsResponse struct {
	// <content needs to be added>
	Events types.List `tfsdk:"events" tf:"optional"`
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	NextPage types.List `tfsdk:"next_page" tf:"optional,object"`
	// The total number of events filtered by the start_time, end_time, and
	// event_types.
	TotalCount types.Int64 `tfsdk:"total_count" tf:"optional"`
}

func (newState *GetEventsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEventsResponse) {
}

func (newState *GetEventsResponse) SyncEffectiveFieldsDuringRead(existingState GetEventsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetEventsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetEventsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetEventsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetEventsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetEventsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetEventsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetEventsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetEventsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"events": basetypes.ListType{
				ElemType: ClusterEvent{}.Type(ctx),
			},
			"next_page": basetypes.ListType{
				ElemType: GetEvents{}.Type(ctx),
			},
			"total_count": types.Int64Type,
		},
	}
}

// Get an init script
type GetGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-"`
}

func (newState *GetGlobalInitScriptRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetGlobalInitScriptRequest) {
}

func (newState *GetGlobalInitScriptRequest) SyncEffectiveFieldsDuringRead(existingState GetGlobalInitScriptRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetGlobalInitScriptRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetGlobalInitScriptRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetGlobalInitScriptRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetGlobalInitScriptRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetGlobalInitScriptRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetGlobalInitScriptRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetGlobalInitScriptRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
	DefaultTags types.Map `tfsdk:"default_tags" tf:"optional"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec types.List `tfsdk:"disk_spec" tf:"optional,object"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes" tf:"optional"`
	// Canonical unique identifier for the pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:""`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name" tf:"optional"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity" tf:"optional"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Custom Docker Image BYOC
	PreloadedDockerImages types.List `tfsdk:"preloaded_docker_images" tf:"optional"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions types.List `tfsdk:"preloaded_spark_versions" tf:"optional"`
	// Current state of the instance pool.
	State types.String `tfsdk:"state" tf:"optional"`
	// Usage statistics about the instance pool.
	Stats types.List `tfsdk:"stats" tf:"optional,object"`
	// Status of failed pending instances in the pool.
	Status types.List `tfsdk:"status" tf:"optional,object"`
}

func (newState *GetInstancePool) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetInstancePool) {
}

func (newState *GetInstancePool) SyncEffectiveFieldsDuringRead(existingState GetInstancePool) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetInstancePool{}

// Equal implements basetypes.ObjectValuable.
func (o GetInstancePool) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetInstancePool) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetInstancePool) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetInstancePool) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetInstancePool) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetInstancePool) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes": basetypes.ListType{
				ElemType: InstancePoolAwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: InstancePoolAzureAttributes{}.Type(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"default_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec": basetypes.ListType{
				ElemType: DiskSpec{}.Type(ctx),
			},
			"enable_elastic_disk": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: InstancePoolGcpAttributes{}.Type(ctx),
			},
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
			"state": types.StringType,
			"stats": basetypes.ListType{
				ElemType: InstancePoolStats{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: InstancePoolStatus{}.Type(ctx),
			},
		},
	}
}

// Get instance pool permission levels
type GetInstancePoolPermissionLevelsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-"`
}

func (newState *GetInstancePoolPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetInstancePoolPermissionLevelsRequest) {
}

func (newState *GetInstancePoolPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetInstancePoolPermissionLevelsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetInstancePoolPermissionLevelsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetInstancePoolPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetInstancePoolPermissionLevelsResponse) {
}

func (newState *GetInstancePoolPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetInstancePoolPermissionLevelsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetInstancePoolPermissionLevelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionLevelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get instance pool permissions
type GetInstancePoolPermissionsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-"`
}

func (newState *GetInstancePoolPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetInstancePoolPermissionsRequest) {
}

func (newState *GetInstancePoolPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetInstancePoolPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetInstancePoolPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

// Get instance pool information
type GetInstancePoolRequest struct {
	// The canonical unique identifier for the instance pool.
	InstancePoolId types.String `tfsdk:"-"`
}

func (newState *GetInstancePoolRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetInstancePoolRequest) {
}

func (newState *GetInstancePoolRequest) SyncEffectiveFieldsDuringRead(existingState GetInstancePoolRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetInstancePoolRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetInstancePoolRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_pool_id": types.StringType,
		},
	}
}

// Get policy family information
type GetPolicyFamilyRequest struct {
	// The family ID about which to retrieve information.
	PolicyFamilyId types.String `tfsdk:"-"`
	// The version number for the family to fetch. Defaults to the latest
	// version.
	Version types.Int64 `tfsdk:"-"`
}

func (newState *GetPolicyFamilyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPolicyFamilyRequest) {
}

func (newState *GetPolicyFamilyRequest) SyncEffectiveFieldsDuringRead(existingState GetPolicyFamilyRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetPolicyFamilyRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetPolicyFamilyRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetPolicyFamilyRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetPolicyFamilyRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetPolicyFamilyRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetPolicyFamilyRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetPolicyFamilyRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Versions types.List `tfsdk:"versions" tf:"optional"`
}

func (newState *GetSparkVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSparkVersionsResponse) {
}

func (newState *GetSparkVersionsResponse) SyncEffectiveFieldsDuringRead(existingState GetSparkVersionsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetSparkVersionsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetSparkVersionsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetSparkVersionsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetSparkVersionsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetSparkVersionsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetSparkVersionsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetSparkVersionsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type GlobalInitScriptCreateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the script
	Name types.String `tfsdk:"name" tf:""`
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
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script" tf:""`
}

func (newState *GlobalInitScriptCreateRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GlobalInitScriptCreateRequest) {
}

func (newState *GlobalInitScriptCreateRequest) SyncEffectiveFieldsDuringRead(existingState GlobalInitScriptCreateRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GlobalInitScriptCreateRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GlobalInitScriptCreateRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GlobalInitScriptCreateRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GlobalInitScriptCreateRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GlobalInitScriptCreateRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GlobalInitScriptCreateRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GlobalInitScriptCreateRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The username of the user who created the script.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the script
	Name types.String `tfsdk:"name" tf:"optional"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id" tf:"optional"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The username of the user who last updated the script
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *GlobalInitScriptDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan GlobalInitScriptDetails) {
}

func (newState *GlobalInitScriptDetails) SyncEffectiveFieldsDuringRead(existingState GlobalInitScriptDetails) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GlobalInitScriptDetails{}

// Equal implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetails) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetails) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetails) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetails) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetails) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetails) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The username of the user who created the script.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the script
	Name types.String `tfsdk:"name" tf:"optional"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script" tf:"optional"`
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id" tf:"optional"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The username of the user who last updated the script
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *GlobalInitScriptDetailsWithContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan GlobalInitScriptDetailsWithContent) {
}

func (newState *GlobalInitScriptDetailsWithContent) SyncEffectiveFieldsDuringRead(existingState GlobalInitScriptDetailsWithContent) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GlobalInitScriptDetailsWithContent{}

// Equal implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetailsWithContent) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetailsWithContent) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetailsWithContent) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetailsWithContent) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetailsWithContent) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GlobalInitScriptDetailsWithContent) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the script
	Name types.String `tfsdk:"name" tf:""`
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
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script" tf:""`
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-"`
}

func (newState *GlobalInitScriptUpdateRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GlobalInitScriptUpdateRequest) {
}

func (newState *GlobalInitScriptUpdateRequest) SyncEffectiveFieldsDuringRead(existingState GlobalInitScriptUpdateRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GlobalInitScriptUpdateRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GlobalInitScriptUpdateRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GlobalInitScriptUpdateRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GlobalInitScriptUpdateRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GlobalInitScriptUpdateRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GlobalInitScriptUpdateRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GlobalInitScriptUpdateRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// The cluster scoped init scripts associated with this cluster event
	Cluster types.List `tfsdk:"cluster" tf:"optional"`
	// The global init scripts associated with this cluster event
	Global types.List `tfsdk:"global" tf:"optional"`
	// The private ip address of the node where the init scripts were run.
	ReportedForNode types.String `tfsdk:"reported_for_node" tf:"optional"`
}

func (newState *InitScriptEventDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan InitScriptEventDetails) {
}

func (newState *InitScriptEventDetails) SyncEffectiveFieldsDuringRead(existingState InitScriptEventDetails) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InitScriptEventDetails{}

// Equal implements basetypes.ObjectValuable.
func (o InitScriptEventDetails) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InitScriptEventDetails) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InitScriptEventDetails) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InitScriptEventDetails) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InitScriptEventDetails) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InitScriptEventDetails) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InitScriptExecutionDetails struct {
	// Addition details regarding errors.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// The duration of the script execution in seconds.
	ExecutionDurationSeconds types.Int64 `tfsdk:"execution_duration_seconds" tf:"optional"`
	// The current status of the script
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *InitScriptExecutionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan InitScriptExecutionDetails) {
}

func (newState *InitScriptExecutionDetails) SyncEffectiveFieldsDuringRead(existingState InitScriptExecutionDetails) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InitScriptExecutionDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a InitScriptExecutionDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InitScriptExecutionDetails{}

// Equal implements basetypes.ObjectValuable.
func (o InitScriptExecutionDetails) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InitScriptExecutionDetails) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InitScriptExecutionDetails) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InitScriptExecutionDetails) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InitScriptExecutionDetails) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InitScriptExecutionDetails) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o InitScriptExecutionDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_message":              types.StringType,
			"execution_duration_seconds": types.Int64Type,
			"status":                     types.StringType,
		},
	}
}

type InitScriptInfo struct {
	// destination needs to be provided. e.g. `{ "abfss" : { "destination" :
	// "abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>"
	// } }
	Abfss types.List `tfsdk:"abfss" tf:"optional,object"`
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs types.List `tfsdk:"dbfs" tf:"optional,object"`
	// destination needs to be provided. e.g. `{ "file" : { "destination" :
	// "file:/my/local/file.sh" } }`
	File types.List `tfsdk:"file" tf:"optional,object"`
	// destination needs to be provided. e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs types.List `tfsdk:"gcs" tf:"optional,object"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 types.List `tfsdk:"s3" tf:"optional,object"`
	// destination needs to be provided. e.g. `{ "volumes" : { "destination" :
	// "/Volumes/my-init.sh" } }`
	Volumes types.List `tfsdk:"volumes" tf:"optional,object"`
	// destination needs to be provided. e.g. `{ "workspace" : { "destination" :
	// "/Users/user1@databricks.com/my-init.sh" } }`
	Workspace types.List `tfsdk:"workspace" tf:"optional,object"`
}

func (newState *InitScriptInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan InitScriptInfo) {
}

func (newState *InitScriptInfo) SyncEffectiveFieldsDuringRead(existingState InitScriptInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InitScriptInfo{}

// Equal implements basetypes.ObjectValuable.
func (o InitScriptInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InitScriptInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InitScriptInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InitScriptInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InitScriptInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InitScriptInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o InitScriptInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"abfss": basetypes.ListType{
				ElemType: Adlsgen2Info{}.Type(ctx),
			},
			"dbfs": basetypes.ListType{
				ElemType: DbfsStorageInfo{}.Type(ctx),
			},
			"file": basetypes.ListType{
				ElemType: LocalFileInfo{}.Type(ctx),
			},
			"gcs": basetypes.ListType{
				ElemType: GcsStorageInfo{}.Type(ctx),
			},
			"s3": basetypes.ListType{
				ElemType: S3StorageInfo{}.Type(ctx),
			},
			"volumes": basetypes.ListType{
				ElemType: VolumesStorageInfo{}.Type(ctx),
			},
			"workspace": basetypes.ListType{
				ElemType: WorkspaceStorageInfo{}.Type(ctx),
			},
		},
	}
}

type InitScriptInfoAndExecutionDetails struct {
	// Details about the script
	ExecutionDetails types.List `tfsdk:"execution_details" tf:"optional,object"`
	// The script
	Script types.List `tfsdk:"script" tf:"optional,object"`
}

func (newState *InitScriptInfoAndExecutionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan InitScriptInfoAndExecutionDetails) {
}

func (newState *InitScriptInfoAndExecutionDetails) SyncEffectiveFieldsDuringRead(existingState InitScriptInfoAndExecutionDetails) {
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
		"execution_details": reflect.TypeOf(InitScriptExecutionDetails{}),
		"script":            reflect.TypeOf(InitScriptInfo{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InitScriptInfoAndExecutionDetails{}

// Equal implements basetypes.ObjectValuable.
func (o InitScriptInfoAndExecutionDetails) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InitScriptInfoAndExecutionDetails) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InitScriptInfoAndExecutionDetails) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InitScriptInfoAndExecutionDetails) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InitScriptInfoAndExecutionDetails) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InitScriptInfoAndExecutionDetails) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o InitScriptInfoAndExecutionDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"execution_details": basetypes.ListType{
				ElemType: InitScriptExecutionDetails{}.Type(ctx),
			},
			"script": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
		},
	}
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// The libraries to install.
	Libraries types.List `tfsdk:"libraries" tf:""`
}

func (newState *InstallLibraries) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstallLibraries) {
}

func (newState *InstallLibraries) SyncEffectiveFieldsDuringRead(existingState InstallLibraries) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstallLibraries{}

// Equal implements basetypes.ObjectValuable.
func (o InstallLibraries) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstallLibraries) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstallLibraries) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstallLibraries) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstallLibraries) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstallLibraries) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InstallLibrariesResponse struct {
}

func (newState *InstallLibrariesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstallLibrariesResponse) {
}

func (newState *InstallLibrariesResponse) SyncEffectiveFieldsDuringRead(existingState InstallLibrariesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstallLibrariesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o InstallLibrariesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type InstancePoolAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *InstancePoolAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAccessControlRequest) {
}

func (newState *InstancePoolAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState InstancePoolAccessControlRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolAccessControlRequest{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *InstancePoolAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAccessControlResponse) {
}

func (newState *InstancePoolAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState InstancePoolAccessControlResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolAccessControlResponse{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolAccessControlResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InstancePoolAndStats struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
	DefaultTags types.Map `tfsdk:"default_tags" tf:"optional"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec types.List `tfsdk:"disk_spec" tf:"optional,object"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes" tf:"optional"`
	// Canonical unique identifier for the pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name" tf:"optional"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity" tf:"optional"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Custom Docker Image BYOC
	PreloadedDockerImages types.List `tfsdk:"preloaded_docker_images" tf:"optional"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions types.List `tfsdk:"preloaded_spark_versions" tf:"optional"`
	// Current state of the instance pool.
	State types.String `tfsdk:"state" tf:"optional"`
	// Usage statistics about the instance pool.
	Stats types.List `tfsdk:"stats" tf:"optional,object"`
	// Status of failed pending instances in the pool.
	Status types.List `tfsdk:"status" tf:"optional,object"`
}

func (newState *InstancePoolAndStats) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAndStats) {
}

func (newState *InstancePoolAndStats) SyncEffectiveFieldsDuringRead(existingState InstancePoolAndStats) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolAndStats{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolAndStats) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolAndStats) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolAndStats) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolAndStats) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolAndStats) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolAndStats) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o InstancePoolAndStats) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_attributes": basetypes.ListType{
				ElemType: InstancePoolAwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: InstancePoolAzureAttributes{}.Type(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"default_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"disk_spec": basetypes.ListType{
				ElemType: DiskSpec{}.Type(ctx),
			},
			"enable_elastic_disk": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: InstancePoolGcpAttributes{}.Type(ctx),
			},
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
			"state": types.StringType,
			"stats": basetypes.ListType{
				ElemType: InstancePoolStats{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: InstancePoolStatus{}.Type(ctx),
			},
		},
	}
}

type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	//
	// The default value is defined by
	// InstancePoolConf.instancePoolDefaultAwsAvailability
	Availability types.String `tfsdk:"availability" tf:"optional"`
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
	//
	// The default value and documentation here should be kept consistent with
	// CommonConf.defaultSpotBidPricePercent and
	// CommonConf.maxSpotBidPricePercent.
	SpotBidPricePercent types.Int64 `tfsdk:"spot_bid_price_percent" tf:"optional"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. The list of available zones as well as the default value
	// can be found by using the `List Zones` method.
	ZoneId types.String `tfsdk:"zone_id" tf:"optional"`
}

func (newState *InstancePoolAwsAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAwsAttributes) {
}

func (newState *InstancePoolAwsAttributes) SyncEffectiveFieldsDuringRead(existingState InstancePoolAwsAttributes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolAwsAttributes{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolAwsAttributes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolAwsAttributes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolAwsAttributes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolAwsAttributes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolAwsAttributes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolAwsAttributes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InstancePoolAzureAttributes struct {
	// Shows the Availability type used for the spot nodes.
	//
	// The default value is defined by
	// InstancePoolConf.instancePoolDefaultAzureAvailability
	Availability types.String `tfsdk:"availability" tf:"optional"`
	// The default value and documentation here should be kept consistent with
	// CommonConf.defaultSpotBidMaxPrice.
	SpotBidMaxPrice types.Float64 `tfsdk:"spot_bid_max_price" tf:"optional"`
}

func (newState *InstancePoolAzureAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolAzureAttributes) {
}

func (newState *InstancePoolAzureAttributes) SyncEffectiveFieldsDuringRead(existingState InstancePoolAzureAttributes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolAzureAttributes{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolAzureAttributes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolAzureAttributes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolAzureAttributes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolAzureAttributes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolAzureAttributes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolAzureAttributes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InstancePoolGcpAttributes struct {
	// This field determines whether the instance pool will contain preemptible
	// VMs, on-demand VMs, or preemptible VMs with a fallback to on-demand VMs
	// if the former is unavailable.
	GcpAvailability types.String `tfsdk:"gcp_availability" tf:"optional"`
	// If provided, each node in the instance pool will have this number of
	// local SSDs attached. Each local SSD is 375GB in size. Refer to [GCP
	// documentation] for the supported number of local SSDs for each instance
	// type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount types.Int64 `tfsdk:"local_ssd_count" tf:"optional"`
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
	ZoneId types.String `tfsdk:"zone_id" tf:"optional"`
}

func (newState *InstancePoolGcpAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolGcpAttributes) {
}

func (newState *InstancePoolGcpAttributes) SyncEffectiveFieldsDuringRead(existingState InstancePoolGcpAttributes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolGcpAttributes{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolGcpAttributes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolGcpAttributes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolGcpAttributes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolGcpAttributes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolGcpAttributes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolGcpAttributes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *InstancePoolPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolPermission) {
}

func (newState *InstancePoolPermission) SyncEffectiveFieldsDuringRead(existingState InstancePoolPermission) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolPermission{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolPermission) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolPermission) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolPermission) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolPermission) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolPermission) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolPermission) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InstancePoolPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *InstancePoolPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolPermissions) {
}

func (newState *InstancePoolPermissions) SyncEffectiveFieldsDuringRead(existingState InstancePoolPermissions) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolPermissions{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolPermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolPermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolPermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolPermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolPermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolPermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InstancePoolPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *InstancePoolPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolPermissionsDescription) {
}

func (newState *InstancePoolPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState InstancePoolPermissionsDescription) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolPermissionsDescription{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsDescription) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsDescription) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsDescription) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsDescription) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsDescription) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsDescription) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-"`
}

func (newState *InstancePoolPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolPermissionsRequest) {
}

func (newState *InstancePoolPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState InstancePoolPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InstancePoolStats struct {
	// Number of active instances in the pool that are NOT part of a cluster.
	IdleCount types.Int64 `tfsdk:"idle_count" tf:"optional"`
	// Number of pending instances in the pool that are NOT part of a cluster.
	PendingIdleCount types.Int64 `tfsdk:"pending_idle_count" tf:"optional"`
	// Number of pending instances in the pool that are part of a cluster.
	PendingUsedCount types.Int64 `tfsdk:"pending_used_count" tf:"optional"`
	// Number of active instances in the pool that are part of a cluster.
	UsedCount types.Int64 `tfsdk:"used_count" tf:"optional"`
}

func (newState *InstancePoolStats) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolStats) {
}

func (newState *InstancePoolStats) SyncEffectiveFieldsDuringRead(existingState InstancePoolStats) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolStats{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolStats) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolStats) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolStats) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolStats) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolStats) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolStats) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	PendingInstanceErrors types.List `tfsdk:"pending_instance_errors" tf:"optional"`
}

func (newState *InstancePoolStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstancePoolStatus) {
}

func (newState *InstancePoolStatus) SyncEffectiveFieldsDuringRead(existingState InstancePoolStatus) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstancePoolStatus{}

// Equal implements basetypes.ObjectValuable.
func (o InstancePoolStatus) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstancePoolStatus) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstancePoolStatus) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstancePoolStatus) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstancePoolStatus) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstancePoolStatus) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InstanceProfile struct {
	// The AWS IAM role ARN of the role associated with the instance profile.
	// This field is required if your role name and instance profile name do not
	// match and you want to use the instance profile with [Databricks SQL
	// Serverless].
	//
	// Otherwise, this field is optional.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	IamRoleArn types.String `tfsdk:"iam_role_arn" tf:"optional"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:""`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile types.Bool `tfsdk:"is_meta_instance_profile" tf:"optional"`
}

func (newState *InstanceProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan InstanceProfile) {
}

func (newState *InstanceProfile) SyncEffectiveFieldsDuringRead(existingState InstanceProfile) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InstanceProfile{}

// Equal implements basetypes.ObjectValuable.
func (o InstanceProfile) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InstanceProfile) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InstanceProfile) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InstanceProfile) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InstanceProfile) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InstanceProfile) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Cran types.List `tfsdk:"cran" tf:"optional,object"`
	// Deprecated. URI of the egg library to install. Installing Python egg
	// files is deprecated and is not supported in Databricks Runtime 14.0 and
	// above.
	Egg types.String `tfsdk:"egg" tf:"optional"`
	// URI of the JAR library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "jar":
	// "/Workspace/path/to/library.jar" }`, `{ "jar" :
	// "/Volumes/path/to/library.jar" }` or `{ "jar":
	// "s3://my-bucket/library.jar" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Jar types.String `tfsdk:"jar" tf:"optional"`
	// Specification of a maven library to be installed. For example: `{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }`
	Maven types.List `tfsdk:"maven" tf:"optional,object"`
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	Pypi types.List `tfsdk:"pypi" tf:"optional,object"`
	// URI of the requirements.txt file to install. Only Workspace paths and
	// Unity Catalog Volumes paths are supported. For example: `{
	// "requirements": "/Workspace/path/to/requirements.txt" }` or `{
	// "requirements" : "/Volumes/path/to/requirements.txt" }`
	Requirements types.String `tfsdk:"requirements" tf:"optional"`
	// URI of the wheel library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "whl":
	// "/Workspace/path/to/library.whl" }`, `{ "whl" :
	// "/Volumes/path/to/library.whl" }` or `{ "whl":
	// "s3://my-bucket/library.whl" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Whl types.String `tfsdk:"whl" tf:"optional"`
}

func (newState *Library) SyncEffectiveFieldsDuringCreateOrUpdate(plan Library) {
}

func (newState *Library) SyncEffectiveFieldsDuringRead(existingState Library) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Library{}

// Equal implements basetypes.ObjectValuable.
func (o Library) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Library) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Library) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Library) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Library) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Library) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Library) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cran": basetypes.ListType{
				ElemType: RCranLibrary{}.Type(ctx),
			},
			"egg": types.StringType,
			"jar": types.StringType,
			"maven": basetypes.ListType{
				ElemType: MavenLibrary{}.Type(ctx),
			},
			"pypi": basetypes.ListType{
				ElemType: PythonPyPiLibrary{}.Type(ctx),
			},
			"requirements": types.StringType,
			"whl":          types.StringType,
		},
	}
}

// The status of the library on a specific cluster.
type LibraryFullStatus struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	IsLibraryForAllClusters types.Bool `tfsdk:"is_library_for_all_clusters" tf:"optional"`
	// Unique identifier for the library.
	Library types.List `tfsdk:"library" tf:"optional,object"`
	// All the info and warning messages that have occurred so far for this
	// library.
	Messages types.List `tfsdk:"messages" tf:"optional"`
	// Status of installing the library on the cluster.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *LibraryFullStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan LibraryFullStatus) {
}

func (newState *LibraryFullStatus) SyncEffectiveFieldsDuringRead(existingState LibraryFullStatus) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LibraryFullStatus{}

// Equal implements basetypes.ObjectValuable.
func (o LibraryFullStatus) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LibraryFullStatus) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LibraryFullStatus) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LibraryFullStatus) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LibraryFullStatus) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LibraryFullStatus) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o LibraryFullStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_library_for_all_clusters": types.BoolType,
			"library": basetypes.ListType{
				ElemType: Library{}.Type(ctx),
			},
			"messages": basetypes.ListType{
				ElemType: types.StringType,
			},
			"status": types.StringType,
		},
	}
}

type ListAllClusterLibraryStatusesResponse struct {
	// A list of cluster statuses.
	Statuses types.List `tfsdk:"statuses" tf:"optional"`
}

func (newState *ListAllClusterLibraryStatusesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAllClusterLibraryStatusesResponse) {
}

func (newState *ListAllClusterLibraryStatusesResponse) SyncEffectiveFieldsDuringRead(existingState ListAllClusterLibraryStatusesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListAllClusterLibraryStatusesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatusesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatusesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatusesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatusesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatusesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListAllClusterLibraryStatusesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListAvailableZonesResponse struct {
	// The availability zone if no `zone_id` is provided in the cluster creation
	// request.
	DefaultZone types.String `tfsdk:"default_zone" tf:"optional"`
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	Zones types.List `tfsdk:"zones" tf:"optional"`
}

func (newState *ListAvailableZonesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAvailableZonesResponse) {
}

func (newState *ListAvailableZonesResponse) SyncEffectiveFieldsDuringRead(existingState ListAvailableZonesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListAvailableZonesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListAvailableZonesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListAvailableZonesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListAvailableZonesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListAvailableZonesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListAvailableZonesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListAvailableZonesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// List cluster policy compliance
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

func (newState *ListClusterCompliancesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClusterCompliancesRequest) {
}

func (newState *ListClusterCompliancesRequest) SyncEffectiveFieldsDuringRead(existingState ListClusterCompliancesRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListClusterCompliancesRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListClusterCompliancesRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListClusterCompliancesRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListClusterCompliancesRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListClusterCompliancesRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListClusterCompliancesRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListClusterCompliancesRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Clusters types.List `tfsdk:"clusters" tf:"optional"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`
}

func (newState *ListClusterCompliancesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClusterCompliancesResponse) {
}

func (newState *ListClusterCompliancesResponse) SyncEffectiveFieldsDuringRead(existingState ListClusterCompliancesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListClusterCompliancesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListClusterCompliancesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListClusterCompliancesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListClusterCompliancesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListClusterCompliancesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListClusterCompliancesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListClusterCompliancesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// List cluster policies
type ListClusterPoliciesRequest struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	SortColumn types.String `tfsdk:"-"`
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	SortOrder types.String `tfsdk:"-"`
}

func (newState *ListClusterPoliciesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClusterPoliciesRequest) {
}

func (newState *ListClusterPoliciesRequest) SyncEffectiveFieldsDuringRead(existingState ListClusterPoliciesRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListClusterPoliciesRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListClusterPoliciesRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListClusterPoliciesRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListClusterPoliciesRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListClusterPoliciesRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListClusterPoliciesRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListClusterPoliciesRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ClusterSources types.List `tfsdk:"cluster_sources" tf:"optional"`
	// The current state of the clusters.
	ClusterStates types.List `tfsdk:"cluster_states" tf:"optional"`
	// Whether the clusters are pinned or not.
	IsPinned types.Bool `tfsdk:"is_pinned" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
}

func (newState *ListClustersFilterBy) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClustersFilterBy) {
}

func (newState *ListClustersFilterBy) SyncEffectiveFieldsDuringRead(existingState ListClustersFilterBy) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListClustersFilterBy{}

// Equal implements basetypes.ObjectValuable.
func (o ListClustersFilterBy) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListClustersFilterBy) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListClustersFilterBy) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListClustersFilterBy) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListClustersFilterBy) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListClustersFilterBy) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// List clusters
type ListClustersRequest struct {
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

func (newState *ListClustersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClustersRequest) {
}

func (newState *ListClustersRequest) SyncEffectiveFieldsDuringRead(existingState ListClustersRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListClustersRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListClustersRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListClustersRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListClustersRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListClustersRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListClustersRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListClustersRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListClustersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_by": basetypes.ListType{
				ElemType: ListClustersFilterBy{}.Type(ctx),
			},
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"sort_by": basetypes.ListType{
				ElemType: ListClustersSortBy{}.Type(ctx),
			},
		},
	}
}

type ListClustersResponse struct {
	// <needs content added>
	Clusters types.List `tfsdk:"clusters" tf:"optional"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`
}

func (newState *ListClustersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClustersResponse) {
}

func (newState *ListClustersResponse) SyncEffectiveFieldsDuringRead(existingState ListClustersResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListClustersResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListClustersResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListClustersResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListClustersResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListClustersResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListClustersResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListClustersResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListClustersSortBy struct {
	// The direction to sort by.
	Direction types.String `tfsdk:"direction" tf:"optional"`
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	Field types.String `tfsdk:"field" tf:"optional"`
}

func (newState *ListClustersSortBy) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListClustersSortBy) {
}

func (newState *ListClustersSortBy) SyncEffectiveFieldsDuringRead(existingState ListClustersSortBy) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListClustersSortBy{}

// Equal implements basetypes.ObjectValuable.
func (o ListClustersSortBy) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListClustersSortBy) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListClustersSortBy) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListClustersSortBy) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListClustersSortBy) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListClustersSortBy) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListGlobalInitScriptsResponse struct {
	Scripts types.List `tfsdk:"scripts" tf:"optional"`
}

func (newState *ListGlobalInitScriptsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListGlobalInitScriptsResponse) {
}

func (newState *ListGlobalInitScriptsResponse) SyncEffectiveFieldsDuringRead(existingState ListGlobalInitScriptsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListGlobalInitScriptsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListGlobalInitScriptsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListInstancePools struct {
	InstancePools types.List `tfsdk:"instance_pools" tf:"optional"`
}

func (newState *ListInstancePools) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListInstancePools) {
}

func (newState *ListInstancePools) SyncEffectiveFieldsDuringRead(existingState ListInstancePools) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListInstancePools{}

// Equal implements basetypes.ObjectValuable.
func (o ListInstancePools) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListInstancePools) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListInstancePools) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListInstancePools) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListInstancePools) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListInstancePools) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListInstanceProfilesResponse struct {
	// A list of instance profiles that the user can access.
	InstanceProfiles types.List `tfsdk:"instance_profiles" tf:"optional"`
}

func (newState *ListInstanceProfilesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListInstanceProfilesResponse) {
}

func (newState *ListInstanceProfilesResponse) SyncEffectiveFieldsDuringRead(existingState ListInstanceProfilesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListInstanceProfilesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListInstanceProfilesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListInstanceProfilesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListInstanceProfilesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListInstanceProfilesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListInstanceProfilesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListInstanceProfilesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListNodeTypesResponse struct {
	// The list of available Spark node types.
	NodeTypes types.List `tfsdk:"node_types" tf:"optional"`
}

func (newState *ListNodeTypesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNodeTypesResponse) {
}

func (newState *ListNodeTypesResponse) SyncEffectiveFieldsDuringRead(existingState ListNodeTypesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListNodeTypesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListNodeTypesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListNodeTypesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListNodeTypesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListNodeTypesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListNodeTypesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListNodeTypesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListPoliciesResponse struct {
	// List of policies.
	Policies types.List `tfsdk:"policies" tf:"optional"`
}

func (newState *ListPoliciesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPoliciesResponse) {
}

func (newState *ListPoliciesResponse) SyncEffectiveFieldsDuringRead(existingState ListPoliciesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListPoliciesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListPoliciesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListPoliciesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListPoliciesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListPoliciesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListPoliciesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListPoliciesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// List policy families
type ListPolicyFamiliesRequest struct {
	// Maximum number of policy families to return.
	MaxResults types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListPolicyFamiliesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPolicyFamiliesRequest) {
}

func (newState *ListPolicyFamiliesRequest) SyncEffectiveFieldsDuringRead(existingState ListPolicyFamiliesRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListPolicyFamiliesRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// List of policy families.
	PolicyFamilies types.List `tfsdk:"policy_families" tf:"optional"`
}

func (newState *ListPolicyFamiliesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPolicyFamiliesResponse) {
}

func (newState *ListPolicyFamiliesResponse) SyncEffectiveFieldsDuringRead(existingState ListPolicyFamiliesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListPolicyFamiliesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListPolicyFamiliesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type LocalFileInfo struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	Destination types.String `tfsdk:"destination" tf:""`
}

func (newState *LocalFileInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan LocalFileInfo) {
}

func (newState *LocalFileInfo) SyncEffectiveFieldsDuringRead(existingState LocalFileInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LocalFileInfo{}

// Equal implements basetypes.ObjectValuable.
func (o LocalFileInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LocalFileInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LocalFileInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LocalFileInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LocalFileInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LocalFileInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// <needs content added>
	LogAnalyticsPrimaryKey types.String `tfsdk:"log_analytics_primary_key" tf:"optional"`
	// <needs content added>
	LogAnalyticsWorkspaceId types.String `tfsdk:"log_analytics_workspace_id" tf:"optional"`
}

func (newState *LogAnalyticsInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogAnalyticsInfo) {
}

func (newState *LogAnalyticsInfo) SyncEffectiveFieldsDuringRead(existingState LogAnalyticsInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogAnalyticsInfo{}

// Equal implements basetypes.ObjectValuable.
func (o LogAnalyticsInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogAnalyticsInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogAnalyticsInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogAnalyticsInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogAnalyticsInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogAnalyticsInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type LogSyncStatus struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	LastAttempted types.Int64 `tfsdk:"last_attempted" tf:"optional"`
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	LastException types.String `tfsdk:"last_exception" tf:"optional"`
}

func (newState *LogSyncStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogSyncStatus) {
}

func (newState *LogSyncStatus) SyncEffectiveFieldsDuringRead(existingState LogSyncStatus) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogSyncStatus{}

// Equal implements basetypes.ObjectValuable.
func (o LogSyncStatus) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogSyncStatus) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogSyncStatus) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogSyncStatus) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogSyncStatus) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogSyncStatus) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Coordinates types.String `tfsdk:"coordinates" tf:""`
	// List of dependences to exclude. For example: `["slf4j:slf4j",
	// "*:hadoop-client"]`.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	Exclusions types.List `tfsdk:"exclusions" tf:"optional"`
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	Repo types.String `tfsdk:"repo" tf:"optional"`
}

func (newState *MavenLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan MavenLibrary) {
}

func (newState *MavenLibrary) SyncEffectiveFieldsDuringRead(existingState MavenLibrary) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = MavenLibrary{}

// Equal implements basetypes.ObjectValuable.
func (o MavenLibrary) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o MavenLibrary) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o MavenLibrary) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o MavenLibrary) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o MavenLibrary) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o MavenLibrary) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type NodeInstanceType struct {
	InstanceTypeId types.String `tfsdk:"instance_type_id" tf:"optional"`

	LocalDiskSizeGb types.Int64 `tfsdk:"local_disk_size_gb" tf:"optional"`

	LocalDisks types.Int64 `tfsdk:"local_disks" tf:"optional"`

	LocalNvmeDiskSizeGb types.Int64 `tfsdk:"local_nvme_disk_size_gb" tf:"optional"`

	LocalNvmeDisks types.Int64 `tfsdk:"local_nvme_disks" tf:"optional"`
}

func (newState *NodeInstanceType) SyncEffectiveFieldsDuringCreateOrUpdate(plan NodeInstanceType) {
}

func (newState *NodeInstanceType) SyncEffectiveFieldsDuringRead(existingState NodeInstanceType) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = NodeInstanceType{}

// Equal implements basetypes.ObjectValuable.
func (o NodeInstanceType) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o NodeInstanceType) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o NodeInstanceType) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o NodeInstanceType) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o NodeInstanceType) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o NodeInstanceType) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type NodeType struct {
	Category types.String `tfsdk:"category" tf:"optional"`
	// A string description associated with this node type, e.g., "r3.xlarge".
	Description types.String `tfsdk:"description" tf:""`

	DisplayOrder types.Int64 `tfsdk:"display_order" tf:"optional"`
	// An identifier for the type of hardware that this node runs on, e.g.,
	// "r3.2xlarge" in AWS.
	InstanceTypeId types.String `tfsdk:"instance_type_id" tf:""`
	// Whether the node type is deprecated. Non-deprecated node types offer
	// greater performance.
	IsDeprecated types.Bool `tfsdk:"is_deprecated" tf:"optional"`
	// AWS specific, whether this instance supports encryption in transit, used
	// for hipaa and pci workloads.
	IsEncryptedInTransit types.Bool `tfsdk:"is_encrypted_in_transit" tf:"optional"`

	IsGraviton types.Bool `tfsdk:"is_graviton" tf:"optional"`

	IsHidden types.Bool `tfsdk:"is_hidden" tf:"optional"`

	IsIoCacheEnabled types.Bool `tfsdk:"is_io_cache_enabled" tf:"optional"`
	// Memory (in MB) available for this node type.
	MemoryMb types.Int64 `tfsdk:"memory_mb" tf:""`

	NodeInfo types.List `tfsdk:"node_info" tf:"optional,object"`

	NodeInstanceType types.List `tfsdk:"node_instance_type" tf:"optional,object"`
	// Unique identifier for this node type.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:""`
	// Number of CPU cores available for this node type. Note that this can be
	// fractional, e.g., 2.5 cores, if the the number of cores on a machine
	// instance is not divisible by the number of Spark nodes on that machine.
	NumCores types.Float64 `tfsdk:"num_cores" tf:""`

	NumGpus types.Int64 `tfsdk:"num_gpus" tf:"optional"`

	PhotonDriverCapable types.Bool `tfsdk:"photon_driver_capable" tf:"optional"`

	PhotonWorkerCapable types.Bool `tfsdk:"photon_worker_capable" tf:"optional"`

	SupportClusterTags types.Bool `tfsdk:"support_cluster_tags" tf:"optional"`

	SupportEbsVolumes types.Bool `tfsdk:"support_ebs_volumes" tf:"optional"`

	SupportPortForwarding types.Bool `tfsdk:"support_port_forwarding" tf:"optional"`
	// Indicates if this node type can be used for an instance pool or cluster
	// with elastic disk enabled. This is true for most node types.
	SupportsElasticDisk types.Bool `tfsdk:"supports_elastic_disk" tf:"optional"`
}

func (newState *NodeType) SyncEffectiveFieldsDuringCreateOrUpdate(plan NodeType) {
}

func (newState *NodeType) SyncEffectiveFieldsDuringRead(existingState NodeType) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = NodeType{}

// Equal implements basetypes.ObjectValuable.
func (o NodeType) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o NodeType) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o NodeType) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o NodeType) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o NodeType) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o NodeType) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
			"node_info": basetypes.ListType{
				ElemType: CloudProviderNodeInfo{}.Type(ctx),
			},
			"node_instance_type": basetypes.ListType{
				ElemType: NodeInstanceType{}.Type(ctx),
			},
			"node_type_id":            types.StringType,
			"num_cores":               types.Float64Type,
			"num_gpus":                types.Int64Type,
			"photon_driver_capable":   types.BoolType,
			"photon_worker_capable":   types.BoolType,
			"support_cluster_tags":    types.BoolType,
			"support_ebs_volumes":     types.BoolType,
			"support_port_forwarding": types.BoolType,
			"supports_elastic_disk":   types.BoolType,
		},
	}
}

type PendingInstanceError struct {
	InstanceId types.String `tfsdk:"instance_id" tf:"optional"`

	Message types.String `tfsdk:"message" tf:"optional"`
}

func (newState *PendingInstanceError) SyncEffectiveFieldsDuringCreateOrUpdate(plan PendingInstanceError) {
}

func (newState *PendingInstanceError) SyncEffectiveFieldsDuringRead(existingState PendingInstanceError) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PendingInstanceError{}

// Equal implements basetypes.ObjectValuable.
func (o PendingInstanceError) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PendingInstanceError) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PendingInstanceError) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PendingInstanceError) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PendingInstanceError) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PendingInstanceError) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

func (newState *PermanentDeleteCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermanentDeleteCluster) {
}

func (newState *PermanentDeleteCluster) SyncEffectiveFieldsDuringRead(existingState PermanentDeleteCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PermanentDeleteCluster{}

// Equal implements basetypes.ObjectValuable.
func (o PermanentDeleteCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PermanentDeleteCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PermanentDeleteCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PermanentDeleteCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PermanentDeleteCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PermanentDeleteCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PermanentDeleteClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PermanentDeleteClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type PinCluster struct {
	// <needs content added>
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

func (newState *PinCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan PinCluster) {
}

func (newState *PinCluster) SyncEffectiveFieldsDuringRead(existingState PinCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PinCluster{}

// Equal implements basetypes.ObjectValuable.
func (o PinCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PinCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PinCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PinCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PinCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PinCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PinClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o PinClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PinClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PinClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PinClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PinClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PinClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	CreatedAtTimestamp types.Int64 `tfsdk:"created_at_timestamp" tf:"optional"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition" tf:"optional"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description" tf:"optional"`
	// If true, policy is a default policy created and managed by Databricks.
	// Default policies cannot be deleted, and their policy families cannot be
	// changed.
	IsDefault types.Bool `tfsdk:"is_default" tf:"optional"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries types.List `tfsdk:"libraries" tf:"optional"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user" tf:"optional"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides" tf:"optional"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId types.String `tfsdk:"policy_family_id" tf:"optional"`
	// Canonical unique identifier for the Cluster Policy.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
}

func (newState *Policy) SyncEffectiveFieldsDuringCreateOrUpdate(plan Policy) {
}

func (newState *Policy) SyncEffectiveFieldsDuringRead(existingState Policy) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Policy{}

// Equal implements basetypes.ObjectValuable.
func (o Policy) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Policy) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Policy) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Policy) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Policy) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Policy) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type PolicyFamily struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition" tf:"optional"`
	// Human-readable description of the purpose of the policy family.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Name of the policy family.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Unique identifier for the policy family.
	PolicyFamilyId types.String `tfsdk:"policy_family_id" tf:"optional"`
}

func (newState *PolicyFamily) SyncEffectiveFieldsDuringCreateOrUpdate(plan PolicyFamily) {
}

func (newState *PolicyFamily) SyncEffectiveFieldsDuringRead(existingState PolicyFamily) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PolicyFamily{}

// Equal implements basetypes.ObjectValuable.
func (o PolicyFamily) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PolicyFamily) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PolicyFamily) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PolicyFamily) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PolicyFamily) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PolicyFamily) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Package types.String `tfsdk:"package" tf:""`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo types.String `tfsdk:"repo" tf:"optional"`
}

func (newState *PythonPyPiLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan PythonPyPiLibrary) {
}

func (newState *PythonPyPiLibrary) SyncEffectiveFieldsDuringRead(existingState PythonPyPiLibrary) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PythonPyPiLibrary{}

// Equal implements basetypes.ObjectValuable.
func (o PythonPyPiLibrary) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PythonPyPiLibrary) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PythonPyPiLibrary) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PythonPyPiLibrary) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PythonPyPiLibrary) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PythonPyPiLibrary) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Package types.String `tfsdk:"package" tf:""`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo types.String `tfsdk:"repo" tf:"optional"`
}

func (newState *RCranLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan RCranLibrary) {
}

func (newState *RCranLibrary) SyncEffectiveFieldsDuringRead(existingState RCranLibrary) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RCranLibrary{}

// Equal implements basetypes.ObjectValuable.
func (o RCranLibrary) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RCranLibrary) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RCranLibrary) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RCranLibrary) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RCranLibrary) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RCranLibrary) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:""`
}

func (newState *RemoveInstanceProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan RemoveInstanceProfile) {
}

func (newState *RemoveInstanceProfile) SyncEffectiveFieldsDuringRead(existingState RemoveInstanceProfile) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RemoveInstanceProfile{}

// Equal implements basetypes.ObjectValuable.
func (o RemoveInstanceProfile) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RemoveInstanceProfile) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RemoveInstanceProfile) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RemoveInstanceProfile) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RemoveInstanceProfile) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RemoveInstanceProfile) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RemoveResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RemoveResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RemoveResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RemoveResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RemoveResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RemoveResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RemoveResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
	// The cluster to be resized.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
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
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
}

func (newState *ResizeCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResizeCluster) {
}

func (newState *ResizeCluster) SyncEffectiveFieldsDuringRead(existingState ResizeCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ResizeCluster{}

// Equal implements basetypes.ObjectValuable.
func (o ResizeCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ResizeCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ResizeCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ResizeCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ResizeCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ResizeCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ResizeCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale": basetypes.ListType{
				ElemType: AutoScale{}.Type(ctx),
			},
			"cluster_id":  types.StringType,
			"num_workers": types.Int64Type,
		},
	}
}

type ResizeClusterResponse struct {
}

func (newState *ResizeClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResizeClusterResponse) {
}

func (newState *ResizeClusterResponse) SyncEffectiveFieldsDuringRead(existingState ResizeClusterResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ResizeClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ResizeClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ResizeClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ResizeClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ResizeClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ResizeClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ResizeClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ResizeClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestartCluster struct {
	// The cluster to be started.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// <needs content added>
	RestartUser types.String `tfsdk:"restart_user" tf:"optional"`
}

func (newState *RestartCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestartCluster) {
}

func (newState *RestartCluster) SyncEffectiveFieldsDuringRead(existingState RestartCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RestartCluster{}

// Equal implements basetypes.ObjectValuable.
func (o RestartCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RestartCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RestartCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RestartCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RestartCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RestartCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RestartClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RestartClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RestartClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RestartClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RestartClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RestartClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RestartClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RestartClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Results struct {
	// The cause of the error
	Cause types.String `tfsdk:"cause" tf:"optional"`

	Data any `tfsdk:"data" tf:"optional"`
	// The image filename
	FileName types.String `tfsdk:"fileName" tf:"optional"`

	FileNames types.List `tfsdk:"fileNames" tf:"optional"`
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	IsJsonSchema types.Bool `tfsdk:"isJsonSchema" tf:"optional"`
	// internal field used by SDK
	Pos types.Int64 `tfsdk:"pos" tf:"optional"`

	ResultType types.String `tfsdk:"resultType" tf:"optional"`
	// The table schema
	Schema types.List `tfsdk:"schema" tf:"optional"`
	// The summary of the error
	Summary types.String `tfsdk:"summary" tf:"optional"`
	// true if partial results are returned.
	Truncated types.Bool `tfsdk:"truncated" tf:"optional"`
}

func (newState *Results) SyncEffectiveFieldsDuringCreateOrUpdate(plan Results) {
}

func (newState *Results) SyncEffectiveFieldsDuringRead(existingState Results) {
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
		"schema":    reflect.TypeOf(struct{}{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Results{}

// Equal implements basetypes.ObjectValuable.
func (o Results) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Results) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Results) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Results) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Results) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Results) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	CannedAcl types.String `tfsdk:"canned_acl" tf:"optional"`
	// S3 destination, e.g. `s3://my-bucket/some-prefix` Note that logs will be
	// delivered using cluster iam role, please make sure you set cluster iam
	// role and the role has write access to the destination. Please also note
	// that you cannot use AWS keys to deliver logs.
	Destination types.String `tfsdk:"destination" tf:""`
	// (Optional) Flag to enable server side encryption, `false` by default.
	EnableEncryption types.Bool `tfsdk:"enable_encryption" tf:"optional"`
	// (Optional) The encryption type, it could be `sse-s3` or `sse-kms`. It
	// will be used only when encryption is enabled and the default type is
	// `sse-s3`.
	EncryptionType types.String `tfsdk:"encryption_type" tf:"optional"`
	// S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`. Either region or
	// endpoint needs to be set. If both are set, endpoint will be used.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// (Optional) Kms key which will be used if encryption is enabled and
	// encryption type is set to `sse-kms`.
	KmsKey types.String `tfsdk:"kms_key" tf:"optional"`
	// S3 region, e.g. `us-west-2`. Either region or endpoint needs to be set.
	// If both are set, endpoint will be used.
	Region types.String `tfsdk:"region" tf:"optional"`
}

func (newState *S3StorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan S3StorageInfo) {
}

func (newState *S3StorageInfo) SyncEffectiveFieldsDuringRead(existingState S3StorageInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = S3StorageInfo{}

// Equal implements basetypes.ObjectValuable.
func (o S3StorageInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o S3StorageInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o S3StorageInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o S3StorageInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o S3StorageInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o S3StorageInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type SparkNode struct {
	// The private IP address of the host instance.
	HostPrivateIp types.String `tfsdk:"host_private_ip" tf:"optional"`
	// Globally unique identifier for the host instance from the cloud provider.
	InstanceId types.String `tfsdk:"instance_id" tf:"optional"`
	// Attributes specific to AWS for a Spark node.
	NodeAwsAttributes types.List `tfsdk:"node_aws_attributes" tf:"optional,object"`
	// Globally unique identifier for this node.
	NodeId types.String `tfsdk:"node_id" tf:"optional"`
	// Private IP address (typically a 10.x.x.x address) of the Spark node. Note
	// that this is different from the private IP address of the host instance.
	PrivateIp types.String `tfsdk:"private_ip" tf:"optional"`
	// Public DNS address of this node. This address can be used to access the
	// Spark JDBC server on the driver node. To communicate with the JDBC
	// server, traffic must be manually authorized by adding security group
	// rules to the "worker-unmanaged" security group via the AWS console.
	//
	// Actually it's the public DNS address of the host instance.
	PublicDns types.String `tfsdk:"public_dns" tf:"optional"`
	// The timestamp (in millisecond) when the Spark node is launched.
	//
	// The start_timestamp is set right before the container is being launched.
	// The timestamp when the container is placed on the ResourceManager, before
	// its launch and setup by the NodeDaemon. This timestamp is the same as the
	// creation timestamp in the database.
	StartTimestamp types.Int64 `tfsdk:"start_timestamp" tf:"optional"`
}

func (newState *SparkNode) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkNode) {
}

func (newState *SparkNode) SyncEffectiveFieldsDuringRead(existingState SparkNode) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SparkNode{}

// Equal implements basetypes.ObjectValuable.
func (o SparkNode) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SparkNode) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SparkNode) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SparkNode) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SparkNode) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SparkNode) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o SparkNode) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"host_private_ip": types.StringType,
			"instance_id":     types.StringType,
			"node_aws_attributes": basetypes.ListType{
				ElemType: SparkNodeAwsAttributes{}.Type(ctx),
			},
			"node_id":         types.StringType,
			"private_ip":      types.StringType,
			"public_dns":      types.StringType,
			"start_timestamp": types.Int64Type,
		},
	}
}

type SparkNodeAwsAttributes struct {
	// Whether this node is on an Amazon spot instance.
	IsSpot types.Bool `tfsdk:"is_spot" tf:"optional"`
}

func (newState *SparkNodeAwsAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkNodeAwsAttributes) {
}

func (newState *SparkNodeAwsAttributes) SyncEffectiveFieldsDuringRead(existingState SparkNodeAwsAttributes) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SparkNodeAwsAttributes{}

// Equal implements basetypes.ObjectValuable.
func (o SparkNodeAwsAttributes) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SparkNodeAwsAttributes) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SparkNodeAwsAttributes) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SparkNodeAwsAttributes) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SparkNodeAwsAttributes) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SparkNodeAwsAttributes) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Key types.String `tfsdk:"key" tf:"optional"`
	// A descriptive name for this Spark version, for example "Spark 2.1".
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *SparkVersion) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkVersion) {
}

func (newState *SparkVersion) SyncEffectiveFieldsDuringRead(existingState SparkVersion) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SparkVersion{}

// Equal implements basetypes.ObjectValuable.
func (o SparkVersion) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SparkVersion) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SparkVersion) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SparkVersion) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SparkVersion) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SparkVersion) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type StartCluster struct {
	// The cluster to be started.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

func (newState *StartCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartCluster) {
}

func (newState *StartCluster) SyncEffectiveFieldsDuringRead(existingState StartCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = StartCluster{}

// Equal implements basetypes.ObjectValuable.
func (o StartCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o StartCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o StartCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o StartCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o StartCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o StartCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = StartClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o StartClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o StartClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o StartClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o StartClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o StartClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o StartClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o StartClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code types.String `tfsdk:"code" tf:"optional"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters types.Map `tfsdk:"parameters" tf:"optional"`
	// type of the termination
	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *TerminationReason) SyncEffectiveFieldsDuringCreateOrUpdate(plan TerminationReason) {
}

func (newState *TerminationReason) SyncEffectiveFieldsDuringRead(existingState TerminationReason) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TerminationReason{}

// Equal implements basetypes.ObjectValuable.
func (o TerminationReason) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TerminationReason) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TerminationReason) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TerminationReason) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TerminationReason) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TerminationReason) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type UninstallLibraries struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// The libraries to uninstall.
	Libraries types.List `tfsdk:"libraries" tf:""`
}

func (newState *UninstallLibraries) SyncEffectiveFieldsDuringCreateOrUpdate(plan UninstallLibraries) {
}

func (newState *UninstallLibraries) SyncEffectiveFieldsDuringRead(existingState UninstallLibraries) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UninstallLibraries{}

// Equal implements basetypes.ObjectValuable.
func (o UninstallLibraries) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UninstallLibraries) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UninstallLibraries) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UninstallLibraries) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UninstallLibraries) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UninstallLibraries) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type UninstallLibrariesResponse struct {
}

func (newState *UninstallLibrariesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UninstallLibrariesResponse) {
}

func (newState *UninstallLibrariesResponse) SyncEffectiveFieldsDuringRead(existingState UninstallLibrariesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UninstallLibrariesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UninstallLibrariesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UnpinCluster struct {
	// <needs content added>
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

func (newState *UnpinCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpinCluster) {
}

func (newState *UnpinCluster) SyncEffectiveFieldsDuringRead(existingState UnpinCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UnpinCluster{}

// Equal implements basetypes.ObjectValuable.
func (o UnpinCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UnpinCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UnpinCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UnpinCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UnpinCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UnpinCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UnpinClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UnpinClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UnpinClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UnpinClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UnpinClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UnpinClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UnpinClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UnpinClusterResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateCluster struct {
	// The cluster to be updated.
	Cluster types.List `tfsdk:"cluster" tf:"optional,object"`
	// ID of the cluster.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// Specifies which fields of the cluster will be updated. This is required
	// in the POST request. The update mask should be supplied as a single
	// string. To specify multiple fields, separate them with commas (no
	// spaces). To delete a field from a cluster configuration, add it to the
	// `update_mask` string but omit it from the `cluster` object.
	UpdateMask types.String `tfsdk:"update_mask" tf:""`
}

func (newState *UpdateCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCluster) {
}

func (newState *UpdateCluster) SyncEffectiveFieldsDuringRead(existingState UpdateCluster) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateCluster{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateCluster) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateCluster) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateCluster) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateCluster) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateCluster) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateCluster) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster": basetypes.ListType{
				ElemType: UpdateClusterResource{}.Type(ctx),
			},
			"cluster_id":  types.StringType,
			"update_mask": types.StringType,
		},
	}
}

type UpdateClusterResource struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf" tf:"optional,object"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode types.String `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage types.List `tfsdk:"docker_image" tf:"optional,object"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
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
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine types.String `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf types.Map `tfsdk:"spark_conf" tf:"optional"`
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
	SparkEnvVars types.Map `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:"optional"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType types.List `tfsdk:"workload_type" tf:"optional,object"`
}

func (newState *UpdateClusterResource) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateClusterResource) {
}

func (newState *UpdateClusterResource) SyncEffectiveFieldsDuringRead(existingState UpdateClusterResource) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateClusterResource{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateClusterResource) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateClusterResource) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateClusterResource) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateClusterResource) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateClusterResource) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateClusterResource) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateClusterResource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscale": basetypes.ListType{
				ElemType: AutoScale{}.Type(ctx),
			},
			"autotermination_minutes": types.Int64Type,
			"aws_attributes": basetypes.ListType{
				ElemType: AwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: AzureAttributes{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: ClusterLogConf{}.Type(ctx),
			},
			"cluster_name": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"data_security_mode": types.StringType,
			"docker_image": basetypes.ListType{
				ElemType: DockerImage{}.Type(ctx),
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_elastic_disk":          types.BoolType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: GcpAttributes{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id": types.StringType,
			"node_type_id":     types.StringType,
			"num_workers":      types.Int64Type,
			"policy_id":        types.StringType,
			"runtime_engine":   types.StringType,
			"single_user_name": types.StringType,
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
			"workload_type": basetypes.ListType{
				ElemType: WorkloadType{}.Type(ctx),
			},
		},
	}
}

type UpdateClusterResponse struct {
}

func (newState *UpdateClusterResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateClusterResponse) {
}

func (newState *UpdateClusterResponse) SyncEffectiveFieldsDuringRead(existingState UpdateClusterResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateClusterResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateClusterResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateClusterResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateClusterResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateClusterResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateClusterResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateClusterResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type VolumesStorageInfo struct {
	// Unity Catalog Volumes file destination, e.g. `/Volumes/my-init.sh`
	Destination types.String `tfsdk:"destination" tf:""`
}

func (newState *VolumesStorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan VolumesStorageInfo) {
}

func (newState *VolumesStorageInfo) SyncEffectiveFieldsDuringRead(existingState VolumesStorageInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = VolumesStorageInfo{}

// Equal implements basetypes.ObjectValuable.
func (o VolumesStorageInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o VolumesStorageInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o VolumesStorageInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o VolumesStorageInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o VolumesStorageInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o VolumesStorageInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o VolumesStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

type WorkloadType struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	Clients types.List `tfsdk:"clients" tf:"object"`
}

func (newState *WorkloadType) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkloadType) {
}

func (newState *WorkloadType) SyncEffectiveFieldsDuringRead(existingState WorkloadType) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = WorkloadType{}

// Equal implements basetypes.ObjectValuable.
func (o WorkloadType) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o WorkloadType) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o WorkloadType) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o WorkloadType) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o WorkloadType) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o WorkloadType) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o WorkloadType) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clients": basetypes.ListType{
				ElemType: ClientsTypes{}.Type(ctx),
			},
		},
	}
}

type WorkspaceStorageInfo struct {
	// workspace files destination, e.g.
	// `/Users/user1@databricks.com/my-init.sh`
	Destination types.String `tfsdk:"destination" tf:""`
}

func (newState *WorkspaceStorageInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceStorageInfo) {
}

func (newState *WorkspaceStorageInfo) SyncEffectiveFieldsDuringRead(existingState WorkspaceStorageInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = WorkspaceStorageInfo{}

// Equal implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceStorageInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
		},
	}
}

// Availability type used for all subsequent nodes past the `first_on_demand`
// ones.
//
// Note: If `first_on_demand` is zero, this availability type will be used for
// the entire cluster.

// Availability type used for all subsequent nodes past the `first_on_demand`
// ones. Note: If `first_on_demand` is zero (which only happens on pool
// clusters), this availability type will be used for the entire cluster.

// Permission level

// Permission level

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.

// <needs content added>

// Data security mode decides what data governance model to use when accessing
// data from a cluster.
//
// * `NONE`: No security isolation for multiple users sharing the cluster. Data
// governance features are not available in this mode. * `SINGLE_USER`: A secure
// cluster that can only be exclusively used by a single user specified in
// `single_user_name`. Most programming languages, cluster features and data
// governance features are available in this mode. * `USER_ISOLATION`: A secure
// cluster that can be shared by multiple users. Cluster users are fully
// isolated so that they cannot see each other's data and credentials. Most data
// governance features are supported in this mode. But programming languages and
// cluster features might be limited.
//
// The following modes are deprecated starting with Databricks Runtime 15.0 and
// will be removed for future Databricks Runtime versions:
//
// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table ACL
// clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating from
// legacy Passthrough on high concurrency clusters. * `LEGACY_SINGLE_USER`: This
// mode is for users migrating from legacy Passthrough on standard clusters. *
// `LEGACY_SINGLE_USER_STANDARD`: This mode provides a way that doesn’t have
// UC nor passthrough enabled.

// The type of EBS volumes that will be launched with this cluster.

// The cause of a change in target size.

// This field determines whether the instance pool will contain preemptible VMs,
// on-demand VMs, or preemptible VMs with a fallback to on-demand VMs if the
// former is unavailable.

// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".

// The current status of the script

// Availability type used for the spot nodes.
//
// The default value is defined by
// InstancePoolConf.instancePoolDefaultAwsAvailability

// Shows the Availability type used for the spot nodes.
//
// The default value is defined by
// InstancePoolConf.instancePoolDefaultAzureAvailability

// Permission level

// Current state of the instance pool.

// The status of a library on a specific cluster.

// The direction to sort by.

// The sorting criteria. By default, clusters are sorted by 3 columns from
// highest to lowest precedence: cluster state, pinned or unpinned, then cluster
// name.

// A generic ordering enum for list-based queries.

// Determines the cluster's runtime engine, either standard or Photon.
//
// This field is not compatible with legacy `spark_version` values that contain
// `-photon-`. Remove `-photon-` from the `spark_version` and set
// `runtime_engine` to `PHOTON`.
//
// If left unspecified, the runtime engine defaults to standard unless the
// spark_version contains -photon-, in which case Photon will be used.

// Current state of the cluster.

// status code indicating why the cluster was terminated

// type of the termination
