package gcp

import (
	"context"
	"reflect"

	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const vpcDataSourceName = "gcp_vpc_policy"

func DataSourceGcpVpcPolicy() datasource.DataSource {
	return &GcpVpcPolicyDataSource{}
}

var _ datasource.DataSourceWithConfigure = &GcpVpcPolicyDataSource{}

type GcpVpcPolicyDataSource struct {
	Client *common.DatabricksClient
}

// GcpVpcPolicyData holds the input/output for the VPC project custom role.
// In shared VPC setups, the VPC may reside in a different GCP project than the
// workspace. This data source covers the permissions required in that VPC project.
// See https://docs.databricks.com/gcp/en/admin/cloud-configurations/gcp/permissions
type GcpVpcPolicyData struct {
	EnableByovpc types.Bool `tfsdk:"enable_byovpc"`
	EnableCmk    types.Bool `tfsdk:"enable_cmk"`
	EnablePsc    types.Bool `tfsdk:"enable_psc"`
	Permissions  types.List `tfsdk:"permissions"`
	tfschema.Namespace
}

func (GcpVpcPolicyData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enable_byovpc"] = attrs["enable_byovpc"].SetOptional()
	attrs["enable_cmk"] = attrs["enable_cmk"].SetOptional()
	attrs["enable_psc"] = attrs["enable_psc"].SetOptional()
	attrs["permissions"] = attrs["permissions"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (GcpVpcPolicyData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions":     reflect.TypeOf(types.String{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *GcpVpcPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(vpcDataSourceName)
}

func (d *GcpVpcPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, GcpVpcPolicyData{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *GcpVpcPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *GcpVpcPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, vpcDataSourceName)

	var data GcpVpcPolicyData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	enableByovpc := !data.EnableByovpc.IsNull() && data.EnableByovpc.ValueBool()
	enableCmk := !data.EnableCmk.IsNull() && data.EnableCmk.ValueBool()
	enablePsc := !data.EnablePsc.IsNull() && data.EnablePsc.ValueBool()
	permissions := computeGcpVpcPermissions(enableByovpc, enableCmk, enablePsc)

	permValues := make([]attr.Value, len(permissions))
	for i, p := range permissions {
		permValues[i] = types.StringValue(p)
	}
	data.Permissions = types.ListValueMust(types.StringType, permValues)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func computeGcpVpcPermissions(enableByovpc, enableCmk, enablePsc bool) []string {
	permissions := make([]string, len(gcpVpcBasePermissions))
	copy(permissions, gcpVpcBasePermissions)
	if enableByovpc {
		permissions = append(permissions, gcpByovpcPermissions...)
	}
	if enableCmk {
		permissions = append(permissions, gcpCmkPermissions...)
	}
	if enablePsc {
		permissions = append(permissions, gcpPscPermissions...)
	}
	return permissions
}

// gcpVpcBasePermissions is the set of permissions always required for the Databricks
// workspace creator custom role in the VPC GCP project.
// See https://docs.databricks.com/gcp/en/admin/cloud-configurations/gcp/permissions
var gcpVpcBasePermissions = []string{
	// IAM Role Management (note: iam.roles.delete is workspace project only)
	"iam.roles.create",
	"iam.roles.get",
	"iam.roles.update",
	// Firewall Management
	"compute.firewalls.create",
	"compute.firewalls.get",
	// Network Management
	"compute.networks.get",
	"compute.networks.updatePolicy",
	"compute.projects.get",
	// Project Management (read-only in VPC project)
	"resourcemanager.projects.get",
	"resourcemanager.projects.getIamPolicy",
	// Service Usage (read-only in VPC project)
	"serviceusage.services.get",
	"serviceusage.services.list",
}

// gcpByovpcPermissions are the additional permissions required when using
// a customer-managed (Bring Your Own) VPC.
var gcpByovpcPermissions = []string{
	"compute.subnetworks.get",
	"compute.subnetworks.getIamPolicy",
	"compute.subnetworks.setIamPolicy",
}

// gcpCmkPermissions are the additional permissions required when using
// Customer-Managed Keys (CMK) for encryption. These belong to the VPC project.
var gcpCmkPermissions = []string{
	"cloudkms.cryptoKeys.getIamPolicy",
	"cloudkms.cryptoKeys.setIamPolicy",
}

// gcpPscPermissions are the additional permissions required when using
// Private Service Connect (PSC).
var gcpPscPermissions = []string{
	"compute.forwardingRules.get",
	"compute.forwardingRules.list",
}
