package exporter

import (
	"encoding/base64"
	"fmt"
	"log"
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
	fileName, err := ic.saveContentIn("notebooks", name, []byte(content))
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

func searchRepoByPath(ic *importContext, r *resource) error {
	repoDir, err := ic.workspaceClient.Workspace.GetStatusByPath(ic.Context, r.Value)
	if err != nil {
		return err
	}
	if repoDir.ObjectType != sdk_workspace.ObjectTypeRepo {
		return fmt.Errorf("object %s is not a repo", r.Value)
	}
	if repoDir.ResourceId != "" {
		r.ID = repoDir.ResourceId
	} else {
		r.ID = strconv.FormatInt(repoDir.ObjectId, 10)
	}
	return nil
}

func listRepos(ic *importContext) error {
	it := ic.workspaceClient.Repos.List(ic.Context, sdk_workspace.ListReposRequest{PathPrefix: "/Workspace"})
	i := 1
	for it.HasNext(ic.Context) {
		repo, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		if !ic.MatchesName(repo.Path) {
			log.Printf("[INFO] Repo %s doesn't match %s filter", repo.Path, ic.match)
			continue
		}
		if repo.Url != "" {
			ic.Emit(&resource{
				Resource: "databricks_repo",
				ID:       strconv.FormatInt(repo.Id, 10),
			})
		} else {
			log.Printf("[WARN] ignoring databricks_repo without Git provider. Path: %s", repo.Path)
			ic.addIgnoredResource(fmt.Sprintf("databricks_repo. path=%s", repo.Path))
		}
		if i%50 == 0 {
			log.Printf("[INFO] Scanned %d repos", i)
		}
		i++
	}
	return nil
}

func importRepo(ic *importContext, r *resource) error {
	path := maybeStripWorkspacePrefix(r.Data.Get("path").(string))
	if strings.HasPrefix(path, "/Repos") {
		ic.emitUserOrServicePrincipalForPath(path, "/Repos")
	} else if strings.HasPrefix(path, "/Users") {
		ic.emitUserOrServicePrincipalForPath(path, "/Users")
	}
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/repos/%s", r.ID),
		"repo_"+ic.Importables["databricks_repo"].Name(ic, r.Data))
	return nil
}

func importWorkspaceFile(ic *importContext, r *resource) error {
	ic.emitUserOrServicePrincipalForPath(r.ID, "/Users")
	reader, err := ic.workspaceClient.Workspace.Download(ic.Context, r.ID,
		func(q map[string]any) {
			q["format"] = "AUTO"
		})
	if err != nil {
		if apierr.IsMissing(err) {
			ic.addIgnoredResource(fmt.Sprintf("databricks_workspace_file. path=%s", r.ID))
		}
		return err
	}
	defer reader.Close()
	objectId := r.Data.Get("object_id").(int)
	parts := strings.Split(r.ID, "/")
	plen := len(parts)
	if idx := strings.Index(parts[plen-1], "."); idx != -1 {
		parts[plen-1] = parts[plen-1][:idx] + "_" + strconv.Itoa(objectId) + parts[plen-1][idx:]
	} else {
		parts[plen-1] = parts[plen-1] + "_" + strconv.Itoa(objectId)
	}
	name := fileNameNormalizationRegex.ReplaceAllString(strings.Join(parts, "/")[1:], "_")
	fileName, err := ic.saveReaderIn("workspace_files", name, reader)
	if err != nil {
		return err
	}

	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/files/%d", objectId),
		"ws_file_"+ic.Importables["databricks_workspace_file"].Name(ic, r.Data))

	// TODO: it's not completely correct condition - we need to make emit smarter -
	// emit only if permissions are different from their parent's permission.
	ic.emitWorkspaceObjectParentDirectory(r)
	log.Printf("[TRACE] Creating %s for %s", fileName, r)
	return r.Data.Set("source", fileName)
}
