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

// SqlPermissions defines table access control
type SqlPermissions struct {
	Table                string                `json:"table,omitempty"`
	View                 string                `json:"view,omitempty"`
	Database             string                `json:"database,omitempty"`
	Catalog              bool                  `json:"catalog,omitempty"`
	AnyFile              bool                  `json:"any_file,omitempty"`
	AnonymousFunction    bool                  `json:"anonymous_function,omitempty"`
	ClusterID            string                `json:"cluster_id,omitempty" tf:"computed"`
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`

	exec common.CommandExecutor
}

// PrivilegeAssignment ...
type PrivilegeAssignment struct {
	Principal  string   `json:"principal"`
	Privileges []string `json:"privileges" tf:"slice_set"`
}

func (ta *SqlPermissions) actualDatabase() string {
	if ta.Database == "" {
		return "default"
	}
	return ta.Database
}

// typeAndKey returns ACL object type and key
func (ta *SqlPermissions) typeAndKey() (string, string) {
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
func (ta *SqlPermissions) ID() string {
	objectType, key := ta.typeAndKey()
	if objectType == "" && key == "" {
		return ""
	}
	noBackticks := strings.ReplaceAll(key, "`", "")
	return fmt.Sprintf("%s/%s", strings.ToLower(objectType), noBackticks)
}

func loadTableACL(id string) (SqlPermissions, error) {
	ta := SqlPermissions{}
	split := strings.SplitN(id, "/", 2)
	if len(split) != 2 {
		return ta, fmt.Errorf("ID must be two elements: %s", id)
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

func (ta *SqlPermissions) read() error {
	thisType, thisKey := ta.typeAndKey()
	if thisType == "" && thisKey == "" {
		return fmt.Errorf("invalid ID")
	}
	currentGrantsOnThis := ta.exec.Execute(ta.ClusterID, "sql", fmt.Sprintf(
		"SHOW GRANT ON %s %s", thisType, thisKey))
	if currentGrantsOnThis.Failed() {
		failure := currentGrantsOnThis.Error()
		if strings.Contains(failure, "does not exist") ||
			strings.Contains(failure, "RESOURCE_DOES_NOT_EXIST") {
			return common.NotFound(failure)
		}
		return fmt.Errorf(failure)
	}
	// clear any previous entries
	ta.PrivilegeAssignments = []PrivilegeAssignment{}

	// iterate over existing permissions over given data object
	var currentPrincipal, currentAction, currentType, currentKey string
	for currentGrantsOnThis.Scan(&currentPrincipal, &currentAction, &currentType, &currentKey) {
		if !strings.EqualFold(currentType, thisType) {
			continue
		}
		if !strings.EqualFold(currentKey, thisKey) {
			continue
		}
		if strings.HasPrefix(currentAction, "DENIED_") {
			// DENY statements are intentionally not supported.
			continue
		}
		if currentAction == "OWN" {
			// skip table ownership definitions for now
			continue
		}
		// find existing grants for all principals
		var privileges *[]string
		for i, privilegeAssignment := range ta.PrivilegeAssignments {
			// correct all privileges for the same principal into a slide
			if privilegeAssignment.Principal == currentPrincipal {
				privileges = &ta.PrivilegeAssignments[i].Privileges
			}
		}
		if privileges == nil {
			// initialize permissions wrapper for a principal not seen
			// in previous iterations
			firstSeenPrincipalPermissions := PrivilegeAssignment{
				Principal:  currentPrincipal,
				Privileges: []string{},
			}
			// point privileges to be of the newly added principal
			ta.PrivilegeAssignments = append(ta.PrivilegeAssignments, firstSeenPrincipalPermissions)
			privileges = &ta.PrivilegeAssignments[len(ta.PrivilegeAssignments)-1].Privileges
		}
		// add action for the principal on current iteration
		*privileges = append(*privileges, currentAction)
	}
	return nil
}

func (ta *SqlPermissions) revoke() error {
	existing, err := loadTableACL(ta.ID())
	if err != nil {
		return err
	}
	existing.exec = ta.exec
	existing.ClusterID = ta.ClusterID
	if err = existing.read(); err != nil {
		return err
	}
	for _, privilegeAssignment := range existing.PrivilegeAssignments {
		if err = ta.apply(func(objType, key string) string {
			return fmt.Sprintf("REVOKE ALL PRIVILEGES ON %s %s FROM `%s`",
				objType, key, privilegeAssignment.Principal)
		}); err != nil {
			return err
		}
	}
	return nil
}

func (ta *SqlPermissions) enforce() (err error) {
	if err = ta.revoke(); err != nil {
		return err
	}
	for _, privilegeAssignment := range ta.PrivilegeAssignments {
		if err = ta.apply(func(objType, key string) string {
			privileges := strings.Join(privilegeAssignment.Privileges, ", ")
			return fmt.Sprintf("GRANT %s ON %s %s TO `%s`",
				privileges, objType, key, privilegeAssignment.Principal)
		}); err != nil {
			return err
		}
	}
	return nil
}

func (ta *SqlPermissions) apply(qb func(objType, key string) string) error {
	objType, key := ta.typeAndKey()
	if objType == "" && key == "" {
		return fmt.Errorf("invalid ID")
	}
	sqlQuery := qb(objType, key)
	log.Printf("[INFO] Executing SQL: %s", sqlQuery)
	r := ta.exec.Execute(ta.ClusterID, "sql", sqlQuery)
	return r.Err()
}

func (ta *SqlPermissions) initCluster(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) (err error) {
	clustersAPI := compute.NewClustersAPI(ctx, c)
	if ci, ok := d.GetOk("cluster_id"); ok {
		ta.ClusterID = ci.(string)
	} else {
		ta.ClusterID, err = ta.getOrCreateCluster(clustersAPI)
		if err != nil {
			return
		}
	}
	clusterInfo, err := clustersAPI.StartAndGetInfo(ta.ClusterID)
	if e, ok := err.(common.APIError); ok && e.IsMissing() {
		// cluster that was previously in a tfstate was deleted
		ta.ClusterID, err = ta.getOrCreateCluster(clustersAPI)
		if err != nil {
			return
		}
		clusterInfo, err = clustersAPI.StartAndGetInfo(ta.ClusterID)
	}
	if err != nil {
		return
	}
	if v, ok := clusterInfo.SparkConf["spark.databricks.acl.dfAclsEnabled"]; !ok || v != "true" {
		err = fmt.Errorf("cluster_id: not a High-Concurrency cluster: %s (%s)",
			clusterInfo.ClusterName, clusterInfo.ClusterID)
		return
	}
	ta.exec = c.CommandExecutor(ctx)
	return nil
}

func (ta *SqlPermissions) getOrCreateCluster(clustersAPI compute.ClustersAPI) (string, error) {
	sparkVersion := clustersAPI.LatestSparkVersionOrDefault(compute.SparkVersionRequest{
		Latest: true,
	})
	nodeType := clustersAPI.GetSmallestNodeType(compute.NodeTypeRequest{LocalDisk: true})
	aclCluster, err := clustersAPI.GetOrCreateRunningCluster(
		"terrraform-table-acl", compute.Cluster{
			ClusterName:            "terrraform-table-acl",
			SparkVersion:           sparkVersion,
			NodeTypeID:             nodeType,
			AutoterminationMinutes: 10,
			SparkConf: map[string]string{
				"spark.databricks.acl.dfAclsEnabled":     "true",
				"spark.databricks.repl.allowedLanguages": "python,sql",
				"spark.databricks.cluster.profile":       "serverless",
				"spark.master":                           "local[*]",
			},
			CustomTags: map[string]string{
				"ResourceClass": "SingleNode",
			},
		})
	if err != nil {
		return "", err
	}
	return aclCluster.ClusterID, nil
}

func tableAclForUpdate(ctx context.Context, d *schema.ResourceData,
	s map[string]*schema.Schema, c *common.DatabricksClient) (ta SqlPermissions, err error) {
	if err = common.DataToStructPointer(d, s, &ta); err != nil {
		return
	}
	err = ta.initCluster(ctx, d, c)
	return
}

func tableAclForLoad(ctx context.Context, d *schema.ResourceData,
	s map[string]*schema.Schema, c *common.DatabricksClient) (ta SqlPermissions, err error) {
	ta, err = loadTableACL(d.Id())
	if err != nil {
		return
	}
	err = ta.initCluster(ctx, d, c)
	return
}

// ResourceSqlPermissions manages table ACLs
func ResourceSqlPermissions() *schema.Resource {
	s := common.StructToSchema(SqlPermissions{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		alof := []string{"database", "table", "view", "catalog", "any_file", "anonymous_function"}
		for _, field := range alof {
			s[field].ForceNew = true
			s[field].Optional = true
			s[field].AtLeastOneOf = alof
		}
		s["cluster_id"].Computed = true
		s["database"].Default = "default"
		return s
	})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ta, err := tableAclForUpdate(ctx, d, s, c)
			if err != nil {
				return err
			}
			if err = ta.enforce(); err != nil {
				return err
			}
			d.SetId(ta.ID())
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ta, err := tableAclForLoad(ctx, d, s, c)
			if err != nil {
				return err
			}
			if err = ta.read(); err != nil {
				return err
			}
			if len(ta.PrivilegeAssignments) == 0 {
				// reflect resource is skipping empty privilege_assignments
				d.Set("privilege_assignments", []interface{}{})
			}
			return common.StructToData(ta, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ta, err := tableAclForUpdate(ctx, d, s, c)
			if err != nil {
				return err
			}
			return ta.enforce()
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ta, err := tableAclForLoad(ctx, d, s, c)
			if err != nil {
				return err
			}
			return ta.revoke()
		},
	}.ToResource()
}
