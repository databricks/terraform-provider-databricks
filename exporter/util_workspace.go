package exporter

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"

	"golang.org/x/exp/slices"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func isSupportedWorkspaceObject(obj workspace.ObjectStatus) bool {
	switch obj.ObjectType {
	case workspace.Directory, workspace.Notebook, workspace.File:
		return true
	}
	return false
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

func isRepoPath(path string) bool {
	return strings.HasPrefix(path, "/Repos") || strings.HasPrefix(path, "/Workspace/Repos")
}

func maybeStringWorkspacePrefix(path string) string {
	if strings.HasPrefix(path, "/Workspace/") {
		return path[10:]
	}
	return path
}

func (ic *importContext) emitWorkspaceObject(objType, path string) {
	path = maybeStringWorkspacePrefix(path)
	if isRepoPath(path) {
		ic.emitRepoByPath(path)
	} else {
		ic.maybeEmitWorkspaceObject(objType, path, nil)
	}
}

func (ic *importContext) emitDirectoryOrRepo(path string) {
	ic.emitWorkspaceObject("databricks_directory", path)
}

func (ic *importContext) emitWorkspaceFileOrRepo(path string) {
	ic.emitWorkspaceObject("databricks_workspace_file", path)
}

func (ic *importContext) emitNotebookOrRepo(path string) {
	ic.emitWorkspaceObject("databricks_notebook", path)
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

func (ic *importContext) maybeEmitWorkspaceObject(resourceType, path string, obj *workspace.ObjectStatus) {
	if ic.shouldEmitForPath(path) {
		var data *schema.ResourceData
		if obj != nil {
			switch resourceType {
			case "databricks_notebook":
				data = workspace.ResourceNotebook().ToResource().TestResourceData()
			case "databricks_workspace_file":
				data = workspace.ResourceWorkspaceFile().ToResource().TestResourceData()
			case "databricks_directory":
				data = workspace.ResourceDirectory().ToResource().TestResourceData()
			}
			if data != nil {
				scm := ic.Resources[resourceType].Schema
				data.MarkNewResource()
				data.SetId(path)
				err := common.StructToData(obj, scm, data)
				if err != nil {
					log.Printf("[ERROR] can't convert %s object to data: %v. obj=%v", resourceType, err, obj)
					data = nil
				}
			}
		}
		ic.Emit(&resource{
			Resource:    resourceType,
			ID:          path,
			Data:        data,
			Incremental: ic.incremental,
		})
	} else {
		log.Printf("[WARN] Not emitting a workspace object %s for deleted user. Path='%s'", resourceType, path)
		ic.addIgnoredResource(fmt.Sprintf("%s. path=%s", resourceType, path))
	}
}

func (ic *importContext) shouldSkipWorkspaceObject(object workspace.ObjectStatus, updatedSinceMs int64) bool {
	if ic.incremental && object.ObjectType == workspace.Directory {
		return true
	}
	if !(object.ObjectType == workspace.Notebook || object.ObjectType == workspace.File) ||
		strings.HasPrefix(object.Path, "/Repos") {
		log.Printf("[DEBUG] Skipping unsupported entry %v", object)
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
		ic.maybeEmitWorkspaceObject("databricks_notebook", object.Path, &object)
	case workspace.File:
		ic.maybeEmitWorkspaceObject("databricks_workspace_file", object.Path, &object)
	case workspace.Directory:
		ic.maybeEmitWorkspaceObject("databricks_directory", object.Path, &object)
	default:
		log.Printf("[WARN] unknown type %s for path %s", object.ObjectType, object.Path)
	}
}

func listWorkspaceObjects(ic *importContext) error {
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
				emitWorkpaceObject(ic, object)
				ic.waitGroup.Done()
			}
			log.Printf("[DEBUG] channel %d for workspace objects is finished", num)
			ic.waitGroup.Done()
		}()
	}
	// There are two use cases - this function will handle listing, or it will receive listing
	updatedSinceMs := ic.getUpdatedSinceMs()
	isNotebooksListingEnabled := ic.isServiceInListing("notebooks")
	isDirectoryListingEnabled := ic.isServiceInListing("directories")
	isWsFilesListingEnabled := ic.isServiceInListing("wsfiles")
	allObjects := ic.getAllWorkspaceObjects(func(objects []workspace.ObjectStatus) {
		for _, object := range objects {
			if object.ObjectType == workspace.Directory {
				if !ic.incremental && object.Path != "/" && isDirectoryListingEnabled {
					objectsChannel <- object
				}
			} else {
				if ic.shouldSkipWorkspaceObject(object, updatedSinceMs) {
					continue
				}
				object := object
				switch object.ObjectType {
				case workspace.Notebook:
					if isNotebooksListingEnabled {
						objectsChannel <- object
					}
				case workspace.File:
					if isWsFilesListingEnabled {
						objectsChannel <- object
					}
				default:
					log.Printf("[WARN] unknown type %s for path %s", object.ObjectType, object.Path)
				}
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
			if !ic.incremental && isDirectoryListingEnabled && object.ObjectType == workspace.Directory && object.Path != "/" {
				emitWorkpaceObject(ic, object)
			} else if isNotebooksListingEnabled && object.ObjectType == workspace.Notebook {
				emitWorkpaceObject(ic, object)
			} else if isWsFilesListingEnabled && object.ObjectType == workspace.File {
				emitWorkpaceObject(ic, object)
			}
		}
	}
	return nil
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

func (ic *importContext) emitWorkspaceObjectParentDirectory(r *resource) {
	if !ic.isServiceEnabled("directories") {
		return
	}
	if idx := strings.LastIndex(r.ID, "/"); idx > 0 { // not found, or directly in the root...
		directoryPath := r.ID[:idx]
		ic.Emit(&resource{
			Resource: "databricks_directory",
			ID:       directoryPath,
		})
		r.AddExtraData(ParentDirectoryExtraKey, directoryPath)
	}
}
