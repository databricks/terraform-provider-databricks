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

	"github.com/databricks/databricks-sdk-go/service/catalog"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	if ic.match == "" {
		return true
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
	if strings.HasPrefix(path, "/Workspace/") {
		normalPath := strings.TrimPrefix(path, "/Workspace")
		ic.emitWorkspaceFileOrRepo(normalPath)
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
		return a.(string)
	}
	if b != nil {
		return b.(string)
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

func defaultShouldOmitFieldFunc(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
	if as.Computed {
		return true
	} else if as.Default != nil && d.Get(pathString) == as.Default {
		return true
	}

	return false
}

func resourceOrDataBlockBody(ic *importContext, body *hclwrite.Body, r *resource) error {
	blockType := "resource"
	if r.Mode == "data" {
		blockType = r.Mode
	}
	resourceBlock := body.AppendNewBlock(blockType, []string{r.Resource, r.Name})
	return ic.dataToHcl(ic.Importables[r.Resource],
		[]string{}, ic.Resources[r.Resource], r, resourceBlock.Body())
}

func generateUniqueID(v string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(v)))[:10]
}

func (ic *importContext) enableServices(services string) {
	ic.services = map[string]struct{}{}
	for _, s := range strings.Split(services, ",") {
		ic.services[strings.TrimSpace(s)] = struct{}{}
	}
	for s := range ic.listing { // Add all services mentioned in the listing
		ic.services[strings.TrimSpace(s)] = struct{}{}
	}
}

func (ic *importContext) enableListing(listing string) {
	ic.listing = map[string]struct{}{}
	for _, s := range strings.Split(listing, ",") {
		ic.listing[strings.TrimSpace(s)] = struct{}{}
		ic.services[strings.TrimSpace(s)] = struct{}{}
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

func shouldOmitForUnityCatalog(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
	if pathString == "owner" {
		return d.Get(pathString).(string) == ""
	}
	return defaultShouldOmitFieldFunc(ic, pathString, as, d)
}

func appendEndingSlashToDirName(dir string) string {
	if dir == "" || dir[len(dir)-1] == '/' {
		return dir
	}
	return dir + "/"
}

func isMatchingCatalogAndSchema(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	res_catalog_name := res.Data.Get("catalog_name").(string)
	res_schema_name := res.Data.Get("schema_name").(string)
	ra_catalog_name, cat_found := ra.Get("catalog_name")
	ra_schema_name, schema_found := ra.Get("name")
	if !cat_found || !schema_found {
		log.Printf("[WARN] Can't find attributes in approximation: %s %s, catalog='%v' (found? %v) schema='%v' (found? %v). Resource: %s, catalog='%s', schema='%s'",
			ra.Type, ra.Name, ra_catalog_name, cat_found, ra_schema_name, schema_found, res.Resource, res_catalog_name, res_schema_name)
		return true
	}
	result := ra_catalog_name.(string) == res_catalog_name && ra_schema_name.(string) == res_schema_name
	return result
}

func isMatchingCatalogAndSchemaInModelServing(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	res_catalog_name := res.Data.Get("config.0.auto_capture_config.0.catalog_name").(string)
	res_schema_name := res.Data.Get("config.0.auto_capture_config.0.schema_name").(string)
	ra_catalog_name, cat_found := ra.Get("catalog_name")
	ra_schema_name, schema_found := ra.Get("name")
	if !cat_found || !schema_found {
		log.Printf("[WARN] Can't find attributes in approximation: %s %s, catalog='%v' (found? %v) schema='%v' (found? %v). Resource: %s, catalog='%s', schema='%s'",
			ra.Type, ra.Name, ra_catalog_name, cat_found, ra_schema_name, schema_found, res.Resource, res_catalog_name, res_schema_name)
		return true
	}

	result := ra_catalog_name.(string) == res_catalog_name && ra_schema_name.(string) == res_schema_name
	return result
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

func isMatchingAllowListArtifact(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	objPath := strings.Replace(origPath, ".artifact", ".match_type", 1)
	matchType, ok := res.Data.GetOk(objPath)
	artifactType := res.Data.Get("artifact_type").(string)
	return ok && matchType.(string) == "PREFIX_MATCH" && (artifactType == "LIBRARY_JAR" || artifactType == "INIT_SCRIPT")
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

func (ic *importContext) emitUCGrantsWithOwner(id string, parentResource *resource) (string, *resource) {
	gr := &resource{
		Resource: "databricks_grants",
		ID:       id,
	}
	var owner string
	if parentResource.Data != nil {
		ignoreFunc := ic.Importables[parentResource.Resource].Ignore
		if ignoreFunc != nil && ignoreFunc(ic, parentResource) {
			return "", nil
		}
		ownerRaw, ok := parentResource.Data.GetOk("owner")
		if ok {
			gr.AddExtraData("owner", ownerRaw)
			owner = ownerRaw.(string)
		}
	}
	ic.Emit(gr)
	return owner, gr
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

func dltIsMatchingCatalogAndSchema(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	res_catalog_name := res.Data.Get("catalog").(string)
	if res_catalog_name == "" {
		return false
	}
	res_schema_name := res.Data.Get("target").(string)
	ra_catalog_name, cat_found := ra.Get("catalog_name")
	ra_schema_name, schema_found := ra.Get("name")
	if !cat_found || !schema_found {
		log.Printf("[WARN] Can't find attributes in approximation: %s %s, catalog='%v' (found? %v) schema='%v' (found? %v). Resource: %s, catalog='%s', schema='%s'",
			ra.Type, ra.Name, ra_catalog_name, cat_found, ra_schema_name, schema_found, res.Resource, res_catalog_name, res_schema_name)
		return true
	}

	result := ra_catalog_name.(string) == res_catalog_name && ra_schema_name.(string) == res_schema_name
	return result
}

func (ic *importContext) emitWorkspaceBindings(securableType, securableName string) {
	bindings, err := ic.workspaceClient.WorkspaceBindings.GetBindingsAll(ic.Context, catalog.GetBindingsRequest{
		SecurableName: securableName,
		SecurableType: catalog.GetBindingsSecurableType(securableType),
	})
	if err != nil {
		log.Printf("[ERROR] listing %s bindings for %s: %s", securableType, securableName, err.Error())
		return
	}
	for _, binding := range bindings {
		id := fmt.Sprintf("%d|%s|%s", binding.WorkspaceId, securableType, securableName)
		// We were creating Data instance explicitly because of the bug in the databricks_catalog_workspace_binding
		// implementation. Technically, after the fix is merged we can remove this, but we're keeping it as-is now
		// to decrease a number of API calls.
		d := ic.Resources["databricks_workspace_binding"].Data(
			&terraform.InstanceState{
				ID: id,
				Attributes: map[string]string{
					"workspace_id":   fmt.Sprintf("%d", binding.WorkspaceId),
					"securable_type": securableType,
					"securable_name": securableName,
					"binding_type":   binding.BindingType.String(),
				},
			})
		ic.Emit(&resource{
			Resource: "databricks_workspace_binding",
			ID:       id,
			Name:     fmt.Sprintf("%s_%s_ws_%d", securableType, securableName, binding.WorkspaceId),
			Data:     d,
		})
	}
}

func isMatchingSecurableTypeAndName(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	res_securable_type := res.Data.Get("securable_type").(string)
	res_securable_name := res.Data.Get("securable_name").(string)
	ra_name, _ := ra.Get("name")
	return ra.Type == ("databricks_"+res_securable_type) && ra_name.(string) == res_securable_name
}
