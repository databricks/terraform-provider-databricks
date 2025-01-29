package exporter

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"golang.org/x/exp/maps"

	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/workspace"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/** High level overview of importer design:

                                    +----------+  Add  +--------------------+
                                    | resource +-------> stateApproximation |
                                    +--^-------+       +----|-------------|-+
                                       |                    |             |
  +------------------------+           | Emit               |             |
  | "normal provider flow" |    +------^-----+        +-----V-----+   +---V---+
  +------------^-----------+    | importable +--------> reference |   | scope |
               |                +------^-----+        +--------|--+   +---V---+
+--------------+--------------+        |                       |          |
|terraform-provider-databricks|        |                       |          |
+--------------+--------------+        |                       |          |
               |                       | List               +--V----------V---+
    +----------v---------+        +----^------------+       |                 |
    |                    |        |                 |       |    Generated    |
    |  importer command  +-------->  importContext  |       |    HCL Files    |
    |                    |        |                 |       |                 |
    +--------------------+        +-----------------+       +-----------------+
*/

type resourceChannel chan *resource

type gitInfoCacheEntry struct {
	IsPresent bool
	RepoId    int64
}

type importContext struct {
	// not modified/used only in single thread
	Module            string
	Context           context.Context
	Client            *common.DatabricksClient
	Importables       map[string]importable
	Resources         map[string]*schema.Resource
	Directory         string
	nameFixes         []regexFix
	hclFixes          []regexFix
	variables         map[string]string
	variablesLock     sync.Mutex
	workspaceConfKeys map[string]any

	workspaceClient *databricks.WorkspaceClient
	accountClient   *databricks.AccountClient

	channels                 map[string]resourceChannel
	defaultChannel           resourceChannel
	defaultHanlerChannelSize int

	// mutable resources
	State *stateApproximation
	Scope importedResources

	// command-line resources (immutable, or set by the single thread)
	includeUserDomains                      bool
	importAllUsers                          bool
	exportDeletedUsersAssets                bool
	incremental                             bool
	mounts                                  bool
	noFormat                                bool
	nativeImportSupported                   bool
	services                                map[string]struct{}
	listing                                 map[string]struct{}
	match                                   string
	matchRegexStr                           string
	matchRegex                              *regexp.Regexp
	excludeRegexStr                         string
	excludeRegex                            *regexp.Regexp
	filterDirectoriesDuringWorkspaceWalking bool
	lastActiveDays                          int64
	lastActiveMs                            int64
	generateDeclaration                     bool
	exportSecrets                           bool
	meAdmin                                 bool
	meUserName                              string
	prefix                                  string
	accountLevel                            bool
	shImports                               map[string]bool
	notebooksFormat                         string
	updatedSinceStr                         string
	updatedSinceMs                          int64

	waitGroup *sync.WaitGroup

	// TODO: protect by mutex?
	mountMap map[string]mount

	testEmits      map[string]bool
	testEmitsMutex sync.Mutex

	allGroups   []scim.Group
	groupsMutex sync.Mutex

	allUsers        map[string]scim.User
	usersMutex      sync.RWMutex
	allUsersMapping map[string]string // maps user_name -> internal ID
	allUsersMutex   sync.RWMutex

	allSps        map[string]scim.User
	allSpsMapping map[string]string // maps application_id -> internal ID
	spsMutex      sync.RWMutex

	importing      map[string]bool
	importingMutex sync.RWMutex

	sqlDatasources      map[string]string
	sqlDatasourcesMutex sync.Mutex

	// workspace-related objects & corresponding mutex
	allDirectories            []workspace.ObjectStatus
	allWorkspaceObjects       []workspace.ObjectStatus
	wsObjectsMutex            sync.RWMutex
	oldWorkspaceObjects       []workspace.ObjectStatus
	oldWorkspaceObjectMapping map[int64]string

	gitInfoCache      map[string]gitInfoCacheEntry
	gitInfoCacheMutex sync.RWMutex

	builtInPolicies      map[string]compute.PolicyFamily
	builtInPoliciesMutex sync.Mutex

	// Workspace-level UC Metastore information
	currentMetastore *catalog.GetMetastoreSummaryResponse

	// tracking ignored objects
	ignoredResourcesMutex sync.Mutex
	ignoredResources      map[string]struct{}

	deletedResources map[string]struct{}

	// emitting of users/SPs
	emittedUsers      map[string]struct{}
	emittedUsersMutex sync.RWMutex

	userOrSpDirectories      map[string]bool
	userOrSpDirectoriesMutex sync.RWMutex

	tfvarsMutex sync.Mutex
	tfvars      map[string]string
}

type mount struct {
	URL             string
	InstanceProfile string
	ClusterID       string
}

var nameFixes = []regexFix{
	{regexp.MustCompile(`[0-9a-f]{8}[_-][0-9a-f]{4}[_-][0-9a-f]{4}` +
		`[_-][0-9a-f]{4}[_-][0-9a-f]{12}[_-]`), ""},
	{regexp.MustCompile(`[-\s\.\|]`), "_"},
	{regexp.MustCompile(`\W+`), "_"},
	{regexp.MustCompile(`[_]{2,}`), "_"},
}

// less aggressive name normalizations
var simpleNameFixes = []regexFix{
	{nameNormalizationRegex, "_"},
}

var workspaceConfKeys = map[string]any{
	"enableIpAccessLists":                              false,
	"enableTokensConfig":                               false,
	"maxTokenLifetimeDays":                             0,
	"maxUserInactiveDays":                              0,
	"storeInteractiveNotebookResultsInCustomerAccount": false,
	"enableDeprecatedClusterNamedInitScripts":          false,
	"enableDeprecatedGlobalInitScripts":                false,
	"enforceWorkspaceViewAcls":                         false,
	"enforceUserIsolation":                             false,
	"enableProjectTypeInWorkspace":                     false,
	"enableWorkspaceFilesystem":                        false,
	"enableProjectsAllowList":                          false,
	"projectsAllowList":                                "",
	"reposIpynbResultsExportPermissions":               "ALLOW",
	"enable-X-Frame-Options":                           false,
	"enable-X-Content-Type-Options":                    false,
	"enable-X-XSS-Protection":                          false,
	"enableResultsDownloading":                         false,
	"enableUploadDataUis":                              false,
	"enableExportNotebook":                             false,
	"enableNotebookTableClipboard":                     false,
	"enableWebTerminal":                                false,
	"enableDbfsFileBrowser":                            false,
	"enableDatabricksAutologgingAdminConf":             false,
	"mlflowRunArtifactDownloadEnabled":                 false,
	"mlflowModelServingEndpointCreationEnabled":        false,
	"mlflowModelRegistryEmailNotificationsEnabled":     false,
	"rStudioUserDefaultHomeBase":                       false,
	"enableVerboseAuditLogs":                           false,
	"enableEnforceImdsV2":                              false,
	"enableLibraryAndInitScriptOnSharedCluster":        false,
	"enablePipelinesDataSample":                        false,
	"customerApprovedWSLoginExpirationTime":            false,
	"enableLegacyNotebookVisualizations":               "indefinite",
	"enableJobViewAcls":                                false,
	"enforceClusterViewAcls":                           false,
}

const (
	defaultChannelSize = 100000
	defaultNumRoutines = 2
	envVariablePrefix  = "EXPORTER_PARALLELISM_"
)

// increased concurrency limits, could be also overridden via environment variables with name: envVariablePrefix + resource type
var goroutinesNumber = map[string]int{
	"databricks_notebook":          10,
	"databricks_directory":         5,
	"databricks_workspace_file":    5,
	"databricks_dbfs_file":         3,
	"databricks_user":              1,
	"databricks_service_principal": 1,
	"databricks_dashboard":         4,
	"databricks_sql_dashboard":     3,
	"databricks_sql_widget":        4,
	"databricks_sql_visualization": 4,
	"databricks_query":             6,
	"databricks_alert":             2,
	"databricks_permissions":       11,
}

func makeResourcesChannels() map[string]resourceChannel {
	resources := []string{"databricks_user", "databricks_service_principal", "databricks_group"}
	if val, exists := os.LookupEnv("EXPORTER_DEDICATED_RESOUSE_CHANNELS"); exists {
		resources = strings.Split(val, ",")
	}
	channels := make(map[string]resourceChannel, len(resources))
	for _, r := range resources {
		channels[r] = make(resourceChannel, defaultChannelSize)
	}
	return channels
}

func newImportContext(c *common.DatabricksClient) *importContext {
	p := sdkv2.DatabricksProvider()
	p.TerraformVersion = "exporter"
	p.SetMeta(c)
	ctx := context.WithValue(context.Background(), common.Provider, p)
	ctx = context.WithValue(ctx, common.ResourceName, "exporter")
	c.WithCommandExecutor(func(
		ctx context.Context,
		c *common.DatabricksClient,
	) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, c)
	})

	defaultHanlerChannelSize := getEnvAsInt("EXPORTER_DEFAULT_HANDLER_CHANNEL_SIZE", defaultChannelSize*3)

	supportedResources := maps.Keys(resourcesMap)
	return &importContext{
		Client:                    c,
		Context:                   ctx,
		State:                     newStateApproximation(supportedResources),
		Importables:               resourcesMap,
		Resources:                 p.ResourcesMap,
		Scope:                     importedResources{},
		importing:                 map[string]bool{},
		nameFixes:                 nameFixes,
		hclFixes:                  []regexFix{}, // Be careful with that! it may break working code
		variables:                 map[string]string{},
		allDirectories:            []workspace.ObjectStatus{},
		allWorkspaceObjects:       []workspace.ObjectStatus{},
		oldWorkspaceObjects:       []workspace.ObjectStatus{},
		oldWorkspaceObjectMapping: map[int64]string{},
		gitInfoCache:              map[string]gitInfoCacheEntry{},
		workspaceConfKeys:         workspaceConfKeys,
		shImports:                 map[string]bool{},
		notebooksFormat:           "SOURCE",
		allUsers:                  map[string]scim.User{},
		allSps:                    map[string]scim.User{},
		waitGroup:                 &sync.WaitGroup{},
		channels:                  makeResourcesChannels(),
		defaultHanlerChannelSize:  defaultHanlerChannelSize,
		defaultChannel:            make(resourceChannel, defaultHanlerChannelSize),
		ignoredResources:          map[string]struct{}{},
		deletedResources:          map[string]struct{}{},
		emittedUsers:              map[string]struct{}{},
		userOrSpDirectories:       map[string]bool{},
		services:                  map[string]struct{}{},
		listing:                   map[string]struct{}{},
		tfvars:                    map[string]string{},
	}
}

func getLastRunString(fileName string) string {
	var updatedSinceStr string
	statsData, err := os.ReadFile(fileName)
	if err == nil {
		var m map[string]any
		err = json.Unmarshal(statsData, &m)
		s, exists := m["startTime"]
		if err == nil && exists {
			updatedSinceStr = s.(string)
		} else {
			log.Printf("[WARN] Can't decode data: err=%v or startTime field doesn't exist. data: '%s'",
				err, string(statsData))
		}
	} else {
		log.Printf("[WARN] Can't load data from file %s: %v", fileName, err)
	}
	return updatedSinceStr
}

func (ic *importContext) Run() error {
	startTime := time.Now()
	statsFileName := ic.Directory + "/exporter-run-stats.json"
	wsObjectsFileName := ic.Directory + "/ws_objects.json"
	if len(ic.services) == 0 {
		return fmt.Errorf("no services to import")
	}

	if ic.matchRegexStr != "" {
		log.Printf("[DEBUG] Using regex '%s' to filter resources", ic.matchRegexStr)
		re, err := regexp.Compile(ic.matchRegexStr)
		if err != nil {
			log.Printf("[ERROR] can't compile regex '%s': %v", ic.matchRegexStr, err)
			return err
		}
		ic.matchRegex = re
	}
	if ic.excludeRegexStr != "" {
		log.Printf("[DEBUG] Using regex '%s' to filter resources", ic.excludeRegexStr)
		re, err := regexp.Compile(ic.excludeRegexStr)
		if err != nil {
			log.Printf("[ERROR] can't compile regex '%s': %v", ic.excludeRegexStr, err)
			return err
		}
		ic.excludeRegex = re
	}
	if ic.incremental {
		if ic.updatedSinceStr == "" {
			ic.updatedSinceStr = getLastRunString(statsFileName)
		}
		if ic.updatedSinceStr == "" {
			return fmt.Errorf("-updated-since is required with -interactive parameter if %s file doesn't exist",
				statsFileName)
		}
		tm, err := time.Parse(time.RFC3339, ic.updatedSinceStr)
		if err != nil {
			return fmt.Errorf("can't parse value '%s' please specify it in ISO8601 format, i.e. 2023-07-01T00:00:00Z",
				ic.updatedSinceStr)
		}
		ic.updatedSinceStr = tm.UTC().Format(time.RFC3339)
		tm, _ = time.Parse(time.RFC3339, ic.updatedSinceStr)
		ic.updatedSinceMs = tm.UnixMilli()

		ic.loadOldWorkspaceObjects(wsObjectsFileName)
	}

	log.Printf("[INFO] Importing %s module into %s directory Databricks resources of %s services. Listing %s",
		ic.Module, ic.Directory, maps.Keys(ic.services), maps.Keys(ic.listing))

	ic.notebooksFormat = strings.ToUpper(ic.notebooksFormat)
	_, supportedFormat := fileExtensionFormatMapping[ic.notebooksFormat]
	if !supportedFormat && ic.notebooksFormat != "SOURCE" {
		return fmt.Errorf("unsupported notebook format: '%s'", ic.notebooksFormat)
	}

	info, err := os.Stat(ic.Directory)
	if os.IsNotExist(err) {
		err = os.MkdirAll(ic.Directory, 0o755)
		if err != nil {
			return fmt.Errorf("can't create directory %s", ic.Directory)
		}
	} else if !info.IsDir() {
		return fmt.Errorf("the path %s is not a directory", ic.Directory)
	}

	ic.accountLevel = ic.Client.Config.IsAccountClient()
	if ic.accountLevel {
		ic.meAdmin = true
		// TODO: check if we can get the current user from the account client
		ic.accountClient, err = ic.Client.AccountClient()
		if err != nil {
			return err
		}
	} else {
		ic.workspaceClient, err = ic.Client.WorkspaceClient()
		if err != nil {
			return err
		}
		me, err := ic.workspaceClient.CurrentUser.Me(ic.Context)
		if err != nil {
			return err
		}
		ic.meUserName = me.UserName
		for _, g := range me.Groups {
			if g.Display == "admins" {
				ic.meAdmin = true
				break
			}
		}
		currentMetastore, err := ic.workspaceClient.Metastores.Summary(ic.Context)
		if err == nil {
			ic.currentMetastore = currentMetastore
		} else {
			log.Printf("[WARN] can't get current UC metastore: %v", err)
		}
	}
	// Concurrent execution part
	if ic.waitGroup == nil {
		ic.waitGroup = &sync.WaitGroup{}
	}
	// Start goroutines for each resource type
	ic.startImportChannels()

	// Start listing of objects
	listWorkspaceObjectsAlreadyRunning := false
	for rnLoop, irLoop := range ic.Importables {
		resourceName := rnLoop
		ir := irLoop
		// TODO: extend this to other services?  Like, Git Folders
		if !ic.accountLevel && (ir.Service == "notebooks" || ir.Service == "wsfiles" || (ir.Service == "directories" && !ic.incremental)) {
			if _, exists := ic.listing[ir.Service]; exists && !listWorkspaceObjectsAlreadyRunning {
				ic.waitGroup.Add(1)
				log.Printf("[DEBUG] Starting listing of workspace objects")
				go func() {
					if err := listWorkspaceObjects(ic); err != nil {
						log.Printf("[ERROR] listing of workspace objects failed %s", err)
					}
					log.Print("[DEBUG] Finished listing of workspace objects")
					ic.waitGroup.Done()
				}()
				listWorkspaceObjectsAlreadyRunning = true
			}
			continue
		}
		if ir.List == nil {
			continue
		}
		if _, exists := ic.listing[ir.Service]; !exists {
			log.Printf("[DEBUG] %s (%s service) is not part of listing", resourceName, ir.Service)
			continue
		}
		if ic.accountLevel && !ir.AccountLevel {
			log.Printf("[DEBUG] %s (%s service) is not a account level resource", resourceName, ir.Service)
			continue
		}
		if !ic.accountLevel && !ir.WorkspaceLevel {
			log.Printf("[DEBUG] %s (%s service) is not a workspace level resource", resourceName, ir.Service)
			continue
		}
		ic.waitGroup.Add(1)
		go func() {
			if err := ir.List(ic); err != nil {
				log.Printf("[ERROR] %s (%s service) listing failed: %s", resourceName, ir.Service, err)
			}
			log.Printf("[DEBUG] Finished listing for service %s", resourceName)
			ic.waitGroup.Done()
		}()
	}

	ic.waitGroup.Wait()
	// close channels
	ic.closeImportChannels()

	// Generating the code
	ic.findDeletedResources()
	if ic.Scope.Len() == 0 && len(ic.deletedResources) == 0 {
		return fmt.Errorf("no resources to import or delete")
	}
	shFileName := fmt.Sprintf("%s/import.sh", ic.Directory)
	if ic.incremental {
		shFile, err := os.Open(shFileName)
		if err == nil {
			defer shFile.Close()
			fileScanner := bufio.NewScanner(shFile)
			fileScanner.Split(bufio.ScanLines)
			for fileScanner.Scan() {
				line := fileScanner.Text()
				if strings.HasPrefix(line, "terraform import ") {
					ic.shImports[strings.TrimRight(line, "\n")] = true
				}
			}
		} else {
			log.Printf("[ERROR] opening %s: %v", shFileName, err)
		}
	}
	sh, err := os.OpenFile(shFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}
	defer sh.Close()
	// nolint
	sh.WriteString("#!/bin/sh\n\nset -e\n\n")

	if ic.generateDeclaration {
		dcfile, err := os.Create(fmt.Sprintf("%s/databricks.tf", ic.Directory))
		if err != nil {
			return err
		}
		// nolint
		dcfile.WriteString(
			`terraform {
  required_providers {
    databricks = {
      source  = "databricks/databricks"
      version = "` + common.Version() + `"
    }
  }
}

provider "databricks" {
`)
		if ic.accountLevel {
			dcfile.WriteString(fmt.Sprintf(`  host       = "%s"
  account_id = "%s"
`, ic.Client.Config.Host, ic.Client.Config.AccountID))
		}
		dcfile.WriteString(`}`)
		dcfile.Close()
	}
	//
	ic.generateAndWriteResources(sh)
	err = ic.generateVariables()
	if err != nil {
		log.Printf("[ERROR] can't write variables file: %s", err.Error())
	}

	err = ic.generateTfvars()
	if err != nil {
		log.Printf("[ERROR] can't write terraform.tfvars file: %s", err.Error())
	}

	// Write stats file
	if stats, err := os.Create(statsFileName); err == nil {
		defer stats.Close()
		statsData := map[string]any{
			"startTime":       startTime.UTC().Format(time.RFC3339),
			"duration":        fmt.Sprintf("%f sec", time.Since(startTime).Seconds()),
			"exportedObjects": ic.Scope.Len(),
		}
		statsBytes, _ := json.Marshal(statsData)
		if _, err = stats.Write(statsBytes); err != nil {
			log.Printf("[ERROR] can't write stats into the %s: %s", statsFileName, err.Error())
		}
	}

	// Write workspace objects file
	if len(ic.allWorkspaceObjects) > 0 {
		if wsObjects, err := os.Create(wsObjectsFileName); err == nil {
			defer wsObjects.Close()
			wsObjectsBytes, _ := json.Marshal(ic.allWorkspaceObjects)
			if _, err = wsObjects.Write(wsObjectsBytes); err != nil {
				log.Printf("[ERROR] can't write workspace objects into the %s: %s", wsObjectsFileName, err.Error())
			}
		} else {
			log.Printf("[ERROR] can't open %s: %s", wsObjectsFileName, err.Error())
		}
	}

	// output ignored resources...
	ignoredResourcesFileName := fmt.Sprintf("%s/ignored_resources.txt", ic.Directory)
	if ignored, err := os.Create(ignoredResourcesFileName); err == nil {
		defer ignored.Close()
		ic.ignoredResourcesMutex.Lock()
		keys := maps.Keys(ic.ignoredResources)
		sort.Strings(keys)
		for _, s := range keys {
			ignored.WriteString(s + "\n")
		}
		ic.ignoredResourcesMutex.Unlock()
	} else {
		log.Printf("[ERROR] can't open %s: %s", ignoredResourcesFileName, err.Error())
	}

	if !ic.noFormat {
		// format generated source code
		cmd := exec.CommandContext(context.Background(), "terraform", "fmt")
		cmd.Dir = ic.Directory
		err = cmd.Run()
		if err != nil {
			log.Printf("[ERROR] problems when formatting the generated code: %v", err)
			return err
		}
	}
	log.Printf("[INFO] Done. Please edit the files and roll out new environment.")
	return nil
}

func (ic *importContext) resourceHandler(num int, resourceType string, ch resourceChannel) {
	log.Printf("[DEBUG] Starting goroutine %d for resource %s", num, resourceType)
	for r := range ch {
		log.Printf("[DEBUG] channel for %s, channel size=%d got %v", resourceType, len(ch), r)
		if r != nil {
			r.ImportResource(ic)
			log.Printf("[DEBUG] Finished importing %s, %v", resourceType, r)
		}
	}
}

func (ic *importContext) startImportChannels() {
	for rt, c := range ic.channels {
		ch := c
		resourceType := rt
		numRoutines, exists := goroutinesNumber[resourceType]
		if !exists {
			numRoutines = defaultNumRoutines
		}
		numRoutines = getEnvAsInt(envVariablePrefix+resourceType, numRoutines)

		for i := 0; i < numRoutines; i++ {
			num := i
			go func() {
				ic.resourceHandler(num, resourceType, ch)
			}()
		}
	}

	numRoutines := getEnvAsInt(envVariablePrefix+"default", 15)
	for i := 0; i < numRoutines; i++ {
		num := i
		go func() {
			ic.resourceHandler(num, "default", ic.defaultChannel)
		}()
	}
}

func (ic *importContext) closeImportChannels() {
	for rt, ch := range ic.channels {
		log.Printf("[DEBUG] Closing channel for resource %s", rt)
		close(ch)
	}
	close(ic.defaultChannel)
}

func (ic *importContext) HasInState(r *resource) bool {
	return ic.State.Has(r)
}

func (ic *importContext) Add(r *resource) {
	if ic.HasInState(r) { // resource must exist in the state
		return
	}
	rString := r.String()
	ic.importingMutex.Lock()
	isAdded, ok := ic.importing[rString]
	if ok && isAdded {
		ic.importingMutex.Unlock()
		log.Printf("[DEBUG] %s is already added", rString)
		return
	}
	ic.importing[rString] = true // mark resource as added
	ic.importingMutex.Unlock()
	state := r.Data.State()
	if state == nil {
		log.Printf("[ERROR] state is nil for %s", r)
		return
	}
	inst := instanceApproximation{
		Attributes: map[string]any{},
	}
	for k, v := range state.Attributes {
		inst.Attributes[k] = v
	}
	if r.Mode == "" {
		r.Mode = "managed"
	}
	inst.Attributes["id"] = r.ID
	ic.State.Append(resourceApproximation{
		Mode:      r.Mode,
		Type:      r.Resource,
		Name:      r.Name,
		Instances: []instanceApproximation{inst},
		Resource:  r,
	})
	ic.Scope.Append(r)
}

func (ic *importContext) regexFix(s string, fixes []regexFix) string {
	for _, x := range fixes {
		s = x.Regex.ReplaceAllString(s, x.Replacement)
	}
	return s
}

func (ic *importContext) ResourceName(r *resource) string {
	name := r.Name
	if name == "" && ic.Importables[r.Resource].Name != nil {
		name = ic.Importables[r.Resource].Name(ic, r.Data)
	}
	if name == "" {
		name = r.ID
	}
	name = ic.prefix + name
	origCaseName := name
	name = strings.ToLower(name)
	name = ic.regexFix(name, ic.nameFixes)
	// this is either numeric id or all-non-ascii
	if regexp.MustCompile(`^\d`).MatchString(name) || name == "" {
		if name == "" {
			origCaseName = r.ID
		}
		name = fmt.Sprintf("r%x", md5.Sum([]byte(origCaseName)))[0:12]
	}
	return name
}

func (ic *importContext) EmitIfUpdatedAfterMillis(r *resource, modifiedAt int64, message string) {
	updatedSinceMs := ic.getUpdatedSinceMs()
	if ic.incremental && modifiedAt < updatedSinceMs {
		log.Printf("[DEBUG] skipping %s that was modified at %d (last active=%d)",
			message, modifiedAt, updatedSinceMs)
		return
	}
	ic.Emit(r)
}

func (ic *importContext) EmitIfUpdatedAfterMillisAndNameMatches(r *resource, name string, modifiedAt int64, message string) {
	if ic.MatchesName(name) {
		ic.EmitIfUpdatedAfterMillis(r, modifiedAt, message)
	}
}

func (ic *importContext) EmitIfUpdatedAfterIsoString(r *resource, updatedAt, message string) {
	updatedSinceStr := ic.getUpdatedSinceStr()
	if ic.incremental && updatedAt < updatedSinceStr {
		log.Printf("[DEBUG] skipping %s that was modified at %s (updatedSince=%s)", message,
			updatedAt, updatedSinceStr)
		return
	}
	ic.Emit(r)
}

func (ic *importContext) Emit(r *resource) {
	_, v := r.MatchPair()
	if v == "" {
		log.Printf("[DEBUG] %s has got empty identifier", r)
		return
	}
	ir, ok := ic.Importables[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available for import", r)
		return
	}
	if !ic.isServiceEnabled(ir.Service) {
		log.Printf("[DEBUG] %s (%s service) is not part of the import", r.Resource, ir.Service)
		return
	}
	rString := r.String()
	if ic.testEmits != nil {
		log.Printf("[INFO] %s is emitted in test mode", r)
		ic.testEmitsMutex.Lock()
		ic.testEmits[rString] = true
		ic.testEmitsMutex.Unlock()
		return
	}
	// we need to check that we're not importing the same resource twice - this may happen
	// under high concurrency for specific resources, for example, directories when they
	// aren't part of the listing
	ic.importingMutex.Lock()
	res, ok := ic.importing[rString]
	if ok {
		ic.importingMutex.Unlock()
		log.Printf("[DEBUG] %s already being imported: %v", rString, res)
		return
	}
	ic.importing[rString] = false // we're starting to add a new resource
	ic.importingMutex.Unlock()
	_, ok = ic.Resources[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available in provider", r)
		return
	}

	if ic.accountLevel && !ir.AccountLevel {
		log.Printf("[DEBUG] %s (%s service) is not part of the account level export",
			r.Resource, ir.Service)
		return
	}
	if !ic.accountLevel && !ir.WorkspaceLevel {
		log.Printf("[DEBUG] %s (%s service) is not part of the workspace level export",
			r.Resource, ir.Service)
		return
	}
	// from here, it should be done by the goroutine...  send resource into the channel
	ch, exists := ic.channels[r.Resource]
	if exists {
		log.Printf("[TRACE] increasing counter & sending to the channel for resource %s", r.Resource)
		ic.waitGroup.Add(1)
		ch <- r
	} else {
		log.Print("[TRACE] increasing counter & sending to the default channel")
		ic.waitGroup.Add(1)
		ic.defaultChannel <- r
	}
}
