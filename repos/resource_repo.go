package repos

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ReposAPI exposes the Repos API
type ReposAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewReposAPI creates ReposAPI instance from provider meta
func NewReposAPI(ctx context.Context, m any) ReposAPI {
	return ReposAPI{m.(*common.DatabricksClient), ctx}
}

type ReposSparseCheckout struct {
	Patterns []string `json:"patterns"`
}

// ReposInformation provides information about given repository
type ReposInformation struct {
	common.Namespace
	ID             int64                `json:"id"`
	Url            string               `json:"url" tf:"force_new"`
	Provider       string               `json:"provider,omitempty" tf:"computed,alias:git_provider,force_new"`
	SparseCheckout *ReposSparseCheckout `json:"sparse_checkout,omitempty" tf:"force_new"`
	Path           string               `json:"path,omitempty" tf:"computed,force_new"` // TODO: remove force_new after the Update API will support changing the path
	Branch         string               `json:"branch,omitempty" tf:"computed"`
	HeadCommitID   string               `json:"head_commit_id,omitempty" tf:"computed,alias:commit_hash"`
}

// RepoID returns job id as string
func (r ReposInformation) RepoID() string {
	return fmt.Sprintf("%d", r.ID)
}

type reposCreateRequest struct {
	Url            string               `json:"url"`
	Provider       string               `json:"provider"`
	Path           string               `json:"path,omitempty"`
	SparseCheckout *ReposSparseCheckout `json:"sparse_checkout,omitempty"`
}

func (a ReposAPI) Create(r reposCreateRequest) (ReposInformation, error) {
	var resp ReposInformation
	if r.Provider == "" { // trying to infer Git Provider from the URL
		r.Provider = GetGitProviderFromUrl(r.Url)
	}
	if r.Provider == "" {
		return resp, fmt.Errorf("git_provider isn't specified and we can't detect provider from URL")
	}
	if r.Path != "" {
		p := path.Dir(strings.TrimSuffix(r.Path, "/"))
		if err := workspace.NewNotebooksAPI(a.context, a.client).Mkdirs(p); err != nil {
			return resp, err
		}
	}

	err := a.client.Post(a.context, "/repos", r, &resp)
	return resp, err
}

func (a ReposAPI) Delete(id string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/repos/%s", id), nil)
}

func (a ReposAPI) Update(id string, r map[string]any) error {
	if len(r) == 0 {
		return nil
	}
	// TODO: update may change ONE OF (url AND provider (optional)), (path), or (branch OR tag).
	// for URL/provider force re-create as there are limits on what could be done for changing URL/provider
	if path, ok := r["path"]; ok {
		err := a.client.Patch(a.context, fmt.Sprintf("/repos/%s", id), map[string]any{"path": path})
		if err != nil {
			return err
		}
		delete(r, "path")
	}
	return a.client.Patch(a.context, fmt.Sprintf("/repos/%s", id), r)
}

func (a ReposAPI) Read(id string) (ReposInformation, error) {
	var resp ReposInformation
	err := a.client.Get(a.context, fmt.Sprintf("/repos/%s", id), nil, &resp)
	return resp, err
}

type ReposListResponse struct {
	NextPageToken string             `json:"next_page_token,omitempty"`
	Repos         []ReposInformation `json:"repos"`
}

func (a ReposAPI) List(prefix string) ([]ReposInformation, error) {
	req := map[string]string{}
	if prefix != "" {
		req["path_prefix"] = prefix
	}
	reposList := []ReposInformation{}
	for {
		var resp ReposListResponse
		err := a.client.Get(a.context, "/repos", req, &resp)
		if err != nil {
			return nil, err
		}
		reposList = append(reposList, resp.Repos...)
		if resp.NextPageToken == "" {
			break
		}
		req["next_page_token"] = resp.NextPageToken
	}
	return reposList, nil
}

func (a ReposAPI) ListAll() ([]ReposInformation, error) {
	return a.List("")
}

var (
	gitProvidersMap = map[string]string{
		"github.com":    "gitHub",
		"dev.azure.com": "azureDevOpsServices",
		"gitlab.com":    "gitLab",
		"bitbucket.org": "bitbucketCloud",
	}
	awsCodeCommitRegex = regexp.MustCompile(`^git-codecommit\.[^.]+\.amazonaws\.com$`)
)

func GetGitProviderFromUrl(uri string) string {
	provider := ""
	u, err := url.Parse(uri)
	if err == nil {
		lhost := strings.ToLower(u.Host)
		provider = gitProvidersMap[lhost]
		if provider == "" && awsCodeCommitRegex.FindStringSubmatch(lhost) != nil {
			provider = "awsCodeCommit"
		}
	}
	return provider
}

func validatePath(i interface{}, k string) (_ []string, errors []error) {
	v := i.(string)
	if v == "" || !strings.HasPrefix(v, "/Repos/") {
		return
	}
	v = strings.TrimSuffix(v, "/")
	parts := strings.Split(v, "/")
	if len(parts) != 4 { // we require 3 path parts + starting /
		errors = append(errors, fmt.Errorf("should have 3 components (/Repos/<directory>/<repo>), got %d", len(parts)-1))
	}
	return
}

func ResourceRepo() common.Resource {
	s := common.StructToSchema(ReposInformation{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["url"].ValidateFunc = validation.IsURLWithScheme([]string{"https", "http"})
		s["git_provider"].DiffSuppressFunc = common.EqualFoldDiffSuppress
		s["branch"].ConflictsWith = []string{"tag"}
		s["branch"].ValidateFunc = validation.StringIsNotWhiteSpace
		s["path"].ValidateFunc = validatePath

		s["tag"] = &schema.Schema{
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"branch"},
			ValidateFunc:  validation.StringIsNotWhiteSpace,
		}
		s["workspace_path"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}

		delete(s, "id")
		return s
	})

	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			reposAPI := NewReposAPI(ctx, c)
			var repo ReposInformation
			common.DataToStructPointer(d, s, &repo)

			req := reposCreateRequest{Path: repo.Path, Provider: repo.Provider,
				Url: repo.Url, SparseCheckout: repo.SparseCheckout}
			resp, err := reposAPI.Create(req)
			if err != nil {
				return err
			}
			d.SetId(resp.RepoID())
			branch := d.Get("branch").(string)
			tag := d.Get("tag").(string)
			updateReq := map[string]any{}
			if tag != "" {
				updateReq["tag"] = tag
			} else if branch != "" && branch != resp.Branch {
				updateReq["branch"] = branch
			}
			return reposAPI.Update(d.Id(), updateReq)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			reposAPI := NewReposAPI(ctx, c)
			resp, err := reposAPI.Read(d.Id())
			if err != nil {
				return err
			}
			err = common.StructToData(resp, s, d)
			if err != nil {
				return err
			}
			d.Set("workspace_path", "/Workspace"+resp.Path)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var repo ReposInformation
			common.DataToStructPointer(d, s, &repo)

			reposAPI := NewReposAPI(ctx, c)
			req := map[string]any{}
			// Not working yet, wait until API is ready
			// if d.HasChange("path") {
			// 	req["path"] = d.Get("path").(string)
			// }
			if d.HasChange("tag") {
				req["tag"] = d.Get("tag").(string)
				d.Set("branch", "")
			} else if d.HasChange("branch") {
				req["branch"] = repo.Branch
				d.Set("tag", "")
			} else {
				if repo.Branch != "" {
					req["branch"] = repo.Branch
				} else if v := d.Get("tag").(string); v != "" {
					req["tag"] = v
				}
			}
			if repo.SparseCheckout != nil {
				req["sparse_checkout"] = map[string]any{"patterns": repo.SparseCheckout.Patterns}
			}
			return reposAPI.Update(d.Id(), req)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewReposAPI(ctx, c).Delete(d.Id())
		},
	}
}
