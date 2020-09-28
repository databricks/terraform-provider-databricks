package access

import (
	"context"
	"fmt"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// https://docs.databricks.com/security/access-control/table-acls/object-privileges.html#operations-and-privileges

// TableACL defines table access control
type TableACL struct {
	Table     string   `json:"table,omitempty"`
	View      string   `json:"view,omitempty"`
	Database  string   `json:"database,omitempty"`
	Principal string   `json:"principal"`
	Grants    []string `json:"grants,omitempty" tf:"slice_set"`
	Denies    []string `json:"denies,omitempty" tf:"slice_set"`
	Owner     string   `json:"owner,omitempty"`
	ClusterID string   `json:"cluster_id,omitempty"`
	// SanityCheck bool     `json:"sanity_check,omitempty"`
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
	return "", ""
}

// ID returns Terraform resource ID
func (ta *TableACL) ID() string {
	objectType, key := ta.TypeAndKey()
	noBackticks := strings.ReplaceAll(key, "`", "")
	return fmt.Sprintf("%s/%s/%s", objectType, noBackticks, ta.Principal)
}

func loadTableACL(id string) (TableACL, error) {
	ta := TableACL{}
	split := strings.SplitN(id, "/", 3)
	if len(split) != 3 {
		return ta, fmt.Errorf("ID must be three elements: %s", id)
	}
	ta.Principal = split[2]
	switch strings.ToLower(split[0]) {
	case "database":
		ta.Database = split[1]
	case "view":
		dav := strings.SplitN(split[1], ".", 2)
		if len(dav) != 2 {
			return ta, fmt.Errorf("View must be two elements")
		}
		ta.Database = dav[0]
		ta.View = dav[1]
	case "table":
		dav := strings.SplitN(split[1], ".", 2)
		if len(dav) != 2 {
			return ta, fmt.Errorf("Table must be two elements")
		}
		ta.Database = dav[0]
		ta.Table = dav[1]
	default:
		return ta, fmt.Errorf("Illegal ID type: %s", split[0])
	}
	return ta, nil
}

// Read shows all grants
func (ta *TableACL) Read(exec common.CommandExecutor) error {
	objectType, key := ta.TypeAndKey()
	result := exec.Execute(ta.ClusterID, "sql", fmt.Sprintf(
		"SHOW GRANT `%s` ON %s %s", ta.Principal, objectType, key))
	if result.Failed() {
		failure := result.Error()
		if strings.Contains(failure, "does not exist") ||
			strings.Contains(failure, "RESOURCE_DOES_NOT_EXIST") {
			return common.APIError{
				Message:    failure,
				StatusCode: 404,
			}
		}
		return fmt.Errorf(failure)
	}
	// clear any previous entries
	ta.Grants = []string{}
	ta.Denies = []string{}
	var principal, actionType, objType, objectKey string
	for result.Scan(&principal, &actionType, &objType, &objectKey) {
		if principal != ta.Principal {
			continue
		}
		if objectType != objType {
			continue
		}
		if objectKey != key {
			continue
		}
		if strings.HasPrefix(actionType, "DENIED_") {
			ta.Denies = append(ta.Denies, strings.Replace(actionType, "DENIED_", "", 1))
			continue
		}
		ta.Grants = append(ta.Grants, actionType)
	}
	return nil
}

// Apply ACL to an object
func (ta *TableACL) Apply(exec common.CommandExecutor, action, privilege, direction string) error {
	objectType, key := ta.TypeAndKey()
	r := exec.Execute(ta.ClusterID, "sql", fmt.Sprintf(
		"%s %s ON %s %s %s %s", action, privilege,
		objectType, key, direction, ta.Principal))
	return r.Err()
}

func newTableACLExecutor(ta *TableACL, d *schema.ResourceData, m interface{}) (common.CommandExecutor, error) {
	clustersAPI := compute.NewClustersAPI(m)
	// TODO: create cluster if not exists
	if ci, ok := d.GetOk("cluster_id"); ok {
		ta.ClusterID = ci.(string)
	}
	clusterInfo, err := clustersAPI.StartAndGetInfo(ta.ClusterID)
	// if common.IsMissing(err) {
	// 	return nil, err
	// }
	if err != nil {
		return nil, err
	}
	if v, ok := clusterInfo.SparkConf["spark.databricks.acl.dfAclsEnabled"]; !ok || v != "true" {
		return nil, fmt.Errorf("cluster_id: not a High-Concurrency cluster: %s (%s)",
			clusterInfo.ClusterName, clusterInfo.ClusterID)
	}
	return compute.NewCommandsAPI(m), nil
}

// ResourceTableACL manages table ACLs
func ResourceTableACL() *schema.Resource {
	s := internal.StructToSchema(TableACL{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["table"].ForceNew = true
		s["view"].ForceNew = true
		s["database"].ForceNew = true
		s["database"].Default = "default"

		s["principal"].ValidateFunc = validation.StringIsNotEmpty

		alof := []string{"database", "table", "view"}
		s["table"].AtLeastOneOf = alof
		s["view"].AtLeastOneOf = alof
		s["database"].AtLeastOneOf = alof

		s["table"].ConflictsWith = []string{"view"}
		s["view"].ConflictsWith = []string{"table"}

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
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		ta, err := loadTableACL(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		commandsAPI, err := newTableACLExecutor(&ta, d, m)
		if err != nil {
			return diag.FromErr(err)
		}
		err = ta.Read(commandsAPI)
		if common.IsMissing(err) {
			d.SetId("")
			return nil
		}
		err = internal.StructToData(ta, s, d)
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ta TableACL
			err := internal.DataToStructPointer(d, s, &ta)
			if err != nil {
				return diag.FromErr(err)
			}
			commandsAPI, err := newTableACLExecutor(&ta, d, m)
			if common.IsMissing(err) {
				d.SetId("")
				return nil
			}
			if err != nil {
				return diag.FromErr(err)
			}
			err = ta.Apply(commandsAPI, "REVOKE", "ALL PRIVILEGES", "FROM")
			if err != nil {
				return diag.FromErr(err)
			}
			for _, grant := range ta.Grants {
				err = ta.Apply(commandsAPI, "GRANT", grant, "TO")
				if err != nil {
					return diag.FromErr(err)
				}
			}
			for _, deny := range ta.Denies {
				err = ta.Apply(commandsAPI, "DENY", deny, "TO")
				if err != nil {
					return diag.FromErr(err)
				}
			}
			d.SetId(ta.ID())
			// TODO: owner
			return readContext(ctx, d, m)
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			old, err := loadTableACL(d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			var new TableACL
			err = internal.DataToStructPointer(d, s, &new)
			if err != nil {
				return diag.FromErr(err)
			}
			commandsAPI, err := newTableACLExecutor(&new, d, m)
			if err != nil {
				return diag.FromErr(err)
			}
			applyDiff := func(a, b []string, action, direction string) (err error) {
				for _, newGrant := range a {
					for _, oldGrant := range b {
						if oldGrant == newGrant {
							return
						}
					}
					err = new.Apply(commandsAPI, action, newGrant, direction)
					if err != nil {
						return
					}
				}
				return
			}
			// at first glance, we may not need Update here, though if we revoke all priviliges 
			// from the group during the lifetime of streaming query, it might fail. so that is
			// why we're carefully revoking grants only when they change
			for _, err := range []error{
				applyDiff(new.Grants, old.Grants, "GRANT", "TO"),
				applyDiff(old.Grants, new.Grants, "REVOKE", "FROM"),
				applyDiff(new.Denies, old.Denies, "DENY", "TO"),
				applyDiff(old.Denies, new.Denies, "REVOKE", "FROM"),
			} {
				if err != nil {
					return diag.FromErr(err)
				}
			}
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ta TableACL
			err := internal.DataToStructPointer(d, s, &ta)
			if err != nil {
				return diag.FromErr(err)
			}
			commandsAPI, err := newTableACLExecutor(&ta, d, m)
			if err != nil {
				return diag.FromErr(err)
			}
			err = ta.Apply(commandsAPI, "REVOKE", "ALL PRIVILEGES", "FROM")
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Schema: s,
	}
}
