package importer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/storage"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
)

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
		Name: func(d *schema.ResourceData) string {
			raw, ok := d.GetOk("instance_pool_name")
			if !ok {
				return strings.Split(d.Id(), "-")[2]
			}
			name := raw.(string)
			if name == "" {
				return strings.Split(d.Id(), "-")[2]
			}
			return name
		},
		Import: func(ic *importContext, r *resource) error {
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       fmt.Sprintf("/instance-pools/%s", r.ID),
				Name:     ic.Importables["databricks_instance_pool"].Name(r.Data),
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
			ipa := "instance_profile_arn"
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			b.SetAttributeRaw(ipa, ic.reference(ic.Importables[r.Resource],
				[]string{ipa}, r.Data.Get(ipa).(string)))
			b.SetAttributeValue("skip_validation", cty.BoolVal(false))
			return nil
		},
	},
	"databricks_group_instance_profile": {
		Service: "access",
		Depends: []reference{
			{Path: "group_id", Resource: "databricks_group"},
			{Path: "instance_profile_id", Resource: "databricks_instance_profile"},
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
			{Path: "aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "init_scripts.dbfs.destination", Resource: "databricks_dbfs_file"},
			{Path: "library.jar", Resource: "databricks_dbfs_file"},
			{Path: "library.whl", Resource: "databricks_dbfs_file"},
			{Path: "library.egg", Resource: "databricks_dbfs_file"},
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
				if !ic.MatchesName(c.ClusterName) {
					continue
				}
				if c.LastActivityTime < time.Now().Unix()-lastActiveMs {
					log.Printf("[INFO] Older inactive cluster %s", c.ClusterID)
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_cluster",
					ID:       c.ClusterID,
				})
				log.Printf("[INFO] Scanned %d of %d clusters", offset+1, len(clusters))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			var c compute.Cluster
			s := ic.Resources["databricks_cluster"].Schema
			if err := internal.DataToStructPointer(r.Data, s, &c); err != nil {
				return err
			}
			if err := ic.importCluster(&c); err != nil {
				return err
			}
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       fmt.Sprintf("/clusters/%s", r.ID),
				Name:     r.Data.Get("cluster_name").(string),
			})
			return ic.importLibraries(r.Data, s)
		},
	},
	"databricks_job": {
		Service: "jobs",
		Name: func(d *schema.ResourceData) string {
			return d.Get("name").(string)
		},
		Depends: []reference{
			{Path: "email_notifications.on_failure", Resource: "databricks_user", Match: "user_name"},
			{Path: "email_notifications.on_success", Resource: "databricks_user", Match: "user_name"},
			{Path: "email_notifications.on_start", Resource: "databricks_user", Match: "user_name"},
			{Path: "new_cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "new_cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file"},
			{Path: "new_cluster.instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "existing_cluster_id", Resource: "databricks_cluster"},
			{Path: "library.jar", Resource: "databricks_dbfs_file"},
			{Path: "library.whl", Resource: "databricks_dbfs_file"},
			{Path: "library.egg", Resource: "databricks_dbfs_file"},
			{Path: "spark_python_task.python_file", Resource: "databricks_dbfs_file"},
			{Path: "spark_python_task.parameters", Resource: "databricks_dbfs_file"},
			{Path: "spark_jar_task.jar_uri", Resource: "databricks_dbfs_file"},
		},
		Import: func(ic *importContext, r *resource) error {
			var job compute.JobSettings
			s := ic.Resources["databricks_job"].Schema
			if err := internal.DataToStructPointer(r.Data, s, &job); err != nil {
				return err
			}
			if err := ic.importCluster(job.NewCluster); err != nil {
				return err
			}
			ic.Emit(&resource{
				Resource: "databricks_cluster",
				ID:       job.ExistingClusterID,
			})
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       fmt.Sprintf("/jobs/%s", r.ID),
				Name:     r.Data.Get("name").(string),
			})
			if job.SparkPythonTask != nil {
				ic.emitIfDbfsFile(job.SparkPythonTask.PythonFile)
				for _, p := range job.SparkPythonTask.Parameters {
					ic.emitIfDbfsFile(p)
				}
			}
			if job.SparkJarTask != nil {
				jarURI := job.SparkJarTask.JarURI
				if jarURI != "" {
					if libs, ok := r.Data.Get("library").(*schema.Set); ok {
						// nolint remove legacy jar uri support
						r.Data.Set("spark_jar_task", []interface{}{
							map[string]interface{}{
								"main_class_name": job.SparkJarTask.MainClassName,
								"parameters":      job.SparkJarTask.Parameters,
							},
						})
						// if variable doesn't contain a sad face, it's a job jar
						if !strings.Contains(jarURI, ":/") {
							jarURI = fmt.Sprintf("dbfs:/FileStore/job-jars/%s", jarURI)
						}
						ic.emitIfDbfsFile(jarURI)
						libs.Add(map[string]interface{}{
							"jar": jarURI,
						})
						// nolint
						r.Data.Set("library", libs)
					}
				}
			}
			return ic.importLibraries(r.Data, s)
		},
		List: func(ic *importContext) error {
			a := compute.NewJobsAPI(ic.Client)
			nowSeconds := time.Now().Unix()
			starterAfter := (nowSeconds - (ic.lastActiveDays * 24 * 60 * 60)) * 1000
			if l, err := a.List(); err == nil {
				i := 0
				for _, job := range l.Jobs {
					if !ic.MatchesName(job.Settings.Name) {
						continue
					}
					if ic.lastActiveDays != 3650 {
						rl, err := a.RunsList(compute.JobRunsListRequest{
							JobID:         job.JobID,
							CompletedOnly: true,
							Limit:         1,
						})
						if err != nil {
							log.Printf("[WARN] Failed to get runs: %s", err)
							continue
						}
						if len(rl.Runs) == 0 {
							log.Printf("[INFO] Job %#v (%d) did never run. Skipping", job.Settings.Name, job.JobID)
							continue
						}
						if rl.Runs[0].StartTime < starterAfter {
							log.Printf("[INFO] Job %#v (%d) didn't run for %d days. Skipping",
								job.Settings.Name, job.JobID,
								(nowSeconds*1000-rl.Runs[0].StartTime)/24*60*60/1000)
							continue
						}
					}
					ic.Emit(&resource{
						Resource: "databricks_job",
						ID:       job.ID(),
					})
					i++
					log.Printf("[INFO] Imported %d of total %d jobs", i, len(l.Jobs))
				}
			}
			return nil
		},
	},
	"databricks_cluster_policy": {
		Service: "compute",
		Name: func(d *schema.ResourceData) string {
			return d.Get("name").(string)
		},
		Import: func(ic *importContext, r *resource) error {
			ic.Emit(&resource{
				Resource: "databricks_permissions",
				ID:       fmt.Sprintf("/cluster-policies/%s", r.ID),
				Name:     r.Data.Get("name").(string),
			})
			var definition map[string]map[string]interface{}
			err := json.Unmarshal([]byte(r.Data.Get("definition").(string)), &definition)
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
		Service: "groups",
		Name: func(d *schema.ResourceData) string {
			return d.Get("display_name").(string)
		},
		List: func(ic *importContext) error {
			if err := ic.cacheGroups(); err != nil {
				return err
			}
			for _, g := range ic.allGroups {
				if !ic.MatchesName(g.DisplayName) {
					continue
				}
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
				if g.DisplayName == r.Value && r.Attribute == "display_name" {
					r.ID = g.ID
					return nil
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			if r.Name == "admins" || r.Name == "users" {
				// admins & users are to be imported through "data block"
				r.Mode = "data"
				r.Data.State().Set(&terraform.InstanceState{
					ID: r.ID,
					Attributes: map[string]string{
						"display_name": r.Name,
					},
				})
			}
			if err := ic.cacheGroups(); err != nil {
				return err
			}
			for _, g := range ic.allGroups {
				if r.ID != g.ID {
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
					// TODO: make flag
					log.Printf("[INFO] Skipping import of entire user directory ...")
					continue
				}
				if len(g.Members) > 10 {
					log.Printf("[INFO] Importing %d members of %s",
						len(g.Members), g.DisplayName)
				}
				for _, parent := range g.Parents {
					ic.Emit(&resource{
						Resource: "databricks_group",
						ID:       parent.Value,
					})
					if "direct" == parent.Type {
						ic.Emit(&resource{
							Resource: "databricks_group_member",
							ID:       fmt.Sprintf("%s|%s", parent.Value, g.ID),
							Name:     fmt.Sprintf("%s_%s", parent.Display, g.DisplayName),
						})
					}
				}
				for i, x := range g.Members {
					if strings.Contains(x.Ref, "Users/") {
						ic.Emit(&resource{
							Resource: "databricks_user",
							ID:       x.Value,
						})
					}
					if strings.Contains(x.Ref, "Groups/") {
						ic.Emit(&resource{
							Resource: "databricks_group",
							ID:       x.Value,
						})
						ic.Emit(&resource{
							Resource: "databricks_group_member",
							ID:       fmt.Sprintf("%s|%s", g.ID, x.Value),
							Name:     fmt.Sprintf("%s_%s", g.DisplayName, x.Display),
						})
					}
					if len(g.Members) > 10 {
						log.Printf("[INFO] Imported %d of %d members of %s",
							i, len(g.Members), g.DisplayName)
					}
				}
			}
			return nil
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			blockType := "resource"
			if r.Mode == "data" {
				blockType = r.Mode
			}
			resourceBlock := body.AppendNewBlock(blockType, []string{r.Resource, r.Name})
			return ic.dataToHcl(ic.Importables[r.Resource],
				[]string{}, ic.Resources[r.Resource], r.Data, resourceBlock.Body())
		},
	},
	"databricks_group_member": {
		Service: "groups",
		Depends: []reference{
			{Path: "group_id", Resource: "databricks_group"},
			{Path: "member_id", Resource: "databricks_user"},
			{Path: "member_id", Resource: "databricks_group"},
		},
	},
	"databricks_user": {
		Service: "users",
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
		Import: func(ic *importContext, r *resource) error {
			u, err := ic.findUserByName(r.Data.Get("user_name").(string))
			if err != nil {
				return err
			}
			for _, g := range u.Groups {
				if "direct" != g.Type {
					continue
				}
				ic.Emit(&resource{
					Resource: "databricks_group",
					ID:       g.Value,
				})
				ic.Emit(&resource{
					Resource: "databricks_group_member",
					ID:       fmt.Sprintf("%s|%s", g.Value, u.ID),
					Name:     fmt.Sprintf("%s_%s", g.Display, u.DisplayName),
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
			{Path: "job_id", Resource: "databricks_job"},
			{Path: "cluster_id", Resource: "databricks_cluster"},
			{Path: "instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "cluster_policy_id", Resource: "databricks_cluster_policy"},
			{Path: "access_control.user_name", Resource: "databricks_user", Match: "user_name"},
			{Path: "access_control.group_name", Resource: "databricks_group", Match: "display_name"},
		},
		Import: func(ic *importContext, r *resource) error {
			var permissions access.PermissionsEntity
			s := ic.Resources["databricks_permissions"].Schema
			err := internal.DataToStructPointer(r.Data, s, &permissions)
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
	"databricks_secret_scope": {
		Service: "secrets",
		Name: func(d *schema.ResourceData) string {
			return d.Get("name").(string)
		},
		List: func(ic *importContext) error {
			ssAPI := access.NewSecretScopesAPI(ic.Client)
			if scopes, err := ssAPI.List(); err == nil {
				for i, scope := range scopes {
					if !ic.MatchesName(scope.Name) {
						continue
					}
					ic.Emit(&resource{
						Resource: "databricks_secret_scope",
						ID:       scope.Name,
					})
					log.Printf("[INFO] Imported %d of %d secret scopes",
						i, len(scopes))
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			if l, err := access.NewSecretsAPI(ic.Client).List(r.ID); err == nil {
				for _, secret := range l {
					ic.Emit(&resource{
						Resource: "databricks_secret",
						ID:       fmt.Sprintf("%s|||%s", r.ID, secret.Key),
					})
				}
			}
			if l, err := access.NewSecretAclsAPI(ic.Client).List(r.ID); err == nil {
				for _, acl := range l {
					ic.Emit(&resource{
						Resource: "databricks_secret_acl",
						ID:       fmt.Sprintf("%s|||%s", r.ID, acl.Principal),
					})
				}
			}
			return nil
		},
	},
	"databricks_secret": {
		Service: "secrets",
		Depends: []reference{
			{Path: "scope", Resource: "databricks_secret_scope"},
			{Path: "string_value", Resource: "vault_generic_secret", Match: "data"},
			{Path: "string_value", Resource: "aws_kms_secrets", Match: "plaintext"},
			{Path: "string_value", Resource: "azurerm_key_vault_secret", Match: "value"},
			{Path: "string_value", Resource: "aws_secretsmanager_secret_version", Match: "secret_string"},
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			b.SetAttributeRaw("scope", ic.reference(ic.Importables[r.Resource],
				[]string{"scope"}, r.Data.Get("scope").(string)))
			// secret data is exposed only within notebooks
			b.SetAttributeRaw("string_value", ic.variable(
				r.Name, fmt.Sprintf("Secret %s from %s scope",
					r.Data.Get("key"), r.Data.Get("scope"))))
			b.SetAttributeValue("key", cty.StringVal(r.Data.Get("key").(string)))
			return nil
		},
	},
	"databricks_secret_acl": {
		Service: "secrets",
		Depends: []reference{
			{Path: "scope", Resource: "databricks_secret_scope"},
			{Path: "principal", Resource: "databricks_group", Match: "display_name"},
			{Path: "principal", Resource: "databricks_user", Match: "user_name"},
		},
	},
}
