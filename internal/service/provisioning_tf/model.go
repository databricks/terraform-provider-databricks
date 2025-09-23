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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AwsCredentials struct {
	StsRole types.Object `tfsdk:"sts_role"`
}

func (to *AwsCredentials) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsCredentials) {
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

func (to *AwsCredentials) SyncFieldsDuringRead(ctx context.Context, from AwsCredentials) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				toStsRole.SyncFieldsDuringRead(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (c AwsCredentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sts_role"] = attrs["sts_role"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsCredentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(StsRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsCredentials
// only implements ToObjectValue() and Type().
func (o AwsCredentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": o.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsCredentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": StsRole{}.Type(ctx),
		},
	}
}

// GetStsRole returns the value of the StsRole field in AwsCredentials as
// a StsRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsCredentials) GetStsRole(ctx context.Context) (StsRole, bool) {
	var e StsRole
	if o.StsRole.IsNull() || o.StsRole.IsUnknown() {
		return e, false
	}
	var v StsRole
	d := o.StsRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStsRole sets the value of the StsRole field in AwsCredentials.
func (o *AwsCredentials) SetStsRole(ctx context.Context, v StsRole) {
	vs := v.ToObjectValue(ctx)
	o.StsRole = vs
}

type AwsDbManagedNetworkExtraInfo struct {
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
	NpipVpcInfra types.Object `tfsdk:"npip_vpc_infra"`
	// Security group which is given to the user to manage without Databricks
	// interference.
	UnmanagedSecurityGroup types.String `tfsdk:"unmanaged_security_group"`
	// Contents of the secret key which gives ssh access to the workers.
	WorkerKeyContents types.String `tfsdk:"worker_key_contents"`
	// Name of the keypair in AWS which allows sshing into the workers.
	WorkerKeypairName types.String `tfsdk:"worker_keypair_name"`
}

func (to *AwsDbManagedNetworkExtraInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsDbManagedNetworkExtraInfo) {
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

func (to *AwsDbManagedNetworkExtraInfo) SyncFieldsDuringRead(ctx context.Context, from AwsDbManagedNetworkExtraInfo) {
	if !from.NpipVpcInfra.IsNull() && !from.NpipVpcInfra.IsUnknown() {
		if toNpipVpcInfra, ok := to.GetNpipVpcInfra(ctx); ok {
			if fromNpipVpcInfra, ok := from.GetNpipVpcInfra(ctx); ok {
				toNpipVpcInfra.SyncFieldsDuringRead(ctx, fromNpipVpcInfra)
				to.SetNpipVpcInfra(ctx, toNpipVpcInfra)
			}
		}
	}
}

func (c AwsDbManagedNetworkExtraInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dhcp_options_id"] = attrs["dhcp_options_id"].SetOptional()
	attrs["gateway_id"] = attrs["gateway_id"].SetOptional()
	attrs["managed_security_group"] = attrs["managed_security_group"].SetOptional()
	attrs["npip_vpc_infra"] = attrs["npip_vpc_infra"].SetOptional()
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
func (a AwsDbManagedNetworkExtraInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"npip_vpc_infra": reflect.TypeOf(NpipVpcInfra{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsDbManagedNetworkExtraInfo
// only implements ToObjectValue() and Type().
func (o AwsDbManagedNetworkExtraInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AwsDbManagedNetworkExtraInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dhcp_options_id":          types.StringType,
			"gateway_id":               types.StringType,
			"managed_security_group":   types.StringType,
			"npip_vpc_infra":           NpipVpcInfra{}.Type(ctx),
			"unmanaged_security_group": types.StringType,
			"worker_key_contents":      types.StringType,
			"worker_keypair_name":      types.StringType,
		},
	}
}

// GetNpipVpcInfra returns the value of the NpipVpcInfra field in AwsDbManagedNetworkExtraInfo as
// a NpipVpcInfra value.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsDbManagedNetworkExtraInfo) GetNpipVpcInfra(ctx context.Context) (NpipVpcInfra, bool) {
	var e NpipVpcInfra
	if o.NpipVpcInfra.IsNull() || o.NpipVpcInfra.IsUnknown() {
		return e, false
	}
	var v NpipVpcInfra
	d := o.NpipVpcInfra.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNpipVpcInfra sets the value of the NpipVpcInfra field in AwsDbManagedNetworkExtraInfo.
func (o *AwsDbManagedNetworkExtraInfo) SetNpipVpcInfra(ctx context.Context, v NpipVpcInfra) {
	vs := v.ToObjectValue(ctx)
	o.NpipVpcInfra = vs
}

type AwsKeyInfo struct {
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

func (to *AwsKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsKeyInfo) {
}

func (to *AwsKeyInfo) SyncFieldsDuringRead(ctx context.Context, from AwsKeyInfo) {
}

func (c AwsKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AwsKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsKeyInfo
// only implements ToObjectValue() and Type().
func (o AwsKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AwsKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"key_region":                    types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type AwsNetworkInfo struct {
	// Additional information for DB managed VPC, which is mainly used to
	// populate WorkerEnvironment.
	DbManagedVpcExtraInfo types.Object `tfsdk:"db_managed_vpc_extra_info"`
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

func (to *AwsNetworkInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsNetworkInfo) {
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

func (to *AwsNetworkInfo) SyncFieldsDuringRead(ctx context.Context, from AwsNetworkInfo) {
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

func (c AwsNetworkInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["db_managed_vpc_extra_info"] = attrs["db_managed_vpc_extra_info"].SetOptional()
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
func (a AwsNetworkInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"db_managed_vpc_extra_info": reflect.TypeOf(AwsDbManagedNetworkExtraInfo{}),
		"security_group_ids":        reflect.TypeOf(types.String{}),
		"subnet_ids":                reflect.TypeOf(types.String{}),
		"subnets":                   reflect.TypeOf(SubnetInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsNetworkInfo
// only implements ToObjectValue() and Type().
func (o AwsNetworkInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AwsNetworkInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"db_managed_vpc_extra_info": AwsDbManagedNetworkExtraInfo{}.Type(ctx),
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnets": basetypes.ListType{
				ElemType: SubnetInfo{}.Type(ctx),
			},
			"vpc_address_space": types.StringType,
			"vpc_id":            types.StringType,
		},
	}
}

// GetDbManagedVpcExtraInfo returns the value of the DbManagedVpcExtraInfo field in AwsNetworkInfo as
// a AwsDbManagedNetworkExtraInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsNetworkInfo) GetDbManagedVpcExtraInfo(ctx context.Context) (AwsDbManagedNetworkExtraInfo, bool) {
	var e AwsDbManagedNetworkExtraInfo
	if o.DbManagedVpcExtraInfo.IsNull() || o.DbManagedVpcExtraInfo.IsUnknown() {
		return e, false
	}
	var v AwsDbManagedNetworkExtraInfo
	d := o.DbManagedVpcExtraInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDbManagedVpcExtraInfo sets the value of the DbManagedVpcExtraInfo field in AwsNetworkInfo.
func (o *AwsNetworkInfo) SetDbManagedVpcExtraInfo(ctx context.Context, v AwsDbManagedNetworkExtraInfo) {
	vs := v.ToObjectValue(ctx)
	o.DbManagedVpcExtraInfo = vs
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in AwsNetworkInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsNetworkInfo) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
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

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in AwsNetworkInfo.
func (o *AwsNetworkInfo) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in AwsNetworkInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsNetworkInfo) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
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

// SetSubnetIds sets the value of the SubnetIds field in AwsNetworkInfo.
func (o *AwsNetworkInfo) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetSubnets returns the value of the Subnets field in AwsNetworkInfo as
// a slice of SubnetInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsNetworkInfo) GetSubnets(ctx context.Context) ([]SubnetInfo, bool) {
	if o.Subnets.IsNull() || o.Subnets.IsUnknown() {
		return nil, false
	}
	var v []SubnetInfo
	d := o.Subnets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnets sets the value of the Subnets field in AwsNetworkInfo.
func (o *AwsNetworkInfo) SetSubnets(ctx context.Context, v []SubnetInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subnets = types.ListValueMust(t, vs)
}

type AzureKeyInfo struct {
	// The Disk Encryption Set id that is used to represent the key info used
	// for Managed Disk BYOK use case
	DiskEncryptionSetId types.String `tfsdk:"disk_encryption_set_id"`
	// The structure to store key access credential This is set if the Managed
	// Identity is being used to access the Azure Key Vault key.
	KeyAccessConfiguration types.Object `tfsdk:"key_access_configuration"`
	// The name of the key in KeyVault.
	KeyName types.String `tfsdk:"key_name"`
	// The base URI of the KeyVault.
	KeyVaultUri types.String `tfsdk:"key_vault_uri"`
	// The tenant id where the KeyVault lives.
	TenantId types.String `tfsdk:"tenant_id"`
	// The current key version.
	Version types.String `tfsdk:"version"`
}

func (to *AzureKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureKeyInfo) {
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

func (to *AzureKeyInfo) SyncFieldsDuringRead(ctx context.Context, from AzureKeyInfo) {
	if !from.KeyAccessConfiguration.IsNull() && !from.KeyAccessConfiguration.IsUnknown() {
		if toKeyAccessConfiguration, ok := to.GetKeyAccessConfiguration(ctx); ok {
			if fromKeyAccessConfiguration, ok := from.GetKeyAccessConfiguration(ctx); ok {
				toKeyAccessConfiguration.SyncFieldsDuringRead(ctx, fromKeyAccessConfiguration)
				to.SetKeyAccessConfiguration(ctx, toKeyAccessConfiguration)
			}
		}
	}
}

func (c AzureKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disk_encryption_set_id"] = attrs["disk_encryption_set_id"].SetOptional()
	attrs["key_access_configuration"] = attrs["key_access_configuration"].SetOptional()
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
func (a AzureKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"key_access_configuration": reflect.TypeOf(KeyAccessConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureKeyInfo
// only implements ToObjectValue() and Type().
func (o AzureKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AzureKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disk_encryption_set_id":   types.StringType,
			"key_access_configuration": KeyAccessConfiguration{}.Type(ctx),
			"key_name":                 types.StringType,
			"key_vault_uri":            types.StringType,
			"tenant_id":                types.StringType,
			"version":                  types.StringType,
		},
	}
}

// GetKeyAccessConfiguration returns the value of the KeyAccessConfiguration field in AzureKeyInfo as
// a KeyAccessConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (o *AzureKeyInfo) GetKeyAccessConfiguration(ctx context.Context) (KeyAccessConfiguration, bool) {
	var e KeyAccessConfiguration
	if o.KeyAccessConfiguration.IsNull() || o.KeyAccessConfiguration.IsUnknown() {
		return e, false
	}
	var v KeyAccessConfiguration
	d := o.KeyAccessConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKeyAccessConfiguration sets the value of the KeyAccessConfiguration field in AzureKeyInfo.
func (o *AzureKeyInfo) SetKeyAccessConfiguration(ctx context.Context, v KeyAccessConfiguration) {
	vs := v.ToObjectValue(ctx)
	o.KeyAccessConfiguration = vs
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	ResourceGroup types.String `tfsdk:"resource_group"`
	// Azure Subscription ID
	SubscriptionId types.String `tfsdk:"subscription_id"`
}

func (to *AzureWorkspaceInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureWorkspaceInfo) {
}

func (to *AzureWorkspaceInfo) SyncFieldsDuringRead(ctx context.Context, from AzureWorkspaceInfo) {
}

func (c AzureWorkspaceInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AzureWorkspaceInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureWorkspaceInfo
// only implements ToObjectValue() and Type().
func (o AzureWorkspaceInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resource_group":  o.ResourceGroup,
			"subscription_id": o.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureWorkspaceInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource_group":  types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

type CloudResourceContainer struct {
	Gcp types.Object `tfsdk:"gcp"`
}

func (to *CloudResourceContainer) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CloudResourceContainer) {
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

func (to *CloudResourceContainer) SyncFieldsDuringRead(ctx context.Context, from CloudResourceContainer) {
	if !from.Gcp.IsNull() && !from.Gcp.IsUnknown() {
		if toGcp, ok := to.GetGcp(ctx); ok {
			if fromGcp, ok := from.GetGcp(ctx); ok {
				toGcp.SyncFieldsDuringRead(ctx, fromGcp)
				to.SetGcp(ctx, toGcp)
			}
		}
	}
}

func (c CloudResourceContainer) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp"] = attrs["gcp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloudResourceContainer.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CloudResourceContainer) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp": reflect.TypeOf(CustomerFacingGcpCloudResourceContainer{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudResourceContainer
// only implements ToObjectValue() and Type().
func (o CloudResourceContainer) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp": o.Gcp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloudResourceContainer) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp": CustomerFacingGcpCloudResourceContainer{}.Type(ctx),
		},
	}
}

// GetGcp returns the value of the Gcp field in CloudResourceContainer as
// a CustomerFacingGcpCloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (o *CloudResourceContainer) GetGcp(ctx context.Context) (CustomerFacingGcpCloudResourceContainer, bool) {
	var e CustomerFacingGcpCloudResourceContainer
	if o.Gcp.IsNull() || o.Gcp.IsUnknown() {
		return e, false
	}
	var v CustomerFacingGcpCloudResourceContainer
	d := o.Gcp.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcp sets the value of the Gcp field in CloudResourceContainer.
func (o *CloudResourceContainer) SetGcp(ctx context.Context, v CustomerFacingGcpCloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	o.Gcp = vs
}

type CreateAwsKeyInfo struct {
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

func (to *CreateAwsKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAwsKeyInfo) {
}

func (to *CreateAwsKeyInfo) SyncFieldsDuringRead(ctx context.Context, from CreateAwsKeyInfo) {
}

func (c CreateAwsKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateAwsKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAwsKeyInfo
// only implements ToObjectValue() and Type().
func (o CreateAwsKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateAwsKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"key_region":                    types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type CreateCredentialAwsCredentials struct {
	StsRole types.Object `tfsdk:"sts_role"`
}

func (to *CreateCredentialAwsCredentials) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialAwsCredentials) {
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

func (to *CreateCredentialAwsCredentials) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialAwsCredentials) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				toStsRole.SyncFieldsDuringRead(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (c CreateCredentialAwsCredentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sts_role"] = attrs["sts_role"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialAwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialAwsCredentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(CreateCredentialStsRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialAwsCredentials
// only implements ToObjectValue() and Type().
func (o CreateCredentialAwsCredentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": o.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialAwsCredentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": CreateCredentialStsRole{}.Type(ctx),
		},
	}
}

// GetStsRole returns the value of the StsRole field in CreateCredentialAwsCredentials as
// a CreateCredentialStsRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialAwsCredentials) GetStsRole(ctx context.Context) (CreateCredentialStsRole, bool) {
	var e CreateCredentialStsRole
	if o.StsRole.IsNull() || o.StsRole.IsUnknown() {
		return e, false
	}
	var v CreateCredentialStsRole
	d := o.StsRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStsRole sets the value of the StsRole field in CreateCredentialAwsCredentials.
func (o *CreateCredentialAwsCredentials) SetStsRole(ctx context.Context, v CreateCredentialStsRole) {
	vs := v.ToObjectValue(ctx)
	o.StsRole = vs
}

type CreateCredentialRequest struct {
	AwsCredentials types.Object `tfsdk:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name"`
}

func (to *CreateCredentialRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialRequest) {
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

func (to *CreateCredentialRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialRequest) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				toAwsCredentials.SyncFieldsDuringRead(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (c CreateCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_credentials"] = attrs["aws_credentials"].SetOptional()
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
func (a CreateCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_credentials": reflect.TypeOf(CreateCredentialAwsCredentials{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialRequest
// only implements ToObjectValue() and Type().
func (o CreateCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_credentials":  o.AwsCredentials,
			"credentials_name": o.CredentialsName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_credentials":  CreateCredentialAwsCredentials{}.Type(ctx),
			"credentials_name": types.StringType,
		},
	}
}

// GetAwsCredentials returns the value of the AwsCredentials field in CreateCredentialRequest as
// a CreateCredentialAwsCredentials value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest) GetAwsCredentials(ctx context.Context) (CreateCredentialAwsCredentials, bool) {
	var e CreateCredentialAwsCredentials
	if o.AwsCredentials.IsNull() || o.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v CreateCredentialAwsCredentials
	d := o.AwsCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsCredentials sets the value of the AwsCredentials field in CreateCredentialRequest.
func (o *CreateCredentialRequest) SetAwsCredentials(ctx context.Context, v CreateCredentialAwsCredentials) {
	vs := v.ToObjectValue(ctx)
	o.AwsCredentials = vs
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
type CreateCredentialStsRole struct {
	// Note: This must match the external_id on the parent object.
	//
	// TODO(j): Add validation to ensure this cannot be updated. If the user can
	// override the external_id, that defeats the purpose.
	ExternalId types.String `tfsdk:"external_id"`

	RoleArn types.String `tfsdk:"role_arn"`
}

func (to *CreateCredentialStsRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialStsRole) {
}

func (to *CreateCredentialStsRole) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialStsRole) {
}

func (c CreateCredentialStsRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateCredentialStsRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialStsRole
// only implements ToObjectValue() and Type().
func (o CreateCredentialStsRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
			"role_arn":    o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialStsRole) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
			"role_arn":    types.StringType,
		},
	}
}

type CreateCustomerManagedKeyRequest struct {
	AwsKeyInfo types.Object `tfsdk:"aws_key_info"`

	GcpKeyInfo types.Object `tfsdk:"gcp_key_info"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases"`
}

func (to *CreateCustomerManagedKeyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCustomerManagedKeyRequest) {
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

func (to *CreateCustomerManagedKeyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCustomerManagedKeyRequest) {
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

func (c CreateCustomerManagedKeyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_key_info"] = attrs["aws_key_info"].SetOptional()
	attrs["gcp_key_info"] = attrs["gcp_key_info"].SetOptional()
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
func (a CreateCustomerManagedKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info": reflect.TypeOf(CreateAwsKeyInfo{}),
		"gcp_key_info": reflect.TypeOf(CreateGcpKeyInfo{}),
		"use_cases":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomerManagedKeyRequest
// only implements ToObjectValue() and Type().
func (o CreateCustomerManagedKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_key_info": o.AwsKeyInfo,
			"gcp_key_info": o.GcpKeyInfo,
			"use_cases":    o.UseCases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCustomerManagedKeyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_key_info": CreateAwsKeyInfo{}.Type(ctx),
			"gcp_key_info": CreateGcpKeyInfo{}.Type(ctx),
			"use_cases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAwsKeyInfo returns the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest as
// a CreateAwsKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest) GetAwsKeyInfo(ctx context.Context) (CreateAwsKeyInfo, bool) {
	var e CreateAwsKeyInfo
	if o.AwsKeyInfo.IsNull() || o.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v CreateAwsKeyInfo
	d := o.AwsKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest.
func (o *CreateCustomerManagedKeyRequest) SetAwsKeyInfo(ctx context.Context, v CreateAwsKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.AwsKeyInfo = vs
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest as
// a CreateGcpKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest) GetGcpKeyInfo(ctx context.Context) (CreateGcpKeyInfo, bool) {
	var e CreateGcpKeyInfo
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v CreateGcpKeyInfo
	d := o.GcpKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest.
func (o *CreateCustomerManagedKeyRequest) SetGcpKeyInfo(ctx context.Context, v CreateGcpKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpKeyInfo = vs
}

// GetUseCases returns the value of the UseCases field in CreateCustomerManagedKeyRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest) GetUseCases(ctx context.Context) ([]types.String, bool) {
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

// SetUseCases sets the value of the UseCases field in CreateCustomerManagedKeyRequest.
func (o *CreateCustomerManagedKeyRequest) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UseCases = types.ListValueMust(t, vs)
}

type CreateGcpKeyInfo struct {
	// Globally unique service account email that has access to the KMS key. The
	// service account exists within the Databricks CP project.
	GcpServiceAccount types.Object `tfsdk:"gcp_service_account"`
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId types.String `tfsdk:"kms_key_id"`
}

func (to *CreateGcpKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateGcpKeyInfo) {
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

func (to *CreateGcpKeyInfo) SyncFieldsDuringRead(ctx context.Context, from CreateGcpKeyInfo) {
	if !from.GcpServiceAccount.IsNull() && !from.GcpServiceAccount.IsUnknown() {
		if toGcpServiceAccount, ok := to.GetGcpServiceAccount(ctx); ok {
			if fromGcpServiceAccount, ok := from.GetGcpServiceAccount(ctx); ok {
				toGcpServiceAccount.SyncFieldsDuringRead(ctx, fromGcpServiceAccount)
				to.SetGcpServiceAccount(ctx, toGcpServiceAccount)
			}
		}
	}
}

func (c CreateGcpKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_service_account"] = attrs["gcp_service_account"].SetOptional()
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
func (a CreateGcpKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_service_account": reflect.TypeOf(GcpServiceAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGcpKeyInfo
// only implements ToObjectValue() and Type().
func (o CreateGcpKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_service_account": o.GcpServiceAccount,
			"kms_key_id":          o.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateGcpKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_service_account": GcpServiceAccount{}.Type(ctx),
			"kms_key_id":          types.StringType,
		},
	}
}

// GetGcpServiceAccount returns the value of the GcpServiceAccount field in CreateGcpKeyInfo as
// a GcpServiceAccount value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateGcpKeyInfo) GetGcpServiceAccount(ctx context.Context) (GcpServiceAccount, bool) {
	var e GcpServiceAccount
	if o.GcpServiceAccount.IsNull() || o.GcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v GcpServiceAccount
	d := o.GcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpServiceAccount sets the value of the GcpServiceAccount field in CreateGcpKeyInfo.
func (o *CreateGcpKeyInfo) SetGcpServiceAccount(ctx context.Context, v GcpServiceAccount) {
	vs := v.ToObjectValue(ctx)
	o.GcpServiceAccount = vs
}

type CreateNetworkRequest struct {
	GcpNetworkInfo types.Object `tfsdk:"gcp_network_info"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds types.List `tfsdk:"security_group_ids"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds types.List `tfsdk:"subnet_ids"`

	VpcEndpoints types.Object `tfsdk:"vpc_endpoints"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId types.String `tfsdk:"vpc_id"`
}

func (to *CreateNetworkRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkRequest) {
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

func (to *CreateNetworkRequest) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkRequest) {
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

func (c CreateNetworkRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_network_info"] = attrs["gcp_network_info"].SetOptional()
	attrs["network_name"] = attrs["network_name"].SetOptional()
	attrs["security_group_ids"] = attrs["security_group_ids"].SetOptional()
	attrs["subnet_ids"] = attrs["subnet_ids"].SetOptional()
	attrs["vpc_endpoints"] = attrs["vpc_endpoints"].SetOptional()
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
func (a CreateNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_network_info":   reflect.TypeOf(GcpNetworkInfo{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
		"subnet_ids":         reflect.TypeOf(types.String{}),
		"vpc_endpoints":      reflect.TypeOf(NetworkVpcEndpoints{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkRequest
// only implements ToObjectValue() and Type().
func (o CreateNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateNetworkRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_network_info": GcpNetworkInfo{}.Type(ctx),
			"network_name":     types.StringType,
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"vpc_endpoints": NetworkVpcEndpoints{}.Type(ctx),
			"vpc_id":        types.StringType,
		},
	}
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in CreateNetworkRequest as
// a GcpNetworkInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo, bool) {
	var e GcpNetworkInfo
	if o.GcpNetworkInfo.IsNull() || o.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v GcpNetworkInfo
	d := o.GcpNetworkInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpNetworkInfo = vs
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in CreateNetworkRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
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

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in CreateNetworkRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
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

// SetSubnetIds sets the value of the SubnetIds field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in CreateNetworkRequest as
// a NetworkVpcEndpoints value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints, bool) {
	var e NetworkVpcEndpoints
	if o.VpcEndpoints.IsNull() || o.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v NetworkVpcEndpoints
	d := o.VpcEndpoints.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints) {
	vs := v.ToObjectValue(ctx)
	o.VpcEndpoints = vs
}

type CreatePrivateAccessSettingsRequest struct {
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

func (to *CreatePrivateAccessSettingsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePrivateAccessSettingsRequest) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (to *CreatePrivateAccessSettingsRequest) SyncFieldsDuringRead(ctx context.Context, from CreatePrivateAccessSettingsRequest) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (c CreatePrivateAccessSettingsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreatePrivateAccessSettingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateAccessSettingsRequest
// only implements ToObjectValue() and Type().
func (o CreatePrivateAccessSettingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreatePrivateAccessSettingsRequest) Type(ctx context.Context) attr.Type {
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

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in CreatePrivateAccessSettingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePrivateAccessSettingsRequest) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in CreatePrivateAccessSettingsRequest.
func (o *CreatePrivateAccessSettingsRequest) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type CreateStorageConfigurationRequest struct {
	RootBucketInfo types.Object `tfsdk:"root_bucket_info"`

	StorageConfigurationName types.String `tfsdk:"storage_configuration_name"`
}

func (to *CreateStorageConfigurationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateStorageConfigurationRequest) {
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

func (to *CreateStorageConfigurationRequest) SyncFieldsDuringRead(ctx context.Context, from CreateStorageConfigurationRequest) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				toRootBucketInfo.SyncFieldsDuringRead(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (c CreateStorageConfigurationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["root_bucket_info"] = attrs["root_bucket_info"].SetOptional()
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
func (a CreateStorageConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateStorageConfigurationRequest
// only implements ToObjectValue() and Type().
func (o CreateStorageConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"root_bucket_info":           o.RootBucketInfo,
			"storage_configuration_name": o.StorageConfigurationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateStorageConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"root_bucket_info":           RootBucketInfo{}.Type(ctx),
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in CreateStorageConfigurationRequest as
// a RootBucketInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageConfigurationRequest) GetRootBucketInfo(ctx context.Context) (RootBucketInfo, bool) {
	var e RootBucketInfo
	if o.RootBucketInfo.IsNull() || o.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v RootBucketInfo
	d := o.RootBucketInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in CreateStorageConfigurationRequest.
func (o *CreateStorageConfigurationRequest) SetRootBucketInfo(ctx context.Context, v RootBucketInfo) {
	vs := v.ToObjectValue(ctx)
	o.RootBucketInfo = vs
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id"`
	// The cloud info of this vpc endpoint.
	GcpVpcEndpointInfo types.Object `tfsdk:"gcp_vpc_endpoint_info"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name"`
}

func (to *CreateVpcEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVpcEndpointRequest) {
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

func (to *CreateVpcEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from CreateVpcEndpointRequest) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				toGcpVpcEndpointInfo.SyncFieldsDuringRead(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (c CreateVpcEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_vpc_endpoint_id"] = attrs["aws_vpc_endpoint_id"].SetOptional()
	attrs["gcp_vpc_endpoint_info"] = attrs["gcp_vpc_endpoint_info"].SetOptional()
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
func (a CreateVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (o CreateVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateVpcEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_vpc_endpoint_id":   types.StringType,
			"gcp_vpc_endpoint_info": GcpVpcEndpointInfo{}.Type(ctx),
			"region":                types.StringType,
			"vpc_endpoint_name":     types.StringType,
		},
	}
}

// GetGcpVpcEndpointInfo returns the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest as
// a GcpVpcEndpointInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateVpcEndpointRequest) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo, bool) {
	var e GcpVpcEndpointInfo
	if o.GcpVpcEndpointInfo.IsNull() || o.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v GcpVpcEndpointInfo
	d := o.GcpVpcEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest.
func (o *CreateVpcEndpointRequest) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpVpcEndpointInfo = vs
}

type CreateWorkspaceRequest struct {
	AwsRegion types.String `tfsdk:"aws_region"`
	// The cloud name. This field always has the value `gcp`.
	Cloud types.String `tfsdk:"cloud"`

	CloudResourceContainer types.Object `tfsdk:"cloud_resource_container"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id"`

	CustomTags types.Map `tfsdk:"custom_tags"`

	DeploymentName types.String `tfsdk:"deployment_name"`

	GcpManagedNetworkConfig types.Object `tfsdk:"gcp_managed_network_config"`

	GkeConfig types.Object `tfsdk:"gke_config"`
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

func (to *CreateWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceRequest) {
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

func (to *CreateWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceRequest) {
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

func (c CreateWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_region"] = attrs["aws_region"].SetOptional()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].SetOptional()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["deployment_name"] = attrs["deployment_name"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["gke_config"] = attrs["gke_config"].SetOptional()
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
func (a CreateWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig{}),
		"gke_config":                 reflect.TypeOf(GkeConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o CreateWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_region":               types.StringType,
			"cloud":                    types.StringType,
			"cloud_resource_container": CloudResourceContainer{}.Type(ctx),
			"credentials_id":           types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name":                          types.StringType,
			"gcp_managed_network_config":               GcpManagedNetworkConfig{}.Type(ctx),
			"gke_config":                               GkeConfig{}.Type(ctx),
			"is_no_public_ip_enabled":                  types.BoolType,
			"location":                                 types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network_id":                               types.StringType,
			"pricing_tier":                             types.StringType,
			"private_access_settings_id":               types.StringType,
			"storage_configuration_id":                 types.StringType,
			"storage_customer_managed_key_id":          types.StringType,
			"workspace_name":                           types.StringType,
		},
	}
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in CreateWorkspaceRequest as
// a CloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer, bool) {
	var e CloudResourceContainer
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v CloudResourceContainer
	d := o.CloudResourceContainer.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	o.CloudResourceContainer = vs
}

// GetCustomTags returns the value of the CustomTags field in CreateWorkspaceRequest as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpManagedNetworkConfig
	d := o.GcpManagedNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	o.GcpManagedNetworkConfig = vs
}

// GetGkeConfig returns the value of the GkeConfig field in CreateWorkspaceRequest as
// a GkeConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetGkeConfig(ctx context.Context) (GkeConfig, bool) {
	var e GkeConfig
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v GkeConfig
	d := o.GkeConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGkeConfig sets the value of the GkeConfig field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetGkeConfig(ctx context.Context, v GkeConfig) {
	vs := v.ToObjectValue(ctx)
	o.GkeConfig = vs
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id"`

	AwsCredentials types.Object `tfsdk:"aws_credentials"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Databricks credential configuration ID.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name"`
}

func (to *Credential) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Credential) {
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

func (to *Credential) SyncFieldsDuringRead(ctx context.Context, from Credential) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				toAwsCredentials.SyncFieldsDuringRead(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (c Credential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_credentials"] = attrs["aws_credentials"].SetOptional()
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
func (a Credential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_credentials": reflect.TypeOf(AwsCredentials{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Credential
// only implements ToObjectValue() and Type().
func (o Credential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Credential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":       types.StringType,
			"aws_credentials":  AwsCredentials{}.Type(ctx),
			"creation_time":    types.Int64Type,
			"credentials_id":   types.StringType,
			"credentials_name": types.StringType,
		},
	}
}

// GetAwsCredentials returns the value of the AwsCredentials field in Credential as
// a AwsCredentials value.
// If the field is unknown or null, the boolean return value is false.
func (o *Credential) GetAwsCredentials(ctx context.Context) (AwsCredentials, bool) {
	var e AwsCredentials
	if o.AwsCredentials.IsNull() || o.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v AwsCredentials
	d := o.AwsCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsCredentials sets the value of the AwsCredentials field in Credential.
func (o *Credential) SetAwsCredentials(ctx context.Context, v AwsCredentials) {
	vs := v.ToObjectValue(ctx)
	o.AwsCredentials = vs
}

type CustomerFacingGcpCloudResourceContainer struct {
	ProjectId types.String `tfsdk:"project_id"`
}

func (to *CustomerFacingGcpCloudResourceContainer) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomerFacingGcpCloudResourceContainer) {
}

func (to *CustomerFacingGcpCloudResourceContainer) SyncFieldsDuringRead(ctx context.Context, from CustomerFacingGcpCloudResourceContainer) {
}

func (c CustomerFacingGcpCloudResourceContainer) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CustomerFacingGcpCloudResourceContainer) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerFacingGcpCloudResourceContainer
// only implements ToObjectValue() and Type().
func (o CustomerFacingGcpCloudResourceContainer) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_id": o.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomerFacingGcpCloudResourceContainer) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_id": types.StringType,
		},
	}
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId types.String `tfsdk:"account_id"`

	AwsKeyInfo types.Object `tfsdk:"aws_key_info"`

	AzureKeyInfo types.Object `tfsdk:"azure_key_info"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId types.String `tfsdk:"customer_managed_key_id"`

	GcpKeyInfo types.Object `tfsdk:"gcp_key_info"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases"`
}

func (to *CustomerManagedKey) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomerManagedKey) {
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

func (to *CustomerManagedKey) SyncFieldsDuringRead(ctx context.Context, from CustomerManagedKey) {
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

func (c CustomerManagedKey) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_key_info"] = attrs["aws_key_info"].SetOptional()
	attrs["azure_key_info"] = attrs["azure_key_info"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["customer_managed_key_id"] = attrs["customer_managed_key_id"].SetOptional()
	attrs["gcp_key_info"] = attrs["gcp_key_info"].SetOptional()
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
func (a CustomerManagedKey) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info":   reflect.TypeOf(AwsKeyInfo{}),
		"azure_key_info": reflect.TypeOf(AzureKeyInfo{}),
		"gcp_key_info":   reflect.TypeOf(GcpKeyInfo{}),
		"use_cases":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerManagedKey
// only implements ToObjectValue() and Type().
func (o CustomerManagedKey) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CustomerManagedKey) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":              types.StringType,
			"aws_key_info":            AwsKeyInfo{}.Type(ctx),
			"azure_key_info":          AzureKeyInfo{}.Type(ctx),
			"creation_time":           types.Int64Type,
			"customer_managed_key_id": types.StringType,
			"gcp_key_info":            GcpKeyInfo{}.Type(ctx),
			"use_cases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAwsKeyInfo returns the value of the AwsKeyInfo field in CustomerManagedKey as
// a AwsKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey) GetAwsKeyInfo(ctx context.Context) (AwsKeyInfo, bool) {
	var e AwsKeyInfo
	if o.AwsKeyInfo.IsNull() || o.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v AwsKeyInfo
	d := o.AwsKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CustomerManagedKey.
func (o *CustomerManagedKey) SetAwsKeyInfo(ctx context.Context, v AwsKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.AwsKeyInfo = vs
}

// GetAzureKeyInfo returns the value of the AzureKeyInfo field in CustomerManagedKey as
// a AzureKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey) GetAzureKeyInfo(ctx context.Context) (AzureKeyInfo, bool) {
	var e AzureKeyInfo
	if o.AzureKeyInfo.IsNull() || o.AzureKeyInfo.IsUnknown() {
		return e, false
	}
	var v AzureKeyInfo
	d := o.AzureKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureKeyInfo sets the value of the AzureKeyInfo field in CustomerManagedKey.
func (o *CustomerManagedKey) SetAzureKeyInfo(ctx context.Context, v AzureKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.AzureKeyInfo = vs
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CustomerManagedKey as
// a GcpKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey) GetGcpKeyInfo(ctx context.Context) (GcpKeyInfo, bool) {
	var e GcpKeyInfo
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v GcpKeyInfo
	d := o.GcpKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CustomerManagedKey.
func (o *CustomerManagedKey) SetGcpKeyInfo(ctx context.Context, v GcpKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpKeyInfo = vs
}

// GetUseCases returns the value of the UseCases field in CustomerManagedKey as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey) GetUseCases(ctx context.Context) ([]types.String, bool) {
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

// SetUseCases sets the value of the UseCases field in CustomerManagedKey.
func (o *CustomerManagedKey) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UseCases = types.ListValueMust(t, vs)
}

type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (to *DeleteCredentialRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCredentialRequest) {
}

func (to *DeleteCredentialRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCredentialRequest) {
}

func (c DeleteCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialRequest
// only implements ToObjectValue() and Type().
func (o DeleteCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id": o.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id": types.StringType,
		},
	}
}

type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

func (to *DeleteEncryptionKeyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEncryptionKeyRequest) {
}

func (to *DeleteEncryptionKeyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteEncryptionKeyRequest) {
}

func (c DeleteEncryptionKeyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteEncryptionKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEncryptionKeyRequest
// only implements ToObjectValue() and Type().
func (o DeleteEncryptionKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": o.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteEncryptionKeyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_managed_key_id": types.StringType,
		},
	}
}

type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

func (to *DeleteNetworkRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNetworkRequest) {
}

func (to *DeleteNetworkRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteNetworkRequest) {
}

func (c DeleteNetworkRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkRequest
// only implements ToObjectValue() and Type().
func (o DeleteNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

type DeletePrivateAccesRequest struct {
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *DeletePrivateAccesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePrivateAccesRequest) {
}

func (to *DeletePrivateAccesRequest) SyncFieldsDuringRead(ctx context.Context, from DeletePrivateAccesRequest) {
}

func (c DeletePrivateAccesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeletePrivateAccesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateAccesRequest
// only implements ToObjectValue() and Type().
func (o DeletePrivateAccesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePrivateAccesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

type DeleteStorageRequest struct {
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (to *DeleteStorageRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteStorageRequest) {
}

func (to *DeleteStorageRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteStorageRequest) {
}

func (c DeleteStorageRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteStorageRequest
// only implements ToObjectValue() and Type().
func (o DeleteStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteStorageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

type DeleteVpcEndpointRequest struct {
	VpcEndpointId types.String `tfsdk:"-"`
}

func (to *DeleteVpcEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteVpcEndpointRequest) {
}

func (to *DeleteVpcEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteVpcEndpointRequest) {
}

func (c DeleteVpcEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (o DeleteVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteVpcEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

type DeleteWorkspaceRequest struct {
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceRequest) {
}

func (to *DeleteWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceRequest) {
}

func (c DeleteWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o DeleteWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type ExternalCustomerInfo struct {
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

func (to *ExternalCustomerInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalCustomerInfo) {
}

func (to *ExternalCustomerInfo) SyncFieldsDuringRead(ctx context.Context, from ExternalCustomerInfo) {
}

func (c ExternalCustomerInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExternalCustomerInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalCustomerInfo
// only implements ToObjectValue() and Type().
func (o ExternalCustomerInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExternalCustomerInfo) Type(ctx context.Context) attr.Type {
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
type GcpCommonNetworkConfig struct {
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	GkeClusterMasterIpRange types.String `tfsdk:"gke_cluster_master_ip_range"`
	// The type of network connectivity of the GKE cluster.
	GkeConnectivityType types.String `tfsdk:"gke_connectivity_type"`
}

func (to *GcpCommonNetworkConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpCommonNetworkConfig) {
}

func (to *GcpCommonNetworkConfig) SyncFieldsDuringRead(ctx context.Context, from GcpCommonNetworkConfig) {
}

func (c GcpCommonNetworkConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GcpCommonNetworkConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpCommonNetworkConfig
// only implements ToObjectValue() and Type().
func (o GcpCommonNetworkConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gke_cluster_master_ip_range": o.GkeClusterMasterIpRange,
			"gke_connectivity_type":       o.GkeConnectivityType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpCommonNetworkConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gke_cluster_master_ip_range": types.StringType,
			"gke_connectivity_type":       types.StringType,
		},
	}
}

type GcpKeyInfo struct {
	// Globally unique service account email that has access to the KMS key. The
	// service account exists within the Databricks CP project.
	GcpServiceAccount types.Object `tfsdk:"gcp_service_account"`
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId types.String `tfsdk:"kms_key_id"`
}

func (to *GcpKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpKeyInfo) {
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

func (to *GcpKeyInfo) SyncFieldsDuringRead(ctx context.Context, from GcpKeyInfo) {
	if !from.GcpServiceAccount.IsNull() && !from.GcpServiceAccount.IsUnknown() {
		if toGcpServiceAccount, ok := to.GetGcpServiceAccount(ctx); ok {
			if fromGcpServiceAccount, ok := from.GetGcpServiceAccount(ctx); ok {
				toGcpServiceAccount.SyncFieldsDuringRead(ctx, fromGcpServiceAccount)
				to.SetGcpServiceAccount(ctx, toGcpServiceAccount)
			}
		}
	}
}

func (c GcpKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_service_account"] = attrs["gcp_service_account"].SetOptional()
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
func (a GcpKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_service_account": reflect.TypeOf(GcpServiceAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpKeyInfo
// only implements ToObjectValue() and Type().
func (o GcpKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_service_account": o.GcpServiceAccount,
			"kms_key_id":          o.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_service_account": GcpServiceAccount{}.Type(ctx),
			"kms_key_id":          types.StringType,
		},
	}
}

// GetGcpServiceAccount returns the value of the GcpServiceAccount field in GcpKeyInfo as
// a GcpServiceAccount value.
// If the field is unknown or null, the boolean return value is false.
func (o *GcpKeyInfo) GetGcpServiceAccount(ctx context.Context) (GcpServiceAccount, bool) {
	var e GcpServiceAccount
	if o.GcpServiceAccount.IsNull() || o.GcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v GcpServiceAccount
	d := o.GcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpServiceAccount sets the value of the GcpServiceAccount field in GcpKeyInfo.
func (o *GcpKeyInfo) SetGcpServiceAccount(ctx context.Context, v GcpServiceAccount) {
	vs := v.ToObjectValue(ctx)
	o.GcpServiceAccount = vs
}

// The network configuration for the workspace.
type GcpManagedNetworkConfig struct {
	// The IP range that will be used to allocate GKE cluster Pods from.
	GkeClusterPodIpRange types.String `tfsdk:"gke_cluster_pod_ip_range"`
	// The IP range that will be used to allocate GKE cluster Services from.
	GkeClusterServiceIpRange types.String `tfsdk:"gke_cluster_service_ip_range"`
	// The IP range which will be used to allocate GKE cluster nodes from. Note:
	// Pods, services and master IP range must be mutually exclusive.
	SubnetCidr types.String `tfsdk:"subnet_cidr"`
}

func (to *GcpManagedNetworkConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpManagedNetworkConfig) {
}

func (to *GcpManagedNetworkConfig) SyncFieldsDuringRead(ctx context.Context, from GcpManagedNetworkConfig) {
}

func (c GcpManagedNetworkConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GcpManagedNetworkConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpManagedNetworkConfig
// only implements ToObjectValue() and Type().
func (o GcpManagedNetworkConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gke_cluster_pod_ip_range":     o.GkeClusterPodIpRange,
			"gke_cluster_service_ip_range": o.GkeClusterServiceIpRange,
			"subnet_cidr":                  o.SubnetCidr,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpManagedNetworkConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gke_cluster_pod_ip_range":     types.StringType,
			"gke_cluster_service_ip_range": types.StringType,
			"subnet_cidr":                  types.StringType,
		},
	}
}

type GcpNetworkInfo struct {
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

func (to *GcpNetworkInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpNetworkInfo) {
}

func (to *GcpNetworkInfo) SyncFieldsDuringRead(ctx context.Context, from GcpNetworkInfo) {
}

func (c GcpNetworkInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GcpNetworkInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpNetworkInfo
// only implements ToObjectValue() and Type().
func (o GcpNetworkInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GcpNetworkInfo) Type(ctx context.Context) attr.Type {
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

type GcpServiceAccount struct {
	ServiceAccountEmail types.String `tfsdk:"service_account_email"`
}

func (to *GcpServiceAccount) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpServiceAccount) {
}

func (to *GcpServiceAccount) SyncFieldsDuringRead(ctx context.Context, from GcpServiceAccount) {
}

func (c GcpServiceAccount) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GcpServiceAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpServiceAccount
// only implements ToObjectValue() and Type().
func (o GcpServiceAccount) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_account_email": o.ServiceAccountEmail,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpServiceAccount) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_account_email": types.StringType,
		},
	}
}

type GcpVpcEndpointInfo struct {
	EndpointRegion types.String `tfsdk:"endpoint_region"`

	ProjectId types.String `tfsdk:"project_id"`

	PscConnectionId types.String `tfsdk:"psc_connection_id"`

	PscEndpointName types.String `tfsdk:"psc_endpoint_name"`

	ServiceAttachmentId types.String `tfsdk:"service_attachment_id"`
}

func (to *GcpVpcEndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpVpcEndpointInfo) {
}

func (to *GcpVpcEndpointInfo) SyncFieldsDuringRead(ctx context.Context, from GcpVpcEndpointInfo) {
}

func (c GcpVpcEndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GcpVpcEndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpVpcEndpointInfo
// only implements ToObjectValue() and Type().
func (o GcpVpcEndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GcpVpcEndpointInfo) Type(ctx context.Context) attr.Type {
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

type GetCredentialRequest struct {
	// Credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (to *GetCredentialRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCredentialRequest) {
}

func (to *GetCredentialRequest) SyncFieldsDuringRead(ctx context.Context, from GetCredentialRequest) {
}

func (c GetCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialRequest
// only implements ToObjectValue() and Type().
func (o GetCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id": o.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id": types.StringType,
		},
	}
}

type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

func (to *GetEncryptionKeyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEncryptionKeyRequest) {
}

func (to *GetEncryptionKeyRequest) SyncFieldsDuringRead(ctx context.Context, from GetEncryptionKeyRequest) {
}

func (c GetEncryptionKeyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetEncryptionKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEncryptionKeyRequest
// only implements ToObjectValue() and Type().
func (o GetEncryptionKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": o.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEncryptionKeyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_managed_key_id": types.StringType,
		},
	}
}

type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

func (to *GetNetworkRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNetworkRequest) {
}

func (to *GetNetworkRequest) SyncFieldsDuringRead(ctx context.Context, from GetNetworkRequest) {
}

func (c GetNetworkRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkRequest
// only implements ToObjectValue() and Type().
func (o GetNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNetworkRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

type GetPrivateAccesRequest struct {
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *GetPrivateAccesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPrivateAccesRequest) {
}

func (to *GetPrivateAccesRequest) SyncFieldsDuringRead(ctx context.Context, from GetPrivateAccesRequest) {
}

func (c GetPrivateAccesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetPrivateAccesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateAccesRequest
// only implements ToObjectValue() and Type().
func (o GetPrivateAccesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPrivateAccesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

type GetStorageRequest struct {
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (to *GetStorageRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStorageRequest) {
}

func (to *GetStorageRequest) SyncFieldsDuringRead(ctx context.Context, from GetStorageRequest) {
}

func (c GetStorageRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStorageRequest
// only implements ToObjectValue() and Type().
func (o GetStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStorageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId types.String `tfsdk:"-"`
}

func (to *GetVpcEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetVpcEndpointRequest) {
}

func (to *GetVpcEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from GetVpcEndpointRequest) {
}

func (c GetVpcEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (o GetVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetVpcEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

type GetWorkspaceRequest struct {
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceRequest) {
}

func (to *GetWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceRequest) {
}

func (c GetWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// The configurations of the GKE cluster used by the GCP workspace.
type GkeConfig struct {
	// The type of network connectivity of the GKE cluster.
	ConnectivityType types.String `tfsdk:"connectivity_type"`
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	MasterIpRange types.String `tfsdk:"master_ip_range"`
}

func (to *GkeConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GkeConfig) {
}

func (to *GkeConfig) SyncFieldsDuringRead(ctx context.Context, from GkeConfig) {
}

func (c GkeConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GkeConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GkeConfig
// only implements ToObjectValue() and Type().
func (o GkeConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connectivity_type": o.ConnectivityType,
			"master_ip_range":   o.MasterIpRange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GkeConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connectivity_type": types.StringType,
			"master_ip_range":   types.StringType,
		},
	}
}

// The credential ID that is used to access the key vault.
type KeyAccessConfiguration struct {
	CredentialId types.String `tfsdk:"credential_id"`
}

func (to *KeyAccessConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KeyAccessConfiguration) {
}

func (to *KeyAccessConfiguration) SyncFieldsDuringRead(ctx context.Context, from KeyAccessConfiguration) {
}

func (c KeyAccessConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a KeyAccessConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KeyAccessConfiguration
// only implements ToObjectValue() and Type().
func (o KeyAccessConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o KeyAccessConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.StringType,
		},
	}
}

type ListCredentialsRequest struct {
}

func (to *ListCredentialsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCredentialsRequest) {
}

func (to *ListCredentialsRequest) SyncFieldsDuringRead(ctx context.Context, from ListCredentialsRequest) {
}

func (c ListCredentialsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsRequest
// only implements ToObjectValue() and Type().
func (o ListCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListEncryptionKeysRequest struct {
}

func (to *ListEncryptionKeysRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEncryptionKeysRequest) {
}

func (to *ListEncryptionKeysRequest) SyncFieldsDuringRead(ctx context.Context, from ListEncryptionKeysRequest) {
}

func (c ListEncryptionKeysRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListEncryptionKeysRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEncryptionKeysRequest
// only implements ToObjectValue() and Type().
func (o ListEncryptionKeysRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListEncryptionKeysRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNetworksRequest struct {
}

func (to *ListNetworksRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworksRequest) {
}

func (to *ListNetworksRequest) SyncFieldsDuringRead(ctx context.Context, from ListNetworksRequest) {
}

func (c ListNetworksRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNetworksRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworksRequest
// only implements ToObjectValue() and Type().
func (o ListNetworksRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworksRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListPrivateAccessRequest struct {
}

func (to *ListPrivateAccessRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateAccessRequest) {
}

func (to *ListPrivateAccessRequest) SyncFieldsDuringRead(ctx context.Context, from ListPrivateAccessRequest) {
}

func (c ListPrivateAccessRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListPrivateAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateAccessRequest
// only implements ToObjectValue() and Type().
func (o ListPrivateAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListPrivateAccessRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListStorageRequest struct {
}

func (to *ListStorageRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListStorageRequest) {
}

func (to *ListStorageRequest) SyncFieldsDuringRead(ctx context.Context, from ListStorageRequest) {
}

func (c ListStorageRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStorageRequest
// only implements ToObjectValue() and Type().
func (o ListStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListStorageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListVpcEndpointsRequest struct {
}

func (to *ListVpcEndpointsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVpcEndpointsRequest) {
}

func (to *ListVpcEndpointsRequest) SyncFieldsDuringRead(ctx context.Context, from ListVpcEndpointsRequest) {
}

func (c ListVpcEndpointsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListVpcEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVpcEndpointsRequest
// only implements ToObjectValue() and Type().
func (o ListVpcEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListVpcEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListWorkspacesRequest struct {
}

func (to *ListWorkspacesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspacesRequest) {
}

func (to *ListWorkspacesRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspacesRequest) {
}

func (c ListWorkspacesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListWorkspacesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspacesRequest
// only implements ToObjectValue() and Type().
func (o ListWorkspacesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspacesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	AccountId types.String `tfsdk:"account_id"`

	AwsNetworkInfo types.Object `tfsdk:"aws_network_info"`
	// Time in epoch milliseconds when the network was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Array of error messages about the network configuration.
	ErrorMessages types.List `tfsdk:"error_messages"`

	GcpNetworkInfo types.Object `tfsdk:"gcp_network_info"`
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

	VpcEndpoints types.Object `tfsdk:"vpc_endpoints"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId types.String `tfsdk:"vpc_id"`

	VpcStatus types.String `tfsdk:"vpc_status"`
	// Array of warning messages about the network configuration.
	WarningMessages types.List `tfsdk:"warning_messages"`
	// Workspace ID associated with this network configuration.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *Network) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Network) {
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

func (to *Network) SyncFieldsDuringRead(ctx context.Context, from Network) {
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

func (c Network) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_network_info"] = attrs["aws_network_info"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["error_messages"] = attrs["error_messages"].SetComputed()
	attrs["gcp_network_info"] = attrs["gcp_network_info"].SetOptional()
	attrs["network_id"] = attrs["network_id"].SetOptional()
	attrs["network_name"] = attrs["network_name"].SetOptional()
	attrs["security_group_ids"] = attrs["security_group_ids"].SetOptional()
	attrs["subnet_ids"] = attrs["subnet_ids"].SetOptional()
	attrs["vpc_endpoints"] = attrs["vpc_endpoints"].SetComputed()
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
func (a Network) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_network_info":   reflect.TypeOf(AwsNetworkInfo{}),
		"error_messages":     reflect.TypeOf(NetworkHealth{}),
		"gcp_network_info":   reflect.TypeOf(GcpNetworkInfo{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
		"subnet_ids":         reflect.TypeOf(types.String{}),
		"vpc_endpoints":      reflect.TypeOf(NetworkVpcEndpoints{}),
		"warning_messages":   reflect.TypeOf(NetworkWarning{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Network
// only implements ToObjectValue() and Type().
func (o Network) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Network) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":       types.StringType,
			"aws_network_info": AwsNetworkInfo{}.Type(ctx),
			"creation_time":    types.Int64Type,
			"error_messages": basetypes.ListType{
				ElemType: NetworkHealth{}.Type(ctx),
			},
			"gcp_network_info": GcpNetworkInfo{}.Type(ctx),
			"network_id":       types.StringType,
			"network_name":     types.StringType,
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"vpc_endpoints": NetworkVpcEndpoints{}.Type(ctx),
			"vpc_id":        types.StringType,
			"vpc_status":    types.StringType,
			"warning_messages": basetypes.ListType{
				ElemType: NetworkWarning{}.Type(ctx),
			},
			"workspace_id": types.Int64Type,
		},
	}
}

// GetAwsNetworkInfo returns the value of the AwsNetworkInfo field in Network as
// a AwsNetworkInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetAwsNetworkInfo(ctx context.Context) (AwsNetworkInfo, bool) {
	var e AwsNetworkInfo
	if o.AwsNetworkInfo.IsNull() || o.AwsNetworkInfo.IsUnknown() {
		return e, false
	}
	var v AwsNetworkInfo
	d := o.AwsNetworkInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsNetworkInfo sets the value of the AwsNetworkInfo field in Network.
func (o *Network) SetAwsNetworkInfo(ctx context.Context, v AwsNetworkInfo) {
	vs := v.ToObjectValue(ctx)
	o.AwsNetworkInfo = vs
}

// GetErrorMessages returns the value of the ErrorMessages field in Network as
// a slice of NetworkHealth values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetErrorMessages(ctx context.Context) ([]NetworkHealth, bool) {
	if o.ErrorMessages.IsNull() || o.ErrorMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkHealth
	d := o.ErrorMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetErrorMessages sets the value of the ErrorMessages field in Network.
func (o *Network) SetErrorMessages(ctx context.Context, v []NetworkHealth) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ErrorMessages = types.ListValueMust(t, vs)
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in Network as
// a GcpNetworkInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo, bool) {
	var e GcpNetworkInfo
	if o.GcpNetworkInfo.IsNull() || o.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v GcpNetworkInfo
	d := o.GcpNetworkInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in Network.
func (o *Network) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpNetworkInfo = vs
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in Network as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
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

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in Network.
func (o *Network) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in Network as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
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

// SetSubnetIds sets the value of the SubnetIds field in Network.
func (o *Network) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in Network as
// a NetworkVpcEndpoints value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints, bool) {
	var e NetworkVpcEndpoints
	if o.VpcEndpoints.IsNull() || o.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v NetworkVpcEndpoints
	d := o.VpcEndpoints.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in Network.
func (o *Network) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints) {
	vs := v.ToObjectValue(ctx)
	o.VpcEndpoints = vs
}

// GetWarningMessages returns the value of the WarningMessages field in Network as
// a slice of NetworkWarning values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetWarningMessages(ctx context.Context) ([]NetworkWarning, bool) {
	if o.WarningMessages.IsNull() || o.WarningMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkWarning
	d := o.WarningMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarningMessages sets the value of the WarningMessages field in Network.
func (o *Network) SetWarningMessages(ctx context.Context, v []NetworkWarning) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["warning_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WarningMessages = types.ListValueMust(t, vs)
}

type NetworkHealth struct {
	// Details of the error.
	ErrorMessage types.String `tfsdk:"error_message"`

	ErrorType types.String `tfsdk:"error_type"`
}

func (to *NetworkHealth) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkHealth) {
}

func (to *NetworkHealth) SyncFieldsDuringRead(ctx context.Context, from NetworkHealth) {
}

func (c NetworkHealth) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NetworkHealth) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkHealth
// only implements ToObjectValue() and Type().
func (o NetworkHealth) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_message": o.ErrorMessage,
			"error_type":    o.ErrorType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkHealth) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_message": types.StringType,
			"error_type":    types.StringType,
		},
	}
}

type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay types.List `tfsdk:"dataplane_relay"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi types.List `tfsdk:"rest_api"`
}

func (to *NetworkVpcEndpoints) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkVpcEndpoints) {
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

func (to *NetworkVpcEndpoints) SyncFieldsDuringRead(ctx context.Context, from NetworkVpcEndpoints) {
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

func (c NetworkVpcEndpoints) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NetworkVpcEndpoints) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataplane_relay": reflect.TypeOf(types.String{}),
		"rest_api":        reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkVpcEndpoints
// only implements ToObjectValue() and Type().
func (o NetworkVpcEndpoints) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataplane_relay": o.DataplaneRelay,
			"rest_api":        o.RestApi,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkVpcEndpoints) Type(ctx context.Context) attr.Type {
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

// GetDataplaneRelay returns the value of the DataplaneRelay field in NetworkVpcEndpoints as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkVpcEndpoints) GetDataplaneRelay(ctx context.Context) ([]types.String, bool) {
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

// SetDataplaneRelay sets the value of the DataplaneRelay field in NetworkVpcEndpoints.
func (o *NetworkVpcEndpoints) SetDataplaneRelay(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataplane_relay"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataplaneRelay = types.ListValueMust(t, vs)
}

// GetRestApi returns the value of the RestApi field in NetworkVpcEndpoints as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkVpcEndpoints) GetRestApi(ctx context.Context) ([]types.String, bool) {
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

// SetRestApi sets the value of the RestApi field in NetworkVpcEndpoints.
func (o *NetworkVpcEndpoints) SetRestApi(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rest_api"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RestApi = types.ListValueMust(t, vs)
}

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage types.String `tfsdk:"warning_message"`

	WarningType types.String `tfsdk:"warning_type"`
}

func (to *NetworkWarning) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkWarning) {
}

func (to *NetworkWarning) SyncFieldsDuringRead(ctx context.Context, from NetworkWarning) {
}

func (c NetworkWarning) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NetworkWarning) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkWarning
// only implements ToObjectValue() and Type().
func (o NetworkWarning) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warning_message": o.WarningMessage,
			"warning_type":    o.WarningType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkWarning) Type(ctx context.Context) attr.Type {
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
type NpipVpcInfra struct {
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

func (to *NpipVpcInfra) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NpipVpcInfra) {
}

func (to *NpipVpcInfra) SyncFieldsDuringRead(ctx context.Context, from NpipVpcInfra) {
}

func (c NpipVpcInfra) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NpipVpcInfra) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NpipVpcInfra
// only implements ToObjectValue() and Type().
func (o NpipVpcInfra) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o NpipVpcInfra) Type(ctx context.Context) attr.Type {
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
type PrivateAccessSettings struct {
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

func (to *PrivateAccessSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateAccessSettings) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (to *PrivateAccessSettings) SyncFieldsDuringRead(ctx context.Context, from PrivateAccessSettings) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (c PrivateAccessSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PrivateAccessSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateAccessSettings
// only implements ToObjectValue() and Type().
func (o PrivateAccessSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PrivateAccessSettings) Type(ctx context.Context) attr.Type {
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

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in PrivateAccessSettings as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PrivateAccessSettings) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in PrivateAccessSettings.
func (o *PrivateAccessSettings) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type ReplacePrivateAccessSettingsRequest struct {
	CustomerFacingPrivateAccessSettings types.Object `tfsdk:"customer_facing_private_access_settings"`
	// The ID in the MWS API of the Private Access Settings.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *ReplacePrivateAccessSettingsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReplacePrivateAccessSettingsRequest) {
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

func (to *ReplacePrivateAccessSettingsRequest) SyncFieldsDuringRead(ctx context.Context, from ReplacePrivateAccessSettingsRequest) {
	if !from.CustomerFacingPrivateAccessSettings.IsNull() && !from.CustomerFacingPrivateAccessSettings.IsUnknown() {
		if toCustomerFacingPrivateAccessSettings, ok := to.GetCustomerFacingPrivateAccessSettings(ctx); ok {
			if fromCustomerFacingPrivateAccessSettings, ok := from.GetCustomerFacingPrivateAccessSettings(ctx); ok {
				toCustomerFacingPrivateAccessSettings.SyncFieldsDuringRead(ctx, fromCustomerFacingPrivateAccessSettings)
				to.SetCustomerFacingPrivateAccessSettings(ctx, toCustomerFacingPrivateAccessSettings)
			}
		}
	}
}

func (c ReplacePrivateAccessSettingsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["customer_facing_private_access_settings"] = attrs["customer_facing_private_access_settings"].SetRequired()
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
func (a ReplacePrivateAccessSettingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"customer_facing_private_access_settings": reflect.TypeOf(PrivateAccessSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplacePrivateAccessSettingsRequest
// only implements ToObjectValue() and Type().
func (o ReplacePrivateAccessSettingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_facing_private_access_settings": o.CustomerFacingPrivateAccessSettings,
			"private_access_settings_id":              o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ReplacePrivateAccessSettingsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_facing_private_access_settings": PrivateAccessSettings{}.Type(ctx),
			"private_access_settings_id":              types.StringType,
		},
	}
}

// GetCustomerFacingPrivateAccessSettings returns the value of the CustomerFacingPrivateAccessSettings field in ReplacePrivateAccessSettingsRequest as
// a PrivateAccessSettings value.
// If the field is unknown or null, the boolean return value is false.
func (o *ReplacePrivateAccessSettingsRequest) GetCustomerFacingPrivateAccessSettings(ctx context.Context) (PrivateAccessSettings, bool) {
	var e PrivateAccessSettings
	if o.CustomerFacingPrivateAccessSettings.IsNull() || o.CustomerFacingPrivateAccessSettings.IsUnknown() {
		return e, false
	}
	var v PrivateAccessSettings
	d := o.CustomerFacingPrivateAccessSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomerFacingPrivateAccessSettings sets the value of the CustomerFacingPrivateAccessSettings field in ReplacePrivateAccessSettingsRequest.
func (o *ReplacePrivateAccessSettingsRequest) SetCustomerFacingPrivateAccessSettings(ctx context.Context, v PrivateAccessSettings) {
	vs := v.ToObjectValue(ctx)
	o.CustomerFacingPrivateAccessSettings = vs
}

type RootBucketInfo struct {
	// Name of the bucket
	BucketName types.String `tfsdk:"bucket_name"`
}

func (to *RootBucketInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RootBucketInfo) {
}

func (to *RootBucketInfo) SyncFieldsDuringRead(ctx context.Context, from RootBucketInfo) {
}

func (c RootBucketInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RootBucketInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RootBucketInfo
// only implements ToObjectValue() and Type().
func (o RootBucketInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bucket_name": o.BucketName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RootBucketInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bucket_name": types.StringType,
		},
	}
}

type StorageConfiguration struct {
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

	RootBucketInfo types.Object `tfsdk:"root_bucket_info"`

	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`

	StorageConfigurationName types.String `tfsdk:"storage_configuration_name"`
}

func (to *StorageConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StorageConfiguration) {
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

func (to *StorageConfiguration) SyncFieldsDuringRead(ctx context.Context, from StorageConfiguration) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				toRootBucketInfo.SyncFieldsDuringRead(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (c StorageConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["role_arn"] = attrs["role_arn"].SetOptional()
	attrs["root_bucket_info"] = attrs["root_bucket_info"].SetOptional()
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
func (a StorageConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StorageConfiguration
// only implements ToObjectValue() and Type().
func (o StorageConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o StorageConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":                 types.StringType,
			"creation_time":              types.Int64Type,
			"role_arn":                   types.StringType,
			"root_bucket_info":           RootBucketInfo{}.Type(ctx),
			"storage_configuration_id":   types.StringType,
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in StorageConfiguration as
// a RootBucketInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageConfiguration) GetRootBucketInfo(ctx context.Context) (RootBucketInfo, bool) {
	var e RootBucketInfo
	if o.RootBucketInfo.IsNull() || o.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v RootBucketInfo
	d := o.RootBucketInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in StorageConfiguration.
func (o *StorageConfiguration) SetRootBucketInfo(ctx context.Context, v RootBucketInfo) {
	vs := v.ToObjectValue(ctx)
	o.RootBucketInfo = vs
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
type StsRole struct {
	// Note: This must match the external_id on the parent object.
	//
	// TODO(j): Add validation to ensure this cannot be updated. If the user can
	// override the external_id, that defeats the purpose.
	ExternalId types.String `tfsdk:"external_id"`

	RoleArn types.String `tfsdk:"role_arn"`
}

func (to *StsRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StsRole) {
}

func (to *StsRole) SyncFieldsDuringRead(ctx context.Context, from StsRole) {
}

func (c StsRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a StsRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StsRole
// only implements ToObjectValue() and Type().
func (o StsRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
			"role_arn":    o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StsRole) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
			"role_arn":    types.StringType,
		},
	}
}

// Describes a single subnet, which is associated with a particular AWS AZ and a
// particular address space which is a subset of the overall vpc_address_space.
type SubnetInfo struct {
	// Example: us-west-2a
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	// Example: 10.0.0.0/17.
	SubnetAddressSpace types.String `tfsdk:"subnet_address_space"`

	SubnetId types.String `tfsdk:"subnet_id"`
}

func (to *SubnetInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SubnetInfo) {
}

func (to *SubnetInfo) SyncFieldsDuringRead(ctx context.Context, from SubnetInfo) {
}

func (c SubnetInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SubnetInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubnetInfo
// only implements ToObjectValue() and Type().
func (o SubnetInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"availability_zone":    o.AvailabilityZone,
			"subnet_address_space": o.SubnetAddressSpace,
			"subnet_id":            o.SubnetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubnetInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"availability_zone":    types.StringType,
			"subnet_address_space": types.StringType,
			"subnet_id":            types.StringType,
		},
	}
}

type UpdateWorkspaceRequest struct {
	CustomerFacingWorkspace types.Object `tfsdk:"customer_facing_workspace"`
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

func (to *UpdateWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceRequest) {
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

func (to *UpdateWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceRequest) {
	if !from.CustomerFacingWorkspace.IsNull() && !from.CustomerFacingWorkspace.IsUnknown() {
		if toCustomerFacingWorkspace, ok := to.GetCustomerFacingWorkspace(ctx); ok {
			if fromCustomerFacingWorkspace, ok := from.GetCustomerFacingWorkspace(ctx); ok {
				toCustomerFacingWorkspace.SyncFieldsDuringRead(ctx, fromCustomerFacingWorkspace)
				to.SetCustomerFacingWorkspace(ctx, toCustomerFacingWorkspace)
			}
		}
	}
}

func (c UpdateWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["customer_facing_workspace"] = attrs["customer_facing_workspace"].SetRequired()
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
func (a UpdateWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"customer_facing_workspace": reflect.TypeOf(Workspace{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_facing_workspace": o.CustomerFacingWorkspace,
			"update_mask":               o.UpdateMask,
			"workspace_id":              o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_facing_workspace": Workspace{}.Type(ctx),
			"update_mask":               types.StringType,
			"workspace_id":              types.Int64Type,
		},
	}
}

// GetCustomerFacingWorkspace returns the value of the CustomerFacingWorkspace field in UpdateWorkspaceRequest as
// a Workspace value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceRequest) GetCustomerFacingWorkspace(ctx context.Context) (Workspace, bool) {
	var e Workspace
	if o.CustomerFacingWorkspace.IsNull() || o.CustomerFacingWorkspace.IsUnknown() {
		return e, false
	}
	var v Workspace
	d := o.CustomerFacingWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomerFacingWorkspace sets the value of the CustomerFacingWorkspace field in UpdateWorkspaceRequest.
func (o *UpdateWorkspaceRequest) SetCustomerFacingWorkspace(ctx context.Context, v Workspace) {
	vs := v.ToObjectValue(ctx)
	o.CustomerFacingWorkspace = vs
}

// *
type VpcEndpoint struct {
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
	GcpVpcEndpointInfo types.Object `tfsdk:"gcp_vpc_endpoint_info"`
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

func (to *VpcEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VpcEndpoint) {
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

func (to *VpcEndpoint) SyncFieldsDuringRead(ctx context.Context, from VpcEndpoint) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				toGcpVpcEndpointInfo.SyncFieldsDuringRead(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (c VpcEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_account_id"] = attrs["aws_account_id"].SetOptional()
	attrs["aws_endpoint_service_id"] = attrs["aws_endpoint_service_id"].SetOptional()
	attrs["aws_vpc_endpoint_id"] = attrs["aws_vpc_endpoint_id"].SetOptional()
	attrs["gcp_vpc_endpoint_info"] = attrs["gcp_vpc_endpoint_info"].SetOptional()
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
func (a VpcEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VpcEndpoint
// only implements ToObjectValue() and Type().
func (o VpcEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o VpcEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":              types.StringType,
			"aws_account_id":          types.StringType,
			"aws_endpoint_service_id": types.StringType,
			"aws_vpc_endpoint_id":     types.StringType,
			"gcp_vpc_endpoint_info":   GcpVpcEndpointInfo{}.Type(ctx),
			"region":                  types.StringType,
			"state":                   types.StringType,
			"use_case":                types.StringType,
			"vpc_endpoint_id":         types.StringType,
			"vpc_endpoint_name":       types.StringType,
		},
	}
}

// GetGcpVpcEndpointInfo returns the value of the GcpVpcEndpointInfo field in VpcEndpoint as
// a GcpVpcEndpointInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *VpcEndpoint) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo, bool) {
	var e GcpVpcEndpointInfo
	if o.GcpVpcEndpointInfo.IsNull() || o.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v GcpVpcEndpointInfo
	d := o.GcpVpcEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in VpcEndpoint.
func (o *VpcEndpoint) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpVpcEndpointInfo = vs
}

type Workspace struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id"`

	AwsRegion types.String `tfsdk:"aws_region"`

	AzureWorkspaceInfo types.Object `tfsdk:"azure_workspace_info"`
	// The cloud name. This field always has the value `gcp`.
	Cloud types.String `tfsdk:"cloud"`

	CloudResourceContainer types.Object `tfsdk:"cloud_resource_container"`
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
	ExternalCustomerInfo types.Object `tfsdk:"external_customer_info"`

	GcpManagedNetworkConfig types.Object `tfsdk:"gcp_managed_network_config"`

	GkeConfig types.Object `tfsdk:"gke_config"`
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
	Network types.Object `tfsdk:"network"`
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

func (to *Workspace) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Workspace) {
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

func (to *Workspace) SyncFieldsDuringRead(ctx context.Context, from Workspace) {
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

func (c Workspace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_region"] = attrs["aws_region"].SetOptional()
	attrs["azure_workspace_info"] = attrs["azure_workspace_info"].SetComputed()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].SetOptional()
	attrs["compute_mode"] = attrs["compute_mode"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["deployment_name"] = attrs["deployment_name"].SetOptional()
	attrs["external_customer_info"] = attrs["external_customer_info"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["gke_config"] = attrs["gke_config"].SetOptional()
	attrs["is_no_public_ip_enabled"] = attrs["is_no_public_ip_enabled"].SetOptional()
	attrs["location"] = attrs["location"].SetOptional()
	attrs["managed_services_customer_managed_key_id"] = attrs["managed_services_customer_managed_key_id"].SetOptional()
	attrs["network"] = attrs["network"].SetOptional()
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
func (a Workspace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_workspace_info":       reflect.TypeOf(AzureWorkspaceInfo{}),
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"external_customer_info":     reflect.TypeOf(ExternalCustomerInfo{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig{}),
		"gke_config":                 reflect.TypeOf(GkeConfig{}),
		"network":                    reflect.TypeOf(WorkspaceNetwork{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Workspace
// only implements ToObjectValue() and Type().
func (o Workspace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Workspace) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":               types.StringType,
			"aws_region":               types.StringType,
			"azure_workspace_info":     AzureWorkspaceInfo{}.Type(ctx),
			"cloud":                    types.StringType,
			"cloud_resource_container": CloudResourceContainer{}.Type(ctx),
			"compute_mode":             types.StringType,
			"creation_time":            types.Int64Type,
			"credentials_id":           types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name":                          types.StringType,
			"external_customer_info":                   ExternalCustomerInfo{}.Type(ctx),
			"gcp_managed_network_config":               GcpManagedNetworkConfig{}.Type(ctx),
			"gke_config":                               GkeConfig{}.Type(ctx),
			"is_no_public_ip_enabled":                  types.BoolType,
			"location":                                 types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network":                                  WorkspaceNetwork{}.Type(ctx),
			"network_connectivity_config_id":           types.StringType,
			"network_id":                               types.StringType,
			"pricing_tier":                             types.StringType,
			"private_access_settings_id":               types.StringType,
			"storage_configuration_id":                 types.StringType,
			"storage_customer_managed_key_id":          types.StringType,
			"storage_mode":                             types.StringType,
			"workspace_id":                             types.Int64Type,
			"workspace_name":                           types.StringType,
			"workspace_status":                         types.StringType,
			"workspace_status_message":                 types.StringType,
		},
	}
}

// GetAzureWorkspaceInfo returns the value of the AzureWorkspaceInfo field in Workspace as
// a AzureWorkspaceInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetAzureWorkspaceInfo(ctx context.Context) (AzureWorkspaceInfo, bool) {
	var e AzureWorkspaceInfo
	if o.AzureWorkspaceInfo.IsNull() || o.AzureWorkspaceInfo.IsUnknown() {
		return e, false
	}
	var v AzureWorkspaceInfo
	d := o.AzureWorkspaceInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureWorkspaceInfo sets the value of the AzureWorkspaceInfo field in Workspace.
func (o *Workspace) SetAzureWorkspaceInfo(ctx context.Context, v AzureWorkspaceInfo) {
	vs := v.ToObjectValue(ctx)
	o.AzureWorkspaceInfo = vs
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in Workspace as
// a CloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer, bool) {
	var e CloudResourceContainer
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v CloudResourceContainer
	d := o.CloudResourceContainer.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in Workspace.
func (o *Workspace) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	o.CloudResourceContainer = vs
}

// GetCustomTags returns the value of the CustomTags field in Workspace as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in Workspace.
func (o *Workspace) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetExternalCustomerInfo returns the value of the ExternalCustomerInfo field in Workspace as
// a ExternalCustomerInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetExternalCustomerInfo(ctx context.Context) (ExternalCustomerInfo, bool) {
	var e ExternalCustomerInfo
	if o.ExternalCustomerInfo.IsNull() || o.ExternalCustomerInfo.IsUnknown() {
		return e, false
	}
	var v ExternalCustomerInfo
	d := o.ExternalCustomerInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalCustomerInfo sets the value of the ExternalCustomerInfo field in Workspace.
func (o *Workspace) SetExternalCustomerInfo(ctx context.Context, v ExternalCustomerInfo) {
	vs := v.ToObjectValue(ctx)
	o.ExternalCustomerInfo = vs
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in Workspace as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpManagedNetworkConfig
	d := o.GcpManagedNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in Workspace.
func (o *Workspace) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	o.GcpManagedNetworkConfig = vs
}

// GetGkeConfig returns the value of the GkeConfig field in Workspace as
// a GkeConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetGkeConfig(ctx context.Context) (GkeConfig, bool) {
	var e GkeConfig
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v GkeConfig
	d := o.GkeConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGkeConfig sets the value of the GkeConfig field in Workspace.
func (o *Workspace) SetGkeConfig(ctx context.Context, v GkeConfig) {
	vs := v.ToObjectValue(ctx)
	o.GkeConfig = vs
}

// GetNetwork returns the value of the Network field in Workspace as
// a WorkspaceNetwork value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetNetwork(ctx context.Context) (WorkspaceNetwork, bool) {
	var e WorkspaceNetwork
	if o.Network.IsNull() || o.Network.IsUnknown() {
		return e, false
	}
	var v WorkspaceNetwork
	d := o.Network.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNetwork sets the value of the Network field in Workspace.
func (o *Workspace) SetNetwork(ctx context.Context, v WorkspaceNetwork) {
	vs := v.ToObjectValue(ctx)
	o.Network = vs
}

// The network configuration for workspaces.
type WorkspaceNetwork struct {
	// The shared network config for GCP workspace. This object has common
	// network configurations that are network attributions of a workspace. This
	// object is input-only.
	GcpCommonNetworkConfig types.Object `tfsdk:"gcp_common_network_config"`
	// The mutually exclusive network deployment modes. The option decides which
	// network mode the workspace will use. The network config for GCP workspace
	// with Databricks managed network. This object is input-only and will not
	// be provided when listing workspaces. See go/gcp-byovpc-alpha-design for
	// interface decisions.
	GcpManagedNetworkConfig types.Object `tfsdk:"gcp_managed_network_config"`
	// The ID of the network object, if the workspace is a BYOVPC workspace.
	// This should apply to workspaces on all clouds in internal services. In
	// accounts-rest-api, user will use workspace.network_id for input and
	// output instead. Currently (2021-06-19) the network ID is only used by
	// GCP.
	NetworkId types.String `tfsdk:"network_id"`
}

func (to *WorkspaceNetwork) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceNetwork) {
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

func (to *WorkspaceNetwork) SyncFieldsDuringRead(ctx context.Context, from WorkspaceNetwork) {
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

func (c WorkspaceNetwork) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_common_network_config"] = attrs["gcp_common_network_config"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
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
func (a WorkspaceNetwork) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_common_network_config":  reflect.TypeOf(GcpCommonNetworkConfig{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceNetwork
// only implements ToObjectValue() and Type().
func (o WorkspaceNetwork) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_common_network_config":  o.GcpCommonNetworkConfig,
			"gcp_managed_network_config": o.GcpManagedNetworkConfig,
			"network_id":                 o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceNetwork) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_common_network_config":  GcpCommonNetworkConfig{}.Type(ctx),
			"gcp_managed_network_config": GcpManagedNetworkConfig{}.Type(ctx),
			"network_id":                 types.StringType,
		},
	}
}

// GetGcpCommonNetworkConfig returns the value of the GcpCommonNetworkConfig field in WorkspaceNetwork as
// a GcpCommonNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceNetwork) GetGcpCommonNetworkConfig(ctx context.Context) (GcpCommonNetworkConfig, bool) {
	var e GcpCommonNetworkConfig
	if o.GcpCommonNetworkConfig.IsNull() || o.GcpCommonNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpCommonNetworkConfig
	d := o.GcpCommonNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpCommonNetworkConfig sets the value of the GcpCommonNetworkConfig field in WorkspaceNetwork.
func (o *WorkspaceNetwork) SetGcpCommonNetworkConfig(ctx context.Context, v GcpCommonNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	o.GcpCommonNetworkConfig = vs
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in WorkspaceNetwork as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceNetwork) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpManagedNetworkConfig
	d := o.GcpManagedNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in WorkspaceNetwork.
func (o *WorkspaceNetwork) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	o.GcpManagedNetworkConfig = vs
}
