package workspace

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

// ReposResponse provides information about given repository
type ReposResponse struct {
	Id           int64  `json:"id"`
	Url          string `json:"url"`
	Provider     string `json:"provider"`
	Path         string `json:"path"`
	Branch       string `json:"branch"`
	HeadCommitId string `json:"head_commit_id"`
}

type createRequest struct {
	Url      string `json:"url"`
	Provider string `json:"provider"`
	Path     string `json:"path,omitempty"`
}

func (a ReposAPI) Create(r createRequest) (ReposResponse, error) {
	var resp ReposResponse
	err := a.client.Post(a.context, "/repos", r, &resp)
	return resp, err
}

func (a ReposAPI) Delete(id string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/repos/%s", id), nil)
}

func (a ReposAPI) Update(id string, r map[string]string) error {
	return a.client.Patch(a.context, fmt.Sprintf("/repos/%s", id), r)
}

func (a ReposAPI) Read(id string) (ReposResponse, error) {
	var resp ReposResponse
	err := a.client.Get(a.context, fmt.Sprintf("/repos/%s", id), nil, &resp)
	return resp, err
}

type ReposListResponse struct {
	NextPageToken string          `json:"next_page_token,omitempty"`
	Repos         []ReposResponse `json:"repos"`
}

func (a ReposAPI) List(prefix string) ([]ReposResponse, error) {
	req := map[string]string{}
	if prefix != "" {
		req["path_prefix"] = prefix
	}
	reposList := []ReposResponse{}
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

func (a ReposAPI) ListAll() ([]ReposResponse, error) {
	return a.List("")
}

var gitProvidersMap = map[string]string{
	"github.com":    "gitHub",
	"dev.azure.com": "azureDevOpsServices",
	"gitlab.com":    "gitLab",
	"bitbucket.org": "bitbucketCloud",
}

func getProviderFromUrl(uri string) string {
	provider := ""
	u, err := url.Parse(uri)
	if err == nil {
		provider = gitProvidersMap[strings.ToLower(u.Host)]
	}
	return provider
}

func ResourceRepo() *schema.Resource {
	s := map[string]*schema.Schema{
		"url": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"git_provider": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return strings.EqualFold(old, new)
			},
		},
		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true, // TODO: remove after the Update API will support changing the path
		},
		"branch": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"tag"},
		},
		"tag": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"branch"},
		},
		"commit_hash": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			reposAPI := NewReposAPI(ctx, c)
			url := d.Get("url").(string)
			provider := d.Get("git_provider").(string)
			if provider == "" { // trying to infer Git Provider from the URL
				provider = getProviderFromUrl(url)
			}
			if provider == "" {
				return fmt.Errorf("git_provider isn't specified and we can't detect provider from URL")
			}
			req := createRequest{Path: d.Get("path").(string), Provider: provider, Url: url}
			resp, err := reposAPI.Create(req)
			if err != nil {
				return err
			}
			d.SetId(fmt.Sprintf("%d", resp.Id))
			branch := d.Get("branch").(string)
			tag := d.Get("tag").(string)
			if tag != "" {
				d.Set("branch", "")
				return reposAPI.Update(d.Id(), map[string]string{"tag": tag})
			} else if branch != "" && branch != resp.Branch {
				return reposAPI.Update(d.Id(), map[string]string{"branch": branch})
			}

			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			reposAPI := NewReposAPI(ctx, c)
			resp, err := reposAPI.Read(d.Id())
			if err != nil {
				return err
			}
			d.SetId(fmt.Sprintf("%d", resp.Id))
			d.Set("url", resp.Url)
			d.Set("path", resp.Path)
			d.Set("git_provider", resp.Provider)
			d.Set("commit_hash", resp.HeadCommitId)
			if d.Get("branch").(string) == "" {
				d.Set("branch", resp.Branch)
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			// TODO: update may change ONE OF (url AND provider (optional)), (path), or (branch OR tag).
			// for URL/provider force re-create as there are limits on what could be done for changing URL/provider
			reposAPI := NewReposAPI(ctx, c)
			// Not working yet, wait until API is ready
			// if d.HasChange("path") {
			// 	err := reposAPI.Update(d.Id(), map[string]string{"path": d.Get("path").(string)})
			// 	if err != nil {
			// 		return err
			// 	}
			// }
			if d.HasChange("branch") || d.HasChange("tag") {
				branch := d.Get("branch").(string)
				tag := d.Get("tag").(string)
				if tag != "" {
					d.Set("branch", "")
					return reposAPI.Update(d.Id(), map[string]string{"tag": tag})
				}
				return reposAPI.Update(d.Id(), map[string]string{"branch": branch})
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewReposAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
