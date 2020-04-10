package service

import (
	"encoding/json"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"net/http"
)

// TokensAPI exposes the Secrets API
type NotebooksAPI struct {
	Client DBApiClient
}

func (a NotebooksAPI) init(client DBApiClient) NotebooksAPI {
	a.Client = client
	return a
}

func (a NotebooksAPI) Create(path string, content string, language model.Language, format model.ExportFormat, overwrite bool) error {
	notebookCreateRequest := struct {
		Content   string             `json:"content,omitempty"`
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

	_, err := a.Client.performQuery(http.MethodPost, "/workspace/import", "2.0", nil, notebookCreateRequest)
	return err
}

func (a NotebooksAPI) Read(path string) (model.NotebookInfo, error) {
	var notebookInfo model.NotebookInfo
	notebookGetStatusRequest := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{}
	notebookGetStatusRequest.Path = path
	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/get-status", "2.0", nil, notebookGetStatusRequest)
	if err != nil {
		return notebookInfo, err
	}

	err = json.Unmarshal(resp, &notebookInfo)
	return notebookInfo, err
}

func (a NotebooksAPI) Export(path string, format model.ExportFormat) (string, error) {
	var notebookContent map[string]string
	notebookExportRequest := struct {
		Path   string             `json:"path,omitempty" url:"path,omitempty"`
		Format model.ExportFormat `json:"format,omitempty" url:"format,omitempty"`
	}{}
	notebookExportRequest.Path = path
	notebookExportRequest.Format = format
	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/export", "2.0", nil, notebookExportRequest)
	if err != nil {
		return notebookContent["content"], err
	}

	err = json.Unmarshal(resp, &notebookContent)
	return notebookContent["content"], err
}

func (a NotebooksAPI) Mkdirs(path string) error {
	mkDirsRequest := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{}
	mkDirsRequest.Path = path

	_, err := a.Client.performQuery(http.MethodPost, "/workspace/mkdirs", "2.0", nil, mkDirsRequest)

	return err
}

func (a NotebooksAPI) List(path string, recursive bool) ([]model.NotebookInfo, error) {
	if recursive == true {
		var paths []model.NotebookInfo
		a.recursiveAddPaths(path, &paths)
		return paths, nil
	} else {
		return a.list(path)
	}
}

func (a NotebooksAPI) recursiveAddPaths(path string, pathList *[]model.NotebookInfo) {
	notebookInfoList, _ := a.list(path)
	for _, v := range notebookInfoList {
		if v.ObjectType == model.Notebook {
			*pathList = append(*pathList, v)
		} else if v.ObjectType == model.Directory {
			a.recursiveAddPaths(v.Path, pathList)
		}
	}
}

func (a NotebooksAPI) list(path string) ([]model.NotebookInfo, error) {
	var notebookList struct {
		Objects []model.NotebookInfo `json:"objects,omitempty" url:"objects,omitempty"`
	}
	listRequest := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{}
	listRequest.Path = path

	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/list", "2.0", nil, listRequest)
	if err != nil {
		return notebookList.Objects, err
	}

	err = json.Unmarshal(resp, &notebookList)
	return notebookList.Objects, err
}

func (a NotebooksAPI) Delete(path string, recursive bool) error {
	notebookDelete := struct {
		Path      string `json:"path,omitempty"`
		Recursive bool   `json:"recursive,omitempty"`
	}{}
	notebookDelete.Path = path
	notebookDelete.Recursive = recursive
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/delete", "2.0", nil, notebookDelete)
	return err
}
