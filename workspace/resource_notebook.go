package workspace

import (
	"context"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/pkg/errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Language is a custom type for language types in Databricks notebooks
type Language string

// ObjectType is a custom type for object types in Databricks workspaces
type ObjectType string

// ExportFormat is a custom type for formats in which you can export Databricks workspace components
type ExportFormat string

// Different types of export formats available on Databricks
const (
	Source  ExportFormat = "SOURCE"
	HTML    ExportFormat = "HTML"
	Jupyter ExportFormat = "JUPYTER"
	DBC     ExportFormat = "DBC"
)

// Different types of language formats available on Databricks
const (
	Scala  Language = "SCALA"
	Python Language = "PYTHON"
	SQL    Language = "SQL"
	R      Language = "R"
)

// Different types of export formats available on Databricks
const (
	Notebook      ObjectType = "NOTEBOOK"
	Directory     ObjectType = "DIRECTORY"
	LibraryObject ObjectType = "LIBRARY"
)

// ObjectStatus contains information when doing a get request or list request on the workspace api
type ObjectStatus struct {
	ObjectID   int64      `json:"object_id" tf:"computed"`
	ObjectType ObjectType `json:"object_type" tf:"computed"`
	Path       string     `json:"path"`
	Language   Language   `json:"language,omitempty"`
}

// NotebookContent contains the base64 content of the notebook
type NotebookContent struct {
	Content string `json:"content,omitempty"`
}

// ImportRequest contains the payload to import a notebook
type ImportRequest struct {
	Content   string       `json:"content"`
	Path      string       `json:"path"`
	Language  string     `json:"language,omitempty"`
	Format    string `json:"format,omitempty"`
	Overwrite bool         `json:"overwrite,omitempty"`
}

// NotebookDeleteRequest contains the payload to delete a notebook
type NotebookDeleteRequest struct {
	Path      string `json:"path,omitempty"`
	Recursive bool   `json:"recursive,omitempty"`
}

// NewNotebooksAPI creates NotebooksAPI instance from provider meta
func NewNotebooksAPI(ctx context.Context, m interface{}) NotebooksAPI {
	return NotebooksAPI{client: m.(*common.DatabricksClient)}
}

// NotebooksAPI exposes the Notebooks API
type NotebooksAPI struct {
	client *common.DatabricksClient
}

// Mutex for synchronous deletes (api has poor limits in terms of allowed parallelism this increases stability of the deletes)
// sometimes there will be two folders with the same name at the same level due to issues with creating directories in
// parallel. This mutex just synchronizes everything to create folders one at a time. This mutex will be removed when mkdirs
// is removed from the notebooks resource. Then we will switch to TF resource retry.
var mkdirMtx = &sync.Mutex{}

// Create creates a notebook given the content and path
func (a NotebooksAPI) Create(r ImportRequest) error {
	mkdirMtx.Lock()
	defer mkdirMtx.Unlock()
	return a.client.Post("/workspace/import", r, nil)
}

// Read returns the notebook metadata and not the contents
func (a NotebooksAPI) Read(path string) (ObjectStatus, error) {
	var notebookInfo ObjectStatus
	err := a.client.Get("/workspace/get-status", map[string]string{
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
	err := a.client.Get("/workspace/export", workspacePathRequest{
		Format: format,
		Path:   path,
	}, &notebookContent)
	return notebookContent.Content, err
}

// Mkdirs will make folders in a workspace recursively given a path
func (a NotebooksAPI) Mkdirs(path string) error {
	// This mutex will be removed when mkdirs is removed from the notebooks resource.
	// Then we will switch to TF resource retry.
	mkdirMtx.Lock()
	defer mkdirMtx.Unlock()

	return a.client.Post("/workspace/mkdirs", map[string]string{
		"path": path,
	}, nil)
}

// List will list all objects in a path on the workspace and with the recursive flag it will recursively list
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
	Objects []ObjectStatus `json:"objects,omitempty" url:"objects,omitempty"`
}

func (a NotebooksAPI) list(path string) ([]ObjectStatus, error) {
	var notebookList objectList
	err := a.client.Get("/workspace/list", map[string]string{
		"path": path,
	}, &notebookList)
	return notebookList.Objects, err
}

// Delete will delete folders given a path and recursive flag
func (a NotebooksAPI) Delete(path string, recursive bool) error {
	return a.client.Post("/workspace/delete", NotebookDeleteRequest{
		Path:      path,
		Recursive: recursive,
	}, nil)
}

// ResourceNotebook manages notebooks
func ResourceNotebook() *schema.Resource {
	s := internal.StructToSchema(ImportRequest{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		return internal.StructToSchema(ObjectStatus{}, func(s2 map[string]*schema.Schema) map[string]*schema.Schema {
			s["content"].StateFunc = func(i interface{}) string {
				base64String := i.(string)
				base64, err := convertBase64ToCheckSum(base64String)
				if err != nil {
					return ""
				}
				return base64
			}
			s["content"].ValidateFunc = validation.StringIsBase64
			s["path"].ValidateFunc = ValidateNotebookPath
			s["language"].ValidateFunc = validation.StringInSlice([]string{
				string(Scala),
				string(Python),
				string(R),
				string(SQL),
			}, false)
			delete(s, "overwrite")
			delete(s, "mkdirs")
			// TODO: deprecate for v0.3
			//s["overwrite"].Default = true
			// TODO: deprecate for v0.3
			//s["mkdirs"].Default = true
			// TODO: deprecate FORMAT for v0.3
			for k,v := range s {
				if v.Computed {
					v.ForceNew = true
				}
				s2[k] = v
			}
			return s2
		})
	})
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		notebooksAPI := NewNotebooksAPI(ctx, m)
		content, err := notebooksAPI.Export(d.Id(), Source)
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		d.Partial(true)
		crc, err := convertBase64ToCheckSum(content)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("content", crc)
		objectStatus, err := notebooksAPI.Read(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		err = internal.StructToData(objectStatus, s, d)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Partial(false)
		return nil
	}
	return &schema.Resource{
		Schema: s,
		SchemaVersion: 2,
		// TODO: state migrate
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var notebookImport ImportRequest
			err := internal.DataToStructPointer(d, s, &notebookImport)
			if err != nil {
				return diag.FromErr(err)
			}
			notebookImport.Format = "SOURCE"
			notebookImport.Overwrite = true
			notebooksAPI := NewNotebooksAPI(ctx, m)
			parent := filepath.Dir(notebookImport.Path)
			if parent != "" {
				err = notebooksAPI.Mkdirs(parent)
				if err != nil {
					// TODO: handle RESOURCE_ALREADY_EXISTS
					return diag.FromErr(err)
				}
			}
			err = notebooksAPI.Create(notebookImport)
			if err != nil {
				return diag.FromErr(err)
			}
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			err := NewNotebooksAPI(ctx, m).Delete(d.Id(), true)
			if err == nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func convertBase64ToCheckSum(b64 string) (string, error) {
	dataArr, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "error", errors.Wrap(err, "error while trying to decode base64 content")
	}
	// TODO: change it to something else
	return strconv.Itoa(int(crc32.ChecksumIEEE(dataArr))), nil
}

// ValidateNotebookPath ...
func ValidateNotebookPath(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	switch {
	case v == "":
		errs = append(errs, fmt.Errorf("%s is empty must have a value", key))
		fallthrough
	case !strings.HasPrefix(v, "/"):
		errs = append(errs, fmt.Errorf("%s must start with /, got: %s", key, v))
	}

	return nil, errs
}
