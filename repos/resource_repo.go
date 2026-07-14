package repos

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// repoInfo wraps the Go SDK workspace.RepoInfo so the databricks_repo schema can be
// generated directly from the SDK struct. The Terraform-specific attribute names
// (git_provider, commit_hash) are supplied via Aliases() rather than a parallel
// hand-maintained struct.
type repoInfo struct {
	workspace.RepoInfo
	common.Namespace
}

func (repoInfo) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"repos.repoInfo": {
			"provider":       "git_provider",
			"head_commit_id": "commit_hash",
		},
	}
}

func (repoInfo) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	s.SchemaPath("url").SetForceNew().SetValidateFunc(validation.IsURLWithScheme([]string{"https", "http"}))
	s.SchemaPath("git_provider").SetComputed().SetForceNew().SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	s.SchemaPath("sparse_checkout").SetForceNew()
	s.SchemaPath("path").SetComputed().SetForceNew().SetValidateFunc(validatePath)
	s.SchemaPath("branch").SetComputed().SetConflictsWith([]string{"tag"}).SetValidateFunc(validation.StringIsNotWhiteSpace)
	s.SchemaPath("commit_hash").SetComputed()

	s.AddNewField("tag", &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"branch"},
		ValidateFunc:  validation.StringIsNotWhiteSpace,
	})
	s.AddNewField("workspace_path", &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	})

	s.RemoveField("id")
	common.NamespaceCustomizeSchema(s)
	return s
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
	s := common.StructToSchema(repoInfo{}, nil)

	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var repo repoInfo
			common.DataToStructPointer(d, s, &repo)

			provider := repo.Provider
			if provider == "" { // trying to infer Git Provider from the URL
				provider = GetGitProviderFromUrl(repo.Url)
			}
			if provider == "" {
				return fmt.Errorf("git_provider isn't specified and we can't detect provider from URL")
			}
			if repo.Path != "" {
				p := path.Dir(strings.TrimSuffix(repo.Path, "/"))
				if err := w.Workspace.MkdirsByPath(ctx, p); err != nil {
					return err
				}
			}

			createReq := workspace.CreateRepoRequest{
				Url:      repo.Url,
				Provider: provider,
				Path:     repo.Path,
			}
			if repo.SparseCheckout != nil {
				createReq.SparseCheckout = &workspace.SparseCheckout{Patterns: repo.SparseCheckout.Patterns}
			}
			resp, err := w.Repos.Create(ctx, createReq)
			if err != nil {
				return err
			}
			d.SetId(strconv.FormatInt(resp.Id, 10))

			branch := d.Get("branch").(string)
			tag := d.Get("tag").(string)
			updateReq := workspace.UpdateRepoRequest{RepoId: resp.Id}
			needsUpdate := false
			if tag != "" {
				updateReq.Tag = tag
				needsUpdate = true
			} else if branch != "" && branch != resp.Branch {
				updateReq.Branch = branch
				needsUpdate = true
			}
			if !needsUpdate {
				return nil
			}
			return w.Repos.Update(ctx, updateReq)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			resp, err := w.Repos.GetByRepoId(ctx, id)
			if err != nil {
				return err
			}
			// GetByRepoId returns a GetRepoResponse; the schema is built from RepoInfo.
			// The two are structurally identical, so copy across into the schema struct.
			repo := repoInfo{RepoInfo: workspace.RepoInfo{
				Id:              resp.Id,
				Url:             resp.Url,
				Provider:        resp.Provider,
				Path:            resp.Path,
				Branch:          resp.Branch,
				HeadCommitId:    resp.HeadCommitId,
				SparseCheckout:  resp.SparseCheckout,
				ForceSendFields: resp.ForceSendFields,
			}}
			if err = common.StructToData(repo, s, d); err != nil {
				return err
			}
			d.Set("workspace_path", "/Workspace"+resp.Path)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var repo repoInfo
			common.DataToStructPointer(d, s, &repo)

			id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			req := workspace.UpdateRepoRequest{RepoId: id}
			// Not working yet, wait until API is ready:
			// path updates are gated on the Update API supporting them.
			if d.HasChange("tag") {
				req.Tag = d.Get("tag").(string)
				d.Set("branch", "")
			} else if d.HasChange("branch") {
				req.Branch = repo.Branch
				d.Set("tag", "")
			} else {
				if repo.Branch != "" {
					req.Branch = repo.Branch
				} else if v := d.Get("tag").(string); v != "" {
					req.Tag = v
				}
			}
			if repo.SparseCheckout != nil {
				req.SparseCheckout = &workspace.SparseCheckoutUpdate{Patterns: repo.SparseCheckout.Patterns}
			}
			return w.Repos.Update(ctx, req)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			id, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				return err
			}
			return w.Repos.DeleteByRepoId(ctx, id)
		},
	}
}
