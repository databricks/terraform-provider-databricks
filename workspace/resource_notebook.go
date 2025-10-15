package workspace

import (
	"context"
	"encoding/base64"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ...
const (
	Notebook  string = "NOTEBOOK"
	File      string = "FILE"
	Directory string = "DIRECTORY"
	Repo      string = "REPO"
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
	GitInfo               *workspace.RepoInfo    `json:"git_info,omitempty"`
	ResourceId            string                 `json:"resource_id,omitempty"`
}

type ImportResponse struct {
	ObjectID int64 `json:"object_id,omitempty"`
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
	return NotebooksAPI{m.(*common.DatabricksClient), ctx}
}

// NotebooksAPI exposes the Notebooks API
type NotebooksAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Mutex for synchronous deletes (api has poor limits in terms of allowed parallelism this increases stability of the deletes)
// sometimes there will be two folders with the same name at the same level due to issues with creating directories in
// parallel. This mutex just synchronizes everything to create folders one at a time. This mutex will be removed when mkdirs
// is removed from the notebooks resource. Then we will switch to TF resource retry.
var mtx = &sync.Mutex{}

// Create creates a notebook given the content and path
func (a NotebooksAPI) Create(r ImportPath) (ImportResponse, error) {
	if r.Format == "DBC" {
		mtx.Lock()
		defer mtx.Unlock()
	}
	var response ImportResponse
	err := a.client.Post(a.context, "/workspace/import", r, &response)
	return response, err
}

// Read returns the notebook metadata and not the contents
func (a NotebooksAPI) Read(path string) (ObjectStatus, error) {
	return a.GetStatus(path, false)
}

// Read returns the notebook metadata and not the contents
func (a NotebooksAPI) GetStatus(path string, returnGitInfo bool) (ObjectStatus, error) {
	var notebookInfo ObjectStatus
	params := map[string]string{
		"path": path,
	}
	if returnGitInfo {
		params["return_git_info"] = "true"
	}
	_, err := common.RetryOnTimeout(a.context, func(ctx context.Context) (*ObjectStatus, error) {
		err := a.client.Get(a.context, "/workspace/get-status", params, &notebookInfo)
		return nil, err
	})
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
	return a.ListInternalImpl(path)
}

func (a NotebooksAPI) recursiveAddPaths(path string, pathList *[]ObjectStatus, ignoreErrors bool) error {
	notebookInfoList, err := a.ListInternalImpl(path)
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

func (a NotebooksAPI) ListInternalImpl(path string) ([]ObjectStatus, error) {
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

func SetWorkspaceObjectComputedProperties(d *schema.ResourceData, c *common.DatabricksClient) {
	d.Set("url", c.FormatURL("#workspace", d.Id()))
	d.Set("workspace_path", "/Workspace"+d.Id())
}

// ResourceNotebook manages notebooks
func ResourceNotebook() common.Resource {
	s := FileContentSchema(map[string]*schema.Schema{
		"language": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true, // we need it because it will be filled by the provider or backend
			ValidateFunc: validation.StringInSlice([]string{
				Scala,
				Python,
				R,
				SQL,
			}, false),
		},
		"format": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"workspace_path": {
			Type:     schema.TypeString,
			Computed: true,
		},
	})
	s["content_base64"].RequiredWith = []string{"language"}
	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		CanSkipReadAfterCreateAndUpdate: func(d *schema.ResourceData) bool {
			return d.Get("format").(string) == "SOURCE"
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
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
			if createNotebook.Format == "" && createNotebook.Language != "" {
				createNotebook.Format = "SOURCE"
				d.Set("format", createNotebook.Format)
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
				d.Set("language", createNotebook.Language)
			}
			resp, err := notebooksAPI.Create(createNotebook)
			if err != nil {
				if isParentDoesntExistError(err) {
					parent := filepath.ToSlash(filepath.Dir(path))
					log.Printf("[DEBUG] Parent folder '%s' doesn't exist, creating...", parent)
					err = notebooksAPI.Mkdirs(parent)
					if err != nil {
						return err
					}
					resp, err = notebooksAPI.Create(createNotebook)
				}
				if err != nil {
					return err
				}
			}
			if d.Get("object_type").(string) == "" {
				d.Set("object_type", Notebook)
			}
			d.Set("object_id", resp.ObjectID)
			d.SetId(path)
			SetWorkspaceObjectComputedProperties(d, c)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			oldFormat := d.Get("format").(string)
			if oldFormat == "" {
				source := d.Get("source").(string)
				// check if `source` is set, and if it is, use file extension to determine format
				if source != "" {
					ext := strings.ToLower(filepath.Ext(source))
					oldFormat = extMap[ext].Format
				} else {
					oldFormat = "SOURCE"
				}
			}
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			objectStatus, err := common.RetryOnTimeout(ctx, func(ctx context.Context) (*workspace.ObjectInfo, error) {
				return w.Workspace.GetStatusByPath(ctx, d.Id())
			})
			if err != nil {
				return err
			}
			SetWorkspaceObjectComputedProperties(d, c)
			err = common.StructToData(objectStatus, s, d)
			if err != nil {
				return err
			}
			return d.Set("format", oldFormat)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			notebooksAPI := NewNotebooksAPI(ctx, c)
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			format := d.Get("format").(string)
			var resp ImportResponse
			if format == "DBC" {
				// Overwrite cannot be used for source format when importing a folder
				err = notebooksAPI.Delete(d.Id(), true)
				if err != nil {
					return err
				}
				resp, err = notebooksAPI.Create(ImportPath{
					Content: base64.StdEncoding.EncodeToString(content),
					Format:  format,
					Path:    d.Id(),
				})
			} else {
				resp, err = notebooksAPI.Create(ImportPath{
					Content:   base64.StdEncoding.EncodeToString(content),
					Language:  d.Get("language").(string),
					Format:    format,
					Overwrite: true,
					Path:      d.Id(),
				})
			}
			if err != nil {
				return err
			}
			d.Set("object_id", resp.ObjectID)
			SetWorkspaceObjectComputedProperties(d, c)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			objType := d.Get("object_type")
			return NewNotebooksAPI(ctx, c).Delete(d.Id(), !(objType == Notebook || objType == File))
		},
	}
}

func isParentDoesntExistError(err error) bool {
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "parent folder ") && strings.Contains(errStr, " does not exist")
}
