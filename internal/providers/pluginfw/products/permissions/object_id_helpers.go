package permissions

import (
	"context"
	"maps"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ObjectIdentifiers is a struct containing all object identifier fields
// This is used by both PermissionResourceModel and the validator
type ObjectIdentifiers struct {
	ClusterId              types.String `tfsdk:"cluster_id"`
	ClusterPolicyId        types.String `tfsdk:"cluster_policy_id"`
	InstancePoolId         types.String `tfsdk:"instance_pool_id"`
	JobId                  types.String `tfsdk:"job_id"`
	PipelineId             types.String `tfsdk:"pipeline_id"`
	NotebookId             types.String `tfsdk:"notebook_id"`
	NotebookPath           types.String `tfsdk:"notebook_path"`
	DirectoryId            types.String `tfsdk:"directory_id"`
	DirectoryPath          types.String `tfsdk:"directory_path"`
	WorkspaceFileId        types.String `tfsdk:"workspace_file_id"`
	WorkspaceFilePath      types.String `tfsdk:"workspace_file_path"`
	RegisteredModelId      types.String `tfsdk:"registered_model_id"`
	ExperimentId           types.String `tfsdk:"experiment_id"`
	SqlDashboardId         types.String `tfsdk:"sql_dashboard_id"`
	SqlEndpointId          types.String `tfsdk:"sql_endpoint_id"`
	SqlQueryId             types.String `tfsdk:"sql_query_id"`
	SqlAlertId             types.String `tfsdk:"sql_alert_id"`
	DashboardId            types.String `tfsdk:"dashboard_id"`
	RepoId                 types.String `tfsdk:"repo_id"`
	RepoPath               types.String `tfsdk:"repo_path"`
	Authorization          types.String `tfsdk:"authorization"`
	ServingEndpointId      types.String `tfsdk:"serving_endpoint_id"`
	VectorSearchEndpointId types.String `tfsdk:"vector_search_endpoint_id"`
	AppName                types.String `tfsdk:"app_name"`
	DatabaseInstanceName   types.String `tfsdk:"database_instance_name"`
	AlertV2Id              types.String `tfsdk:"alert_v2_id"`
}

// ToFieldValuesMap converts the ObjectIdentifiers struct to a map of field names to values
func (o *ObjectIdentifiers) ToFieldValuesMap() map[string]string {
	return map[string]string{
		"cluster_id":                o.ClusterId.ValueString(),
		"cluster_policy_id":         o.ClusterPolicyId.ValueString(),
		"instance_pool_id":          o.InstancePoolId.ValueString(),
		"job_id":                    o.JobId.ValueString(),
		"pipeline_id":               o.PipelineId.ValueString(),
		"notebook_id":               o.NotebookId.ValueString(),
		"notebook_path":             o.NotebookPath.ValueString(),
		"directory_id":              o.DirectoryId.ValueString(),
		"directory_path":            o.DirectoryPath.ValueString(),
		"workspace_file_id":         o.WorkspaceFileId.ValueString(),
		"workspace_file_path":       o.WorkspaceFilePath.ValueString(),
		"registered_model_id":       o.RegisteredModelId.ValueString(),
		"experiment_id":             o.ExperimentId.ValueString(),
		"sql_dashboard_id":          o.SqlDashboardId.ValueString(),
		"sql_endpoint_id":           o.SqlEndpointId.ValueString(),
		"sql_query_id":              o.SqlQueryId.ValueString(),
		"sql_alert_id":              o.SqlAlertId.ValueString(),
		"dashboard_id":              o.DashboardId.ValueString(),
		"repo_id":                   o.RepoId.ValueString(),
		"repo_path":                 o.RepoPath.ValueString(),
		"authorization":             o.Authorization.ValueString(),
		"serving_endpoint_id":       o.ServingEndpointId.ValueString(),
		"vector_search_endpoint_id": o.VectorSearchEndpointId.ValueString(),
		"app_name":                  o.AppName.ValueString(),
		"database_instance_name":    o.DatabaseInstanceName.ValueString(),
		"alert_v2_id":               o.AlertV2Id.ValueString(),
	}
}

// GetObjectIdentifierFields returns all possible object identifier field names
// This is derived from the keys of ToFieldValuesMap() to maintain single source of truth
func GetObjectIdentifierFields() []string {
	var empty ObjectIdentifiers
	return slices.Collect(maps.Keys(empty.ToFieldValuesMap()))
}

// ExtractObjectIdentifiersFromConfig reads object identifiers from a tfsdk.Config
// Returns nil if there are errors reading the config
func ExtractObjectIdentifiersFromConfig(ctx context.Context, config tfsdk.Config) *ObjectIdentifiers {
	var objectIds ObjectIdentifiers
	diags := config.Get(ctx, &objectIds)
	if diags.HasError() {
		return nil
	}
	return &objectIds
}
