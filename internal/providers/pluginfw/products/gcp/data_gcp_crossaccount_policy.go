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

const crossAccountDataSourceName = "gcp_crossaccount_policy"

func DataSourceGcpCrossaccountPolicy() datasource.DataSource {
	return &GcpCrossaccountPolicyDataSource{}
}

var _ datasource.DataSourceWithConfigure = &GcpCrossaccountPolicyDataSource{}

type GcpCrossaccountPolicyDataSource struct {
	Client *common.DatabricksClient
}

// GcpCrossaccountPolicyData holds the input/output for the workspace project
// custom role. Compute/networking permissions belong to databricks_gcp_vpc_policy.
type GcpCrossaccountPolicyData struct {
	Permissions types.List `tfsdk:"permissions"`
	tfschema.Namespace
}

func (GcpCrossaccountPolicyData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permissions"] = attrs["permissions"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (GcpCrossaccountPolicyData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions":     reflect.TypeOf(types.String{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *GcpCrossaccountPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(crossAccountDataSourceName)
}

func (d *GcpCrossaccountPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, GcpCrossaccountPolicyData{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *GcpCrossaccountPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *GcpCrossaccountPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, crossAccountDataSourceName)

	var data GcpCrossaccountPolicyData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	permissions := computeGcpCrossaccountPermissions()

	permValues := make([]attr.Value, len(permissions))
	for i, p := range permissions {
		permValues[i] = types.StringValue(p)
	}
	data.Permissions = types.ListValueMust(types.StringType, permValues)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func computeGcpCrossaccountPermissions() []string {
	permissions := make([]string, len(gcpWorkspaceCreatorBasePermissions))
	copy(permissions, gcpWorkspaceCreatorBasePermissions)
	return permissions
}

// gcpWorkspaceCreatorBasePermissions is the set of permissions required for the
// Databricks workspace creator custom role in the workspace GCP project.
// See https://docs.databricks.com/gcp/en/admin/cloud-configurations/gcp/permissions
// Compute and networking permissions (needed in the VPC project) are covered by
// databricks_gcp_vpc_policy.
var gcpWorkspaceCreatorBasePermissions = []string{
	// IAM Role Management
	"iam.roles.create",
	"iam.roles.delete",
	"iam.roles.get",
	"iam.roles.update",
	// Service Account Management
	"iam.serviceAccounts.create",
	"iam.serviceAccounts.get",
	"iam.serviceAccounts.getIamPolicy",
	"iam.serviceAccounts.setIamPolicy",
	// Project Management
	"resourcemanager.projects.get",
	"resourcemanager.projects.getIamPolicy",
	"resourcemanager.projects.setIamPolicy",
	// Service Usage
	"serviceusage.services.get",
	"serviceusage.services.list",
	"serviceusage.services.enable",
}
