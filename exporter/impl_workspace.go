package exporter

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	sdk_workspace "github.com/databricks/databricks-sdk-go/service/workspace"
)

func ImportNotebook(ic *importContext, r *resource) error {
	ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
	resp, err := ic.workspaceClient.Workspace.Export(ic.Context, sdk_workspace.ExportRequest{
		Path:   r.ID,
		Format: sdk_workspace.ExportFormat(ic.notebooksFormat),
	})
	if err != nil {
		if apierr.IsMissing(err) {
			ic.addIgnoredResource(fmt.Sprintf("databricks_notebook. path=%s", r.ID))
		}
		return err
	}
	var fileExtension string
	if ic.notebooksFormat == "SOURCE" {
		language := r.Data.Get("language").(string)
		fileExtension = fileExtensionLanguageMapping[language]
		r.Data.Set("language", "")
	} else {
		fileExtension = fileExtensionFormatMapping[ic.notebooksFormat]
	}
	r.Data.Set("format", ic.notebooksFormat)
	objectId := r.Data.Get("object_id").(int)
	name := fileNameNormalizationRegex.ReplaceAllString(r.ID[1:], "_") + "_" + strconv.Itoa(objectId) + fileExtension
	content, _ := base64.StdEncoding.DecodeString(resp.Content)
	fileName, err := ic.saveFileIn("notebooks", name, []byte(content))
	if err != nil {
		return err
	}
	if ic.notebooksFormat == "SOURCE" && r.Data.Get("language").(string) == "PYTHON" {
		analyzeNotebook(ic, string(content))
	}
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/notebooks/%d", objectId),
		"notebook_"+ic.Importables["databricks_notebook"].Name(ic, r.Data))
	// TODO: it's not completely correct condition - we need to make emit smarter -
	// emit only if permissions are different from their parent's permission.
	ic.emitWorkspaceObjectParentDirectory(r)
	return r.Data.Set("source", fileName)
}

func analyzeNotebook(ic *importContext, content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}

		if strings.HasPrefix(strings.TrimSpace(line), "%pip") && strings.Contains(line, " install ") {
			parts := strings.Fields(line)
			if len(parts) < 3 {
				continue
			}
			for _, part := range parts[2:] {
				if strings.HasPrefix(part, "-") || strings.Contains(part, "==") {
					continue
				}
				if strings.HasPrefix(part, "/dbfs/") {
					ic.Emit(&resource{
						Resource: "databricks_dbfs_file",
						ID:       strings.TrimPrefix(part, "/dbfs"),
					})
				} else if strings.HasPrefix(part, "/Workspace/") {
					ic.Emit(&resource{
						Resource: "databricks_workspace_file",
						ID:       strings.TrimPrefix(part, "/Workspace"),
					})
				} else if strings.HasPrefix(part, "/Volumes/") {
					ic.Emit(&resource{
						Resource: "databricks_file",
						ID:       part,
					})
				}
			}
		}
	}

}
