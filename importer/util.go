package importer

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (ic *importContext) importCluster(c *compute.Cluster) error {
	if c == nil {
		return nil
	}
	for _, is := range c.InitScripts {
		if is.Dbfs != nil {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       is.Dbfs.Destination,
			})
		}
	}
	if c.AwsAttributes != nil {
		ic.Emit(&resource{
			Resource: "databricks_instance_profile",
			ID:       c.AwsAttributes.InstanceProfileArn,
		})
	}
	if c.InstancePoolID != "" {
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       c.InstancePoolID,
		})
	}
	if c.PolicyID != "" {
		ic.Emit(&resource{
			Resource: "databricks_cluster_policy",
			ID:       c.PolicyID,
		})
	}
	return nil
}

func (ic *importContext) importLibraries(d *schema.ResourceData, s map[string]*schema.Schema) error {
	var cll compute.ClusterLibraryList
	err := internal.DataToStructPointer(d, s, &cll)
	if err != nil {
		return err
	}
	for _, lib := range cll.Libraries {
		if strings.HasPrefix(lib.Whl, "dbfs:") {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       lib.Whl,
			})
		}
		if strings.HasPrefix(lib.Jar, "dbfs:") {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       lib.Jar,
			})
		}
		if strings.HasPrefix(lib.Egg, "dbfs:") {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       lib.Egg,
			})
		}
	}
	return nil
}

func (ic *importContext) cacheGroups() error {
	if len(ic.allGroups) == 0 {
		log.Printf("[INFO] Caching groups in memory ...")
		groupsAPI := identity.NewGroupsAPI(ic.Client)
		g, err := groupsAPI.Filter("")
		if err != nil {
			return err
		}
		ic.allGroups = g.Resources
		log.Printf("[INFO] Cached %d groups", len(ic.allGroups))
	}
	return nil
}

// func (ic *importContext) cacheUsers() error {
// 	if len(ic.allUsers) == 0 {
// 		// workspace has at least one user, always.
// 		log.Printf("[INFO] Fetching users into in-memory cache")
// 		usersAPI := identity.NewUsersAPI(ic.Client)
// 		users, err := usersAPI.Filter("active eq true")
// 		if err != nil {
// 			return err
// 		}
// 		ic.allUsers = users
// 		log.Printf("[INFO] Cached %d users", len(users))
// 	}
// 	return nil
// }

func (ic *importContext) findUserByName(name string) (u identity.ScimUser, err error) {
	a := identity.NewUsersAPI(ic.Client)
	users, err := a.Filter(fmt.Sprintf("userName eq '%s'", name))
	if err != nil {
		return
	}
	if len(users) == 0 {
		err = fmt.Errorf("User %s not found", name)
		return
	}
	u = users[0]
	return
	// if err := ic.cacheUsers(); err != nil {
	// 	return
	// }
	// for _, _u := range ic.allUsers {
	// 	if _u.UserName == name {
	// 		u = _u
	// 		return
	// 	}
	// }
	// err = fmt.Errorf("User %s not found", name)
	// return
}

func (ic *importContext) emitIfDbfsFile(path string) {
	if strings.HasPrefix(path, "dbfs:") {
		ic.Emit(&resource{
			Resource: "databricks_dbfs_file",
			ID:       path,
		})
	}
}

func (ic *importContext) refreshMounts() error {
	if ic.mountMap != nil {
		return nil
	}
	clustersAPI := compute.NewClustersAPI(ic.Client)
	commandAPI := ic.Client.CommandExecutor()
	cluster, err := clustersAPI.GetOrCreateRunningCluster("terraform-mount")
	if err != nil {
		return err
	}
	// TODO: edit instance profile of a cluster in a loop
	j, err := commandAPI.Execute(cluster.ClusterID, "python", `
	import json
	json.dumps({mp.replace('/mnt/', ''):source for mp, source, _ in dbutils.fs.mounts()})`)
	if err != nil {
		return err
	}
	ic.mountMap = map[string]string{}
	return json.Unmarshal([]byte(j), &ic.mountMap)
}
