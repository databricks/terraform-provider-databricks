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
	"sync"
	"sync/atomic"
	"time"

	"github.com/databricks/terraform-provider-databricks/aws"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/workspace"

	"github.com/databricks/databricks-sdk-go/service/compute"
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
			ic.emitWorkspaceFileOrRepo(is.Workspace.Destination)
		}
		if is.Volumes != nil {
			ic.emitIfVolumeFile(is.Volumes.Destination)
		}
	}
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

func (ic *importContext) importCluster(c *clusters.Cluster) {
	if c == nil {
		return
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
	ic.emitInitScripts(c.InitScripts)
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

func (ic *importContext) emitListOfUsers(users []string) {
	for _, user := range users {
		if user != "" {
			ic.Emit(&resource{
				Resource:  "databricks_user",
				Attribute: "user_name",
				Value:     user,
			})
		}
	}
}

func (ic *importContext) emitUserOrServicePrincipal(userOrSPName string) {
	if userOrSPName == "" || !ic.isServiceEnabled("users") {
		return
	}
	// Cache check here to avoid emitting
	ic.emittedUsersMutex.RLock()
	_, exists := ic.emittedUsers[userOrSPName]
	ic.emittedUsersMutex.RUnlock()
	if exists {
		// log.Printf("[DEBUG] user or SP %s already emitted...", userOrSPName)
		return
	}
	if common.StringIsUUID(userOrSPName) {
		user, err := ic.findSpnByAppID(userOrSPName, false)
		if err != nil {
			log.Printf("[ERROR] Can't find SP with application ID %s", userOrSPName)
			ic.addIgnoredResource(fmt.Sprintf("databricks_service_principal. application_id=%s", userOrSPName))
		} else {
			ic.Emit(&resource{
				Resource: "databricks_service_principal",
				ID:       user.ID,
			})
		}
	} else {
		user, err := ic.findUserByName(strings.ToLower(userOrSPName), false)
		if err != nil {
			log.Printf("[ERROR] Can't find user with name %s", userOrSPName)
			ic.addIgnoredResource(fmt.Sprintf("databricks_user. user_name=%s", userOrSPName))
		} else {
			ic.Emit(&resource{
				Resource: "databricks_user",
				ID:       user.ID,
			})
		}
	}
	ic.emittedUsersMutex.Lock()
	ic.emittedUsers[userOrSPName] = struct{}{}
	ic.emittedUsersMutex.Unlock()
}

func getUserOrSpNameAndDirectory(path, prefix string) (string, string) {
	if !strings.HasPrefix(path, prefix) {
		return "", ""
	}
	pathLen := len(path)
	prefixLen := len(prefix)
	searchStart := prefixLen + 1
	if pathLen <= searchStart {
		return "", ""
	}
	pos := strings.Index(path[searchStart:pathLen], "/")
	if pos == -1 { // we have only user directory...
		return path[searchStart:pathLen], path
	}
	return path[searchStart : pos+searchStart], path[0 : pos+searchStart]
}

func (ic *importContext) emitUserOrServicePrincipalForPath(path, prefix string) {
	userOrSpName, _ := getUserOrSpNameAndDirectory(path, prefix)
	if userOrSpName != "" {
		ic.emitUserOrServicePrincipal(userOrSpName)
	}
}

func (ic *importContext) IsUserOrServicePrincipalDirectory(path, prefix string, strict bool) bool {
	userOrSPName, userDir := getUserOrSpNameAndDirectory(path, prefix)
	if userOrSPName == "" {
		return false
	}
	// strict mode means that it should be exactly user dir, maybe with trailing `/`
	if strict && !(len(path) == len(userDir) || (len(path) == len(userDir)+1 && path[len(path)-1] == '/')) {
		return false
	}
	ic.userOrSpDirectoriesMutex.RLock()
	result, exists := ic.userOrSpDirectories[userDir]
	ic.userOrSpDirectoriesMutex.RUnlock()
	if exists {
		// log.Printf("[DEBUG] Directory %s already checked. Result=%v", userDir, result)
		return result
	}
	var err error
	if common.StringIsUUID(userOrSPName) {
		_, err = ic.findSpnByAppID(userOrSPName, true)
		if err != nil {
			ic.addIgnoredResource(fmt.Sprintf("databricks_service_principal. application_id=%s", userOrSPName))
		}
	} else {
		_, err = ic.findUserByName(strings.ToLower(userOrSPName), true)
		if err != nil {
			ic.addIgnoredResource(fmt.Sprintf("databricks_user. user_name=%s", userOrSPName))
		}
	}
	ic.userOrSpDirectoriesMutex.Lock()
	ic.userOrSpDirectories[userDir] = (err == nil)
	ic.userOrSpDirectoriesMutex.Unlock()
	return err == nil
}

func (ic *importContext) emitRepoByPath(path string) {
	// Path to Repos objects consits of following parts: /Repos, folder, repository, path inside Repo.
	// Because it starts with `/`, it will produce empty string as first element in the slice.
	// And we're stopping splitting to avoid producing too many not necessary parts, so we have 5 parts only.
	parts := strings.SplitN(path, "/", 5)
	if len(parts) >= 4 {
		ic.Emit(&resource{
			Resource:  "databricks_repo",
			Attribute: "path",
			Value:     strings.Join(parts[:4], "/"),
		})
	} else {
		log.Printf("[WARN] Incorrect Repos path")
	}
}

func (ic *importContext) emitWorkspaceFileOrRepo(path string) {
	if strings.HasPrefix(path, "/Repos") {
		ic.emitRepoByPath(path)
	} else {
		// TODO: wrap this into ic.shouldEmit...
		// TODO: strip /Workspace prefix if it's provided
		ic.Emit(&resource{
			Resource: "databricks_workspace_file",
			ID:       path,
		})
	}
}

func (ic *importContext) emitNotebookOrRepo(path string) {
	if strings.HasPrefix(path, "/Repos") {
		ic.emitRepoByPath(path)
	} else {
		// TODO: strip /Workspace prefix if it's provided
		ic.maybeEmitWorkspaceObject("databricks_notebook", path)
	}
}

func (ic *importContext) getAllDirectories() []workspace.ObjectStatus {
	if len(ic.allDirectories) == 0 {
		objects := ic.getAllWorkspaceObjects(nil)
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

// TODO: add ignoring directories of deleted users?  This could potentially decrease the number of processed objects...
func excludeAuxiliaryDirectories(v workspace.ObjectStatus) bool {
	if v.ObjectType != workspace.Directory {
		return true
	}
	// TODO: rewrite to use suffix check, etc., instead of split and slice contains?
	parts := strings.Split(v.Path, "/")
	result := len(parts) > 1 && slices.Contains[[]string, string](directoriesToIgnore, parts[len(parts)-1])
	if result {
		log.Printf("[DEBUG] Ignoring directory %s", v.Path)
	}
	return !result
}

func (ic *importContext) getAllWorkspaceObjects(visitor func([]workspace.ObjectStatus)) []workspace.ObjectStatus {
	ic.wsObjectsMutex.Lock()
	defer ic.wsObjectsMutex.Unlock()
	if len(ic.allWorkspaceObjects) == 0 {
		t1 := time.Now()
		log.Print("[INFO] Starting to list all workspace objects")
		notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
		ic.allWorkspaceObjects, _ = ListParallel(notebooksAPI, "/", excludeAuxiliaryDirectories, visitor)
		log.Printf("[INFO] Finished listing of all workspace objects. %d objects in total. %v seconds",
			len(ic.allWorkspaceObjects), time.Since(t1).Seconds())
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

func (ic *importContext) emitLibraries(libs []libraries.Library) {
	for _, lib := range libs {
		// Files on DBFS
		ic.emitIfDbfsFile(lib.Whl)
		ic.emitIfDbfsFile(lib.Jar)
		ic.emitIfDbfsFile(lib.Egg)
		// Files on WSFS
		ic.emitIfWsfsFile(lib.Whl)
		ic.emitIfWsfsFile(lib.Jar)
		ic.emitIfWsfsFile(lib.Egg)
		// Files on UC Volumes
		ic.emitIfVolumeFile(lib.Whl)
		ic.emitIfVolumeFile(lib.Jar)
	}

}

func (ic *importContext) importLibraries(d *schema.ResourceData, s map[string]*schema.Schema) error {
	var cll libraries.ClusterLibraryList
	common.DataToStructPointer(d, s, &cll)
	ic.emitLibraries(cll.Libraries)
	return nil
}

func (ic *importContext) importClusterLibraries(d *schema.ResourceData, s map[string]*schema.Schema) error {
	cll, err := libraries.NewLibrariesAPI(ic.Context, ic.Client).ClusterStatus(d.Id())
	if err != nil {
		return err
	}
	for _, lib := range cll.LibraryStatuses {
		ic.emitIfDbfsFile(lib.Library.Egg)
		ic.emitIfDbfsFile(lib.Library.Jar)
		ic.emitIfDbfsFile(lib.Library.Whl)
		// Files on UC Volumes
		ic.emitIfVolumeFile(lib.Library.Whl)
		ic.emitIfVolumeFile(lib.Library.Jar)
		// Files on WSFS
		ic.emitIfWsfsFile(lib.Library.Whl)
		ic.emitIfWsfsFile(lib.Library.Jar)
	}
	return nil
}

func (ic *importContext) cacheGroups() error {
	ic.groupsMutex.Lock()
	defer ic.groupsMutex.Unlock()
	if ic.allGroups == nil {
		log.Printf("[INFO] Caching groups in memory ...")
		var groups []iam.Group
		var err error
		if ic.accountLevel {
			groups, err = ic.accountClient.Groups.ListAll(ic.Context, iam.ListAccountGroupsRequest{
				Attributes: "id",
			})
		} else {
			groups, err = ic.workspaceClient.Groups.ListAll(ic.Context, iam.ListGroupsRequest{
				Attributes: "id",
			})
		}
		if err != nil {
			log.Printf("[ERROR] can't fetch list of groups")
			return err
		}
		api := scim.NewGroupsAPI(ic.Context, ic.Client)
		ic.allGroups = make([]scim.Group, 0, len(groups))
		for i, g := range groups {
			group, err := api.Read(g.Id, "id,displayName,active,externalId,entitlements,groups,roles,members")
			if err != nil {
				log.Printf("[ERROR] Error reading group with ID %s", g.Id)
				continue
			}
			ic.allGroups = append(ic.allGroups, group)
			if (i+1)%10 == 0 {
				log.Printf("[DEBUG] Read %d out of %d groups", i+1, len(groups))
			}
		}
		log.Printf("[INFO] Cached %d groups", len(ic.allGroups))
	}
	return nil
}

func (ic *importContext) addIgnoredResource(msg string) {
	ic.ignoredResourcesMutex.Lock()
	defer ic.ignoredResourcesMutex.Unlock()
	ic.ignoredResources[msg] = struct{}{}
}

const (
	nonExistingUserOrSp = "__USER_OR_SPN_DOES_NOT_EXIST__"
)

func (ic *importContext) getUsersMapping() {
	ic.allUsersMutex.RLocker().Lock()
	userMapping := ic.allUsersMapping
	ic.allUsersMutex.RLocker().Unlock()
	if userMapping == nil {
		ic.allUsersMutex.Lock()
		defer ic.allUsersMutex.Unlock()
		if ic.allUsersMapping != nil {
			return
		}
		ic.allUsersMapping = make(map[string]string)
		var users []iam.User
		var err error
		if ic.accountLevel {
			users, err = ic.accountClient.Users.ListAll(ic.Context, iam.ListAccountUsersRequest{
				Attributes: "id,userName",
			})
		} else {
			users, err = ic.workspaceClient.Users.ListAll(ic.Context, iam.ListUsersRequest{
				Attributes: "id,userName",
			})
		}
		if err != nil {
			log.Printf("[ERROR] can't fetch list of users")
			return
		}
		for _, user := range users {
			// log.Printf("[DEBUG] adding user %v into the map. %d out of %d", user, i+1, len(users))
			ic.allUsersMapping[user.UserName] = user.Id
		}
		log.Printf("[DEBUG] users are copied")
	}
}

func (ic *importContext) findUserByName(name string, fastCheck bool) (u scim.User, err error) {
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
	ic.allUsersMutex.RLocker().Lock()
	userId, exists := ic.allUsersMapping[name]
	ic.allUsersMutex.RLocker().Unlock()
	if !exists {
		err = fmt.Errorf("there is no user '%s'", name)
		u = scim.User{UserName: nonExistingUserOrSp}
	} else {
		if fastCheck {
			return scim.User{UserName: name}, nil
		}
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
				Attributes: "id,userName",
			})
		} else {
			sps, err = ic.workspaceClient.ServicePrincipals.ListAll(ic.Context, iam.ListServicePrincipalsRequest{
				Attributes: "id,userName",
			})
		}
		if err != nil {
			log.Printf("[ERROR] can't fetch list of service principals")
			return
		}
		for _, sp := range sps {
			ic.allSpsMapping[sp.ApplicationId] = sp.Id
		}
	}
}

func (ic *importContext) getBuiltinPolicyFamilies() map[string]compute.PolicyFamily {
	ic.builtInPoliciesMutex.Lock()
	defer ic.builtInPoliciesMutex.Unlock()
	if ic.builtInPolicies == nil {
		if !ic.accountLevel {
			log.Printf("[DEBUG] Going to initialize ic.builtInPolicies. Getting policy families...")
			families, err := ic.workspaceClient.PolicyFamilies.ListAll(ic.Context, compute.ListPolicyFamiliesRequest{})
			log.Printf("[DEBUG] Going to initialize ic.builtInPolicies. Getting policy families...")
			if err == nil {
				ic.builtInPolicies = make(map[string]compute.PolicyFamily, len(families))
				for _, f := range families {
					f2 := f
					ic.builtInPolicies[f2.PolicyFamilyId] = f2
				}
			} else {
				log.Printf("[ERROR] Can't fetch cluster policy families: %v", err)
				ic.builtInPolicies = map[string]compute.PolicyFamily{}
			}
		} else {
			log.Print("[WARN] Can't list cluster policy families on account level")
			ic.builtInPolicies = map[string]compute.PolicyFamily{}
		}
	}
	return ic.builtInPolicies
}

func (ic *importContext) findSpnByAppID(applicationID string, fastCheck bool) (u scim.User, err error) {
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
		if fastCheck {
			return scim.User{ApplicationID: applicationID}, nil
		}
		a := scim.NewServicePrincipalsAPI(ic.Context, ic.Client)
		u, err = a.Read(spId, "userName,displayName,active,externalId,entitlements,groups,roles")
		if err != nil {
			log.Printf("[WARN] error reading service principal with AppID '%s', SP ID: %s", applicationID, spId)
			u = scim.User{ApplicationID: nonExistingUserOrSp}
		}
	}
	ic.spsMutex.Lock()
	defer ic.spsMutex.Unlock()
	ic.allSps[applicationID] = u

	return
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
		[]string{}, ic.Resources[r.Resource], r, resourceBlock.Body())
}

func generateUniqueID(v string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(v)))[:10]
}

func shouldOmitMd5Field(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
	if pathString == "md5" { // `md5` is kind of computed, but not declared as it...
		return true
	}
	return defaultShouldOmitFieldFunc(ic, pathString, as, d)
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

func (ic *importContext) shouldEmitForPath(path string) bool {
	if !ic.exportDeletedUsersAssets && strings.HasPrefix(path, "/Users/") {
		return ic.IsUserOrServicePrincipalDirectory(path, "/Users", false)
	}
	return true
}

func (ic *importContext) maybeEmitWorkspaceObject(resourceType, path string) {
	if ic.shouldEmitForPath(path) {
		ic.Emit(&resource{
			Resource:    resourceType,
			ID:          path,
			Incremental: ic.incremental,
		})
	} else {
		log.Printf("[WARN] Not emitting a workspace object %s for deleted user. Path='%s'", resourceType, path)
		ic.addIgnoredResource(fmt.Sprintf("%s. path=%s", resourceType, path))
	}
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
	if parent == "" {
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

func (ic *importContext) shouldSkipWorkspaceObject(object workspace.ObjectStatus, updatedSinceMs int64) bool {
	if !(object.ObjectType == workspace.Notebook || object.ObjectType == workspace.File) ||
		strings.HasPrefix(object.Path, "/Repos") {
		// log.Printf("[DEBUG] Skipping unsupported entry %v", object)
		return true
	}
	if res := ignoreIdeFolderRegex.FindStringSubmatch(object.Path); res != nil {
		return true
	}
	modifiedAt := wsObjectGetModifiedAt(object)
	if ic.incremental && modifiedAt < updatedSinceMs {
		p := ic.oldWorkspaceObjectMapping[object.ObjectID]
		if p == "" || p == object.Path {
			log.Printf("[DEBUG] skipping '%s' that was modified at %d (last active=%d)",
				object.Path, modifiedAt, updatedSinceMs)
			return true
		}
		log.Printf("[DEBUG] Different path for object %d. Old='%s', New='%s'", object.ObjectID, p, object.Path)
	}
	if !ic.MatchesName(object.Path) {
		return true
	}
	return false
}

func emitWorkpaceObject(ic *importContext, object workspace.ObjectStatus) {
	// check the size of the default channel, and add delays if it has less than %20 capacity left.
	// In this case we won't need to have increase the size of the default channel to extended capacity.
	defChannelSize := len(ic.defaultChannel)
	if float64(defChannelSize) > float64(ic.defaultHanlerChannelSize)*0.8 {
		log.Printf("[DEBUG] waiting a bit before emitting a resource because default channel is 80%% full (%d): %v",
			defChannelSize, object)
		time.Sleep(1 * time.Second)
	}
	switch object.ObjectType {
	case workspace.Notebook:
		ic.maybeEmitWorkspaceObject("databricks_notebook", object.Path)
	case workspace.File:
		ic.maybeEmitWorkspaceObject("databricks_workspace_file", object.Path)
	default:
		log.Printf("[WARN] unknown type %s for path %s", object.ObjectType, object.Path)
	}
}

func listNotebooksAndWorkspaceFiles(ic *importContext) error {
	objectsChannel := make(chan workspace.ObjectStatus, defaultChannelSize)
	numRoutines := 2 // TODO: make configurable? together with the channel size?
	var processedObjects atomic.Uint64
	for i := 0; i < numRoutines; i++ {
		num := i
		ic.waitGroup.Add(1)
		go func() {
			log.Printf("[DEBUG] Starting channel %d for workspace objects", num)
			for object := range objectsChannel {
				processedObjects.Add(1)
				ic.waitGroup.Add(1)
				// log.Printf("[DEBUG] channel %d for workspace objects, channel size=%d got %v",
				// 	num, len(objectsChannel), object)
				emitWorkpaceObject(ic, object)
				ic.waitGroup.Done()
			}
			log.Printf("[DEBUG] channel %d for workspace objects is finished", num)
			ic.waitGroup.Done()
		}()
	}
	// There are two use cases - this function will handle listing, or it will receive listing
	updatedSinceMs := ic.getUpdatedSinceMs()
	allObjects := ic.getAllWorkspaceObjects(func(objects []workspace.ObjectStatus) {
		for _, object := range objects {
			if ic.shouldSkipWorkspaceObject(object, updatedSinceMs) {
				continue
			}
			object := object
			switch object.ObjectType {
			case workspace.Notebook, workspace.File:
				objectsChannel <- object
			default:
				log.Printf("[WARN] unknown type %s for path %s", object.ObjectType, object.Path)
			}
		}
	})
	close(objectsChannel)
	log.Printf("[DEBUG] processedObjects=%d", processedObjects.Load())
	if processedObjects.Load() == 0 { // we didn't have side effect from listing as it was already happened
		log.Printf("[DEBUG] ic.getAllWorkspaceObjects already was called before, so we need to explicitly submit all objects")
		for _, object := range allObjects {
			if ic.shouldSkipWorkspaceObject(object, updatedSinceMs) {
				continue
			}
			emitWorkpaceObject(ic, object)
		}
	}
	return nil
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

// Parallel listing implementation
type syncAnswer struct {
	MU   sync.Mutex
	data []workspace.ObjectStatus
}

func (a *syncAnswer) append(objs []workspace.ObjectStatus) {
	a.MU.Lock()
	a.data = append(a.data, objs...)
	a.MU.Unlock()
}

type directoryInfo struct {
	Path     string
	Attempts int
}

// constants related to the parallel listing
const (
	envVarListParallelism       = "EXPORTER_WS_LIST_PARALLELISM"
	envVarDirectoryChannelSize  = "EXPORTER_DIRECTORIES_CHANNEL_SIZE"
	defaultWorkersPoolSize      = 10
	defaultDirectoryChannelSize = 100000
)

func recursiveAddPathsParallel(a workspace.NotebooksAPI, directory directoryInfo, dirChannel chan directoryInfo,
	answer *syncAnswer, wg *sync.WaitGroup, shouldIncludeDir func(workspace.ObjectStatus) bool, visitor func([]workspace.ObjectStatus)) {
	defer wg.Done()
	notebookInfoList, err := a.ListInternalImpl(directory.Path)
	if err != nil {
		log.Printf("[WARN] error listing '%s': %v", directory.Path, err)
		if isRetryableError(err.Error(), directory.Attempts) {
			wg.Add(1)
			log.Printf("[INFO] attempt %d of retrying listing of '%s' after error: %v",
				directory.Attempts+1, directory.Path, err)
			time.Sleep(time.Duration(retryDelaySeconds) * time.Second)
			dirChannel <- directoryInfo{Path: directory.Path, Attempts: directory.Attempts + 1}
		}
	}

	newList := make([]workspace.ObjectStatus, 0, len(notebookInfoList))
	directories := make([]workspace.ObjectStatus, 0, len(notebookInfoList))
	for _, v := range notebookInfoList {
		if v.ObjectType == workspace.Directory {
			if shouldIncludeDir(v) {
				newList = append(newList, v)
				directories = append(directories, v)
			}
		} else {
			newList = append(newList, v)
		}
	}
	answer.append(newList)
	for _, v := range directories {
		wg.Add(1)
		log.Printf("[DEBUG] putting directory '%s' into channel. Channel size: %d", v.Path, len(dirChannel))
		dirChannel <- directoryInfo{Path: v.Path}
		time.Sleep(3 * time.Millisecond)
	}
	if visitor != nil {
		visitor(newList)
	}
}

func ListParallel(a workspace.NotebooksAPI, path string, shouldIncludeDir func(workspace.ObjectStatus) bool,
	visitor func([]workspace.ObjectStatus)) ([]workspace.ObjectStatus, error) {
	var answer syncAnswer
	wg := &sync.WaitGroup{}

	if shouldIncludeDir == nil {
		shouldIncludeDir = func(workspace.ObjectStatus) bool { return true }
	}

	numWorkers := getEnvAsInt(envVarListParallelism, defaultWorkersPoolSize)
	channelSize := getEnvAsInt(envVarDirectoryChannelSize, defaultDirectoryChannelSize)
	dirChannel := make(chan directoryInfo, channelSize)
	for i := 0; i < numWorkers; i++ {
		t := i
		go func() {
			log.Printf("[DEBUG] starting go routine %d", t)
			for directory := range dirChannel {
				log.Printf("[DEBUG] processing directory %s", directory.Path)
				recursiveAddPathsParallel(a, directory, dirChannel, &answer, wg, shouldIncludeDir, visitor)
			}
		}()

	}
	log.Print("[DEBUG] pushing initial path to channel")
	wg.Add(1)
	recursiveAddPathsParallel(a, directoryInfo{Path: path}, dirChannel, &answer, wg, shouldIncludeDir, visitor)
	log.Print("[DEBUG] starting to wait")
	wg.Wait()
	log.Print("[DEBUG] closing the directory channel")
	close(dirChannel)

	answer.MU.Lock()
	defer answer.MU.Unlock()
	return answer.data, nil
}

var (
	maxRetries        = 5
	retryDelaySeconds = 2
	retriableErrors   = []string{"context deadline exceeded", "Error handling request", "Timed out after "}
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
	// log.Printf("[DEBUG] matchingCatalogAndSchema: resource: %s, origPath=%s", res.Resource, origPath)
	res_catalog_name := res.Data.Get("catalog_name").(string)
	res_schema_name := res.Data.Get("schema_name").(string)
	// log.Printf("[DEBUG] matchingCatalogAndSchema: resource: %s, catalog='%s' schema='%s'",
	// 	res.Resource, res_catalog_name, res_schema_name)
	ra_catalog_name, cat_found := ra.Get("catalog_name")
	ra_schema_name, schema_found := ra.Get("name")
	// log.Printf("[DEBUG] matchingCatalogAndSchema: approximation: %s %s, catalog='%v' (found? %v) schema='%v' (found? %v)",
	// 	ra.Type, ra.Name, ra_catalog_name, cat_found, ra_schema_name, schema_found)
	if !cat_found || !schema_found {
		log.Printf("[WARN] Can't find attributes in approximation: %s %s, catalog='%v' (found? %v) schema='%v' (found? %v). Resource: %s, catalog='%s', schema='%s'",
			ra.Type, ra.Name, ra_catalog_name, cat_found, ra_schema_name, schema_found, res.Resource, res_catalog_name, res_schema_name)
		return true
	}

	result := ra_catalog_name.(string) == res_catalog_name && ra_schema_name.(string) == res_schema_name
	// log.Printf("[DEBUG] matchingCatalogAndSchema: result: %v approximation: catalog='%v' schema='%v', res: catalog='%s' schema='%s'",
	// 	result, ra_catalog_name, ra_schema_name, res_catalog_name, res_schema_name)
	return result
}

func isMatchingShareRecipient(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	shareName, ok := res.Data.GetOk("share")
	// principal := res.Data.Get(origPath)
	// log.Printf("[DEBUG] isMatchingShareRecipient: origPath='%s', ra.Type='%s', shareName='%v', ok? %v, principal='%v'",
	// 	origPath, ra.Type, shareName, ok, principal)

	return ok && shareName.(string) != ""
}

func isMatchignShareObject(obj string) isValidAproximationFunc {
	return func(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
		objPath := strings.Replace(origPath, ".name", ".data_object_type", 1)
		objType, ok := res.Data.GetOk(objPath)
		// name := res.Data.Get(origPath)
		// log.Printf("[DEBUG] isMatchignShareObject: %s origPath='%s', ra.Type='%s', name='%v', objPath='%s' objType='%v' ok? %v",
		// 	obj, origPath, ra.Type, name, objPath, objType, ok)

		return ok && objType.(string) == obj
	}
}

func isMatchingAllowListArtifact(ic *importContext, res *resource, ra *resourceApproximation, origPath string) bool {
	objPath := strings.Replace(origPath, ".artifact", ".match_type", 1)
	matchType, ok := res.Data.GetOk(objPath)
	artifactType := res.Data.Get("artifact_type").(string)
	// artifact := res.Data.Get(origPath)
	// log.Printf("[DEBUG] isMatchingAllowListArtifact: origPath='%s', ra.Type='%s', artifactType='%v' artifact='%v', objPath='%s' matchType='%v' ok? %v",
	// 	origPath, ra.Type, artifactType, artifact, objPath, matchType, ok)

	return ok && matchType.(string) == "PREFIX_MATCH" && (artifactType == "LIBRARY_JAR" || artifactType == "INIT_SCRIPT")
}

func generateIgnoreObjectWithoutName(resourceType string) func(ic *importContext, r *resource) bool {
	return func(ic *importContext, r *resource) bool {
		res := (r.Data != nil && r.Data.Get("name").(string) == "")
		if res {
			ic.addIgnoredResource(fmt.Sprintf("%s. ID=%s", resourceType, r.ID))
		}
		return res
	}
}

func (ic *importContext) emitUCGrantsWithOwner(id string, parentResource *resource) {
	gr := &resource{
		Resource: "databricks_grants",
		ID:       id,
	}
	if parentResource.Data != nil {
		if owner, ok := parentResource.Data.GetOk("owner"); ok {
			gr.AddExtraData("owner", owner)
		}
	}
	ic.Emit(gr)
}

func (ic *importContext) addTfVar(name, value string) {
	ic.tfvarsMutex.Lock()
	defer ic.tfvarsMutex.Unlock()
	ic.tfvars[name] = value
}
