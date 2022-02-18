package exporter

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/databrickslabs/terraform-provider-databricks/access"
	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/jobs"
	"github.com/databrickslabs/terraform-provider-databricks/permissions"
	"github.com/databrickslabs/terraform-provider-databricks/repos"
	"github.com/databrickslabs/terraform-provider-databricks/secrets"
	"github.com/databrickslabs/terraform-provider-databricks/workspace"

	"github.com/databrickslabs/terraform-provider-databricks/storage"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
)

var (
	adlsGen2Regex           = regexp.MustCompile(`^(abfss?)://([^@]+)@([^.]+)\.(?:[^/]+)(/.*)?$`)
	adlsGen1Regex           = regexp.MustCompile(`^(adls?)://([^.]+)\.(?:[^/]+)(/.*)?$`)
	wasbsRegex              = regexp.MustCompile(`^(wasbs?)://([^@]+)@([^.]+)\.(?:[^/]+)(/.*)?$`)
	s3Regex                 = regexp.MustCompile(`^(s3a?)://([^/]+)(/.*)?$`)
	gsRegex                 = regexp.MustCompile(`^gs://([^/]+)(/.*)?$`)
	globalWorkspaceConfName = "global_workspace_conf"
)

func generateMountBody(ic *importContext, body *hclwrite.Body, r *resource) error {
	mount := ic.mountMap[r.ID]

	b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
	b.SetAttributeValue("name", cty.StringVal(strings.Replace(r.ID, "/mnt/", "", 1)))
	if res := s3Regex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("s3", nil).Body()
		block.SetAttributeValue("bucket_name", cty.StringVal(res[2]))
		if mount.InstanceProfile != "" {
			block.SetAttributeValue("instance_profile", cty.StringVal(mount.InstanceProfile))
		} else if mount.ClusterID != "" {
			b.SetAttributeValue("cluster_id", cty.StringVal(mount.ClusterID))
		}
	} else if res := gsRegex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("gs", nil).Body()
		block.SetAttributeValue("bucket_name", cty.StringVal(res[1]))
		if mount.ClusterID != "" {
			b.SetAttributeValue("cluster_id", cty.StringVal(mount.ClusterID))
		}
	} else if res := adlsGen2Regex.FindStringSubmatch(mount.URL); res != nil {
		containerName := res[2]
		storageAccountName := res[3]
		block := b.AppendNewBlock("abfs", nil).Body()
		block.SetAttributeValue("container_name", cty.StringVal(containerName))
		block.SetAttributeValue("storage_account_name", cty.StringVal(storageAccountName))
		if res[4] != "" && res[4] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[4]))
		}

		varName := ic.regexFix("_"+storageAccountName+"_"+containerName+"_abfs", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting ADLSv2 resource %s://%s@%s",
			res[1], containerName, storageAccountName)

		block.SetAttributeRaw("client_id", ic.variable(
			"client_id"+varName, "Client ID"+textStr))
		block.SetAttributeRaw("tenant_id", ic.variable(
			"tenant_id"+varName, "Tenant ID"+textStr))
		block.SetAttributeRaw("client_secret_scope", ic.variable(
			"client_secret_scope"+varName,
			"Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("client_secret_key", ic.variable(
			"client_secret_key"+varName,
			"Key in secret scope that stores app client secret"+textStr))
	} else if res := adlsGen1Regex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("adl", nil).Body()
		storageResourceName := res[2]
		block.SetAttributeValue("storage_resource_name", cty.StringVal(storageResourceName))
		if res[3] != "" && res[3] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[3]))
		}
		varName := ic.regexFix("_"+storageResourceName+"_adl", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting ADLSv1 resource %s://%s", res[1], storageResourceName)

		block.SetAttributeRaw("client_id", ic.variable("client_id"+varName, "Client ID"+textStr))
		block.SetAttributeRaw("tenant_id", ic.variable("tenant_id"+varName, "Tenant IDs"+textStr))
		block.SetAttributeRaw("client_secret_scope", ic.variable(
			"client_secret_scope"+varName, "Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("client_secret_key", ic.variable(
			"client_secret_key"+varName, "Key in secret scope that stores app client secret"+textStr))
	} else if res := wasbsRegex.FindStringSubmatch(mount.URL); res != nil {
		containerName := res[2]
		storageAccountName := res[3]
		block := b.AppendNewBlock("wasb", nil).Body()
		block.SetAttributeValue("container_name", cty.StringVal(containerName))
		block.SetAttributeValue("storage_account_name", cty.StringVal(storageAccountName))
		if res[4] != "" && res[4] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[4]))
		}
		block.SetAttributeValue("auth_type", cty.StringVal("ACCESS_KEY"))

		varName := ic.regexFix("_"+storageAccountName+"_"+containerName+"_wasb", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting WASB resource %s://%s@%s",
			res[1], containerName, storageAccountName)

		block.SetAttributeRaw("token_secret_scope", ic.variable(
			"client_secret_scope"+varName,
			"Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("token_secret_key", ic.variable(
			"client_secret_key"+varName,
			"Key in secret scope that stores app client secret"+textStr))
	} else {
		return fmt.Errorf("no matching handler for: %s", mount.URL)
	}
	body.AppendNewline()

	return nil
}

var resourcesMap map[string]importable = map[string]importable{
	"databricks_dbfs_file": {
		Service: "storage",
		Name: func(d *schema.ResourceData) string {
			fileNameMd5 := fmt.Sprintf("%x", md5.Sum([]byte(d.Id())))
			s := strings.Split(d.Id(), "/")
			name := "_" + s[len(s)-1] + "_" + fileNameMd5
			return name
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			dbfsAPI := storage.NewDbfsAPI(ic.Context, ic.Client)
			content, err := dbfsAPI.Read(r.ID)
			if err != nil {
				return err
			}
			name := ic.Importables["databricks_dbfs_file"].Name(r.Data)
			fileName, err := ic.createFile(name, content)
			log.Printf("Creating %s for %s", fileName, r)
			if err != nil {
				return err
			}
			// libraries installed with init scripts won't be exported.
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			relativeFile := fmt.Sprintf("${path.module}/%s", fileName)
			b.SetAttributeValue("path", cty.StringVal(strings.Replace(r.ID, "dbfs:", "", 1)))
			b.SetAttributeRaw("source", hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}},
				&hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(relativeFile)},
				&hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}},
			})
			return nil
		},
	},
	"databricks_instance_pool": {
		Service: "compute",
		Name: func(d *schema.ResourceData) string {
			raw, ok := d.GetOk("instance_pool_name")
			if !ok || raw.(string) == "" {
				return strings.Split(d.Id(), "-")[2]
			}
			return raw.(string)
		},
		Import: func(ic *importContext, r *resource) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/instance-pools/%s", r.ID),
					Name:     "inst_pool_" + ic.Importables["databricks_instance_pool"].Name(r.Data),
				})
			}
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
			return fmt.Sprintf("%s_%s", name, d.Id())
		},
		Depends: []reference{
			{Path: "aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "init_scripts.dbfs.destination", Resource: "databricks_dbfs_file"},
			{Path: "library.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "policy_id", Resource: "databricks_cluster_policy"},
		},
		List: func(ic *importContext) error {
			clusters, err := clusters.NewClustersAPI(ic.Context, ic.Client).List()
			if err != nil {
				return err
			}
			lastActiveMs := ic.lastActiveDays * 24 * 60 * 60 * 1000
			for offset, c := range clusters {
				if c.ClusterSource == "JOB" {
					log.Printf("[INFO] Skipping job cluster %s", c.ClusterID)
					continue
				}
				if strings.HasPrefix(c.ClusterName, "terraform-") {
					log.Printf("[INFO] Skipping terraform-specific cluster %s", c.ClusterName)
					continue
				}
				if !ic.MatchesName(c.ClusterName) {
					log.Printf("[INFO] Skipping %s because it doesn't match %s", c.ClusterName, ic.match)
					continue
				}
				if c.LastActivityTime < time.Now().Unix()-lastActiveMs {
					log.Printf("[INFO] Older inactive cluster %s", c.ClusterName)
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
			var c clusters.Cluster
			s := ic.Resources["databricks_cluster"].Schema
			common.DataToStructPointer(r.Data, s, &c)
			ic.importCluster(&c)
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/clusters/%s", r.ID),
					Name:     "cluster_" + ic.Importables["databricks_cluster"].Name(r.Data),
				})
			}
			return ic.importLibraries(r.Data, s)
		},
	},
	"databricks_job": {
		Service: "jobs",
		Name: func(d *schema.ResourceData) string {
			return fmt.Sprintf("%s_%s", d.Get("name").(string), d.Id())
		},
		Depends: []reference{
			{Path: "email_notifications.on_failure", Resource: "databricks_user", Match: "user_name"},
			{Path: "email_notifications.on_success", Resource: "databricks_user", Match: "user_name"},
			{Path: "email_notifications.on_start", Resource: "databricks_user", Match: "user_name"},
			{Path: "new_cluster.aws_attributes.instance_profile_arn", Resource: "databricks_instance_profile"},
			{Path: "new_cluster.init_scripts.dbfs.destination", Resource: "databricks_dbfs_file"},
			{Path: "new_cluster.instance_pool_id", Resource: "databricks_instance_pool"},
			{Path: "existing_cluster_id", Resource: "databricks_cluster"},
			{Path: "library.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "library.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "spark_python_task.python_file", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "spark_python_task.parameters", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
			{Path: "spark_jar_task.jar_uri", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
		},
		Import: func(ic *importContext, r *resource) error {
			var job jobs.JobSettings
			s := ic.Resources["databricks_job"].Schema
			common.DataToStructPointer(r.Data, s, &job)
			ic.importCluster(job.NewCluster)
			ic.Emit(&resource{
				Resource: "databricks_cluster",
				ID:       job.ExistingClusterID,
			})
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/jobs/%s", r.ID),
					Name:     "job_" + ic.Importables["databricks_job"].Name(r.Data),
				})
			}
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
			if l, err := jobs.NewJobsAPI(ic.Context, ic.Client).List(); err == nil {
				ic.importJobs(l)
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
				Name:     "clust_policy_" + ic.Importables["databricks_cluster_policy"].Name(r.Data),
			})
			var definition map[string]map[string]interface{}
			err := json.Unmarshal([]byte(r.Data.Get("definition").(string)), &definition)
			if err != nil {
				return err
			}
			for k, policy := range definition {
				value, vok := policy["value"]
				defaultValue, dok := policy["defaultValue"]
				if !vok && !dok {
					continue
				}
				if k == "aws_attributes.instance_profile_arn" {
					ic.Emit(&resource{
						Resource: "databricks_instance_profile",
						ID:       eitherString(value, defaultValue),
					})
				}
				if k == "instance_pool_id" {
					ic.Emit(&resource{
						Resource: "databricks_instance_pool",
						ID:       eitherString(value, defaultValue),
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
					log.Printf("[INFO] Group %s doesn't match %s filter", g.DisplayName, ic.match)
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
				for _, parent := range g.Groups {
					ic.Emit(&resource{
						Resource: "databricks_group",
						ID:       parent.Value,
					})
					if parent.Type == "direct" {
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
			// TODO: if I have 2 users from different domains: test@domain1.com & test@domain2.com - then I'll generate the same name
			// use another algorithm for name generation, like, just replace non-word characters with '_'
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
			username := r.Data.Get("user_name").(string)
			u, err := ic.findUserByName(username)
			if err != nil {
				return err
			}
			for _, g := range u.Groups {
				if g.Type != "direct" {
					log.Printf("Skipping non-direct group %s/%s for user %s", g.Value, g.Display, username)
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
		Ignore: func(ic *importContext, r *resource) bool {
			var permissions permissions.PermissionsEntity
			s := ic.Resources["databricks_permissions"].Schema
			common.DataToStructPointer(r.Data, s, &permissions)
			return (len(permissions.AccessControlList) == 0)
		},
		Import: func(ic *importContext, r *resource) error {
			var permissions permissions.PermissionsEntity
			s := ic.Resources["databricks_permissions"].Schema
			common.DataToStructPointer(r.Data, s, &permissions)
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
			ssAPI := secrets.NewSecretScopesAPI(ic.Context, ic.Client)
			if scopes, err := ssAPI.List(); err == nil {
				for i, scope := range scopes {
					if !ic.MatchesName(scope.Name) {
						log.Printf("[INFO] Secret scope %s doesn't match %s filter", scope.Name, ic.match)
						continue
					}
					ic.Emit(&resource{
						Resource: "databricks_secret_scope",
						ID:       scope.Name,
						Name:     scope.Name,
					})
					log.Printf("[INFO] Imported %d of %d secret scopes",
						i, len(scopes))
				}
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			backendType, _ := r.Data.GetOk("backend_type")
			if backendType != "AZURE_KEYVAULT" {
				if l, err := secrets.NewSecretsAPI(ic.Context, ic.Client).List(r.ID); err == nil {
					for _, secret := range l {
						ic.Emit(&resource{
							Resource: "databricks_secret",
							ID:       fmt.Sprintf("%s|||%s", r.ID, secret.Key),
						})
					}
				}
			}
			if l, err := secrets.NewSecretAclsAPI(ic.Context, ic.Client).List(r.ID); err == nil {
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
	"databricks_mount": {
		Service: "mounts",
		Body:    generateMountBody,
		List: func(ic *importContext) error {
			if !ic.mounts {
				return nil
			}
			if err := ic.refreshMounts(); err != nil {
				return err
			}
			for mountName, source := range ic.mountMap {
				if !ic.MatchesName(mountName) {
					continue
				}
				if strings.HasPrefix(source.URL, "s3a://") {
					log.Printf("[INFO] Emitting databricks_mount: %s", source.URL)
					if source.InstanceProfile != "" {
						ic.Emit(&resource{
							Resource: "databricks_instance_profile",
							ID:       source.InstanceProfile,
						})
					} else if source.ClusterID != "" {
						ic.Emit(&resource{
							Resource: "databricks_cluster",
							ID:       source.ClusterID,
						})
					}
				} else if strings.HasPrefix(source.URL, "gs://") {
					if source.ClusterID != "" {
						ic.Emit(&resource{
							Resource: "databricks_cluster",
							ID:       source.ClusterID,
						})
					}
				} else if res := adlsGen2Regex.FindStringSubmatch(source.URL); res != nil {
				} else if res := adlsGen1Regex.FindStringSubmatch(source.URL); res != nil {
				} else if res := wasbsRegex.FindStringSubmatch(source.URL); res != nil {
				} else {
					log.Printf("[INFO] No matching handler for: %s", source.URL)
					continue
				}
				log.Printf("[INFO] Emitting databricks_mount: %s", source.URL)
				ic.Emit(&resource{
					Resource: "databricks_mount",
					ID:       mountName,
					Data: ic.Resources["databricks_mount"].Data(
						&terraform.InstanceState{
							ID:         mountName,
							Attributes: map[string]string{},
						}),
				})

			}
			return nil
		},
		Depends: []reference{
			{Path: "s3_bucket_name", Resource: "aws_s3_bucket", Match: "bucket"}, // this should be changed somehow & avoid clashes with GCS bucket_name
			{Path: "instance_profile", Resource: "databricks_instance_profile"},
			{Path: "cluster_id", Resource: "databricks_cluster"},
			{Path: "storage_account_name", Resource: "azurerm_storage_account", Match: "name"}, // similarly for WASBS vs ABFSS
			{Path: "container_name", Resource: "azurerm_storage_container", Match: "name"},
			{Path: "storage_resource_name", Resource: "azurerm_data_lake_store", Match: "name"},
		},
	},
	"databricks_global_init_script": {
		Service: "workspace",
		Name: func(d *schema.ResourceData) string {
			name := d.Get("name").(string)
			if name == "" {
				return d.Id()
			}
			re := regexp.MustCompile(`[^0-9A-Za-z_]`)
			return re.ReplaceAllString(name, "_")
		},
		List: func(ic *importContext) error {
			globalInitScripts, err := workspace.NewGlobalInitScriptsAPI(ic.Context, ic.Client).List()
			if err != nil {
				return err
			}
			for offset, gis := range globalInitScripts {
				ic.Emit(&resource{
					Resource: "databricks_global_init_script",
					ID:       gis.ScriptID,
				})
				log.Printf("[INFO] Scanned %d of %d global init scripts", offset+1, len(globalInitScripts))
			}
			return nil
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			gis, err := workspace.NewGlobalInitScriptsAPI(ic.Context, ic.Client).Get(r.ID)
			if err != nil {
				return err
			}
			content, err := base64.StdEncoding.DecodeString(gis.ContentBase64)
			if err != nil {
				return err
			}
			fileName, err := ic.createFile(path.Base(r.Name), content)
			log.Printf("Creating %s for %s", fileName, r)
			if err != nil {
				return err
			}
			relativeFile := fmt.Sprintf("${path.module}/%s", fileName)
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			b.SetAttributeValue("name", cty.StringVal(gis.Name))
			b.SetAttributeValue("enabled", cty.BoolVal(gis.Enabled))
			b.SetAttributeRaw("source", hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}},
				&hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(relativeFile)},
				&hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}},
			})
			return nil
		},
	},
	"databricks_repo": {
		Service: "repos",
		Name: func(d *schema.ResourceData) string {
			name := d.Get("path").(string)
			if name == "" {
				return d.Id()
			} else {
				name = strings.TrimPrefix(name, "/")
			}
			re := regexp.MustCompile(`[^0-9A-Za-z_]`)
			return re.ReplaceAllString(name, "_")
		},
		List: func(ic *importContext) error {
			repoList, err := repos.NewReposAPI(ic.Context, ic.Client).ListAll()
			if err != nil {
				return err
			}
			for offset, repo := range repoList {
				if repo.Url != "" {
					ic.Emit(&resource{
						Resource: "databricks_repo",
						ID:       fmt.Sprintf("%d", repo.ID),
					})
				}
				log.Printf("[INFO] Scanned %d of %d repos", offset+1, len(repoList))
			}
			return nil
		},
		Import: func(ic *importContext, r *resource) error {
			if ic.meAdmin {
				ic.Emit(&resource{
					Resource: "databricks_permissions",
					ID:       fmt.Sprintf("/repos/%s", r.ID),
					Name:     "repo_" + ic.Importables["databricks_repo"].Name(r.Data),
				})
			}
			return nil
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			b.SetAttributeValue("url", cty.StringVal(r.Data.Get("url").(string)))
			b.SetAttributeValue("git_provider", cty.StringVal(r.Data.Get("git_provider").(string)))
			t := r.Data.Get("branch").(string)
			if t != "" {
				b.SetAttributeValue("branch", cty.StringVal(t))
			}
			t = r.Data.Get("path").(string)
			if t != "" {
				b.SetAttributeValue("path", cty.StringVal(t))
			}
			return nil
		},
	},
	"databricks_git_credential": {
		Service: "repos",
		Name: func(d *schema.ResourceData) string {
			return d.Get("git_provider").(string) + "_" + d.Get("git_username").(string) + "_" + d.Id()
		},
		List: func(ic *importContext) error {
			creds, err := repos.NewGitCredentialsAPI(ic.Context, ic.Client).List()
			if err != nil {
				return err
			}
			for offset, cred := range creds {
				ic.Emit(&resource{
					Resource: "databricks_git_credential",
					ID:       fmt.Sprintf("%d", cred.ID),
				})
				log.Printf("[INFO] Scanned %d of %d Git credentials", offset+1, len(creds))
			}
			return nil
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			b.SetAttributeRaw("personal_access_token", ic.variable("git_token_"+r.Name, "Git token for "+r.Name))
			b.SetAttributeValue("git_provider", cty.StringVal(r.Data.Get("git_provider").(string)))
			b.SetAttributeValue("git_username", cty.StringVal(r.Data.Get("git_username").(string)))
			return nil
		},
	},
	"databricks_workspace_conf": {
		Service: "workspace",
		Name: func(d *schema.ResourceData) string {
			return globalWorkspaceConfName
		},
		Import: func(ic *importContext, r *resource) error {
			wsConfAPI := workspace.NewWorkspaceConfAPI(ic.Context, ic.Client)
			keys := map[string]interface{}{"enableIpAccessLists": false, "maxTokenLifetimeDays": 0, "enableTokensConfig": false}
			err := wsConfAPI.Read(&keys)
			if err != nil {
				return err
			}
			r.Data.Set("custom_config", keys)
			return nil
		},
	},
	"databricks_ip_access_list": {
		Service: "access",
		Name: func(d *schema.ResourceData) string {
			return d.Get("list_type").(string) + "_" + d.Get("label").(string)
		},
		List: func(ic *importContext) error {
			ipListsResp, err := access.NewIPAccessListsAPI(ic.Context, ic.Client).List()
			if err != nil {
				return err
			}
			ipLists := ipListsResp.ListIPAccessListsResponse
			for offset, ipList := range ipLists {
				ic.Emit(&resource{
					Resource: "databricks_ip_access_list",
					ID:       ipList.ListID,
				})
				log.Printf("[INFO] Scanned %d of %d IP Access Lists", offset+1, len(ipLists))
			}
			if len(ipLists) > 0 {
				ic.Emit(&resource{
					Resource: "databricks_workspace_conf",
					ID:       globalWorkspaceConfName,
					Data: ic.Resources["databricks_workspace_conf"].Data(
						&terraform.InstanceState{
							ID:         globalWorkspaceConfName,
							Attributes: map[string]string{},
						}),
				})
			}
			return nil
		},
	},
	"databricks_notebook": {
		Service: "notebooks",
		Name: func(d *schema.ResourceData) string {
			name := d.Get("path").(string)
			if name == "" {
				return d.Id()
			} else {
				name = strings.TrimPrefix(name, "/")
			}
			re := regexp.MustCompile(`[^0-9A-Za-z_]`)
			return re.ReplaceAllString(name, "_")
		},
		List: func(ic *importContext) error {
			notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
			notebookList, err := notebooksAPI.List("/", true)
			if err != nil {
				return err
			}
			// TODO: emit permissions for notebook folders if non-default,
			// as per-notebook permission entry would be a noise in the state
			for offset, notebook := range notebookList {
				ic.Emit(&resource{
					Resource: "databricks_notebook",
					ID:       notebook.Path,
				})
				if offset%50 == 0 {
					log.Printf("[INFO] Scanned %d of %d notebooks",
						offset+1, len(notebookList))
				}
			}
			return nil
		},
		Body: func(ic *importContext, body *hclwrite.Body, r *resource) error {
			notebooksAPI := workspace.NewNotebooksAPI(ic.Context, ic.Client)
			status, err := notebooksAPI.Read(r.ID)
			if err != nil {
				return err
			}
			contentB64, err := notebooksAPI.Export(r.ID, "SOURCE")
			if err != nil {
				return err
			}
			name := r.ID[1:] + status.Extension() // todo: replace non-alphanum+/ with _
			content, _ := base64.StdEncoding.DecodeString(contentB64)
			fileName, err := ic.createFileIn("notebooks", name, []byte(content))
			log.Printf("Creating %s for %s", fileName, r)
			if err != nil {
				return err
			}
			// libraries installed with init scripts won't be exported.
			b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			relativeFile := fmt.Sprintf("${path.module}/%s", fileName)
			b.SetAttributeValue("path", cty.StringVal(r.ID))
			b.SetAttributeRaw("source", hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}},
				&hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(relativeFile)},
				&hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}},
			})
			return nil
		},
	},
}
