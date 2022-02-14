package workspace

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ReposAPI exposes the Repos API
type ReposAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewReposAPI creates ReposAPI instance from provider meta
func NewReposAPI(ctx context.Context, m interface{}) ReposAPI {
	return ReposAPI{m.(*common.DatabricksClient), ctx}
}

// ReposInformation provides information about given repository
type ReposInformation struct {
	ID           int64  `json:"id"`
	Url          string `json:"url" tf:"force_new"`
	Provider     string `json:"provider,omitempty" tf:"computed,alias:git_provider,force_new"`
	Path         string `json:"path,omitempty" tf:"computed,force_new"` // TODO: remove force_new after the Update API will support changing the path
	Branch       string `json:"branch,omitempty" tf:"computed"`
	HeadCommitID string `json:"head_commit_id,omitempty" tf:"computed,alias:commit_hash"`
}

// ID returns job id as string
func (r ReposInformation) RepoID() string {
	return fmt.Sprintf("%d", r.ID)
}

type createRequest struct {
	Url      string `json:"url"`
	Provider string `json:"provider"`
	Path     string `json:"path,omitempty"`
}

func (a ReposAPI) Create(r createRequest) (ReposInformation, error) {
	var resp ReposInformation
	if r.Provider == "" { // trying to infer Git Provider from the URL
		r.Provider = GetGitProviderFromUrl(r.Url)
	}
	if r.Provider == "" {
		return resp, fmt.Errorf("git_provider isn't specified and we can't detect provider from URL")
	}
	if r.Path != "" {
		if !strings.HasPrefix(r.Path, "/Repos/") {
			return resp, fmt.Errorf("path should start with /Repos/")
		}
		p := r.Path
		if strings.HasSuffix(r.Path, "/") {
			p = strings.TrimSuffix(r.Path, "/")
		}
		p = path.Dir(p)
		if err := NewNotebooksAPI(a.context, a.client).Mkdirs(p); err != nil {
			return resp, err
		}
	}

	err := a.client.Post(a.context, "/repos", r, &resp)
	return resp, err
}

func (a ReposAPI) Delete(id string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/repos/%s", id), nil)
}

func (a ReposAPI) Update(id string, r map[string]string) error {
	if len(r) == 0 {
		return nil
	}
	// TODO: update may change ONE OF (url AND provider (optional)), (path), or (branch OR tag).
	// for URL/provider force re-create as there are limits on what could be done for changing URL/provider
	if path, ok := r["path"]; ok {
		err := a.client.Patch(a.context, fmt.Sprintf("/repos/%s", id), map[string]string{"path": path})
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

var gitProvidersMap = map[string]string{
	"github.com":    "gitHub",
	"dev.azure.com": "azureDevOpsServices",
	"gitlab.com":    "gitLab",
	"bitbucket.org": "bitbucketCloud",
}

func GetGitProviderFromUrl(uri string) string {
	provider := ""
	u, err := url.Parse(uri)
	if err == nil {
		provider = gitProvidersMap[strings.ToLower(u.Host)]
	}
	return provider
}

func ResourceRepo() *schema.Resource {
	s := common.StructToSchema(ReposInformation{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["url"].ValidateFunc = validation.IsURLWithScheme([]string{"https", "http"})
		s["git_provider"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
			return strings.EqualFold(old, new)
		}
		s["branch"].ConflictsWith = []string{"tag"}
		s["branch"].ValidateFunc = validation.StringIsNotWhiteSpace

		s["tag"] = &schema.Schema{
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"branch"},
			ValidateFunc:  validation.StringIsNotWhiteSpace,
		}

		delete(s, "id")
		return s
	})

	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			reposAPI := NewReposAPI(ctx, c)
			req := createRequest{Path: d.Get("path").(string), Provider: d.Get("git_provider").(string), Url: d.Get("url").(string)}
			resp, err := reposAPI.Create(req)
			if err != nil {
				return err
			}
			d.SetId(resp.RepoID())
			branch := d.Get("branch").(string)
			tag := d.Get("tag").(string)
			updateReq := map[string]string{}
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
			return common.StructToData(resp, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			reposAPI := NewReposAPI(ctx, c)
			req := map[string]string{}
			// Not working yet, wait until API is ready
			// if d.HasChange("path") {
			// 	req["path"] = d.Get("path").(string)
			// }
			if d.HasChange("tag") {
				req["tag"] = d.Get("tag").(string)
				d.Set("branch", "")
			} else if d.HasChange("branch") {
				req["branch"] = d.Get("branch").(string)
			}
			return reposAPI.Update(d.Id(), req)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewReposAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
