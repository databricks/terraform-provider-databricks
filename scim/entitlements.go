package scim

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type entitlements struct {
	AllowClusterCreate      bool `json:"allow_cluster_create,omitempty"`
	AllowInstancePoolCreate bool `json:"allow_instance_pool_create,omitempty"`
	DatabricksSQLAccess     bool `json:"databricks_sql_access,omitempty"`
	WorkspaceAccess         bool `json:"workspace_access,omitempty"`
	WorkspaceConsume        bool `json:"workspace_consume,omitempty"`
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
	if e.WorkspaceConsume {
		result = append(result, ComplexValue{
			Value: "workspace-consume",
		})
	}
	return result
}

func newEntitlements(ctx context.Context, cv []ComplexValue) entitlements {
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
		case "workspace-consume":
			e.WorkspaceConsume = true
		default:
			tflog.Warn(ctx, fmt.Sprintf("Ignoring unknown entitlement: %s", c.Value))
		}
	}
	return e
}

func mergeEntitlements(e1, e2 entitlements) entitlements {
	return entitlements{
		AllowClusterCreate:      e1.AllowClusterCreate || e2.AllowClusterCreate,
		AllowInstancePoolCreate: e1.AllowInstancePoolCreate || e2.AllowInstancePoolCreate,
		DatabricksSQLAccess:     e1.DatabricksSQLAccess || e2.DatabricksSQLAccess,
		WorkspaceAccess:         e1.WorkspaceAccess || e2.WorkspaceAccess,
		WorkspaceConsume:        e1.WorkspaceConsume || e2.WorkspaceConsume,
	}
}

func customizeEntitlementsSchema(m map[string]*schema.Schema) map[string]*schema.Schema {
	m["workspace_consume"].ConflictsWith = []string{"workspace_access", "databricks_sql_access"}
	return m
}
