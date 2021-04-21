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
	Denies []TablePermissions `json:"deny,omitempty"` // TODO: remove denies

	exec common.CommandExecutor
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

// typeAndKey returns ACL object type and key
func (ta *TableACL) typeAndKey() (string, string) {
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
	objectType, key := ta.typeAndKey()
	if objectType == "" && key == "" {
		return ""
	}
	noBackticks := strings.ReplaceAll(key, "`", "")
	return fmt.Sprintf("%s/%s", strings.ToLower(objectType), noBackticks)
}

func loadTableACL(id string) (TableACL, error) {
	ta := TableACL{}
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

func (ta *TableACL) read() error {
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
	ta.Grants = []TablePermissions{}
	ta.Denies = []TablePermissions{}

	// iterate over existing permissions over given data object
	var currentPrincipal, currentAction, currentType, currentKey string
	for currentGrantsOnThis.Scan(&currentPrincipal, &currentAction, &currentType, &currentKey) {
		if !strings.EqualFold(currentType, thisType) {
			continue
		}
		if !strings.EqualFold(currentKey, thisKey) {
			continue
		}
		// find existing grants or denies for all principals
		target := ta.Grants
		var isDeny = strings.HasPrefix(currentAction, "DENIED_")
		if isDeny {
			// TODO: this is going to be removed
			currentAction = strings.Replace(currentAction, "DENIED_", "", 1)
			target = ta.Denies
		}
		var privileges *[]string
		for i, tablePermissions := range target {
			// correct all privileges for the same principal into a slide
			if tablePermissions.Principal == currentPrincipal {
				if isDeny {
					privileges = &ta.Denies[i].Privileges
				} else {
					privileges = &ta.Grants[i].Privileges
				}
			}
		}
		if privileges == nil {
			// initialize permissions wrapper for a principal not seen
			// in previous iterations
			firstSeenPrincipalPermissions := TablePermissions{
				Principal:  currentPrincipal,
				Privileges: []string{},
			}
			if isDeny {
				// TODO: this will be removed
				ta.Denies = append(ta.Denies, firstSeenPrincipalPermissions)
				privileges = &ta.Denies[len(ta.Denies)-1].Privileges
			} else {
				// point privileges to be of the newly added principal
				ta.Grants = append(ta.Grants, firstSeenPrincipalPermissions)
				privileges = &ta.Grants[len(ta.Grants)-1].Privileges
			}
		}
		// add action for the principal on current iteration
		*privileges = append(*privileges, currentAction)
	}
	return nil
}

func (ta *TableACL) revoke() error {
	existing, err := loadTableACL(ta.ID())
	if err != nil {
		return err
	}
	existing.exec = ta.exec
	existing.ClusterID = ta.ClusterID
	if err = existing.read(); err != nil {
		return err
	}
	for _, tps := range existing.permissions() {
		for _, priv := range tps {
			if err = ta.apply(func(objType, key string) string {
				return fmt.Sprintf("REVOKE ALL PRIVILEGES ON %s %s FROM `%s`",
					objType, key, priv.Principal)
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ta *TableACL) enforce() (err error) {
	if err = ta.revoke(); err != nil {
		return err
	}
	for action, grantsOrDenies := range ta.permissions() {
		for _, grant := range grantsOrDenies {
			if err = ta.apply(func(objType, key string) string {
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

func (ta *TableACL) apply(qb func(objType, key string) string) error {
	objType, key := ta.typeAndKey()
	if objType == "" && key == "" {
		return fmt.Errorf("invalid ID")
	}
	sqlQuery := qb(objType, key)
	log.Printf("[INFO] Executing SQL: %s", sqlQuery)
	r := ta.exec.Execute(ta.ClusterID, "sql", sqlQuery)
	return r.Err()
}

func (ta *TableACL) initCluster(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) (err error) {
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

func (ta *TableACL) getOrCreateCluster(clustersAPI compute.ClustersAPI) (string, error) {
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
			"spark.databricks.cluster.profile":   "singleNode",
			"spark.master":                       "local[*]",
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
	s map[string]*schema.Schema, c *common.DatabricksClient) (ta TableACL, err error) {
	if err = common.DataToStructPointer(d, s, &ta); err != nil {
		return
	}
	err = ta.initCluster(ctx, d, c)
	return
}

func tableAclForLoad(ctx context.Context, d *schema.ResourceData,
	s map[string]*schema.Schema, c *common.DatabricksClient) (ta TableACL, err error) {
	ta, err = loadTableACL(d.Id())
	if err != nil {
		return
	}
	err = ta.initCluster(ctx, d, c)
	return
}

// ResourceSqlPermissions manages table ACLs
func ResourceSqlPermissions() *schema.Resource {
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
		// 	"USAGE",
		// 	"READ_METADATA",
		// 	"CREATE_NAMED_FUNCTION",
		// 	"MODIFY_CLASSPATH",
		// 	"ALL PRIVILEGES",
		// }, false)
		// s["grant"].ValidateDiagFunc = validation.ToDiagFunc(validateGrants)
		// s["deny"].ValidateDiagFunc = validation.ToDiagFunc(validateGrants)
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
