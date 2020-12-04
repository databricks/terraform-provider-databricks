package workspace

import (
	"context"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Language is a custom type for language types in Databricks notebooks
type Language string

// ObjectType is a custom type for object types in Databricks workspaces
type ObjectType string

// ExportFormat is a custom type for formats in which you can export Databricks workspace components
type ExportFormat string

// ...
const (
	Source  ExportFormat = "SOURCE"
	HTML    ExportFormat = "HTML"
	Jupyter ExportFormat = "JUPYTER"
	DBC     ExportFormat = "DBC"

	Scala  Language = "SCALA"
	Python Language = "PYTHON"
	SQL    Language = "SQL"
	R      Language = "R"

	Notebook      ObjectType = "NOTEBOOK"
	Directory     ObjectType = "DIRECTORY"
	LibraryObject ObjectType = "LIBRARY"
)

var extMap = map[string]string{
	".scala": "SCALA",
	".py":    "PYTHON",
	".sql":   "SQL",
	".r":     "R",
}

// ObjectStatus contains information when doing a get request or list request on the workspace api
type ObjectStatus struct {
	ObjectID   int64      `json:"object_id,omitempty" tf:"computed"`
	ObjectType ObjectType `json:"object_type,omitempty" tf:"computed"`
	Path       string     `json:"path"`
	Language   Language   `json:"language,omitempty"`
}

// NotebookContent contains the base64 content of the notebook
type NotebookContent struct {
	Content string `json:"content,omitempty"`
}

// ImportRequest contains the payload to import a notebook
type ImportRequest struct {
	Content   string `json:"content"`
	Path      string `json:"path"`
	Language  string `json:"language,omitempty"`
	Format    string `json:"format,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty"`
}

// NotebookDeleteRequest contains the payload to delete a notebook
type NotebookDeleteRequest struct {
	Path      string `json:"path,omitempty"`
	Recursive bool   `json:"recursive,omitempty"`
}

// NewNotebooksAPI creates NotebooksAPI instance from provider meta
func NewNotebooksAPI(ctx context.Context, m interface{}) NotebooksAPI {
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
func (a NotebooksAPI) Create(r ImportRequest) error {
	mtx.Lock()
	defer mtx.Unlock()
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
	Format ExportFormat `url:"format,omitempty"`
	Path   string       `url:"path,omitempty"`
}

// Export returns the notebook content as a base64 string
func (a NotebooksAPI) Export(path string, format ExportFormat) (string, error) {
	var notebookContent NotebookContent
	err := a.client.Get(a.context, "/workspace/export", workspacePathRequest{
		Format: format,
		Path:   path,
	}, &notebookContent)
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
func (a NotebooksAPI) List(path string, recursive bool) ([]ObjectStatus, error) {
	if recursive {
		var paths []ObjectStatus
		err := a.recursiveAddPaths(path, &paths)
		if err != nil {
			return nil, err
		}
		return paths, err
	}
	return a.list(path)
}

func (a NotebooksAPI) recursiveAddPaths(path string, pathList *[]ObjectStatus) error {
	notebookInfoList, err := a.list(path)
	if err != nil {
		return err
	}
	for _, v := range notebookInfoList {
		if v.ObjectType == Notebook {
			*pathList = append(*pathList, v)
		} else if v.ObjectType == Directory {
			err := a.recursiveAddPaths(v.Path, pathList)
			if err != nil {
				return err
			}
		}
	}
	return err
}

type objectList struct {
	Objects []ObjectStatus `json:"objects,omitempty"`
}

func (a NotebooksAPI) list(path string) ([]ObjectStatus, error) {
	var notebookList objectList
	err := a.client.Get(a.context, "/workspace/list", map[string]string{
		"path": path,
	}, &notebookList)
	return notebookList.Objects, err
}

// Delete will delete folders given a path and recursive flag
func (a NotebooksAPI) Delete(path string, recursive bool) error {
	return a.client.Post(a.context, "/workspace/delete", NotebookDeleteRequest{
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
				string(Scala),
				string(Python),
				string(R),
				string(SQL),
			}, false),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				source := d.Get("source").(string)
				if source == "" {
					return false
				}
				return old == extMap[strings.ToLower(filepath.Ext(source))]
			},
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"object_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"object_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	})
	s["content_base64"].RequiredWith = []string{"language"}
	return util.CommonResource{
		Schema:        s,
		SchemaVersion: 2,
		// TODO: state migrate
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			notebooksAPI := NewNotebooksAPI(ctx, c)
			path := d.Get("path").(string)
			parent := filepath.Dir(path)
			if parent != "/" {
				err = notebooksAPI.Mkdirs(parent)
				if err != nil {
					// TODO: handle RESOURCE_ALREADY_EXISTS
					return err
				}
			}
			lang := d.Get("language").(string)
			if lang == "" {
				// TODO: check what happens with empty source
				lang = extMap[strings.ToLower(filepath.Ext(d.Get("source").(string)))]
			}
			if err = notebooksAPI.Create(ImportRequest{
				Content:   base64.StdEncoding.EncodeToString(content),
				Language:  lang,
				Format:    "SOURCE",
				Overwrite: true,
				Path:      path,
			}); err != nil {
				return err
			}
			d.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			notebooksAPI := NewNotebooksAPI(ctx, c)
			objectStatus, err := notebooksAPI.Read(d.Id())
			if err != nil {
				return err
			}
			d.Set("url", fmt.Sprintf("%s#workspace%s", c.Host, d.Id()))
			return internal.StructToData(objectStatus, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			notebooksAPI := NewNotebooksAPI(ctx, c)
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			return notebooksAPI.Create(ImportRequest{
				Content:   base64.StdEncoding.EncodeToString(content),
				Language:  d.Get("language").(string),
				Format:    "SOURCE",
				Overwrite: true,
				Path:      d.Id(),
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewNotebooksAPI(ctx, c).Delete(d.Id(), true)
		},
	}.ToResource()
}
