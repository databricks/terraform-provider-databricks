package service

import (
	"encoding/json"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// NotebooksAPI exposes the Notebooks API
type NotebooksAPI struct {
	Client *DBApiClient
}

// Create creates a notebook given the content and path
func (a NotebooksAPI) Create(path string, content string, language model.Language, format model.ExportFormat, overwrite bool) error {
	notebookCreateRequest := struct {
		Content   string             `json:"content,omitempty" mask:"true"`
		Path      string             `json:"path,omitempty"`
		Language  model.Language     `json:"language,omitempty"`
		Overwrite bool               `json:"overwrite,omitempty"`
		Format    model.ExportFormat `json:"format,omitempty"`
	}{}
	notebookCreateRequest.Content = content
	notebookCreateRequest.Language = language
	notebookCreateRequest.Path = path
	notebookCreateRequest.Format = format
	notebookCreateRequest.Overwrite = overwrite

	_, err := a.Client.performQuery(http.MethodPost, "/workspace/import", "2.0", nil, notebookCreateRequest, nil)
	return err
}

// Read returns the notebook metadata and not the contents
func (a NotebooksAPI) Read(path string) (model.NotebookInfo, error) {
	var notebookInfo model.NotebookInfo
	notebookGetStatusRequest := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{}
	notebookGetStatusRequest.Path = path
	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/get-status", "2.0", nil, notebookGetStatusRequest, nil)
	if err != nil {
		return notebookInfo, err
	}

	err = json.Unmarshal(resp, &notebookInfo)
	return notebookInfo, err
}

// Export returns the notebook content as a base64 string
func (a NotebooksAPI) Export(path string, format model.ExportFormat) (string, error) {
	var notebookContent map[string]string
	notebookExportRequest := struct {
		Path   string             `json:"path,omitempty" url:"path,omitempty"`
		Format model.ExportFormat `json:"format,omitempty" url:"format,omitempty"`
	}{}
	notebookExportRequest.Path = path
	notebookExportRequest.Format = format
	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/export", "2.0", nil, notebookExportRequest, nil)
	if err != nil {
		return notebookContent["content"], err
	}

	err = json.Unmarshal(resp, &notebookContent)
	return notebookContent["content"], err
}

// Mkdirs will make folders in a workspace recursively given a path
func (a NotebooksAPI) Mkdirs(path string) error {
	mkDirsRequest := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{}
	mkDirsRequest.Path = path

	_, err := a.Client.performQuery(http.MethodPost, "/workspace/mkdirs", "2.0", nil, mkDirsRequest, nil)

	return err
}

// List will list all objects in a path on the workspace and with the recursive flag it will recursively list
// all the objects
func (a NotebooksAPI) List(path string, recursive bool) ([]model.NotebookInfo, error) {
	if recursive {
		var paths []model.NotebookInfo
		err := a.recursiveAddPaths(path, &paths)
		if err != nil {
			return nil, err
		}
		return paths, err
	}
	return a.list(path)
}

func (a NotebooksAPI) recursiveAddPaths(path string, pathList *[]model.NotebookInfo) error {
	notebookInfoList, err := a.list(path)
	if err != nil {
		return err
	}
	for _, v := range notebookInfoList {
		if v.ObjectType == model.Notebook {
			*pathList = append(*pathList, v)
		} else if v.ObjectType == model.Directory {
			err := a.recursiveAddPaths(v.Path, pathList)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (a NotebooksAPI) list(path string) ([]model.NotebookInfo, error) {
	var notebookList struct {
		Objects []model.NotebookInfo `json:"objects,omitempty" url:"objects,omitempty"`
	}
	listRequest := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{}
	listRequest.Path = path

	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/list", "2.0", nil, listRequest, nil)
	if err != nil {
		return notebookList.Objects, err
	}

	err = json.Unmarshal(resp, &notebookList)
	return notebookList.Objects, err
}

// Delete will delete folders given a path and recursive flag
func (a NotebooksAPI) Delete(path string, recursive bool) error {
	notebookDelete := struct {
		Path      string `json:"path,omitempty"`
		Recursive bool   `json:"recursive,omitempty"`
	}{}
	notebookDelete.Path = path
	notebookDelete.Recursive = recursive
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/delete", "2.0", nil, notebookDelete, nil)
	return err
}
