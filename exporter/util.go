package exporter

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/databrickslabs/terraform-provider-databricks/aws"
	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/jobs"
	"github.com/databrickslabs/terraform-provider-databricks/libraries"
	"github.com/databrickslabs/terraform-provider-databricks/scim"
	"github.com/databrickslabs/terraform-provider-databricks/storage"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (ic *importContext) importCluster(c *clusters.Cluster) {
	if c == nil {
		return
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
}

func (ic *importContext) importLibraries(d *schema.ResourceData, s map[string]*schema.Schema) error {
	var cll libraries.ClusterLibraryList
	err := common.DataToStructPointer(d, s, &cll)
	if err != nil {
		return err
	}
	for _, lib := range cll.Libraries {
		ic.emitIfDbfsFile(lib.Whl)
		ic.emitIfDbfsFile(lib.Jar)
		ic.emitIfDbfsFile(lib.Egg)
	}
	return nil
}

func (ic *importContext) cacheGroups() error {
	if len(ic.allGroups) == 0 {
		log.Printf("[INFO] Caching groups in memory ...")
		groupsAPI := scim.NewGroupsAPI(ic.Context, ic.Client)
		g, err := groupsAPI.Filter("")
		if err != nil {
			return err
		}
		ic.allGroups = g.Resources
		log.Printf("[INFO] Cached %d groups", len(ic.allGroups))
	}
	return nil
}

func (ic *importContext) findUserByName(name string) (u scim.User, err error) {
	a := scim.NewUsersAPI(ic.Context, ic.Client)
	users, err := a.Filter(fmt.Sprintf("userName eq '%s'", name))
	if err != nil {
		return
	}
	if len(users) == 0 {
		err = fmt.Errorf("user %s not found", name)
		return
	}
	u = users[0]
	return
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
	commandAPI := ic.Client.CommandExecutor(ic.Context)
	clustersAPI := clusters.NewClustersAPI(ic.Context, ic.Client)
	cluster, err := clustersAPI.GetOrCreateRunningCluster("terraform-mount")
	if err != nil {
		return err
	}
	log.Printf("[INFO] Refreshing worskpace-wide mounts")
	mountMap, err := ic.getMountsThroughCluster(commandAPI, cluster.ClusterID)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Found %d worskpace-wide mounts", len(mountMap))
	ic.mountMap = map[string]mount{}
	for k, v := range mountMap {
		ic.mountMap[k] = mount{
			URL: v,
			// cluster id is needed for AWS S3 mounts, that may
			// be visible for every cluster
			ClusterID: cluster.ClusterID,
		}
	}
	if ic.Client.IsAws() {
		profiles, err := aws.NewInstanceProfilesAPI(ic.Context, ic.Client).List()
		if err != nil {
			return err
		}
		for _, instanceProfile := range profiles {
			log.Printf("[INFO] Refreshing mounts accessible by %s", instanceProfile.InstanceProfileArn)
			profileCluster, err := storage.GetOrCreateMountingClusterWithInstanceProfile(
				clustersAPI, instanceProfile.InstanceProfileArn)
			if err != nil {
				return err
			}
			profileMountMap, err := ic.getMountsThroughCluster(commandAPI, profileCluster.ClusterID)
			if err != nil {
				return err
			}
			i := 0
			for k, v := range profileMountMap {
				if _, has := ic.mountMap[k]; has {
					continue
				}
				i++
				ic.mountMap[k] = mount{
					URL:             v,
					InstanceProfile: instanceProfile.InstanceProfileArn,
				}
			}
			if i > 0 {
				log.Printf("[INFO] Found %d mounts accessible by %s",
					len(profileMountMap), instanceProfile.InstanceProfileArn)
			}
		}
	}
	return nil
}

var getReadableMountsCommand = `
import scala.concurrent._
import scala.concurrent.duration._
import ExecutionContext.Implicits.global
import scala.concurrent.{Await, Future}
import com.fasterxml.jackson.databind.{DeserializationFeature, ObjectMapper}
import com.fasterxml.jackson.module.scala.experimental.ScalaObjectMapper
import com.fasterxml.jackson.module.scala.DefaultScalaModule

val readableMounts = dbutils.fs.mounts
  .filter(_.mountPoint.startsWith("/mnt"))
  .par.map { mount =>
    try {
        Await.result(Future {
            dbutils.fs.ls(mount.mountPoint)
            (mount.mountPoint
                .replace("/mnt/", "")
                .stripSuffix("/"), 
             mount.source)
        }, 5.second)
    } catch {
        case _ : Throwable => (null, mount.source)
    }
  }.seq.filter {
      mount => mount._1 != null
  } toMap

val mapper = new ObjectMapper() with ScalaObjectMapper
mapper.registerModule(DefaultScalaModule)
mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false)

println(mapper.writeValueAsString(readableMounts))`

func (ic *importContext) getMountsThroughCluster(
	commandAPI common.CommandExecutor, clusterID string) (mm map[string]string, err error) {
	// Scala has actually working timeout handling, compared to Python
	result := commandAPI.Execute(clusterID, "scala", getReadableMountsCommand)
	if result.Failed() {
		err = result.Err()
		return
	}
	lines := strings.Split(result.Text(), "\n")
	err = json.Unmarshal([]byte(lines[0]), &mm)
	return
}

func eitherString(a interface{}, b interface{}) string {
	if a != nil {
		return a.(string)
	}
	if b != nil {
		return b.(string)
	}
	return ""
}

func (ic *importContext) importJobs(l jobs.JobList) {
	nowSeconds := time.Now().Unix()
	a := jobs.NewJobsAPI(ic.Context, ic.Client)
	starterAfter := (nowSeconds - (ic.lastActiveDays * 24 * 60 * 60)) * 1000
	i := 0
	for _, job := range l.Jobs {
		if !ic.MatchesName(job.Settings.Name) {
			log.Printf("[INFO] Job name %s doesn't match selection %s", job.Settings.Name, ic.match)
			continue
		}
		if ic.lastActiveDays != 3650 {
			rl, err := a.RunsList(jobs.JobRunsListRequest{
				JobID:         job.JobID,
				CompletedOnly: true,
				Limit:         1,
			})
			if err != nil {
				log.Printf("[WARN] Failed to get runs: %s", err)
				continue
			}
			if len(rl.Runs) == 0 {
				log.Printf("[INFO] Job %#v (%d) did never run. Skipping", job.Settings.Name, job.JobID)
				continue
			}
			if rl.Runs[0].StartTime < starterAfter {
				log.Printf("[INFO] Job %#v (%d) didn't run for %d days. Skipping",
					job.Settings.Name, job.JobID,
					(nowSeconds*1000-rl.Runs[0].StartTime)/24*60*60/1000)
				continue
			}
		}
		ic.Emit(&resource{
			Resource: "databricks_job",
			ID:       job.ID(),
		})
		i++
		log.Printf("[INFO] Imported %d of total %d jobs", i, len(l.Jobs))
	}
}
