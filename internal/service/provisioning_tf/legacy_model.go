// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package provisioning_tf

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

type AwsCredentials_SdkV2 struct {
	StsRole types.List `tfsdk:"sts_role"`
}

func (to *AwsCredentials_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsCredentials_SdkV2) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				// Recursively sync the fields of StsRole
				toStsRole.SyncFieldsDuringCreateOrUpdate(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (to *AwsCredentials_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AwsCredentials_SdkV2) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				toStsRole.SyncFieldsDuringRead(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (c AwsCredentials_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sts_role"] = attrs["sts_role"].SetOptional()
	attrs["sts_role"] = attrs["sts_role"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsCredentials_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(StsRole_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsCredentials_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsCredentials_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": o.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsCredentials_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": basetypes.ListType{
				ElemType: StsRole_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStsRole returns the value of the StsRole field in AwsCredentials_SdkV2 as
// a StsRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsCredentials_SdkV2) GetStsRole(ctx context.Context) (StsRole_SdkV2, bool) {
	var e StsRole_SdkV2
	if o.StsRole.IsNull() || o.StsRole.IsUnknown() {
		return e, false
	}
	var v []StsRole_SdkV2
	d := o.StsRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStsRole sets the value of the StsRole field in AwsCredentials_SdkV2.
func (o *AwsCredentials_SdkV2) SetStsRole(ctx context.Context, v StsRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sts_role"]
	o.StsRole = types.ListValueMust(t, vs)
}

type AwsDbManagedNetworkExtraInfo_SdkV2 struct {
	// This field is need to populate worker env for DB managed VPC. It is
	// likely only for resource tracking/deletion purpose.
	DhcpOptionsId types.String `tfsdk:"dhcp_options_id"`
	// This is the internal gateway which is different from the NAT gateway in
	// the NPIP VPC Infra. It is likely only for resource tracking/deletion
	// purpose.
	GatewayId types.String `tfsdk:"gateway_id"`
	// Security group which the Vault will control, ensuring that
	// worker_opened_ports are actually open.
	ManagedSecurityGroup types.String `tfsdk:"managed_security_group"`
	// Resources description for no public IP shard environment.
	NpipVpcInfra types.List `tfsdk:"npip_vpc_infra"`
	// Security group which is given to the user to manage without Databricks
	// interference.
	UnmanagedSecurityGroup types.String `tfsdk:"unmanaged_security_group"`
	// Contents of the secret key which gives ssh access to the workers.
	WorkerKeyContents types.String `tfsdk:"worker_key_contents"`
	// Name of the keypair in AWS which allows sshing into the workers.
	WorkerKeypairName types.String `tfsdk:"worker_keypair_name"`
}

func (to *AwsDbManagedNetworkExtraInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsDbManagedNetworkExtraInfo_SdkV2) {
	if !from.NpipVpcInfra.IsNull() && !from.NpipVpcInfra.IsUnknown() {
		if toNpipVpcInfra, ok := to.GetNpipVpcInfra(ctx); ok {
			if fromNpipVpcInfra, ok := from.GetNpipVpcInfra(ctx); ok {
				// Recursively sync the fields of NpipVpcInfra
				toNpipVpcInfra.SyncFieldsDuringCreateOrUpdate(ctx, fromNpipVpcInfra)
				to.SetNpipVpcInfra(ctx, toNpipVpcInfra)
			}
		}
	}
}

func (to *AwsDbManagedNetworkExtraInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AwsDbManagedNetworkExtraInfo_SdkV2) {
	if !from.NpipVpcInfra.IsNull() && !from.NpipVpcInfra.IsUnknown() {
		if toNpipVpcInfra, ok := to.GetNpipVpcInfra(ctx); ok {
			if fromNpipVpcInfra, ok := from.GetNpipVpcInfra(ctx); ok {
				toNpipVpcInfra.SyncFieldsDuringRead(ctx, fromNpipVpcInfra)
				to.SetNpipVpcInfra(ctx, toNpipVpcInfra)
			}
		}
	}
}

func (c AwsDbManagedNetworkExtraInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dhcp_options_id"] = attrs["dhcp_options_id"].SetOptional()
	attrs["gateway_id"] = attrs["gateway_id"].SetOptional()
	attrs["managed_security_group"] = attrs["managed_security_group"].SetOptional()
	attrs["npip_vpc_infra"] = attrs["npip_vpc_infra"].SetOptional()
	attrs["npip_vpc_infra"] = attrs["npip_vpc_infra"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["unmanaged_security_group"] = attrs["unmanaged_security_group"].SetOptional()
	attrs["worker_key_contents"] = attrs["worker_key_contents"].SetOptional()
	attrs["worker_keypair_name"] = attrs["worker_keypair_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsDbManagedNetworkExtraInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsDbManagedNetworkExtraInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"npip_vpc_infra": reflect.TypeOf(NpipVpcInfra_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsDbManagedNetworkExtraInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsDbManagedNetworkExtraInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dhcp_options_id":          o.DhcpOptionsId,
			"gateway_id":               o.GatewayId,
			"managed_security_group":   o.ManagedSecurityGroup,
			"npip_vpc_infra":           o.NpipVpcInfra,
			"unmanaged_security_group": o.UnmanagedSecurityGroup,
			"worker_key_contents":      o.WorkerKeyContents,
			"worker_keypair_name":      o.WorkerKeypairName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsDbManagedNetworkExtraInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dhcp_options_id":        types.StringType,
			"gateway_id":             types.StringType,
			"managed_security_group": types.StringType,
			"npip_vpc_infra": basetypes.ListType{
				ElemType: NpipVpcInfra_SdkV2{}.Type(ctx),
			},
			"unmanaged_security_group": types.StringType,
			"worker_key_contents":      types.StringType,
			"worker_keypair_name":      types.StringType,
		},
	}
}

// GetNpipVpcInfra returns the value of the NpipVpcInfra field in AwsDbManagedNetworkExtraInfo_SdkV2 as
// a NpipVpcInfra_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsDbManagedNetworkExtraInfo_SdkV2) GetNpipVpcInfra(ctx context.Context) (NpipVpcInfra_SdkV2, bool) {
	var e NpipVpcInfra_SdkV2
	if o.NpipVpcInfra.IsNull() || o.NpipVpcInfra.IsUnknown() {
		return e, false
	}
	var v []NpipVpcInfra_SdkV2
	d := o.NpipVpcInfra.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNpipVpcInfra sets the value of the NpipVpcInfra field in AwsDbManagedNetworkExtraInfo_SdkV2.
func (o *AwsDbManagedNetworkExtraInfo_SdkV2) SetNpipVpcInfra(ctx context.Context, v NpipVpcInfra_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["npip_vpc_infra"]
	o.NpipVpcInfra = types.ListValueMust(t, vs)
}

type AwsKeyInfo_SdkV2 struct {
	// The alias name of the KMS key.
	KeyAlias types.String `tfsdk:"key_alias"`
	// The the arn of the KMS key.
	KeyArn types.String `tfsdk:"key_arn"`
	// The region of the KMS key.
	KeyRegion types.String `tfsdk:"key_region"`
	// Indicates if the key should be used for cluster volumes. Can only be set
	// if the CMK can be used as a data plane key (use case storage)
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes"`
}

func (to *AwsKeyInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsKeyInfo_SdkV2) {
}

func (to *AwsKeyInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AwsKeyInfo_SdkV2) {
}

func (c AwsKeyInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key_alias"] = attrs["key_alias"].SetOptional()
	attrs["key_arn"] = attrs["key_arn"].SetRequired()
	attrs["key_region"] = attrs["key_region"].SetRequired()
	attrs["reuse_key_for_cluster_volumes"] = attrs["reuse_key_for_cluster_volumes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key_alias":                     o.KeyAlias,
			"key_arn":                       o.KeyArn,
			"key_region":                    o.KeyRegion,
			"reuse_key_for_cluster_volumes": o.ReuseKeyForClusterVolumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"key_region":                    types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type AwsNetworkInfo_SdkV2 struct {
	// Additional information for DB managed VPC, which is mainly used to
	// populate WorkerEnvironment.
	DbManagedVpcExtraInfo types.List `tfsdk:"db_managed_vpc_extra_info"`
	// The cloud-provided Security Group IDs that will be determine ingress and
	// egress rules for Cluster nodes.
	SecurityGroupIds types.List `tfsdk:"security_group_ids"`
	// The cloud-provided Subnet IDs that will be available to Clusters in
	// Workspaces using this Network.
	SubnetIds types.List `tfsdk:"subnet_ids"`
	// Details information of each individual subnet, including
	// availability_zone and address_space. This field is populated during
	// workspace creation and used for WorkerEnvironment.
	Subnets types.List `tfsdk:"subnets"`
	// CIDR that used for routing tables and security groups. Example:
	// 10.0.0.0/16. CIDR blocks can now be inferred from instance metadata
	// during setup so theoretically it is no longer necessary to populate the
	// `vpcAddressSpace` field. But there is a unknown bug which causes errors
	// when listing existing clusters and preventing customers from creating new
	// clusters under workspace `Compute` page. This field is populated during
	// workspace creation and used for WorkerEnvironment.
	VpcAddressSpace types.String `tfsdk:"vpc_address_space"`
	// The cloud-provided VPC ID.
	VpcId types.String `tfsdk:"vpc_id"`
}

func (to *AwsNetworkInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsNetworkInfo_SdkV2) {
	if !from.DbManagedVpcExtraInfo.IsNull() && !from.DbManagedVpcExtraInfo.IsUnknown() {
		if toDbManagedVpcExtraInfo, ok := to.GetDbManagedVpcExtraInfo(ctx); ok {
			if fromDbManagedVpcExtraInfo, ok := from.GetDbManagedVpcExtraInfo(ctx); ok {
				// Recursively sync the fields of DbManagedVpcExtraInfo
				toDbManagedVpcExtraInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromDbManagedVpcExtraInfo)
				to.SetDbManagedVpcExtraInfo(ctx, toDbManagedVpcExtraInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.Subnets.IsNull() && !from.Subnets.IsUnknown() && to.Subnets.IsNull() && len(from.Subnets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Subnets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Subnets = from.Subnets
	}
}

func (to *AwsNetworkInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AwsNetworkInfo_SdkV2) {
	if !from.DbManagedVpcExtraInfo.IsNull() && !from.DbManagedVpcExtraInfo.IsUnknown() {
		if toDbManagedVpcExtraInfo, ok := to.GetDbManagedVpcExtraInfo(ctx); ok {
			if fromDbManagedVpcExtraInfo, ok := from.GetDbManagedVpcExtraInfo(ctx); ok {
				toDbManagedVpcExtraInfo.SyncFieldsDuringRead(ctx, fromDbManagedVpcExtraInfo)
				to.SetDbManagedVpcExtraInfo(ctx, toDbManagedVpcExtraInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.Subnets.IsNull() && !from.Subnets.IsUnknown() && to.Subnets.IsNull() && len(from.Subnets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Subnets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Subnets = from.Subnets
	}
}

func (c AwsNetworkInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["db_managed_vpc_extra_info"] = attrs["db_managed_vpc_extra_info"].SetOptional()
	attrs["db_managed_vpc_extra_info"] = attrs["db_managed_vpc_extra_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["security_group_ids"] = attrs["security_group_ids"].SetOptional()
	attrs["subnet_ids"] = attrs["subnet_ids"].SetOptional()
	attrs["subnets"] = attrs["subnets"].SetOptional()
	attrs["vpc_address_space"] = attrs["vpc_address_space"].SetOptional()
	attrs["vpc_id"] = attrs["vpc_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsNetworkInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsNetworkInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"db_managed_vpc_extra_info": reflect.TypeOf(AwsDbManagedNetworkExtraInfo_SdkV2{}),
		"security_group_ids":        reflect.TypeOf(types.String{}),
		"subnet_ids":                reflect.TypeOf(types.String{}),
		"subnets":                   reflect.TypeOf(SubnetInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsNetworkInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsNetworkInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"db_managed_vpc_extra_info": o.DbManagedVpcExtraInfo,
			"security_group_ids":        o.SecurityGroupIds,
			"subnet_ids":                o.SubnetIds,
			"subnets":                   o.Subnets,
			"vpc_address_space":         o.VpcAddressSpace,
			"vpc_id":                    o.VpcId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsNetworkInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"db_managed_vpc_extra_info": basetypes.ListType{
				ElemType: AwsDbManagedNetworkExtraInfo_SdkV2{}.Type(ctx),
			},
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnets": basetypes.ListType{
				ElemType: SubnetInfo_SdkV2{}.Type(ctx),
			},
			"vpc_address_space": types.StringType,
			"vpc_id":            types.StringType,
		},
	}
}

// GetDbManagedVpcExtraInfo returns the value of the DbManagedVpcExtraInfo field in AwsNetworkInfo_SdkV2 as
// a AwsDbManagedNetworkExtraInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsNetworkInfo_SdkV2) GetDbManagedVpcExtraInfo(ctx context.Context) (AwsDbManagedNetworkExtraInfo_SdkV2, bool) {
	var e AwsDbManagedNetworkExtraInfo_SdkV2
	if o.DbManagedVpcExtraInfo.IsNull() || o.DbManagedVpcExtraInfo.IsUnknown() {
		return e, false
	}
	var v []AwsDbManagedNetworkExtraInfo_SdkV2
	d := o.DbManagedVpcExtraInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbManagedVpcExtraInfo sets the value of the DbManagedVpcExtraInfo field in AwsNetworkInfo_SdkV2.
func (o *AwsNetworkInfo_SdkV2) SetDbManagedVpcExtraInfo(ctx context.Context, v AwsDbManagedNetworkExtraInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["db_managed_vpc_extra_info"]
	o.DbManagedVpcExtraInfo = types.ListValueMust(t, vs)
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in AwsNetworkInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsNetworkInfo_SdkV2) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
	if o.SecurityGroupIds.IsNull() || o.SecurityGroupIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SecurityGroupIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in AwsNetworkInfo_SdkV2.
func (o *AwsNetworkInfo_SdkV2) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in AwsNetworkInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsNetworkInfo_SdkV2) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
	if o.SubnetIds.IsNull() || o.SubnetIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SubnetIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnetIds sets the value of the SubnetIds field in AwsNetworkInfo_SdkV2.
func (o *AwsNetworkInfo_SdkV2) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetSubnets returns the value of the Subnets field in AwsNetworkInfo_SdkV2 as
// a slice of SubnetInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsNetworkInfo_SdkV2) GetSubnets(ctx context.Context) ([]SubnetInfo_SdkV2, bool) {
	if o.Subnets.IsNull() || o.Subnets.IsUnknown() {
		return nil, false
	}
	var v []SubnetInfo_SdkV2
	d := o.Subnets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnets sets the value of the Subnets field in AwsNetworkInfo_SdkV2.
func (o *AwsNetworkInfo_SdkV2) SetSubnets(ctx context.Context, v []SubnetInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subnets = types.ListValueMust(t, vs)
}

type AzureKeyInfo_SdkV2 struct {
	// The Disk Encryption Set id that is used to represent the key info used
	// for Managed Disk BYOK use case
	DiskEncryptionSetId types.String `tfsdk:"disk_encryption_set_id"`
	// The structure to store key access credential This is set if the Managed
	// Identity is being used to access the Azure Key Vault key.
	KeyAccessConfiguration types.List `tfsdk:"key_access_configuration"`
	// The name of the key in KeyVault.
	KeyName types.String `tfsdk:"key_name"`
	// The base URI of the KeyVault.
	KeyVaultUri types.String `tfsdk:"key_vault_uri"`
	// The tenant id where the KeyVault lives.
	TenantId types.String `tfsdk:"tenant_id"`
	// The current key version.
	Version types.String `tfsdk:"version"`
}

func (to *AzureKeyInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureKeyInfo_SdkV2) {
	if !from.KeyAccessConfiguration.IsNull() && !from.KeyAccessConfiguration.IsUnknown() {
		if toKeyAccessConfiguration, ok := to.GetKeyAccessConfiguration(ctx); ok {
			if fromKeyAccessConfiguration, ok := from.GetKeyAccessConfiguration(ctx); ok {
				// Recursively sync the fields of KeyAccessConfiguration
				toKeyAccessConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromKeyAccessConfiguration)
				to.SetKeyAccessConfiguration(ctx, toKeyAccessConfiguration)
			}
		}
	}
}

func (to *AzureKeyInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AzureKeyInfo_SdkV2) {
	if !from.KeyAccessConfiguration.IsNull() && !from.KeyAccessConfiguration.IsUnknown() {
		if toKeyAccessConfiguration, ok := to.GetKeyAccessConfiguration(ctx); ok {
			if fromKeyAccessConfiguration, ok := from.GetKeyAccessConfiguration(ctx); ok {
				toKeyAccessConfiguration.SyncFieldsDuringRead(ctx, fromKeyAccessConfiguration)
				to.SetKeyAccessConfiguration(ctx, toKeyAccessConfiguration)
			}
		}
	}
}

func (c AzureKeyInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disk_encryption_set_id"] = attrs["disk_encryption_set_id"].SetOptional()
	attrs["key_access_configuration"] = attrs["key_access_configuration"].SetOptional()
	attrs["key_access_configuration"] = attrs["key_access_configuration"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["key_name"] = attrs["key_name"].SetOptional()
	attrs["key_vault_uri"] = attrs["key_vault_uri"].SetOptional()
	attrs["tenant_id"] = attrs["tenant_id"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"key_access_configuration": reflect.TypeOf(KeyAccessConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disk_encryption_set_id":   o.DiskEncryptionSetId,
			"key_access_configuration": o.KeyAccessConfiguration,
			"key_name":                 o.KeyName,
			"key_vault_uri":            o.KeyVaultUri,
			"tenant_id":                o.TenantId,
			"version":                  o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disk_encryption_set_id": types.StringType,
			"key_access_configuration": basetypes.ListType{
				ElemType: KeyAccessConfiguration_SdkV2{}.Type(ctx),
			},
			"key_name":      types.StringType,
			"key_vault_uri": types.StringType,
			"tenant_id":     types.StringType,
			"version":       types.StringType,
		},
	}
}

// GetKeyAccessConfiguration returns the value of the KeyAccessConfiguration field in AzureKeyInfo_SdkV2 as
// a KeyAccessConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AzureKeyInfo_SdkV2) GetKeyAccessConfiguration(ctx context.Context) (KeyAccessConfiguration_SdkV2, bool) {
	var e KeyAccessConfiguration_SdkV2
	if o.KeyAccessConfiguration.IsNull() || o.KeyAccessConfiguration.IsUnknown() {
		return e, false
	}
	var v []KeyAccessConfiguration_SdkV2
	d := o.KeyAccessConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetKeyAccessConfiguration sets the value of the KeyAccessConfiguration field in AzureKeyInfo_SdkV2.
func (o *AzureKeyInfo_SdkV2) SetKeyAccessConfiguration(ctx context.Context, v KeyAccessConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["key_access_configuration"]
	o.KeyAccessConfiguration = types.ListValueMust(t, vs)
}

type AzureWorkspaceInfo_SdkV2 struct {
	// Azure Resource Group name
	ResourceGroup types.String `tfsdk:"resource_group"`
	// Azure Subscription ID
	SubscriptionId types.String `tfsdk:"subscription_id"`
}

func (to *AzureWorkspaceInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureWorkspaceInfo_SdkV2) {
}

func (to *AzureWorkspaceInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AzureWorkspaceInfo_SdkV2) {
}

func (c AzureWorkspaceInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["resource_group"] = attrs["resource_group"].SetOptional()
	attrs["subscription_id"] = attrs["subscription_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureWorkspaceInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureWorkspaceInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureWorkspaceInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureWorkspaceInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resource_group":  o.ResourceGroup,
			"subscription_id": o.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureWorkspaceInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource_group":  types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

type CloudResourceContainer_SdkV2 struct {
	Gcp types.List `tfsdk:"gcp"`
}

func (to *CloudResourceContainer_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CloudResourceContainer_SdkV2) {
	if !from.Gcp.IsNull() && !from.Gcp.IsUnknown() {
		if toGcp, ok := to.GetGcp(ctx); ok {
			if fromGcp, ok := from.GetGcp(ctx); ok {
				// Recursively sync the fields of Gcp
				toGcp.SyncFieldsDuringCreateOrUpdate(ctx, fromGcp)
				to.SetGcp(ctx, toGcp)
			}
		}
	}
}

func (to *CloudResourceContainer_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CloudResourceContainer_SdkV2) {
	if !from.Gcp.IsNull() && !from.Gcp.IsUnknown() {
		if toGcp, ok := to.GetGcp(ctx); ok {
			if fromGcp, ok := from.GetGcp(ctx); ok {
				toGcp.SyncFieldsDuringRead(ctx, fromGcp)
				to.SetGcp(ctx, toGcp)
			}
		}
	}
}

func (c CloudResourceContainer_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp"] = attrs["gcp"].SetOptional()
	attrs["gcp"] = attrs["gcp"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloudResourceContainer.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CloudResourceContainer_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp": reflect.TypeOf(CustomerFacingGcpCloudResourceContainer_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudResourceContainer_SdkV2
// only implements ToObjectValue() and Type().
func (o CloudResourceContainer_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp": o.Gcp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloudResourceContainer_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp": basetypes.ListType{
				ElemType: CustomerFacingGcpCloudResourceContainer_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetGcp returns the value of the Gcp field in CloudResourceContainer_SdkV2 as
// a CustomerFacingGcpCloudResourceContainer_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CloudResourceContainer_SdkV2) GetGcp(ctx context.Context) (CustomerFacingGcpCloudResourceContainer_SdkV2, bool) {
	var e CustomerFacingGcpCloudResourceContainer_SdkV2
	if o.Gcp.IsNull() || o.Gcp.IsUnknown() {
		return e, false
	}
	var v []CustomerFacingGcpCloudResourceContainer_SdkV2
	d := o.Gcp.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcp sets the value of the Gcp field in CloudResourceContainer_SdkV2.
func (o *CloudResourceContainer_SdkV2) SetGcp(ctx context.Context, v CustomerFacingGcpCloudResourceContainer_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp"]
	o.Gcp = types.ListValueMust(t, vs)
}

type CreateAwsKeyInfo_SdkV2 struct {
	// The alias name of the KMS key.
	KeyAlias types.String `tfsdk:"key_alias"`
	// The the arn of the KMS key.
	KeyArn types.String `tfsdk:"key_arn"`
	// The region of the KMS key.
	KeyRegion types.String `tfsdk:"key_region"`
	// Indicates if the key should be used for cluster volumes. Can only be set
	// if the CMK can be used as a data plane key (use case storage)
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes"`
}

func (to *CreateAwsKeyInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAwsKeyInfo_SdkV2) {
}

func (to *CreateAwsKeyInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateAwsKeyInfo_SdkV2) {
}

func (c CreateAwsKeyInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key_alias"] = attrs["key_alias"].SetOptional()
	attrs["key_arn"] = attrs["key_arn"].SetRequired()
	attrs["key_region"] = attrs["key_region"].SetOptional()
	attrs["reuse_key_for_cluster_volumes"] = attrs["reuse_key_for_cluster_volumes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAwsKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAwsKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAwsKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateAwsKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key_alias":                     o.KeyAlias,
			"key_arn":                       o.KeyArn,
			"key_region":                    o.KeyRegion,
			"reuse_key_for_cluster_volumes": o.ReuseKeyForClusterVolumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAwsKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"key_region":                    types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type CreateCredentialAwsCredentials_SdkV2 struct {
	StsRole types.List `tfsdk:"sts_role"`
}

func (to *CreateCredentialAwsCredentials_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialAwsCredentials_SdkV2) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				// Recursively sync the fields of StsRole
				toStsRole.SyncFieldsDuringCreateOrUpdate(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (to *CreateCredentialAwsCredentials_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialAwsCredentials_SdkV2) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				toStsRole.SyncFieldsDuringRead(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (c CreateCredentialAwsCredentials_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sts_role"] = attrs["sts_role"].SetOptional()
	attrs["sts_role"] = attrs["sts_role"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialAwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialAwsCredentials_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(CreateCredentialStsRole_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialAwsCredentials_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialAwsCredentials_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": o.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialAwsCredentials_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": basetypes.ListType{
				ElemType: CreateCredentialStsRole_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStsRole returns the value of the StsRole field in CreateCredentialAwsCredentials_SdkV2 as
// a CreateCredentialStsRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialAwsCredentials_SdkV2) GetStsRole(ctx context.Context) (CreateCredentialStsRole_SdkV2, bool) {
	var e CreateCredentialStsRole_SdkV2
	if o.StsRole.IsNull() || o.StsRole.IsUnknown() {
		return e, false
	}
	var v []CreateCredentialStsRole_SdkV2
	d := o.StsRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStsRole sets the value of the StsRole field in CreateCredentialAwsCredentials_SdkV2.
func (o *CreateCredentialAwsCredentials_SdkV2) SetStsRole(ctx context.Context, v CreateCredentialStsRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sts_role"]
	o.StsRole = types.ListValueMust(t, vs)
}

type CreateCredentialRequest_SdkV2 struct {
	AwsCredentials types.List `tfsdk:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name"`
}

func (to *CreateCredentialRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialRequest_SdkV2) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				// Recursively sync the fields of AwsCredentials
				toAwsCredentials.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (to *CreateCredentialRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialRequest_SdkV2) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				toAwsCredentials.SyncFieldsDuringRead(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (c CreateCredentialRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_credentials"] = attrs["aws_credentials"].SetOptional()
	attrs["aws_credentials"] = attrs["aws_credentials"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["credentials_name"] = attrs["credentials_name"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_credentials": reflect.TypeOf(CreateCredentialAwsCredentials_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_credentials":  o.AwsCredentials,
			"credentials_name": o.CredentialsName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_credentials": basetypes.ListType{
				ElemType: CreateCredentialAwsCredentials_SdkV2{}.Type(ctx),
			},
			"credentials_name": types.StringType,
		},
	}
}

// GetAwsCredentials returns the value of the AwsCredentials field in CreateCredentialRequest_SdkV2 as
// a CreateCredentialAwsCredentials_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest_SdkV2) GetAwsCredentials(ctx context.Context) (CreateCredentialAwsCredentials_SdkV2, bool) {
	var e CreateCredentialAwsCredentials_SdkV2
	if o.AwsCredentials.IsNull() || o.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v []CreateCredentialAwsCredentials_SdkV2
	d := o.AwsCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsCredentials sets the value of the AwsCredentials field in CreateCredentialRequest_SdkV2.
func (o *CreateCredentialRequest_SdkV2) SetAwsCredentials(ctx context.Context, v CreateCredentialAwsCredentials_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_credentials"]
	o.AwsCredentials = types.ListValueMust(t, vs)
}

// * Use Amazon's STS service to assume a specified IAM role. The
// `longLivedProvider` is required to grant permission to assume `roleArn`. As
// an example, consider the vault creating the vpc in the customer account. The
// customer may provide her credentials as a role that we can assume. To create
// the VPC, the vault will use the "sts:AssumeRole" permission in its IAM role
// to assume the customer role. In this case, the vault's role is the long lived
// provider. @param roleArn The role to assume @param externalId An identifier
// that enables cross account role assumption @param longLivedProvider The
// credentials with which to assume the role
type CreateCredentialStsRole_SdkV2 struct {
	// Note: This must match the external_id on the parent object.
	//
	// TODO(j): Add validation to ensure this cannot be updated. If the user can
	// override the external_id, that defeats the purpose.
	ExternalId types.String `tfsdk:"external_id"`

	RoleArn types.String `tfsdk:"role_arn"`
}

func (to *CreateCredentialStsRole_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialStsRole_SdkV2) {
}

func (to *CreateCredentialStsRole_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialStsRole_SdkV2) {
}

func (c CreateCredentialStsRole_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["role_arn"] = attrs["role_arn"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialStsRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialStsRole_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialStsRole_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialStsRole_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
			"role_arn":    o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialStsRole_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
			"role_arn":    types.StringType,
		},
	}
}

type CreateCustomerManagedKeyRequest_SdkV2 struct {
	AwsKeyInfo types.List `tfsdk:"aws_key_info"`

	GcpKeyInfo types.List `tfsdk:"gcp_key_info"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases"`
}

func (to *CreateCustomerManagedKeyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCustomerManagedKeyRequest_SdkV2) {
	if !from.AwsKeyInfo.IsNull() && !from.AwsKeyInfo.IsUnknown() {
		if toAwsKeyInfo, ok := to.GetAwsKeyInfo(ctx); ok {
			if fromAwsKeyInfo, ok := from.GetAwsKeyInfo(ctx); ok {
				// Recursively sync the fields of AwsKeyInfo
				toAwsKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsKeyInfo)
				to.SetAwsKeyInfo(ctx, toAwsKeyInfo)
			}
		}
	}
	if !from.GcpKeyInfo.IsNull() && !from.GcpKeyInfo.IsUnknown() {
		if toGcpKeyInfo, ok := to.GetGcpKeyInfo(ctx); ok {
			if fromGcpKeyInfo, ok := from.GetGcpKeyInfo(ctx); ok {
				// Recursively sync the fields of GcpKeyInfo
				toGcpKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpKeyInfo)
				to.SetGcpKeyInfo(ctx, toGcpKeyInfo)
			}
		}
	}
	if !from.UseCases.IsNull() && !from.UseCases.IsUnknown() && to.UseCases.IsNull() && len(from.UseCases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UseCases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UseCases = from.UseCases
	}
}

func (to *CreateCustomerManagedKeyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCustomerManagedKeyRequest_SdkV2) {
	if !from.AwsKeyInfo.IsNull() && !from.AwsKeyInfo.IsUnknown() {
		if toAwsKeyInfo, ok := to.GetAwsKeyInfo(ctx); ok {
			if fromAwsKeyInfo, ok := from.GetAwsKeyInfo(ctx); ok {
				toAwsKeyInfo.SyncFieldsDuringRead(ctx, fromAwsKeyInfo)
				to.SetAwsKeyInfo(ctx, toAwsKeyInfo)
			}
		}
	}
	if !from.GcpKeyInfo.IsNull() && !from.GcpKeyInfo.IsUnknown() {
		if toGcpKeyInfo, ok := to.GetGcpKeyInfo(ctx); ok {
			if fromGcpKeyInfo, ok := from.GetGcpKeyInfo(ctx); ok {
				toGcpKeyInfo.SyncFieldsDuringRead(ctx, fromGcpKeyInfo)
				to.SetGcpKeyInfo(ctx, toGcpKeyInfo)
			}
		}
	}
	if !from.UseCases.IsNull() && !from.UseCases.IsUnknown() && to.UseCases.IsNull() && len(from.UseCases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UseCases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UseCases = from.UseCases
	}
}

func (c CreateCustomerManagedKeyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_key_info"] = attrs["aws_key_info"].SetOptional()
	attrs["aws_key_info"] = attrs["aws_key_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["gcp_key_info"] = attrs["gcp_key_info"].SetOptional()
	attrs["gcp_key_info"] = attrs["gcp_key_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["use_cases"] = attrs["use_cases"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomerManagedKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCustomerManagedKeyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info": reflect.TypeOf(CreateAwsKeyInfo_SdkV2{}),
		"gcp_key_info": reflect.TypeOf(CreateGcpKeyInfo_SdkV2{}),
		"use_cases":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomerManagedKeyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCustomerManagedKeyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_key_info": o.AwsKeyInfo,
			"gcp_key_info": o.GcpKeyInfo,
			"use_cases":    o.UseCases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCustomerManagedKeyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_key_info": basetypes.ListType{
				ElemType: CreateAwsKeyInfo_SdkV2{}.Type(ctx),
			},
			"gcp_key_info": basetypes.ListType{
				ElemType: CreateGcpKeyInfo_SdkV2{}.Type(ctx),
			},
			"use_cases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAwsKeyInfo returns the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest_SdkV2 as
// a CreateAwsKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest_SdkV2) GetAwsKeyInfo(ctx context.Context) (CreateAwsKeyInfo_SdkV2, bool) {
	var e CreateAwsKeyInfo_SdkV2
	if o.AwsKeyInfo.IsNull() || o.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v []CreateAwsKeyInfo_SdkV2
	d := o.AwsKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest_SdkV2.
func (o *CreateCustomerManagedKeyRequest_SdkV2) SetAwsKeyInfo(ctx context.Context, v CreateAwsKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_key_info"]
	o.AwsKeyInfo = types.ListValueMust(t, vs)
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest_SdkV2 as
// a CreateGcpKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest_SdkV2) GetGcpKeyInfo(ctx context.Context) (CreateGcpKeyInfo_SdkV2, bool) {
	var e CreateGcpKeyInfo_SdkV2
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v []CreateGcpKeyInfo_SdkV2
	d := o.GcpKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest_SdkV2.
func (o *CreateCustomerManagedKeyRequest_SdkV2) SetGcpKeyInfo(ctx context.Context, v CreateGcpKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_key_info"]
	o.GcpKeyInfo = types.ListValueMust(t, vs)
}

// GetUseCases returns the value of the UseCases field in CreateCustomerManagedKeyRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest_SdkV2) GetUseCases(ctx context.Context) ([]types.String, bool) {
	if o.UseCases.IsNull() || o.UseCases.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.UseCases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUseCases sets the value of the UseCases field in CreateCustomerManagedKeyRequest_SdkV2.
func (o *CreateCustomerManagedKeyRequest_SdkV2) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UseCases = types.ListValueMust(t, vs)
}

type CreateGcpKeyInfo_SdkV2 struct {
	// Globally unique service account email that has access to the KMS key. The
	// service account exists within the Databricks CP project.
	GcpServiceAccount types.List `tfsdk:"gcp_service_account"`
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId types.String `tfsdk:"kms_key_id"`
}

func (to *CreateGcpKeyInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateGcpKeyInfo_SdkV2) {
	if !from.GcpServiceAccount.IsNull() && !from.GcpServiceAccount.IsUnknown() {
		if toGcpServiceAccount, ok := to.GetGcpServiceAccount(ctx); ok {
			if fromGcpServiceAccount, ok := from.GetGcpServiceAccount(ctx); ok {
				// Recursively sync the fields of GcpServiceAccount
				toGcpServiceAccount.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpServiceAccount)
				to.SetGcpServiceAccount(ctx, toGcpServiceAccount)
			}
		}
	}
}

func (to *CreateGcpKeyInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateGcpKeyInfo_SdkV2) {
	if !from.GcpServiceAccount.IsNull() && !from.GcpServiceAccount.IsUnknown() {
		if toGcpServiceAccount, ok := to.GetGcpServiceAccount(ctx); ok {
			if fromGcpServiceAccount, ok := from.GetGcpServiceAccount(ctx); ok {
				toGcpServiceAccount.SyncFieldsDuringRead(ctx, fromGcpServiceAccount)
				to.SetGcpServiceAccount(ctx, toGcpServiceAccount)
			}
		}
	}
}

func (c CreateGcpKeyInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_service_account"] = attrs["gcp_service_account"].SetOptional()
	attrs["gcp_service_account"] = attrs["gcp_service_account"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["kms_key_id"] = attrs["kms_key_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateGcpKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateGcpKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_service_account": reflect.TypeOf(GcpServiceAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGcpKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateGcpKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_service_account": o.GcpServiceAccount,
			"kms_key_id":          o.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateGcpKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_service_account": basetypes.ListType{
				ElemType: GcpServiceAccount_SdkV2{}.Type(ctx),
			},
			"kms_key_id": types.StringType,
		},
	}
}

// GetGcpServiceAccount returns the value of the GcpServiceAccount field in CreateGcpKeyInfo_SdkV2 as
// a GcpServiceAccount_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateGcpKeyInfo_SdkV2) GetGcpServiceAccount(ctx context.Context) (GcpServiceAccount_SdkV2, bool) {
	var e GcpServiceAccount_SdkV2
	if o.GcpServiceAccount.IsNull() || o.GcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []GcpServiceAccount_SdkV2
	d := o.GcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpServiceAccount sets the value of the GcpServiceAccount field in CreateGcpKeyInfo_SdkV2.
func (o *CreateGcpKeyInfo_SdkV2) SetGcpServiceAccount(ctx context.Context, v GcpServiceAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_service_account"]
	o.GcpServiceAccount = types.ListValueMust(t, vs)
}

type CreateNetworkRequest_SdkV2 struct {
	GcpNetworkInfo types.List `tfsdk:"gcp_network_info"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds types.List `tfsdk:"security_group_ids"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds types.List `tfsdk:"subnet_ids"`

	VpcEndpoints types.List `tfsdk:"vpc_endpoints"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId types.String `tfsdk:"vpc_id"`
}

func (to *CreateNetworkRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkRequest_SdkV2) {
	if !from.GcpNetworkInfo.IsNull() && !from.GcpNetworkInfo.IsUnknown() {
		if toGcpNetworkInfo, ok := to.GetGcpNetworkInfo(ctx); ok {
			if fromGcpNetworkInfo, ok := from.GetGcpNetworkInfo(ctx); ok {
				// Recursively sync the fields of GcpNetworkInfo
				toGcpNetworkInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpNetworkInfo)
				to.SetGcpNetworkInfo(ctx, toGcpNetworkInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.VpcEndpoints.IsNull() && !from.VpcEndpoints.IsUnknown() {
		if toVpcEndpoints, ok := to.GetVpcEndpoints(ctx); ok {
			if fromVpcEndpoints, ok := from.GetVpcEndpoints(ctx); ok {
				// Recursively sync the fields of VpcEndpoints
				toVpcEndpoints.SyncFieldsDuringCreateOrUpdate(ctx, fromVpcEndpoints)
				to.SetVpcEndpoints(ctx, toVpcEndpoints)
			}
		}
	}
}

func (to *CreateNetworkRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkRequest_SdkV2) {
	if !from.GcpNetworkInfo.IsNull() && !from.GcpNetworkInfo.IsUnknown() {
		if toGcpNetworkInfo, ok := to.GetGcpNetworkInfo(ctx); ok {
			if fromGcpNetworkInfo, ok := from.GetGcpNetworkInfo(ctx); ok {
				toGcpNetworkInfo.SyncFieldsDuringRead(ctx, fromGcpNetworkInfo)
				to.SetGcpNetworkInfo(ctx, toGcpNetworkInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.VpcEndpoints.IsNull() && !from.VpcEndpoints.IsUnknown() {
		if toVpcEndpoints, ok := to.GetVpcEndpoints(ctx); ok {
			if fromVpcEndpoints, ok := from.GetVpcEndpoints(ctx); ok {
				toVpcEndpoints.SyncFieldsDuringRead(ctx, fromVpcEndpoints)
				to.SetVpcEndpoints(ctx, toVpcEndpoints)
			}
		}
	}
}

func (c CreateNetworkRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_network_info"] = attrs["gcp_network_info"].SetOptional()
	attrs["gcp_network_info"] = attrs["gcp_network_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["network_name"] = attrs["network_name"].SetOptional()
	attrs["security_group_ids"] = attrs["security_group_ids"].SetOptional()
	attrs["subnet_ids"] = attrs["subnet_ids"].SetOptional()
	attrs["vpc_endpoints"] = attrs["vpc_endpoints"].SetOptional()
	attrs["vpc_endpoints"] = attrs["vpc_endpoints"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["vpc_id"] = attrs["vpc_id"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateNetworkRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_network_info":   reflect.TypeOf(GcpNetworkInfo_SdkV2{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
		"subnet_ids":         reflect.TypeOf(types.String{}),
		"vpc_endpoints":      reflect.TypeOf(NetworkVpcEndpoints_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateNetworkRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_network_info":   o.GcpNetworkInfo,
			"network_name":       o.NetworkName,
			"security_group_ids": o.SecurityGroupIds,
			"subnet_ids":         o.SubnetIds,
			"vpc_endpoints":      o.VpcEndpoints,
			"vpc_id":             o.VpcId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateNetworkRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_network_info": basetypes.ListType{
				ElemType: GcpNetworkInfo_SdkV2{}.Type(ctx),
			},
			"network_name": types.StringType,
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"vpc_endpoints": basetypes.ListType{
				ElemType: NetworkVpcEndpoints_SdkV2{}.Type(ctx),
			},
			"vpc_id": types.StringType,
		},
	}
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in CreateNetworkRequest_SdkV2 as
// a GcpNetworkInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest_SdkV2) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo_SdkV2, bool) {
	var e GcpNetworkInfo_SdkV2
	if o.GcpNetworkInfo.IsNull() || o.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v []GcpNetworkInfo_SdkV2
	d := o.GcpNetworkInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in CreateNetworkRequest_SdkV2.
func (o *CreateNetworkRequest_SdkV2) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_network_info"]
	o.GcpNetworkInfo = types.ListValueMust(t, vs)
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in CreateNetworkRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest_SdkV2) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
	if o.SecurityGroupIds.IsNull() || o.SecurityGroupIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SecurityGroupIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in CreateNetworkRequest_SdkV2.
func (o *CreateNetworkRequest_SdkV2) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in CreateNetworkRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest_SdkV2) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
	if o.SubnetIds.IsNull() || o.SubnetIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SubnetIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnetIds sets the value of the SubnetIds field in CreateNetworkRequest_SdkV2.
func (o *CreateNetworkRequest_SdkV2) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in CreateNetworkRequest_SdkV2 as
// a NetworkVpcEndpoints_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest_SdkV2) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints_SdkV2, bool) {
	var e NetworkVpcEndpoints_SdkV2
	if o.VpcEndpoints.IsNull() || o.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v []NetworkVpcEndpoints_SdkV2
	d := o.VpcEndpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in CreateNetworkRequest_SdkV2.
func (o *CreateNetworkRequest_SdkV2) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["vpc_endpoints"]
	o.VpcEndpoints = types.ListValueMust(t, vs)
}

type CreatePrivateAccessSettingsRequest_SdkV2 struct {
	// The MWS API ID of VPC Endpoints that can access this workspace - only
	// filled if privateAccessLevel is ENDPOINT
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids"`
	// The level of isolation of a workspace attached to this settings object
	PrivateAccessLevel types.String `tfsdk:"private_access_level"`
	// The friendly user-facing name of the Private Access Settings (i.e. jake's
	// private access settings)
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name"`
	// Whether or not public traffic can enter this workspace. True for hybrid
	// workspaces, false otherwise.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled"`
	// The region in which this private access settings is valid
	Region types.String `tfsdk:"region"`
}

func (to *CreatePrivateAccessSettingsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePrivateAccessSettingsRequest_SdkV2) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (to *CreatePrivateAccessSettingsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePrivateAccessSettingsRequest_SdkV2) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (c CreatePrivateAccessSettingsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_vpc_endpoint_ids"] = attrs["allowed_vpc_endpoint_ids"].SetOptional()
	attrs["private_access_level"] = attrs["private_access_level"].SetOptional()
	attrs["private_access_settings_name"] = attrs["private_access_settings_name"].SetOptional()
	attrs["public_access_enabled"] = attrs["public_access_enabled"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateAccessSettingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePrivateAccessSettingsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateAccessSettingsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreatePrivateAccessSettingsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_vpc_endpoint_ids":     o.AllowedVpcEndpointIds,
			"private_access_level":         o.PrivateAccessLevel,
			"private_access_settings_name": o.PrivateAccessSettingsName,
			"public_access_enabled":        o.PublicAccessEnabled,
			"region":                       o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePrivateAccessSettingsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_vpc_endpoint_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"private_access_level":         types.StringType,
			"private_access_settings_name": types.StringType,
			"public_access_enabled":        types.BoolType,
			"region":                       types.StringType,
		},
	}
}

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in CreatePrivateAccessSettingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePrivateAccessSettingsRequest_SdkV2) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
	if o.AllowedVpcEndpointIds.IsNull() || o.AllowedVpcEndpointIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.AllowedVpcEndpointIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in CreatePrivateAccessSettingsRequest_SdkV2.
func (o *CreatePrivateAccessSettingsRequest_SdkV2) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type CreateStorageConfigurationRequest_SdkV2 struct {
	RootBucketInfo types.List `tfsdk:"root_bucket_info"`

	StorageConfigurationName types.String `tfsdk:"storage_configuration_name"`
}

func (to *CreateStorageConfigurationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateStorageConfigurationRequest_SdkV2) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				// Recursively sync the fields of RootBucketInfo
				toRootBucketInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (to *CreateStorageConfigurationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateStorageConfigurationRequest_SdkV2) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				toRootBucketInfo.SyncFieldsDuringRead(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (c CreateStorageConfigurationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["root_bucket_info"] = attrs["root_bucket_info"].SetOptional()
	attrs["root_bucket_info"] = attrs["root_bucket_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["storage_configuration_name"] = attrs["storage_configuration_name"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateStorageConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateStorageConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateStorageConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateStorageConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"root_bucket_info":           o.RootBucketInfo,
			"storage_configuration_name": o.StorageConfigurationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateStorageConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"root_bucket_info": basetypes.ListType{
				ElemType: RootBucketInfo_SdkV2{}.Type(ctx),
			},
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in CreateStorageConfigurationRequest_SdkV2 as
// a RootBucketInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageConfigurationRequest_SdkV2) GetRootBucketInfo(ctx context.Context) (RootBucketInfo_SdkV2, bool) {
	var e RootBucketInfo_SdkV2
	if o.RootBucketInfo.IsNull() || o.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v []RootBucketInfo_SdkV2
	d := o.RootBucketInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in CreateStorageConfigurationRequest_SdkV2.
func (o *CreateStorageConfigurationRequest_SdkV2) SetRootBucketInfo(ctx context.Context, v RootBucketInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["root_bucket_info"]
	o.RootBucketInfo = types.ListValueMust(t, vs)
}

type CreateVpcEndpointRequest_SdkV2 struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id"`
	// The cloud info of this vpc endpoint.
	GcpVpcEndpointInfo types.List `tfsdk:"gcp_vpc_endpoint_info"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name"`
}

func (to *CreateVpcEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVpcEndpointRequest_SdkV2) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				// Recursively sync the fields of GcpVpcEndpointInfo
				toGcpVpcEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (to *CreateVpcEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateVpcEndpointRequest_SdkV2) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				toGcpVpcEndpointInfo.SyncFieldsDuringRead(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (c CreateVpcEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_vpc_endpoint_id"] = attrs["aws_vpc_endpoint_id"].SetOptional()
	attrs["gcp_vpc_endpoint_info"] = attrs["gcp_vpc_endpoint_info"].SetOptional()
	attrs["gcp_vpc_endpoint_info"] = attrs["gcp_vpc_endpoint_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["region"] = attrs["region"].SetOptional()
	attrs["vpc_endpoint_name"] = attrs["vpc_endpoint_name"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVpcEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVpcEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateVpcEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_vpc_endpoint_id":   o.AwsVpcEndpointId,
			"gcp_vpc_endpoint_info": o.GcpVpcEndpointInfo,
			"region":                o.Region,
			"vpc_endpoint_name":     o.VpcEndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVpcEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_vpc_endpoint_id": types.StringType,
			"gcp_vpc_endpoint_info": basetypes.ListType{
				ElemType: GcpVpcEndpointInfo_SdkV2{}.Type(ctx),
			},
			"region":            types.StringType,
			"vpc_endpoint_name": types.StringType,
		},
	}
}

// GetGcpVpcEndpointInfo returns the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest_SdkV2 as
// a GcpVpcEndpointInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateVpcEndpointRequest_SdkV2) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo_SdkV2, bool) {
	var e GcpVpcEndpointInfo_SdkV2
	if o.GcpVpcEndpointInfo.IsNull() || o.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v []GcpVpcEndpointInfo_SdkV2
	d := o.GcpVpcEndpointInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest_SdkV2.
func (o *CreateVpcEndpointRequest_SdkV2) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_vpc_endpoint_info"]
	o.GcpVpcEndpointInfo = types.ListValueMust(t, vs)
}

type CreateWorkspaceRequest_SdkV2 struct {
	AwsRegion types.String `tfsdk:"aws_region"`
	// The cloud name. This field always has the value `gcp`.
	Cloud types.String `tfsdk:"cloud"`

	CloudResourceContainer types.List `tfsdk:"cloud_resource_container"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id"`

	CustomTags types.Map `tfsdk:"custom_tags"`

	DeploymentName types.String `tfsdk:"deployment_name"`

	GcpManagedNetworkConfig types.List `tfsdk:"gcp_managed_network_config"`

	GkeConfig types.List `tfsdk:"gke_config"`
	// Whether No Public IP is enabled for the workspace
	IsNoPublicIpEnabled types.Bool `tfsdk:"is_no_public_ip_enabled"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location types.String `tfsdk:"location"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id"`

	NetworkId types.String `tfsdk:"network_id"`

	PricingTier types.String `tfsdk:"pricing_tier"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types. Before
	// configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id"`
	// The human-readable name of the workspace.
	WorkspaceName types.String `tfsdk:"workspace_name"`
}

func (to *CreateWorkspaceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceRequest_SdkV2) {
	if !from.CloudResourceContainer.IsNull() && !from.CloudResourceContainer.IsUnknown() {
		if toCloudResourceContainer, ok := to.GetCloudResourceContainer(ctx); ok {
			if fromCloudResourceContainer, ok := from.GetCloudResourceContainer(ctx); ok {
				// Recursively sync the fields of CloudResourceContainer
				toCloudResourceContainer.SyncFieldsDuringCreateOrUpdate(ctx, fromCloudResourceContainer)
				to.SetCloudResourceContainer(ctx, toCloudResourceContainer)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				// Recursively sync the fields of GcpManagedNetworkConfig
				toGcpManagedNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
	if !from.GkeConfig.IsNull() && !from.GkeConfig.IsUnknown() {
		if toGkeConfig, ok := to.GetGkeConfig(ctx); ok {
			if fromGkeConfig, ok := from.GetGkeConfig(ctx); ok {
				// Recursively sync the fields of GkeConfig
				toGkeConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGkeConfig)
				to.SetGkeConfig(ctx, toGkeConfig)
			}
		}
	}
}

func (to *CreateWorkspaceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceRequest_SdkV2) {
	if !from.CloudResourceContainer.IsNull() && !from.CloudResourceContainer.IsUnknown() {
		if toCloudResourceContainer, ok := to.GetCloudResourceContainer(ctx); ok {
			if fromCloudResourceContainer, ok := from.GetCloudResourceContainer(ctx); ok {
				toCloudResourceContainer.SyncFieldsDuringRead(ctx, fromCloudResourceContainer)
				to.SetCloudResourceContainer(ctx, toCloudResourceContainer)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				toGcpManagedNetworkConfig.SyncFieldsDuringRead(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
	if !from.GkeConfig.IsNull() && !from.GkeConfig.IsUnknown() {
		if toGkeConfig, ok := to.GetGkeConfig(ctx); ok {
			if fromGkeConfig, ok := from.GetGkeConfig(ctx); ok {
				toGkeConfig.SyncFieldsDuringRead(ctx, fromGkeConfig)
				to.SetGkeConfig(ctx, toGkeConfig)
			}
		}
	}
}

func (c CreateWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_region"] = attrs["aws_region"].SetOptional()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["deployment_name"] = attrs["deployment_name"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["gke_config"] = attrs["gke_config"].SetOptional()
	attrs["gke_config"] = attrs["gke_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["is_no_public_ip_enabled"] = attrs["is_no_public_ip_enabled"].SetOptional()
	attrs["location"] = attrs["location"].SetOptional()
	attrs["managed_services_customer_managed_key_id"] = attrs["managed_services_customer_managed_key_id"].SetOptional()
	attrs["network_id"] = attrs["network_id"].SetOptional()
	attrs["pricing_tier"] = attrs["pricing_tier"].SetOptional()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetOptional()
	attrs["storage_customer_managed_key_id"] = attrs["storage_customer_managed_key_id"].SetOptional()
	attrs["workspace_name"] = attrs["workspace_name"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer_SdkV2{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig_SdkV2{}),
		"gke_config":                 reflect.TypeOf(GkeConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_region":                 o.AwsRegion,
			"cloud":                      o.Cloud,
			"cloud_resource_container":   o.CloudResourceContainer,
			"credentials_id":             o.CredentialsId,
			"custom_tags":                o.CustomTags,
			"deployment_name":            o.DeploymentName,
			"gcp_managed_network_config": o.GcpManagedNetworkConfig,
			"gke_config":                 o.GkeConfig,
			"is_no_public_ip_enabled":    o.IsNoPublicIpEnabled,
			"location":                   o.Location,
			"managed_services_customer_managed_key_id": o.ManagedServicesCustomerManagedKeyId,
			"network_id":                      o.NetworkId,
			"pricing_tier":                    o.PricingTier,
			"private_access_settings_id":      o.PrivateAccessSettingsId,
			"storage_configuration_id":        o.StorageConfigurationId,
			"storage_customer_managed_key_id": o.StorageCustomerManagedKeyId,
			"workspace_name":                  o.WorkspaceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_region": types.StringType,
			"cloud":      types.StringType,
			"cloud_resource_container": basetypes.ListType{
				ElemType: CloudResourceContainer_SdkV2{}.Type(ctx),
			},
			"credentials_id": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name": types.StringType,
			"gcp_managed_network_config": basetypes.ListType{
				ElemType: GcpManagedNetworkConfig_SdkV2{}.Type(ctx),
			},
			"gke_config": basetypes.ListType{
				ElemType: GkeConfig_SdkV2{}.Type(ctx),
			},
			"is_no_public_ip_enabled": types.BoolType,
			"location":                types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network_id":                      types.StringType,
			"pricing_tier":                    types.StringType,
			"private_access_settings_id":      types.StringType,
			"storage_configuration_id":        types.StringType,
			"storage_customer_managed_key_id": types.StringType,
			"workspace_name":                  types.StringType,
		},
	}
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in CreateWorkspaceRequest_SdkV2 as
// a CloudResourceContainer_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest_SdkV2) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer_SdkV2, bool) {
	var e CloudResourceContainer_SdkV2
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v []CloudResourceContainer_SdkV2
	d := o.CloudResourceContainer.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in CreateWorkspaceRequest_SdkV2.
func (o *CreateWorkspaceRequest_SdkV2) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloud_resource_container"]
	o.CloudResourceContainer = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in CreateWorkspaceRequest_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in CreateWorkspaceRequest_SdkV2.
func (o *CreateWorkspaceRequest_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest_SdkV2 as
// a GcpManagedNetworkConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest_SdkV2) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig_SdkV2, bool) {
	var e GcpManagedNetworkConfig_SdkV2
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v []GcpManagedNetworkConfig_SdkV2
	d := o.GcpManagedNetworkConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest_SdkV2.
func (o *CreateWorkspaceRequest_SdkV2) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_managed_network_config"]
	o.GcpManagedNetworkConfig = types.ListValueMust(t, vs)
}

// GetGkeConfig returns the value of the GkeConfig field in CreateWorkspaceRequest_SdkV2 as
// a GkeConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest_SdkV2) GetGkeConfig(ctx context.Context) (GkeConfig_SdkV2, bool) {
	var e GkeConfig_SdkV2
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v []GkeConfig_SdkV2
	d := o.GkeConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGkeConfig sets the value of the GkeConfig field in CreateWorkspaceRequest_SdkV2.
func (o *CreateWorkspaceRequest_SdkV2) SetGkeConfig(ctx context.Context, v GkeConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gke_config"]
	o.GkeConfig = types.ListValueMust(t, vs)
}

type Credential_SdkV2 struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id"`

	AwsCredentials types.List `tfsdk:"aws_credentials"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Databricks credential configuration ID.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name"`
}

func (to *Credential_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Credential_SdkV2) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				// Recursively sync the fields of AwsCredentials
				toAwsCredentials.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (to *Credential_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Credential_SdkV2) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				toAwsCredentials.SyncFieldsDuringRead(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (c Credential_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_credentials"] = attrs["aws_credentials"].SetOptional()
	attrs["aws_credentials"] = attrs["aws_credentials"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["credentials_name"] = attrs["credentials_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Credential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Credential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_credentials": reflect.TypeOf(AwsCredentials_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Credential_SdkV2
// only implements ToObjectValue() and Type().
func (o Credential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":       o.AccountId,
			"aws_credentials":  o.AwsCredentials,
			"creation_time":    o.CreationTime,
			"credentials_id":   o.CredentialsId,
			"credentials_name": o.CredentialsName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Credential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"aws_credentials": basetypes.ListType{
				ElemType: AwsCredentials_SdkV2{}.Type(ctx),
			},
			"creation_time":    types.Int64Type,
			"credentials_id":   types.StringType,
			"credentials_name": types.StringType,
		},
	}
}

// GetAwsCredentials returns the value of the AwsCredentials field in Credential_SdkV2 as
// a AwsCredentials_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Credential_SdkV2) GetAwsCredentials(ctx context.Context) (AwsCredentials_SdkV2, bool) {
	var e AwsCredentials_SdkV2
	if o.AwsCredentials.IsNull() || o.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v []AwsCredentials_SdkV2
	d := o.AwsCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsCredentials sets the value of the AwsCredentials field in Credential_SdkV2.
func (o *Credential_SdkV2) SetAwsCredentials(ctx context.Context, v AwsCredentials_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_credentials"]
	o.AwsCredentials = types.ListValueMust(t, vs)
}

type CustomerFacingGcpCloudResourceContainer_SdkV2 struct {
	ProjectId types.String `tfsdk:"project_id"`
}

func (to *CustomerFacingGcpCloudResourceContainer_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomerFacingGcpCloudResourceContainer_SdkV2) {
}

func (to *CustomerFacingGcpCloudResourceContainer_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CustomerFacingGcpCloudResourceContainer_SdkV2) {
}

func (c CustomerFacingGcpCloudResourceContainer_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomerFacingGcpCloudResourceContainer.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CustomerFacingGcpCloudResourceContainer_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerFacingGcpCloudResourceContainer_SdkV2
// only implements ToObjectValue() and Type().
func (o CustomerFacingGcpCloudResourceContainer_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_id": o.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomerFacingGcpCloudResourceContainer_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_id": types.StringType,
		},
	}
}

type CustomerManagedKey_SdkV2 struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId types.String `tfsdk:"account_id"`

	AwsKeyInfo types.List `tfsdk:"aws_key_info"`

	AzureKeyInfo types.List `tfsdk:"azure_key_info"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId types.String `tfsdk:"customer_managed_key_id"`

	GcpKeyInfo types.List `tfsdk:"gcp_key_info"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases"`
}

func (to *CustomerManagedKey_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomerManagedKey_SdkV2) {
	if !from.AwsKeyInfo.IsNull() && !from.AwsKeyInfo.IsUnknown() {
		if toAwsKeyInfo, ok := to.GetAwsKeyInfo(ctx); ok {
			if fromAwsKeyInfo, ok := from.GetAwsKeyInfo(ctx); ok {
				// Recursively sync the fields of AwsKeyInfo
				toAwsKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsKeyInfo)
				to.SetAwsKeyInfo(ctx, toAwsKeyInfo)
			}
		}
	}
	if !from.AzureKeyInfo.IsNull() && !from.AzureKeyInfo.IsUnknown() {
		if toAzureKeyInfo, ok := to.GetAzureKeyInfo(ctx); ok {
			if fromAzureKeyInfo, ok := from.GetAzureKeyInfo(ctx); ok {
				// Recursively sync the fields of AzureKeyInfo
				toAzureKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureKeyInfo)
				to.SetAzureKeyInfo(ctx, toAzureKeyInfo)
			}
		}
	}
	if !from.GcpKeyInfo.IsNull() && !from.GcpKeyInfo.IsUnknown() {
		if toGcpKeyInfo, ok := to.GetGcpKeyInfo(ctx); ok {
			if fromGcpKeyInfo, ok := from.GetGcpKeyInfo(ctx); ok {
				// Recursively sync the fields of GcpKeyInfo
				toGcpKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpKeyInfo)
				to.SetGcpKeyInfo(ctx, toGcpKeyInfo)
			}
		}
	}
	if !from.UseCases.IsNull() && !from.UseCases.IsUnknown() && to.UseCases.IsNull() && len(from.UseCases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UseCases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UseCases = from.UseCases
	}
}

func (to *CustomerManagedKey_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CustomerManagedKey_SdkV2) {
	if !from.AwsKeyInfo.IsNull() && !from.AwsKeyInfo.IsUnknown() {
		if toAwsKeyInfo, ok := to.GetAwsKeyInfo(ctx); ok {
			if fromAwsKeyInfo, ok := from.GetAwsKeyInfo(ctx); ok {
				toAwsKeyInfo.SyncFieldsDuringRead(ctx, fromAwsKeyInfo)
				to.SetAwsKeyInfo(ctx, toAwsKeyInfo)
			}
		}
	}
	if !from.AzureKeyInfo.IsNull() && !from.AzureKeyInfo.IsUnknown() {
		if toAzureKeyInfo, ok := to.GetAzureKeyInfo(ctx); ok {
			if fromAzureKeyInfo, ok := from.GetAzureKeyInfo(ctx); ok {
				toAzureKeyInfo.SyncFieldsDuringRead(ctx, fromAzureKeyInfo)
				to.SetAzureKeyInfo(ctx, toAzureKeyInfo)
			}
		}
	}
	if !from.GcpKeyInfo.IsNull() && !from.GcpKeyInfo.IsUnknown() {
		if toGcpKeyInfo, ok := to.GetGcpKeyInfo(ctx); ok {
			if fromGcpKeyInfo, ok := from.GetGcpKeyInfo(ctx); ok {
				toGcpKeyInfo.SyncFieldsDuringRead(ctx, fromGcpKeyInfo)
				to.SetGcpKeyInfo(ctx, toGcpKeyInfo)
			}
		}
	}
	if !from.UseCases.IsNull() && !from.UseCases.IsUnknown() && to.UseCases.IsNull() && len(from.UseCases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UseCases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UseCases = from.UseCases
	}
}

func (c CustomerManagedKey_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_key_info"] = attrs["aws_key_info"].SetOptional()
	attrs["aws_key_info"] = attrs["aws_key_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_key_info"] = attrs["azure_key_info"].SetOptional()
	attrs["azure_key_info"] = attrs["azure_key_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["customer_managed_key_id"] = attrs["customer_managed_key_id"].SetOptional()
	attrs["gcp_key_info"] = attrs["gcp_key_info"].SetOptional()
	attrs["gcp_key_info"] = attrs["gcp_key_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["use_cases"] = attrs["use_cases"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomerManagedKey.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CustomerManagedKey_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info":   reflect.TypeOf(AwsKeyInfo_SdkV2{}),
		"azure_key_info": reflect.TypeOf(AzureKeyInfo_SdkV2{}),
		"gcp_key_info":   reflect.TypeOf(GcpKeyInfo_SdkV2{}),
		"use_cases":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerManagedKey_SdkV2
// only implements ToObjectValue() and Type().
func (o CustomerManagedKey_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":              o.AccountId,
			"aws_key_info":            o.AwsKeyInfo,
			"azure_key_info":          o.AzureKeyInfo,
			"creation_time":           o.CreationTime,
			"customer_managed_key_id": o.CustomerManagedKeyId,
			"gcp_key_info":            o.GcpKeyInfo,
			"use_cases":               o.UseCases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomerManagedKey_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"aws_key_info": basetypes.ListType{
				ElemType: AwsKeyInfo_SdkV2{}.Type(ctx),
			},
			"azure_key_info": basetypes.ListType{
				ElemType: AzureKeyInfo_SdkV2{}.Type(ctx),
			},
			"creation_time":           types.Int64Type,
			"customer_managed_key_id": types.StringType,
			"gcp_key_info": basetypes.ListType{
				ElemType: GcpKeyInfo_SdkV2{}.Type(ctx),
			},
			"use_cases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAwsKeyInfo returns the value of the AwsKeyInfo field in CustomerManagedKey_SdkV2 as
// a AwsKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey_SdkV2) GetAwsKeyInfo(ctx context.Context) (AwsKeyInfo_SdkV2, bool) {
	var e AwsKeyInfo_SdkV2
	if o.AwsKeyInfo.IsNull() || o.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v []AwsKeyInfo_SdkV2
	d := o.AwsKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CustomerManagedKey_SdkV2.
func (o *CustomerManagedKey_SdkV2) SetAwsKeyInfo(ctx context.Context, v AwsKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_key_info"]
	o.AwsKeyInfo = types.ListValueMust(t, vs)
}

// GetAzureKeyInfo returns the value of the AzureKeyInfo field in CustomerManagedKey_SdkV2 as
// a AzureKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey_SdkV2) GetAzureKeyInfo(ctx context.Context) (AzureKeyInfo_SdkV2, bool) {
	var e AzureKeyInfo_SdkV2
	if o.AzureKeyInfo.IsNull() || o.AzureKeyInfo.IsUnknown() {
		return e, false
	}
	var v []AzureKeyInfo_SdkV2
	d := o.AzureKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureKeyInfo sets the value of the AzureKeyInfo field in CustomerManagedKey_SdkV2.
func (o *CustomerManagedKey_SdkV2) SetAzureKeyInfo(ctx context.Context, v AzureKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_key_info"]
	o.AzureKeyInfo = types.ListValueMust(t, vs)
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CustomerManagedKey_SdkV2 as
// a GcpKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey_SdkV2) GetGcpKeyInfo(ctx context.Context) (GcpKeyInfo_SdkV2, bool) {
	var e GcpKeyInfo_SdkV2
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v []GcpKeyInfo_SdkV2
	d := o.GcpKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CustomerManagedKey_SdkV2.
func (o *CustomerManagedKey_SdkV2) SetGcpKeyInfo(ctx context.Context, v GcpKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_key_info"]
	o.GcpKeyInfo = types.ListValueMust(t, vs)
}

// GetUseCases returns the value of the UseCases field in CustomerManagedKey_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey_SdkV2) GetUseCases(ctx context.Context) ([]types.String, bool) {
	if o.UseCases.IsNull() || o.UseCases.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.UseCases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUseCases sets the value of the UseCases field in CustomerManagedKey_SdkV2.
func (o *CustomerManagedKey_SdkV2) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UseCases = types.ListValueMust(t, vs)
}

type DeleteCredentialRequest_SdkV2 struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (to *DeleteCredentialRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCredentialRequest_SdkV2) {
}

func (to *DeleteCredentialRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCredentialRequest_SdkV2) {
}

func (c DeleteCredentialRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["credentials_id"] = attrs["credentials_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id": o.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id": types.StringType,
		},
	}
}

type DeleteEncryptionKeyRequest_SdkV2 struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

func (to *DeleteEncryptionKeyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEncryptionKeyRequest_SdkV2) {
}

func (to *DeleteEncryptionKeyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteEncryptionKeyRequest_SdkV2) {
}

func (c DeleteEncryptionKeyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["customer_managed_key_id"] = attrs["customer_managed_key_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEncryptionKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteEncryptionKeyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEncryptionKeyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteEncryptionKeyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": o.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteEncryptionKeyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_managed_key_id": types.StringType,
		},
	}
}

type DeleteNetworkRequest_SdkV2 struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

func (to *DeleteNetworkRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNetworkRequest_SdkV2) {
}

func (to *DeleteNetworkRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteNetworkRequest_SdkV2) {
}

func (c DeleteNetworkRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_id"] = attrs["network_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteNetworkRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

type DeletePrivateAccesRequest_SdkV2 struct {
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *DeletePrivateAccesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePrivateAccesRequest_SdkV2) {
}

func (to *DeletePrivateAccesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeletePrivateAccesRequest_SdkV2) {
}

func (c DeletePrivateAccesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePrivateAccesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePrivateAccesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateAccesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePrivateAccesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePrivateAccesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

type DeleteStorageRequest_SdkV2 struct {
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (to *DeleteStorageRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteStorageRequest_SdkV2) {
}

func (to *DeleteStorageRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteStorageRequest_SdkV2) {
}

func (c DeleteStorageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteStorageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteStorageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteStorageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteStorageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

type DeleteVpcEndpointRequest_SdkV2 struct {
	VpcEndpointId types.String `tfsdk:"-"`
}

func (to *DeleteVpcEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteVpcEndpointRequest_SdkV2) {
}

func (to *DeleteVpcEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteVpcEndpointRequest_SdkV2) {
}

func (c DeleteVpcEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["vpc_endpoint_id"] = attrs["vpc_endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteVpcEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVpcEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteVpcEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteVpcEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

type DeleteWorkspaceRequest_SdkV2 struct {
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceRequest_SdkV2) {
}

func (to *DeleteWorkspaceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceRequest_SdkV2) {
}

func (c DeleteWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type ExternalCustomerInfo_SdkV2 struct {
	// Email of the authoritative user.
	AuthoritativeUserEmail types.String `tfsdk:"authoritative_user_email"`
	// The authoritative user full name.
	AuthoritativeUserFullName types.String `tfsdk:"authoritative_user_full_name"`
	// The legal entity name for the external workspace
	CustomerName types.String `tfsdk:"customer_name"`

	OptOutExternalCustomerTosWorkflow types.Bool `tfsdk:"opt_out_external_customer_tos_workflow"`
	// The email of the authoritative user that signed the Terms of service.
	TosAcceptedByEmail types.String `tfsdk:"tos_accepted_by_email"`
	// The full name of the authoritative user that signed the Terms of service.
	TosAcceptedByFullName types.String `tfsdk:"tos_accepted_by_full_name"`
	// Indicates when the Terms of service was signed. None if it has not been
	// signed.
	TosAcceptedTimestamp types.Int64 `tfsdk:"tos_accepted_timestamp"`
}

func (to *ExternalCustomerInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalCustomerInfo_SdkV2) {
}

func (to *ExternalCustomerInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalCustomerInfo_SdkV2) {
}

func (c ExternalCustomerInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authoritative_user_email"] = attrs["authoritative_user_email"].SetOptional()
	attrs["authoritative_user_full_name"] = attrs["authoritative_user_full_name"].SetOptional()
	attrs["customer_name"] = attrs["customer_name"].SetOptional()
	attrs["opt_out_external_customer_tos_workflow"] = attrs["opt_out_external_customer_tos_workflow"].SetOptional()
	attrs["tos_accepted_by_email"] = attrs["tos_accepted_by_email"].SetOptional()
	attrs["tos_accepted_by_full_name"] = attrs["tos_accepted_by_full_name"].SetOptional()
	attrs["tos_accepted_timestamp"] = attrs["tos_accepted_timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalCustomerInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalCustomerInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalCustomerInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ExternalCustomerInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authoritative_user_email":               o.AuthoritativeUserEmail,
			"authoritative_user_full_name":           o.AuthoritativeUserFullName,
			"customer_name":                          o.CustomerName,
			"opt_out_external_customer_tos_workflow": o.OptOutExternalCustomerTosWorkflow,
			"tos_accepted_by_email":                  o.TosAcceptedByEmail,
			"tos_accepted_by_full_name":              o.TosAcceptedByFullName,
			"tos_accepted_timestamp":                 o.TosAcceptedTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalCustomerInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authoritative_user_email":               types.StringType,
			"authoritative_user_full_name":           types.StringType,
			"customer_name":                          types.StringType,
			"opt_out_external_customer_tos_workflow": types.BoolType,
			"tos_accepted_by_email":                  types.StringType,
			"tos_accepted_by_full_name":              types.StringType,
			"tos_accepted_timestamp":                 types.Int64Type,
		},
	}
}

// The shared network config for GCP workspace. This object has common network
// configurations that are network attributions of a workspace. DEPRECATED. Use
// GkeConfig instead.
type GcpCommonNetworkConfig_SdkV2 struct {
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	GkeClusterMasterIpRange types.String `tfsdk:"gke_cluster_master_ip_range"`
	// The type of network connectivity of the GKE cluster.
	GkeConnectivityType types.String `tfsdk:"gke_connectivity_type"`
}

func (to *GcpCommonNetworkConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpCommonNetworkConfig_SdkV2) {
}

func (to *GcpCommonNetworkConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GcpCommonNetworkConfig_SdkV2) {
}

func (c GcpCommonNetworkConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gke_cluster_master_ip_range"] = attrs["gke_cluster_master_ip_range"].SetOptional()
	attrs["gke_connectivity_type"] = attrs["gke_connectivity_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpCommonNetworkConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpCommonNetworkConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpCommonNetworkConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpCommonNetworkConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gke_cluster_master_ip_range": o.GkeClusterMasterIpRange,
			"gke_connectivity_type":       o.GkeConnectivityType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpCommonNetworkConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gke_cluster_master_ip_range": types.StringType,
			"gke_connectivity_type":       types.StringType,
		},
	}
}

type GcpKeyInfo_SdkV2 struct {
	// Globally unique service account email that has access to the KMS key. The
	// service account exists within the Databricks CP project.
	GcpServiceAccount types.List `tfsdk:"gcp_service_account"`
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId types.String `tfsdk:"kms_key_id"`
}

func (to *GcpKeyInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpKeyInfo_SdkV2) {
	if !from.GcpServiceAccount.IsNull() && !from.GcpServiceAccount.IsUnknown() {
		if toGcpServiceAccount, ok := to.GetGcpServiceAccount(ctx); ok {
			if fromGcpServiceAccount, ok := from.GetGcpServiceAccount(ctx); ok {
				// Recursively sync the fields of GcpServiceAccount
				toGcpServiceAccount.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpServiceAccount)
				to.SetGcpServiceAccount(ctx, toGcpServiceAccount)
			}
		}
	}
}

func (to *GcpKeyInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GcpKeyInfo_SdkV2) {
	if !from.GcpServiceAccount.IsNull() && !from.GcpServiceAccount.IsUnknown() {
		if toGcpServiceAccount, ok := to.GetGcpServiceAccount(ctx); ok {
			if fromGcpServiceAccount, ok := from.GetGcpServiceAccount(ctx); ok {
				toGcpServiceAccount.SyncFieldsDuringRead(ctx, fromGcpServiceAccount)
				to.SetGcpServiceAccount(ctx, toGcpServiceAccount)
			}
		}
	}
}

func (c GcpKeyInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_service_account"] = attrs["gcp_service_account"].SetOptional()
	attrs["gcp_service_account"] = attrs["gcp_service_account"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["kms_key_id"] = attrs["kms_key_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_service_account": reflect.TypeOf(GcpServiceAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_service_account": o.GcpServiceAccount,
			"kms_key_id":          o.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_service_account": basetypes.ListType{
				ElemType: GcpServiceAccount_SdkV2{}.Type(ctx),
			},
			"kms_key_id": types.StringType,
		},
	}
}

// GetGcpServiceAccount returns the value of the GcpServiceAccount field in GcpKeyInfo_SdkV2 as
// a GcpServiceAccount_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GcpKeyInfo_SdkV2) GetGcpServiceAccount(ctx context.Context) (GcpServiceAccount_SdkV2, bool) {
	var e GcpServiceAccount_SdkV2
	if o.GcpServiceAccount.IsNull() || o.GcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []GcpServiceAccount_SdkV2
	d := o.GcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpServiceAccount sets the value of the GcpServiceAccount field in GcpKeyInfo_SdkV2.
func (o *GcpKeyInfo_SdkV2) SetGcpServiceAccount(ctx context.Context, v GcpServiceAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_service_account"]
	o.GcpServiceAccount = types.ListValueMust(t, vs)
}

// The network configuration for the workspace.
type GcpManagedNetworkConfig_SdkV2 struct {
	// The IP range that will be used to allocate GKE cluster Pods from.
	GkeClusterPodIpRange types.String `tfsdk:"gke_cluster_pod_ip_range"`
	// The IP range that will be used to allocate GKE cluster Services from.
	GkeClusterServiceIpRange types.String `tfsdk:"gke_cluster_service_ip_range"`
	// The IP range which will be used to allocate GKE cluster nodes from. Note:
	// Pods, services and master IP range must be mutually exclusive.
	SubnetCidr types.String `tfsdk:"subnet_cidr"`
}

func (to *GcpManagedNetworkConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpManagedNetworkConfig_SdkV2) {
}

func (to *GcpManagedNetworkConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GcpManagedNetworkConfig_SdkV2) {
}

func (c GcpManagedNetworkConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gke_cluster_pod_ip_range"] = attrs["gke_cluster_pod_ip_range"].SetOptional()
	attrs["gke_cluster_service_ip_range"] = attrs["gke_cluster_service_ip_range"].SetOptional()
	attrs["subnet_cidr"] = attrs["subnet_cidr"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpManagedNetworkConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpManagedNetworkConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpManagedNetworkConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpManagedNetworkConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gke_cluster_pod_ip_range":     o.GkeClusterPodIpRange,
			"gke_cluster_service_ip_range": o.GkeClusterServiceIpRange,
			"subnet_cidr":                  o.SubnetCidr,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpManagedNetworkConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gke_cluster_pod_ip_range":     types.StringType,
			"gke_cluster_service_ip_range": types.StringType,
			"subnet_cidr":                  types.StringType,
		},
	}
}

type GcpNetworkInfo_SdkV2 struct {
	// The GCP project ID for network resources. This project is where the VPC
	// and subnet resides.
	NetworkProjectId types.String `tfsdk:"network_project_id"`
	// Name of the secondary range within the subnet that will be used by GKE as
	// Pod IP range. This is BYO VPC specific. DB VPC uses
	// network.getGcpManagedNetworkConfig.getGkeClusterPodIpRange
	PodIpRangeName types.String `tfsdk:"pod_ip_range_name"`
	// Name of the secondary range within the subnet that will be used by GKE as
	// Service IP range.
	ServiceIpRangeName types.String `tfsdk:"service_ip_range_name"`
	// The customer-provided Subnet ID that will be available to Clusters in
	// Workspaces using this Network.
	SubnetId types.String `tfsdk:"subnet_id"`

	SubnetRegion types.String `tfsdk:"subnet_region"`
	// The customer-provided VPC ID.
	VpcId types.String `tfsdk:"vpc_id"`
}

func (to *GcpNetworkInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpNetworkInfo_SdkV2) {
}

func (to *GcpNetworkInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GcpNetworkInfo_SdkV2) {
}

func (c GcpNetworkInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_project_id"] = attrs["network_project_id"].SetRequired()
	attrs["pod_ip_range_name"] = attrs["pod_ip_range_name"].SetRequired()
	attrs["service_ip_range_name"] = attrs["service_ip_range_name"].SetRequired()
	attrs["subnet_id"] = attrs["subnet_id"].SetRequired()
	attrs["subnet_region"] = attrs["subnet_region"].SetRequired()
	attrs["vpc_id"] = attrs["vpc_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpNetworkInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpNetworkInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpNetworkInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpNetworkInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_project_id":    o.NetworkProjectId,
			"pod_ip_range_name":     o.PodIpRangeName,
			"service_ip_range_name": o.ServiceIpRangeName,
			"subnet_id":             o.SubnetId,
			"subnet_region":         o.SubnetRegion,
			"vpc_id":                o.VpcId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpNetworkInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_project_id":    types.StringType,
			"pod_ip_range_name":     types.StringType,
			"service_ip_range_name": types.StringType,
			"subnet_id":             types.StringType,
			"subnet_region":         types.StringType,
			"vpc_id":                types.StringType,
		},
	}
}

type GcpServiceAccount_SdkV2 struct {
	ServiceAccountEmail types.String `tfsdk:"service_account_email"`
}

func (to *GcpServiceAccount_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpServiceAccount_SdkV2) {
}

func (to *GcpServiceAccount_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GcpServiceAccount_SdkV2) {
}

func (c GcpServiceAccount_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_account_email"] = attrs["service_account_email"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpServiceAccount.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpServiceAccount_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpServiceAccount_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpServiceAccount_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_account_email": o.ServiceAccountEmail,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpServiceAccount_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_account_email": types.StringType,
		},
	}
}

type GcpVpcEndpointInfo_SdkV2 struct {
	EndpointRegion types.String `tfsdk:"endpoint_region"`

	ProjectId types.String `tfsdk:"project_id"`

	PscConnectionId types.String `tfsdk:"psc_connection_id"`

	PscEndpointName types.String `tfsdk:"psc_endpoint_name"`

	ServiceAttachmentId types.String `tfsdk:"service_attachment_id"`
}

func (to *GcpVpcEndpointInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpVpcEndpointInfo_SdkV2) {
}

func (to *GcpVpcEndpointInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GcpVpcEndpointInfo_SdkV2) {
}

func (c GcpVpcEndpointInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_region"] = attrs["endpoint_region"].SetRequired()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["psc_connection_id"] = attrs["psc_connection_id"].SetOptional()
	attrs["psc_endpoint_name"] = attrs["psc_endpoint_name"].SetRequired()
	attrs["service_attachment_id"] = attrs["service_attachment_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpVpcEndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpVpcEndpointInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpVpcEndpointInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpVpcEndpointInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_region":       o.EndpointRegion,
			"project_id":            o.ProjectId,
			"psc_connection_id":     o.PscConnectionId,
			"psc_endpoint_name":     o.PscEndpointName,
			"service_attachment_id": o.ServiceAttachmentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpVpcEndpointInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_region":       types.StringType,
			"project_id":            types.StringType,
			"psc_connection_id":     types.StringType,
			"psc_endpoint_name":     types.StringType,
			"service_attachment_id": types.StringType,
		},
	}
}

type GetCredentialRequest_SdkV2 struct {
	// Credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (to *GetCredentialRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCredentialRequest_SdkV2) {
}

func (to *GetCredentialRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCredentialRequest_SdkV2) {
}

func (c GetCredentialRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["credentials_id"] = attrs["credentials_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id": o.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id": types.StringType,
		},
	}
}

type GetEncryptionKeyRequest_SdkV2 struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

func (to *GetEncryptionKeyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEncryptionKeyRequest_SdkV2) {
}

func (to *GetEncryptionKeyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEncryptionKeyRequest_SdkV2) {
}

func (c GetEncryptionKeyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["customer_managed_key_id"] = attrs["customer_managed_key_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEncryptionKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEncryptionKeyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEncryptionKeyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetEncryptionKeyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": o.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEncryptionKeyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_managed_key_id": types.StringType,
		},
	}
}

type GetNetworkRequest_SdkV2 struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

func (to *GetNetworkRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNetworkRequest_SdkV2) {
}

func (to *GetNetworkRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetNetworkRequest_SdkV2) {
}

func (c GetNetworkRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_id"] = attrs["network_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNetworkRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetNetworkRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNetworkRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

type GetPrivateAccesRequest_SdkV2 struct {
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *GetPrivateAccesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPrivateAccesRequest_SdkV2) {
}

func (to *GetPrivateAccesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPrivateAccesRequest_SdkV2) {
}

func (c GetPrivateAccesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPrivateAccesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPrivateAccesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateAccesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPrivateAccesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPrivateAccesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

type GetStorageRequest_SdkV2 struct {
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (to *GetStorageRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStorageRequest_SdkV2) {
}

func (to *GetStorageRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetStorageRequest_SdkV2) {
}

func (c GetStorageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStorageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStorageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetStorageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStorageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

type GetVpcEndpointRequest_SdkV2 struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId types.String `tfsdk:"-"`
}

func (to *GetVpcEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetVpcEndpointRequest_SdkV2) {
}

func (to *GetVpcEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetVpcEndpointRequest_SdkV2) {
}

func (c GetVpcEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["vpc_endpoint_id"] = attrs["vpc_endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetVpcEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVpcEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetVpcEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetVpcEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

type GetWorkspaceRequest_SdkV2 struct {
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceRequest_SdkV2) {
}

func (to *GetWorkspaceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceRequest_SdkV2) {
}

func (c GetWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// The configurations of the GKE cluster used by the GCP workspace.
type GkeConfig_SdkV2 struct {
	// The type of network connectivity of the GKE cluster.
	ConnectivityType types.String `tfsdk:"connectivity_type"`
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	MasterIpRange types.String `tfsdk:"master_ip_range"`
}

func (to *GkeConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GkeConfig_SdkV2) {
}

func (to *GkeConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GkeConfig_SdkV2) {
}

func (c GkeConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connectivity_type"] = attrs["connectivity_type"].SetOptional()
	attrs["master_ip_range"] = attrs["master_ip_range"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GkeConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GkeConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GkeConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o GkeConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connectivity_type": o.ConnectivityType,
			"master_ip_range":   o.MasterIpRange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GkeConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connectivity_type": types.StringType,
			"master_ip_range":   types.StringType,
		},
	}
}

// The credential ID that is used to access the key vault.
type KeyAccessConfiguration_SdkV2 struct {
	CredentialId types.String `tfsdk:"credential_id"`
}

func (to *KeyAccessConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KeyAccessConfiguration_SdkV2) {
}

func (to *KeyAccessConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KeyAccessConfiguration_SdkV2) {
}

func (c KeyAccessConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KeyAccessConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a KeyAccessConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KeyAccessConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o KeyAccessConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o KeyAccessConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.StringType,
		},
	}
}

type ListCredentialsRequest_SdkV2 struct {
}

func (to *ListCredentialsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCredentialsRequest_SdkV2) {
}

func (to *ListCredentialsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCredentialsRequest_SdkV2) {
}

func (c ListCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListEncryptionKeysRequest_SdkV2 struct {
}

func (to *ListEncryptionKeysRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEncryptionKeysRequest_SdkV2) {
}

func (to *ListEncryptionKeysRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEncryptionKeysRequest_SdkV2) {
}

func (c ListEncryptionKeysRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEncryptionKeysRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListEncryptionKeysRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEncryptionKeysRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListEncryptionKeysRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListEncryptionKeysRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNetworksRequest_SdkV2 struct {
}

func (to *ListNetworksRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworksRequest_SdkV2) {
}

func (to *ListNetworksRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListNetworksRequest_SdkV2) {
}

func (c ListNetworksRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworksRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNetworksRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworksRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNetworksRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworksRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListPrivateAccessRequest_SdkV2 struct {
}

func (to *ListPrivateAccessRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateAccessRequest_SdkV2) {
}

func (to *ListPrivateAccessRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPrivateAccessRequest_SdkV2) {
}

func (c ListPrivateAccessRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPrivateAccessRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateAccessRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPrivateAccessRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListPrivateAccessRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListStorageRequest_SdkV2 struct {
}

func (to *ListStorageRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListStorageRequest_SdkV2) {
}

func (to *ListStorageRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListStorageRequest_SdkV2) {
}

func (c ListStorageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListStorageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStorageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListStorageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListStorageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListVpcEndpointsRequest_SdkV2 struct {
}

func (to *ListVpcEndpointsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVpcEndpointsRequest_SdkV2) {
}

func (to *ListVpcEndpointsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListVpcEndpointsRequest_SdkV2) {
}

func (c ListVpcEndpointsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVpcEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVpcEndpointsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVpcEndpointsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListVpcEndpointsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListVpcEndpointsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListWorkspacesRequest_SdkV2 struct {
}

func (to *ListWorkspacesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspacesRequest_SdkV2) {
}

func (to *ListWorkspacesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspacesRequest_SdkV2) {
}

func (c ListWorkspacesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspacesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspacesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspacesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWorkspacesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspacesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Network_SdkV2 struct {
	// The Databricks account ID associated with this network configuration.
	AccountId types.String `tfsdk:"account_id"`

	AwsNetworkInfo types.List `tfsdk:"aws_network_info"`
	// Time in epoch milliseconds when the network was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Array of error messages about the network configuration.
	ErrorMessages types.List `tfsdk:"error_messages"`

	GcpNetworkInfo types.List `tfsdk:"gcp_network_info"`
	// The Databricks network configuration ID.
	NetworkId types.String `tfsdk:"network_id"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds types.List `tfsdk:"security_group_ids"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds types.List `tfsdk:"subnet_ids"`

	VpcEndpoints types.List `tfsdk:"vpc_endpoints"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId types.String `tfsdk:"vpc_id"`

	VpcStatus types.String `tfsdk:"vpc_status"`
	// Array of warning messages about the network configuration.
	WarningMessages types.List `tfsdk:"warning_messages"`
	// Workspace ID associated with this network configuration.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *Network_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Network_SdkV2) {
	if !from.AwsNetworkInfo.IsNull() && !from.AwsNetworkInfo.IsUnknown() {
		if toAwsNetworkInfo, ok := to.GetAwsNetworkInfo(ctx); ok {
			if fromAwsNetworkInfo, ok := from.GetAwsNetworkInfo(ctx); ok {
				// Recursively sync the fields of AwsNetworkInfo
				toAwsNetworkInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsNetworkInfo)
				to.SetAwsNetworkInfo(ctx, toAwsNetworkInfo)
			}
		}
	}
	if !from.ErrorMessages.IsNull() && !from.ErrorMessages.IsUnknown() && to.ErrorMessages.IsNull() && len(from.ErrorMessages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ErrorMessages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ErrorMessages = from.ErrorMessages
	}
	if !from.GcpNetworkInfo.IsNull() && !from.GcpNetworkInfo.IsUnknown() {
		if toGcpNetworkInfo, ok := to.GetGcpNetworkInfo(ctx); ok {
			if fromGcpNetworkInfo, ok := from.GetGcpNetworkInfo(ctx); ok {
				// Recursively sync the fields of GcpNetworkInfo
				toGcpNetworkInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpNetworkInfo)
				to.SetGcpNetworkInfo(ctx, toGcpNetworkInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.VpcEndpoints.IsNull() && !from.VpcEndpoints.IsUnknown() {
		if toVpcEndpoints, ok := to.GetVpcEndpoints(ctx); ok {
			if fromVpcEndpoints, ok := from.GetVpcEndpoints(ctx); ok {
				// Recursively sync the fields of VpcEndpoints
				toVpcEndpoints.SyncFieldsDuringCreateOrUpdate(ctx, fromVpcEndpoints)
				to.SetVpcEndpoints(ctx, toVpcEndpoints)
			}
		}
	}
	if !from.WarningMessages.IsNull() && !from.WarningMessages.IsUnknown() && to.WarningMessages.IsNull() && len(from.WarningMessages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WarningMessages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WarningMessages = from.WarningMessages
	}
}

func (to *Network_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Network_SdkV2) {
	if !from.AwsNetworkInfo.IsNull() && !from.AwsNetworkInfo.IsUnknown() {
		if toAwsNetworkInfo, ok := to.GetAwsNetworkInfo(ctx); ok {
			if fromAwsNetworkInfo, ok := from.GetAwsNetworkInfo(ctx); ok {
				toAwsNetworkInfo.SyncFieldsDuringRead(ctx, fromAwsNetworkInfo)
				to.SetAwsNetworkInfo(ctx, toAwsNetworkInfo)
			}
		}
	}
	if !from.ErrorMessages.IsNull() && !from.ErrorMessages.IsUnknown() && to.ErrorMessages.IsNull() && len(from.ErrorMessages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ErrorMessages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ErrorMessages = from.ErrorMessages
	}
	if !from.GcpNetworkInfo.IsNull() && !from.GcpNetworkInfo.IsUnknown() {
		if toGcpNetworkInfo, ok := to.GetGcpNetworkInfo(ctx); ok {
			if fromGcpNetworkInfo, ok := from.GetGcpNetworkInfo(ctx); ok {
				toGcpNetworkInfo.SyncFieldsDuringRead(ctx, fromGcpNetworkInfo)
				to.SetGcpNetworkInfo(ctx, toGcpNetworkInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.VpcEndpoints.IsNull() && !from.VpcEndpoints.IsUnknown() {
		if toVpcEndpoints, ok := to.GetVpcEndpoints(ctx); ok {
			if fromVpcEndpoints, ok := from.GetVpcEndpoints(ctx); ok {
				toVpcEndpoints.SyncFieldsDuringRead(ctx, fromVpcEndpoints)
				to.SetVpcEndpoints(ctx, toVpcEndpoints)
			}
		}
	}
	if !from.WarningMessages.IsNull() && !from.WarningMessages.IsUnknown() && to.WarningMessages.IsNull() && len(from.WarningMessages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WarningMessages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WarningMessages = from.WarningMessages
	}
}

func (c Network_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_network_info"] = attrs["aws_network_info"].SetOptional()
	attrs["aws_network_info"] = attrs["aws_network_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["error_messages"] = attrs["error_messages"].SetComputed()
	attrs["gcp_network_info"] = attrs["gcp_network_info"].SetOptional()
	attrs["gcp_network_info"] = attrs["gcp_network_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["network_id"] = attrs["network_id"].SetOptional()
	attrs["network_name"] = attrs["network_name"].SetOptional()
	attrs["security_group_ids"] = attrs["security_group_ids"].SetOptional()
	attrs["subnet_ids"] = attrs["subnet_ids"].SetOptional()
	attrs["vpc_endpoints"] = attrs["vpc_endpoints"].SetComputed()
	attrs["vpc_endpoints"] = attrs["vpc_endpoints"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["vpc_id"] = attrs["vpc_id"].SetOptional()
	attrs["vpc_status"] = attrs["vpc_status"].SetComputed()
	attrs["warning_messages"] = attrs["warning_messages"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Network.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Network_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_network_info":   reflect.TypeOf(AwsNetworkInfo_SdkV2{}),
		"error_messages":     reflect.TypeOf(NetworkHealth_SdkV2{}),
		"gcp_network_info":   reflect.TypeOf(GcpNetworkInfo_SdkV2{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
		"subnet_ids":         reflect.TypeOf(types.String{}),
		"vpc_endpoints":      reflect.TypeOf(NetworkVpcEndpoints_SdkV2{}),
		"warning_messages":   reflect.TypeOf(NetworkWarning_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Network_SdkV2
// only implements ToObjectValue() and Type().
func (o Network_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":         o.AccountId,
			"aws_network_info":   o.AwsNetworkInfo,
			"creation_time":      o.CreationTime,
			"error_messages":     o.ErrorMessages,
			"gcp_network_info":   o.GcpNetworkInfo,
			"network_id":         o.NetworkId,
			"network_name":       o.NetworkName,
			"security_group_ids": o.SecurityGroupIds,
			"subnet_ids":         o.SubnetIds,
			"vpc_endpoints":      o.VpcEndpoints,
			"vpc_id":             o.VpcId,
			"vpc_status":         o.VpcStatus,
			"warning_messages":   o.WarningMessages,
			"workspace_id":       o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Network_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"aws_network_info": basetypes.ListType{
				ElemType: AwsNetworkInfo_SdkV2{}.Type(ctx),
			},
			"creation_time": types.Int64Type,
			"error_messages": basetypes.ListType{
				ElemType: NetworkHealth_SdkV2{}.Type(ctx),
			},
			"gcp_network_info": basetypes.ListType{
				ElemType: GcpNetworkInfo_SdkV2{}.Type(ctx),
			},
			"network_id":   types.StringType,
			"network_name": types.StringType,
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"vpc_endpoints": basetypes.ListType{
				ElemType: NetworkVpcEndpoints_SdkV2{}.Type(ctx),
			},
			"vpc_id":     types.StringType,
			"vpc_status": types.StringType,
			"warning_messages": basetypes.ListType{
				ElemType: NetworkWarning_SdkV2{}.Type(ctx),
			},
			"workspace_id": types.Int64Type,
		},
	}
}

// GetAwsNetworkInfo returns the value of the AwsNetworkInfo field in Network_SdkV2 as
// a AwsNetworkInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetAwsNetworkInfo(ctx context.Context) (AwsNetworkInfo_SdkV2, bool) {
	var e AwsNetworkInfo_SdkV2
	if o.AwsNetworkInfo.IsNull() || o.AwsNetworkInfo.IsUnknown() {
		return e, false
	}
	var v []AwsNetworkInfo_SdkV2
	d := o.AwsNetworkInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsNetworkInfo sets the value of the AwsNetworkInfo field in Network_SdkV2.
func (o *Network_SdkV2) SetAwsNetworkInfo(ctx context.Context, v AwsNetworkInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_network_info"]
	o.AwsNetworkInfo = types.ListValueMust(t, vs)
}

// GetErrorMessages returns the value of the ErrorMessages field in Network_SdkV2 as
// a slice of NetworkHealth_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetErrorMessages(ctx context.Context) ([]NetworkHealth_SdkV2, bool) {
	if o.ErrorMessages.IsNull() || o.ErrorMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkHealth_SdkV2
	d := o.ErrorMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetErrorMessages sets the value of the ErrorMessages field in Network_SdkV2.
func (o *Network_SdkV2) SetErrorMessages(ctx context.Context, v []NetworkHealth_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ErrorMessages = types.ListValueMust(t, vs)
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in Network_SdkV2 as
// a GcpNetworkInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo_SdkV2, bool) {
	var e GcpNetworkInfo_SdkV2
	if o.GcpNetworkInfo.IsNull() || o.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v []GcpNetworkInfo_SdkV2
	d := o.GcpNetworkInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in Network_SdkV2.
func (o *Network_SdkV2) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_network_info"]
	o.GcpNetworkInfo = types.ListValueMust(t, vs)
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in Network_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
	if o.SecurityGroupIds.IsNull() || o.SecurityGroupIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SecurityGroupIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in Network_SdkV2.
func (o *Network_SdkV2) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in Network_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
	if o.SubnetIds.IsNull() || o.SubnetIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SubnetIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnetIds sets the value of the SubnetIds field in Network_SdkV2.
func (o *Network_SdkV2) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in Network_SdkV2 as
// a NetworkVpcEndpoints_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints_SdkV2, bool) {
	var e NetworkVpcEndpoints_SdkV2
	if o.VpcEndpoints.IsNull() || o.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v []NetworkVpcEndpoints_SdkV2
	d := o.VpcEndpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in Network_SdkV2.
func (o *Network_SdkV2) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["vpc_endpoints"]
	o.VpcEndpoints = types.ListValueMust(t, vs)
}

// GetWarningMessages returns the value of the WarningMessages field in Network_SdkV2 as
// a slice of NetworkWarning_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetWarningMessages(ctx context.Context) ([]NetworkWarning_SdkV2, bool) {
	if o.WarningMessages.IsNull() || o.WarningMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkWarning_SdkV2
	d := o.WarningMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarningMessages sets the value of the WarningMessages field in Network_SdkV2.
func (o *Network_SdkV2) SetWarningMessages(ctx context.Context, v []NetworkWarning_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["warning_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WarningMessages = types.ListValueMust(t, vs)
}

type NetworkHealth_SdkV2 struct {
	// Details of the error.
	ErrorMessage types.String `tfsdk:"error_message"`

	ErrorType types.String `tfsdk:"error_type"`
}

func (to *NetworkHealth_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkHealth_SdkV2) {
}

func (to *NetworkHealth_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NetworkHealth_SdkV2) {
}

func (c NetworkHealth_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["error_type"] = attrs["error_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkHealth.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkHealth_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkHealth_SdkV2
// only implements ToObjectValue() and Type().
func (o NetworkHealth_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_message": o.ErrorMessage,
			"error_type":    o.ErrorType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkHealth_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_message": types.StringType,
			"error_type":    types.StringType,
		},
	}
}

type NetworkVpcEndpoints_SdkV2 struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay types.List `tfsdk:"dataplane_relay"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi types.List `tfsdk:"rest_api"`
}

func (to *NetworkVpcEndpoints_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkVpcEndpoints_SdkV2) {
	if !from.DataplaneRelay.IsNull() && !from.DataplaneRelay.IsUnknown() && to.DataplaneRelay.IsNull() && len(from.DataplaneRelay.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataplaneRelay, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataplaneRelay = from.DataplaneRelay
	}
	if !from.RestApi.IsNull() && !from.RestApi.IsUnknown() && to.RestApi.IsNull() && len(from.RestApi.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RestApi, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RestApi = from.RestApi
	}
}

func (to *NetworkVpcEndpoints_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NetworkVpcEndpoints_SdkV2) {
	if !from.DataplaneRelay.IsNull() && !from.DataplaneRelay.IsUnknown() && to.DataplaneRelay.IsNull() && len(from.DataplaneRelay.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataplaneRelay, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataplaneRelay = from.DataplaneRelay
	}
	if !from.RestApi.IsNull() && !from.RestApi.IsUnknown() && to.RestApi.IsNull() && len(from.RestApi.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RestApi, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RestApi = from.RestApi
	}
}

func (c NetworkVpcEndpoints_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataplane_relay"] = attrs["dataplane_relay"].SetOptional()
	attrs["rest_api"] = attrs["rest_api"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkVpcEndpoints.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkVpcEndpoints_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataplane_relay": reflect.TypeOf(types.String{}),
		"rest_api":        reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkVpcEndpoints_SdkV2
// only implements ToObjectValue() and Type().
func (o NetworkVpcEndpoints_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataplane_relay": o.DataplaneRelay,
			"rest_api":        o.RestApi,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkVpcEndpoints_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataplane_relay": basetypes.ListType{
				ElemType: types.StringType,
			},
			"rest_api": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDataplaneRelay returns the value of the DataplaneRelay field in NetworkVpcEndpoints_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkVpcEndpoints_SdkV2) GetDataplaneRelay(ctx context.Context) ([]types.String, bool) {
	if o.DataplaneRelay.IsNull() || o.DataplaneRelay.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DataplaneRelay.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataplaneRelay sets the value of the DataplaneRelay field in NetworkVpcEndpoints_SdkV2.
func (o *NetworkVpcEndpoints_SdkV2) SetDataplaneRelay(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataplane_relay"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataplaneRelay = types.ListValueMust(t, vs)
}

// GetRestApi returns the value of the RestApi field in NetworkVpcEndpoints_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkVpcEndpoints_SdkV2) GetRestApi(ctx context.Context) ([]types.String, bool) {
	if o.RestApi.IsNull() || o.RestApi.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RestApi.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestApi sets the value of the RestApi field in NetworkVpcEndpoints_SdkV2.
func (o *NetworkVpcEndpoints_SdkV2) SetRestApi(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rest_api"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RestApi = types.ListValueMust(t, vs)
}

type NetworkWarning_SdkV2 struct {
	// Details of the warning.
	WarningMessage types.String `tfsdk:"warning_message"`

	WarningType types.String `tfsdk:"warning_type"`
}

func (to *NetworkWarning_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkWarning_SdkV2) {
}

func (to *NetworkWarning_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NetworkWarning_SdkV2) {
}

func (c NetworkWarning_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["warning_message"] = attrs["warning_message"].SetOptional()
	attrs["warning_type"] = attrs["warning_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkWarning.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkWarning_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkWarning_SdkV2
// only implements ToObjectValue() and Type().
func (o NetworkWarning_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warning_message": o.WarningMessage,
			"warning_type":    o.WarningType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkWarning_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warning_message": types.StringType,
			"warning_type":    types.StringType,
		},
	}
}

// Describes AWS resources allocations for NPIP shard environments. Used to
// track and delete resources during worker (and shard) environment deletion.
// Should only be used for MT NPIP shard environments currently.
type NpipVpcInfra_SdkV2 struct {
	// Elastic IP allocation id. Example: eipalloc-0df89abd3b5a548af
	NatEipAllocationId types.String `tfsdk:"nat_eip_allocation_id"`
	// NAT gateway id. Example: nat-0ae5b2f027fe7221a
	NatGatewayId types.String `tfsdk:"nat_gateway_id"`
	// Route table association id. Example: rtbassoc-089a9a9037542a912
	NatRouteTableAssociationId types.String `tfsdk:"nat_route_table_association_id"`
	// Route table id. Example: rtb-06118dc3003ee809b
	NatRouteTableId types.String `tfsdk:"nat_route_table_id"`
	// Subnet id. Example: subnet-0f6f001e243e00c10
	NatSubnetId types.String `tfsdk:"nat_subnet_id"`
	// VPC endpoint id. Example: vpce-08f210093b4e5ecb5
	NatVpcEndpointId types.String `tfsdk:"nat_vpc_endpoint_id"`
}

func (to *NpipVpcInfra_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NpipVpcInfra_SdkV2) {
}

func (to *NpipVpcInfra_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NpipVpcInfra_SdkV2) {
}

func (c NpipVpcInfra_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["nat_eip_allocation_id"] = attrs["nat_eip_allocation_id"].SetOptional()
	attrs["nat_gateway_id"] = attrs["nat_gateway_id"].SetOptional()
	attrs["nat_route_table_association_id"] = attrs["nat_route_table_association_id"].SetOptional()
	attrs["nat_route_table_id"] = attrs["nat_route_table_id"].SetOptional()
	attrs["nat_subnet_id"] = attrs["nat_subnet_id"].SetOptional()
	attrs["nat_vpc_endpoint_id"] = attrs["nat_vpc_endpoint_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NpipVpcInfra.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NpipVpcInfra_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NpipVpcInfra_SdkV2
// only implements ToObjectValue() and Type().
func (o NpipVpcInfra_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"nat_eip_allocation_id":          o.NatEipAllocationId,
			"nat_gateway_id":                 o.NatGatewayId,
			"nat_route_table_association_id": o.NatRouteTableAssociationId,
			"nat_route_table_id":             o.NatRouteTableId,
			"nat_subnet_id":                  o.NatSubnetId,
			"nat_vpc_endpoint_id":            o.NatVpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NpipVpcInfra_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"nat_eip_allocation_id":          types.StringType,
			"nat_gateway_id":                 types.StringType,
			"nat_route_table_association_id": types.StringType,
			"nat_route_table_id":             types.StringType,
			"nat_subnet_id":                  types.StringType,
			"nat_vpc_endpoint_id":            types.StringType,
		},
	}
}

// *
type PrivateAccessSettings_SdkV2 struct {
	// The MWS Account in which the Private Access Settings exists.
	AccountId types.String `tfsdk:"account_id"`
	// The MWS API ID of VPC Endpoints that can access this workspace - only
	// filled if privateAccessLevel is ENDPOINT
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids"`
	// The level of isolation of a workspace attached to this settings object
	PrivateAccessLevel types.String `tfsdk:"private_access_level"`
	// The ID in the MWS API of the Private Access Settings.
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// The friendly user-facing name of the Private Access Settings (i.e. jake's
	// private access settings)
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name"`
	// Whether or not public traffic can enter this workspace. True for hybrid
	// workspaces, false otherwise.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled"`
	// The region in which this private access settings is valid
	Region types.String `tfsdk:"region"`
}

func (to *PrivateAccessSettings_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateAccessSettings_SdkV2) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (to *PrivateAccessSettings_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PrivateAccessSettings_SdkV2) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (c PrivateAccessSettings_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["allowed_vpc_endpoint_ids"] = attrs["allowed_vpc_endpoint_ids"].SetOptional()
	attrs["private_access_level"] = attrs["private_access_level"].SetOptional()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetOptional()
	attrs["private_access_settings_name"] = attrs["private_access_settings_name"].SetOptional()
	attrs["public_access_enabled"] = attrs["public_access_enabled"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateAccessSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrivateAccessSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateAccessSettings_SdkV2
// only implements ToObjectValue() and Type().
func (o PrivateAccessSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                   o.AccountId,
			"allowed_vpc_endpoint_ids":     o.AllowedVpcEndpointIds,
			"private_access_level":         o.PrivateAccessLevel,
			"private_access_settings_id":   o.PrivateAccessSettingsId,
			"private_access_settings_name": o.PrivateAccessSettingsName,
			"public_access_enabled":        o.PublicAccessEnabled,
			"region":                       o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PrivateAccessSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"allowed_vpc_endpoint_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"private_access_level":         types.StringType,
			"private_access_settings_id":   types.StringType,
			"private_access_settings_name": types.StringType,
			"public_access_enabled":        types.BoolType,
			"region":                       types.StringType,
		},
	}
}

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in PrivateAccessSettings_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PrivateAccessSettings_SdkV2) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
	if o.AllowedVpcEndpointIds.IsNull() || o.AllowedVpcEndpointIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.AllowedVpcEndpointIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in PrivateAccessSettings_SdkV2.
func (o *PrivateAccessSettings_SdkV2) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type ReplacePrivateAccessSettingsRequest_SdkV2 struct {
	CustomerFacingPrivateAccessSettings types.List `tfsdk:"customer_facing_private_access_settings"`
	// The ID in the MWS API of the Private Access Settings.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *ReplacePrivateAccessSettingsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReplacePrivateAccessSettingsRequest_SdkV2) {
	if !from.CustomerFacingPrivateAccessSettings.IsNull() && !from.CustomerFacingPrivateAccessSettings.IsUnknown() {
		if toCustomerFacingPrivateAccessSettings, ok := to.GetCustomerFacingPrivateAccessSettings(ctx); ok {
			if fromCustomerFacingPrivateAccessSettings, ok := from.GetCustomerFacingPrivateAccessSettings(ctx); ok {
				// Recursively sync the fields of CustomerFacingPrivateAccessSettings
				toCustomerFacingPrivateAccessSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromCustomerFacingPrivateAccessSettings)
				to.SetCustomerFacingPrivateAccessSettings(ctx, toCustomerFacingPrivateAccessSettings)
			}
		}
	}
}

func (to *ReplacePrivateAccessSettingsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ReplacePrivateAccessSettingsRequest_SdkV2) {
	if !from.CustomerFacingPrivateAccessSettings.IsNull() && !from.CustomerFacingPrivateAccessSettings.IsUnknown() {
		if toCustomerFacingPrivateAccessSettings, ok := to.GetCustomerFacingPrivateAccessSettings(ctx); ok {
			if fromCustomerFacingPrivateAccessSettings, ok := from.GetCustomerFacingPrivateAccessSettings(ctx); ok {
				toCustomerFacingPrivateAccessSettings.SyncFieldsDuringRead(ctx, fromCustomerFacingPrivateAccessSettings)
				to.SetCustomerFacingPrivateAccessSettings(ctx, toCustomerFacingPrivateAccessSettings)
			}
		}
	}
}

func (c ReplacePrivateAccessSettingsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["customer_facing_private_access_settings"] = attrs["customer_facing_private_access_settings"].SetRequired()
	attrs["customer_facing_private_access_settings"] = attrs["customer_facing_private_access_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplacePrivateAccessSettingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplacePrivateAccessSettingsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"customer_facing_private_access_settings": reflect.TypeOf(PrivateAccessSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplacePrivateAccessSettingsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ReplacePrivateAccessSettingsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_facing_private_access_settings": o.CustomerFacingPrivateAccessSettings,
			"private_access_settings_id":              o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ReplacePrivateAccessSettingsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_facing_private_access_settings": basetypes.ListType{
				ElemType: PrivateAccessSettings_SdkV2{}.Type(ctx),
			},
			"private_access_settings_id": types.StringType,
		},
	}
}

// GetCustomerFacingPrivateAccessSettings returns the value of the CustomerFacingPrivateAccessSettings field in ReplacePrivateAccessSettingsRequest_SdkV2 as
// a PrivateAccessSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ReplacePrivateAccessSettingsRequest_SdkV2) GetCustomerFacingPrivateAccessSettings(ctx context.Context) (PrivateAccessSettings_SdkV2, bool) {
	var e PrivateAccessSettings_SdkV2
	if o.CustomerFacingPrivateAccessSettings.IsNull() || o.CustomerFacingPrivateAccessSettings.IsUnknown() {
		return e, false
	}
	var v []PrivateAccessSettings_SdkV2
	d := o.CustomerFacingPrivateAccessSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCustomerFacingPrivateAccessSettings sets the value of the CustomerFacingPrivateAccessSettings field in ReplacePrivateAccessSettingsRequest_SdkV2.
func (o *ReplacePrivateAccessSettingsRequest_SdkV2) SetCustomerFacingPrivateAccessSettings(ctx context.Context, v PrivateAccessSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["customer_facing_private_access_settings"]
	o.CustomerFacingPrivateAccessSettings = types.ListValueMust(t, vs)
}

type RootBucketInfo_SdkV2 struct {
	// Name of the bucket
	BucketName types.String `tfsdk:"bucket_name"`
}

func (to *RootBucketInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RootBucketInfo_SdkV2) {
}

func (to *RootBucketInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RootBucketInfo_SdkV2) {
}

func (c RootBucketInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bucket_name"] = attrs["bucket_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RootBucketInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RootBucketInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RootBucketInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RootBucketInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bucket_name": o.BucketName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RootBucketInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bucket_name": types.StringType,
		},
	}
}

type StorageConfiguration_SdkV2 struct {
	AccountId types.String `tfsdk:"account_id"`

	CreationTime types.Int64 `tfsdk:"creation_time"`
	// The IAM role that is used to access the workspace catalog which is
	// created during workspace creation for UC by Default. If a storage
	// configuration that has this field populated is used to create a
	// workspace, then a workspace catalog is created together with the
	// workspace. The workspace catalog shares the root bucket with internal
	// workspace storage (including DBFS root) but uses a dedicated bucket path
	// prefix.
	RoleArn types.String `tfsdk:"role_arn"`

	RootBucketInfo types.List `tfsdk:"root_bucket_info"`

	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`

	StorageConfigurationName types.String `tfsdk:"storage_configuration_name"`
}

func (to *StorageConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StorageConfiguration_SdkV2) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				// Recursively sync the fields of RootBucketInfo
				toRootBucketInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (to *StorageConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StorageConfiguration_SdkV2) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				toRootBucketInfo.SyncFieldsDuringRead(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (c StorageConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["role_arn"] = attrs["role_arn"].SetOptional()
	attrs["root_bucket_info"] = attrs["root_bucket_info"].SetOptional()
	attrs["root_bucket_info"] = attrs["root_bucket_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetOptional()
	attrs["storage_configuration_name"] = attrs["storage_configuration_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StorageConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StorageConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StorageConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o StorageConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                 o.AccountId,
			"creation_time":              o.CreationTime,
			"role_arn":                   o.RoleArn,
			"root_bucket_info":           o.RootBucketInfo,
			"storage_configuration_id":   o.StorageConfigurationId,
			"storage_configuration_name": o.StorageConfigurationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StorageConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":    types.StringType,
			"creation_time": types.Int64Type,
			"role_arn":      types.StringType,
			"root_bucket_info": basetypes.ListType{
				ElemType: RootBucketInfo_SdkV2{}.Type(ctx),
			},
			"storage_configuration_id":   types.StringType,
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in StorageConfiguration_SdkV2 as
// a RootBucketInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageConfiguration_SdkV2) GetRootBucketInfo(ctx context.Context) (RootBucketInfo_SdkV2, bool) {
	var e RootBucketInfo_SdkV2
	if o.RootBucketInfo.IsNull() || o.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v []RootBucketInfo_SdkV2
	d := o.RootBucketInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in StorageConfiguration_SdkV2.
func (o *StorageConfiguration_SdkV2) SetRootBucketInfo(ctx context.Context, v RootBucketInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["root_bucket_info"]
	o.RootBucketInfo = types.ListValueMust(t, vs)
}

// * Use Amazon's STS service to assume a specified IAM role. The
// `longLivedProvider` is required to grant permission to assume `roleArn`. As
// an example, consider the vault creating the vpc in the customer account. The
// customer may provide her credentials as a role that we can assume. To create
// the VPC, the vault will use the "sts:AssumeRole" permission in its IAM role
// to assume the customer role. In this case, the vault's role is the long lived
// provider. @param roleArn The role to assume @param externalId An identifier
// that enables cross account role assumption @param longLivedProvider The
// credentials with which to assume the role
type StsRole_SdkV2 struct {
	// Note: This must match the external_id on the parent object.
	//
	// TODO(j): Add validation to ensure this cannot be updated. If the user can
	// override the external_id, that defeats the purpose.
	ExternalId types.String `tfsdk:"external_id"`

	RoleArn types.String `tfsdk:"role_arn"`
}

func (to *StsRole_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StsRole_SdkV2) {
}

func (to *StsRole_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StsRole_SdkV2) {
}

func (c StsRole_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["role_arn"] = attrs["role_arn"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StsRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StsRole_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StsRole_SdkV2
// only implements ToObjectValue() and Type().
func (o StsRole_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
			"role_arn":    o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StsRole_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
			"role_arn":    types.StringType,
		},
	}
}

// Describes a single subnet, which is associated with a particular AWS AZ and a
// particular address space which is a subset of the overall vpc_address_space.
type SubnetInfo_SdkV2 struct {
	// Example: us-west-2a
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	// Example: 10.0.0.0/17.
	SubnetAddressSpace types.String `tfsdk:"subnet_address_space"`

	SubnetId types.String `tfsdk:"subnet_id"`
}

func (to *SubnetInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SubnetInfo_SdkV2) {
}

func (to *SubnetInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SubnetInfo_SdkV2) {
}

func (c SubnetInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["availability_zone"] = attrs["availability_zone"].SetOptional()
	attrs["subnet_address_space"] = attrs["subnet_address_space"].SetOptional()
	attrs["subnet_id"] = attrs["subnet_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubnetInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubnetInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubnetInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o SubnetInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability_zone":    o.AvailabilityZone,
			"subnet_address_space": o.SubnetAddressSpace,
			"subnet_id":            o.SubnetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubnetInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability_zone":    types.StringType,
			"subnet_address_space": types.StringType,
			"subnet_id":            types.StringType,
		},
	}
}

type UpdateWorkspaceRequest_SdkV2 struct {
	CustomerFacingWorkspace types.List `tfsdk:"customer_facing_workspace"`
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
	UpdateMask types.String `tfsdk:"-"`
	// A unique integer ID for the workspace
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *UpdateWorkspaceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceRequest_SdkV2) {
	if !from.CustomerFacingWorkspace.IsNull() && !from.CustomerFacingWorkspace.IsUnknown() {
		if toCustomerFacingWorkspace, ok := to.GetCustomerFacingWorkspace(ctx); ok {
			if fromCustomerFacingWorkspace, ok := from.GetCustomerFacingWorkspace(ctx); ok {
				// Recursively sync the fields of CustomerFacingWorkspace
				toCustomerFacingWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromCustomerFacingWorkspace)
				to.SetCustomerFacingWorkspace(ctx, toCustomerFacingWorkspace)
			}
		}
	}
}

func (to *UpdateWorkspaceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceRequest_SdkV2) {
	if !from.CustomerFacingWorkspace.IsNull() && !from.CustomerFacingWorkspace.IsUnknown() {
		if toCustomerFacingWorkspace, ok := to.GetCustomerFacingWorkspace(ctx); ok {
			if fromCustomerFacingWorkspace, ok := from.GetCustomerFacingWorkspace(ctx); ok {
				toCustomerFacingWorkspace.SyncFieldsDuringRead(ctx, fromCustomerFacingWorkspace)
				to.SetCustomerFacingWorkspace(ctx, toCustomerFacingWorkspace)
			}
		}
	}
}

func (c UpdateWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["customer_facing_workspace"] = attrs["customer_facing_workspace"].SetRequired()
	attrs["customer_facing_workspace"] = attrs["customer_facing_workspace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"customer_facing_workspace": reflect.TypeOf(Workspace_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_facing_workspace": o.CustomerFacingWorkspace,
			"update_mask":               o.UpdateMask,
			"workspace_id":              o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_facing_workspace": basetypes.ListType{
				ElemType: Workspace_SdkV2{}.Type(ctx),
			},
			"update_mask":  types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// GetCustomerFacingWorkspace returns the value of the CustomerFacingWorkspace field in UpdateWorkspaceRequest_SdkV2 as
// a Workspace_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceRequest_SdkV2) GetCustomerFacingWorkspace(ctx context.Context) (Workspace_SdkV2, bool) {
	var e Workspace_SdkV2
	if o.CustomerFacingWorkspace.IsNull() || o.CustomerFacingWorkspace.IsUnknown() {
		return e, false
	}
	var v []Workspace_SdkV2
	d := o.CustomerFacingWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCustomerFacingWorkspace sets the value of the CustomerFacingWorkspace field in UpdateWorkspaceRequest_SdkV2.
func (o *UpdateWorkspaceRequest_SdkV2) SetCustomerFacingWorkspace(ctx context.Context, v Workspace_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["customer_facing_workspace"]
	o.CustomerFacingWorkspace = types.ListValueMust(t, vs)
}

// *
type VpcEndpoint_SdkV2 struct {
	// The Databricks account ID that hosts the VPC endpoint configuration. TODO
	// - This may signal an OpenAPI diff; it does not show up in the generated
	// spec
	AccountId types.String `tfsdk:"account_id"`
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId types.String `tfsdk:"aws_account_id"`
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	AwsEndpointServiceId types.String `tfsdk:"aws_endpoint_service_id"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id"`
	// The cloud info of this vpc endpoint. Info for a GCP vpc endpoint.
	GcpVpcEndpointInfo types.List `tfsdk:"gcp_vpc_endpoint_info"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State types.String `tfsdk:"state"`

	UseCase types.String `tfsdk:"use_case"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId types.String `tfsdk:"vpc_endpoint_id"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name"`
}

func (to *VpcEndpoint_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VpcEndpoint_SdkV2) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				// Recursively sync the fields of GcpVpcEndpointInfo
				toGcpVpcEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (to *VpcEndpoint_SdkV2) SyncFieldsDuringRead(ctx context.Context, from VpcEndpoint_SdkV2) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				toGcpVpcEndpointInfo.SyncFieldsDuringRead(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (c VpcEndpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_account_id"] = attrs["aws_account_id"].SetOptional()
	attrs["aws_endpoint_service_id"] = attrs["aws_endpoint_service_id"].SetOptional()
	attrs["aws_vpc_endpoint_id"] = attrs["aws_vpc_endpoint_id"].SetOptional()
	attrs["gcp_vpc_endpoint_info"] = attrs["gcp_vpc_endpoint_info"].SetOptional()
	attrs["gcp_vpc_endpoint_info"] = attrs["gcp_vpc_endpoint_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["region"] = attrs["region"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["use_case"] = attrs["use_case"].SetOptional()
	attrs["vpc_endpoint_id"] = attrs["vpc_endpoint_id"].SetOptional()
	attrs["vpc_endpoint_name"] = attrs["vpc_endpoint_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VpcEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a VpcEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VpcEndpoint_SdkV2
// only implements ToObjectValue() and Type().
func (o VpcEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":              o.AccountId,
			"aws_account_id":          o.AwsAccountId,
			"aws_endpoint_service_id": o.AwsEndpointServiceId,
			"aws_vpc_endpoint_id":     o.AwsVpcEndpointId,
			"gcp_vpc_endpoint_info":   o.GcpVpcEndpointInfo,
			"region":                  o.Region,
			"state":                   o.State,
			"use_case":                o.UseCase,
			"vpc_endpoint_id":         o.VpcEndpointId,
			"vpc_endpoint_name":       o.VpcEndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VpcEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":              types.StringType,
			"aws_account_id":          types.StringType,
			"aws_endpoint_service_id": types.StringType,
			"aws_vpc_endpoint_id":     types.StringType,
			"gcp_vpc_endpoint_info": basetypes.ListType{
				ElemType: GcpVpcEndpointInfo_SdkV2{}.Type(ctx),
			},
			"region":            types.StringType,
			"state":             types.StringType,
			"use_case":          types.StringType,
			"vpc_endpoint_id":   types.StringType,
			"vpc_endpoint_name": types.StringType,
		},
	}
}

// GetGcpVpcEndpointInfo returns the value of the GcpVpcEndpointInfo field in VpcEndpoint_SdkV2 as
// a GcpVpcEndpointInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *VpcEndpoint_SdkV2) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo_SdkV2, bool) {
	var e GcpVpcEndpointInfo_SdkV2
	if o.GcpVpcEndpointInfo.IsNull() || o.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v []GcpVpcEndpointInfo_SdkV2
	d := o.GcpVpcEndpointInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in VpcEndpoint_SdkV2.
func (o *VpcEndpoint_SdkV2) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_vpc_endpoint_info"]
	o.GcpVpcEndpointInfo = types.ListValueMust(t, vs)
}

type Workspace_SdkV2 struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id"`

	AwsRegion types.String `tfsdk:"aws_region"`

	AzureWorkspaceInfo types.List `tfsdk:"azure_workspace_info"`
	// The cloud name. This field always has the value `gcp`.
	Cloud types.String `tfsdk:"cloud"`

	CloudResourceContainer types.List `tfsdk:"cloud_resource_container"`
	// The compute mode of the workspace.
	ComputeMode types.String `tfsdk:"compute_mode"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id"`

	CustomTags types.Map `tfsdk:"custom_tags"`

	DeploymentName types.String `tfsdk:"deployment_name"`
	// maps to external_customer_info from workspace proto this will contains
	// fields for the customers
	ExternalCustomerInfo types.List `tfsdk:"external_customer_info"`

	GcpManagedNetworkConfig types.List `tfsdk:"gcp_managed_network_config"`

	GkeConfig types.List `tfsdk:"gke_config"`
	// Whether No Public IP is enabled for the workspace
	IsNoPublicIpEnabled types.Bool `tfsdk:"is_no_public_ip_enabled"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location types.String `tfsdk:"location"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id"`
	// The network configuration for the workspace.
	//
	// DEPRECATED. Use `network_id` instead.
	Network types.List `tfsdk:"network"`
	// The object ID of network connectivity config.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// If this workspace is BYO VPC, then the network_id will be populated. If
	// this workspace is not BYO VPC, then the network_id will be empty.
	NetworkId types.String `tfsdk:"network_id"`

	PricingTier types.String `tfsdk:"pricing_tier"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id"`
	// The storage mode of the workspace.
	StorageMode types.String `tfsdk:"storage_mode"`
	// A unique integer ID for the workspace
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
	// The human-readable name of the workspace.
	WorkspaceName types.String `tfsdk:"workspace_name"`
	// The status of a workspace
	WorkspaceStatus types.String `tfsdk:"workspace_status"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage types.String `tfsdk:"workspace_status_message"`
}

func (to *Workspace_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Workspace_SdkV2) {
	if !from.AzureWorkspaceInfo.IsNull() && !from.AzureWorkspaceInfo.IsUnknown() {
		if toAzureWorkspaceInfo, ok := to.GetAzureWorkspaceInfo(ctx); ok {
			if fromAzureWorkspaceInfo, ok := from.GetAzureWorkspaceInfo(ctx); ok {
				// Recursively sync the fields of AzureWorkspaceInfo
				toAzureWorkspaceInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureWorkspaceInfo)
				to.SetAzureWorkspaceInfo(ctx, toAzureWorkspaceInfo)
			}
		}
	}
	if !from.CloudResourceContainer.IsNull() && !from.CloudResourceContainer.IsUnknown() {
		if toCloudResourceContainer, ok := to.GetCloudResourceContainer(ctx); ok {
			if fromCloudResourceContainer, ok := from.GetCloudResourceContainer(ctx); ok {
				// Recursively sync the fields of CloudResourceContainer
				toCloudResourceContainer.SyncFieldsDuringCreateOrUpdate(ctx, fromCloudResourceContainer)
				to.SetCloudResourceContainer(ctx, toCloudResourceContainer)
			}
		}
	}
	if !from.ExternalCustomerInfo.IsNull() && !from.ExternalCustomerInfo.IsUnknown() {
		if toExternalCustomerInfo, ok := to.GetExternalCustomerInfo(ctx); ok {
			if fromExternalCustomerInfo, ok := from.GetExternalCustomerInfo(ctx); ok {
				// Recursively sync the fields of ExternalCustomerInfo
				toExternalCustomerInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromExternalCustomerInfo)
				to.SetExternalCustomerInfo(ctx, toExternalCustomerInfo)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				// Recursively sync the fields of GcpManagedNetworkConfig
				toGcpManagedNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
	if !from.GkeConfig.IsNull() && !from.GkeConfig.IsUnknown() {
		if toGkeConfig, ok := to.GetGkeConfig(ctx); ok {
			if fromGkeConfig, ok := from.GetGkeConfig(ctx); ok {
				// Recursively sync the fields of GkeConfig
				toGkeConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGkeConfig)
				to.SetGkeConfig(ctx, toGkeConfig)
			}
		}
	}
	if !from.Network.IsNull() && !from.Network.IsUnknown() {
		if toNetwork, ok := to.GetNetwork(ctx); ok {
			if fromNetwork, ok := from.GetNetwork(ctx); ok {
				// Recursively sync the fields of Network
				toNetwork.SyncFieldsDuringCreateOrUpdate(ctx, fromNetwork)
				to.SetNetwork(ctx, toNetwork)
			}
		}
	}
}

func (to *Workspace_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Workspace_SdkV2) {
	if !from.AzureWorkspaceInfo.IsNull() && !from.AzureWorkspaceInfo.IsUnknown() {
		if toAzureWorkspaceInfo, ok := to.GetAzureWorkspaceInfo(ctx); ok {
			if fromAzureWorkspaceInfo, ok := from.GetAzureWorkspaceInfo(ctx); ok {
				toAzureWorkspaceInfo.SyncFieldsDuringRead(ctx, fromAzureWorkspaceInfo)
				to.SetAzureWorkspaceInfo(ctx, toAzureWorkspaceInfo)
			}
		}
	}
	if !from.CloudResourceContainer.IsNull() && !from.CloudResourceContainer.IsUnknown() {
		if toCloudResourceContainer, ok := to.GetCloudResourceContainer(ctx); ok {
			if fromCloudResourceContainer, ok := from.GetCloudResourceContainer(ctx); ok {
				toCloudResourceContainer.SyncFieldsDuringRead(ctx, fromCloudResourceContainer)
				to.SetCloudResourceContainer(ctx, toCloudResourceContainer)
			}
		}
	}
	if !from.ExternalCustomerInfo.IsNull() && !from.ExternalCustomerInfo.IsUnknown() {
		if toExternalCustomerInfo, ok := to.GetExternalCustomerInfo(ctx); ok {
			if fromExternalCustomerInfo, ok := from.GetExternalCustomerInfo(ctx); ok {
				toExternalCustomerInfo.SyncFieldsDuringRead(ctx, fromExternalCustomerInfo)
				to.SetExternalCustomerInfo(ctx, toExternalCustomerInfo)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				toGcpManagedNetworkConfig.SyncFieldsDuringRead(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
	if !from.GkeConfig.IsNull() && !from.GkeConfig.IsUnknown() {
		if toGkeConfig, ok := to.GetGkeConfig(ctx); ok {
			if fromGkeConfig, ok := from.GetGkeConfig(ctx); ok {
				toGkeConfig.SyncFieldsDuringRead(ctx, fromGkeConfig)
				to.SetGkeConfig(ctx, toGkeConfig)
			}
		}
	}
	if !from.Network.IsNull() && !from.Network.IsUnknown() {
		if toNetwork, ok := to.GetNetwork(ctx); ok {
			if fromNetwork, ok := from.GetNetwork(ctx); ok {
				toNetwork.SyncFieldsDuringRead(ctx, fromNetwork)
				to.SetNetwork(ctx, toNetwork)
			}
		}
	}
}

func (c Workspace_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_region"] = attrs["aws_region"].SetOptional()
	attrs["azure_workspace_info"] = attrs["azure_workspace_info"].SetComputed()
	attrs["azure_workspace_info"] = attrs["azure_workspace_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["compute_mode"] = attrs["compute_mode"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["deployment_name"] = attrs["deployment_name"].SetOptional()
	attrs["external_customer_info"] = attrs["external_customer_info"].SetOptional()
	attrs["external_customer_info"] = attrs["external_customer_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["gke_config"] = attrs["gke_config"].SetOptional()
	attrs["gke_config"] = attrs["gke_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["is_no_public_ip_enabled"] = attrs["is_no_public_ip_enabled"].SetOptional()
	attrs["location"] = attrs["location"].SetOptional()
	attrs["managed_services_customer_managed_key_id"] = attrs["managed_services_customer_managed_key_id"].SetOptional()
	attrs["network"] = attrs["network"].SetOptional()
	attrs["network"] = attrs["network"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetOptional()
	attrs["network_id"] = attrs["network_id"].SetOptional()
	attrs["pricing_tier"] = attrs["pricing_tier"].SetOptional()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetOptional()
	attrs["storage_customer_managed_key_id"] = attrs["storage_customer_managed_key_id"].SetOptional()
	attrs["storage_mode"] = attrs["storage_mode"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_name"] = attrs["workspace_name"].SetOptional()
	attrs["workspace_status"] = attrs["workspace_status"].SetComputed()
	attrs["workspace_status_message"] = attrs["workspace_status_message"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Workspace.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Workspace_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_workspace_info":       reflect.TypeOf(AzureWorkspaceInfo_SdkV2{}),
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer_SdkV2{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"external_customer_info":     reflect.TypeOf(ExternalCustomerInfo_SdkV2{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig_SdkV2{}),
		"gke_config":                 reflect.TypeOf(GkeConfig_SdkV2{}),
		"network":                    reflect.TypeOf(WorkspaceNetwork_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Workspace_SdkV2
// only implements ToObjectValue() and Type().
func (o Workspace_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                 o.AccountId,
			"aws_region":                 o.AwsRegion,
			"azure_workspace_info":       o.AzureWorkspaceInfo,
			"cloud":                      o.Cloud,
			"cloud_resource_container":   o.CloudResourceContainer,
			"compute_mode":               o.ComputeMode,
			"creation_time":              o.CreationTime,
			"credentials_id":             o.CredentialsId,
			"custom_tags":                o.CustomTags,
			"deployment_name":            o.DeploymentName,
			"external_customer_info":     o.ExternalCustomerInfo,
			"gcp_managed_network_config": o.GcpManagedNetworkConfig,
			"gke_config":                 o.GkeConfig,
			"is_no_public_ip_enabled":    o.IsNoPublicIpEnabled,
			"location":                   o.Location,
			"managed_services_customer_managed_key_id": o.ManagedServicesCustomerManagedKeyId,
			"network":                         o.Network,
			"network_connectivity_config_id":  o.NetworkConnectivityConfigId,
			"network_id":                      o.NetworkId,
			"pricing_tier":                    o.PricingTier,
			"private_access_settings_id":      o.PrivateAccessSettingsId,
			"storage_configuration_id":        o.StorageConfigurationId,
			"storage_customer_managed_key_id": o.StorageCustomerManagedKeyId,
			"storage_mode":                    o.StorageMode,
			"workspace_id":                    o.WorkspaceId,
			"workspace_name":                  o.WorkspaceName,
			"workspace_status":                o.WorkspaceStatus,
			"workspace_status_message":        o.WorkspaceStatusMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Workspace_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"aws_region": types.StringType,
			"azure_workspace_info": basetypes.ListType{
				ElemType: AzureWorkspaceInfo_SdkV2{}.Type(ctx),
			},
			"cloud": types.StringType,
			"cloud_resource_container": basetypes.ListType{
				ElemType: CloudResourceContainer_SdkV2{}.Type(ctx),
			},
			"compute_mode":   types.StringType,
			"creation_time":  types.Int64Type,
			"credentials_id": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name": types.StringType,
			"external_customer_info": basetypes.ListType{
				ElemType: ExternalCustomerInfo_SdkV2{}.Type(ctx),
			},
			"gcp_managed_network_config": basetypes.ListType{
				ElemType: GcpManagedNetworkConfig_SdkV2{}.Type(ctx),
			},
			"gke_config": basetypes.ListType{
				ElemType: GkeConfig_SdkV2{}.Type(ctx),
			},
			"is_no_public_ip_enabled": types.BoolType,
			"location":                types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network": basetypes.ListType{
				ElemType: WorkspaceNetwork_SdkV2{}.Type(ctx),
			},
			"network_connectivity_config_id":  types.StringType,
			"network_id":                      types.StringType,
			"pricing_tier":                    types.StringType,
			"private_access_settings_id":      types.StringType,
			"storage_configuration_id":        types.StringType,
			"storage_customer_managed_key_id": types.StringType,
			"storage_mode":                    types.StringType,
			"workspace_id":                    types.Int64Type,
			"workspace_name":                  types.StringType,
			"workspace_status":                types.StringType,
			"workspace_status_message":        types.StringType,
		},
	}
}

// GetAzureWorkspaceInfo returns the value of the AzureWorkspaceInfo field in Workspace_SdkV2 as
// a AzureWorkspaceInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetAzureWorkspaceInfo(ctx context.Context) (AzureWorkspaceInfo_SdkV2, bool) {
	var e AzureWorkspaceInfo_SdkV2
	if o.AzureWorkspaceInfo.IsNull() || o.AzureWorkspaceInfo.IsUnknown() {
		return e, false
	}
	var v []AzureWorkspaceInfo_SdkV2
	d := o.AzureWorkspaceInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureWorkspaceInfo sets the value of the AzureWorkspaceInfo field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetAzureWorkspaceInfo(ctx context.Context, v AzureWorkspaceInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_workspace_info"]
	o.AzureWorkspaceInfo = types.ListValueMust(t, vs)
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in Workspace_SdkV2 as
// a CloudResourceContainer_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer_SdkV2, bool) {
	var e CloudResourceContainer_SdkV2
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v []CloudResourceContainer_SdkV2
	d := o.CloudResourceContainer.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloud_resource_container"]
	o.CloudResourceContainer = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in Workspace_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetExternalCustomerInfo returns the value of the ExternalCustomerInfo field in Workspace_SdkV2 as
// a ExternalCustomerInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetExternalCustomerInfo(ctx context.Context) (ExternalCustomerInfo_SdkV2, bool) {
	var e ExternalCustomerInfo_SdkV2
	if o.ExternalCustomerInfo.IsNull() || o.ExternalCustomerInfo.IsUnknown() {
		return e, false
	}
	var v []ExternalCustomerInfo_SdkV2
	d := o.ExternalCustomerInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalCustomerInfo sets the value of the ExternalCustomerInfo field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetExternalCustomerInfo(ctx context.Context, v ExternalCustomerInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_customer_info"]
	o.ExternalCustomerInfo = types.ListValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in Workspace_SdkV2 as
// a GcpManagedNetworkConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig_SdkV2, bool) {
	var e GcpManagedNetworkConfig_SdkV2
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v []GcpManagedNetworkConfig_SdkV2
	d := o.GcpManagedNetworkConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_managed_network_config"]
	o.GcpManagedNetworkConfig = types.ListValueMust(t, vs)
}

// GetGkeConfig returns the value of the GkeConfig field in Workspace_SdkV2 as
// a GkeConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetGkeConfig(ctx context.Context) (GkeConfig_SdkV2, bool) {
	var e GkeConfig_SdkV2
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v []GkeConfig_SdkV2
	d := o.GkeConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGkeConfig sets the value of the GkeConfig field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetGkeConfig(ctx context.Context, v GkeConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gke_config"]
	o.GkeConfig = types.ListValueMust(t, vs)
}

// GetNetwork returns the value of the Network field in Workspace_SdkV2 as
// a WorkspaceNetwork_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetNetwork(ctx context.Context) (WorkspaceNetwork_SdkV2, bool) {
	var e WorkspaceNetwork_SdkV2
	if o.Network.IsNull() || o.Network.IsUnknown() {
		return e, false
	}
	var v []WorkspaceNetwork_SdkV2
	d := o.Network.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNetwork sets the value of the Network field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetNetwork(ctx context.Context, v WorkspaceNetwork_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["network"]
	o.Network = types.ListValueMust(t, vs)
}

// The network configuration for workspaces.
type WorkspaceNetwork_SdkV2 struct {
	// The shared network config for GCP workspace. This object has common
	// network configurations that are network attributions of a workspace. This
	// object is input-only.
	GcpCommonNetworkConfig types.List `tfsdk:"gcp_common_network_config"`
	// The mutually exclusive network deployment modes. The option decides which
	// network mode the workspace will use. The network config for GCP workspace
	// with Databricks managed network. This object is input-only and will not
	// be provided when listing workspaces. See go/gcp-byovpc-alpha-design for
	// interface decisions.
	GcpManagedNetworkConfig types.List `tfsdk:"gcp_managed_network_config"`
	// The ID of the network object, if the workspace is a BYOVPC workspace.
	// This should apply to workspaces on all clouds in internal services. In
	// accounts-rest-api, user will use workspace.network_id for input and
	// output instead. Currently (2021-06-19) the network ID is only used by
	// GCP.
	NetworkId types.String `tfsdk:"network_id"`
}

func (to *WorkspaceNetwork_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceNetwork_SdkV2) {
	if !from.GcpCommonNetworkConfig.IsNull() && !from.GcpCommonNetworkConfig.IsUnknown() {
		if toGcpCommonNetworkConfig, ok := to.GetGcpCommonNetworkConfig(ctx); ok {
			if fromGcpCommonNetworkConfig, ok := from.GetGcpCommonNetworkConfig(ctx); ok {
				// Recursively sync the fields of GcpCommonNetworkConfig
				toGcpCommonNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpCommonNetworkConfig)
				to.SetGcpCommonNetworkConfig(ctx, toGcpCommonNetworkConfig)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				// Recursively sync the fields of GcpManagedNetworkConfig
				toGcpManagedNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
}

func (to *WorkspaceNetwork_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceNetwork_SdkV2) {
	if !from.GcpCommonNetworkConfig.IsNull() && !from.GcpCommonNetworkConfig.IsUnknown() {
		if toGcpCommonNetworkConfig, ok := to.GetGcpCommonNetworkConfig(ctx); ok {
			if fromGcpCommonNetworkConfig, ok := from.GetGcpCommonNetworkConfig(ctx); ok {
				toGcpCommonNetworkConfig.SyncFieldsDuringRead(ctx, fromGcpCommonNetworkConfig)
				to.SetGcpCommonNetworkConfig(ctx, toGcpCommonNetworkConfig)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				toGcpManagedNetworkConfig.SyncFieldsDuringRead(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
}

func (c WorkspaceNetwork_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_common_network_config"] = attrs["gcp_common_network_config"].SetOptional()
	attrs["gcp_common_network_config"] = attrs["gcp_common_network_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["network_id"] = attrs["network_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceNetwork.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceNetwork_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_common_network_config":  reflect.TypeOf(GcpCommonNetworkConfig_SdkV2{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceNetwork_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceNetwork_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_common_network_config":  o.GcpCommonNetworkConfig,
			"gcp_managed_network_config": o.GcpManagedNetworkConfig,
			"network_id":                 o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceNetwork_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_common_network_config": basetypes.ListType{
				ElemType: GcpCommonNetworkConfig_SdkV2{}.Type(ctx),
			},
			"gcp_managed_network_config": basetypes.ListType{
				ElemType: GcpManagedNetworkConfig_SdkV2{}.Type(ctx),
			},
			"network_id": types.StringType,
		},
	}
}

// GetGcpCommonNetworkConfig returns the value of the GcpCommonNetworkConfig field in WorkspaceNetwork_SdkV2 as
// a GcpCommonNetworkConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceNetwork_SdkV2) GetGcpCommonNetworkConfig(ctx context.Context) (GcpCommonNetworkConfig_SdkV2, bool) {
	var e GcpCommonNetworkConfig_SdkV2
	if o.GcpCommonNetworkConfig.IsNull() || o.GcpCommonNetworkConfig.IsUnknown() {
		return e, false
	}
	var v []GcpCommonNetworkConfig_SdkV2
	d := o.GcpCommonNetworkConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpCommonNetworkConfig sets the value of the GcpCommonNetworkConfig field in WorkspaceNetwork_SdkV2.
func (o *WorkspaceNetwork_SdkV2) SetGcpCommonNetworkConfig(ctx context.Context, v GcpCommonNetworkConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_common_network_config"]
	o.GcpCommonNetworkConfig = types.ListValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in WorkspaceNetwork_SdkV2 as
// a GcpManagedNetworkConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceNetwork_SdkV2) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig_SdkV2, bool) {
	var e GcpManagedNetworkConfig_SdkV2
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v []GcpManagedNetworkConfig_SdkV2
	d := o.GcpManagedNetworkConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in WorkspaceNetwork_SdkV2.
func (o *WorkspaceNetwork_SdkV2) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_managed_network_config"]
	o.GcpManagedNetworkConfig = types.ListValueMust(t, vs)
}
