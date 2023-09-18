package exporter

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/databricks/terraform-provider-databricks/aws"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/sql"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/workspace"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (ic *importContext) emitInitScripts(initScripts []clusters.InitScriptStorageInfo) {
	for _, is := range initScripts {
		if is.Dbfs != nil {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       is.Dbfs.Destination,
			})
		}
		if is.Workspace != nil {
			ic.Emit(&resource{
				Resource: "databricks_workspace_file",
				ID:       is.Workspace.Destination,
			})
		}
	}

}

func (ic *importContext) importCluster(c *clusters.Cluster) {
	if c == nil {
		return
	}
	ic.emitInitScripts(c.InitScripts)
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
	ic.emitSecretsFromSecretsPath(c.SparkConf)
	ic.emitSecretsFromSecretsPath(c.SparkEnvVars)
	ic.emitUserOrServicePrincipal(c.SingleUserName)
}

func (ic *importContext) emitSecretsFromSecretsPath(m map[string]string) {
	for _, v := range m {
		if res := secretPathRegex.FindStringSubmatch(v); res != nil {
			ic.Emit(&resource{
				Resource: "databricks_secret_scope",
				ID:       res[1],
			})
		}
	}
}

func (ic *importContext) emitUserOrServicePrincipal(userOrSPName string) {
	if userOrSPName == "" {
		return
	}
	if strings.Contains(userOrSPName, "@") {
		ic.Emit(&resource{
			Resource:  "databricks_user",
			Attribute: "user_name",
			Value:     userOrSPName,
		})
	} else if uuidRegex.MatchString(userOrSPName) {
		ic.Emit(&resource{
			Resource:  "databricks_service_principal",
			Attribute: "application_id",
			Value:     userOrSPName,
		})
	}
}

func (ic *importContext) emitUserOrServicePrincipalForPath(path, prefix string) {
	if strings.HasPrefix(path, prefix) {
		parts := strings.SplitN(path, "/", 4)
		if len(parts) >= 3 {
			ic.emitUserOrServicePrincipal(parts[2])
		}
	}
}

func (ic *importContext) IsUserOrServicePrincipalDirectory(path, prefix string) bool {
	if !strings.HasPrefix(path, prefix) {
		return false
	}
	parts := strings.SplitN(path, "/", 4)
	if len(parts) == 3 || (len(parts) == 4 && parts[3] == "") {
		userOrSPName := parts[2]
		return strings.Contains(userOrSPName, "@") || uuidRegex.MatchString(userOrSPName)
	}
	return false
}

func (ic *importContext) emitNotebookOrRepo(path string) {
	if strings.HasPrefix(path, "/Repos") {
		ic.Emit(&resource{
			Resource:  "databricks_repo",
			Attribute: "path",
			Value:     strings.Join(strings.SplitN(path, "/", 5)[:4], "/"),
		})
	} else {
		ic.Emit(&resource{
			Resource: "databricks_notebook",
			ID:       path,
		})
	}
}

func (ic *importContext) getAllDirectories() []workspace.ObjectStatus {
	if len(ic.allDirectories) == 0 {
		objects := ic.getAllWorkspaceObjects()
		for _, v := range objects {
			if v.ObjectType == workspace.Directory {
				ic.allDirectories = append(ic.allDirectories, v)
			}
		}
	}
	return ic.allDirectories
}

func (ic *importContext) getAllWorkspaceObjects() []workspace.ObjectStatus {
	if len(ic.allWorkspaceObjects) == 0 {
		t1 := time.Now()
		log.Printf("[DEBUG] %v. Starting to list all workspace objects", t1.Local().Format(time.RFC3339))
		notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
		ic.allWorkspaceObjects, _ = notebooksAPI.ListParallel("/", true, true)
		t2 := time.Now()
		log.Printf("[DEBUG] %v. Finished listing of all workspace objects. %d objects in total. %v seconds",
			t2.Local().Format(time.RFC3339), len(ic.allWorkspaceObjects), t2.Sub(t1).Seconds())
	}
	return ic.allWorkspaceObjects
}

func (ic *importContext) emitGroups(u scim.User, principal string) {
	for _, g := range u.Groups {
		if g.Type != "direct" {
			log.Printf("[DEBUG] Skipping non-direct group %s/%s for %s", g.Value, g.Display, principal)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_group",
			ID:       g.Value,
		})
		ic.Emit(&resource{
			Resource: "databricks_group_member",
			ID:       fmt.Sprintf("%s|%s", g.Value, u.ID),
			Name:     fmt.Sprintf("%s_%s_%s", g.Display, g.Value, principal),
		})
	}
}

func (ic *importContext) emitRoles(objType string, id string, roles []scim.ComplexValue) {
	log.Printf("[DEBUG] emitting roles for object type: %s, ID: %s, roles: %v", objType, id, roles)
	for _, role := range roles {
		if role.Type != "direct" {
			continue
		}
		if !ic.accountLevel {
			ic.Emit(&resource{
				Resource: "databricks_instance_profile",
				ID:       role.Value,
			})
		}
		ic.Emit(&resource{
			Resource: fmt.Sprintf("databricks_%s_role", objType),
			ID:       fmt.Sprintf("%s|%s", id, role.Value),
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

func (ic *importContext) importClusterLibraries(d *schema.ResourceData, s map[string]*schema.Schema) error {
	cll, err := libraries.NewLibrariesAPI(ic.Context, ic.Client).ClusterStatus(d.Id())
	if err != nil {
		return err
	}
	if cll.LibraryStatuses != nil {
		for _, lib := range cll.LibraryStatuses {
			ic.Emit(&resource{
				Resource: "databricks_library",
				ID:       lib.Library.GetID(d.Id()),
			})
			ic.emitIfDbfsFile(lib.Library.Egg)
			ic.emitIfDbfsFile(lib.Library.Jar)
			ic.emitIfDbfsFile(lib.Library.Whl)
		}
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
	users, err := a.Filter(fmt.Sprintf("userName eq '%s'", name), false)
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

func (ic *importContext) findSpnByAppID(applicationID string) (u scim.User, err error) {
	a := scim.NewServicePrincipalsAPI(ic.Context, ic.Client)
	users, err := a.Filter(fmt.Sprintf("applicationId eq '%s'", strings.ReplaceAll(applicationID, "'", "")), false)
	if err != nil {
		return
	}
	if len(users) == 0 {
		err = fmt.Errorf("service principal %s not found", applicationID)
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

// todo: generic with go1.18
type dbsqlListResponse struct {
	Results    []map[string]any `json:"results"`
	Page       int64            `json:"page"`
	TotalCount int64            `json:"count"`
	PageSize   int64            `json:"page_size"`
}

// Generic function to list objects related to the DBSQL
func dbsqlListObjects(ic *importContext, path string) (events []map[string]any, err error) {
	// TODO: create API method & use it also for data resource
	var listResponse dbsqlListResponse
	page_size := 100
	err = ic.Client.Get(ic.Context, path, map[string]any{"page_size": page_size}, &listResponse)
	if err != nil {
		return nil, err
	}
	totalCount := int(listResponse.TotalCount)
	if totalCount == 0 {
		return events, nil
	}
	events = append(events, listResponse.Results...)
	page := 2
	for len(events) < totalCount {
		var listResponse dbsqlListResponse
		err := ic.Client.Get(ic.Context, path,
			map[string]any{"page_size": page_size, "page": page},
			&listResponse)
		if err != nil {
			return nil, err
		}
		events = append(events, listResponse.Results...)
		page++
	}
	return events, err
}

func (ic *importContext) getSqlDataSources() (map[string]string, error) {
	if ic.sqlDatasources == nil {
		var dss []sql.DataSource
		err := ic.Client.Get(ic.Context, "/preview/sql/data_sources", nil, &dss)
		if err != nil {
			return nil, err
		}
		ic.sqlDatasources = make(map[string]string, len(dss))
		for _, ds := range dss {
			ic.sqlDatasources[ds.ID] = ds.EndpointID
		}
	}
	return ic.sqlDatasources, nil
}

func (ic *importContext) getSqlEndpoint(dataSourceId string) (string, error) {
	sources, err := ic.getSqlDataSources()
	if err != nil {
		return "", err
	}
	endpointID, ok := sources[dataSourceId]
	if !ok {
		return "", fmt.Errorf("can't find data source for SQL endpoint %s", dataSourceId)
	}
	return endpointID, nil
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

func (ic *importContext) importJobs(l []jobs.Job) {
	nowSeconds := time.Now().Unix()
	a := jobs.NewJobsAPI(ic.Context, ic.Client)
	starterAfter := (nowSeconds - (ic.lastActiveDays * 24 * 60 * 60)) * 1000
	i := 0
	for offset, job := range l {
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
		log.Printf("[INFO] Scanned %d of total %d jobs", offset+1, len(l))
	}
	log.Printf("[INFO] %d of total %d jobs are going to be imported", i, len(l))
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

func defaultShouldOmitFieldFunc(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
	if as.Computed {
		return true
	} else if as.Default != nil && d.Get(pathString) == as.Default {
		return true
	}

	return false
}

func makeShouldOmitFieldForCluster(regex *regexp.Regexp) func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
	return func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
		prefix := ""
		if regex != nil {
			if res := regex.FindStringSubmatch(pathString); res != nil {
				prefix = res[0]
			} else {
				return false
			}
		}
		raw := d.Get(pathString)
		if raw != nil {
			v := reflect.ValueOf(raw)
			if as.Optional && v.IsZero() {
				return true
			}
		}
		workerInstPoolID := d.Get(prefix + "instance_pool_id").(string)
		switch pathString {
		case prefix + "node_type_id":
			return workerInstPoolID != ""
		case prefix + "driver_node_type_id":
			driverInstPoolID := d.Get(prefix + "driver_instance_pool_id").(string)
			nodeTypeID := d.Get(prefix + "node_type_id").(string)
			return workerInstPoolID != "" || driverInstPoolID != "" || raw.(string) == nodeTypeID
		case prefix + "driver_instance_pool_id":
			return raw.(string) == workerInstPoolID
		case prefix + "enable_elastic_disk", prefix + "aws_attributes", prefix + "azure_attributes", prefix + "gcp_attributes":
			return workerInstPoolID != ""
		case prefix + "enable_local_disk_encryption":
			return false
		}

		return defaultShouldOmitFieldFunc(ic, pathString, as, d)
	}
}

func resourceOrDataBlockBody(ic *importContext, body *hclwrite.Body, r *resource) error {
	blockType := "resource"
	if r.Mode == "data" {
		blockType = r.Mode
	}
	resourceBlock := body.AppendNewBlock(blockType, []string{r.Resource, r.Name})
	return ic.dataToHcl(ic.Importables[r.Resource],
		[]string{}, ic.Resources[r.Resource], r.Data, resourceBlock.Body())
}

func generateUniqueID(v string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(v)))[:10]
}

func workspaceObjectResouceName(ic *importContext, d *schema.ResourceData) string {
	name := d.Get("path").(string)
	if name == "" {
		return d.Id()
	} else {
		name = nameNormalizationRegex.ReplaceAllString(name[1:], "_") + "_" +
			strconv.FormatInt(int64(d.Get("object_id").(int)), 10)
	}
	return name
}

func wsObjectGetModifiedAt(obs workspace.ObjectStatus) int64 {
	if obs.ModifiedAtInteractive != nil && obs.ModifiedAtInteractive.TimeMillis != 0 {
		return obs.ModifiedAtInteractive.TimeMillis
	}
	return obs.ModifiedAt
}

func createListWorkspaceObjectsFunc(objType string, resourceType string, objName string) func(ic *importContext) error {
	return func(ic *importContext) error {
		objectsList := ic.getAllWorkspaceObjects()
		updatedSinceMs := ic.getUpdatedSinceMs()
		for offset, object := range objectsList {
			if object.ObjectType != objType || strings.HasPrefix(object.Path, "/Repos") {
				continue
			}
			if res := ignoreIdeFolderRegex.FindStringSubmatch(object.Path); res != nil {
				continue
			}
			modifiedAt := wsObjectGetModifiedAt(object)
			if ic.incremental && modifiedAt != 0 && modifiedAt < updatedSinceMs {
				log.Printf("[DEBUG] skipping '%s' that was modified at %d (last active=%d)", object.Path,
					modifiedAt, updatedSinceMs)
				continue
			}
			if !ic.MatchesName(object.Path) {
				continue
			}
			ic.Emit(&resource{
				Resource:    resourceType,
				ID:          object.Path,
				Incremental: ic.incremental,
			})

			if offset%50 == 0 {
				log.Printf("[INFO] Scanned %d of %d %ss", offset+1, len(objectsList), objName)
			}
		}
		return nil
	}
}

func (ic *importContext) getLastActiveMs() int64 {
	if ic.lastActiveMs == 0 {
		ic.lastActiveMs = (time.Now().Unix() - ic.lastActiveDays*24*60*60) * 1000
	}
	return ic.lastActiveMs
}

func (ic *importContext) getUpdatedSinceStr() string {
	return ic.updatedSinceStr
}

func (ic *importContext) getUpdatedSinceMs() int64 {
	if ic.updatedSinceMs == 0 {
		tm, _ := time.Parse(time.RFC3339, ic.updatedSinceStr)
		ic.updatedSinceMs = tm.UnixMilli()
	}
	return ic.updatedSinceMs
}
