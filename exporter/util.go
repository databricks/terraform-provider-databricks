package exporter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/databricks/terraform-provider-databricks/aws"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/sql"
	"github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/databricks/terraform-provider-databricks/storage"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (ic *importContext) importCluster(c *clusters.Cluster) {
	if c == nil {
		return
	}
	// TODO: Remove PYSPARK_PYTHON from spark_env_vars & spark.databricks.delta.preview.enabled from spark_conf
	// but don't forget to convert back from structure to data
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
		// set enable_elastic_disk to false, and remove aws/gcp/azure_attributes
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       c.InstancePoolID,
		})
	}
	if c.DriverInstancePoolID != "" {
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       c.DriverInstancePoolID,
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
	common.DataToStructPointer(d, s, &cll)
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

func (ic *importContext) getSqlEndpoint(dataSourceId string) (string, error) {
	if ic.sqlDatasources == nil {
		var dss []sql.DataSource
		err := ic.Client.Get(ic.Context, "/preview/sql/data_sources", nil, &dss)
		if err != nil {
			return "", err
		}
		ic.sqlDatasources = make(map[string]string, len(dss))
		for _, ds := range dss {
			ic.sqlDatasources[ds.ID] = ds.EndpointID
		}
	}
	endpointID, ok := ic.sqlDatasources[dataSourceId]
	if !ok {
		return "", fmt.Errorf("can't find data source for SQL endpoint %s", dataSourceId)
	}

	return endpointID, nil
}

func (ic *importContext) getSqlVisualizations() (map[string]string, error) {
	if ic.sqlVisualizations == nil {
		ic.sqlVisualizations = make(map[string]string)
		qs, err := dbsqlListObjects(ic, "/preview/sql/queries")
		if err != nil {
			return nil, err
		}
		queryApi := sql.NewQueryAPI(ic.Context, ic.Client)
		for _, q := range qs {
			queryID := q["id"].(string)
			fullQuery, err := queryApi.Read(queryID)
			if err != nil {
				log.Printf("[WARN] Problems getting query with ID: %s", queryID)
				continue
			}
			for _, rv := range fullQuery.Visualizations {
				var vis api.Visualization
				err = json.Unmarshal(rv, &vis)
				if err != nil {
					log.Printf("[WARN] Problems decoding visualization for query with ID: %s", queryID)
					continue
				}
				ic.sqlVisualizations[vis.ID.String()] = queryID + "/" + vis.ID.String()
			}
		}
	}

	return ic.sqlVisualizations, nil
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
			ic.addAwsMounts(instanceProfile.InstanceProfileArn, profileMountMap)
		}
	}
	return nil
}

func (ic *importContext) addAwsMounts(arn string, profileMountMap map[string]string) {
	i := 0
	for k, v := range profileMountMap {
		if _, has := ic.mountMap[k]; has {
			continue
		}
		i++
		ic.mountMap[k] = mount{
			URL:             v,
			InstanceProfile: arn,
		}
	}
	if i > 0 {
		log.Printf("[INFO] Found %d mounts accessible by %s", len(profileMountMap), arn)
	}
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

func eitherString(a any, b any) string {
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

// returns created file name in "files" directory for the export and error if any
func (ic *importContext) createFile(name string, content []byte) (string, error) {
	return ic.createFileIn("files", name, content)
}

func (ic *importContext) createFileIn(dir, name string, content []byte) (string, error) {
	fileName := ic.prefix + name
	localFileName := fmt.Sprintf("%s/%s/%s", ic.Directory, dir, fileName)
	err := os.MkdirAll(path.Dir(localFileName), 0755)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	local, err := os.Create(localFileName)
	if err != nil {
		return "", err
	}
	defer local.Close()
	_, err = local.Write(content)
	if err != nil {
		return "", err
	}
	relativeName := strings.Replace(localFileName, ic.Directory+"/", "", 1)
	return relativeName, nil
}
