package workspace

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ...
const (
	Notebook  string = "NOTEBOOK"
	File      string = "FILE"
	Directory string = "DIRECTORY"
	Scala     string = "SCALA"
	Python    string = "PYTHON"
	SQL       string = "SQL"
	R         string = "R"
	Jupyter   string = "JUPYTER"
	Auto      string = "AUTO"
)

type notebookLanguageFormat struct {
	Language  string
	Format    string
	Overwrite bool
}

var extMap = map[string]notebookLanguageFormat{
	".scala": {"SCALA", "SOURCE", true},
	".py":    {"PYTHON", "SOURCE", true},
	".sql":   {"SQL", "SOURCE", true},
	".r":     {"R", "SOURCE", true},
	".ipynb": {"", "JUPYTER", true},
	".dbc":   {"", "DBC", false},
}

type ModifiedAtInteractive struct {
	UserID     string `json:"user_id,omitempty"`
	TimeMillis int64  `json:"time_millis,omitempty"`
}

// ObjectStatus contains information when doing a get request or list request on the workspace api
type ObjectStatus struct {
	ObjectID              int64                  `json:"object_id,omitempty" tf:"computed"`
	ObjectType            string                 `json:"object_type,omitempty" tf:"computed"`
	Path                  string                 `json:"path"`
	Language              string                 `json:"language,omitempty"`
	CreatedAt             int64                  `json:"created_at,omitempty"`
	ModifiedAt            int64                  `json:"modified_at,omitempty"`
	ModifiedAtInteractive *ModifiedAtInteractive `json:"modified_at_interactive,omitempty"`
	Size                  int64                  `json:"size,omitempty"`
}

// ExportPath contains the base64 content of the notebook
type ExportPath struct {
	Content string `json:"content,omitempty"`
}

// ImportPath contains the payload to import a notebook
type ImportPath struct {
	Content   string `json:"content"`
	Path      string `json:"path"`
	Language  string `json:"language,omitempty"`
	Format    string `json:"format,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty"`
}

// DeletePath contains the payload to delete a notebook
type DeletePath struct {
	Path      string `json:"path,omitempty"`
	Recursive bool   `json:"recursive,omitempty"`
}

// NewNotebooksAPI creates NotebooksAPI instance from provider meta
func NewNotebooksAPI(ctx context.Context, m any) NotebooksAPI {
	return NotebooksAPI{m.(common.DatabricksAPI), ctx}
}

// NotebooksAPI exposes the Notebooks API
type NotebooksAPI struct {
	client  common.DatabricksAPI
	context context.Context
}

// Mutex for synchronous deletes (api has poor limits in terms of allowed parallelism this increases stability of the deletes)
// sometimes there will be two folders with the same name at the same level due to issues with creating directories in
// parallel. This mutex just synchronizes everything to create folders one at a time. This mutex will be removed when mkdirs
// is removed from the notebooks resource. Then we will switch to TF resource retry.
var mtx = &sync.Mutex{}

// Create creates a notebook given the content and path
func (a NotebooksAPI) Create(r ImportPath) error {
	if r.Format == "DBC" {
		mtx.Lock()
		defer mtx.Unlock()
	}
	return a.client.Post(a.context, "/workspace/import", r, nil)
}

// Read returns the notebook metadata and not the contents
func (a NotebooksAPI) Read(path string) (ObjectStatus, error) {
	var notebookInfo ObjectStatus
	err := a.client.Get(a.context, "/workspace/get-status", map[string]string{
		"path": path,
	}, &notebookInfo)
	return notebookInfo, err
}

type workspacePathRequest struct {
	Format string `url:"format,omitempty"`
	Path   string `url:"path,omitempty"`
}

// Export returns the notebook content as a base64 string
func (a NotebooksAPI) Export(path string, format string) (string, error) {
	var notebookContent ExportPath
	err := a.client.Get(a.context, "/workspace/export", workspacePathRequest{
		Format: format,
		Path:   path,
	}, &notebookContent)
	// TODO: return decoded []byte
	return notebookContent.Content, err
}

// Mkdirs will make folders in a workspace recursively given a path
func (a NotebooksAPI) Mkdirs(path string) error {
	// This mutex will be removed when mkdirs is removed from the notebooks resource.
	// Then we will switch to TF resource retry.
	mtx.Lock()
	defer mtx.Unlock()

	return a.client.Post(a.context, "/workspace/mkdirs", map[string]string{
		"path": path,
	}, nil)
}

type syncAnswer struct {
	MU   sync.Mutex
	data []ObjectStatus
}

func (a *syncAnswer) append(objs []ObjectStatus) {
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
	directoryListingMaxAttempts = 3
	envVarListParallelism       = "EXPORTER_WS_LIST_PARALLELISM"
	envVarDirectoryChannelSize  = "EXPORTER_DIRECTORIES_CHANNEL_SIZE"
	defaultWorkersPoolSize      = 10
	defaultDirectoryChannelSize = 100000
)

func getFormattedNowTime() string {
	return time.Now().Local().Format(time.RFC3339Nano)
}

func (a NotebooksAPI) recursiveAddPathsParallel(directory directoryInfo, dirChannel chan directoryInfo,
	answer *syncAnswer, wg *sync.WaitGroup, shouldIncludeDir func(ObjectStatus) bool) {
	defer wg.Done()
	notebookInfoList, err := a.list(directory.Path)
	if err != nil {
		log.Printf("[WARN] error listing '%s': %v", directory.Path, err)
		if directory.Attempts < directoryListingMaxAttempts {
			wg.Add(1)
			dirChannel <- directoryInfo{Path: directory.Path, Attempts: directory.Attempts + 1}
		}
	}

	newList := make([]ObjectStatus, 0, len(notebookInfoList))
	directories := make([]ObjectStatus, 0, len(notebookInfoList))
	for _, v := range notebookInfoList {
		if v.ObjectType == Directory {
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
		log.Printf("[DEBUG] %s: putting directory '%s' into channel. Channel size: %d",
			getFormattedNowTime(), v.Path, len(dirChannel))
		dirChannel <- directoryInfo{Path: v.Path}
		time.Sleep(15 * time.Millisecond)
	}
}

func getEnvAsInt(envName string, defaultValue int) int {
	if val, exists := os.LookupEnv(envName); exists {
		parsedVal, err := strconv.Atoi(val)
		if err == nil {
			return parsedVal
		}
	}
	return defaultValue
}

func (a NotebooksAPI) ListParallel(path string, shouldIncludeDir func(ObjectStatus) bool) ([]ObjectStatus, error) {
	var answer syncAnswer
	wg := &sync.WaitGroup{}

	if shouldIncludeDir == nil {
		shouldIncludeDir = func(ObjectStatus) bool { return true }
	}

	numWorkers := getEnvAsInt(envVarListParallelism, defaultWorkersPoolSize)
	channelSize := getEnvAsInt(envVarDirectoryChannelSize, defaultDirectoryChannelSize)
	dirChannel := make(chan directoryInfo, channelSize)
	for i := 0; i < numWorkers; i++ {
		t := i
		go func() {
			log.Printf("[DEBUG] %s: starting go routine %d", getFormattedNowTime(), t)
			for directory := range dirChannel {
				log.Printf("[DEBUG] %s: processing directory %s", getFormattedNowTime(), directory.Path)
				a.recursiveAddPathsParallel(directory, dirChannel, &answer, wg, shouldIncludeDir)
			}
		}()

	}
	log.Printf("[DEBUG] %s: pushing initial path to channel", getFormattedNowTime())
	wg.Add(1)
	a.recursiveAddPathsParallel(directoryInfo{Path: path}, dirChannel, &answer, wg, shouldIncludeDir)
	log.Printf("[DEBUG] %s: starting to wait", getFormattedNowTime())
	wg.Wait()
	log.Printf("[DEBUG] %s: closing the directory channel", getFormattedNowTime())
	close(dirChannel)

	answer.MU.Lock()
	defer answer.MU.Unlock()
	return answer.data, nil
}

// List will list all objects in a path on the workspace
// and with the recursive flag it will recursively list
// all the objects
func (a NotebooksAPI) List(path string, recursive bool, ignoreErrors bool) ([]ObjectStatus, error) {
	if recursive {
		var paths []ObjectStatus
		err := a.recursiveAddPaths(path, &paths, ignoreErrors)
		if err != nil {
			return nil, err
		}
		return paths, err
	}
	return a.list(path)
}

func (a NotebooksAPI) recursiveAddPaths(path string, pathList *[]ObjectStatus, ignoreErrors bool) error {
	notebookInfoList, err := a.list(path)
	if err != nil && !ignoreErrors {
		return err
	}
	for _, v := range notebookInfoList {
		*pathList = append(*pathList, v)
		if v.ObjectType == Directory {
			err := a.recursiveAddPaths(v.Path, pathList, ignoreErrors)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type ObjectList struct {
	Objects []ObjectStatus `json:"objects,omitempty"`
}

func (a NotebooksAPI) list(path string) ([]ObjectStatus, error) {
	var notebookList ObjectList
	err := a.client.Get(a.context, "/workspace/list", map[string]string{
		"path": path,
	}, &notebookList)
	return notebookList.Objects, err
}

// Delete will delete folders given a path and recursive flag
func (a NotebooksAPI) Delete(path string, recursive bool) error {
	if recursive {
		log.Printf("[DEBUG] Doing recursive delete of path '%s'", path)
		mtx.Lock()
		defer mtx.Unlock()
	}
	return a.client.Post(a.context, "/workspace/delete", DeletePath{
		Path:      path,
		Recursive: recursive,
	}, nil)
}

// ResourceNotebook manages notebooks
func ResourceNotebook() *schema.Resource {
	s := FileContentSchema(map[string]*schema.Schema{
		"language": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				Scala,
				Python,
				R,
				SQL,
			}, false),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				source := d.Get("source").(string)
				if source == "" {
					return false
				}
				ext := strings.ToLower(filepath.Ext(source))
				return old == extMap[ext].Language
			},
		},
		"format": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SOURCE",
			ValidateFunc: validation.StringInSlice([]string{
				"SOURCE",
				"DBC",
				"JUPYTER",
			}, false),
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"object_type": {
			Type:       schema.TypeString,
			Optional:   true,
			Computed:   true,
			Deprecated: "Always is a notebook",
		},
		"object_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	})
	s["content_base64"].RequiredWith = []string{"language"}
	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		Create: func(ctx context.Context, d *schema.ResourceData, c common.DatabricksAPI) error {
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			notebooksAPI := NewNotebooksAPI(ctx, c)
			path := d.Get("path").(string)
			createNotebook := ImportPath{
				Content:   base64.StdEncoding.EncodeToString(content),
				Language:  d.Get("language").(string),
				Format:    d.Get("format").(string),
				Path:      path,
				Overwrite: true,
			}
			if createNotebook.Language == "" {
				// TODO: check what happens with empty source
				ext := strings.ToLower(filepath.Ext(d.Get("source").(string)))
				createNotebook.Language = extMap[ext].Language
				createNotebook.Format = extMap[ext].Format
				// Overwrite cannot be used for Dbc format
				createNotebook.Overwrite = extMap[ext].Overwrite
				// by default it's SOURCE, but for DBC we have to change it
				d.Set("format", createNotebook.Format)
			}
			err = notebooksAPI.Create(createNotebook)
			if err != nil {
				if isParentDoesntExistError(err) {
					parent := filepath.ToSlash(filepath.Dir(path))
					log.Printf("[DEBUG] Parent folder '%s' doesn't exist, creating...", parent)
					err = notebooksAPI.Mkdirs(parent)
					if err != nil {
						return err
					}
					err = notebooksAPI.Create(createNotebook)
				}
				if err != nil {
					return err
				}
			}
			d.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c common.DatabricksAPI) error {
			notebooksAPI := NewNotebooksAPI(ctx, c)
			objectStatus, err := notebooksAPI.Read(d.Id())
			if err != nil {
				return err
			}
			d.Set("url", c.FormatURL("#workspace", d.Id()))
			return common.StructToData(objectStatus, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c common.DatabricksAPI) error {
			notebooksAPI := NewNotebooksAPI(ctx, c)
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			format := d.Get("format").(string)
			if format == "DBC" {
				// Overwrite cannot be used for source format when importing a folder
				err = notebooksAPI.Delete(d.Id(), true)
				if err != nil {
					return err
				}
				return notebooksAPI.Create(ImportPath{
					Content: base64.StdEncoding.EncodeToString(content),
					Format:  format,
					Path:    d.Id(),
				})
			}
			return notebooksAPI.Create(ImportPath{
				Content:   base64.StdEncoding.EncodeToString(content),
				Language:  d.Get("language").(string),
				Format:    format,
				Overwrite: true,
				Path:      d.Id(),
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c common.DatabricksAPI) error {
			objType := d.Get("object_type")
			return NewNotebooksAPI(ctx, c).Delete(d.Id(), !(objType == Notebook || objType == File))
		},
	}.ToResource()
}

func isParentDoesntExistError(err error) bool {
	errStr := err.Error()
	return strings.HasPrefix(errStr, "The parent folder ") && strings.HasSuffix(errStr, " does not exist.")
}
