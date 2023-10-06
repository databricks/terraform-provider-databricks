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

	"github.com/databricks/databricks-sdk-go/service/iam"

	"golang.org/x/exp/slices"

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
	// TODO: think about another way of checking for a user. ideally we need to check against the
	// list of users/SPs obtained via SCIM API - this will be done in the refactoring requested by the SCIM team
	if strings.Contains(userOrSPName, "@") {
		ic.Emit(&resource{
			Resource:  "databricks_user",
			Attribute: "user_name",
			Value:     userOrSPName,
		})
	} else if common.StringIsUUID(userOrSPName) {
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
		// TODO: think about another way of checking for a user. ideally we need to check against the
		// list of users/SPs obtained via SCIM API - this will be done in the refactoring requested by the SCIM team
		userOrSPName := parts[2]
		return strings.Contains(userOrSPName, "@") || common.StringIsUUID(userOrSPName)
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
		ic.wsObjectsMutex.Lock()
		defer ic.wsObjectsMutex.Unlock()
		if len(ic.allDirectories) == 0 {
			for _, v := range objects {
				if v.ObjectType == workspace.Directory {
					ic.allDirectories = append(ic.allDirectories, v)
				}
			}
		}
	}
	return ic.allDirectories
}

// TODO: Ignore databricks_automl as well?
var directoriesToIgnore = []string{".ide", ".bundle", "__pycache__"}

func excludeAuxiliaryDirectories(v workspace.ObjectStatus) bool {
	if v.ObjectType != workspace.Directory {
		return true
	}
	parts := strings.Split(v.Path, "/")
	result := len(parts) > 1 && slices.Contains[[]string, string](directoriesToIgnore, parts[len(parts)-1])
	if result {
		log.Printf("[DEBUG] Ignoring directory %s", v.Path)
	}
	return !result
}

func (ic *importContext) getAllWorkspaceObjects() []workspace.ObjectStatus {
	ic.wsObjectsMutex.Lock()
	defer ic.wsObjectsMutex.Unlock()
	if len(ic.allWorkspaceObjects) == 0 {
		t1 := time.Now()
		log.Printf("[DEBUG] %v. Starting to list all workspace objects", t1.Local().Format(time.RFC3339))
		notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
		ic.allWorkspaceObjects, _ = notebooksAPI.ListParallel("/", excludeAuxiliaryDirectories)
		t2 := time.Now()
		log.Printf("[DEBUG] %v. Finished listing of all workspace objects. %d objects in total. %v seconds",
			t2.Local().Format(time.RFC3339), len(ic.allWorkspaceObjects), t2.Sub(t1).Seconds())
	}
	return ic.allWorkspaceObjects
}

func (ic *importContext) emitGroups(u scim.User) {
	for _, g := range u.Groups {
		if g.Type != "direct" {
			log.Printf("[DEBUG] Skipping non-direct group %s/%s for %s", g.Value, g.Display, u.DisplayName)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_group",
			ID:       g.Value,
		})
		ic.Emit(&resource{
			Resource: "databricks_group_member",
			ID:       fmt.Sprintf("%s|%s", g.Value, u.ID),
			Name:     fmt.Sprintf("%s_%s_%s_%s", g.Display, g.Value, u.DisplayName, u.ID),
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
	ic.groupsMutex.Lock()
	defer ic.groupsMutex.Unlock()
	if ic.allGroups == nil {
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

const (
	nonExistingUserOrSp = "__USER_OR_SPN_DOES_NOT_EXIST__"
)

func (ic *importContext) getUsersMapping() {
	ic.usersMutex.Lock()
	defer ic.usersMutex.Unlock()
	if ic.allUsersMapping == nil {
		ic.allUsersMapping = make(map[string]string)
		var users []iam.User
		var err error
		if ic.accountLevel {
			users, err = ic.accountClient.Users.ListAll(ic.Context, iam.ListAccountUsersRequest{
				Attributes: "userName,id",
			})
		} else {
			users, err = ic.workspaceClient.Users.ListAll(ic.Context, iam.ListUsersRequest{
				Attributes: "userName,id",
			})
		}
		if err != nil {
			log.Printf("[WARN] can't fetch list of users")
			return
		}
		for _, user := range users {
			// log.Printf("[DEBUG] adding user %v into the map. %d out of %d", user, i+1, len(users))
			ic.allUsersMapping[user.UserName] = user.Id
		}
	}
}

func (ic *importContext) findUserByName(name string) (u scim.User, err error) {
	log.Printf("[DEBUG] Looking for user %s", name)
	ic.usersMutex.RLocker().Lock()
	user, exists := ic.allUsers[name]
	ic.usersMutex.RLocker().Unlock()
	if exists {
		if user.UserName == nonExistingUserOrSp {
			log.Printf("[DEBUG] non-existing user %s is found in the cache", name)
			err = fmt.Errorf("user %s is not found", name)
		} else {
			log.Printf("[DEBUG] existing user %s is found in the cache", name)
			u = user
		}
		return
	}
	ic.getUsersMapping()
	ic.usersMutex.RLocker().Lock()
	userId, exists := ic.allUsersMapping[name]
	ic.usersMutex.RLocker().Unlock()
	if !exists {
		err = fmt.Errorf("there is no user '%s'", name)
		u = scim.User{UserName: nonExistingUserOrSp}
	} else {
		a := scim.NewUsersAPI(ic.Context, ic.Client)
		u, err = a.Read(userId, "id,userName,displayName,active,externalId,entitlements,groups,roles")
		if err != nil {
			log.Printf("[WARN] error reading user with name '%s', user ID: %s", name, userId)
			u = scim.User{UserName: nonExistingUserOrSp}
		}
	}
	ic.usersMutex.Lock()
	defer ic.usersMutex.Unlock()
	ic.allUsers[name] = u
	return
}

func (ic *importContext) getSpsMapping() {
	ic.spsMutex.Lock()
	defer ic.spsMutex.Unlock()
	if ic.allSpsMapping == nil {
		ic.allSpsMapping = make(map[string]string)
		var sps []iam.ServicePrincipal
		var err error
		// Reimplement it myself
		if ic.accountLevel {
			sps, err = ic.accountClient.ServicePrincipals.ListAll(ic.Context, iam.ListAccountServicePrincipalsRequest{
				ExcludedAttributes: "groups,roles,entitlements",
			})
		} else {
			sps, err = ic.workspaceClient.ServicePrincipals.ListAll(ic.Context, iam.ListServicePrincipalsRequest{
				ExcludedAttributes: "groups,roles,entitlements",
			})
		}
		if err != nil {
			log.Printf("[WARN] can't fetch list of service principals")
			return
		}
		for _, sp := range sps {
			log.Printf("[DEBUG] adding sp %v", sp)
			ic.allSpsMapping[sp.ApplicationId] = sp.Id
		}
	}
}

func (ic *importContext) findSpnByAppID(applicationID string) (u scim.User, err error) {
	log.Printf("[DEBUG] Looking for SP %s", applicationID)
	ic.spsMutex.RLocker().Lock()
	sp, exists := ic.allSps[applicationID]
	ic.spsMutex.RLocker().Unlock()
	if exists {
		if sp.ApplicationID == nonExistingUserOrSp {
			log.Printf("[DEBUG] non-existing SP %s is found in the cache", applicationID)
			err = fmt.Errorf("service principal %s is not found", applicationID)
		} else {
			log.Printf("[DEBUG] existing SP %s is found in the cache", applicationID)
			u = sp
		}
		return
	}
	ic.getSpsMapping()
	ic.spsMutex.RLocker().Lock()
	spId, exists := ic.allSpsMapping[applicationID]
	ic.spsMutex.RLocker().Unlock()
	if !exists {
		err = fmt.Errorf("there is no service principal '%s'", applicationID)
		u = scim.User{ApplicationID: nonExistingUserOrSp}
	} else {
		a := scim.NewServicePrincipalsAPI(ic.Context, ic.Client)
		u, err = a.Read(spId)
		if err != nil {
			log.Printf("[WARN] error reading service principal with AppID '%s', SP ID: %s", applicationID, spId)
			u = scim.User{ApplicationID: nonExistingUserOrSp}
		}
	}
	log.Printf("[DEBUG] Adding SP with full details: %v", u)
	ic.spsMutex.Lock()
	defer ic.spsMutex.Unlock()
	ic.allSps[applicationID] = u

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
	ic.sqlDatasourcesMutex.Lock()
	defer ic.sqlDatasourcesMutex.Unlock()
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
	i := 0
	for offset, job := range l {
		if !ic.MatchesName(job.Settings.Name) {
			log.Printf("[INFO] Job name %s doesn't match selection %s", job.Settings.Name, ic.match)
			continue
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
		case prefix + "spark_conf":
			return fmt.Sprintf("%v", d.Get(prefix+"spark_conf")) == "map[spark.databricks.delta.preview.enabled:true]"
		case prefix + "spark_env_vars":
			return fmt.Sprintf("%v", d.Get(prefix+"spark_env_vars")) == "map[PYSPARK_PYTHON:/databricks/python3/bin/python3]"
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
		// TODO: can we pass a visitor here, that will emit corresponding object earlier?
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
			if ic.incremental && modifiedAt < updatedSinceMs {
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
	return ic.updatedSinceMs
}

func getEnvAsInt(envName string, defaultValue int) int {
	if val, exists := os.LookupEnv(envName); exists {
		parsedVal, err := strconv.Atoi(val)
		if err == nil {
			return parsedVal
		}
		log.Printf("[ERROR] Can't parse value '%s' of environment variable '%s'", val, envName)
	}
	return defaultValue
}
