package exporter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	sdk_compute "github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/environments"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	environments_wbe "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/environments_workspace_base_environment"
	tf_dlt "github.com/databricks/terraform-provider-databricks/pipelines"
	tf_workspace "github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

func listClusters(ic *importContext) error {
	lastActiveMs := ic.getLastActiveMs()
	interactiveClusters := []sdk_compute.ClusterSource{sdk_compute.ClusterSourceUi, sdk_compute.ClusterSourceApi}

	it := ic.workspaceClient.Clusters.List(ic.Context, sdk_compute.ListClustersRequest{
		FilterBy: &sdk_compute.ListClustersFilterBy{
			ClusterSources: interactiveClusters,
		},
		PageSize: 100,
	})
	i := 0
	for it.HasNext(ic.Context) {
		c, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++

		if strings.HasPrefix(c.ClusterName, "terraform-") {
			log.Printf("[INFO] Skipping terraform-specific cluster %s", c.ClusterName)
			continue
		}
		if !ic.MatchesName(c.ClusterName) {
			log.Printf("[INFO] Skipping %s because it doesn't match %s", c.ClusterName, ic.match)
			continue
		}
		if c.LastRestartedTime > 0 && c.LastRestartedTime < lastActiveMs {
			log.Printf("[INFO] Old inactive cluster %s", c.ClusterName)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_cluster",
			ID:       c.ClusterId,
		})
		if i%50 == 0 {
			log.Printf("[INFO] Scanned %d clusters", i)
		}
	}
	return nil
}

func (ic *importContext) importCluster(c *sdk_compute.ClusterSpec) {
	if c == nil {
		return
	}
	if c.AwsAttributes != nil && c.AwsAttributes.InstanceProfileArn != "" {
		ic.Emit(&resource{
			Resource: "databricks_instance_profile",
			ID:       c.AwsAttributes.InstanceProfileArn,
		})
	}
	if c.InstancePoolId != "" {
		// set enable_elastic_disk to false, and remove aws/gcp/azure_attributes
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       c.InstancePoolId,
		})
	}
	if c.DriverInstancePoolId != "" {
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       c.DriverInstancePoolId,
		})
	}
	if c.PolicyId != "" {
		ic.Emit(&resource{
			Resource: "databricks_cluster_policy",
			ID:       c.PolicyId,
		})
	}
	if c.DockerImage != nil && c.DockerImage.BasicAuth != nil {
		ic.emitSecretsFromSecretPathString(c.DockerImage.BasicAuth.Password)
		ic.emitSecretsFromSecretPathString(c.DockerImage.BasicAuth.Username)
	}
	ic.emitInitScripts(c.InitScripts)
	ic.emitSecretsFromSecretsPathMap(c.SparkConf)
	ic.emitSecretsFromSecretsPathMap(c.SparkEnvVars)
	ic.emitUserOrServicePrincipal(c.SingleUserName)
	if c.Kind.String() != "" && c.SingleUserName != "" {
		ic.Emit(&resource{
			Resource:  "databricks_group",
			Attribute: "display_name",
			Value:     c.SingleUserName,
		})
	}
}

func listInstancePools(ic *importContext) error {
	it := ic.workspaceClient.InstancePools.List(ic.Context)
	i := 0
	for it.HasNext(ic.Context) {
		pool, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if !ic.MatchesName(pool.InstancePoolName) {
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       pool.InstancePoolId,
		})
		if i%50 == 0 {
			log.Printf("[INFO] Imported %d instance pools", i)
		}
	}
	return nil
}

func instancePoolName(ic *importContext, d *schema.ResourceData) string {
	raw, ok := d.GetOk("instance_pool_name")
	if !ok || raw.(string) == "" {
		return strings.Split(d.Id(), "-")[2]
	}
	return raw.(string)
}

func importInstancePool(ic *importContext, r *resource) error {
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/instance-pools/%s", r.ID),
		"inst_pool_"+ic.Importables["databricks_instance_pool"].Name(ic, r.Data))
	return nil
}

func importCluster(ic *importContext, r *resource) error {
	var c sdk_compute.ClusterSpec
	s := ic.Resources["databricks_cluster"].Schema
	common.DataToStructPointer(r.Data, s, &c)
	ic.importCluster(&c)
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/clusters/%s", r.ID),
		"cluster_"+ic.Importables["databricks_cluster"].Name(ic, r.Data))
	return ic.importClusterLibraries(r.Data)
}

func listClusterPolicies(ic *importContext) error {
	builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
	it := ic.workspaceClient.ClusterPolicies.List(ic.Context, sdk_compute.ListClusterPoliciesRequest{})
	i := 0
	for it.HasNext(ic.Context) {
		policy, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		family, isBuiltin := builtInClusterPolicies[policy.PolicyFamilyId]
		if policy.PolicyFamilyId != "" && isBuiltin && family.Name == policy.Name &&
			policy.PolicyFamilyDefinitionOverrides == "" {
			log.Printf("[DEBUG] Skipping builtin cluster policy '%s' without overrides", policy.Name)
			continue
		}
		if !ic.MatchesName(policy.Name) {
			log.Printf("[DEBUG] Policy %s doesn't match %s filter", policy.Name, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_cluster_policy",
			ID:       policy.PolicyId,
		})
		if i%10 == 0 {
			log.Printf("[INFO] Scanned %d cluster policies", i)
		}
	}
	return nil
}

// clusterPolicyReferences builds the Depends references for databricks_cluster_policy.
// The policy JSON is stored escaped in a single string attribute (either `definition`
// or `policy_family_definition_overrides`), so references to other objects embedded in
// it are extracted with regexps. The same set of references applies to both fields.
// All definition references set ContinueMatch so every one is evaluated and their
// non-overlapping substitutions are combined; secret and init-script references also
// set MultiMatch because a single definition can embed several of them, each pointing
// to a different object.
func clusterPolicyReferences() []reference {
	var refs []reference
	for _, f := range []string{"definition", "policy_family_definition_overrides"} {
		refs = append(refs,
			reference{Path: f, Resource: "databricks_instance_pool", ContinueMatch: true,
				MatchType: MatchRegexp, Regexp: policyInstancePoolIdRegex},
			reference{Path: f, Resource: "databricks_instance_pool", ContinueMatch: true,
				MatchType: MatchRegexp, Regexp: policyDriverInstancePoolIdRegex},
			reference{Path: f, Resource: "databricks_instance_profile", ContinueMatch: true,
				MatchType: MatchRegexp, Regexp: policyInstanceProfileArnRegex},
			reference{Path: f, Resource: "databricks_secret", Match: "config_reference",
				ContinueMatch: true, MultiMatch: true, MatchType: MatchRegexp, Regexp: policySecretRegex},
			// Fallback for scopes without per-secret resources (e.g. Azure Key Vault
			// backed scopes): substitute only the scope, keeping the key literal. It
			// runs after the databricks_secret reference, so tokens already resolved to
			// a secret are left untouched (their span overlaps and is skipped).
			reference{Path: f, Resource: "databricks_secret_scope",
				ContinueMatch: true, MultiMatch: true, MatchType: MatchRegexp, Regexp: policySecretScopeRegex},
			reference{Path: f, Resource: "databricks_dbfs_file", Match: "dbfs_path",
				ContinueMatch: true, MultiMatch: true, MatchType: MatchRegexp, Regexp: policyInitScriptDbfsRegex},
			reference{Path: f, Resource: "databricks_workspace_file", Match: "workspace_path",
				ContinueMatch: true, MultiMatch: true, MatchType: MatchRegexp, Regexp: policyInitScriptWorkspaceRegex},
			reference{Path: f, Resource: "databricks_workspace_file", Match: "path",
				ContinueMatch: true, MultiMatch: true, MatchType: MatchRegexp, Regexp: policyInitScriptWorkspaceRegex},
			reference{Path: f, Resource: "databricks_file", Match: "path",
				ContinueMatch: true, MultiMatch: true, MatchType: MatchRegexp, Regexp: policyInitScriptVolumesRegex},
		)
	}
	refs = append(refs,
		reference{Path: "libraries.jar", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
		reference{Path: "libraries.jar", Resource: "databricks_file"},
		reference{Path: "libraries.jar", Resource: "databricks_workspace_file", Match: "workspace_path"},
		reference{Path: "libraries.whl", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
		reference{Path: "libraries.whl", Resource: "databricks_file"},
		reference{Path: "libraries.whl", Resource: "databricks_workspace_file", Match: "workspace_path"},
		reference{Path: "libraries.egg", Resource: "databricks_dbfs_file", Match: "dbfs_path"},
		reference{Path: "libraries.egg", Resource: "databricks_workspace_file", Match: "workspace_path"},
		reference{Path: "libraries.whl", Resource: "databricks_repo", Match: "workspace_path",
			MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		reference{Path: "libraries.egg", Resource: "databricks_repo", Match: "workspace_path",
			MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		reference{Path: "libraries.jar", Resource: "databricks_repo", Match: "workspace_path",
			MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
	)
	return refs
}

func importClusterPolicy(ic *importContext, r *resource) error {
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/cluster-policies/%s", r.ID),
		"cluster_policy_"+ic.Importables["databricks_cluster_policy"].Name(ic, r.Data))

	var clusterPolicy sdk_compute.Policy
	s := ic.Resources["databricks_cluster_policy"].Schema
	common.DataToStructPointer(r.Data, s, &clusterPolicy)

	// Determine which field to parse and convert
	var policyDefStr string
	var isOverride bool
	if clusterPolicy.PolicyFamilyId != "" && clusterPolicy.PolicyFamilyDefinitionOverrides != "" {
		policyDefStr = clusterPolicy.PolicyFamilyDefinitionOverrides
		isOverride = true
	} else if clusterPolicy.Definition != "" {
		policyDefStr = clusterPolicy.Definition
		isOverride = false
	}

	var definition map[string]map[string]any
	if policyDefStr != "" {
		err := json.Unmarshal([]byte(policyDefStr), &definition)
		if err != nil {
			return err
		}
	} else {
		// No definition to process
		definition = make(map[string]map[string]any)
	}

	// Convert cloud-specific attributes if targetCloud is set
	if ic.targetCloud != "" && len(definition) > 0 {
		sourceCloud := ic.getSourceCloud()
		if ic.convertClusterPolicyDefinition(definition, sourceCloud, ic.targetCloud) {
			// Re-encode the modified definition
			modifiedDef, err := json.Marshal(definition)
			if err != nil {
				return err
			}
			policyDefStr = string(modifiedDef)

			// Update the appropriate field in the resource data
			if isOverride {
				r.Data.Set("policy_family_definition_overrides", policyDefStr)
			} else {
				r.Data.Set("definition", policyDefStr)
			}
			log.Printf("[INFO] Converted cluster policy '%s' for target cloud '%s'", clusterPolicy.Name, ic.targetCloud)
		}
	}

	for k, policy := range definition {
		value, vok := policy["value"]
		defaultValue, dok := policy["defaultValue"]
		typ := policy["type"]
		if !vok && !dok {
			log.Printf("[DEBUG] Skipping policy element as it doesn't have both value and defaultValue. k='%v', policy='%v'",
				k, policy)
			continue
		}
		if k == "aws_attributes.instance_profile_arn" {
			ic.Emit(&resource{
				Resource: "databricks_instance_profile",
				ID:       eitherString(value, defaultValue),
			})
		}
		if k == "instance_pool_id" || k == "driver_instance_pool_id" {
			ic.Emit(&resource{
				Resource: "databricks_instance_pool",
				ID:       eitherString(value, defaultValue),
			})
		}
		if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
			strings.HasSuffix(k, ".dbfs.destination") {
			ic.emitIfDbfsFile(eitherString(value, defaultValue))
		}
		if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
			strings.HasSuffix(k, ".volumes.destination") {
			ic.emitIfVolumeFile(eitherString(value, defaultValue))
		}
		if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
			strings.HasSuffix(k, ".workspace.destination") {
			ic.emitWorkspaceFileOrRepo(eitherString(value, defaultValue))
		}
		// Secret references may appear in any attribute (spark_conf, spark_env_vars,
		// docker_image credentials, ...) and with any type - notably policy families
		// often store them in `defaultValue` with a non-fixed type - so emit the scope
		// for every `{{secrets/scope/key}}` token regardless of key or type.
		ic.emitSecretScopesFromString(eitherString(value, defaultValue))
	}

	for _, lib := range clusterPolicy.Libraries {
		ic.emitIfDbfsFile(lib.Whl)
		ic.emitIfDbfsFile(lib.Jar)
		ic.emitIfDbfsFile(lib.Egg)
		ic.emitIfWsfsFile(lib.Whl)
		ic.emitIfWsfsFile(lib.Jar)
		ic.emitIfWsfsFile(lib.Egg)
		ic.emitIfVolumeFile(lib.Whl)
		ic.emitIfVolumeFile(lib.Jar)
	}

	policyFamilyId := clusterPolicy.PolicyFamilyId
	if policyFamilyId != "" && clusterPolicy.Definition != "" {
		// we need to set definition to empty value because otherwise it will be put into
		// generated HCL code for data source, and it only supports the `name` attribute
		r.Data.Set("definition", "")
		builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
		v, isBuiltin := builtInClusterPolicies[policyFamilyId]
		if isBuiltin && clusterPolicy.PolicyFamilyDefinitionOverrides == "" && v.Name == clusterPolicy.Name {
			r.Mode = "data"
		}
	}

	return nil
}

func listPipelines(ic *importContext) error {
	it := ic.workspaceClient.Pipelines.ListPipelines(ic.Context, pipelines.ListPipelinesRequest{
		MaxResults: 100,
	})
	i := 0
	for it.HasNext(ic.Context) {
		q, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if !ic.MatchesName(q.Name) {
			continue
		}
		var modifiedAt int64
		if ic.incremental {
			pipeline, err := ic.workspaceClient.Pipelines.Get(ic.Context, pipelines.GetPipelineRequest{
				PipelineId: q.PipelineId,
			})
			if err != nil {
				return err
			}
			modifiedAt = pipeline.LastModified
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_pipeline",
			ID:       q.PipelineId,
		}, modifiedAt, fmt.Sprintf("DLT Pipeline '%s'", q.Name))
		if i%100 == 0 {
			log.Printf("[INFO] Imported %d DLT Pipelines", i)
		}
	}
	log.Printf("[INFO] Listed %d DLT pipelines", i)
	return nil
}

func importPipeline(ic *importContext, r *resource) error {
	var pipeline tf_dlt.Pipeline
	s := ic.Resources["databricks_pipeline"].Schema
	common.DataToStructPointer(r.Data, s, &pipeline)
	if pipeline.Deployment != nil && pipeline.Deployment.Kind == "BUNDLE" {
		log.Printf("[INFO] Skipping processing of DLT Pipeline with ID %s (%s) as deployed with DABs",
			r.ID, pipeline.Name)
		return nil
	}
	if pipeline.Catalog != "" {
		var schemaName string
		if pipeline.Target != "" {
			schemaName = pipeline.Target
		} else if pipeline.Schema != "" {
			schemaName = pipeline.Schema
		}
		if schemaName != "" {
			ic.Emit(&resource{
				Resource: "databricks_schema",
				ID:       pipeline.Catalog + "." + pipeline.Target,
			})
		}
	}
	if pipeline.Deployment == nil || pipeline.Deployment.Kind != "BUNDLE" {
		nbAPI := tf_workspace.NewNotebooksAPI(ic.Context, ic.Client)
		for _, lib := range pipeline.Libraries {
			if lib.Notebook != nil {
				ic.emitNotebookOrRepo(lib.Notebook.Path)
			}
			if lib.File != nil {
				ic.emitWorkspaceFileOrRepo(lib.File.Path)
			}
			if lib.Glob != nil {
				if strings.HasSuffix(lib.Glob.Include, "/**") {
					// Emit all files and notebooks under the directory
					dirPath := strings.TrimSuffix(lib.Glob.Include, "/**")
					ic.emitDirectoryOrRepo(dirPath)
					objects, err := nbAPI.List(dirPath, false, true)
					if err == nil {
						for _, object := range objects {
							switch object.ObjectType {
							case tf_workspace.File:
								ic.emitWorkspaceFileOrRepo(object.Path)
							case tf_workspace.Notebook:
								ic.emitNotebookOrRepo(object.Path)
							}
						}
					} else {
						log.Printf("[WARN] Can't list directory %s for DLT pipeline %s", lib.Glob.Include, pipeline.Name)
					}
				} else {
					// Perform get-status call to check what is the object type
					object, err := nbAPI.GetStatus(lib.Glob.Include, false)
					if err == nil {
						switch object.ObjectType {
						case tf_workspace.File:
							ic.emitWorkspaceFileOrRepo(lib.Glob.Include)
						case tf_workspace.Notebook:
							ic.emitNotebookOrRepo(lib.Glob.Include)
						}
					} else {
						log.Printf("[WARN] Can't get status of %s for DLT pipeline %s", lib.Glob.Include, pipeline.Name)
					}
				}
			}
			ic.emitIfDbfsFile(lib.Jar)
			ic.emitIfDbfsFile(lib.Whl)
		}
		if pipeline.RootPath != "" {
			ic.emitDirectoryOrRepo(pipeline.RootPath)
		}
	}
	// Emit clusters
	for _, cluster := range pipeline.Clusters {
		if cluster.AwsAttributes != nil && cluster.AwsAttributes.InstanceProfileArn != "" {
			ic.Emit(&resource{
				Resource: "databricks_instance_profile",
				ID:       cluster.AwsAttributes.InstanceProfileArn,
			})
		}
		if cluster.InstancePoolId != "" {
			ic.Emit(&resource{
				Resource: "databricks_instance_pool",
				ID:       cluster.InstancePoolId,
			})
		}
		if cluster.DriverInstancePoolId != "" {
			ic.Emit(&resource{
				Resource: "databricks_instance_pool",
				ID:       cluster.DriverInstancePoolId,
			})
		}
		if cluster.PolicyId != "" {
			ic.Emit(&resource{
				Resource: "databricks_cluster_policy",
				ID:       cluster.PolicyId,
			})
		}
		ic.emitInitScripts(cluster.InitScripts)
		ic.emitSecretsFromSecretsPathMap(cluster.SparkConf)
		ic.emitSecretsFromSecretsPathMap(cluster.SparkEnvVars)
	}
	ic.emitFilesFromMap(pipeline.Configuration)
	ic.emitSecretsFromSecretsPathMap(pipeline.Configuration)
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/pipelines/%s", r.ID),
		"pipeline_"+ic.Importables["databricks_pipeline"].Name(ic, r.Data))
	if pipeline.Notifications != nil {
		for _, n := range pipeline.Notifications {
			ic.emitListOfUsers(n.EmailRecipients)
		}
	}
	if pipeline.EventLog != nil {
		var catalogName, schemaName string
		if pipeline.EventLog.Catalog != "" {
			catalogName = pipeline.EventLog.Catalog
		} else {
			catalogName = pipeline.Catalog
		}
		if pipeline.EventLog.Schema != "" {
			schemaName = pipeline.EventLog.Schema
		} else {
			schemaName = pipeline.Schema
		}
		if catalogName != "" && schemaName != "" && pipeline.EventLog.Name != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_table",
				ID:       catalogName + "." + schemaName + "." + pipeline.EventLog.Name,
			})
		}
	}
	if pipeline.RunAs != nil {
		if pipeline.RunAs.UserName != "" {
			ic.Emit(&resource{
				Resource:  "databricks_user",
				Attribute: "user_name",
				Value:     pipeline.RunAs.UserName,
			})
		}
		if pipeline.RunAs.ServicePrincipalName != "" {
			ic.Emit(&resource{
				Resource:  "databricks_service_principal",
				Attribute: "application_id",
				Value:     pipeline.RunAs.ServicePrincipalName,
			})
		}
	}
	if pipeline.Environment != nil {
		for _, dep := range pipeline.Environment.Dependencies {
			emitEnvironmentDependency(ic, dep)
		}
	}
	return nil
}

// databricksProvidedBaseEnvPrefix is the resource name prefix used by
// Databricks-provided (system-managed) workspace base environments, such as
// `workspace-base-environments/databricks_ml_v5`. These are built into every
// workspace and should not be exported.
const databricksProvidedBaseEnvPrefix = "workspace-base-environments/databricks_"

func isDatabricksProvidedBaseEnvironment(name string) bool {
	return strings.HasPrefix(name, databricksProvidedBaseEnvPrefix)
}

// isUserManagedBaseEnvironmentRef reports whether a base environment reference
// points to a user-managed environment. Besides the name check (Databricks-provided
// environments are named `workspace-base-environments/databricks_*`), it looks up
// the referenced environment and requires a `filepath` to be set. Relying on
// `filepath` rather than only the name is more robust, since names can change while
// Databricks-managed environments never define a `filepath`.
func (ic *importContext) isUserManagedBaseEnvironmentRef(ref string) bool {
	if ref == "" || isDatabricksProvidedBaseEnvironment(ref) {
		return false
	}
	env, err := ic.workspaceClient.Environments.GetWorkspaceBaseEnvironment(ic.Context,
		environments.GetWorkspaceBaseEnvironmentRequest{Name: ref})
	if err != nil {
		log.Printf("[WARN] Can't get base environment %s: %s", ref, err)
		return false
	}
	return env.Filepath != ""
}

// hasUserManagedBaseEnvironment reports whether any of the given base environment
// references points to a user-managed environment.
func (ic *importContext) hasUserManagedBaseEnvironment(refs ...string) bool {
	for _, ref := range refs {
		if ic.isUserManagedBaseEnvironmentRef(ref) {
			return true
		}
	}
	return false
}

func listWorkspaceBaseEnvironments(ic *importContext) error {
	it := ic.workspaceClient.Environments.ListWorkspaceBaseEnvironments(ic.Context,
		environments.ListWorkspaceBaseEnvironmentsRequest{})
	for it.HasNext(ic.Context) {
		env, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		if env.Name == "" {
			continue
		}
		// Skip Databricks-provided base environments - they exist in every
		// workspace and aren't managed by users.
		if isDatabricksProvidedBaseEnvironment(env.Name) {
			log.Printf("[INFO] Skipping Databricks-provided base environment %s", env.Name)
			continue
		}
		if !ic.MatchesName(env.DisplayName) {
			log.Printf("[INFO] Skipping base environment %s because it doesn't match %s", env.DisplayName, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_environments_workspace_base_environment",
			ID:       env.Name,
		})
	}
	return nil
}

// importWorkspaceBaseEnvironment emits the WSFS or UC Volumes files referenced by
// the base environment: the `filepath` environment YAML, and the WSFS/Volumes
// libraries or `-r requirements.txt` entries listed in `spec.dependencies`. When
// `filepath` is set, the referenced YAML file is downloaded and parsed, and the
// files referenced by its own `dependencies` are emitted as well.
func importWorkspaceBaseEnvironment(ic *importContext, r *resource) error {
	var env environments.WorkspaceBaseEnvironment
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		environments_wbe.WorkspaceBaseEnvironment{}, &env); err != nil {
		return err
	}
	if env.Filepath != "" {
		ic.emitIfWsfsFile(env.Filepath)
		ic.emitIfVolumeFile(env.Filepath)
		ic.emitBaseEnvironmentFileDependencies(env.Filepath)
	}
	if env.Spec != nil {
		for _, dep := range env.Spec.Dependencies {
			emitEnvironmentDependency(ic, dep)
		}
	}
	return nil
}

// baseEnvironmentFileSpec is the subset of the environment YAML file that we need
// to resolve file references from.
type baseEnvironmentFileSpec struct {
	Dependencies []string `yaml:"dependencies"`
}

// emitBaseEnvironmentFileDependencies downloads the environment YAML referenced by
// `filepath` (from WSFS or UC Volumes), parses its `dependencies`, and emits the
// WSFS/Volumes files referenced there (direct library paths and `-r` requirements
// files) so they are exported as well.
func (ic *importContext) emitBaseEnvironmentFileDependencies(filepath string) {
	content, err := ic.readWorkspaceOrVolumeFile(filepath)
	if err != nil {
		log.Printf("[WARN] Can't read base environment file %s: %s", filepath, err)
		return
	}
	var spec baseEnvironmentFileSpec
	if err := yaml.Unmarshal(content, &spec); err != nil {
		log.Printf("[WARN] Can't parse base environment file %s: %s", filepath, err)
		return
	}
	for _, dep := range spec.Dependencies {
		emitEnvironmentDependency(ic, dep)
	}
}

// readWorkspaceOrVolumeFile downloads the content of a file located either on UC
// Volumes (`/Volumes/...`) or in the workspace file system (WSFS).
func (ic *importContext) readWorkspaceOrVolumeFile(path string) ([]byte, error) {
	if strings.HasPrefix(path, "/Volumes/") {
		resp, err := ic.workspaceClient.Files.DownloadByFilePath(ic.Context, path)
		if err != nil {
			return nil, err
		}
		defer resp.Contents.Close()
		return io.ReadAll(resp.Contents)
	}
	reader, err := ic.workspaceClient.Workspace.Download(ic.Context, maybeStripWorkspacePrefix(path),
		func(q map[string]any) {
			q["format"] = "AUTO"
		})
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

// listDefaultWorkspaceBaseEnvironment reads the singleton default workspace base
// environment and emits it only when it points at a user-managed base environment.
// Defaults that only reference Databricks-provided base environments are skipped -
// those are maintained by Databricks and shouldn't be exported.
func listDefaultWorkspaceBaseEnvironment(ic *importContext) error {
	env, err := ic.workspaceClient.Environments.GetDefaultWorkspaceBaseEnvironment(ic.Context,
		environments.GetDefaultWorkspaceBaseEnvironmentRequest{
			Name: "default-workspace-base-environment",
		})
	if err != nil {
		return err
	}
	if !ic.hasUserManagedBaseEnvironment(env.CpuWorkspaceBaseEnvironment, env.GpuWorkspaceBaseEnvironment) {
		log.Printf("[INFO] Skipping default workspace base environment - it only references Databricks-managed base environments")
		return nil
	}
	ic.Emit(&resource{
		Resource: "databricks_environments_default_workspace_base_environment",
		ID:       "default-workspace-base-environment",
	})
	return nil
}

// ignoreDefaultWorkspaceBaseEnvironment ignores the default when it doesn't
// reference any user-managed base environment (in case it was emitted from
// elsewhere) - i.e. it only points at Databricks-managed environments.
func ignoreDefaultWorkspaceBaseEnvironment(ic *importContext, r *resource) bool {
	if r.DataWrapper == nil {
		return false
	}
	cpu, _ := r.DataWrapper.GetOk("cpu_workspace_base_environment")
	gpu, _ := r.DataWrapper.GetOk("gpu_workspace_base_environment")
	cpuStr, _ := cpu.(string)
	gpuStr, _ := gpu.(string)
	return !ic.hasUserManagedBaseEnvironment(cpuStr, gpuStr)
}

// importDefaultWorkspaceBaseEnvironment emits the user-managed base environments
// referenced by the default so they are exported and referenced by name.
func importDefaultWorkspaceBaseEnvironment(ic *importContext, r *resource) error {
	for _, field := range []string{"cpu_workspace_base_environment", "gpu_workspace_base_environment"} {
		v, ok := r.DataWrapper.GetOk(field)
		if !ok {
			continue
		}
		name, ok := v.(string)
		if !ok || name == "" || isDatabricksProvidedBaseEnvironment(name) {
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_environments_workspace_base_environment",
			ID:       name,
		})
	}
	return nil
}
