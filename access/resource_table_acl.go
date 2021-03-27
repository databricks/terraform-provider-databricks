package access

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://docs.databricks.com/security/access-control/table-acls/object-privileges.html#operations-and-privileges

// TableACL defines table access control
type TableACL struct {
	Table             string `json:"table,omitempty"`
	View              string `json:"view,omitempty"`
	Database          string `json:"database,omitempty"`
	Catalog           bool   `json:"catalog,omitempty"`
	AnyFile           bool   `json:"any_file,omitempty"`
	AnonymousFunction bool   `json:"anonymous_function,omitempty"`
	Owner             string `json:"owner,omitempty"`
	ClusterID         string `json:"cluster_id,omitempty" tf:"computed"`

	Grants []TablePermissions `json:"grant,omitempty"`
	Denies []TablePermissions `json:"deny,omitempty"`
	// SanityCheck bool     `json:"sanity_check,omitempty"`
}

// TablePermissions ...
type TablePermissions struct {
	Principal  string   `json:"principal"`
	Privileges []string `json:"privileges" tf:"slice_set"`
}

func (ta *TableACL) permissions() map[string][]TablePermissions {
	return map[string][]TablePermissions{
		"GRANT": ta.Grants,
		"DENY":  ta.Denies,
	}
}

func (ta *TableACL) actualDatabase() string {
	if ta.Database == "" {
		return "default"
	}
	return ta.Database
}

// TypeAndKey returns ACL object type and key
func (ta *TableACL) TypeAndKey() (string, string) {
	if ta.Table != "" {
		return "TABLE", fmt.Sprintf("`%s`.`%s`", ta.actualDatabase(), ta.Table)
	}
	if ta.View != "" {
		return "VIEW", fmt.Sprintf("`%s`.`%s`", ta.actualDatabase(), ta.View)
	}
	if ta.Database != "" {
		return "DATABASE", ta.Database
	}
	if ta.Catalog {
		return "CATALOG", ""
	}
	if ta.AnyFile {
		return "ANY FILE", ""
	}
	if ta.AnonymousFunction {
		return "ANONYMOUS FUNCTION", ""
	}
	return "", ""
}

// ID returns Terraform resource ID
func (ta *TableACL) ID() string {
	objectType, key := ta.TypeAndKey()
	noBackticks := strings.ReplaceAll(key, "`", "")
	return fmt.Sprintf("%s/%s", strings.ToLower(objectType), noBackticks)
}

func loadTableACL(id string) (TableACL, error) {
	ta := TableACL{}
	split := strings.SplitN(id, "/", 2)
	if len(split) != 2 {
		return ta, fmt.Errorf("ID must be three elements: %s", id)
	}
	switch strings.ToLower(split[0]) {
	case "database":
		ta.Database = split[1]
	case "view":
		dav := strings.SplitN(split[1], ".", 2)
		if len(dav) != 2 {
			return ta, fmt.Errorf("view must have two elements")
		}
		ta.Database = dav[0]
		ta.View = dav[1]
	case "table":
		dav := strings.SplitN(split[1], ".", 2)
		if len(dav) != 2 {
			return ta, fmt.Errorf("table must have two elements")
		}
		ta.Database = dav[0]
		ta.Table = dav[1]
	case "catalog":
		ta.Catalog = true
	case "any file":
		ta.AnyFile = true
	case "anonymous function":
		ta.AnonymousFunction = true
	default:
		return ta, fmt.Errorf("illegal ID type: %s", split[0])
	}
	return ta, nil
}

func (ta *TableACL) read(exec common.CommandExecutor) error {
	objectType, key := ta.TypeAndKey()
	result := exec.Execute(ta.ClusterID, "sql", fmt.Sprintf(
		"SHOW GRANT ON %s %s", objectType, key))
	if result.Failed() {
		failure := result.Error()
		if strings.Contains(failure, "does not exist") ||
			strings.Contains(failure, "RESOURCE_DOES_NOT_EXIST") {
			return common.NotFound(failure)
		}
		return fmt.Errorf(failure)
	}
	// clear any previous entries
	ta.Grants = []TablePermissions{}
	ta.Denies = []TablePermissions{}
	var principal, actionType, objType, objectKey string
	for result.Scan(&principal, &actionType, &objType, &objectKey) {
		if !strings.EqualFold(objectType, objType) {
			continue
		}
		if !strings.EqualFold(objectKey, key) {
			continue
		}
		target := ta.Grants
		var isDeny = strings.HasPrefix(actionType, "DENIED_")
		if isDeny {
			actionType = strings.Replace(actionType, "DENIED_", "", 1)
			target = ta.Denies
		}
		var privileges *[]string
		for i, tp := range target {
			if tp.Principal == principal {
				if isDeny {
					privileges = &ta.Denies[i].Privileges
				} else {
					privileges = &ta.Grants[i].Privileges
				}
			}
		}
		if privileges == nil {
			tablePerms := TablePermissions{
				Principal:  principal,
				Privileges: []string{},
			}
			if isDeny {
				ta.Denies = append(ta.Denies, tablePerms)
				privileges = &ta.Denies[len(ta.Denies)-1].Privileges
			} else {
				ta.Grants = append(ta.Grants, tablePerms)
				privileges = &ta.Grants[len(ta.Grants)-1].Privileges
			}
		}
		*privileges = append(*privileges, actionType)
	}
	return nil
}

func (ta *TableACL) revoke(commandsAPI common.CommandExecutor) error {
	existing, err := loadTableACL(ta.ID())
	if err != nil {
		return err
	}
	if err = existing.read(commandsAPI); err != nil {
		return err
	}
	for _, tps := range existing.permissions() {
		for _, priv := range tps {
			if err = ta.apply(commandsAPI, func(objType, key string) string {
				return fmt.Sprintf("REVOKE ALL PRIVILEGES ON %s %s FROM `%s`",
					objType, key, priv.Principal)
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ta *TableACL) enforce(commandsAPI common.CommandExecutor) (err error) {
	if err = ta.revoke(commandsAPI); err != nil {
		return err
	}
	for action, grantsOrDenies := range ta.permissions() {
		for _, grant := range grantsOrDenies {
			if err = ta.apply(commandsAPI, func(objType, key string) string {
				privileges := strings.Join(grant.Privileges, ", ")
				return fmt.Sprintf("%s %s ON %s %s TO `%s`",
					action, privileges, objType, key, grant.Principal)
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ta *TableACL) apply(exec common.CommandExecutor, qb func(objType, key string) string) error {
	sqlQuery := qb(ta.TypeAndKey())
	log.Printf("[INFO] Executing SQL: %s", sqlQuery)
	r := exec.Execute(ta.ClusterID, "sql", sqlQuery)
	return r.Err()
}

func newTableACLExecutor(ctx context.Context,
	ta *TableACL, d *schema.ResourceData,
	c *common.DatabricksClient) (common.CommandExecutor, error) {
	clustersAPI := compute.NewClustersAPI(ctx, c)
	if ci, ok := d.GetOk("cluster_id"); ok {
		ta.ClusterID = ci.(string)
	} else {
		sparkVersion := clustersAPI.LatestSparkVersionOrDefault(compute.SparkVersionRequest{
			Latest: true,
		})
		nodeType := clustersAPI.GetSmallestNodeType(compute.NodeTypeRequest{LocalDisk: true})
		aclCluster, err := clustersAPI.GetOrCreateRunningCluster("terrraform-table-acl", compute.Cluster{
			ClusterName:            "terrraform-table-acl",
			SparkVersion:           sparkVersion,
			NodeTypeID:             nodeType,
			AutoterminationMinutes: 10,
			SparkConf: map[string]string{
				"spark.databricks.acl.dfAclsEnabled": "true",
				"spark.master":                       "local[*]",
			},
			CustomTags: map[string]string{
				"ResourceClass": "SingleNode",
			},
		})
		if err != nil {
			return nil, err
		}
		ta.ClusterID = aclCluster.ClusterID
	}
	clusterInfo, err := clustersAPI.StartAndGetInfo(ta.ClusterID)
	if err != nil {
		return nil, err
	}
	if v, ok := clusterInfo.SparkConf["spark.databricks.acl.dfAclsEnabled"]; !ok || v != "true" {
		return nil, fmt.Errorf("cluster_id: not a High-Concurrency cluster: %s (%s)",
			clusterInfo.ClusterName, clusterInfo.ClusterID)
	}
	return c.CommandExecutor(ctx), nil
}

// ResourceTableACL manages table ACLs
func ResourceTableACL() *schema.Resource {
	s := common.StructToSchema(TableACL{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		alof := []string{"database", "table", "view", "catalog", "any_file", "anonymous_function"}
		for _, field := range alof {
			s[field].ForceNew = true
			s[field].Optional = true
			s[field].AtLeastOneOf = alof
		}
		s["database"].Default = "default"
		// TODO: ignore changes on cluster_id

		// validateGrants := validation.StringInSlice([]string{
		// 	"SELECT",
		// 	"CREATE",
		// 	"MODIFY",
		// 	"READ_METADATA",
		// 	"CREATE_NAMED_FUNCTION",
		// }, false)
		// s["grants"].ValidateFunc = validateGrants
		// s["denies"].ValidateFunc = validateGrants
		return s
	})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ta TableACL
			if err := common.DataToStructPointer(d, s, &ta); err != nil {
				return err
			}
			commandsAPI, err := newTableACLExecutor(ctx, &ta, d, c)
			if err != nil {
				return err
			}
			if err = ta.enforce(commandsAPI); err != nil {
				return err
			}
			d.SetId(ta.ID())
			// TODO: owner
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ta, err := loadTableACL(d.Id())
			if err != nil {
				return err
			}
			commandsAPI, err := newTableACLExecutor(ctx, &ta, d, c)
			if err != nil {
				return err
			}
			if err = ta.read(commandsAPI); err != nil {
				return err
			}
			return common.StructToData(ta, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var new TableACL
			if err := common.DataToStructPointer(d, s, &new); err != nil {
				return err
			}
			commandsAPI, err := newTableACLExecutor(ctx, &new, d, c)
			if err != nil {
				return err
			}
			return new.enforce(commandsAPI)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ta TableACL
			if err := common.DataToStructPointer(d, s, &ta); err != nil {
				return err
			}
			commandsAPI, err := newTableACLExecutor(ctx, &ta, d, c)
			if err != nil {
				return err
			}
			return ta.revoke(commandsAPI)
		},
	}.ToResource()
}
