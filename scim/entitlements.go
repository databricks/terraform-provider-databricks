package scim

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type entitlements struct {
	AllowClusterCreate      bool `json:"allow_cluster_create,omitempty"`
	AllowInstancePoolCreate bool `json:"allow_instance_pool_create,omitempty"`
	DatabricksSQLAccess     bool `json:"databricks_sql_access,omitempty"`
	WorkspaceAccess         bool `json:"workspace_access,omitempty"`
}

func (e entitlements) toComplexValueList() []ComplexValue {
	result := []ComplexValue{}
	if e.AllowClusterCreate {
		result = append(result, ComplexValue{
			Value: "allow-cluster-create",
		})
	}
	if e.AllowInstancePoolCreate {
		result = append(result, ComplexValue{
			Value: "allow-instance-pool-create",
		})
	}
	if e.DatabricksSQLAccess {
		result = append(result, ComplexValue{
			Value: "databricks-sql-access",
		})
	}
	if e.WorkspaceAccess {
		result = append(result, ComplexValue{
			Value: "workspace-access",
		})
	}
	return result
}

func fromComplexValueList(ctx context.Context, cv []ComplexValue) entitlements {
	var e entitlements
	for _, c := range cv {
		switch c.Value {
		case "allow-cluster-create":
			e.AllowClusterCreate = true
		case "allow-instance-pool-create":
			e.AllowInstancePoolCreate = true
		case "databricks-sql-access":
			e.DatabricksSQLAccess = true
		case "workspace-access":
			e.WorkspaceAccess = true
		default:
			tflog.Info(ctx, fmt.Sprintf("Ignoring unknown entitlement: %s", c.Value))
		}
	}
	return e
}
