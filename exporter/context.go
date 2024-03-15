package exporter

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"golang.org/x/exp/maps"

	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/workspace"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
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
	includeUserDomains       bool
	importAllUsers           bool
	exportDeletedUsersAssets bool
	incremental              bool
	mounts                   bool
	noFormat                 bool
	nativeImportSupported    bool
	services                 map[string]struct{}
	listing                  map[string]struct{}
	match                    string
	lastActiveDays           int64
	lastActiveMs             int64
	generateDeclaration      bool
	exportSecrets            bool
	meAdmin                  bool
	meUserName               string
	prefix                   string
	accountLevel             bool
	shImports                map[string]bool
	notebooksFormat          string
	updatedSinceStr          string
	updatedSinceMs           int64

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
	"databricks_sql_dashboard":     3,
	"databricks_sql_widget":        4,
	"databricks_sql_visualization": 4,
	"databricks_sql_query":         5,
	"databricks_sql_alert":         2,
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
	p := provider.DatabricksProvider()
	p.TerraformVersion = "exporter"
	p.SetMeta(c)
	ctx := context.WithValue(context.Background(), common.Provider, p)
	ctx = context.WithValue(ctx, common.ResourceName, "exporter")
	c.WithCommandExecutor(func(
		ctx context.Context,
		c *common.DatabricksClient) common.CommandExecutor {
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
		err = os.MkdirAll(ic.Directory, 0755)
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
		for _, g := range me.Groups {
			if g.Display == "admins" {
				ic.meAdmin = true
				ic.meUserName = me.UserName
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
	for rnLoop, irLoop := range ic.Importables {
		resourceName := rnLoop
		ir := irLoop
		if ir.List == nil {
			continue
		}
		_, exists := ic.listing[ir.Service]
		if !exists {
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
	sh, err := os.OpenFile(shFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
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
			dcfile.WriteString(fmt.Sprintf(`	host       = "%s"
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

func isSupportedWsObject(obj workspace.ObjectStatus) bool {
	switch obj.ObjectType {
	case workspace.Directory, workspace.Notebook, workspace.File:
		return true
	}
	return false
}

func (ic *importContext) generateResourceIdForWsObject(obj workspace.ObjectStatus) (string, string) {
	var rtype string
	switch obj.ObjectType {
	case workspace.Directory:
		rtype = "databricks_directory"
	case workspace.File:
		rtype = "databricks_workspace_file"
	case workspace.Notebook:
		rtype = "databricks_notebook"
	default:
		log.Printf("[WARN] Unsupported WS object type: %s in obj %v", obj.ObjectType, obj)
		return "", ""
	}
	rData := ic.Resources[rtype].Data(
		&terraform.InstanceState{
			ID:         obj.Path,
			Attributes: map[string]string{},
		})
	rData.Set("object_id", obj.ObjectID)
	rData.Set("path", obj.Path)
	name := ic.ResourceName(&resource{
		ID:       obj.Path,
		Resource: rtype,
		Data:     rData,
	})
	return generateResourceName(rtype, name), rtype
}

func (ic *importContext) loadOldWorkspaceObjects(fileName string) {
	ic.oldWorkspaceObjects = []workspace.ObjectStatus{}
	// Read a list of resources from previous run
	oldDataFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("[WARN] Can't open the file (%s) with previous list of workspace objects: %s", fileName, err.Error())
		return
	}
	err = json.Unmarshal(oldDataFile, &ic.oldWorkspaceObjects)
	if err != nil {
		log.Printf("[WARN] Can't desereialize previous list of workspace objects: %s", err.Error())
		return
	}
	log.Printf("[DEBUG] Read previous list of workspace objects. got %d objects", len(ic.oldWorkspaceObjects))
	for _, obj := range ic.oldWorkspaceObjects {
		ic.oldWorkspaceObjectMapping[obj.ObjectID] = obj.Path
	}
}

func (ic *importContext) findDeletedResources() {
	log.Print("[INFO] Starting detection of deleted workspace objects")
	if !ic.incremental || len(ic.allWorkspaceObjects) == 0 {
		return
	}
	if len(ic.oldWorkspaceObjects) == 0 {
		log.Print("[INFO] Previous list of workspace objects is empty")
		return
	}
	// generate IDs of current objects
	currentObjs := map[string]struct{}{}
	for _, obj := range ic.allWorkspaceObjects {
		obj := obj
		if !isSupportedWsObject(obj) {
			continue
		}
		rid, _ := ic.generateResourceIdForWsObject(obj)
		currentObjs[rid] = struct{}{}
	}
	// Loop through previous objects, and if it's missing from the current list, add it to deleted, including permission
	for _, obj := range ic.oldWorkspaceObjects {
		obj := obj
		if !isSupportedWsObject(obj) {
			continue
		}
		rid, rtype := ic.generateResourceIdForWsObject(obj)
		_, exists := currentObjs[rid]
		if exists {
			log.Printf("[DEBUG] object %s still exists", rid) // change to TRACE?
			continue
		}
		log.Printf("[DEBUG] object %s is deleted!", rid)
		ic.deletedResources[rid] = struct{}{}
		// convert into permissions. This is quite fragile right now, need to think how to handle it better
		var permId string
		switch rtype {
		case "databricks_notebook":
			permId = "databricks_permissions.notebook_" + rid[len(rtype)+1:]
		case "databricks_directory":
			permId = "databricks_permissions.directory_" + rid[len(rtype)+1:]
		case "databricks_workspace_file":
			permId = "databricks_permissions.ws_file_" + rid[len(rtype)+1:]
		}
		log.Printf("[DEBUG] deleted permissions object %s", permId)
		if permId != "" {
			ic.deletedResources[permId] = struct{}{}
		}
	}
	log.Printf("[INFO] Finished detection of deleted workspace objects. Detected %d deleted objects.",
		len(ic.deletedResources))
	log.Printf("[DEBUG] Deleted objects. %v", ic.deletedResources) // change to TRACE?
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

func generateResourceName(rtype, rname string) string {
	return rtype + "." + rname
}

func generateBlockFullName(block *hclwrite.Block) string {
	labels := block.Labels()
	return generateResourceName(labels[0], strings.Join(labels[1:], "_"))
}

type resourceWriteData struct {
	BlockName     string
	ResourceBody  string
	ImportCommand string
}

type dataWriteChannel chan *resourceWriteData
type importWriteChannel chan string

func (ic *importContext) handleResourceWrite(generatedFile string, ch dataWriteChannel, importChan importWriteChannel) {
	var existingFile *hclwrite.File
	if ic.incremental {
		log.Printf("[DEBUG] Going to read existing file %s", generatedFile)
		content, err := os.ReadFile(generatedFile)
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("[WARN] File %s doesn't exist when using incremental export", generatedFile)
		} else if err != nil {
			log.Printf("[ERROR] error opening %s", generatedFile)
		} else {
			log.Printf("[DEBUG] Going to parse existing file %s", generatedFile)
			var diags hcl.Diagnostics
			existingFile, diags = hclwrite.ParseConfig(content, generatedFile, hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				log.Printf("[ERROR] parsing of existing file %s failed: %s", generatedFile, diags.Error())
			} else {
				log.Printf("[DEBUG] There are %d objects in existing file %s",
					len(existingFile.Body().Blocks()), generatedFile)
			}
		}
	}
	if existingFile == nil {
		existingFile = hclwrite.NewEmptyFile()
	}

	tf, err := os.Create(generatedFile)
	if err != nil {
		log.Printf("[ERROR] Can't create %s: %v", generatedFile, err)
		return
	}

	//
	newResources := make(map[string]struct{}, 100)
	log.Printf("[DEBUG] started processing new writes for %s", generatedFile)
	for f := range ch {
		if f != nil {
			log.Printf("[DEBUG] started writing resource body for %s", f.BlockName)
			_, err = tf.WriteString(f.ResourceBody)
			if err == nil {
				newResources[f.BlockName] = struct{}{}
				if f.ImportCommand != "" {
					ic.waitGroup.Add(1)
					importChan <- f.ImportCommand
				}
				log.Printf("[DEBUG] finished writing resource body for %s", f.BlockName)
			} else {
				log.Printf("[ERROR] Error when writing to %s: %v", generatedFile, err)
			}
		} else {
			log.Print("[WARN] got nil as resourceWriteData!")
		}
		ic.waitGroup.Done()
	}
	numResources := len(newResources)
	log.Printf("[DEBUG] finished processing new writes for %s. Wrote %d resources", generatedFile, numResources)
	// update existing file if incremental mode
	if ic.incremental {
		log.Printf("[DEBUG] Starting to merge existing resources for %s", generatedFile)
		f := hclwrite.NewEmptyFile()
		for _, block := range existingFile.Body().Blocks() {
			blockName := generateBlockFullName(block)
			_, exists := newResources[blockName]
			_, deleted := ic.deletedResources[blockName]
			if exists {
				log.Printf("[DEBUG] resource %s already generated, skipping...", blockName)
			} else if deleted {
				log.Printf("[DEBUG] resource %s is deleted, skipping...", blockName)
			} else {
				log.Printf("[DEBUG] resource %s doesn't exist, adding...", blockName)
				f.Body().AppendBlock(block)
				numResources = numResources + 1
			}
		}
		_, err = tf.WriteString(string(f.Bytes()))
		if err != nil {
			log.Printf("[ERROR] error when writing existing resources for file %s: %v", generatedFile, err)
		}
		log.Printf("[DEBUG] Finished merging existing resources for %s", generatedFile)
	}
	tf.Close()
	if numResources == 0 {
		log.Printf("[DEBUG] removing empty file %s - no resources for a given service", generatedFile)
		os.Remove(generatedFile)
	}
}

func (ic *importContext) writeShellImports(sh *os.File, importChan importWriteChannel) {
	for importCommand := range importChan {
		if importCommand != "" && sh != nil {
			log.Printf("[DEBUG] writing import command %s", importCommand)
			sh.WriteString(importCommand + "\n")
			delete(ic.shImports, importCommand)
		} else {
			log.Print("[WARN] got empty import command... or file is nil")
		}
		ic.waitGroup.Done()
	}
	if sh != nil {
		log.Printf("[DEBUG] Writing the rest of import commands. len=%d", len(ic.shImports))
		for k := range ic.shImports {
			parts := strings.Split(k, " ")
			if len(parts) > 3 {
				resource := parts[2]
				_, deleted := ic.deletedResources[resource]
				if deleted {
					log.Printf("[DEBUG] Resource %s is deleted. Skipping import command for it", resource)
					continue
				}
			}
			sh.WriteString(k + "\n")
		}
	}
}

func extractResourceIdFromImportBlock(block *hclwrite.Block) string {
	if block.Type() != "import" {
		log.Print("[WARN] it's not an import block!")
		return ""
	}
	idAttr := block.Body().GetAttribute("to")
	if idAttr == nil {
		log.Printf("[WARN] Can't find `to` attribute in the import block")
		return ""
	}
	idVal := string(idAttr.Expr().BuildTokens(nil).Bytes())
	return strings.TrimSpace(idVal)
}

func extractResourceIdFromImportBlockString(importBlock string) string {
	block, diags := hclwrite.ParseConfig([]byte(importBlock), "test.tf", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		log.Printf("[WARN] parsing of import block %s has failed: %s", importBlock, diags.Error())
		return ""
	}
	if len(block.Body().Blocks()) == 0 {
		log.Printf("[WARN] import block %s has 0 blocks!", importBlock)
		return ""
	}
	return extractResourceIdFromImportBlock(block.Body().Blocks()[0])
}

func (ic *importContext) writeNativeImports(importChan importWriteChannel) {
	if !ic.nativeImportSupported {
		log.Print("[DEBUG] Native import is not enabled, skipping...")
		return
	}
	importsFileName := fmt.Sprintf("%s/import.tf", ic.Directory)
	// TODO: in incremental mode read existing file with imports and append them for not processed & not deleted resources
	var existingFile *hclwrite.File
	if ic.incremental {
		log.Printf("[DEBUG] Going to read existing file %s", importsFileName)
		content, err := os.ReadFile(importsFileName)
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("[WARN] File %s doesn't exist when using incremental export", importsFileName)
		} else if err != nil {
			log.Printf("[ERROR] error opening %s", importsFileName)
		} else {
			log.Printf("[DEBUG] Going to parse existing file %s", importsFileName)
			var diags hcl.Diagnostics
			existingFile, diags = hclwrite.ParseConfig(content, importsFileName, hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				log.Printf("[ERROR] parsing of existing file %s failed: %s", importsFileName, diags.Error())
			} else {
				log.Printf("[DEBUG] There are %d objects in existing file %s",
					len(existingFile.Body().Blocks()), importsFileName)
			}
		}
	}
	if existingFile == nil {
		existingFile = hclwrite.NewEmptyFile()
	}

	// do actual writes
	importsFile, err := os.Create(importsFileName)
	if err != nil {
		log.Printf("[ERROR] Can't create %s: %v", importsFileName, err)
		return
	}
	defer importsFile.Close()

	newImports := make(map[string]struct{}, 100)
	log.Printf("[DEBUG] started processing new writes for %s", importsFileName)
	// write native imports
	for importBlock := range importChan {
		if importBlock != "" {
			log.Printf("[TRACE] writing import command %s", importBlock)
			importsFile.WriteString(importBlock)
			id := extractResourceIdFromImportBlockString(importBlock)
			if id != "" {
				newImports[id] = struct{}{}
			}
		} else {
			log.Print("[WARN] got empty import command...")
		}
		ic.waitGroup.Done()
	}
	// write the rest of import blocks
	numResources := len(newImports)
	log.Printf("[DEBUG] finished processing new writes for %s. Wrote %d resources", importsFileName, numResources)
	// update existing file if incremental mode
	if ic.incremental {
		log.Printf("[DEBUG] Starting to merge existing resources for %s", importsFileName)
		f := hclwrite.NewEmptyFile()
		for _, block := range existingFile.Body().Blocks() {
			blockName := extractResourceIdFromImportBlock(block)
			if blockName == "" {
				log.Printf("[WARN] can't extract resource ID from import block: %s",
					string(block.BuildTokens(nil).Bytes()))
				continue
			}
			_, exists := newImports[blockName]
			_, deleted := ic.deletedResources[blockName]
			if exists {
				log.Printf("[DEBUG] resource %s already generated, skipping...", blockName)
			} else if deleted {
				log.Printf("[DEBUG] resource %s is deleted, skipping...", blockName)
			} else {
				log.Printf("[DEBUG] resource %s doesn't exist, adding...", blockName)
				f.Body().AppendBlock(block)
				numResources = numResources + 1
			}
		}
		_, err = importsFile.WriteString(string(f.Bytes()))
		if err != nil {
			log.Printf("[ERROR] error when writing existing resources for file %s: %v", importsFileName, err)
		}
		log.Printf("[DEBUG] Finished merging existing resources for %s", importsFileName)
	}

}

func (ic *importContext) processSingleResource(resourcesChan resourceChannel,
	writerChannels map[string]dataWriteChannel, nativeImportChannel importWriteChannel) {
	processed := 0
	generated := 0
	ignored := 0
	for r := range resourcesChan {
		processed = processed + 1
		if r == nil {
			log.Print("[WARN] Got nil resource...")
			ic.waitGroup.Done()
			continue
		}
		ir := ic.Importables[r.Resource]
		if ir.Ignore != nil && ir.Ignore(ic, r) {
			log.Printf("[WARN] Ignoring resource %s: %s", r.Resource, r.Name)
			ignored = ignored + 1
			ic.waitGroup.Done()
			continue
		}
		var err error
		f := hclwrite.NewEmptyFile()
		log.Printf("[TRACE] Generating %s: %s", r.Resource, r.Name)
		body := f.Body()
		if ir.Body != nil {
			err = ir.Body(ic, body, r)
			if err != nil {
				log.Printf("[ERROR] error calling ir.Body for %v: %s", r, err.Error())
			}
		} else {
			resourceBlock := body.AppendNewBlock("resource", []string{r.Resource, r.Name})
			err = ic.dataToHcl(ir, []string{}, ic.Resources[r.Resource], r, resourceBlock.Body())
			if err != nil {
				log.Printf("[ERROR] error generating body for %v: %s", r, err.Error())
			}
		}
		if err == nil && len(body.Blocks()) > 0 {
			formatted := hclwrite.Format(f.Bytes())
			// fix some formatting in a hacky way instead of writing 100 lines of HCL AST writer code
			formatted = []byte(ic.regexFix(string(formatted), ic.hclFixes))
			writeData := &resourceWriteData{
				ResourceBody: string(formatted),
				BlockName:    generateBlockFullName(body.Blocks()[0]),
			}
			if r.Mode != "data" && ic.Resources[r.Resource].Importer != nil {
				writeData.ImportCommand = r.ImportCommand(ic)
				if ic.nativeImportSupported { // generate import block for native import
					imp := hclwrite.NewEmptyFile()
					imoBlock := imp.Body().AppendNewBlock("import", []string{})
					imoBlock.Body().SetAttributeValue("id", cty.StringVal(r.ID))
					traversal := hcl.Traversal{
						hcl.TraverseRoot{Name: r.Resource},
						hcl.TraverseAttr{Name: r.Name},
					}
					tokens := hclwrite.TokensForTraversal(traversal)
					imoBlock.Body().SetAttributeRaw("to", tokens)
					formattedImp := hclwrite.Format(imp.Bytes())
					//log.Printf("[DEBUG] Import block for %s: %s", r.ID, string(formattedImp))
					ic.waitGroup.Add(1)
					nativeImportChannel <- string(formattedImp)
				}
			}
			ch, exists := writerChannels[ir.Service]
			if exists {
				ic.waitGroup.Add(1)
				ch <- writeData
			} else {
				log.Printf("[WARN] can't find a channel for service: %s, resource: %s", ir.Service, r.Resource)
			}
			log.Printf("[TRACE] Finished generating %s: %s", r.Resource, r.Name)
			generated = generated + 1
		} else {
			log.Printf("[WARN] error generating resource body: %v, or body blocks len is 0", err)
		}
		ic.waitGroup.Done()
	}
	log.Printf("[DEBUG] processed resources: %d, generated: %d, ignored: %d", processed, generated, ignored)
}

func (ic *importContext) generateAndWriteResources(sh *os.File) {
	resources := ic.Scope.Sorted()
	scopeSize := ic.Scope.Len()
	t1 := time.Now()
	log.Printf("[INFO] Generating configuration for %d resources", scopeSize)

	// make configurable via environment variables
	resourceHandlersNumber := getEnvAsInt("EXPORTER_RESOURCE_GENERATORS", 50)
	resourcesChan := make(resourceChannel, defaultChannelSize)

	resourceWriters := make(map[string]dataWriteChannel, len(ic.Resources))
	for service := range ic.services {
		resourceWriters[service] = make(dataWriteChannel, defaultChannelSize)
	}
	writersWaitGroup := &sync.WaitGroup{}
	// write shell script for importing
	shellImportChan := make(importWriteChannel, defaultChannelSize)
	writersWaitGroup.Add(1)
	go func() {
		ic.writeShellImports(sh, shellImportChan)
		writersWaitGroup.Done()
	}()
	//
	nativeImportChan := make(importWriteChannel, defaultChannelSize)
	writersWaitGroup.Add(1)
	go func() {
		ic.writeNativeImports(nativeImportChan)
		writersWaitGroup.Done()
	}()
	// start resource handlers
	for i := 0; i < resourceHandlersNumber; i++ {
		i := i
		go func() {
			log.Printf("[DEBUG] Starting resource handler %d", i)
			ic.processSingleResource(resourcesChan, resourceWriters, nativeImportChan)
		}()
	}
	// start writers for specific services
	for service, ch := range resourceWriters {
		service := service
		ch := ch
		generatedFile := fmt.Sprintf("%s/%s.tf", ic.Directory, service)
		log.Printf("[DEBUG] starting writer for service %s", service)
		writersWaitGroup.Add(1)
		go func() {
			ic.handleResourceWrite(generatedFile, ch, shellImportChan)
			writersWaitGroup.Done()
		}()
	}
	// submit all extracted resources...
	for i, r := range resources {
		ic.waitGroup.Add(1)
		resourcesChan <- r
		if i%500 == 0 {
			log.Printf("[INFO] Submitted %d of %d resources", i+1, scopeSize)
		}
	}
	ic.waitGroup.Wait()
	// close all channels
	close(shellImportChan)
	close(nativeImportChan)
	close(resourcesChan)
	for service, ch := range resourceWriters {
		log.Printf("Closing writer for service %s", service)
		close(ch)
	}
	writersWaitGroup.Wait()

	log.Printf("[INFO] Finished generation of configuration for %d resources (took %v seconds)",
		scopeSize, time.Since(t1).Seconds())
}

func (ic *importContext) generateTfvars() error {
	if len(ic.tfvars) == 0 {
		return nil
	}
	f := hclwrite.NewEmptyFile()
	body := f.Body()
	fileName := fmt.Sprintf("%s/terraform.tfvars", ic.Directory)

	vf, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer vf.Close()

	for k, v := range ic.tfvars {
		body.SetAttributeValue(k, cty.StringVal(v))
	}
	// nolint
	vf.Write(f.Bytes())
	log.Printf("[INFO] Written %d tfvars", len(ic.tfvars))

	return nil
}

func (ic *importContext) generateVariables() error {
	if len(ic.variables) == 0 {
		return nil
	}
	f := hclwrite.NewEmptyFile()
	body := f.Body()
	fileName := fmt.Sprintf("%s/vars.tf", ic.Directory)
	if ic.incremental {
		content, err := os.ReadFile(fileName)
		if err == nil {
			ftmp, diags := hclwrite.ParseConfig(content, fileName, hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				log.Printf("[ERROR] parsing of existing file failed: %s", diags)
			} else {
				tbody := ftmp.Body()
				for _, block := range tbody.Blocks() {
					typ := block.Type()
					labels := block.Labels()
					log.Printf("[DEBUG] blockBody: %v %v\n", typ, labels)
					_, present := ic.variables[labels[0]]
					if typ == "variable" && present {
						log.Printf("[DEBUG] Ignoring variable '%s' that will be re-exported", labels[0])
					} else {
						log.Printf("[DEBUG] Adding not exported object. type='%s', labels=%v", typ, labels)
						body.AppendBlock(block)
					}
				}
			}
		} else {
			log.Printf("[ERROR] opening file %s", fileName)
		}
	}
	vf, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer vf.Close()

	for k, v := range ic.variables {
		b := body.AppendNewBlock("variable", []string{k}).Body()
		b.SetAttributeValue("description", cty.StringVal(v))
	}
	// nolint
	vf.Write(f.Bytes())
	log.Printf("[INFO] Written %d variables", len(ic.variables))
	return nil
}

func (ic *importContext) MatchesName(n string) bool {
	if ic.match == "" {
		return true
	}
	return strings.Contains(strings.ToLower(n), strings.ToLower(ic.match))
}

func genTraversalTokens(sr *resourceApproximation, pick string) hcl.Traversal {
	if sr.Mode == "data" {
		return hcl.Traversal{
			hcl.TraverseRoot{Name: "data"},
			hcl.TraverseAttr{Name: sr.Type},
			hcl.TraverseAttr{Name: sr.Name},
			hcl.TraverseAttr{Name: pick},
		}
	}
	return hcl.Traversal{
		hcl.TraverseRoot{Name: sr.Type},
		hcl.TraverseAttr{Name: sr.Name},
		hcl.TraverseAttr{Name: pick},
	}
}

func (ic *importContext) Find(value, attr string, ref reference, origResource *resource, origPath string) (string, hcl.Traversal, bool) {
	log.Printf("[DEBUG] Starting searching for reference for resource %s, attr='%s', value='%s', ref=%v",
		ref.Resource, attr, value, ref)
	// optimize performance by avoiding doing regexp matching multiple times
	matchValue := ""
	switch ref.MatchType {
	case MatchRegexp:
		if ref.Regexp == nil {
			log.Printf("[WARN] you must provide regular expression for 'regexp' match type")
			return "", nil, false
		}
		res := ref.Regexp.FindStringSubmatch(value)
		if len(res) < 2 {
			log.Printf("[WARN] no match for regexp: %v in string %s", ref.Regexp, value)
			return "", nil, false
		}
		matchValue = res[1]
	case MatchCaseInsensitive:
		matchValue = strings.ToLower(value) // performance optimization to avoid doing it in the loop
	case MatchExact, MatchDefault:
		matchValue = value
	case MatchPrefix, MatchLongestPrefix:
		if ref.MatchValueTransformFunc != nil {
			matchValue = ref.MatchValueTransformFunc(value)
		} else {
			matchValue = value
		}
	}
	// doing explicit lookup in the state.  For case insensitive matches, first attempt to lookup for the value,
	// and do iteration if it's not found
	if (ref.MatchType == MatchExact || ref.MatchType == MatchDefault || ref.MatchType == MatchRegexp ||
		ref.MatchType == MatchCaseInsensitive) && !ref.SkipDirectLookup {
		sr := ic.State.Get(ref.Resource, attr, matchValue)
		if sr != nil && (ref.IsValidApproximation == nil || ref.IsValidApproximation(ic, origResource, sr, origPath)) {
			log.Printf("[DEBUG] Finished direct lookup for reference for resource %s, attr='%s', value='%s', ref=%v. Found: type=%s name=%s",
				ref.Resource, attr, value, ref, sr.Type, sr.Name)
			// TODO: we need to not generate traversals resources for which their Ignore function returns true...
			return matchValue, genTraversalTokens(sr, attr), sr.Mode == "data"
		}
		if ref.MatchType != MatchCaseInsensitive { // for case-insensitive matching we'll try iteration
			log.Printf("[DEBUG] Finished direct lookup for reference for resource %s, attr='%s', value='%s', ref=%v. Not found",
				ref.Resource, attr, value, ref)
			return "", nil, false
		}
	}

	maxPrefixLen := 0
	maxPrefixOrigValue := ""
	var maxPrefixResource *resourceApproximation
	srs := *ic.State.Resources(ref.Resource)
	for _, sr := range srs {
		for _, i := range sr.Instances {
			v := i.Attributes[attr]
			if v == nil {
				log.Printf("[WARN] Can't find instance attribute '%v' in resource: '%v'", attr, ref.Resource)
				continue
			}
			strValue := v.(string)
			origValue := strValue
			if ref.SearchValueTransformFunc != nil {
				strValue = ref.SearchValueTransformFunc(strValue)
				log.Printf("[DEBUG] Resource %s. Transformed value from '%s' to '%s'", ref.Resource, origValue, strValue)
			}
			matched := false
			switch ref.MatchType {
			case MatchCaseInsensitive:
				matched = (strings.ToLower(strValue) == matchValue)
			case MatchPrefix:
				matched = strings.HasPrefix(matchValue, strValue)
			case MatchLongestPrefix:
				if strings.HasPrefix(matchValue, strValue) && len(origValue) > maxPrefixLen {
					maxPrefixLen = len(origValue)
					maxPrefixOrigValue = origValue
					maxPrefixResource = sr
				}
			case MatchExact, MatchDefault:
				matched = (strValue == matchValue)
			default:
				log.Printf("[WARN] Unsupported match type: %s", ref.MatchType)
			}
			if !matched || (ref.IsValidApproximation != nil && !ref.IsValidApproximation(ic, origResource, sr, origPath)) {
				continue
			}
			// TODO: we need to not generate traversals resources for which their Ignore function returns true...
			log.Printf("[DEBUG] Finished searching for reference for resource %s, attr='%s', value='%s', ref=%v. Found: type=%s name=%s",
				ref.Resource, attr, value, ref, sr.Type, sr.Name)
			return origValue, genTraversalTokens(sr, attr), sr.Mode == "data"
		}
	}
	if ref.MatchType == MatchLongestPrefix && maxPrefixResource != nil &&
		(ref.IsValidApproximation == nil || ref.IsValidApproximation(ic, origResource, maxPrefixResource, origPath)) {
		log.Printf("[DEBUG] Finished searching longest prefix for reference for resource %s, attr='%s', value='%s', ref=%v. Found: type=%s name=%s",
			ref.Resource, attr, value, ref, maxPrefixResource.Type, maxPrefixResource.Name)
		return maxPrefixOrigValue, genTraversalTokens(maxPrefixResource, attr), maxPrefixResource.Mode == "data"
	}
	log.Printf("[DEBUG] Finished searching for reference for resource %s, pick=%s, ref=%v. Not found", ref.Resource, attr, ref)
	return "", nil, false
}

// This function checks if resource exist in any state (already added or in process of addition)
func (ic *importContext) Has(r *resource) bool {
	return ic.HasInState(r, false)
}

func (ic *importContext) isImporting(s string) (bool, bool) {
	ic.importingMutex.RLocker().Lock()
	defer ic.importingMutex.RLocker().Unlock()
	v, visiting := ic.importing[s]
	return v, visiting
}

// This function checks if resource exist. onlyAdded flag enforces that true is returned only if it was added with Add()
func (ic *importContext) HasInState(r *resource, onlyAdded bool) bool {
	v, visiting := ic.isImporting(r.String())
	if visiting && (v || !onlyAdded) {
		return true
	}
	return ic.State.Has(r)
}

func (ic *importContext) setImportingState(s string, state bool) {
	ic.importingMutex.Lock()
	defer ic.importingMutex.Unlock()
	ic.importing[s] = state
}

func (ic *importContext) Add(r *resource) {
	if ic.HasInState(r, true) { // resource must exist and already marked as added
		return
	}
	ic.setImportingState(r.String(), true) // mark resource as added
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
		Module:    ic.Module,
		Type:      r.Resource,
		Name:      r.Name,
		Instances: []instanceApproximation{inst},
	})
	// in single-threaded scenario scope is toposorted
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

func (ic *importContext) isServiceEnabled(service string) bool {
	_, exists := ic.services[service]
	return exists
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
	// TODO: change into channels, if stack trace depth issues would surface
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
	if ic.Has(r) {
		log.Printf("[DEBUG] %s already imported", r)
		return
	}
	if ic.testEmits != nil {
		log.Printf("[INFO] %s is emitted in test mode", r)
		ic.testEmitsMutex.Lock()
		ic.testEmits[r.String()] = true
		ic.testEmitsMutex.Unlock()
		return
	}
	ic.setImportingState(r.String(), false) // we're starting to add a new resource
	_, ok = ic.Resources[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available in provider", r)
		return
	}

	if ic.accountLevel && !ir.AccountLevel {
		log.Printf("[DEBUG] %s (%s service) is not part of the account level export", r.Resource, ir.Service)
		return
	}
	// TODO: add similar condition for checking workspace-level objects only. After new ACLs import is merged

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

func maybeAddQuoteCharacter(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}

func (ic *importContext) getTraversalTokens(ref reference, value string, origResource *resource, origPath string) (hclwrite.Tokens, bool) {
	matchType := ref.MatchTypeValue()
	attr := ref.MatchAttribute()
	attrValue, traversal, isData := ic.Find(value, attr, ref, origResource, origPath)
	// at least one invocation of ic.Find will assign Nil to traversal if resource with value is not found
	if traversal == nil {
		return nil, isData
	}
	// capture if it's data?
	switch matchType {
	case MatchExact, MatchDefault, MatchCaseInsensitive:
		return hclwrite.TokensForTraversal(traversal), isData
	case MatchPrefix, MatchLongestPrefix:
		rest := value[len(attrValue):]
		tokens := hclwrite.Tokens{&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"', '$', '{'}}}
		tokens = append(tokens, hclwrite.TokensForTraversal(traversal)...)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'}'}})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(maybeAddQuoteCharacter(rest))})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}})
		return tokens, isData
	case MatchRegexp:
		indices := ref.Regexp.FindStringSubmatchIndex(value)
		if len(indices) == 4 {
			tokens := hclwrite.Tokens{&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}}}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(maybeAddQuoteCharacter(value[0:indices[2]]))})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'$', '{'}})
			tokens = append(tokens, hclwrite.TokensForTraversal(traversal)...)
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'}'}})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(maybeAddQuoteCharacter(value[indices[3]:]))})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}})
			return tokens, isData
		}
		log.Printf("[WARN] Can't match found data in '%s'. Indices: %v", value, indices)
	default:
		log.Printf("[WARN] Unsupported match type: %s", ref.MatchType)
	}
	return nil, false
}

// TODO: move to IC
var dependsRe = regexp.MustCompile(`(\.[\d]+)`)

func (ic *importContext) generateVariableName(attrName, name string) string {
	return fmt.Sprintf("%s_%s", attrName, name)
}

func (ic *importContext) reference(i importable, path []string, value string, ctyValue cty.Value, origResource *resource) hclwrite.Tokens {
	pathString := strings.Join(path, ".")
	match := dependsRe.ReplaceAllString(pathString, "")
	// get reference candidate, but if it's a `data`, then look for another non-data reference if possible..
	var dataTokens hclwrite.Tokens
	for _, d := range i.Depends {
		if d.Path != match {
			continue
		}
		if d.File {
			relativeFile := fmt.Sprintf("${path.module}/%s", value)
			return hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}},
				&hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(relativeFile)},
				&hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}},
			}
		}
		if d.Variable {
			varName := ic.generateVariableName(path[0], value)
			return ic.variable(varName, "")
		}

		tokens, isData := ic.getTraversalTokens(d, value, origResource, pathString)
		if tokens != nil {
			if isData {
				dataTokens = tokens
				log.Printf("[DEBUG] Got reference to data for dependency %v", d)
			} else {
				return tokens
			}
		}
	}
	if len(dataTokens) > 0 {
		return dataTokens
	}
	return hclwrite.TokensForValue(ctyValue)
}

func (ic *importContext) variable(name, desc string) hclwrite.Tokens {
	ic.variables[name] = desc
	return hclwrite.TokensForTraversal(hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: name},
	})
}

type fieldTuple struct {
	Field  string
	Schema *schema.Schema
}

func (ic *importContext) dataToHcl(i importable, path []string,
	pr *schema.Resource, res *resource, body *hclwrite.Body) error {
	d := res.Data
	ss := []fieldTuple{}
	for a, as := range pr.Schema {
		ss = append(ss, fieldTuple{a, as})
	}
	sort.Slice(ss, func(i, j int) bool {
		// it just happens that reverse field order
		// makes the most beautiful configs
		return ss[i].Field > ss[j].Field
	})
	var_cnt := 0
	for _, tuple := range ss {
		a, as := tuple.Field, tuple.Schema
		pathString := strings.Join(append(path, a), ".")
		raw, nonZero := d.GetOk(pathString)
		// log.Printf("[DEBUG] path=%s, raw='%v'", pathString, raw)
		if i.ShouldOmitField == nil { // we don't have custom function, so skip computed & default fields
			if defaultShouldOmitFieldFunc(ic, pathString, as, d) {
				continue
			}
		} else if i.ShouldOmitField(ic, pathString, as, d) {
			continue
		}
		mpath := dependsRe.ReplaceAllString(pathString, "")
		for _, ref := range i.Depends {
			if ref.Path == mpath && ref.Variable {
				// sensitive fields are moved to variable depends, variable name is normalized
				// TODO: handle a case when we have multiple blocks, so names won't be unique
				raw = ic.regexFix(ic.ResourceName(res), simpleNameFixes)
				if var_cnt > 0 {
					raw = fmt.Sprintf("%s_%d", raw, var_cnt)
				}
				nonZero = true
				var_cnt++
			}
		}
		shouldSkip := !nonZero
		if as.Required { // for required fields we must produce a value, even empty...
			shouldSkip = false
		} else if as.Default != nil && !reflect.DeepEqual(raw, as.Default) {
			// In case when have zero value, but there is non-zero default, we also need to produce it
			shouldSkip = false
		}
		if shouldSkip {
			continue
		}
		switch as.Type {
		case schema.TypeString:
			value := raw.(string)
			tokens := ic.reference(i, append(path, a), value, cty.StringVal(value), res)
			body.SetAttributeRaw(a, tokens)
		case schema.TypeBool:
			body.SetAttributeValue(a, cty.BoolVal(raw.(bool)))
		case schema.TypeInt:
			var num int64
			switch iv := raw.(type) {
			case int:
				num = int64(iv)
			case int32:
				num = int64(iv)
			case int64:
				num = iv
			}
			body.SetAttributeRaw(a, ic.reference(i, append(path, a),
				strconv.FormatInt(num, 10), cty.NumberIntVal(num), res))
		case schema.TypeFloat:
			body.SetAttributeValue(a, cty.NumberFloatVal(raw.(float64)))
		case schema.TypeMap:
			// TODO: Resolve references in maps as well, and also support different types inside map...
			ov := map[string]cty.Value{}
			for key, iv := range raw.(map[string]any) {
				v := cty.StringVal(fmt.Sprintf("%v", iv))
				ov[key] = v
			}
			body.SetAttributeValue(a, cty.ObjectVal(ov))
		case schema.TypeSet:
			if rawSet, ok := raw.(*schema.Set); ok {
				rawList := rawSet.List()
				err := ic.readListFromData(i, append(path, a), res, rawList, body, as, func(i int) string {
					return strconv.Itoa(rawSet.F(rawList[i]))
				})
				if err != nil {
					return err
				}
			}
		case schema.TypeList:
			if rawList, ok := raw.([]any); ok {
				err := ic.readListFromData(i, append(path, a), res, rawList, body, as, strconv.Itoa)
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unsupported schema type: %v", path)
		}
	}
	return nil
}

func (ic *importContext) readListFromData(i importable, path []string, res *resource,
	rawList []any, body *hclwrite.Body, as *schema.Schema, offsetConverter func(i int) string) error {
	if len(rawList) == 0 {
		return nil
	}
	name := path[len(path)-1]
	switch elem := as.Elem.(type) {
	case *schema.Resource:
		if as.MaxItems == 1 {
			nestedPath := append(path, offsetConverter(0))
			confBlock := body.AppendNewBlock(name, []string{})
			return ic.dataToHcl(i, nestedPath, elem, res, confBlock.Body())
		}
		for offset := range rawList {
			confBlock := body.AppendNewBlock(name, []string{})
			nestedPath := append(path, offsetConverter(offset))
			err := ic.dataToHcl(i, nestedPath, elem, res, confBlock.Body())
			if err != nil {
				return err
			}
		}
	case *schema.Schema:
		toks := hclwrite.Tokens{}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenOBrack,
			Bytes: []byte{'['},
		})
		for _, raw := range rawList {
			if len(toks) != 1 {
				toks = append(toks, &hclwrite.Token{
					Type:  hclsyntax.TokenComma,
					Bytes: []byte{','},
				})
			}
			switch x := raw.(type) {
			case string:
				value := raw.(string)
				toks = append(toks, ic.reference(i, path, value, cty.StringVal(value), res)...)
			case int:
				// probably we don't even use integer lists?...
				toks = append(toks, hclwrite.TokensForValue(
					cty.NumberIntVal(int64(x)))...)
			default:
				return fmt.Errorf("unsupported primitive list: %#v", path)
			}
		}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenCBrack,
			Bytes: []byte{']'},
		})
		body.SetAttributeRaw(name, toks)
	}
	return nil
}
