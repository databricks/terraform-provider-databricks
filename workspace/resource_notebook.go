package workspace

import (
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/pkg/errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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

// WorkspaceObjectStatus contains information when doing a get request or list request on the workspace api
type WorkspaceObjectStatus struct {
	ObjectID   int64      `json:"object_id,omitempty"`
	ObjectType ObjectType `json:"object_type,omitempty"`
	Path       string     `json:"path,omitempty"`
	Language   Language   `json:"language,omitempty"`
}

// NotebookContent contains the base64 content of the notebook
type NotebookContent struct {
	Content string `json:"content,omitempty"`
}

// NotebookImportRequest contains the payload to import a notebook
type NotebookImportRequest struct {
	Content   string       `json:"content,omitempty"`
	Path      string       `json:"path,omitempty"`
	Language  Language     `json:"language,omitempty"`
	Overwrite bool         `json:"overwrite,omitempty"`
	Format    ExportFormat `json:"format,omitempty"`
}

// NotebookDeleteRequest contains the payload to delete a notebook
type NotebookDeleteRequest struct {
	Path      string `json:"path,omitempty"`
	Recursive bool   `json:"recursive,omitempty"`
}

// NewNotebooksAPI creates NotebooksAPI instance from provider meta
func NewNotebooksAPI(m interface{}) NotebooksAPI {
	return NotebooksAPI{C: m.(*common.DatabricksClient)}
}

// NotebooksAPI exposes the Notebooks API
type NotebooksAPI struct {
	C *common.DatabricksClient
}

// Mutex for synchronous deletes (api has poor limits in terms of allowed parallelism this increases stability of the deletes)
// sometimes there will be two folders with the same name at the same level due to issues with creating directories in
// parallel. This mutex just synchronizes everything to create folders one at a time. This mutex will be removed when mkdirs
// is removed from the notebooks resource. Then we will switch to TF resource retry.
var mkdirMtx = &sync.Mutex{}

// Create creates a notebook given the content and path
func (a NotebooksAPI) Create(path string, content string, language Language, format ExportFormat, overwrite bool) error {
	notebookCreateRequest := NotebookImportRequest{}
	notebookCreateRequest.Content = content
	notebookCreateRequest.Language = language
	notebookCreateRequest.Path = path
	notebookCreateRequest.Format = format
	notebookCreateRequest.Overwrite = overwrite
	return a.C.Post("/workspace/import", notebookCreateRequest, nil)
}

// Read returns the notebook metadata and not the contents
func (a NotebooksAPI) Read(path string) (WorkspaceObjectStatus, error) {
	var notebookInfo WorkspaceObjectStatus
	err := a.C.Get("/workspace/get-status", map[string]string{
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
	err := a.C.Get("/workspace/export", workspacePathRequest{
		Format: format,
		Path:   path,
	}, &notebookContent)
	return notebookContent.Content, err
}

// Mkdirs will make folders in a workspace recursively given a path
func (a NotebooksAPI) Mkdirs(path string) error {
	// This mutex will be removed when mkdirs is removed from the notebooks resource.
	// Then we will switch to TF resource retry.
	mkdirMtx.Lock() // this mutex might also be needed for /workspace/import
	defer mkdirMtx.Unlock()
	return a.C.Post("/workspace/mkdirs", map[string]string{
		"path": path,
	}, nil)
}

// List will list all objects in a path on the workspace and with the recursive flag it will recursively list
// all the objects
func (a NotebooksAPI) List(path string, recursive bool) ([]WorkspaceObjectStatus, error) {
	if recursive {
		var paths []WorkspaceObjectStatus
		err := a.recursiveAddPaths(path, &paths)
		if err != nil {
			return nil, err
		}
		return paths, err
	}
	return a.list(path)
}

func (a NotebooksAPI) recursiveAddPaths(path string, pathList *[]WorkspaceObjectStatus) error {
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

func (a NotebooksAPI) list(path string) ([]WorkspaceObjectStatus, error) {
	var notebookList struct {
		Objects []WorkspaceObjectStatus `json:"objects,omitempty" url:"objects,omitempty"`
	}
	err := a.C.Get("/workspace/list", map[string]string{
		"path": path,
	}, &notebookList)
	return notebookList.Objects, err
}

// Delete will delete folders given a path and recursive flag
func (a NotebooksAPI) Delete(path string, recursive bool) error {
	return a.C.Post("/workspace/delete", NotebookDeleteRequest{
		Path:      path,
		Recursive: recursive,
	}, nil)
}

func ResourceNotebook() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebookCreate,
		Read:   resourceNotebookRead,
		Delete: resourceNotebookDelete,

		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				StateFunc: func(i interface{}) string {
					base64String := i.(string)
					base64, err := convertBase64ToCheckSum(base64String)
					if err != nil {
						return ""
					}
					return base64
				},
				ValidateFunc: validation.StringIsBase64,
			},
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: ValidateNotebookPath,
			},
			"language": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(Scala),
					string(Python),
					string(R),
					string(SQL),
				}, false),
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"mkdirs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  string(Source),
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					// Only supports source format as it is easiest to identify diff/delta
					string(Source),
				}, false),
			},
			"object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceNotebookCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	path := d.Get("path").(string)
	content := d.Get("content").(string)
	language := d.Get("language").(string)
	format := d.Get("format").(string)
	overwrite := d.Get("overwrite").(bool)
	mkdirs := d.Get("mkdirs").(bool)

	if mkdirs {
		parentDir, err := internal.GetParentDirPath(path)
		switch err {
		// Notebook path is root directory so no need to make directory and there is no parent
		case internal.DirPathRootDirError:
			break
		// Notebook path is empty thus a valid error
		case internal.PathEmptyError:
			return err
		//	Notebook path is valid and has a parent directory
		case nil:
			workspaceObj, err := NewNotebooksAPI(client).Read(parentDir)
			// Notebook parent path is not a directory and it could be a notebook
			if err == nil && workspaceObj.ObjectType != Directory {
				return fmt.Errorf("parent path %s should be a directory and not a %s",
					workspaceObj.Path,
					workspaceObj.ObjectType,
				)
			}
			// Parent path is missing thus needs to be created as a directory
			if e, ok := err.(common.APIError); ok && e.IsMissing() {
				err := NewNotebooksAPI(client).Mkdirs(parentDir)
				if err != nil {
					return errors.Wrapf(err, "failed to create directory %s", parentDir)
				}
			}
		}
	}

	err := NewNotebooksAPI(client).Create(path, content, Language(language), ExportFormat(format), overwrite)
	if err != nil {
		return err
	}
	d.SetId(path)

	err = d.Set("format", format)
	if err != nil {
		return err
	}
	err = d.Set("overwrite", overwrite)
	if err != nil {
		return err
	}
	err = d.Set("mkdirs", mkdirs)
	if err != nil {
		return err
	}
	return resourceNotebookRead(d, m)
}

func resourceNotebookRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	format := d.Get("format").(string)
	notebookData, err := NewNotebooksAPI(client).Export(id, ExportFormat(format))
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	notebookInfo, err := NewNotebooksAPI(client).Read(id)
	if err != nil {
		return err
	}
	err = d.Set("path", id)
	if err != nil {
		return err
	}

	crc, err := convertBase64ToCheckSum(notebookData)
	if err != nil {
		return err
	}
	err = d.Set("content", crc)
	if err != nil {
		return err
	}
	err = d.Set("language", string(notebookInfo.Language))
	if err != nil {
		return err
	}
	err = d.Set("object_id", int(notebookInfo.ObjectID))
	if err != nil {
		return err
	}
	err = d.Set("object_type", string(notebookInfo.ObjectType))

	return err
}

func resourceNotebookDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)

	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		err := NewNotebooksAPI(client).Delete(id, true)
		if err == nil {
			return nil
		}
		var e common.APIError
		if errors.As(err, &e) && e.IsTooManyRequests() {
			// Wait for requests to clear up
			baseDuration := 250 * time.Millisecond
			jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
			// So all threads dont wake up at once
			time.Sleep(baseDuration + jitter)
			return resource.RetryableError(fmt.Errorf("request rate limit hit %w", err))
		}
		return resource.NonRetryableError(err)
	})
}

func convertBase64ToCheckSum(b64 string) (string, error) {
	dataArr, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "error", errors.Wrap(err, "error while trying to decode base64 content")
	}
	return strconv.Itoa(int(crc32.ChecksumIEEE(normalizeNotebookSource(dataArr)))), nil
}

// This will normalize the notebooks to only validate the contents of the notebook that is editable
// and not be impacted by autogenerated text by databricks import and export
func normalizeNotebookSource(dataArr []byte) []byte {
	scalaMagic := "// MAGIC "
	pythonRMagic := "# MAGIC "
	sqlMagic := "-- MAGIC "
	filteredDataArr := filter(strings.Split(string(dataArr), "\n"), func(s string) bool {
		trimmedS := strings.TrimRight(s, " ")
		scalaMagicStatements := trimmedS != strings.Trim(scalaMagic, " ")
		pythonRMagicStatements := trimmedS != strings.Trim(pythonRMagic, " ")
		sqlMagicStatements := trimmedS != strings.Trim(sqlMagic, " ")
		ignoreScalaCmd := trimmedS != "// COMMAND ----------"
		ignorePythonRCmd := trimmedS != "# COMMAND ----------"
		ignoreSqlCmd := trimmedS != "-- COMMAND ----------"
		ignoreNotebookHeader := !strings.HasSuffix(trimmedS, "Databricks notebook source")
		return scalaMagicStatements && pythonRMagicStatements && sqlMagicStatements &&
			ignoreScalaCmd && ignorePythonRCmd && ignoreSqlCmd && ignoreNotebookHeader && trimmedS != ""
	})
	transformedDataArr := transform(filteredDataArr, func(s string) string {
		if strings.HasPrefix(s, scalaMagic) {
			return strings.TrimPrefix(s, scalaMagic)
		}
		if strings.HasPrefix(s, pythonRMagic) {
			return strings.TrimPrefix(s, pythonRMagic)
		}
		if strings.HasPrefix(s, sqlMagic) {
			return strings.TrimPrefix(s, sqlMagic)
		}
		return s
	})
	return []byte(strings.Join(transformedDataArr, "\n"))
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func transform(ss []string, t func(string) string) (ret []string) {
	for _, s := range ss {
		ret = append(ret, t(s))
	}
	return
}

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
