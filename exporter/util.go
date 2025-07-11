package exporter

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/databricks/terraform-provider-databricks/aws"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/storage"
	"golang.org/x/exp/maps"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (ic *importContext) isServiceEnabled(service string) bool {
	_, exists := ic.services[service]
	return exists
}

func (ic *importContext) isServiceInListing(service string) bool {
	_, exists := ic.listing[service]
	return exists
}

func (ic *importContext) MatchesName(n string) bool {
	if ic.match == "" && ic.matchRegex == nil && ic.excludeRegex == nil {
		return true
	}
	if ic.excludeRegex != nil && ic.excludeRegex.MatchString(n) {
		return false
	}
	if ic.matchRegex != nil {
		return ic.matchRegex.MatchString(n)
	}
	return strings.Contains(strings.ToLower(n), strings.ToLower(ic.match))
}

func (ic *importContext) emitFilesFromSlice(slice []string) {
	for _, p := range slice {
		ic.emitIfDbfsFile(p)
		ic.emitIfWsfsFile(p)
		ic.emitIfVolumeFile(p)
	}
}

func (ic *importContext) emitFilesFromMap(m map[string]string) {
	for _, p := range m {
		ic.emitIfDbfsFile(p)
		ic.emitIfWsfsFile(p)
		ic.emitIfVolumeFile(p)
	}
}

func (ic *importContext) addIgnoredResource(msg string) {
	ic.ignoredResourcesMutex.Lock()
	defer ic.ignoredResourcesMutex.Unlock()
	ic.ignoredResources[msg] = struct{}{}
}

func (ic *importContext) emitIfDbfsFile(path string) {
	if strings.HasPrefix(path, "dbfs:") {
		if strings.HasPrefix(path, "dbfs:/Volumes/") {
			ic.emitIfVolumeFile(path[5:])
		} else {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       path,
			})
		}
	}
}

func (ic *importContext) emitIfWsfsFile(path string) {
	if hasWorkspacePrefix(path) {
		ic.emitWorkspaceFileOrRepo(maybeStripWorkspacePrefix(path))
	}
}

func (ic *importContext) emitIfVolumeFile(path string) {
	if strings.HasPrefix(path, "/Volumes/") {
		ic.Emit(&resource{
			Resource: "databricks_file",
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
		dss, err := ic.workspaceClient.DataSources.List(ic.Context)
		if err != nil {
			return nil, err
		}
		ic.sqlDatasources = make(map[string]string, len(dss))
		for _, ds := range dss {
			ic.sqlDatasources[ds.Id] = ds.WarehouseId
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
		if t, ok := a.(string); ok {
			return t
		}
	}
	if b != nil {
		if t, ok := b.(string); ok {
			return t
		}
	}
	return ""
}

func (ic *importContext) createFileIn(dir, name string) (*os.File, string, error) {
	fileName := ic.prefix + name
	localFileName := fmt.Sprintf("%s/%s/%s", ic.Directory, dir, fileName)
	err := os.MkdirAll(path.Dir(localFileName), 0755)
	if err != nil && !os.IsExist(err) {
		return nil, "", err
	}
	local, err := os.Create(localFileName)
	if err != nil {
		return nil, "", err
	}
	relativeName := strings.TrimPrefix(localFileName, ic.Directory+"/")
	return local, relativeName, nil
}

func (ic *importContext) saveFileIn(dir, name string, content []byte) (string, error) {
	local, relativeName, err := ic.createFileIn(dir, name)
	if err != nil {
		return "", err
	}
	defer local.Close()
	_, err = local.Write(content)
	if err != nil {
		return "", err
	}
	return relativeName, nil
}

func defaultShouldOmitFieldFunc(_ *importContext, pathString string, as *schema.Schema, d *schema.ResourceData, _ *resource) bool {
	if as.Computed {
		return true
	} else if as.Default != nil && d.Get(pathString) == as.Default {
		return true
	}

	return false
}

func (ic *importContext) generateNewData(data *schema.ResourceData, resourceType, rID string, obj any) *schema.ResourceData {
	data.MarkNewResource()
	data.SetId(rID)
	scm := ic.Resources[resourceType].Schema
	err := common.StructToData(obj, scm, data)
	if err != nil {
		log.Printf("[ERROR] can't convert %s object to data: %v. obj=%v", resourceType, err, obj)
		return nil
	}
	return data
}

func generateUniqueID(v string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(v)))[:10]
}

func (ic *importContext) allServicesAndListing() (string, string) {
	services := map[string]struct{}{}
	listing := map[string]struct{}{}
	for _, ir := range ic.Importables {
		services[ir.Service] = struct{}{}
		if ir.List != nil {
			listing[ir.Service] = struct{}{}
		}
	}
	// We need this to specify default listings of UC & Workspace objects...
	for _, ir := range []string{"uc-schemas", "uc-models", "uc-tables", "uc-volumes",
		"notebooks", "directories", "wsfiles"} {
		listing[ir] = struct{}{}
	}
	return strings.Join(maps.Keys(services), ","), strings.Join(maps.Keys(listing), ",")
}

func (ic *importContext) parseServicesList(services string, isListing bool) []string {
	allEnabledServices, allEnabledListing := ic.allServicesAndListing()
	var allServices []string
	if isListing {
		allServices = strings.Split(allEnabledListing, ",")
	} else {
		allServices = strings.Split(allEnabledServices, ",")
	}
	var allUcServices []string
	for _, s := range allServices {
		if strings.HasPrefix(s, "uc-") {
			allUcServices = append(allUcServices, s)
		}
	}
	allUcServices = append(allUcServices, "vector-search")
	servicesList := map[string]struct{}{}
	for _, s := range strings.Split(services, ",") {
		ss := strings.TrimSpace(s)
		if ss == "all" {
			for _, service := range allServices {
				servicesList[service] = struct{}{}
			}
		} else if ss == "+uc" || ss == "uc" {
			for _, service := range allUcServices {
				servicesList[service] = struct{}{}
			}

		} else if ss == "-uc" {
			for _, service := range allUcServices {
				delete(servicesList, service)
			}
		} else if strings.HasPrefix(ss, "-") {
			delete(servicesList, ss[1:])
		} else if strings.HasPrefix(ss, "+") {
			servicesList[ss[1:]] = struct{}{}
		} else if ss != "" {
			servicesList[ss] = struct{}{}
		}
	}
	return maps.Keys(servicesList)
}

func (ic *importContext) enableServices(services string) {
	ic.services = map[string]struct{}{}
	for _, s := range ic.parseServicesList(services, false) {
		ic.services[s] = struct{}{}
	}
	for s := range ic.listing { // Add all services mentioned in the listing
		ic.services[s] = struct{}{}
	}
}

func (ic *importContext) enableListing(listing string) {
	ic.listing = map[string]struct{}{}
	for _, s := range ic.parseServicesList(listing, true) {
		ic.listing[s] = struct{}{}
		ic.services[s] = struct{}{}
	}
}

func (ic *importContext) emitSqlParentDirectory(parent string) {
	if parent == "" || !ic.isServiceEnabled("directories") {
		return
	}
	res := sqlParentRegexp.FindStringSubmatch(parent)
	if len(res) > 1 {
		ic.Emit(&resource{
			Resource:  "databricks_directory",
			Attribute: "object_id",
			Value:     res[1],
		})
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

var (
	maxRetries        = 5
	retryDelaySeconds = 2
	retriableErrors   = []string{"deadline exceeded", "Error handling request", "Timed out after ", "Operation timed out"}
)

func isRetryableError(err string, i int) bool {
	if i < (maxRetries - 1) {
		for _, msg := range retriableErrors {
			if strings.Contains(err, msg) {
				return true
			}
		}
	}
	return false
}

func runWithRetries[ERR any](runFunc func() ERR, msg string) ERR {
	var err ERR
	delay := 1
	for i := 0; i < maxRetries; i++ {
		err = runFunc()
		valOf := reflect.ValueOf(&err).Elem()
		if valOf.IsNil() || valOf.IsZero() {
			break
		}
		if !isRetryableError(fmt.Sprintf("%v", err), i) {
			log.Printf("[ERROR] Error %s after %d retries: %v", msg, i, err)
			return err
		}
		delay = delay * retryDelaySeconds
		log.Printf("[INFO] next retry (%d) for %s after %d seconds", (i + 1), msg, delay)
		time.Sleep(time.Duration(delay) * time.Second)
	}
	return err
}

func appendEndingSlashToDirName(dir string) string {
	if dir == "" || dir[len(dir)-1] == '/' {
		return dir
	}
	return dir + "/"
}

func isMatchingShareRecipient(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	shareName, ok := res.Data.GetOk("share")
	return ok && shareName.(string) != ""
}

func isMatchignShareObject(obj string) isValidAproximationFunc {
	return func(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
		objPath := strings.Replace(origPath, ".name", ".data_object_type", 1)
		objType, ok := res.Data.GetOk(objPath)
		return ok && objType.(string) == obj
	}
}

func generateIgnoreObjectWithEmptyAttributeValue(resourceType, attrName string) func(ic *importContext, r *resource) bool {
	return func(ic *importContext, r *resource) bool {
		res := (r.Data != nil && r.Data.Get(attrName).(string) == "")
		if res {
			ic.addIgnoredResource(fmt.Sprintf("%s. id=%s", resourceType, r.ID))
		}
		return res
	}
}

func (ic *importContext) addTfVar(name, value string) {
	ic.tfvarsMutex.Lock()
	defer ic.tfvarsMutex.Unlock()
	ic.tfvars[name] = value
}

func (ic *importContext) emitPermissionsIfNotIgnored(r *resource, id, name string) {
	if ic.meAdmin {
		ignoreFunc := ic.Importables[r.Resource].Ignore
		if ignoreFunc == nil || !ignoreFunc(ic, r) {
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       id,
				Name:     name,
			})
		}
	}
}

func makeNamePlusIdFunc(nm string) func(ic *importContext, d *schema.ResourceData) string {
	return func(ic *importContext, d *schema.ResourceData) string {
		return d.Get(nm).(string) + "_" + d.Id()
	}
}

func makeNameOrIdFunc(nm string) func(ic *importContext, d *schema.ResourceData) string {
	return func(ic *importContext, d *schema.ResourceData) string {
		name := d.Get(nm).(string)
		if name == "" {
			return d.Id()
		}
		return name
	}
}
