package importer

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/provider"
	"github.com/databrickslabs/databricks-terraform/storage"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
)

/** High level overview of importer design:

                                    +----------+      +--------------------+
                                    | resource +------> stateApproximation |
                                    +--^-------+      +----|-------------|-+
                                       |                   |             |
  +------------------------+           |                   |             |
  | "normal provider flow" |    +------^-----+       +-----V-----+   +---V---+
  +------------^-----------+    | importable +-------> reference |   | scope |
               |                +------^-----+       +--------|--+   +---V---+
+--------------+--------------+        |                      |          |
|terraform-provider-databricks|        |                      |          |
+--------------+--------------+        |                      |          |
               |                       |                   +--V----------V---+
    +----------v---------+        +----^------------+      |                 |
    |                    |        |                 |      |    Generated    |
    |  importer command  +-------->  importContext  |      |      Files      |
    |                    |        |                 |      |                 |
    +--------------------+        +-----------------+      +-----------------+
*/

type regexFix struct {
	Regex       *regexp.Regexp
	Replacement string
}

type instanceApproximation struct {
	// not really interested in other than strings...
	Attributes map[string]interface{} `json:"attributes"`
}

type resourceApproximation struct {
	Type      string                  `json:"type"`
	Name      string                  `json:"name"`
	Provider  string                  `json:"provider"`
	Mode      string                  `json:"mode"`
	Module    string                  `json:"module,omitempty"`
	Instances []instanceApproximation `json:"instances"`
}

type stateApproximation struct {
	Resources []resourceApproximation `json:"resources"`
}

type importContext struct {
	Module      string
	Client      *common.DatabricksClient
	State       stateApproximation
	Importables map[string]importable
	Resources   map[string]*schema.Resource
	Scope       []*resource
	Files       map[string]*hclwrite.File
	Directory   string
	importing   map[string]bool
	nameFixes   []regexFix
	hclFixes    []regexFix
	allUsers    []identity.ScimUser
	allGroups   []identity.ScimGroup

	debug          bool
	services       string
	lastActiveDays int64
}

func (ic *importContext) Find(r *resource, pick string) hcl.Traversal {
	for _, sr := range ic.State.Resources {
		if sr.Type != r.Resource {
			continue
		}
		for _, i := range sr.Instances {
			if i.Attributes[r.Attribute].(string) == r.Value {
				if "data" == sr.Mode {
					return hcl.Traversal{
						hcl.TraverseRoot{
							Name: "data",
						},
						hcl.TraverseAttr{
							Name: sr.Type,
						},
						hcl.TraverseAttr{
							Name: sr.Name,
						},
						hcl.TraverseAttr{
							Name: pick,
						},
					}
				}
				return hcl.Traversal{
					hcl.TraverseRoot{
						Name: sr.Type,
					},
					hcl.TraverseAttr{
						Name: sr.Name,
					},
					hcl.TraverseAttr{
						Name: pick,
					},
				}
			}
		}
	}
	return nil
}

func (ic *importContext) Has(r *resource) bool {
	if _, visiting := ic.importing[r.String()]; visiting {
		return true
	}
	k, v := r.MatchPair()
	for _, sr := range ic.State.Resources {
		if sr.Type != r.Resource {
			continue
		}
		for _, i := range sr.Instances {
			if i.Attributes[k].(string) == v {
				if "data" == sr.Mode && ic.Module != sr.Module {
					return false
				}
				return true
			}
		}
	}
	return false
}

func (ic *importContext) InstanceState(r *resource) *terraform.InstanceState {
	k, v := r.MatchPair()
	for _, sr := range ic.State.Resources {
		if sr.Type != r.Resource {
			continue
		}
		for _, i := range sr.Instances {
			if i.Attributes[k].(string) == v {
				attrs := map[string]string{}
				for k, v := range i.Attributes {
					if k == "id" {
						continue
					}
					attrs[k] = v.(string)
				}
				return &terraform.InstanceState{
					ID:         r.ID,
					Attributes: attrs,
				}
			}
		}
	}
	return nil
}

func (ic *importContext) Add(r *resource, attrs map[string]string) {
	if ic.Has(r) {
		return
	}
	inst := instanceApproximation{
		Attributes: map[string]interface{}{},
	}
	for k, v := range attrs {
		inst.Attributes[k] = v
	}
	inst.Attributes["id"] = r.ID
	ic.State.Resources = append(ic.State.Resources, resourceApproximation{
		Mode:      "managed",
		Module:    ic.Module,
		Type:      r.Resource,
		Name:      r.Name,
		Instances: []instanceApproximation{inst},
	})
	// scope this way is guaranteed to be topologically sorted
	ic.Scope = append(ic.Scope, r)
	return
}

type importable struct {
	Name    func(d *schema.ResourceData) string
	List    func(ic *importContext) error
	Search  func(ic *importContext, r *resource) error
	Import  func(ic *importContext, d *schema.ResourceData) error
	Depends []reference
	Service string
	Body    func(ic *importContext, body *hclwrite.Body, r *resource) error
}

type reference struct {
	Path     string
	Resource string
	Match    string
	Pick     string
}

type resource struct {
	Resource  string
	ID        string
	Attribute string
	Value     string
	Name      string
}

func (r *resource) MatchPair() (string, string) {
	k := "id"
	v := r.ID
	if r.ID == "" && r.Attribute != "" && r.Value != "" {
		k = r.Attribute
		v = r.Value
	}
	return k, v
}

func (r *resource) String() string {
	n := r.Name
	if n == "" {
		n = "<unknown>"
	}
	k, v := r.MatchPair()
	return fmt.Sprintf("%s[%s] (%s: %s)", r.Resource, n, k, v)
}

func (r *resource) ImportCommand(ic *importContext) string {
	m := ""
	if ic.Module != "" {
		m = ic.Module + "."
	}
	return fmt.Sprintf(`terraform import %s%s.%s "%s"`, m, r.Resource, r.Name, r.ID)
}

func instancePoolName(d *schema.ResourceData) string {
	raw, ok := d.GetOk("instance_pool_name")
	if !ok {
		return strings.Split(d.Id(), "-")[2]
	}
	name := raw.(string)
	if name == "" {
		return strings.Split(d.Id(), "-")[2]
	}
	return name
}

var resourcesMap map[string]importable = map[string]importable{
	"databricks_dbfs_file": {
		Service: "storage",
		Name: func(d *schema.ResourceData) string {
			s := strings.Split(d.Id(), "/")
			return s[len(s)-1]
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			dbfsAPI := storage.NewDBFSAPI(ic.Client)
			b64, err := dbfsAPI.Read(r.ID)
			if err != nil {
				return err
			}
			err = os.Mkdir(fmt.Sprintf("%s/files", ic.Directory), 0755)
			if err != nil && !os.IsExist(err) {
				return err
			}
			local, err := os.Create(fmt.Sprintf("%s/files/%s", ic.Directory, path.Base(r.ID)))
			if err != nil {
				return err
			}
			defer local.Close()
			fileBytes, err := base64.StdEncoding.DecodeString(b64)
			if err != nil {
				return err
			}
			_, err = local.Write(fileBytes)
			if err != nil {
				return err
			}
			// libraries installed with init scripts won't be exported.
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			relativeFile := fmt.Sprintf("${path.module}/files/%s", path.Base(r.ID))
			b.SetAttributeValue("path", cty.StringVal(r.ID))
			b.SetAttributeRaw("source", hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}},
				&hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(relativeFile)},
				&hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}},
			})
			b.SetAttributeRaw("content_b64_md5", hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenStringLit, Bytes: []byte("md5")},
				&hclwrite.Token{Type: hclsyntax.TokenOParen, Bytes: []byte{'('}},
				&hclwrite.Token{Type: hclsyntax.TokenStringLit, Bytes: []byte("filebase64")},
				&hclwrite.Token{Type: hclsyntax.TokenOParen, Bytes: []byte{'('}},
				&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}},
				&hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(relativeFile)},
				&hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}},
				&hclwrite.Token{Type: hclsyntax.TokenCParen, Bytes: []byte{')'}},
				&hclwrite.Token{Type: hclsyntax.TokenCParen, Bytes: []byte{')'}},
			})
			b.SetAttributeValue("validate_remote_file", cty.BoolVal(false))
			b.SetAttributeValue("overwrite", cty.BoolVal(true))
			b.SetAttributeValue("mkdirs", cty.BoolVal(true))
			return nil
		},
	},
	"databricks_instance_pool": {
		Service: "compute",
		Name:    instancePoolName,
		Import: func(ic *importContext, d *schema.ResourceData) error {
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       fmt.Sprintf("/instance-pools/%s", d.Id()),
				Name:     instancePoolName(d),
			})
			return nil
		},
	},
	"databricks_instance_profile": {
		Service: "access",
		Name: func(d *schema.ResourceData) string {
			arn := d.Get("instance_profile_arn").(string)
			splits := strings.Split(arn, "/")
			return splits[len(splits)-1]
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			ir := ic.Importables[r.Resource]
			state := ic.InstanceState(r)
			ipa := "instance_profile_arn"
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			ic.reference(ir, []string{ipa}, ipa, state.Attributes[ipa], b)
			b.SetAttributeValue("skip_validation", cty.BoolVal(false))
			return nil
		},
	},
	"databricks_group_instance_profile": {
		Service: "access",
		Name: func(d *schema.ResourceData) string {
			return d.Id()
		},
		Depends: []reference{
			{
				Path:     "group_id",
				Resource: "databricks_group",
			},
			{
				Path:     "instance_profile_id",
				Resource: "databricks_instance_profile",
			},
		},
	},
	"databricks_cluster": {
		Service: "compute",
		Name: func(d *schema.ResourceData) string {
			name := d.Get("cluster_name").(string)
			if name == "" {
				return strings.Split(d.Id(), "-")[2]
			}
			return name
		},
		Depends: []reference{
			{
				Path:     "aws_attributes.instance_profile_arn",
				Resource: "databricks_instance_profile",
			},
			{
				Path:     "instance_pool_id",
				Resource: "databricks_instance_pool",
			},
			{
				Path:     "init_scripts.dbfs.destination",
				Resource: "databricks_dbfs_file",
			},
			{
				Path:     "library.jar",
				Resource: "databricks_dbfs_file",
			},
			{
				Path:     "library.whl",
				Resource: "databricks_dbfs_file",
			},
			{
				Path:     "library.egg",
				Resource: "databricks_dbfs_file",
			},
		},
		List: func(ic *importContext) error {
			clusters, err := compute.NewClustersAPI(ic.Client).List()
			if err != nil {
				return err
			}
			lastActiveMs := ic.lastActiveDays * 24 * 60 * 60 * 1000
			for offset, c := range clusters {
				if c.ClusterSource == "JOB" {
					log.Printf("[INFO] Skipping job cluster %s", c.ClusterID)
					continue
				}
				if c.ClusterID != "0328-184258-tempi2" {
					continue
				}
				if c.LastActivityTime < time.Now().Unix()-lastActiveMs {
					log.Printf("[INFO] Older inactive cluster %s", c.ClusterID)
					//continue
				}
				ic.Emit(&resource{
					Resource: "databricks_cluster",
					ID:       c.ClusterID,
				})
				log.Printf("[INFO] Scanned %d of %d clusters", offset, len(clusters))
			}
			return nil
		},
		Import: func(ic *importContext, d *schema.ResourceData) error {
			var c compute.Cluster
			s := ic.Resources["databricks_cluster"].Schema
			if err := internal.DataToStructPointer(d, s, &c); err != nil {
				return err
			}
			if err := ic.importCluster(&c); err != nil {
				return err
			}
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       fmt.Sprintf("/clusters/%s", d.Id()),
				Name:     d.Get("cluster_name").(string),
			})
			return ic.importLibraries(d, s)
		},
	},
	"databricks_job": {
		Service: "jobs",
		Name: func(d *schema.ResourceData) string {
			return d.Get("name").(string)
		},
		Import: func(ic *importContext, d *schema.ResourceData) error {
			var job compute.JobSettings
			s := ic.Resources["databricks_job"].Schema
			if err := internal.DataToStructPointer(d, s, &job); err != nil {
				return err
			}
			if err := ic.importCluster(job.NewCluster); err != nil {
				return err
			}
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       fmt.Sprintf("/jobs/%s", d.Id()),
				Name:     d.Get("name").(string),
			})
			return ic.importLibraries(d, s)
		},
		// TODO: add jobs listing
	},
	"databricks_cluster_policy": {
		Service: "compute",
		Name: func(d *schema.ResourceData) string {
			return d.Get("name").(string)
		},
		Import: func(ic *importContext, d *schema.ResourceData) error {
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       fmt.Sprintf("/cluster-policies/%s", d.Id()),
				Name:     d.Get("name").(string),
			})
			var definition map[string]map[string]interface{}
			err := json.Unmarshal([]byte(d.Get("definition").(string)), &definition)
			if err != nil {
				return err
			}
			for k, policy := range definition {
				value, vok := policy["value"]
				defaultValue, dok := policy["defaultValue"]
				if !vok || !dok {
					continue
				}
				if "aws_attributes.instance_profile_arn" == k {
					ic.Emit(&resource{
						Resource: "databricks_instance_profile",
						ID:       fmt.Sprintf("%s%s", value, defaultValue),
					})
				}
				if "instance_pool_id" == k {
					ic.Emit(&resource{
						Resource: "databricks_instance_pool",
						ID:       fmt.Sprintf("%s%s", value, defaultValue),
					})
				}
			}
			return nil
		},
		// TODO: special formatting required, where JSON is written line by line
		// so that we're able to do the references
	},
	"databricks_group": {
		Service: "identity",
		Name: func(d *schema.ResourceData) string {
			return d.Get("display_name").(string)
		},
		List: func(ic *importContext) error {
			if err := ic.cacheGroups(); err != nil {
				return err
			}
			for _, g := range ic.allGroups {
				ic.Emit(&resource{
					Resource: "databricks_group",
					ID:       g.ID,
				})
			}
			return nil
		},
		Search: func(ic *importContext, r *resource) error {
			if err := ic.cacheGroups(); err != nil {
				return err
			}
			for _, g := range ic.allGroups {
				// TODO: special handling for "users" and "admins"
				// if g.DisplayName == "users" {
				// 	continue
				// }
				if g.DisplayName == r.Value && r.Attribute == "display_name" {
					r.ID = g.ID
					return nil
				}
			}
			return nil
		},
		Import: func(ic *importContext, d *schema.ResourceData) error {
			if err := ic.cacheGroups(); err != nil {
				return err
			}
			for _, g := range ic.allGroups {
				if d.Id() != g.ID {
					continue
				}
				for _, instanceProfile := range g.Roles {
					ic.Emit(&resource{
						Resource: "databricks_instance_profile",
						ID:       instanceProfile.Value,
					})
					ic.Emit(&resource{
						Resource: "databricks_group_instance_profile",
						ID:       fmt.Sprintf("%s|%s", g.ID, instanceProfile.Value),
					})
				}
				if g.DisplayName == "users" {
					log.Printf("[INFO] Skipping import of entire user directory ...")
					continue
				}
				if len(g.Members) > 10 {
					log.Printf("[INFO] Importing %d members of %s",
						len(g.Members), g.DisplayName)
				}
				for i, x := range g.Members {
					if strings.Contains(x.Ref, "Users/") {
						ic.Emit(&resource{
							Resource: "databricks_user",
							ID:       x.Value,
						})
					} else {
						ic.Emit(&resource{
							Resource: "databricks_group",
							ID:       x.Value,
						})
					}
					ic.Emit(&resource{
						Resource: "databricks_group_member",
						ID:       fmt.Sprintf("%s|%s", g.ID, x.Value),
					})
					if len(g.Members) > 10 {
						log.Printf("[INFO] Imported %d of %d members of %s",
							i, len(g.Members), g.DisplayName)
					}
				}
			}
			return nil
		},
	},
	"databricks_group_member": {
		Service: "identity",
		Depends: []reference{
			{Path: "group_id", Resource: "databricks_group"},
			{Path: "member_id", Resource: "databricks_user"},
			{Path: "member_id", Resource: "databricks_group"},
		},
	},
	"databricks_user": {
		Service: "identity",
		Name: func(d *schema.ResourceData) string {
			s := strings.Split(d.Get("user_name").(string), "@")
			return s[0]
		},
		Search: func(ic *importContext, r *resource) error {
			u, err := ic.findUserByName(r.Value)
			if err != nil {
				return err
			}
			r.ID = u.ID
			return nil
		},
		Import: func(ic *importContext, d *schema.ResourceData) error {
			u, err := ic.findUserByName(d.Get("user_name").(string))
			if err != nil {
				return err
			}
			for _, g := range u.Groups {
				ic.Emit(&resource{
					Resource: "databricks_group",
					ID:       g.Value,
				})
				ic.Emit(&resource{
					Resource: "databricks_group_member",
					ID:       fmt.Sprintf("%s|%s", g.Value, u.ID),
				})
			}
			return nil
		},
	},
	"databricks_permissions": {
		Service: "access",
		Name: func(d *schema.ResourceData) string {
			s := strings.Split(d.Id(), "/")
			return s[len(s)-1]
		},
		Depends: []reference{
			{
				Path:     "access_control.group_name",
				Resource: "databricks_group",
				Match:    "display_name",
			},
			{
				Path:     "access_control.user_name",
				Resource: "databricks_user",
				Match:    "user_name",
			},
			{
				Path:     "cluster_id",
				Resource: "databricks_cluster",
			},
			{
				Path:     "cluster_policy_id",
				Resource: "databricks_cluster_policy",
			},
			{
				Path:     "instance_pool_id",
				Resource: "databricks_instance_pool",
			},
		},
		Import: func(ic *importContext, d *schema.ResourceData) error {
			var permissions access.PermissionsEntity
			s := ic.Resources["databricks_permissions"].Schema
			err := internal.DataToStructPointer(d, s, &permissions)
			if err != nil {
				return err
			}
			for _, ac := range permissions.AccessControlList {
				ic.Emit(&resource{
					Resource:  "databricks_user",
					Attribute: "user_name",
					Value:     ac.UserName,
				})
				ic.Emit(&resource{
					Resource:  "databricks_group",
					Attribute: "display_name",
					Value:     ac.GroupName,
				})
			}
			return nil
		},
	},
}

func (ic *importContext) importCluster(c *compute.Cluster) error {
	if c == nil {
		return nil
	}
	for _, is := range c.InitScripts {
		if is.Dbfs != nil {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       is.Dbfs.Destination,
			})
		}
	}
	if c.AwsAttributes != nil {
		ic.Emit(&resource{
			Resource: "databricks_instance_profile",
			ID:       c.AwsAttributes.InstanceProfileArn,
		})
	}
	if c.InstancePoolID != "" {
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       c.InstancePoolID,
		})
	}
	if c.PolicyID != "" {
		ic.Emit(&resource{
			Resource: "databricks_cluster_policy",
			ID:       c.PolicyID,
		})
	}
	return nil
}

func (ic *importContext) importLibraries(d *schema.ResourceData, s map[string]*schema.Schema) error {
	var cll compute.ClusterLibraryList
	err := internal.DataToStructPointer(d, s, &cll)
	if err != nil {
		return err
	}
	for _, lib := range cll.Libraries {
		if strings.HasPrefix(lib.Whl, "dbfs:") {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       lib.Whl,
			})
		}
		if strings.HasPrefix(lib.Jar, "dbfs:") {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       lib.Jar,
			})
		}
		if strings.HasPrefix(lib.Egg, "dbfs:") {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       lib.Egg,
			})
		}
	}
	return nil
}

func (ic *importContext) cacheGroups() error {
	if len(ic.allGroups) == 0 {
		log.Printf("[INFO] Caching groups in memory ...")
		groupsAPI := identity.NewGroupsAPI(ic.Client)
		g, err := groupsAPI.Filter("")
		if err != nil {
			return err
		}
		ic.allGroups = g.Resources
		log.Printf("[INFO] Cached %d groups", len(ic.allGroups))
	}
	return nil
}

func (ic *importContext) cacheUsers() error {
	if len(ic.allUsers) == 0 {
		// workspace has at least one user, always.
		log.Printf("[INFO] Fetching users into in-memory cache")
		usersAPI := identity.NewUsersAPI(ic.Client)
		users, err := usersAPI.Filter("active eq true")
		if err != nil {
			return err
		}
		ic.allUsers = users
		log.Printf("[INFO] Cached %d users", len(users))
	}
	return nil
}

func (ic *importContext) findUserByName(name string) (u identity.ScimUser, err error) {
	a := identity.NewUsersAPI(ic.Client)
	users, err := a.Filter(fmt.Sprintf("userName eq '%s'", name))
	if err != nil {
		return
	}
	if len(users) == 0 {
		err = fmt.Errorf("User %s not found", name)
		return
	}
	u = users[0]
	return
	// if err := ic.cacheUsers(); err != nil {
	// 	return
	// }
	// for _, _u := range ic.allUsers {
	// 	if _u.UserName == name {
	// 		u = _u
	// 		return
	// 	}
	// }
	// err = fmt.Errorf("User %s not found", name)
	// return
}

func (ic *importContext) regexFix(s string, fixes []regexFix) string {
	for _, x := range fixes {
		s = x.Regex.ReplaceAllString(s, x.Replacement)
	}
	return s
}

func (ic *importContext) ResourceName(r *resource, d *schema.ResourceData) string {
	name := r.Name
	if name == "" && ic.Importables[r.Resource].Name != nil {
		name = ic.Importables[r.Resource].Name(d)
	}
	if name == "" {
		name = d.Id()
	}
	name = strings.ToLower(name)
	name = ic.regexFix(name, ic.nameFixes)
	// this is either numeric id or all-non-ascii
	if regexp.MustCompile(`^\d`).MatchString(name) || "" == name {
		if "" == name {
			name = d.Id()
		}
		name = fmt.Sprintf("r%x", md5.Sum([]byte(name)))[0:12]
	}
	return name
}

func (ic *importContext) Emit(r *resource) {
	// TODO: change into channels, if stack trace depth issues would surface
	_, v := r.MatchPair()
	if v == "" {
		log.Printf("[INFO] %s has got empty identifier", r)
		return
	}
	if ic.Has(r) {
		log.Printf("[DEBUG] %s already imported", r)
		return
	}
	ic.importing[r.String()] = true
	pr, ok := ic.Resources[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available in provider", r)
		return
	}
	ir, ok := ic.Importables[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available for import", r)
		return
	}
	if !strings.Contains(ic.services, ir.Service) {
		log.Printf("[DEBUG] %s (%s service) is not part of the import",
			r.Resource, ir.Service)
		return
	}
	if r.ID == "" {
		if ir.Search == nil {
			log.Printf("[ERROR] Searching %s is not available", r)
			return
		}
		err := ir.Search(ic, r)
		if err != nil {
			log.Printf("[ERROR] Cannot search for a resource %s: %v", err, r)
			return
		}
		if r.ID == "" {
			log.Printf("[INFO] Cannot find %s", r)
			return
		}
	}
	// empty data with resource schema
	d := pr.Data(&terraform.InstanceState{
		Attributes: map[string]string{},
		ID:         r.ID,
	})
	d.MarkNewResource()

	if pr.Read != nil {
		err := pr.Read(d, ic.Client)
		if err != nil {
			return
		}
	} else {
		dia := pr.ReadContext(context.Background(), d, ic.Client)
		if dia != nil {
			log.Printf("[ERROR] Error reading %s#%s: %v", r.Resource, r.ID, dia)
			return
		}
	}
	r.Name = ic.ResourceName(r, d)
	if ir.Import != nil {
		err := ir.Import(ic, d)
		if err != nil {
			log.Printf("[ERROR] Failed custom import of %s: %s", r, err)
			return
		}
	}
	state := d.State()
	if state == nil {
		s := strings.Split(r.ID, "|")
		if len(s) == 2 {
			g, _ := identity.NewGroupsAPI(ic.Client).Read(s[0])
			u, _ := identity.NewUsersAPI(ic.Client).Read(s[1])
			log.Printf("[INFO] %v ====> %v and %v", r, g, u)
		}

		log.Printf("[ERROR] state is nil for %s", r)
		return
	}
	ic.Add(r, state.Attributes)
}

func newImportContext(c *common.DatabricksClient) *importContext {
	p := provider.DatabricksProvider()
	return &importContext{
		Module:      "",
		Client:      c,
		State:       stateApproximation{},
		Importables: resourcesMap,
		Resources:   p.ResourcesMap,
		Files:       map[string]*hclwrite.File{},
		Scope:       []*resource{},
		Directory:   "/tmp/importer",
		importing:   map[string]bool{},
		nameFixes: []regexFix{
			{regexp.MustCompile(`[0-9a-f]{8}[_-][0-9a-f]{4}[_-][0-9a-f]{4}` +
				`[_-][0-9a-f]{4}[_-][0-9a-f]{12}[_-]`), ""},
			{regexp.MustCompile(`[_-][0-9]+[\._-][0-9]+[\._-].*\.(whl|jar|egg)`), "_$1"},
			{regexp.MustCompile(`[-\s\.\|]`), "_"},
			{regexp.MustCompile(`\W+`), ""},
			{regexp.MustCompile(`[_]{2,}`), "_"},
		},
		hclFixes: []regexFix{
			{regexp.MustCompile(`\{ "`), "{\n\t\t\""},
			{regexp.MustCompile(`", "`), "\",\n\t\t\""},
			{regexp.MustCompile(`" \}`), "\"\t\n}"},
		},
		allUsers: []identity.ScimUser{},
	}
}

func (ic *importContext) Run() error {
	if len(ic.services) == 0 {
		return fmt.Errorf("No services to import")
	}
	log.Printf("[INFO] Importing %s module into %s directory Databricks resources of %s services",
		ic.Module, ic.Directory, ic.services)
	for resourceName, ir := range ic.Importables {
		if ir.List == nil {
			continue
		}
		if !strings.Contains(ic.services, ir.Service) {
			log.Printf("[DEBUG] %s (%s service) is not part of the import",
				resourceName, ir.Service)
			continue
		}
		if err := ir.List(ic); err != nil {
			return err
		}
	}
	if len(ic.Scope) == 0 {
		return fmt.Errorf("No resources to import")
	}
	sh, err := os.Create(fmt.Sprintf("%s/import.sh", ic.Directory))
	if err != nil {
		return err
	}
	defer sh.Close()

	n := len(ic.Scope)
	swap := reflect.Swapper(ic.Scope)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}

	for _, r := range ic.Scope {
		ir := ic.Importables[r.Resource]
		f, ok := ic.Files[ir.Service]
		if !ok {
			f = hclwrite.NewEmptyFile()
			ic.Files[ir.Service] = f
		}
		body := f.Body()
		if ir.Body != nil {
			err := ir.Body(ic, body, r)
			if err != nil {
				return err
			}
		} else {
			pr := ic.Resources[r.Resource]
			resourceBlock := body.AppendNewBlock("resource", []string{r.Resource, r.Name})
			err := ic.dataToHcl(ir, []string{}, pr, pr.Data(ic.InstanceState(r)), resourceBlock.Body())
			if err != nil {
				return err
			}
		}

		// nolint
		//sh.WriteString(r.ImportCommand(ic) + "\n")
	}
	for service, f := range ic.Files {
		formatted := hclwrite.Format(f.Bytes())
		// fix some formatting in a hacky way instead of writing 100 lines
		// of HCL AST writer code
		formatted = []byte(ic.regexFix(string(formatted), ic.hclFixes))
		log.Printf("[INFO] %s", formatted)
		tf, err := os.Create(fmt.Sprintf("%s/%s.tf", ic.Directory, service))
		if err != nil {
			return err
		}
		defer tf.Close()
		_, err = tf.Write(formatted)
		if err != nil {
			return err
		}
	}
	cmd := exec.CommandContext(context.Background(), "terraform", "fmt")
	cmd.Dir = ic.Directory
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

var dependsRe = regexp.MustCompile(`(\.[\d]+)`)

func (ic *importContext) reference(i importable, path []string, key, value string, body *hclwrite.Body) {
	match := dependsRe.ReplaceAllString(strings.Join(append(path, key), "."), "")
	for _, d := range i.Depends {
		if d.Path != match {
			continue
		}
		attr := "id"
		if d.Match != "" {
			attr = d.Match
		}
		if d.Pick == "" {
			d.Pick = "id"
		}
		traversal := ic.Find(&resource{
			Resource:  d.Resource,
			Attribute: attr,
			Value:     value,
		}, d.Pick)
		if traversal != nil {
			body.SetAttributeTraversal(key, traversal)
			return
		}
	}
	body.SetAttributeValue(key, cty.StringVal(value))
}

func (ic *importContext) dataToHcl(i importable, path []string,
	pr *schema.Resource, d *schema.ResourceData, body *hclwrite.Body) error {
	for a, as := range pr.Schema {
		if as.Computed {
			continue
		}
		raw, ok := d.GetOk(strings.Join(append(path, a), "."))
		if !ok {
			continue
		}
		switch as.Type {
		case schema.TypeString:
			ic.reference(i, path, a, raw.(string), body)
		case schema.TypeBool:
			body.SetAttributeValue(a, cty.BoolVal(raw.(bool)))
		case schema.TypeInt:
			switch iv := raw.(type) {
			case int:
				body.SetAttributeValue(a, cty.NumberIntVal(int64(iv)))
			case int32:
				body.SetAttributeValue(a, cty.NumberIntVal(int64(iv)))
			case int64:
				body.SetAttributeValue(a, cty.NumberIntVal(iv))
			}
		case schema.TypeFloat:
			body.SetAttributeValue(a, cty.NumberFloatVal(raw.(float64)))
		case schema.TypeMap:
			// mapBlock := body.AppendNewBlock(a, []string{}).Body()
			ov := map[string]cty.Value{}
			for key, iv := range raw.(map[string]interface{}) {
				v := cty.StringVal(fmt.Sprintf("%v", iv))
				ov[key] = v
				// mapBlock.SetAttributeValue(key, v)
			}
			body.SetAttributeValue(a, cty.ObjectVal(ov))
		case schema.TypeSet:
			if rawSet, ok := raw.(*schema.Set); ok {
				rawList := rawSet.List()
				err := ic.readListFromData(i, append(path, a), d, rawList, body, as, func(i int) string {
					return strconv.Itoa(rawSet.F(rawList[i]))
				})
				if err != nil {
					return err
				}
			}
		case schema.TypeList:
			if rawList, ok := raw.([]interface{}); ok {
				err := ic.readListFromData(i, append(path, a), d, rawList, body, as, strconv.Itoa)
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("Unsupported schema type: %v", path)
		}
	}
	return nil
}

func (ic *importContext) readListFromData(i importable, path []string, d *schema.ResourceData,
	rawList []interface{}, body *hclwrite.Body, as *schema.Schema,
	offsetConverter func(i int) string) error {
	if len(rawList) == 0 {
		return nil
	}
	name := path[len(path)-1]
	switch elem := as.Elem.(type) {
	case *schema.Resource:
		if as.MaxItems == 1 {
			nestedPath := append(path, offsetConverter(0))
			confBlock := body.AppendNewBlock(name, []string{})
			return ic.dataToHcl(i, nestedPath, elem, d, confBlock.Body())
		}
		for offset := range rawList {
			confBlock := body.AppendNewBlock(name, []string{})
			nestedPath := append(path, offsetConverter(offset))
			err := ic.dataToHcl(i, nestedPath, elem, d, confBlock.Body())
			if err != nil {
				return err
			}
		}
	case *schema.Schema:
		primitiveValues := []cty.Value{}
		for _, raw := range rawList {
			switch x := raw.(type) {
			case string:
				primitiveValues = append(primitiveValues, cty.StringVal(x))
			case int:
				primitiveValues = append(primitiveValues, cty.NumberIntVal(int64(x)))
			default:
				return fmt.Errorf("Unsupported primitive list: %#v", path)
			}
		}
		body.SetAttributeValue(name, cty.ListVal(primitiveValues))
	}
	return nil
}
